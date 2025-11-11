package application

import (
	"context"
	"testing"

	"backend/internal/domain/user"

	"github.com/google/uuid"
)

type mockUserRepo struct{}

func (m *mockUserRepo) Get(ctx context.Context, id string) (*user.Model, error) {
	if id == "1" {
		return &user.Model{
			ID:        uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
			UserName:  "testuser",
			FirstName: "Test",
			LastName:  "User",
			Email:     "test@example.com",
		}, nil
	}
	return nil, ErrNotFound
}

// TODO: Complete this
func (m *mockUserRepo) List(ctx context.Context) ([]*user.Model, error) {
	return nil, nil
}

// TODO: Complete this
func (m *mockUserRepo) Create(ctx context.Context, u *user.Model) (*user.Model, error) {
	return nil, nil
}

// TODO: Complete this
func (m *mockUserRepo) Update(ctx context.Context, u *user.Model) error {
	return nil
}

// TODO: Complete this
func (m *mockUserRepo) Delete(ctx context.Context, id string) error {
	return nil
}

type userNotFoundError struct{}

func (e *userNotFoundError) Error() string { return "user not found" }

var ErrNotFound = &userNotFoundError{}

func Test_userApp_Get(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		want    *user.Model
		wantErr bool
	}{
		{
			name: "success: user found",
			id:   "1",
			want: &user.Model{
				ID:        uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
				UserName:  "testuser",
				FirstName: "Test",
				LastName:  "User",
				Email:     "test@example.com",
			},
			wantErr: false,
		},
		{
			name:    "fail: user not found",
			id:      "2",
			want:    nil,
			wantErr: true,
		},
	}
	a := &userApp{userRepo: &mockUserRepo{}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := a.Get(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && got != nil {
				if got.ID != tt.want.ID ||
					got.UserName != tt.want.UserName ||
					got.FirstName != tt.want.FirstName ||
					got.LastName != tt.want.LastName ||
					got.Email != tt.want.Email {
					t.Errorf("Get() = %v, want %v", got, tt.want)
					return
				}
			}
			if tt.want == nil && got != nil {
				t.Errorf("Get() = %v, want nil", got)
			}
		})
	}
}
