package testsupport

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// El sello generado debe verificarse con el certificado de extras/.
func TestSellarCFDI40FirmaVerificaConCertificado(t *testing.T) {
	extras, err := directorioExtras()
	if err != nil {
		t.Fatal(err)
	}
	original, err := os.ReadFile(filepath.Join(extras, "cfdi40.xml"))
	if err != nil {
		t.Fatal(err)
	}
	copia := filepath.Join(t.TempDir(), "cfdi40.xml")
	if err := os.WriteFile(copia, original, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := SellarCFDI40(copia); err != nil {
		t.Fatal(err)
	}

	sellado, err := os.ReadFile(copia)
	if err != nil {
		t.Fatal(err)
	}
	coincidencia := regexp.MustCompile(`Sello="([^"]*)"`).FindSubmatch(sellado)
	firma, err := base64.StdEncoding.DecodeString(string(coincidencia[1]))
	if err != nil {
		t.Fatal(err)
	}

	cadena, err := cadenaOriginal40(sellado)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(cadena, "||4.0|") || !strings.HasSuffix(cadena, "||") {
		t.Fatalf("formato de cadena original inesperado: %.60s", cadena)
	}

	crt, err := os.ReadFile(filepath.Join(extras, "certificado.txt"))
	if err != nil {
		t.Fatal(err)
	}
	der, err := base64.StdEncoding.DecodeString(strings.Join(strings.Fields(string(crt)), ""))
	if err != nil {
		t.Fatal(err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	resumen := sha256.Sum256([]byte(cadena))
	if err := rsa.VerifyPKCS1v15(cert.PublicKey.(*rsa.PublicKey), crypto.SHA256, resumen[:], firma); err != nil {
		t.Fatalf("la firma no verifica con el certificado de prueba: %v", err)
	}
}
