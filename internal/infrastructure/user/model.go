package user

import (
	"backend/internal/constant"
	"backend/internal/domain/user"
	"backend/internal/infrastructure/presistence/postgres"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
)

func ToDomain(u *gocloak.User) *user.User {
	dob, _ := (time.Parse("2006-01-02", getAttr(u, string(constant.UserAttributeDateOfBirth))))
	createdAt := time.UnixMilli(*u.CreatedTimestamp)
	return &user.User{
		ID:          *u.ID,
		FirstName:   getAttr(u, string(constant.UserAttributeFirstName)),
		LastName:    getAttr(u, string(constant.UserAttributeLastName)),
		UserName:    *u.Username,
		Email:       getAttr(u, string(constant.UserAttributeEmail)),
		Address:     getAttr(u, string(constant.UserAttributeAddress)),
		Birthday:    &dob,
		PhoneNumber: getAttr(u, string(constant.UserAttributePhoneNumber)),
		CreatedAt:   &createdAt,
	}
}

func ToCreateUserParams(u *user.User) postgres.CreateUserParams {
	return postgres.CreateUserParams{
		ID: uuid.MustParse(u.ID),
	}
}

func ToUpdateUserParams(u *user.User) gocloak.User {
	attributes := make(map[string][]string)
	attributes["first_name"] = []string{u.FirstName}
	attributes["last_name"] = []string{u.LastName}
	attributes["email"] = []string{u.Email}
	attributes["phone_numer"] = []string{u.PhoneNumber}
	attributes["address"] = []string{u.Address}
	return gocloak.User{
		Attributes: &attributes,
	}
}

func getAttr(u *gocloak.User, key string) string {
	if u.Attributes == nil {
		return ""
	}
	if vals, ok := (*u.Attributes)[key]; ok && len(vals) > 0 {
		return vals[0]
	}
	return ""
}
