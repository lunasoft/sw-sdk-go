package stamp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
)

// Estas pruebas usan un servidor simulado: cubren el envío del timbrado sin
// depender del servicio ni de un XML sellado.

const xmlSimulado = `<?xml version="1.0" encoding="UTF-8"?><cfdi:Comprobante Version="4.0" />`

func escribirXML(t *testing.T) string {
	t.Helper()
	ruta := filepath.Join(t.TempDir(), "cfdi.xml")
	if err := os.WriteFile(ruta, []byte(xmlSimulado), 0o644); err != nil {
		t.Fatal(err)
	}
	return ruta
}

func TestStampEnviaElXMLAlServicio(t *testing.T) {
	casos := []struct {
		version string
		timbrar func(*autenticacion.SWClient, string) (*StampResponse, error)
	}{
		{"v1", StampV1},
		{"v2", StampV2},
		{"v3", StampV3},
		{"v4", StampV4},
	}

	for _, c := range casos {
		t.Run(c.version, func(t *testing.T) {
			servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("método = %s, se esperaba POST", r.Method)
				}
				if r.URL.Path != "/cfdi33/stamp/"+c.version {
					t.Errorf("ruta = %s, se esperaba /cfdi33/stamp/%s", r.URL.Path, c.version)
				}
				if got := r.Header.Get("Authorization"); got != "bearer token-prueba" {
					t.Errorf("Authorization = %q", got)
				}

				archivo, _, err := r.FormFile("xml")
				if err != nil {
					t.Errorf("no llegó el campo xml: %v", err)
				} else {
					contenido, _ := io.ReadAll(archivo)
					if string(contenido) != xmlSimulado {
						t.Errorf("el XML recibido no coincide con el enviado")
					}
				}

				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"status":"success","data":{"tfd":"TFD-PRUEBA","uuid":"UUID-PRUEBA"}}`))
			}))
			defer servidor.Close()
			t.Setenv("SW_BASE_URL", servidor.URL)

			resp, err := c.timbrar(autenticacion.SetUserWithToken("token-prueba"), escribirXML(t))
			if err != nil {
				t.Fatalf("el timbrado falló: %v", err)
			}
			if resp.Status != "success" || resp.Data.TFD != "TFD-PRUEBA" || resp.Data.UUID != "UUID-PRUEBA" {
				t.Fatalf("respuesta inesperada: %+v", resp)
			}
		})
	}
}

func TestStampDevuelveElErrorDelServicio(t *testing.T) {
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"CFDI40102 - rechazado","status":"error"}`))
	}))
	defer servidor.Close()
	t.Setenv("SW_BASE_URL", servidor.URL)

	_, err := StampV1(autenticacion.SetUserWithToken("token-prueba"), escribirXML(t))
	if err == nil {
		t.Fatal("se esperaba un error")
	}
	if !strings.Contains(err.Error(), "400") || !strings.Contains(err.Error(), "CFDI40102") {
		t.Fatalf("el error no incluye el estado y el mensaje del servicio: %v", err)
	}
}

func TestStampRespetaElTimeout(t *testing.T) {
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Leer el cuerpo para que el servidor detecte cuándo el cliente corta la conexión
		io.Copy(io.Discard, r.Body)
		select {
		case <-r.Context().Done():
		case <-time.After(10 * time.Second):
		}
	}))
	defer servidor.Close()
	t.Setenv("SW_BASE_URL", servidor.URL)
	t.Setenv("SW_TIMEOUT", "1")

	inicio := time.Now()
	_, err := StampV1(autenticacion.SetUserWithToken("token-prueba"), escribirXML(t))
	if err == nil {
		t.Fatal("se esperaba un error por timeout")
	}
	if !strings.Contains(err.Error(), "Client.Timeout") {
		t.Fatalf("el error no es de timeout: %v", err)
	}
	if time.Since(inicio) > 5*time.Second {
		t.Fatalf("el timeout tardó %v, se esperaba cerca de 1 s", time.Since(inicio))
	}
}
