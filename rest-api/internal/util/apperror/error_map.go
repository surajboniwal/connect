package apperror

var DBErrorMap = map[string]string{
	"email_unique":                    "Email is already in use",
	"phone_unique":                    "Phone number is already in use",
	"organization_user_unique":        "Looks like this user is already part of this organization",
	"organization_users_user_id_fkey": "Looks like this user does not exists",
}
