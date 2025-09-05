package usuarios

import (
	"sw-sdk-golang/swsdk/autenticacion"
)

// Usuario representa un usuario del sistema
type Usuario struct {
	IDUser             string  `json:"idUser"`
	IDDealer           string  `json:"idDealer"`
	Name               string  `json:"name"`
	TaxID              string  `json:"taxId"`
	Username           string  `json:"username"`
	LastPasswordChange string  `json:"lastPasswordChange"`
	Email              string  `json:"email"`
	NotificationEmail  *string `json:"notificationEmail"`
	Profile            int     `json:"profile"`
	IsActive           bool    `json:"isActive"`
	RegisteredDate     string  `json:"registeredDate"`
	AccessToken        string  `json:"accessToken"`
	Phone              string  `json:"phone"`
	Stamps             int     `json:"stamps"`
	IsUnlimited        bool    `json:"isUnlimited"`
}

// CreateUserRequest representa la petición para crear un usuario
type CreateUserRequest struct {
	Name              string `json:"name"`
	TaxID             string `json:"taxId"`
	Email             string `json:"email"`
	Stamps            int    `json:"stamps"`
	IsUnlimited       bool   `json:"isUnlimited"`
	Password          string `json:"password"`
	NotificationEmail string `json:"notificationEmail"`
	Phone             string `json:"phone"`
}

// UpdateUserRequest representa la petición para actualizar un usuario
type UpdateUserRequest struct {
	Name              *string `json:"name,omitempty"`
	TaxID             *string `json:"taxId,omitempty"`
	IsUnlimited       *bool   `json:"isUnlimited,omitempty"`
	IDUser            string  `json:"iduser"`
	NotificationEmail *string `json:"notificationEmail,omitempty"`
	Phone             *string `json:"phone,omitempty"`
}

// ListUsersResponse representa la respuesta de listar usuarios
type ListUsersResponse struct {
	Data   []Usuario `json:"data"`
	Meta   Meta      `json:"meta"`
	Links  Links     `json:"links"`
	Status string    `json:"status"`
}

// CreateUserResponse representa la respuesta de crear usuario
type CreateUserResponse struct {
	Data   Usuario `json:"data"`
	Meta   *Meta   `json:"meta"`
	Links  *Links  `json:"links"`
	Status string  `json:"status"`
}

// UpdateUserResponse representa la respuesta de actualizar usuario
type UpdateUserResponse struct {
	Data   string `json:"data"`
	Meta   *Meta  `json:"meta"`
	Links  *Links `json:"links"`
	Status string `json:"status"`
}

// Meta representa la información de paginación
type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	PageCount  int `json:"pageCount"`
	TotalCount int `json:"totalCount"`
	TotalPages int `json:"totalPages"`
}

// Links representa los enlaces de paginación
type Links struct {
	Self  string `json:"Self"`
	First string `json:"First"`
	End   string `json:"End"`
	Next  string `json:"Next"`
}

// ListUsersParams representa los parámetros para listar usuarios
type ListUsersParams struct {
	IsActive *bool `json:"isActive,omitempty"`
	Page     *int  `json:"page,omitempty"`
	PerPage  *int  `json:"perPage,omitempty"`
}

// CrearUsuario crea un nuevo usuario
func CrearUsuario(client *autenticacion.SWClient, request *CreateUserRequest) (*CreateUserResponse, error) {
	requestHelper := NewRequestHelper()
	return requestHelper.SendCreateUserRequest(client.BaseURL, client.Token, request)
}

// ListarUsuarios obtiene la lista de usuarios
func ListarUsuarios(client *autenticacion.SWClient, params *ListUsersParams) (*ListUsersResponse, error) {
	requestHelper := NewRequestHelper()
	return requestHelper.SendListUsersRequest(client.BaseURL, client.Token, params)
}

// ActualizarUsuario actualiza un usuario existente
func ActualizarUsuario(client *autenticacion.SWClient, userID string, request *UpdateUserRequest) (*UpdateUserResponse, error) {
	requestHelper := NewRequestHelper()
	return requestHelper.SendUpdateUserRequest(client.BaseURL, client.Token, userID, request)
}

// EliminarUsuario elimina un usuario
func EliminarUsuario(client *autenticacion.SWClient, userID string) error {
	requestHelper := NewRequestHelper()
	return requestHelper.SendDeleteUserRequest(client.BaseURL, client.Token, userID)
}
