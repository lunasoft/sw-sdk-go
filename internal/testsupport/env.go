// Package testsupport reúne utilidades solo para las pruebas del SDK.
package testsupport

import (
	"os"
	"time"
	_ "time/tzdata" // zona horaria incluida, no depende del equipo
)

const (
	baseURLPruebas    = "https://services.test.sw.com.mx"
	apiBaseURLPruebas = "https://api.test.sw.com.mx"
	zonaServicio      = "America/Mexico_City"
)

func init() {
	// Siempre contra el ambiente de pruebas
	os.Setenv("SW_BASE_URL", baseURLPruebas)
	os.Setenv("SW_API_BASE_URL", apiBaseURLPruebas)

	// Hora del servicio, solo dentro del proceso de pruebas
	if zona, err := time.LoadLocation(zonaServicio); err == nil {
		time.Local = zona
	}
}
