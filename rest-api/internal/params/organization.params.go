package params

type CreateOrganization struct {
	Name    string `json:"name" validate:"required"`
	User_Id int64  `json:"user_id" validate:"required"`
}
