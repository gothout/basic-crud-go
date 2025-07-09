package dto

import "time"

type CreateEnterpriseResponse struct {
	Name      string    `json:"name"`
	Cnpj      string    `json:"cnpj"`
	CreatedAt time.Time `json:"createdAt"`
}

type ReadEnterpriseResponse struct {
	Name      string    `json:"name"`
	Cnpj      string    `json:"cnpj"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ReadEnterprisesResponse struct {
	Page        int                      `json:"page"`
	Limit       int                      `json:"limit"`
	Total       int                      `json:"total"`
	Enterprises []ReadEnterpriseResponse `json:"enterprises"`
}
