package consultaestatus

import (
	"testing"
)

func TestConsultarEstatus(t *testing.T) {
	t.Log("=== Test: Consultar Estatus CFDI ===")
	t.Skip("Skip: Consulta de estatus puede fallar por restricciones del SAT")

	rfcEmisor := "EKU9003173C9"
	rfcReceptor := "URE180429TM6"
	total := "199.16"
	uuid := "d9d4dd58-9bfe-4338-a207-d39b4a7044c6"

	resp, err := ConsultarEstatus(rfcEmisor, rfcReceptor, total, uuid)
	if err != nil {
		t.Logf("⚠️  Error consultando estatus: %v", err)
		t.Log("   Nota: Este test puede fallar por restricciones del SAT")
		return
	}

	t.Logf("✅ Estatus consultado exitosamente:")
	t.Logf("   RFC Emisor: %s", rfcEmisor)
	t.Logf("   RFC Receptor: %s", rfcReceptor)
	t.Logf("   Total: %s", total)
	t.Logf("   UUID: %s", uuid)

	// Mostrar resultado
	estado := GetEstadoCFDI(resp)
	t.Logf("   Estado: %s", estado)

	// Verificar estado específico
	if IsVigente(resp) {
		t.Log("   ✅ CFDI está VIGENTE")
	} else if IsCancelado(resp) {
		t.Log("   ❌ CFDI está CANCELADO")
	} else if IsNoEncontrado(resp) {
		t.Log("   ⚠️  CFDI NO ENCONTRADO")
	} else {
		t.Log("   ℹ️  Estado desconocido")
	}

	// Mostrar respuesta completa
	respuestaCompleta := GetRespuestaCompleta(resp)
	if respuestaCompleta != "" {
		t.Logf("   Respuesta SAT: %s", respuestaCompleta)
	}
}
