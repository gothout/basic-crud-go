package dto

type CreateUserResponse struct {
	Cnpj      string `json:"cnpj"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Number    string `json:"number"`
}
