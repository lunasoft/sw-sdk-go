# SW SDK Go

SDK oficial de SW (Servicios Web) para Go que permite interactuar con los servicios.
 
**Versión:** 1.1.0  
**Licencia:** MIT  
**Desarrollador:** David Ernesto Reyes Ayala

## Características

- ✅ **Autenticación** - Gestión de tokens de acceso con soporte para token infinito
- ✅ **Emisión (Issue)** - 4 versiones de respuesta (v1, v2, v3, v4)
- ✅ **Timbrado (Stamp)** - 4 versiones de respuesta (v1, v2, v3, v4)
- ✅ **Cancelación** - 4 métodos de cancelación (UUID, CSD, PFX, XML)
- ✅ **Gestión de Usuarios** - Operaciones CRUD completas
- ✅ **Gestión de Balance** - Añadir, consultar y eliminar timbres
- ✅ **Validación CFDI** - Validar archivos XML de CFDI
- ✅ **Consulta de Estatus** - Consultar estatus de CFDI en el SAT
- ✅ **Variables de Entorno** - Configuración segura mediante variables de entorno
- ✅ **Manejo de errores** robusto
- ✅ **Tests** completos para todas las funcionalidades

## Instalación

```bash
go get github.com/lunasoft/sw-sdk-go
```

### Requisitos

- Go 1.21.0 o superior
- Acceso a los servicios de SW (test o producción)
- Credenciales válidas o token infinito

## Configuración

El SDK se configura automáticamente usando **variables de entorno** o valores por defecto.

### 🔧 Variables de Entorno (Recomendado)

Crea un archivo `.env` en la raíz de tu proyecto:

```bash
# URLs de los servicios
SW_BASE_URL=https://services.test.sw.com.mx
SW_API_BASE_URL=https://api.test.sw.com.mx

# Credenciales (opcional si usas token infinito)
SW_USER=tu-usuario@sw.com.mx
SW_PASSWORD=tu-password

# Token infinito (recomendado para producción)
SW_TOKEN=tu-token-infinito-aqui
```

### 📋 Variables Disponibles

| Variable | Descripción | Valor por defecto |
|----------|-------------|-------------------|
| `SW_BASE_URL` | URL base para servicios CFDI | `https://services.test.sw.com.mx` |
| `SW_API_BASE_URL` | URL base para servicios Management | `https://api.test.sw.com.mx` |
| `SW_USER` | Usuario para autenticación | (vacío) |
| `SW_PASSWORD` | Contraseña para autenticación | (vacío) |
| `SW_TOKEN` | Token infinito/fijo | (vacío) |

### 🔄 Configuración Programática

```go
import "sw-sdk-golang/swsdk"

// El SDK carga automáticamente las variables de entorno
config := swsdk.LoadConfig()

// También puedes configurar programáticamente
config.User = "otro-usuario@sw.com.mx"
config.Token = "otro-token"
```

### 🔑 Sistema de Autenticación

El SDK soporta **dos métodos de autenticación**:

#### 1. **Token Infinito** (Recomendado para producción)
```go
import "sw-sdk-golang/swsdk/autenticacion"

// Usar token fijo/infinito
client := autenticacion.SetUserWithToken("tu-token-infinito")
```

#### 2. **Autenticación Tradicional** (Fallback)
```go
import "sw-sdk-golang/swsdk/autenticacion"

// Usar usuario/contraseña (con fallback automático)
client := autenticacion.SetUser() // Usa config.Token si existe, sino usa User/Password
```

**Prioridad de autenticación:**
1. Si `config.Token` está configurado → Usa token infinito
2. Si `config.Token` está vacío → Usa autenticación tradicional

## Inicio Rápido

### Ejemplo básico con token infinito

```go
package main

import (
    "fmt"
    "log"
    "sw-sdk-golang/swsdk/autenticacion"
    "sw-sdk-golang/swsdk/balance"
)

func main() {
    // Configurar cliente con token infinito
    client := autenticacion.SetUserWithToken("tu-token-infinito")
    
    // Consultar balance
    resp, err := balance.GetBalance()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Timbres disponibles: %d\n", resp.Data.StampsBalance)
}
```

### Ejemplo con autenticación tradicional

```go
package main

import (
    "fmt"
    "log"
    "sw-sdk-golang/swsdk/autenticacion"
    "sw-sdk-golang/swsdk/issue"
)

func main() {
    // Configurar cliente
    client := autenticacion.SetUser()
    
    // Autenticar
    token, err := client.Autenticacion()
    if err != nil {
        log.Fatal(err)
    }
    
    // Emitir CFDI
    resp, err := issue.IssueV4(client, "ruta/al/archivo.xml")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("CFDI emitido: %s\n", resp.Data.UUID)
}
```

### URLs y Endpoints

- **BaseURL**: `https://services.test.sw.com.mx` (para CFDI)
- **APIBaseURL**: `https://api.test.sw.com.mx` (para Management)

#### Endpoints disponibles:
- **Autenticación**: `/security/authenticate`
- **Emisión**: `/cfdi33/issue`
- **Timbrado**: `/cfdi33/stamp`
- **Cancelación**: `/cfdi33/cancel`
- **Usuarios**: `/management/v2/api/dealers/users`
- **Balance**: `/management/v2/api/users/balance`
- **Validación**: `/validate/cfdi`

## Uso

### 1. Autenticación

#### Con Token Infinito (Recomendado)
```go
import "sw-sdk-golang/swsdk/autenticacion"

// Usar token fijo/infinito
client := autenticacion.SetUserWithToken("tu-token-infinito")
// No necesitas llamar Autenticacion() - el token ya está listo
```

#### Con Autenticación Tradicional
```go
import "sw-sdk-golang/swsdk/autenticacion"

client := autenticacion.SetUser()
token, err := client.Autenticacion()
if err != nil {
    log.Fatal(err)
}
```

#### Con Fallback Automático
```go
import "sw-sdk-golang/swsdk/autenticacion"

// Si config.Token está configurado, lo usa automáticamente
// Si no, usa autenticación tradicional
client := autenticacion.SetUser()
token, err := client.Autenticacion() // Usa el método apropiado
```

### 2. Emisión de CFDI

```go
import "sw-sdk-golang/swsdk/issue"

// Emisión v1 (solo TFD)
resp, err := issue.IssueV1(client, "ruta/al/archivo.xml")

// Emisión v2 (TFD + CFDI)
resp, err := issue.IssueV2(client, "ruta/al/archivo.xml")

// Emisión v3 (solo CFDI)
resp, err := issue.IssueV3(client, "ruta/al/archivo.xml")

// Emisión v4 (respuesta completa)
resp, err := issue.IssueV4(client, "ruta/al/archivo.xml")
```

### 3. Timbrado de CFDI

```go
import "sw-sdk-golang/swsdk/stamp"

// Timbrado v1 (solo TFD)
resp, err := stamp.StampV1(client, "ruta/al/archivo.xml")

// Timbrado v2 (TFD + CFDI)
resp, err := stamp.StampV2(client, "ruta/al/archivo.xml")

// Timbrado v3 (solo CFDI)
resp, err := stamp.StampV3(client, "ruta/al/archivo.xml")

// Timbrado v4 (respuesta completa)
resp, err := stamp.StampV4(client, "ruta/al/archivo.xml")
```

### 4. Cancelación de CFDI

```go
import "sw-sdk-golang/swsdk/cancelacion"

// Cancelación por UUID
resp, err := cancelacion.CancelacionPorUUID(client, "RFC", "UUID", "motivo")

// Cancelación por CSD
request := &cancelacion.CancelacionCSDRequest{
    UUID:     "uuid-del-cfdi",
    Password: "password-del-certificado",
    RFC:      "RFC-del-emisor",
    Motivo:   "02",
    B64Cer:   "certificado-en-base64",
    B64Key:   "llave-privada-en-base64",
}
resp, err := cancelacion.CancelacionPorCSD(client, request)

// Cancelación por PFX
request := &cancelacion.CancelacionPFXRequest{
    UUID:     "uuid-del-cfdi",
    Password: "password-del-pfx",
    RFC:      "RFC-del-emisor",
    Motivo:   "02",
    B64Pfx:   "archivo-pfx-en-base64",
}
resp, err := cancelacion.CancelacionPorPFX(client, request)

// Cancelación por XML
resp, err := cancelacion.CancelacionPorXML(client, "ruta/al/xml-de-cancelacion.xml")
```

### 5. Gestión de Usuarios

```go
import "sw-sdk-golang/swsdk/usuarios"

// Crear usuario
request := &usuarios.CreateUserRequest{
    Name:              "Nombre del Usuario",
    TaxID:             "RFC123456789",
    Email:             "usuario@ejemplo.com",
    Stamps:            100,
    IsUnlimited:       false,
    Password:          "Password123!",
    NotificationEmail: "notificaciones@ejemplo.com",
    Phone:             "1234567890",
}
resp, err := usuarios.CrearUsuario(client, request)

// Listar usuarios
params := &usuarios.ListUsersParams{
    IsActive: &[]bool{true}[0],
    Page:     &[]int{1}[0],
    PerPage:  &[]int{10}[0],
}
resp, err := usuarios.ListarUsuarios(client, params)

// Actualizar usuario
request := &usuarios.UpdateUserRequest{
    IDUser:            "id-del-usuario",
    NotificationEmail: &[]string{"nuevo@email.com"}[0],
    Phone:             &[]string{"9876543210"}[0],
}
resp, err := usuarios.ActualizarUsuario(client, "id-del-usuario", request)

// Eliminar usuario
err := usuarios.EliminarUsuario(client, "id-del-usuario")
```

### 6. Gestión de Balance (Timbres)

```go
import "sw-sdk-golang/swsdk/balance"

// Añadir timbres a un usuario
resp, err := balance.AddStamps("user-id", 10, "Se abonan 10 timbres")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Total de timbres después del abono: %d\n", resp.Data)

// Consultar balance de timbres
resp, err := balance.GetBalance()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Timbres disponibles: %d\n", resp.Data.StampsBalance)
fmt.Printf("Timbres usados: %d\n", resp.Data.StampsUsed)
fmt.Printf("Timbres asignados: %d\n", resp.Data.StampsAssigned)

// Eliminar timbres de un usuario
resp, err := balance.RemoveStamps("user-id", 5, "Se retiran 5 timbres")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Total de timbres después de la eliminación: %d\n", resp.Data)
```

### 7. Validación CFDI

```go
import "sw-sdk-golang/swsdk/validacion"

// Validar un archivo XML de CFDI
resp, err := validacion.ValidarCFDI("ruta/al/archivo.xml")
if err != nil {
    log.Fatal(err)
}

// Imprimir resultado de forma legible
validacion.PrintValidacionResult(resp)

// Verificar si el CFDI es válido
if validacion.IsValidCFDI(resp) {
    fmt.Println("CFDI es válido")
} else {
    fmt.Println("CFDI tiene errores")
}

// Verificar si hay errores
if validacion.HasErrors(resp) {
    errors := validacion.GetErrors(resp)
    for _, err := range errors {
        fmt.Printf("Error: %s\n", err)
    }
}
```

## Estructura del Proyecto

```
sw-sdk-golang/
├── swsdk/
│   ├── autenticacion/     # Gestión de autenticación
│   ├── issue/            # Emisión de CFDI
│   ├── stamp/            # Timbrado de CFDI
│   ├── cancelacion/      # Cancelación de CFDI
│   ├── usuarios/         # Gestión de usuarios
│   ├── balance/          # Gestión de balance/timbres
│   ├── validacion/       # Validación de CFDI
│   └── config.go         # Configuración del SDK
├── cmd/
│   └── main.go           # Ejemplo de uso
├── extras/               # Archivos de prueba
└── README.md
```

## Testing

Ejecutar todas las pruebas:

```bash
go test ./...
```

Ejecutar pruebas específicas:

```bash
# Solo emisión
go test ./swsdk/issue -v

# Solo timbrado
go test ./swsdk/stamp -v

# Solo cancelación
go test ./swsdk/cancelacion -v

# Solo usuarios
go test ./swsdk/usuarios -v

# Solo balance
go test ./swsdk/balance -v

# Solo validación
go test ./swsdk/validacion -v
```

### Notas sobre las pruebas

- Las pruebas de **emisión** y **timbrado** requieren XML sellado válido
- Las pruebas de **cancelación** requieren certificados válidos
- Las pruebas de **usuarios** tienen algunos tests con `t.Skip()` por restricciones del API
- Las pruebas de **balance** funcionan correctamente y muestran información detallada
- Las pruebas de **validación** funcionan correctamente con archivos XML reales
- Para skipear pruebas problemáticas: `go test -skip`

## Requisitos

- Go 1.19 o superior
- Acceso a los servicios de SW (credenciales válidas)

## Licencia

Este proyecto está bajo la licencia MIT. Ver el archivo LICENSE para más detalles.

## Troubleshooting

### Problemas comunes

**Error de autenticación:**
```go
// Verificar variables de entorno
config := swsdk.LoadConfig()
if config.Token == "" && config.User == "" {
    // Configurar variables de entorno o usar configuración programática
    config.User = "tu-usuario@sw.com.mx"
    config.Password = "tu-password"
    // O configurar token infinito
    config.Token = "tu-token-infinito"
}
```

**Variables de entorno no cargan:**
- Verificar que el archivo `.env` esté en la raíz del proyecto
- Verificar que las variables estén configuradas en el sistema
- Usar configuración programática como fallback

**Error de conexión:**
- Verificar conectividad a internet
- Verificar que las URLs estén correctas
- Verificar que las credenciales sean válidas

**Error de validación CFDI:**
- Verificar que el archivo XML existe
- Verificar que el archivo sea un CFDI válido
- Verificar permisos de lectura del archivo

### Logs y debugging

```go
// Habilitar logs detallados
import "log"

// En caso de error, revisar el mensaje completo
if err != nil {
    log.Printf("Error detallado: %+v", err)
}
```

## Soporte

**Desarrollador:** David Ernesto Reyes Ayala  
**Email:** david.reyes@sw.com.mx  
**GitHub:** [david-reyes/sw-sdk-golang](https://github.com/david-reyes/sw-sdk-golang)

Para soporte técnico o reportar bugs:
1. Revisar la sección de Troubleshooting
2. Crear un issue en GitHub
3. Contactar al desarrollador

## Changelog

### v1.1.0 (2025-01-04)
- 🔑 **NUEVO**: Sistema de Token Infinito
  - Soporte para tokens fijos/infinitos
  - Fallback automático a autenticación tradicional
  - Función `SetUserWithToken()` para tokens fijos
  - Prioridad: Token infinito > Autenticación tradicional
- 🔧 **NUEVO**: Variables de Entorno
  - Configuración segura sin credenciales hardcodeadas
  - Soporte para archivo `.env`
  - Variables: `SW_USER`, `SW_PASSWORD`, `SW_TOKEN`, `SW_BASE_URL`, `SW_API_BASE_URL`
- ✅ Mejoras en configuración
- ✅ Tests actualizados para token infinito
- ✅ Documentación actualizada

### v1.0.0
- ✅ Implementación inicial del SDK
- ✅ Soporte para emisión, timbrado, cancelación y gestión de usuarios
- ✅ Gestión de balance/timbres
- ✅ Validación de CFDI
- ✅ Tests completos
- ✅ Documentación completa
