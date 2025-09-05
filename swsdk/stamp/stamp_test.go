package stamp

import (
	"os"
	"testing"

	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
	"github.com/lunasoft/sw-sdk-go/swsdk/helpers"
)

func TestStampV1(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Preparar XML listo para timbrar
	xmlHelper := helpers.NewXMLHelper()
	xmlPath := "test_cfdi_stamp_v1.xml"
	defer os.Remove(xmlPath) // Limpiar al final

	err = xmlHelper.PrepararXMLParaTimbrado("cfdi40.xml", xmlPath, "Stamp_1")
	if err != nil {
		t.Fatalf("Error al preparar XML para timbrado: %v", err)
	}

	resp, err := StampV1(client, xmlPath)
	if err != nil {
		t.Fatalf("El timbrado v1 falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("El timbrado falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Timbrado v1 exitoso:")
	t.Logf("   TFD: %d caracteres", len(resp.Data.TFD))
	t.Logf("   TFD (primeros 100 chars): %s", truncateString(resp.Data.TFD, 100))
}

func TestStampV2(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Preparar XML listo para timbrar
	xmlHelper := helpers.NewXMLHelper()
	xmlPath := "test_cfdi_stamp_v2.xml"
	defer os.Remove(xmlPath) // Limpiar al final

	err = xmlHelper.PrepararXMLParaTimbrado("cfdi40.xml", xmlPath, "Stamp_2")
	if err != nil {
		t.Fatalf("Error al preparar XML para timbrado: %v", err)
	}

	resp, err := StampV2(client, xmlPath)
	if err != nil {
		t.Fatalf("El timbrado v2 falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("El timbrado falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Timbrado v2 exitoso:")
	t.Logf("   TFD: %d caracteres", len(resp.Data.TFD))
	t.Logf("   CFDI: %d caracteres", len(resp.Data.CFDI))
	t.Logf("   TFD (primeros 100 chars): %s", truncateString(resp.Data.TFD, 100))
	t.Logf("   CFDI (primeros 100 chars): %s", truncateString(resp.Data.CFDI, 100))
}

func TestStampV3(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Preparar XML listo para timbrar
	xmlHelper := helpers.NewXMLHelper()
	xmlPath := "test_cfdi_stamp_v3.xml"
	defer os.Remove(xmlPath) // Limpiar al final

	err = xmlHelper.PrepararXMLParaTimbrado("cfdi40.xml", xmlPath, "Stamp_3")
	if err != nil {
		t.Fatalf("Error al preparar XML para timbrado: %v", err)
	}

	resp, err := StampV3(client, xmlPath)
	if err != nil {
		t.Fatalf("El timbrado v3 falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("El timbrado falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Timbrado v3 exitoso:")
	t.Logf("   CFDI: %d caracteres", len(resp.Data.CFDI))
	t.Logf("   CFDI (primeros 100 chars): %s", truncateString(resp.Data.CFDI, 100))
}

func TestStampV4(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Preparar XML listo para timbrar
	xmlHelper := helpers.NewXMLHelper()
	xmlPath := "test_cfdi_stamp_v4.xml"
	defer os.Remove(xmlPath) // Limpiar al final

	err = xmlHelper.PrepararXMLParaTimbrado("cfdi40.xml", xmlPath, "Stamp_4")
	if err != nil {
		t.Fatalf("Error al preparar XML para timbrado: %v", err)
	}

	resp, err := StampV4(client, xmlPath)
	if err != nil {
		t.Fatalf("El timbrado v4 falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("El timbrado falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Timbrado v4 exitoso:")
	t.Logf("   UUID: %s", resp.Data.UUID)
	t.Logf("   Fecha Timbrado: %s", resp.Data.FechaTimbrado)
	t.Logf("   Cadena Original SAT: %s", resp.Data.CadenaOriginalSAT)
	t.Logf("   No Certificado SAT: %s", resp.Data.NoCertificadoSAT)
	t.Logf("   No Certificado CFDI: %s", resp.Data.NoCertificadoCFDI)
	t.Logf("   Sello SAT: %s", resp.Data.SelloSAT)
	t.Logf("   Sello CFDI: %s", resp.Data.SelloCFDI)
	t.Logf("   QR Code: %s", resp.Data.QRCode)
}

// Función auxiliar para truncar strings largos
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
