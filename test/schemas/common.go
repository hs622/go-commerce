package schemas_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hs622/ecommerce-cart/schemas"
)

var (
	validate = validator.New()
)

type TestType struct {
	Name     string
	Input    interface{}
	WantErr  bool
	ErrField string
}

func Runner(tests []TestType, t *testing.T) {
	fmt.Println("Runner started")
	for _, newCase := range tests {

		t.Run(newCase.Name, func(t *testing.T) {
			err := validate.Struct(newCase.Input)

			if newCase.WantErr {
				if err == nil {
					t.Errorf("Expected error validation, but got none")
					return
				}

				if newCase.ErrField != "" {
					var found bool
					for _, c := range err.(validator.ValidationErrors) {
						if c.Field() == newCase.ErrField {
							found = true
							break
						}
					}

					if !found {
						t.Errorf("Expected error on field %q, got: %v", newCase.ErrField, err)
					}
				}

			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v\n", err)
				}
			}
		})
	}
}

func ValidToken() string {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": uuid.NewString(),
		"email":   "example@gmail.com",
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	signed, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	return signed
}

func DeviceInfo() schemas.DeviceInfo {
	return schemas.DeviceInfo{
		Platform:   "liunx",
		Agent:      "Mozilla/5.0",
		IP:         "192.168.165.58",
		Cordinates: []float64{25.1548, 36.4587},
	}
}
