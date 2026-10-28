package stamp

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/lunasoft/sw-sdk-go/internal/testsupport"
	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
	"github.com/lunasoft/sw-sdk-go/swsdk/helpers"
	"github.com/lunasoft/sw-sdk-go/swsdk/issue"
)

// Emite un CFDI y lo envía a timbrar de nuevo: el servicio responde 307
// (timbre previo). Cubre el envío del timbrado sin necesitar un XML sellado.
func TestStampComprobanteTimbrado(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	dir := t.TempDir()
	xmlPath := filepath.Join(dir, "cfdi_emitir.xml")
	if err := helpers.NewXMLHelper().PrepararXMLParaTimbrado("cfdi40.xml", xmlPath, "Stamp_307"); err != nil {
		t.Fatalf("Error al preparar XML para timbrado: %v", err)
	}

	emitido, err := issue.IssueV2(client, xmlPath)
	if err != nil {
		t.Fatalf("La emisión falló: %v", err)
	}
	timbradoPath := filepath.Join(dir, "cfdi_timbrado.xml")
	if err := os.WriteFile(timbradoPath, []byte(emitido.Data.CFDI), 0o644); err != nil {
		t.Fatal(err)
	}

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
			_, err := c.timbrar(client, timbradoPath)
			if err == nil {
				t.Fatal("se esperaba el error 307 de timbre previo")
			}
			if !strings.Contains(err.Error(), "307") {
				t.Fatalf("se esperaba el error 307, se obtuvo: %v", err)
			}
		})
	}
}
