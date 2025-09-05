package stamp

import (
	"github.com/lunasoft/sw-sdk-go/swsdk"
	"github.com/lunasoft/sw-sdk-go/swsdk/autenticacion"
)

// StampResponse representa la respuesta del endpoint de timbrado
type StampResponse struct {
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

// StampV1 timbra CFDI versión 1 (solo TFD)
func StampV1(client *autenticacion.SWClient, xmlPath string) (*StampResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendStampRequest(config.BaseURL, client.Token, xmlPath, "v1", "/cfdi33/stamp/")
}

// StampV2 timbra CFDI versión 2 (TFD + CFDI)
func StampV2(client *autenticacion.SWClient, xmlPath string) (*StampResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendStampRequest(config.BaseURL, client.Token, xmlPath, "v2", "/cfdi33/stamp/")
}

// StampV3 timbra CFDI versión 3 (solo CFDI)
func StampV3(client *autenticacion.SWClient, xmlPath string) (*StampResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendStampRequest(config.BaseURL, client.Token, xmlPath, "v3", "/cfdi33/stamp/")
}

// StampV4 timbra CFDI versión 4 (respuesta completa)
func StampV4(client *autenticacion.SWClient, xmlPath string) (*StampResponse, error) {
	config := swsdk.LoadConfig()
	requestHelper := NewRequestHelper()
	return requestHelper.SendStampRequest(config.BaseURL, client.Token, xmlPath, "v4", "/cfdi33/stamp/")
}
