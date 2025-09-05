package stamp

import (
	"encoding/json"
	"fmt"
)

// ResponseHelper maneja respuestas y errores
type ResponseHelper struct{}

// NewResponseHelper crea una nueva instancia del helper de respuestas
func NewResponseHelper() *ResponseHelper {
	return &ResponseHelper{}
}

// ToErrorResponse convierte un mensaje en una respuesta de error estándar
func (r *ResponseHelper) ToErrorResponse(message, messageDetail string) *StampResponse {
	errorData := map[string]interface{}{
		"status":        "error",
		"message":       message,
		"messageDetail": messageDetail,
	}

	// Convertir a JSON y luego a StampResponse para mantener consistencia
	jsonData, _ := json.Marshal(errorData)
	var response StampResponse
	json.Unmarshal(jsonData, &response)

	return &response
}

// HandleError convierte errores en respuestas estructuradas
func (r *ResponseHelper) HandleError(err error) *StampResponse {
	return r.ToErrorResponse(err.Error(), fmt.Sprintf("%+v", err))
}

// ValidateResponse verifica que la respuesta sea exitosa
func (r *ResponseHelper) ValidateResponse(resp *StampResponse) error {
	if resp.Status != "success" {
		return fmt.Errorf("timbrado falló con status: %s", resp.Status)
	}
	return nil
}
