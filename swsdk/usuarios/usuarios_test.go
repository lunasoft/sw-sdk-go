package usuarios

import (
	"testing"

	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
)

func TestCrearUsuario(t *testing.T) {
	t.Skip("Skip: Crear usuario puede fallar por restricciones del API")

	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	request := &CreateUserRequest{
		Name:              "Prueba SDK Golang",
		TaxID:             "XIA190128J61",
		Email:             "correo.example_sdk_golang@sw.com",
		Stamps:            10,
		IsUnlimited:       false,
		Password:          "Swpass1!",
		NotificationEmail: "correo.example_sdk_golang@sw.com",
		Phone:             "0000000000",
	}

	resp, err := CrearUsuario(client, request)
	if err != nil {
		t.Fatalf("La creación de usuario falló: %v", err)
	}

	if resp.Status != "success" {
		t.Fatalf("La creación falló con status: %s", resp.Status)
	}

	t.Logf("✅ Usuario creado exitosamente:")
	t.Logf("   ID: %s", resp.Data.IDUser)
	t.Logf("   Nombre: %s", resp.Data.Name)
	t.Logf("   Email: %s", resp.Data.Email)
	t.Logf("   Stamps: %d", resp.Data.Stamps)
	t.Logf("   IsUnlimited: %t", resp.Data.IsUnlimited)
}

func TestListarUsuarios(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	params := &ListUsersParams{
		IsActive: &[]bool{true}[0],
	}

	resp, err := ListarUsuarios(client, params)
	if err != nil {
		t.Fatalf("La lista de usuarios falló: %v", err)
	}

	if resp.Status != "success" {
		t.Fatalf("La lista falló con status: %s", resp.Status)
	}

	t.Logf("✅ Lista de usuarios obtenida exitosamente:")
	t.Logf("   Total de usuarios: %d", resp.Meta.TotalCount)
	t.Logf("   Página actual: %d", resp.Meta.Page)
	t.Logf("   Total de páginas: %d", resp.Meta.TotalPages)
	t.Logf("   Usuarios en esta página: %d", len(resp.Data))

	for i, user := range resp.Data {
		if i >= 3 {
			break
		}
		t.Logf("   Usuario %d: %s (%s) - %s", i+1, user.Name, user.Email, user.IDUser)
	}
}

func TestActualizarUsuario(t *testing.T) {
	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	userID := "a1b5e93f-022a-4982-ad2e-117f5c59f930"

	request := &UpdateUserRequest{
		Name:              nil,
		TaxID:             nil,
		IsUnlimited:       nil,
		IDUser:            userID,
		NotificationEmail: &[]string{"correo.cambio.sdk@gmail.com"}[0],
		Phone:             &[]string{"1234567890"}[0],
	}

	resp, err := ActualizarUsuario(client, userID, request)
	if err != nil {
		t.Fatalf("La actualización de usuario falló: %v", err)
	}

	if resp.Status != "success" {
		t.Fatalf("La actualización falló con status: %s", resp.Status)
	}

	t.Logf("✅ Usuario actualizado exitosamente:")
	t.Logf("   ID actualizado: %s", resp.Data)
}

func TestEliminarUsuario(t *testing.T) {
	t.Skip("Skip: Eliminar usuario puede fallar por restricciones del API")

	client := autenticacion.SetUser()
	_, err := client.Autenticacion()
	if err != nil {
		t.Fatalf("La autenticación falló: %v", err)
	}

	userID := "35e40184-7aa1-4d60-a375-44f3407a95e1"

	err = EliminarUsuario(client, userID)
	if err != nil {
		if err.Error() == "error en la respuesta del servidor: 500 500 Internal Server Error - {\"message\":\"El usuario no se puede eliminar, tiene saldo de 10\",\"status\":\"error\"}" ||
			err.Error() == "error en la respuesta del servidor: 500 500 Internal Server Error - {\"message\":\"El usuario no se puede eliminar, tiene timbrado ilimitado\",\"status\":\"error\"}" {
			t.Logf("✅ Test de eliminación completado (usuario no se puede eliminar por restricciones):")
			t.Logf("   ID: %s", userID)
			t.Logf("   Motivo: Usuario tiene restricciones (saldo o timbrado ilimitado)")
			return
		}
		t.Fatalf("La eliminación de usuario falló: %v", err)
	}

	t.Logf("✅ Usuario eliminado exitosamente:")
	t.Logf("   ID eliminado: %s", userID)
}
