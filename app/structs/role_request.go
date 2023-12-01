package structs

type RoleStoreRequest struct {
	Name string `json:"name" validate:"required,role_name_available,lowercase"`
}
