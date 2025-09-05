package usuarios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"sw-sdk-golang/swsdk"
)

// RequestHelper maneja las peticiones HTTP para usuarios
type RequestHelper struct{}

// NewRequestHelper crea una nueva instancia de RequestHelper
func NewRequestHelper() *RequestHelper {
	return &RequestHelper{}
}

// SendCreateUserRequest envía petición para crear usuario
func (r *RequestHelper) SendCreateUserRequest(baseURL, token string, request *CreateUserRequest) (*CreateUserResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error al convertir a JSON: %w", err)
	}

	config := swsdk.LoadConfig()
	fullURL := config.APIBaseURL + "/management/v2/api/dealers/users"

	// Crear la petición HTTP
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error al crear petición: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en petición HTTP: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer respuesta: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("error en la respuesta del servidor: %d %s - %s",
			resp.StatusCode, resp.Status, string(responseBody))
	}

	var createUserResp CreateUserResponse
	err = json.Unmarshal(responseBody, &createUserResp)
	if err != nil {
		return nil, fmt.Errorf("error al parsear respuesta JSON: %w", err)
	}

	return &createUserResp, nil
}

// SendListUsersRequest envía petición para listar usuarios
func (r *RequestHelper) SendListUsersRequest(baseURL, token string, params *ListUsersParams) (*ListUsersResponse, error) {
	config := swsdk.LoadConfig()
	fullURL := config.APIBaseURL + "/management/v2/api/dealers/users"

	if params != nil {
		queryParams := url.Values{}
		if params.IsActive != nil {
			queryParams.Add("IsActive", strconv.FormatBool(*params.IsActive))
		}
		if params.Page != nil {
			queryParams.Add("page", strconv.Itoa(*params.Page))
		}
		if params.PerPage != nil {
			queryParams.Add("perPage", strconv.Itoa(*params.PerPage))
		}

		if len(queryParams) > 0 {
			fullURL += "?" + queryParams.Encode()
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear petición: %w", err)
	}

	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en petición HTTP: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer respuesta: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta del servidor: %d %s - %s",
			resp.StatusCode, resp.Status, string(responseBody))
	}

	var listUsersResp ListUsersResponse
	err = json.Unmarshal(responseBody, &listUsersResp)
	if err != nil {
		return nil, fmt.Errorf("error al parsear respuesta JSON: %w", err)
	}

	return &listUsersResp, nil
}

// SendUpdateUserRequest envía petición para actualizar usuario
func (r *RequestHelper) SendUpdateUserRequest(baseURL, token, userID string, request *UpdateUserRequest) (*UpdateUserResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error al convertir a JSON: %w", err)
	}

	config := swsdk.LoadConfig()
	fullURL := config.APIBaseURL + "/management/v2/api/dealers/users/" + userID

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error al crear petición: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en petición HTTP: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer respuesta: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta del servidor: %d %s - %s",
			resp.StatusCode, resp.Status, string(responseBody))
	}

	var updateUserResp UpdateUserResponse
	err = json.Unmarshal(responseBody, &updateUserResp)
	if err != nil {
		return nil, fmt.Errorf("error al parsear respuesta JSON: %w", err)
	}

	return &updateUserResp, nil
}

// SendDeleteUserRequest envía petición para eliminar usuario
func (r *RequestHelper) SendDeleteUserRequest(baseURL, token, userID string) error {
	config := swsdk.LoadConfig()
	fullURL := config.APIBaseURL + "/management/v2/api/dealers/users/" + userID

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return fmt.Errorf("error al crear petición: %w", err)
	}

	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error en petición HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		responseBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error en la respuesta del servidor: %d %s - %s",
			resp.StatusCode, resp.Status, string(responseBody))
	}

	return nil
}
