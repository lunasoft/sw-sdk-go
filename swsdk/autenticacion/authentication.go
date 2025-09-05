package autenticacion

import (
	"io"
	"net/http"

	"github.com/lunasoft/sw-sdk-go/swsdk"
)

// SWClient represents the main client for SW services
type SWClient struct {
	BaseURL    string
	HTTPClient *http.Client
	User       string
	Password   string
	Token      string
}

// AuthResponse represents the response from the authentication endpoint
type AuthResponse struct {
	Data struct {
		Token     string `json:"token,omitempty"`
		ExpiresIn int64  `json:"expires_in,omitempty"`
		TokenType string `json:"tokeny_type,omitempty"`
	} `json:"data,omitempty"`
	Status string `json:"status,omitempty"`
}

// SetUser crea una nueva instancia del cliente SW
// Si existe un token infinito en la configuración, lo usa directamente
// Si no, usa las credenciales de usuario/contraseña
func SetUser() *SWClient {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()

	client := &SWClient{
		BaseURL:    config.BaseURL,
		HTTPClient: requestHelper.ConfigurarClienteHTTP(config.Timeout),
		User:       config.User,
		Password:   config.Password,
	}

	// Si hay token infinito configurado, usarlo directamente
	if config.Token != "" {
		client.Token = config.Token
	}

	return client
}

// SetUserWithCredentials crea un nuevo cliente SW con credenciales personalizadas
func SetUserWithCredentials(user, password string) *SWClient {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()

	return &SWClient{
		BaseURL:    config.BaseURL,
		HTTPClient: requestHelper.ConfigurarClienteHTTP(config.Timeout),
		User:       user,
		Password:   password,
	}
}

// SetUserWithToken crea un nuevo cliente SW con token infinito/fijo
func SetUserWithToken(token string) *SWClient {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()

	return &SWClient{
		BaseURL:    config.BaseURL,
		HTTPClient: requestHelper.ConfigurarClienteHTTP(config.Timeout),
		User:       config.User,
		Password:   config.Password,
		Token:      token, // Token infinito
	}
}

// Autenticacion realiza la autenticación con los servicios SW y devuelve el token
// Si ya existe un token infinito, lo retorna directamente sin hacer autenticación
func (c *SWClient) Autenticacion() (string, error) {
	// Si ya hay un token infinito, usarlo directamente
	if c.Token != "" {
		return c.Token, nil
	}

	// Si no hay token, proceder con autenticación tradicional
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	responseHelper := NewResponseHelper()

	// Construir y validar URL
	url := config.BaseURL + config.AuthEndpoint
	if err := requestHelper.ValidarURL(url); err != nil {
		return "", err
	}

	// Crear y ejecutar request
	req, err := requestHelper.CrearRequestAutenticacion(url, c.User, c.Password)
	if err != nil {
		return "", err
	}

	resp, err := requestHelper.EjecutarRequest(c.HTTPClient, req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Procesar respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	authResp, err := responseHelper.ParseAuthResponse(body)
	if err != nil {
		return "", err
	}

	if err := responseHelper.ValidateResponse(authResp); err != nil {
		return "", err
	}

	// Almacenar token y retornar
	c.Token = authResp.Data.Token
	return authResp.Data.Token, nil
}

// GetToken devuelve el token actual
func (c *SWClient) GetToken() string {
	return c.Token
}
