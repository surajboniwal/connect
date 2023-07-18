package params

type Register struct {
	Name              string `json:"name" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required"`
	Organization_Name string `json:"organization_name" validate:"required"`
}
