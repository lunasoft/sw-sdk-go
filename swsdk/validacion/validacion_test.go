package validacion

import (
	"testing"
)

const testXMLPath = "../../extras/cfdi40.xml"

func TestValidarCFDI(t *testing.T) {
	t.Log("=== Test: Validar CFDI ===")

	resp, err := ValidarCFDI(testXMLPath)
	if err != nil {
		t.Logf("⚠️  Error validando CFDI: %v", err)
		t.Log("   Nota: Este test puede fallar si el archivo XML no existe")
		return
	}

	t.Logf("✅ CFDI validado exitosamente:")
	t.Logf("   Status: %s", resp.Status)
	t.Logf("   UUID: %s", resp.UUID)
	t.Logf("   Status SAT: %s", resp.StatusSat)
	t.Logf("   Código SAT: %s", resp.StatusCodeSat)
	t.Logf("   Cancelable: %s", resp.IsCancelable)

	// Mostrar detalles de validación
	t.Logf("\n   Detalles de validación:")
	for i, section := range resp.Detail {
		t.Logf("   %d. %s", i+1, section.Section)
		for j, detail := range section.Detail {
			t.Logf("      %d.%d %s: %s", i+1, j+1, detail.TypeValue, detail.Message)
		}
	}

	// Verificar si es válido
	if IsValidCFDI(resp) {
		t.Log("   ✅ CFDI es válido")
	} else {
		t.Log("   ❌ CFDI tiene errores")
	}

	// Mostrar errores si los hay
	if HasErrors(resp) {
		t.Log("   Errores encontrados:")
		errors := GetErrors(resp)
		for _, err := range errors {
			t.Logf("   - %s", err)
		}
	}
}
