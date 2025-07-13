package dto

import (
	enterprise "basic-crud-go/internal/app/admin/enterprise/dto"
	"time"
)

type CreateUserResponse struct {
	Cnpj      string `json:"cnpj"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Number    string `json:"number"`
}

type ReadUserResponse struct {
	Id         string                            `json:"id"`
	FirstName  string                            `json:"first_name"`
	LastName   string                            `json:"last_name"`
	Email      string                            `json:"email"`
	Number     string                            `json:"number"`
	CreatedAt  time.Time                         `json:"created_at"`
	UpdatedAt  time.Time                         `json:"updated_at"`
	Enterprise enterprise.ReadEnterpriseResponse `json:"enterprise"`
}

type ReadUsersResponse struct {
	Users []ReadUserResponse `json:"users"`
}

type UpdateUserResponse struct {
	UpdatedUser ReadUserResponse `json:"updated"`
}
