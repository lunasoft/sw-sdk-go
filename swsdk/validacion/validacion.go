package validacion

import (
	"sw-sdk-golang/swsdk/autenticacion"
)

// ValidarCFDI valida un archivo XML de CFDI
func ValidarCFDI(xmlPath string) (*ValidacionResponse, error) {
	client := autenticacion.SetUser()
	return validarCFDIHelper(client, xmlPath)
}
