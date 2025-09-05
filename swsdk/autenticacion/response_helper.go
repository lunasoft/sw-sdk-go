package autenticacion

import (
	"encoding/json"
	"fmt"
)

// ResponseHelper maneja respuestas y errores de autenticación
type ResponseHelper struct{}

// NewResponseHelper crea una nueva instancia del helper de respuestas
func NewResponseHelper() *ResponseHelper {
	return &ResponseHelper{}
}

// ToErrorResponse convierte un mensaje en una respuesta de error estándar
func (r *ResponseHelper) ToErrorResponse(message, messageDetail string) *AuthResponse {
	errorData := map[string]interface{}{
		"status":        "error",
		"message":       message,
		"messageDetail": messageDetail,
	}

	// Convertir a JSON y luego a AuthResponse para mantener consistencia
	jsonData, _ := json.Marshal(errorData)
	var response AuthResponse
	json.Unmarshal(jsonData, &response)

	return &response
}

// HandleError convierte errores en respuestas estructuradas
func (r *ResponseHelper) HandleError(err error) *AuthResponse {
	return r.ToErrorResponse(err.Error(), fmt.Sprintf("%+v", err))
}

// ValidateResponse verifica que la respuesta de autenticación sea exitosa
func (r *ResponseHelper) ValidateResponse(resp *AuthResponse) error {
	if resp.Status != "success" {
		return fmt.Errorf("autenticación falló con status: %s", resp.Status)
	}
	if resp.Data.Token == "" {
		return fmt.Errorf("token no recibido en la respuesta")
	}
	return nil
}

// ParseAuthResponse parsea una respuesta JSON en AuthResponse
func (r *ResponseHelper) ParseAuthResponse(data []byte) (*AuthResponse, error) {
	var authResp AuthResponse
	if err := json.Unmarshal(data, &authResp); err != nil {
		return nil, fmt.Errorf("error al analizar respuesta de autenticación: %w", err)
	}
	return &authResp, nil
}
