package cancelacion

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// RequestHelper maneja todas las peticiones HTTP de cancelación
type RequestHelper struct{}

// NewRequestHelper crea una nueva instancia del helper de requests
func NewRequestHelper() *RequestHelper {
	return &RequestHelper{}
}

// SendCancelacionUUIDRequest envía petición de cancelación por UUID
func (r *RequestHelper) SendCancelacionUUIDRequest(baseURL, token, rfc, uuid, motivo string) (*CancelacionResponse, error) {
	// Construir URL completa
	fullURL := fmt.Sprintf("%s/cfdi33/cancel/%s/%s/%s", baseURL, rfc, uuid, motivo)

	// Crear la petición HTTP
	req, err := http.NewRequest("POST", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear petición: %w", err)
	}

	// Configurar headers
	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Accept", "*/*")

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
	var cancelacionResp CancelacionResponse
	err = json.Unmarshal(responseBody, &cancelacionResp)
	if err != nil {
		return nil, fmt.Errorf("error al parsear respuesta JSON: %w", err)
	}

	return &cancelacionResp, nil
}

// SendCancelacionCSDRequest envía petición de cancelación por CSD
func (r *RequestHelper) SendCancelacionCSDRequest(baseURL, token string, request *CancelacionCSDRequest) (*CancelacionResponse, error) {
	// Convertir a JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error al convertir a JSON: %w", err)
	}

	// Construir URL completa
	fullURL := baseURL + "/cfdi33/cancel/csd"

	// Crear la petición HTTP
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error al crear petición: %w", err)
	}

	// Configurar headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Accept", "*/*")

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
	var cancelacionResp CancelacionResponse
	err = json.Unmarshal(responseBody, &cancelacionResp)
	if err != nil {
		return nil, fmt.Errorf("error al parsear respuesta JSON: %w", err)
	}

	return &cancelacionResp, nil
}

// SendCancelacionPFXRequest envía petición de cancelación por PFX
func (r *RequestHelper) SendCancelacionPFXRequest(baseURL, token string, request *CancelacionPFXRequest) (*CancelacionResponse, error) {
	// Convertir a JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error al convertir a JSON: %w", err)
	}

	// Construir URL completa
	fullURL := baseURL + "/cfdi33/cancel/pfx"

	// Crear la petición HTTP
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error al crear petición: %w", err)
	}

	// Configurar headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Accept", "*/*")

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
	var cancelacionResp CancelacionResponse
	err = json.Unmarshal(responseBody, &cancelacionResp)
	if err != nil {
		return nil, fmt.Errorf("error al parsear respuesta JSON: %w", err)
	}

	return &cancelacionResp, nil
}

// SendCancelacionXMLRequest envía petición de cancelación por XML
func (r *RequestHelper) SendCancelacionXMLRequest(baseURL, token, xmlPath string) (*CancelacionResponse, error) {
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
	fullURL := baseURL + "/cfdi33/cancel/xml"

	// Crear la petición HTTP
	req, err := http.NewRequest("POST", fullURL, &body)
	if err != nil {
		return nil, fmt.Errorf("error al crear petición: %w", err)
	}

	// Configurar headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Accept", "*/*")

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
	var cancelacionResp CancelacionResponse
	err = json.Unmarshal(responseBody, &cancelacionResp)
	if err != nil {
		return nil, fmt.Errorf("error al parsear respuesta JSON: %w", err)
	}

	return &cancelacionResp, nil
}

// ReadFileAsBase64 lee un archivo y lo convierte a base64
// Si el archivo termina en .txt, asume que ya está en base64
func (r *RequestHelper) ReadFileAsBase64(filePath string) (string, error) {
	// Obtener el directorio de trabajo actual
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error al obtener directorio de trabajo: %w", err)
	}

	// Ajustar para trabajar desde subdirectorios de swsdk
	if filepath.Base(wd) == "helpers" || filepath.Base(wd) == "autenticacion" ||
		filepath.Base(wd) == "issue" || filepath.Base(wd) == "stamp" ||
		filepath.Base(wd) == "cancelacion" {
		wd = filepath.Dir(filepath.Dir(wd))
	} else if filepath.Base(wd) == "swsdk" {
		wd = filepath.Dir(wd)
	}

	// Ruta del archivo en extras
	fullPath := filepath.Join(wd, "extras", filePath)

	// Leer el archivo
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("error al leer archivo %s: %w", fullPath, err)
	}

	if filepath.Ext(filePath) == ".txt" {
		return string(content), nil
	}

	return base64.StdEncoding.EncodeToString(content), nil
}
