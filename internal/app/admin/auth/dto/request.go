package dto

import "time"

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenUserDTO struct {
	Email string `json:"email" binding:"required"`
}

type LogoutUserDTO struct {
	Email string `json:"email" binding:"required"`
}

type CreateTokenAPIDto struct {
	Email     string    `json:"email" binding:"required" example:"admin@test.com"`
	ExpiresAt time.Time `json:"expires" binding:"required" example:"2025-12-31T23:59:59Z"`
}
