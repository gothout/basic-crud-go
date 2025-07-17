package model

type ModulePermission struct {
	ID      int64
	Name    string
	Actions []PermissionAction
}

type PermissionAction struct {
	ID   int64
	Name string
}
