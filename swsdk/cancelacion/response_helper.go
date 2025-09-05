package cancelacion

import (
	"encoding/json"
	"fmt"
)

// ResponseHelper maneja respuestas y errores de cancelación
type ResponseHelper struct{}

// NewResponseHelper crea una nueva instancia del helper de respuestas
func NewResponseHelper() *ResponseHelper {
	return &ResponseHelper{}
}

// ToErrorResponse convierte un mensaje en una respuesta de error estándar
func (r *ResponseHelper) ToErrorResponse(message, messageDetail string) *CancelacionResponse {
	errorData := map[string]interface{}{
		"status":        "error",
		"message":       message,
		"messageDetail": messageDetail,
	}

	// Convertir a JSON y luego a CancelacionResponse para mantener consistencia
	jsonData, _ := json.Marshal(errorData)
	var response CancelacionResponse
	json.Unmarshal(jsonData, &response)

	return &response
}

// HandleError convierte errores en respuestas estructuradas
func (r *ResponseHelper) HandleError(err error) *CancelacionResponse {
	return r.ToErrorResponse(err.Error(), fmt.Sprintf("%+v", err))
}

// ValidateResponse verifica que la respuesta sea exitosa
func (r *ResponseHelper) ValidateResponse(resp *CancelacionResponse) error {
	if resp.Status != "success" {
		return fmt.Errorf("cancelación falló con status: %s", resp.Status)
	}
	return nil
}

// ValidateCancelacionData verifica que los datos de cancelación sean válidos
func (r *ResponseHelper) ValidateCancelacionData(resp *CancelacionResponse) error {
	if resp.Data.UUID == nil {
		return fmt.Errorf("UUID no encontrado en la respuesta")
	}
	if resp.Data.AcuseCancelacion == "" {
		return fmt.Errorf("acuse de cancelación no encontrado en la respuesta")
	}
	return nil
}
