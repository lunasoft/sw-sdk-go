package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// XMLHelper proporciona funciones para manipular archivos XML
type XMLHelper struct{}

// NewXMLHelper crea una nueva instancia del helper XML
func NewXMLHelper() *XMLHelper {
	return &XMLHelper{}
}

// PrepararXMLParaTimbrado copia el XML base y lo prepara con fecha y folio únicos
func (h *XMLHelper) PrepararXMLParaTimbrado(nombreArchivoBase, outputPath, folio string) error {
	// Copiar XML desde extras
	err := h.CopiarXMLDeExtras(nombreArchivoBase, outputPath)
	if err != nil {
		return fmt.Errorf("error al copiar XML: %w", err)
	}

	// Leer el archivo XML copiado
	content, err := os.ReadFile(outputPath)
	if err != nil {
		return fmt.Errorf("error al leer el archivo XML: %w", err)
	}

	xmlContent := string(content)

	// Obtener fecha actual en formato ISO con hora
	fechaActual := time.Now().Format("2006-01-02T15:04:05")

	// Actualizar fecha y folio
	xmlContent = h.actualizarAtributoFecha(xmlContent, fechaActual)
	xmlContent = h.actualizarAtributoFolio(xmlContent, folio)

	// Escribir el archivo actualizado
	err = os.WriteFile(outputPath, []byte(xmlContent), 0644)
	if err != nil {
		return fmt.Errorf("error al escribir el archivo XML actualizado: %w", err)
	}

	return nil
}

// actualizarAtributoFecha actualiza el atributo Fecha del CFDI
func (h *XMLHelper) actualizarAtributoFecha(xmlContent, nuevaFecha string) string {
	// Buscar y reemplazar el atributo Fecha
	re := regexp.MustCompile(`Fecha="[^"]*"`)
	return re.ReplaceAllString(xmlContent, fmt.Sprintf(`Fecha="%s"`, nuevaFecha))
}

// actualizarAtributoFolio actualiza el atributo Folio del CFDI
func (h *XMLHelper) actualizarAtributoFolio(xmlContent, nuevoFolio string) string {
	// Buscar y reemplazar el atributo Folio
	re := regexp.MustCompile(`Folio="[^"]*"`)
	return re.ReplaceAllString(xmlContent, fmt.Sprintf(`Folio="%s"`, nuevoFolio))
}

// CopiarXMLDeExtras copia el XML real desde la carpeta extras
func (h *XMLHelper) CopiarXMLDeExtras(nombreArchivo, outputPath string) error {
	// Obtener el directorio de trabajo actual
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error al obtener directorio de trabajo: %w", err)
	}

	// Ajustar para trabajar desde subdirectorios de swsdk
	if filepath.Base(wd) == "helpers" || filepath.Base(wd) == "autenticacion" || filepath.Base(wd) == "issue" || filepath.Base(wd) == "stamp" {
		wd = filepath.Dir(filepath.Dir(wd))
	} else if filepath.Base(wd) == "swsdk" {
		wd = filepath.Dir(wd)
	}

	// Ruta del archivo en extras (ruta absoluta)
	xmlPath := filepath.Join(wd, "extras", nombreArchivo)

	// Verificar que el archivo existe
	if _, err := os.Stat(xmlPath); os.IsNotExist(err) {
		return fmt.Errorf("el archivo %s no existe en la carpeta extras", nombreArchivo)
	}

	// Copiar el archivo
	return h.CopiarXML(xmlPath, outputPath)
}

// CopiarXML copia un archivo XML a una nueva ubicación
func (h *XMLHelper) CopiarXML(origen, destino string) error {
	// Abrir archivo origen
	src, err := os.Open(origen)
	if err != nil {
		return fmt.Errorf("error al abrir archivo origen: %w", err)
	}
	defer src.Close()

	// Crear archivo destino
	dst, err := os.Create(destino)
	if err != nil {
		return fmt.Errorf("error al crear archivo destino: %w", err)
	}
	defer dst.Close()

	// Copiar contenido
	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("error al copiar contenido: %w", err)
	}

	return nil
}
