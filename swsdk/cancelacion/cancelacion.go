package cancelacion

import (
	"github.com/lunasoft/sw-sdk-go/swsdk"
	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
)

// CancelacionResponse representa la respuesta del endpoint de cancelación
type CancelacionResponse struct {
	Data struct {
		AcuseCancelacion string      `json:"acuseCancelacion,omitempty"`
		Acuse            string      `json:"acuse,omitempty"`
		UUID             interface{} `json:"uuid,omitempty"`
		Status           string      `json:"status,omitempty"`
		Message          string      `json:"message,omitempty"`
	} `json:"data,omitempty"`
	Status string `json:"status,omitempty"`
}

// CancelacionUUIDRequest representa la petición para cancelación por UUID
type CancelacionUUIDRequest struct {
	RFC    string `json:"rfc"`
	UUID   string `json:"uuid"`
	Motivo string `json:"motivo"`
}

// CancelacionCSDRequest representa la petición para cancelación por CSD
type CancelacionCSDRequest struct {
	UUID     string `json:"uuid"`
	Password string `json:"password"`
	RFC      string `json:"rfc"`
	Motivo   string `json:"motivo"`
	B64Cer   string `json:"b64Cer"`
	B64Key   string `json:"b64Key"`
}

// CancelacionPFXRequest representa la petición para cancelación por PFX
type CancelacionPFXRequest struct {
	UUID     string `json:"uuid"`
	Password string `json:"password"`
	RFC      string `json:"rfc"`
	Motivo   string `json:"motivo"`
	B64Pfx   string `json:"b64Pfx"`
}

func CancelacionPorUUID(client *autenticacion.SWClient, rfc, uuid, motivo string) (*CancelacionResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendCancelacionUUIDRequest(config.BaseURL, client.Token, rfc, uuid, motivo)
}

func CancelacionPorCSD(client *autenticacion.SWClient, request *CancelacionCSDRequest) (*CancelacionResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendCancelacionCSDRequest(config.BaseURL, client.Token, request)
}

func CancelacionPorPFX(client *autenticacion.SWClient, request *CancelacionPFXRequest) (*CancelacionResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendCancelacionPFXRequest(config.BaseURL, client.Token, request)
}

func CancelacionPorXML(client *autenticacion.SWClient, xmlPath string) (*CancelacionResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendCancelacionXMLRequest(config.BaseURL, client.Token, xmlPath)
}
