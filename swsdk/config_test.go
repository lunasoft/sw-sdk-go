package swsdk

import (
	"testing"
	"time"
)

func TestTimeoutDesdeEntorno(t *testing.T) {
	casos := []struct {
		nombre string
		valor  string
		quiere time.Duration
	}{
		{"sin definir usa el valor por defecto", "", 30 * time.Second},
		{"segundos válidos", "90", 90 * time.Second},
		{"cero significa sin límite", "0", 0},
		{"valor inválido usa el valor por defecto", "abc", 30 * time.Second},
		{"negativo usa el valor por defecto", "-5", 30 * time.Second},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			t.Setenv("SW_TIMEOUT", c.valor)
			if got := LoadConfig().Timeout; got != c.quiere {
				t.Fatalf("SW_TIMEOUT=%q: timeout %v, se esperaba %v", c.valor, got, c.quiere)
			}
			if got := NewHTTPClient().Timeout; got != c.quiere {
				t.Fatalf("NewHTTPClient: timeout %v, se esperaba %v", got, c.quiere)
			}
		})
	}
}
