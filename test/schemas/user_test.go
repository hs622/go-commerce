package schemas_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hs622/ecommerce-cart/schemas"
)

func TestCreateUserRequest(t *testing.T) {
	validate.SetTagName("binding")
	tt := time.Now()

	tests := []TestType{
		{
			Name: "Valid - all fields",
			Input: schemas.CreateUserRequest{
				UserID:       uuid.NewString(),
				FirstName:    "lara",
				LastName:     "smith",
				PrimaryEmail: "lara@gmail.com",
				Password:     "securePassword123",
				CreatedAt:    &tt,
				UpdatedAt:    &tt,
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
				CreatedAt:      &tt,
				UpdatedAt:      &tt,
			},
			WantErr:  true,
			ErrField: "SecondaryEmail",
		},
	}

	Runner(tests, t)
}

func TestPatchUserRequest(t *testing.T) {
	validate.SetTagName("binding")
	tt := time.Now()

	tests := []TestType{
		{
			Name: "Valid - all fields",
			Input: schemas.PatchUserRequest{
				FirstName:      "Evan",
				LastName:       "Gabriel",
				PrimaryEmail:   "evan@gamil.com",
				SecondaryEmail: "backup-evan@gmail.com",
				UpdatedAt:      &tt,
			},
			WantErr: false,
		},
		{
			Name: "Valid - optional fields omitted",
			Input: schemas.PatchUserRequest{
				UpdatedAt: &tt,
			},
			WantErr: false,
		},
		{
			Name: "Invalid - first_name is too long.",
			Input: schemas.PatchUserRequest{
				FirstName: "Evanssssssssssssssssssssssssssssssss",
				UpdatedAt: &tt,
			},
			WantErr:  true,
			ErrField: "FirstName",
		},
		{
			Name: "Invalid - first_name is too short.",
			Input: schemas.PatchUserRequest{
				FirstName: "E",
				UpdatedAt: &tt,
			},
			WantErr:  true,
			ErrField: "FirstName",
		},
		{
			Name: "Invalid - last_name is too long.",
			Input: schemas.PatchUserRequest{
				LastName:  "Gabrielssssssssssssssssssssssssss",
				UpdatedAt: &tt,
			},
			WantErr:  true,
			ErrField: "LastName",
		},
		{
			Name: "Invalid - last_name is too short.",
			Input: schemas.PatchUserRequest{
				LastName:  "G",
				UpdatedAt: &tt,
			},
			WantErr:  true,
			ErrField: "LastName",
		},
		{
			Name: "Invalid - malformed email.",
			Input: schemas.PatchUserRequest{
				PrimaryEmail: "not-a-valid-email",
				UpdatedAt:    &tt,
			},
			WantErr:  true,
			ErrField: "PrimaryEmail",
		},
		{
			Name: "Invalid - malformed secondary email.",
			Input: schemas.PatchUserRequest{
				SecondaryEmail: "not-a-valid-email",
				UpdatedAt:      &tt,
			},
			WantErr:  true,
			ErrField: "SecondaryEmail",
		},
	}

	Runner(tests, t)
}
