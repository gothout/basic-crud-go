package dto

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
