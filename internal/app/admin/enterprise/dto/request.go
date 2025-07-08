package dto

type CreateEnterpriseDTO struct {
	Name string `json:"name" binding:"required"`
	Cnpj string `json:"cnpj" binding:"required"`
}
