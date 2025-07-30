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
