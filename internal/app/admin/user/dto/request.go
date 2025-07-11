package dto

type CreateUserDTO struct {
	Cnpj      string `json:"cnpj" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Number    string `json:"number" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type ReadUserDTO struct {
	Email string `uri:"email" binding:"required"`
}
