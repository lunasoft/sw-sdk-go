package balance

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// BalanceInfo representa la información detallada del balance
type BalanceInfo struct {
	IDUserBalance   string          `json:"idUserBalance"`
	IDUser          string          `json:"idUser"`
	StampsBalance   int             `json:"stampsBalance"`
	StampsUsed      int             `json:"stampsUsed"`
	StampsAssigned  int             `json:"stampsAssigned"`
	IsUnlimited     bool            `json:"isUnlimited"`
	ExpirationDate  *string         `json:"expirationDate"`
	LastTransaction LastTransaction `json:"lastTransaction"`
}

// LastTransaction representa la última transacción de timbres
type LastTransaction struct {
	Folio          int     `json:"folio"`
	IDUser         string  `json:"idUser"`
	IDUserReceiver string  `json:"idUserReceiver"`
	NameReceiver   string  `json:"nameReceiver"`
	StampsIn       *int    `json:"stampsIn"`
	StampsOut      *int    `json:"stampsOut"`
	StampsCurrent  int     `json:"stampsCurrent"`
	Comment        *string `json:"comment"`
	Date           string  `json:"date"`
	IsEmailSent    bool    `json:"isEmailSent"`
}

// ProcessBalanceResponse procesa la respuesta de consulta de balance
func ProcessBalanceResponse(resp *http.Response) (*BalanceResponse, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result BalanceResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %v", err)
	}

	return &result, nil
}

// ProcessSimpleResponse procesa la respuesta simple de añadir/eliminar timbres
func ProcessSimpleResponse(resp *http.Response) (*SimpleBalanceResponse, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result SimpleBalanceResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %v", err)
	}

	return &result, nil
}

// ValidateBalanceResponse valida que la respuesta de balance sea válida
func ValidateBalanceResponse(resp *BalanceResponse) error {
	if resp.Status != "success" {
		return fmt.Errorf("respuesta no exitosa: %s", resp.Status)
	}

	if resp.Data.IDUser == "" {
		return fmt.Errorf("ID de usuario no encontrado en la respuesta")
	}

	return nil
}

// ValidateSimpleResponse valida que la respuesta simple sea válida
func ValidateSimpleResponse(resp *SimpleBalanceResponse) error {
	if resp.Status != "success" {
		return fmt.Errorf("respuesta no exitosa: %s", resp.Status)
	}

	return nil
}
