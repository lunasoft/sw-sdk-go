package issue

import (
	"os"
	"testing"

	"sw-sdk-golang/swsdk/autenticacion"
	"sw-sdk-golang/swsdk/helpers"
)

func TestIssueV1(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Preparar XML listo para timbrar
	xmlHelper := helpers.NewXMLHelper()
	xmlPath := "test_cfdi_v1.xml"
	defer os.Remove(xmlPath) // Limpiar al final

	err = xmlHelper.PrepararXMLParaTimbrado("cfdi40.xml", xmlPath, "Issue_1")
	if err != nil {
		t.Fatalf("Error al preparar XML para timbrado: %v", err)
	}

	resp, err := IssueV1(client, xmlPath)
	if err != nil {
		t.Fatalf("La emisión v1 falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("La emisión falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Emisión v1 exitosa:")
	t.Logf("   TFD: %d caracteres", len(resp.Data.TFD))
	t.Logf("   TFD (primeros 100 chars): %s", truncateString(resp.Data.TFD, 100))
}

func TestIssueV2(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Preparar XML listo para timbrar
	xmlHelper := helpers.NewXMLHelper()
	xmlPath := "test_cfdi_v2.xml"
	defer os.Remove(xmlPath) // Limpiar al final

	err = xmlHelper.PrepararXMLParaTimbrado("cfdi40.xml", xmlPath, "Issue_2")
	if err != nil {
		t.Fatalf("Error al preparar XML para timbrado: %v", err)
	}

	resp, err := IssueV2(client, xmlPath)
	if err != nil {
		t.Fatalf("La emisión v2 falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("La emisión falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Emisión v2 exitosa:")
	t.Logf("   TFD: %d caracteres", len(resp.Data.TFD))
	t.Logf("   CFDI: %d caracteres", len(resp.Data.CFDI))
	t.Logf("   TFD (primeros 100 chars): %s", truncateString(resp.Data.TFD, 100))
	t.Logf("   CFDI (primeros 100 chars): %s", truncateString(resp.Data.CFDI, 100))
}

func TestIssueV3(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Preparar XML listo para timbrar
	xmlHelper := helpers.NewXMLHelper()
	xmlPath := "test_cfdi_v3.xml"
	defer os.Remove(xmlPath) // Limpiar al final

	err = xmlHelper.PrepararXMLParaTimbrado("cfdi40.xml", xmlPath, "Issue_3")
	if err != nil {
		t.Fatalf("Error al preparar XML para timbrado: %v", err)
	}

	resp, err := IssueV3(client, xmlPath)
	if err != nil {
		t.Fatalf("La emisión v3 falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("La emisión falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Emisión v3 exitosa:")
	t.Logf("   CFDI: %d caracteres", len(resp.Data.CFDI))
	t.Logf("   CFDI (primeros 100 chars): %s", truncateString(resp.Data.CFDI, 100))
}

func TestIssueV4(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	// Preparar XML listo para timbrar
	xmlHelper := helpers.NewXMLHelper()
	xmlPath := "test_cfdi_v4.xml"
	defer os.Remove(xmlPath) // Limpiar al final

	err = xmlHelper.PrepararXMLParaTimbrado("cfdi40.xml", xmlPath, "Issue_4")
	if err != nil {
		t.Fatalf("Error al preparar XML para timbrado: %v", err)
	}

	resp, err := IssueV4(client, xmlPath)
	if err != nil {
		t.Fatalf("La emisión v4 falló: %v", err)
	}

	// Verificar respuesta exitosa
	if resp.Status != "success" {
		t.Fatalf("La emisión falló con status: %s", resp.Status)
	}

	// Mostrar datos de respuesta
	t.Logf("✅ Emisión v4 exitosa:")
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
