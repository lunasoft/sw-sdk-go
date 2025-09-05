package consultaestatus

import (
	"fmt"
	"strings"
)

// PrintConsultaResult imprime el resultado de la consulta de forma legible
func PrintConsultaResult(resp *ConsultaEstatusResponse) {
	fmt.Printf("=== Resultado de Consulta de Estatus ===\n")

	// Extraer información del resultado
	result := resp.Body.ConsultaResponse.ConsultaResult.Value
	fmt.Printf("Respuesta SAT: %s\n", result)

	// Parsear la respuesta para extraer información específica
	if strings.Contains(result, "Vigente") {
		fmt.Printf("✅ CFDI VIGENTE\n")
	} else if strings.Contains(result, "Cancelado") {
		fmt.Printf("❌ CFDI CANCELADO\n")
	} else if strings.Contains(result, "No Encontrado") {
		fmt.Printf("⚠️  CFDI NO ENCONTRADO\n")
	} else {
		fmt.Printf("ℹ️  Estado desconocido\n")
	}
}

// IsVigente verifica si el CFDI está vigente
func IsVigente(resp *ConsultaEstatusResponse) bool {
	result := resp.Body.ConsultaResponse.ConsultaResult.Value
	return strings.Contains(result, "Vigente")
}

// IsCancelado verifica si el CFDI está cancelado
func IsCancelado(resp *ConsultaEstatusResponse) bool {
	result := resp.Body.ConsultaResponse.ConsultaResult.Value
	return strings.Contains(result, "Cancelado")
}

// IsNoEncontrado verifica si el CFDI no fue encontrado
func IsNoEncontrado(resp *ConsultaEstatusResponse) bool {
	result := resp.Body.ConsultaResponse.ConsultaResult.Value
	return strings.Contains(result, "No Encontrado")
}

// GetEstadoCFDI obtiene el estado del CFDI de forma legible
func GetEstadoCFDI(resp *ConsultaEstatusResponse) string {
	result := resp.Body.ConsultaResponse.ConsultaResult.Value

	if strings.Contains(result, "Vigente") {
		return "Vigente"
	} else if strings.Contains(result, "Cancelado") {
		return "Cancelado"
	} else if strings.Contains(result, "No Encontrado") {
		return "No Encontrado"
	}

	return "Desconocido"
}

// GetRespuestaCompleta obtiene la respuesta completa del SAT
func GetRespuestaCompleta(resp *ConsultaEstatusResponse) string {
	return resp.Body.ConsultaResponse.ConsultaResult.Value
}
