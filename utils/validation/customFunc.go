package validation

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomRequestValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("required_if_item_is_available", RequiredIfItemIsAvailable)
	}
}

func RequiredIfItemIsAvailable(field validator.FieldLevel) bool {

	parent := field.Parent()
	itemsField := parent.FieldByName("Items")

	if !itemsField.IsValid() || itemsField.Kind() != reflect.Slice {
		return true
	}

	if itemsField.Len() > 0 {
		fieldValue := field.Field()
		return !fieldValue.IsZero()
	}

	return true
}
