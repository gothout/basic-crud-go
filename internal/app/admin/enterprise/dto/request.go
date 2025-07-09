package dto

type CreateEnterpriseDTO struct {
	Name string `json:"name" binding:"required"`
	Cnpj string `json:"cnpj" binding:"required"`
}

type ReadEnterpriseDTO struct {
	Cnpj string `uri:"cnpj" binding:"required"`
}

type ReadEnterprisesDTO struct {
	Page  int `form:"page" binding:"omitempty"`
	Limit int `form:"limit" binding:"omitempty"`
}
