package user

import (
	"backend/internal/domain/user"
	"backend/internal/infrastructure/presistence/postgres"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToDomain(u postgres.User) *user.User {
	var birthday *time.Time
	if u.Birthday.Valid {
		t := u.Birthday.Time
		birthday = &t
	}

	var createdAt *time.Time
	if u.CreatedAt.Valid {
		t := u.CreatedAt.Time
		createdAt = &t
	}

	var deletedAt *time.Time
	if u.DeletedAt.Valid {
		t := u.DeletedAt.Time
		deletedAt = &t
	}

	return &user.User{
		ID:          u.ID.String(),
		Avatar:      u.Avatar.String,
		Birthday:    birthday,
		PhoneNumber: u.PhoneNumber.String,
		CreatedAt:   createdAt,
		DeletedAt:   deletedAt,
	}
}

func ToCreateUserParams(u *user.User) postgres.CreateUserParams {
	return postgres.CreateUserParams{
		Avatar:      stringToPgText(u.Avatar),
		Birthday:    timeToPgDate(u.Birthday),
		PhoneNumber: stringToPgText(u.PhoneNumber),
	}
}

func ToUpdateUserParams(u *user.User) postgres.UpdateUserParams {
	return postgres.UpdateUserParams{
		ID:          uuid.MustParse(u.ID),
		Avatar:      stringToPgText(u.Avatar),
		Birthday:    timeToPgDate(u.Birthday),
		PhoneNumber: stringToPgText(u.PhoneNumber),
	}
}

func stringToPgText(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: true}
}

func timeToPgDate(t *time.Time) pgtype.Date {
	return pgtype.Date{Time: *t, Valid: true}
}
