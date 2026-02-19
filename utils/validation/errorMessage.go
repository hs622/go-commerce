package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/hs622/ecommerce-cart/utils"
)

func GetCustomErrorMessage(fe validator.FieldError) string {
	var field string = fe.Field()

	if !utils.IsAcronym(fe.Field()) {
		field = utils.ToPlainText(fe.Field())
	}

	switch fe.Tag() {
	case "required":
		return "The " + field + " field is mandatory."
	case "min":
		if fe.Kind() == reflect.Int || fe.Kind() == reflect.Float64 {
			return field + " must be at least " + fe.Param()
		}
		return field + " must be at least " + fe.Param() + " characters long."
	case "max":
		return field + " cannot exceed " + fe.Param() + " characters."
	case "dive":
		return "One of the items in " + field + " is invalid."
	}
	return "The value for " + field + " is invalid."
}
