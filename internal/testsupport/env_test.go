package testsupport

import (
	"testing"
	"time"
)

func TestZonaHorariaDeLasPruebas(t *testing.T) {
	if got := time.Now().Location().String(); got != zonaServicio {
		t.Fatalf("zona de las pruebas = %q, se esperaba %q", got, zonaServicio)
	}
}
