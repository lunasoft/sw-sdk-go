package usuarios

import (
	"encoding/json"
	"fmt"
)

// ErrorResponse representa una respuesta de error
type ErrorResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// ToErrorResponse convierte una respuesta de error a ErrorResponse
func ToErrorResponse(data []byte) *ErrorResponse {
	var errorResp ErrorResponse
	if err := json.Unmarshal(data, &errorResp); err != nil {
		return &ErrorResponse{
			Message: "Error desconocido",
			Status:  "error",
		}
	}
	return &errorResp
}

// HandleError maneja errores de respuesta
func HandleError(statusCode int, responseBody []byte) error {
	errorResp := ToErrorResponse(responseBody)
	return fmt.Errorf("error del servidor (%d): %s", statusCode, errorResp.Message)
}

// ValidateResponse valida que la respuesta sea exitosa
func ValidateResponse(status string) error {
	if status != "success" {
		return fmt.Errorf("respuesta no exitosa: %s", status)
	}
	return nil
}

// ValidateUserData valida que los datos del usuario sean válidos
func ValidateUserData(user *Usuario) error {
	if user.IDUser == "" {
		return fmt.Errorf("ID de usuario requerido")
	}
	if user.Email == "" {
		return fmt.Errorf("email requerido")
	}
	return nil
}
