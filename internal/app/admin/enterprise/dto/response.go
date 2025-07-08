package dto

import "time"

type CreateEnterpriseResponse struct {
	Name      string    `json:"name"`
	Cnpj      string    `json:"cnpj"`
	CreatedAt time.Time `json:"createdAt"`
}
