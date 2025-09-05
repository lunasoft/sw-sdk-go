package consultaestatus

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sw-sdk-golang/swsdk"
)

// ConsultaEstatusResponse representa la respuesta de consulta de estatus
type ConsultaEstatusResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		ConsultaResponse struct {
			ConsultaResult struct {
				XMLName xml.Name `xml:"ConsultaResult"`
				Value   string   `xml:",chardata"`
			} `xml:"ConsultaResult"`
		} `xml:"ConsultaResponse"`
	} `xml:"Body"`
}

// ConsultarEstatusHelper maneja la lógica de consulta de estatus
func ConsultarEstatusHelper(rfcEmisor, rfcReceptor, total, uuid string) (*ConsultaEstatusResponse, error) {
	config := swsdk.LoadConfig()
	url := config.ConsultaEstatusEndpoint

	// Crear el SOAP envelope
	soapEnvelope := fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
    <soapenv:Header/>
    <soapenv:Body>
        <tem:Consulta>
            <tem:expresionImpresa>
                <![CDATA[?re=%s&rr=%s&tt=%s&id=%s]]>
            </tem:expresionImpresa>
        </tem:Consulta>
    </soapenv:Body>
</soapenv:Envelope>`, rfcEmisor, rfcReceptor, total, uuid)

	// Crear la petición
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(soapEnvelope))
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición: %v", err)
	}

	// Headers SOAP
	req.Header.Set("Accept", "text/xml")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "text/xml;charset=\"utf-8\"")
	req.Header.Set("SOAPAction", "http://tempuri.org/IConsultaCFDIService/Consulta")

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

	// Parsear la respuesta XML
	var result ConsultaEstatusResponse
	if err := xml.Unmarshal(responseBody, &result); err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta XML: %v", err)
	}

	return &result, nil
}
