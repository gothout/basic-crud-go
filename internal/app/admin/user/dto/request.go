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

type ReadUsersDTO struct {
	Page  int `form:"page" binding:"omitempty"`
	Limit int `form:"limit" binding:"omitempty"`
}

type UpdateUserDTO struct {
	Email        string `uri:"email" binding:"required"`
	FirstName    string `json:"first_name" binding:"omitempty"`
	LastName     string `json:"last_name" binding:"omitempty"`
	EmailUpdated string `json:"email" binding:"omitempty"`
	Number       string `json:"number" binding:"omitempty"`
	Password     string `json:"password" binding:"omitempty"`
}
