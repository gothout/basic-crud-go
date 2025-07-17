package dto

type ReadPermissionResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ReadPermissionsResponse struct {
	Permissions []ReadPermissionResponse `json:"permissions"`
}
