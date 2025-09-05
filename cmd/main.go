package main

import (
	"fmt"

	"github.com/lunasoft/sw-sdk-go/swsdk"
	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
)

func main() {
	fmt.Println("=== SW SDK Go ===")

	fmt.Println("1. Configurando cliente SW...")
	config := swsdk.LoadConfig()

	fmt.Printf("   Usuario: %s\n", config.User)
	fmt.Printf("   Contraseña: %s\n", config.Password)
	fmt.Printf("   URL Base: %s\n", config.BaseURL)

	fmt.Println("\n2. Configurando cliente de autenticación...")

	// Opción 1: Usar token infinito (recomendado)
	fmt.Println("   Opción 1: Usando token infinito...")
	_ = autenticacion.SetUserWithToken("tu_token_infinito_aqui")
	fmt.Printf("   Cliente configurado con token infinito\n")

	fmt.Println("   Opción 2: Usando autenticación tradicional...")
	clientFallback := autenticacion.SetUser()
	fmt.Printf("   Cliente configurado para usuario: %s\n", clientFallback.User)

	_ = clientFallback

	fmt.Println("   Nota: Si config.Token está vacío, usará autenticación tradicional")

	fmt.Println("\n3. Ejemplo de uso de Issue (emisión)...")
	fmt.Println("   - IssueV1: Solo TFD")
	fmt.Println("   - IssueV2: TFD + CFDI")
	fmt.Println("   - IssueV3: Solo CFDI")
	fmt.Println("   - IssueV4: Respuesta completa")
	fmt.Println("   Uso: issue.IssueV1(client, \"ruta/al/archivo.xml\")")

	fmt.Println("\n4. Ejemplo de uso de Stamp (timbrado)...")
	fmt.Println("   - StampV1: Solo TFD")
	fmt.Println("   - StampV2: TFD + CFDI")
	fmt.Println("   - StampV3: Solo CFDI")
	fmt.Println("   - StampV4: Respuesta completa")
	fmt.Println("   Uso: stamp.StampV1(client, \"ruta/al/archivo.xml\")")

	fmt.Println("\n5. Ejemplo de uso de Cancelación...")
	fmt.Println("   - CancelacionPorUUID: Cancelación simple por UUID")
	fmt.Println("   - CancelacionPorCSD: Cancelación con certificado y llave")
	fmt.Println("   - CancelacionPorPFX: Cancelación con archivo PFX")
	fmt.Println("   - CancelacionPorXML: Cancelación con XML de cancelación")
	fmt.Println("   Uso: cancelacion.CancelacionPorUUID(client, \"RFC\", \"UUID\", \"motivo\")")

	fmt.Println("\n6. Ejemplo de uso de Usuarios (Management)...")
	fmt.Println("   - CrearUsuario: Crear nuevo usuario")
	fmt.Println("   - ListarUsuarios: Obtener lista de usuarios")
	fmt.Println("   - ActualizarUsuario: Actualizar usuario existente")
	fmt.Println("   - EliminarUsuario: Eliminar usuario")
	fmt.Println("   Uso: usuarios.CrearUsuario(client, &CreateUserRequest{...})")

	fmt.Println("\n7. Ejemplo de uso de Balance (Gestión de Timbres)...")
	fmt.Println("   - AddStamps: Añadir timbres a un usuario")
	fmt.Println("   - GetBalance: Consultar balance de timbres")
	fmt.Println("   - RemoveStamps: Eliminar timbres de un usuario")
	fmt.Println("   Uso: balance.AddStamps(\"userID\", 10, \"Comentario\")")

	fmt.Println("\n8. Ejemplo de uso de Validación CFDI...")
	fmt.Println("   - ValidarCFDI: Validar un archivo XML de CFDI")
	fmt.Println("   - PrintValidacionResult: Imprimir resultado de validación")
	fmt.Println("   - IsValidCFDI: Verificar si el CFDI es válido")
	fmt.Println("   - HasErrors: Verificar si hay errores")
	fmt.Println("   Uso: validacion.ValidarCFDI(\"ruta/al/archivo.xml\")")

	fmt.Println("\n9. Ejemplo de uso de Consulta de Estatus...")
	fmt.Println("   - ConsultarEstatus: Consultar estatus de CFDI en el SAT")
	fmt.Println("   - PrintConsultaResult: Imprimir resultado de consulta")
	fmt.Println("   - IsVigente: Verificar si el CFDI está vigente")
	fmt.Println("   - IsCancelado: Verificar si el CFDI está cancelado")
	fmt.Println("   - GetEstadoCFDI: Obtener estado del CFDI")
	fmt.Println("   Uso: consultaestatus.ConsultarEstatus(\"RFC\", \"RFC\", \"total\", \"uuid\")")

	fmt.Println("\n10. Notas importantes:")
	fmt.Println("   - Las pruebas pueden fallar porque requieren XML sellado")
	fmt.Println("   - Para skipear pruebas: go test -skip")
	fmt.Println("   - El endpoint de stamp es: /cfdi33/stamp/")
	fmt.Println("   - El endpoint de issue es: /cfdi33/issue/")
	fmt.Println("   - El endpoint de cancelación es: /cfdi33/cancel/")
	fmt.Println("   - El endpoint de usuarios es: /management/v2/api/dealers/users")
	fmt.Println("   - El endpoint de balance es: /management/v2/api/users/balance")
	fmt.Println("   - El endpoint de validación es: /validate/cfdi")
	fmt.Println("   - El endpoint de consulta estatus es: https://pruebacfdiconsultaqr.cloudapp.net/ConsultaCFDIService.svc")

	fmt.Println("\n=== Fin del ejemplo ===")
}
