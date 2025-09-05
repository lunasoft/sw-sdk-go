package issue

import (
	"github.com/lunasoft/sw-sdk-go/swsdk"
	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
)

// IssueResponse representa la respuesta del endpoint de emisión
type IssueResponse struct {
	Data struct {
		// Para v1: solo TFD
		TFD string `json:"tfd,omitempty"`

		// Para v2: TFD y CFDI
		CFDI string `json:"cfdi,omitempty"`

		// Para v3: solo CFDI

		// Para v4: respuesta completa
		CadenaOriginalSAT string `json:"cadenaOriginalSAT,omitempty"`
		NoCertificadoSAT  string `json:"noCertificadoSAT,omitempty"`
		NoCertificadoCFDI string `json:"noCertificadoCFDI,omitempty"`
		UUID              string `json:"uuid,omitempty"`
		SelloSAT          string `json:"selloSAT,omitempty"`
		SelloCFDI         string `json:"selloCFDI,omitempty"`
		FechaTimbrado     string `json:"fechaTimbrado,omitempty"`
		QRCode            string `json:"qrCode,omitempty"`
	} `json:"data,omitempty"`
	Status string `json:"status,omitempty"`
}

// IssueV1 emite CFDI versión 1 (solo TFD)
func IssueV1(client *autenticacion.SWClient, xmlPath string) (*IssueResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendStampRequest(config.BaseURL, client.Token, xmlPath, "v1", "/cfdi33/issue/")
}

// IssueV2 emite CFDI versión 2 (TFD + CFDI)
func IssueV2(client *autenticacion.SWClient, xmlPath string) (*IssueResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendStampRequest(config.BaseURL, client.Token, xmlPath, "v2", "/cfdi33/issue/")
}

// IssueV3 emite CFDI versión 3 (solo CFDI)
func IssueV3(client *autenticacion.SWClient, xmlPath string) (*IssueResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendStampRequest(config.BaseURL, client.Token, xmlPath, "v3", "/cfdi33/issue/")
}

// IssueV4 emite CFDI versión 4 (respuesta completa)
func IssueV4(client *autenticacion.SWClient, xmlPath string) (*IssueResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendStampRequest(config.BaseURL, client.Token, xmlPath, "v4", "/cfdi33/issue/")
}
