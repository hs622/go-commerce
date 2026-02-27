package schemas_test

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hs622/ecommerce-cart/schemas"
)

func TestCreateUserRequest(t *testing.T) {
	validate.SetTagName("binding")

	tests := []struct {
		Name     string
		Input    schemas.CreateUserRequest
		WantErr  bool
		ErrField string
	}{
		{
			Name: "Valid - all fields",
			Input: schemas.CreateUserRequest{
				UserID:       uuid.NewString(),
				FirstName:    "lara",
				LastName:     "smith",
				PrimaryEmail: "lara@gmail.com",
				Password:     "securePassword123",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			WantErr: false,
		},
		{
			Name: "Valid - optional fields omitted",
			Input: schemas.CreateUserRequest{
				UserID:       uuid.NewString(),
				FirstName:    "lara",
				LastName:     "smith",
				PrimaryEmail: "lara@gmail.com",
				Password:     "securePassword123",
			},
			WantErr: false,
		},
		{
			Name: "Invalid - missing user_id",
			Input: schemas.CreateUserRequest{
				FirstName:    "lara",
				LastName:     "simth",
				PrimaryEmail: "lara@gmail.com",
				Password:     "securePassword123",
			},
			WantErr:  true,
			ErrField: "UserID",
		},
		{
			Name: "Invalid - user_id not uuid4",
			Input: schemas.CreateUserRequest{
				UserID:       "not-a-valid-uuid",
				FirstName:    "lara",
				LastName:     "simth",
				PrimaryEmail: "lara@gmail.com",
				Password:     "securePassword123",
			},
			WantErr:  true,
			ErrField: "UserID",
		},
		{
			Name: "Invalid - first_name too short",
			Input: schemas.CreateUserRequest{
				UserID:       uuid.NewString(),
				FirstName:    "l",
				LastName:     "simth",
				PrimaryEmail: "lara@gmail.com",
				Password:     "securePassword123",
			},
			WantErr:  true,
			ErrField: "FirstName",
		},
		{
			Name: "Invalid - first_name too long",
			Input: schemas.CreateUserRequest{
				UserID:       uuid.NewString(),
				FirstName:    "larassssssssssssssssssssssssssss", // 32
				LastName:     "simth",
				PrimaryEmail: "lara@gmail.com",
				Password:     "securePassword123",
			},
			WantErr:  true,
			ErrField: "FirstName",
		},
		{
			Name: "Invalid - last_name too short",
			Input: schemas.CreateUserRequest{
				UserID:       uuid.NewString(),
				FirstName:    "lara",
				LastName:     "s",
				PrimaryEmail: "lara@gmail.com",
				Password:     "securePassword123",
			},
			WantErr:  true,
			ErrField: "LastName",
		},
		{
			Name: "Invalid - last_name too long",
			Input: schemas.CreateUserRequest{
				UserID:       uuid.NewString(),
				FirstName:    "lara",
				LastName:     "simthsssssssssssssssssssssssssss", // 32
				PrimaryEmail: "lara@gmail.com",
				Password:     "securePassword123",
			},
			WantErr:  true,
			ErrField: "LastName",
		},
		{
			Name: "Invalid - missing email",
			Input: schemas.CreateUserRequest{
				UserID:    uuid.NewString(),
				FirstName: "lara",
				LastName:  "simth",
				Password:  "securePassword123",
			},
			WantErr:  true,
			ErrField: "PrimaryEmail",
		},
		{
			Name: "Invalid - malformed email",
			Input: schemas.CreateUserRequest{
				UserID:       uuid.NewString(),
				FirstName:    "lara",
				LastName:     "smith",
				PrimaryEmail: "not-an-email",
				Password:     "securePassword123",
			},
			WantErr:  true,
			ErrField: "PrimaryEmail",
		},
		{
			Name: "Invalid - missing password",
			Input: schemas.CreateUserRequest{
				UserID:       uuid.NewString(),
				FirstName:    "lara",
				LastName:     "smith",
				PrimaryEmail: "lara@gmail.com",
			},
			WantErr:  true,
			ErrField: "Password",
		},
		{
			Name: "Invalid - malformed secondary email.",
			Input: schemas.CreateUserRequest{
				UserID:         uuid.NewString(),
				FirstName:      "lara",
				LastName:       "smith",
				PrimaryEmail:   "lara@gamil.com",
				Password:       "securePassowrd123",
				SecondaryEmail: "not-a-email",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			WantErr:  true,
			ErrField: "SecondaryEmail",
		},
	}

	for _, Case := range tests {
		t.Run(Case.Name, func(t *testing.T) {
			err := validate.Struct(Case.Input)

			if Case.WantErr {
				if err == nil {
					t.Errorf("Expected validation error but got none.")
					return
				}

				if Case.ErrField != "" {
					var found bool
					for _, e := range err.(validator.ValidationErrors) {
						if e.Field() == Case.ErrField {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Expected error on field %q, got %v", Case.ErrField, err)
					}
				}
			} else {
				if err != nil {
					t.Errorf("Expected no errors but got: %v\n", err)
				}
			}
		})
	}
}

func TestPatchUserRequest(t *testing.T) {
	validate.SetTagName("binding")

	tests := []struct {
		Name     string
		Input    schemas.PatchUserRequest
		WantErr  bool
		ErrField string
	}{
		{
			Name: "Valid - all fields",
			Input: schemas.PatchUserRequest{
				FirstName:      "Evan",
				LastName:       "Gabriel",
				PrimaryEmail:   "evan@gamil.com",
				SecondaryEmail: "backup-evan@gmail.com",
				UpdatedAt:      time.Now(),
			},
			WantErr: false,
		},
		{
			Name: "Valid - optional fields omitted",
			Input: schemas.PatchUserRequest{
				FirstName:    "Evan",
				LastName:     "Gabriel",
				PrimaryEmail: "evan@gamil.com",
				UpdatedAt:    time.Now(),
			},
			WantErr: false,
		},
		{
			Name: "Invalid - first_name is missing.",
			Input: schemas.PatchUserRequest{
				LastName:     "Gabriel",
				PrimaryEmail: "evan@gamil.com",
				UpdatedAt:    time.Now(),
			},
			WantErr:  true,
			ErrField: "FirstName",
		},
		{
			Name: "Invalid - first_name is too long.",
			Input: schemas.PatchUserRequest{
				FirstName:    "Evanssssssssssssssssssssssssssssssss",
				LastName:     "Gabriel",
				PrimaryEmail: "evan@gamil.com",
				UpdatedAt:    time.Now(),
			},
			WantErr:  true,
			ErrField: "FirstName",
		},
		{
			Name: "Invalid - first_name is too short.",
			Input: schemas.PatchUserRequest{
				FirstName:    "E",
				LastName:     "Gabriel",
				PrimaryEmail: "evan@gamil.com",
				UpdatedAt:    time.Now(),
			},
			WantErr:  true,
			ErrField: "FirstName",
		},
		{
			Name: "Invalid - last_name is too long.",
			Input: schemas.PatchUserRequest{
				FirstName:    "Evan",
				LastName:     "Gabrielssssssssssssssssssssssssss",
				PrimaryEmail: "evan@gamil.com",
				UpdatedAt:    time.Now(),
			},
			WantErr:  true,
			ErrField: "LastName",
		},
		{
			Name: "Invalid - last_name is too short.",
			Input: schemas.PatchUserRequest{
				FirstName:    "Evan",
				LastName:     "G",
				PrimaryEmail: "evan@gamil.com",
				UpdatedAt:    time.Now(),
			},
			WantErr:  true,
			ErrField: "LastName",
		},
		{
			Name: "Invalid - last_name is missing.",
			Input: schemas.PatchUserRequest{
				FirstName:    "Evan",
				PrimaryEmail: "evan@gamil.com",
				UpdatedAt:    time.Now(),
			},
			WantErr:  true,
			ErrField: "LastName",
		},
		{
			Name: "Invalid - email is missing.",
			Input: schemas.PatchUserRequest{
				FirstName: "Evan",
				LastName:  "Gabriel",
				UpdatedAt: time.Now(),
			},
			WantErr:  true,
			ErrField: "PrimaryEmail",
		},
		{
			Name: "Invalid - malformed email.",
			Input: schemas.PatchUserRequest{
				FirstName: "Evan",
				LastName:  "Gabriel",
				UpdatedAt: time.Now(),
			},
			WantErr:  true,
			ErrField: "PrimaryEmail",
		},
		{
			Name: "Invalid - malformed secondary email.",
			Input: schemas.PatchUserRequest{
				FirstName:      "Evan",
				LastName:       "Gabriel",
				PrimaryEmail:   "evan@gamil.com",
				SecondaryEmail: "not-a-email",
				UpdatedAt:      time.Now(),
			},
			WantErr:  true,
			ErrField: "SecondaryEmail",
		},
	}

	for _, Case := range tests {
		t.Run(Case.Name, func(t *testing.T) {
			err := validate.Struct(Case.Input)

			if Case.WantErr {
				if err == nil {
					t.Errorf("Expected validation error, but got none!")
					return
				}

				if Case.ErrField != "" {
					var found bool
					for _, e := range err.(validator.ValidationErrors) {
						if e.Field() == Case.ErrField {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Expected error on field %q, got: %v", Case.ErrField, err)
					}
				}
			} else {
				if err != nil {
					t.Errorf("Expected no errors, but we got: %v\n", err)
				}
			}
		})
	}
}
