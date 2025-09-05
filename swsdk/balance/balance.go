package balance

import (
	"sw-sdk-golang/swsdk/autenticacion"
)

// AddStamps añade timbres a un usuario específico
func AddStamps(userID string, stamps int, comment string) (*SimpleBalanceResponse, error) {
	client := autenticacion.SetUser()
	return addStampsHelper(client, userID, stamps, comment)
}

// GetBalance consulta el balance de timbres del usuario autenticado
func GetBalance() (*BalanceResponse, error) {
	client := autenticacion.SetUser()
	return getBalanceHelper(client)
}

// RemoveStamps elimina timbres de un usuario específico
func RemoveStamps(userID string, stamps int, comment string) (*SimpleBalanceResponse, error) {
	client := autenticacion.SetUser()
	return removeStampsHelper(client, userID, stamps, comment)
}
