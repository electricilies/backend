package user

import (
	"time"

	"backend/internal/constant"
	"backend/internal/domain/user"
	"backend/internal/infrastructure/persistence/postgres"

	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
)

func ToDomain(u *gocloak.User) *user.Model {
	dob, _ := (time.Parse("2006-01-02", *getAttr(u, string(constant.UserAttributeDateOfBirth))))
	createdAt := time.UnixMilli(*u.CreatedTimestamp)
	id := uuid.MustParse(*u.ID)
	return &user.Model{
		ID:          &id,
		FirstName:   getAttr(u, string(constant.UserAttributeFirstName)),
		LastName:    getAttr(u, string(constant.UserAttributeLastName)),
		UserName:    u.Username,
		Email:       getAttr(u, string(constant.UserAttributeEmail)),
		Address:     getAttr(u, string(constant.UserAttributeAddress)),
		DateOfBirth: &dob,
		PhoneNumber: getAttr(u, string(constant.UserAttributePhoneNumber)),
		CreatedAt:   &createdAt,
	}
}

func ToCreateUserParams(model *user.Model) postgres.CreateUserParams {
	return postgres.CreateUserParams{
		ID: *model.ID,
	}
}

func ToUpdateUserParams(model *user.Model, id *uuid.UUID) *gocloak.User {
	idString := id.String()
	attributes := make(map[string][]string)
	attributes["email"] = []string{*model.Email}
	attributes["phone_numer"] = []string{*model.PhoneNumber}
	attributes["address"] = []string{*model.Address}
	return &gocloak.User{
		ID:         &idString,
		FirstName:  model.FirstName,
		LastName:   model.LastName,
		Attributes: &attributes,
	}
}

func getAttr(u *gocloak.User, key string) *string {
	if u.Attributes == nil {
		return nil
	}
	if vals, ok := (*u.Attributes)[key]; ok && len(vals) > 0 {
		return &vals[0]
	}
	return nil
}
