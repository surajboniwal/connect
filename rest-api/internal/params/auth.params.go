package params

type Register struct {
	Name              string `json:"name" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required"`
	Phone             string `json:"phone" validate:"required,e164"`
	Organization_Name string `json:"organization_name" validate:"required"`
}
