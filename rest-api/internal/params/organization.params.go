package params

type CreateOrganization struct {
	Name string `json:"name" validate:"required"`
}
