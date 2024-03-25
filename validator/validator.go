package validator

import (
	"fmt"
	"strings"

	"github.com/fatih/camelcase"
	vali "github.com/go-playground/validator/v10"
)

type Validator struct {
	messages  map[string]string
	validator *vali.Validate
}

var Instance *Validator

func Init() {
	validate := vali.New(vali.WithRequiredStructEnabled())

	Instance = &Validator{
		messages:  Messages,
		validator: validate,
	}
}

func (v *Validator) Validate(class interface{}) []string {
	var errors []string

	err := v.validator.Struct(class)
	if err != nil {
		for _, err := range err.(vali.ValidationErrors) {
			field := strings.ToLower(strings.Join(camelcase.Split(err.Field()), " "))
			param := strings.ToLower(strings.Join(camelcase.Split(err.Param()), " "))

			errors = append(errors, fmt.Sprintf(v.messages[err.Tag()], field, param))
		}
	}

	return errors
}
