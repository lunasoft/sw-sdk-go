package cancelacion

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
)

func TestCancelacionPorUUID(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	rfc := "EKU9003173C9"
	uuid := "FE4E71B0-8959-4FB9-8091-F5AC4FB0FEF8"
	motivo := "02"

	resp, err := CancelacionPorUUID(client, rfc, uuid, motivo)
	if err != nil {
		t.Fatalf("La cancelación por UUID falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("La cancelación falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Cancelación por UUID exitosa:")
	t.Logf("   UUID: %v", resp.Data.UUID)
	t.Logf("   Status: %s", resp.Data.Status)
	t.Logf("   Message: %s", resp.Data.Message)
	t.Logf("   Acuse Cancelación: %d caracteres", len(resp.Data.AcuseCancelacion))
	t.Logf("   Acuse: %d caracteres", len(resp.Data.Acuse))
}

func TestCancelacionPorCSD(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Leer certificado y llave privada en base64
	requestHelper := NewRequestHelper()
	certificado, err := requestHelper.ReadFileAsBase64("certificado.txt")
	if err != nil {
		t.Fatalf("Error al leer certificado: %v", err)
	}

	llavePrivada, err := requestHelper.ReadFileAsBase64("llave_privada.txt")
	if err != nil {
		t.Fatalf("Error al leer llave privada: %v", err)
	}

	request := &CancelacionCSDRequest{
		UUID:     "24b927ff-e5fa-4662-9ba0-9176a4c218b1",
		Password: "12345678a",
		RFC:      "EKU9003173C9",
		Motivo:   "02",
		B64Cer:   certificado,
		B64Key:   llavePrivada,
	}

	resp, err := CancelacionPorCSD(client, request)
	if err != nil {
		t.Fatalf("La cancelación por CSD falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("La cancelación falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Cancelación por CSD exitosa:")
	t.Logf("   UUID: %v", resp.Data.UUID)
	t.Logf("   Status: %s", resp.Data.Status)
	t.Logf("   Message: %s", resp.Data.Message)
	t.Logf("   Acuse Cancelación: %d caracteres", len(resp.Data.AcuseCancelacion))
	t.Logf("   Acuse: %d caracteres", len(resp.Data.Acuse))
}

func TestCancelacionPorPFX(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Leer archivo PFX
	requestHelper := NewRequestHelper()
	pfxData, err := requestHelper.ReadFileAsBase64("pfx.txt")
	if err != nil {
		t.Fatalf("Error al leer archivo PFX: %v", err)
	}

	request := &CancelacionPFXRequest{
		UUID:     "15b0cdf5-7cc6-4f6f-815a-5f101402f185",
		Password: "12345678a",
		RFC:      "EKU9003173C9",
		Motivo:   "02",
		B64Pfx:   pfxData,
	}

	resp, err := CancelacionPorPFX(client, request)
	if err != nil {
		t.Fatalf("La cancelación por PFX falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("La cancelación falló con status: %s", resp.Status)
	}

	// Mostrar resultado
	t.Logf("✅ Cancelación por PFX exitosa:")
	t.Logf("   UUID: %v", resp.Data.UUID)
	t.Logf("   Status: %s", resp.Data.Status)
	t.Logf("   Message: %s", resp.Data.Message)
	t.Logf("   Acuse Cancelación: %d caracteres", len(resp.Data.AcuseCancelacion))
	t.Logf("   Acuse: %d caracteres", len(resp.Data.Acuse))
}

func TestCancelacionPorXML(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Preparar XML de cancelación
	xmlPath := "test_cancelacion.xml"
	defer os.Remove(xmlPath) // Limpiar al final

	// Copiar XML de cancelación desde extras
	err = copyXMLFromExtras("cfdi40Cancelacion.xml", xmlPath)
	if err != nil {
		t.Fatalf("Error al preparar XML de cancelación: %v", err)
	}

	resp, err := CancelacionPorXML(client, xmlPath)
	if err != nil {
		t.Fatalf("La cancelación por XML falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("La cancelación falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Cancelación por XML exitosa:")
	t.Logf("   UUID: %v", resp.Data.UUID)
	t.Logf("   Status: %s", resp.Data.Status)
	t.Logf("   Message: %s", resp.Data.Message)
	t.Logf("   Acuse Cancelación: %d caracteres", len(resp.Data.AcuseCancelacion))
	t.Logf("   Acuse: %d caracteres", len(resp.Data.Acuse))
}

func copyXMLFromExtras(sourceFile, destFile string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error al obtener directorio de trabajo: %w", err)
	}

	if filepath.Base(wd) == "helpers" || filepath.Base(wd) == "autenticacion" ||
		filepath.Base(wd) == "issue" || filepath.Base(wd) == "stamp" ||
		filepath.Base(wd) == "cancelacion" {
		wd = filepath.Dir(filepath.Dir(wd))
	} else if filepath.Base(wd) == "swsdk" {
		wd = filepath.Dir(wd)
	}

	sourcePath := filepath.Join(wd, "extras", sourceFile)

	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("error al leer archivo fuente: %w", err)
	}

	err = os.WriteFile(destFile, content, 0644)
	if err != nil {
		return fmt.Errorf("error al escribir archivo destino: %w", err)
	}

	return nil
}
