package dto

import (
	permDto "basic-crud-go/internal/app/admin/permission/dto"
	userDTO "basic-crud-go/internal/app/admin/user/dto"
)

type LoginUserResponse struct {
	User        userDTO.ReadUserResponse        `json:"user"`
	Token       string                          `json:"token"`
	Permissions permDto.ReadPermissionsResponse `json:"permissions"`
}

type RefreshTokenUserResponse struct {
	Message string `json:"message" example:"Token refreshed successfully"`
}

type CreateTokenAPIResponse struct {
	Token string `json:"token" example:"645fe33b232a1b0c19f4e2bf1e475df2aa381cf45457ab6885cfb9c4bcd9e288"`
}
