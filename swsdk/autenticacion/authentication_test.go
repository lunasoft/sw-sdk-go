package autenticacion

import (
	"testing"
)

// Espera HTTP 200 y token válido
func TestAuthenticationSuccessWithValidCredentials(t *testing.T) {
	client := SetUser()

	token, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	t.Logf("✅ Autenticación exitosa. Token obtenido: %d caracteres", len(token))
}

// Espera error (HTTP 400 no se obtiene token)
func TestAuthenticationErrorWithInvalidCredentials(t *testing.T) {
	client := SetUserWithCredentials("error@usuario.com", "password_invalida")

	_, err := client.Autenticacion()

	if err != nil {
		t.Logf("✅ La autenticación falló correctamente. Error: %v", err)
	} else {
		t.Error("Se esperaba que la autenticación falle")
	}
}
