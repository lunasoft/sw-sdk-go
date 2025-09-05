package stamp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// RequestHelper maneja todas las peticiones HTTP
type RequestHelper struct{}

// NewRequestHelper crea una nueva instancia del helper de requests
func NewRequestHelper() *RequestHelper {
	return &RequestHelper{}
}

// SendStampRequest envía una petición de timbrado con XML
func (r *RequestHelper) SendStampRequest(baseURL, token, xmlPath, version, endpoint string) (*StampResponse, error) {
	// Preparar el cuerpo multipart
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Agregar el archivo XML
	file, err := os.Open(xmlPath)
	if err != nil {
		return nil, fmt.Errorf("error al abrir archivo XML: %w", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("xml", filepath.Base(xmlPath))
	if err != nil {
		return nil, fmt.Errorf("error al crear form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("error al copiar archivo: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error al cerrar writer: %w", err)
	}

	// Construir URL completa
	fullURL := fmt.Sprintf("%s%s%s", baseURL, endpoint, version)

	// Crear la petición HTTP
	req, err := http.NewRequest("POST", fullURL, &body)
	if err != nil {
		return nil, fmt.Errorf("error al crear petición: %w", err)
	}

	// Configurar headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "bearer "+token)

	// Ejecutar la petición
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en petición HTTP: %w", err)
	}
	defer resp.Body.Close()

	// Leer respuesta
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer respuesta: %w", err)
	}

	// Verificar código de estado HTTP
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta del servidor: %d %s - %s",
			resp.StatusCode, resp.Status, string(responseBody))
	}

	// Parsear respuesta JSON
	var stampResp StampResponse
	err = json.Unmarshal(responseBody, &stampResp)
	if err != nil {
		return nil, fmt.Errorf("error al parsear respuesta JSON: %w", err)
	}

	return &stampResp, nil
}
