package swsdk

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config mantiene la configuración para el SDK de SW
type Config struct {
	// URLs de los servicios
	BaseURL    string
	APIBaseURL string

	// Credenciales de autenticación
	User     string
	Password string
	Token    string // Token infinito/fijo (opcional)

	// Configuración del cliente HTTP
	Timeout time.Duration

	// Endpoints específicos
	AuthEndpoint            string
	IssueEndpoint           string
	StampEndpoint           string
	CancelEndpoint          string
	UsersEndpoint           string
	BalanceEndpoint         string
	ValidacionEndpoint      string
	ConsultaEstatusEndpoint string
}

// LoadConfig carga la configuración con valores por defecto y variables de entorno
func LoadConfig() *Config {
	// Cargar archivo .env si existe
	godotenv.Load()

	// Obtener valores de variables de entorno o usar valores por defecto
	baseURL := getEnv("SW_BASE_URL", "https://services.test.sw.com.mx")
	apiBaseURL := getEnv("SW_API_BASE_URL", "https://api.test.sw.com.mx")
	user := getEnv("SW_USER", "")
	password := getEnv("SW_PASSWORD", "")
	token := getEnv("SW_TOKEN", "")

	return &Config{
		BaseURL:    baseURL,
		APIBaseURL: apiBaseURL,
		User:       user,
		Password:   password,
		Token:      token,
		Timeout:    timeoutDesdeEntorno(),

		AuthEndpoint:            "/security/authenticate",
		IssueEndpoint:           "/cfdi33/issue",
		StampEndpoint:           "/cfdi33/stamp",
		CancelEndpoint:          "/cfdi33/cancel",
		UsersEndpoint:           "/management/v2/api/dealers/users",
		BalanceEndpoint:         "/management/v2/api/users/balance",
		ValidacionEndpoint:      "/validate/cfdi",
		ConsultaEstatusEndpoint: "https://pruebacfdiconsultaqr.cloudapp.net/ConsultaCFDIService.svc",
	}
}

// getEnv obtiene una variable de entorno o retorna un valor por defecto
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// timeoutPorDefecto es el tiempo máximo de cada petición HTTP.
const timeoutPorDefecto = 30 * time.Second

// timeoutDesdeEntorno lee SW_TIMEOUT (segundos enteros). Con 0 no hay límite;
// un valor ausente o inválido usa el valor por defecto.
func timeoutDesdeEntorno() time.Duration {
	valor := os.Getenv("SW_TIMEOUT")
	if valor == "" {
		return timeoutPorDefecto
	}
	segundos, err := strconv.Atoi(valor)
	if err != nil || segundos < 0 {
		return timeoutPorDefecto
	}
	return time.Duration(segundos) * time.Second
}

// NewHTTPClient crea un cliente HTTP con el timeout de la configuración.
func NewHTTPClient() *http.Client {
	return &http.Client{Timeout: LoadConfig().Timeout}
}
