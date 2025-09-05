package validacion

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"github.com/lunasoft/sw-sdk-go/swsdk"
	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
)

// ValidacionResponse representa la respuesta de validación CFDI
type ValidacionResponse struct {
	Status string `json:"status"`
	Detail []struct {
		Detail []struct {
			Message       string `json:"message"`
			MessageDetail string `json:"messageDetail"`
			Type          int    `json:"type"`
			TypeValue     string `json:"typeValue"`
		} `json:"detail"`
		Section string `json:"section"`
	} `json:"detail"`
	CadenaOriginalSAT         string  `json:"cadenaOriginalSAT"`
	CadenaOriginalComprobante string  `json:"cadenaOriginalComprobante"`
	UUID                      string  `json:"uuid"`
	StatusSat                 string  `json:"statusSat"`
	StatusCodeSat             string  `json:"statusCodeSat"`
	IsCancelable              string  `json:"isCancelable"`
	StatusCancelation         *string `json:"statusCancelation"`
}

// validarCFDIHelper maneja la lógica de validación CFDI
func validarCFDIHelper(client *autenticacion.SWClient, xmlPath string) (*ValidacionResponse, error) {
	token, err := client.Autenticacion()
	if err != nil {
		return nil, fmt.Errorf("error en autenticación: %v", err)
	}

	config := swsdk.LoadConfig()
	url := config.BaseURL + config.ValidacionEndpoint

	// Crear el multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Abrir el archivo XML
	file, err := os.Open(xmlPath)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo XML: %v", err)
	}
	defer file.Close()

	// Crear el campo del archivo
	part, err := writer.CreateFormFile("xml", filepath.Base(xmlPath))
	if err != nil {
		return nil, fmt.Errorf("error al crear el campo del archivo: %v", err)
	}

	// Copiar el contenido del archivo
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("error al copiar el contenido del archivo: %v", err)
	}

	// Cerrar el writer
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error al cerrar el writer: %v", err)
	}

	// Crear la petición
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición: %v", err)
	}

	// Headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Cliente HTTP con timeout
	httpClient := &http.Client{
		Timeout: config.Timeout,
	}

	// Realizar la petición
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en la petición: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error HTTP %d: %s", resp.StatusCode, string(responseBody))
	}

	// Verificar si la respuesta está comprimida con gzip
	if resp.Header.Get("Content-Encoding") == "gzip" {
		// Descomprimir gzip
		reader, err := gzip.NewReader(bytes.NewReader(responseBody))
		if err != nil {
			return nil, fmt.Errorf("error al crear reader gzip: %v", err)
		}
		defer reader.Close()

		responseBody, err = io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("error al descomprimir gzip: %v", err)
		}
	}

	// Parsear la respuesta JSON
	var result ValidacionResponse
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %v", err)
	}

	return &result, nil
}
