package constant

type (
	Key           string
	UserAttribute string
)

const (
	TokenKey Key = "token"

	UserAttributeFirstName   UserAttribute = "first_name"
	UserAttributeLastName    UserAttribute = "last_name"
	UserAttributeEmail       UserAttribute = "email"
	UserAttributeAddress     UserAttribute = "address"
	UserAttributeDateOfBirth UserAttribute = "date_of_birth"
	UserAttributePhoneNumber UserAttribute = "phone_number"
	UserAttributeCreatedAt   UserAttribute = "created_at"
)
