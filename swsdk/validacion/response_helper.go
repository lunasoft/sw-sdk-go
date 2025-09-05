package validacion

import (
	"fmt"
)

// PrintValidacionResult imprime el resultado de la validación de forma legible
func PrintValidacionResult(resp *ValidacionResponse) {
	fmt.Printf("=== Resultado de Validación CFDI ===\n")
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("UUID: %s\n", resp.UUID)
	fmt.Printf("Status SAT: %s\n", resp.StatusSat)
	fmt.Printf("Código SAT: %s\n", resp.StatusCodeSat)
	fmt.Printf("Cancelable: %s\n", resp.IsCancelable)
	
	if resp.StatusCancelation != nil {
		fmt.Printf("Status Cancelación: %s\n", *resp.StatusCancelation)
	}

	fmt.Printf("\n=== Detalles de Validación ===\n")
	for i, section := range resp.Detail {
		fmt.Printf("%d. %s\n", i+1, section.Section)
		for j, detail := range section.Detail {
			fmt.Printf("   %d.%d %s: %s\n", i+1, j+1, detail.TypeValue, detail.Message)
			if detail.MessageDetail != "" {
				fmt.Printf("      Detalle: %s\n", detail.MessageDetail)
			}
		}
		fmt.Println()
	}

	if resp.CadenaOriginalSAT != "" {
		fmt.Printf("=== Cadena Original SAT ===\n")
		fmt.Printf("%s\n", resp.CadenaOriginalSAT)
	}

	if resp.CadenaOriginalComprobante != "" {
		fmt.Printf("\n=== Cadena Original Comprobante ===\n")
		fmt.Printf("%s\n", resp.CadenaOriginalComprobante)
	}
}

// IsValidCFDI verifica si el CFDI es válido
func IsValidCFDI(resp *ValidacionResponse) bool {
	return resp.Status == "success"
}

// HasErrors verifica si hay errores en la validación
func HasErrors(resp *ValidacionResponse) bool {
	for _, section := range resp.Detail {
		for _, detail := range section.Detail {
			if detail.Type == 0 { // Error
				return true
			}
		}
	}
	return false
}

// GetErrors obtiene solo los errores de la validación
func GetErrors(resp *ValidacionResponse) []string {
	var errors []string
	for _, section := range resp.Detail {
		for _, detail := range section.Detail {
			if detail.Type == 0 { // Error
				errors = append(errors, fmt.Sprintf("%s: %s", section.Section, detail.Message))
			}
		}
	}
	return errors
}
