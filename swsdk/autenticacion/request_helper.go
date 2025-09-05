package autenticacion

import (
	"fmt"
	"net/http"
	"time"
)

// RequestHelper maneja las peticiones HTTP de autenticación
type RequestHelper struct{}

// NewRequestHelper crea una nueva instancia del helper de requests
func NewRequestHelper() *RequestHelper {
	return &RequestHelper{}
}

// CrearRequestAutenticacion crea una petición HTTP para autenticación
func (r *RequestHelper) CrearRequestAutenticacion(url, user, password string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición: %w", err)
	}

	req.Header.Set("password", password)
	req.Header.Set("user", user)

	return req, nil
}

// ConfigurarClienteHTTP configura un cliente HTTP con timeout
func (r *RequestHelper) ConfigurarClienteHTTP(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

// ValidarURL valida que la URL no esté vacía
func (r *RequestHelper) ValidarURL(url string) error {
	if url == "" {
		return fmt.Errorf("la URL no puede estar vacía")
	}
	return nil
}

// EjecutarRequest ejecuta una petición HTTP y retorna la respuesta
func (r *RequestHelper) EjecutarRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la petición: %w", err)
	}
	return resp, nil
}
