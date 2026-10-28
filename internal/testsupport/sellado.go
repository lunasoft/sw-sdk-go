package testsupport

import (
	"crypto"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Contraseña pública del CSD de pruebas del SAT incluido en extras/.
const PasswordCSDPrueba = "12345678a"

var (
	oidPBES2  = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 13}
	oidPBKDF2 = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 12}
	oid3DES   = asn1.ObjectIdentifier{1, 2, 840, 113549, 3, 7}
)

// SellarCFDI40 vuelve a sellar el XML (CFDI 4.0) con la llave del CSD de
// pruebas de extras/. El timbrado exige un sello válido y, al cambiar Fecha y
// Folio, el original deja de serlo.
func SellarCFDI40(xmlPath string) error {
	extras, err := directorioExtras()
	if err != nil {
		return err
	}
	clave, err := cargarLlave(filepath.Join(extras, "llave_privada.txt"), PasswordCSDPrueba)
	if err != nil {
		return fmt.Errorf("no se pudo abrir la llave de prueba: %w", err)
	}

	contenido, err := os.ReadFile(xmlPath)
	if err != nil {
		return err
	}
	cadena, err := cadenaOriginal40(contenido)
	if err != nil {
		return fmt.Errorf("no se pudo generar la cadena original: %w", err)
	}
	sello, err := firmar(clave, cadena)
	if err != nil {
		return err
	}

	reSello := regexp.MustCompile(`Sello="[^"]*"`)
	if !reSello.Match(contenido) {
		return errors.New("el XML no tiene atributo Sello")
	}
	actualizado := reSello.ReplaceAll(contenido, []byte(`Sello="`+sello+`"`))
	return os.WriteFile(xmlPath, actualizado, 0o644)
}

// directorioExtras sube desde el directorio de trabajo hasta encontrar go.mod.
func directorioExtras() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return filepath.Join(dir, "extras"), nil
		}
		padre := filepath.Dir(dir)
		if padre == dir {
			return "", errors.New("no se encontró go.mod")
		}
		dir = padre
	}
}

// cargarLlave descifra una llave PKCS#8 (PBES2, PBKDF2-SHA1, 3DES-CBC) en base64.
func cargarLlave(ruta, password string) (*rsa.PrivateKey, error) {
	crudo, err := os.ReadFile(ruta)
	if err != nil {
		return nil, err
	}
	der, err := base64.StdEncoding.DecodeString(strings.Join(strings.Fields(string(crudo)), ""))
	if err != nil {
		return nil, err
	}

	var info struct {
		Algoritmo struct {
			OID    asn1.ObjectIdentifier
			Params asn1.RawValue
		}
		Datos []byte
	}
	if _, err := asn1.Unmarshal(der, &info); err != nil {
		return nil, err
	}
	if !info.Algoritmo.OID.Equal(oidPBES2) {
		return nil, errors.New("la llave no usa PBES2")
	}

	var pbes2 struct {
		KDF struct {
			OID    asn1.ObjectIdentifier
			Params asn1.RawValue
		}
		Cifrado struct {
			OID asn1.ObjectIdentifier
			IV  []byte
		}
	}
	if _, err := asn1.Unmarshal(info.Algoritmo.Params.FullBytes, &pbes2); err != nil {
		return nil, err
	}
	if !pbes2.KDF.OID.Equal(oidPBKDF2) || !pbes2.Cifrado.OID.Equal(oid3DES) {
		return nil, errors.New("algoritmo de llave no soportado en las pruebas")
	}

	var kdf struct {
		Sal         []byte
		Iteraciones int
	}
	if _, err := asn1.Unmarshal(pbes2.KDF.Params.FullBytes, &kdf); err != nil {
		return nil, err
	}

	derivada := pbkdf2SHA1([]byte(password), kdf.Sal, kdf.Iteraciones, 24)
	bloque, err := des.NewTripleDESCipher(derivada)
	if err != nil {
		return nil, err
	}
	if len(info.Datos)%bloque.BlockSize() != 0 {
		return nil, errors.New("datos de llave con longitud inválida")
	}
	plano := make([]byte, len(info.Datos))
	cipher.NewCBCDecrypter(bloque, pbes2.Cifrado.IV).CryptBlocks(plano, info.Datos)
	if n := int(plano[len(plano)-1]); n > 0 && n <= bloque.BlockSize() {
		plano = plano[:len(plano)-n]
	}

	clave, err := x509.ParsePKCS8PrivateKey(plano)
	if err != nil {
		return nil, fmt.Errorf("contraseña incorrecta o llave dañada: %w", err)
	}
	rsaKey, ok := clave.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("la llave no es RSA")
	}
	return rsaKey, nil
}

// pbkdf2SHA1 implementa PBKDF2 con HMAC-SHA1 (crypto/pbkdf2 exige Go 1.24).
func pbkdf2SHA1(password, sal []byte, iteraciones, longitud int) []byte {
	prf := hmac.New(sha1.New, password)
	hLen := prf.Size()
	bloques := (longitud + hLen - 1) / hLen
	resultado := make([]byte, 0, bloques*hLen)
	for i := 1; i <= bloques; i++ {
		prf.Reset()
		prf.Write(sal)
		prf.Write([]byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)})
		u := prf.Sum(nil)
		t := append([]byte(nil), u...)
		for j := 1; j < iteraciones; j++ {
			prf.Reset()
			prf.Write(u)
			u = prf.Sum(nil)
			for k := range t {
				t[k] ^= u[k]
			}
		}
		resultado = append(resultado, t...)
	}
	return resultado[:longitud]
}

func firmar(clave *rsa.PrivateKey, cadena string) (string, error) {
	resumen := sha256.Sum256([]byte(cadena))
	firma, err := rsa.SignPKCS1v15(rand.Reader, clave, crypto.SHA256, resumen[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(firma), nil
}

// ---- Cadena original CFDI 4.0 (solo los nodos que usan los XML de prueba) ----

type nodo struct {
	nombre   string
	attrs    map[string]string
	hijos    []*nodo
	esperaEn string // nombre del padre
}

func (n *nodo) hijo(nombre string) *nodo {
	for _, h := range n.hijos {
		if h.nombre == nombre {
			return h
		}
	}
	return nil
}

func (n *nodo) hijosNombrados(nombre string) []*nodo {
	var res []*nodo
	for _, h := range n.hijos {
		if h.nombre == nombre {
			res = append(res, h)
		}
	}
	return res
}

func parsear(contenido []byte) (*nodo, error) {
	dec := xml.NewDecoder(strings.NewReader(string(contenido)))
	var raiz *nodo
	var pila []*nodo
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}
		switch t := tok.(type) {
		case xml.StartElement:
			n := &nodo{nombre: t.Name.Local, attrs: map[string]string{}}
			for _, a := range t.Attr {
				n.attrs[a.Name.Local] = a.Value
			}
			if len(pila) > 0 {
				p := pila[len(pila)-1]
				n.esperaEn = p.nombre
				p.hijos = append(p.hijos, n)
			} else {
				raiz = n
			}
			pila = append(pila, n)
		case xml.EndElement:
			pila = pila[:len(pila)-1]
		}
	}
	if raiz == nil {
		return nil, errors.New("XML vacío o inválido")
	}
	return raiz, nil
}

// valores devuelve, en el orden dado, los atributos presentes (los opcionales
// ausentes no se incluyen, como en la cadena original del SAT).
func (n *nodo) valores(orden ...string) []string {
	var res []string
	for _, a := range orden {
		if v, ok := n.attrs[a]; ok {
			res = append(res, v)
		}
	}
	return res
}

func cadenaOriginal40(contenido []byte) (string, error) {
	c, err := parsear(contenido)
	if err != nil {
		return "", err
	}
	if c.nombre != "Comprobante" {
		return "", errors.New("la raíz no es cfdi:Comprobante")
	}

	var p []string
	p = append(p, c.valores("Version", "Serie", "Folio", "Fecha", "FormaPago", "NoCertificado",
		"CondicionesDePago", "SubTotal", "Descuento", "Moneda", "TipoCambio", "Total",
		"TipoDeComprobante", "Exportacion", "MetodoPago", "LugarExpedicion", "Confirmacion")...)

	if n := c.hijo("InformacionGlobal"); n != nil {
		p = append(p, n.valores("Periodicidad", "Meses", "Año")...)
	}
	if n := c.hijo("CfdiRelacionados"); n != nil {
		p = append(p, n.valores("TipoRelacion")...)
		for _, r := range n.hijosNombrados("CfdiRelacionado") {
			p = append(p, r.valores("UUID")...)
		}
	}
	if n := c.hijo("Emisor"); n != nil {
		p = append(p, n.valores("Rfc", "Nombre", "RegimenFiscal", "FacAtrAdquirente")...)
	}
	if n := c.hijo("Receptor"); n != nil {
		p = append(p, n.valores("Rfc", "Nombre", "DomicilioFiscalReceptor", "ResidenciaFiscal",
			"NumRegIdTrib", "RegimenFiscalReceptor", "UsoCFDI")...)
	}
	if conceptos := c.hijo("Conceptos"); conceptos != nil {
		for _, con := range conceptos.hijosNombrados("Concepto") {
			p = append(p, con.valores("ClaveProdServ", "NoIdentificacion", "Cantidad", "ClaveUnidad",
				"Unidad", "Descripcion", "ValorUnitario", "Importe", "Descuento", "ObjetoImp")...)
			if imp := con.hijo("Impuestos"); imp != nil {
				if t := imp.hijo("Traslados"); t != nil {
					for _, x := range t.hijosNombrados("Traslado") {
						p = append(p, x.valores("Base", "Impuesto", "TipoFactor", "TasaOCuota", "Importe")...)
					}
				}
				if r := imp.hijo("Retenciones"); r != nil {
					for _, x := range r.hijosNombrados("Retencion") {
						p = append(p, x.valores("Base", "Impuesto", "TipoFactor", "TasaOCuota", "Importe")...)
					}
				}
			}
		}
	}
	if imp := c.hijo("Impuestos"); imp != nil {
		if r := imp.hijo("Retenciones"); r != nil {
			for _, x := range r.hijosNombrados("Retencion") {
				p = append(p, x.valores("Impuesto", "Importe")...)
			}
		}
		p = append(p, imp.valores("TotalImpuestosRetenidos")...)
		if t := imp.hijo("Traslados"); t != nil {
			for _, x := range t.hijosNombrados("Traslado") {
				p = append(p, x.valores("Base", "Impuesto", "TipoFactor", "TasaOCuota", "Importe")...)
			}
		}
		p = append(p, imp.valores("TotalImpuestosTrasladados")...)
	}
	return "||" + strings.Join(p, "|") + "||", nil
}
