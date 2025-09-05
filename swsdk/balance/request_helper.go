package balance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sw-sdk-golang/swsdk"
	"sw-sdk-golang/swsdk/autenticacion"
)

// AddStampsRequest representa la estructura para añadir timbres
type AddStampsRequest struct {
	Stamps  int    `json:"stamps"`
	Comment string `json:"comment"`
}

// RemoveStampsRequest representa la estructura para eliminar timbres
type RemoveStampsRequest struct {
	Stamps  int    `json:"stamps"`
	Comment string `json:"comment"`
}

// BalanceResponse representa la respuesta de consulta de balance
type BalanceResponse struct {
	Data struct {
		IDUserBalance   string  `json:"idUserBalance"`
		IDUser          string  `json:"idUser"`
		StampsBalance   int     `json:"stampsBalance"`
		StampsUsed      int     `json:"stampsUsed"`
		StampsAssigned  int     `json:"stampsAssigned"`
		IsUnlimited     bool    `json:"isUnlimited"`
		ExpirationDate  *string `json:"expirationDate"`
		LastTransaction struct {
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
		} `json:"lastTransaction"`
	} `json:"data"`
	Meta   interface{} `json:"meta"`
	Links  interface{} `json:"links"`
	Status string      `json:"status"`
}

// SimpleBalanceResponse representa la respuesta simple de añadir/eliminar timbres
type SimpleBalanceResponse struct {
	Data   int         `json:"data"`
	Meta   interface{} `json:"meta"`
	Links  interface{} `json:"links"`
	Status string      `json:"status"`
}

// addStampsHelper maneja la lógica de añadir timbres
func addStampsHelper(client *autenticacion.SWClient, userID string, stamps int, comment string) (*SimpleBalanceResponse, error) {
	token, err := client.Autenticacion()
	if err != nil {
		return nil, fmt.Errorf("error en autenticación: %v", err)
	}

	config := swsdk.LoadConfig()
	url := config.APIBaseURL + "/management/v2/api/dealers/users/" + userID + "/stamps"

	request := &AddStampsRequest{
		Stamps:  stamps,
		Comment: comment,
	}

	resp, err := makeRequest("POST", url, request, token)
	if err != nil {
		return nil, fmt.Errorf("error en la petición: %v", err)
	}

	var result SimpleBalanceResponse
	if err := handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// getBalanceHelper maneja la lógica de consultar balance
func getBalanceHelper(client *autenticacion.SWClient) (*BalanceResponse, error) {
	token, err := client.Autenticacion()
	if err != nil {
		return nil, fmt.Errorf("error en autenticación: %v", err)
	}

	config := swsdk.LoadConfig()
	url := config.APIBaseURL + "/management/v2/api/users/balance"

	resp, err := makeRequest("GET", url, nil, token)
	if err != nil {
		return nil, fmt.Errorf("error en la petición: %v", err)
	}

	var result BalanceResponse
	if err := handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// removeStampsHelper maneja la lógica de eliminar timbres
func removeStampsHelper(client *autenticacion.SWClient, userID string, stamps int, comment string) (*SimpleBalanceResponse, error) {
	token, err := client.Autenticacion()
	if err != nil {
		return nil, fmt.Errorf("error en autenticación: %v", err)
	}

	config := swsdk.LoadConfig()
	url := config.APIBaseURL + "/management/v2/api/dealers/users/" + userID + "/stamps"

	request := &RemoveStampsRequest{
		Stamps:  stamps,
		Comment: comment,
	}

	resp, err := makeRequest("DELETE", url, request, token)
	if err != nil {
		return nil, fmt.Errorf("error en la petición: %v", err)
	}

	var result SimpleBalanceResponse
	if err := handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// makeRequest realiza una petición HTTP con autenticación
func makeRequest(method, url string, body interface{}, token string) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error al serializar el cuerpo de la petición: %v", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición: %v", err)
	}

	// Headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Connection", "keep-alive")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Cliente HTTP con timeout
	client := &http.Client{
		Timeout: swsdk.LoadConfig().Timeout,
	}

	return client.Do(req)
}

// handleResponse maneja la respuesta HTTP y extrae el JSON
func handleResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error al leer la respuesta: %v", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("error al deserializar la respuesta: %v", err)
	}

	return nil
}
