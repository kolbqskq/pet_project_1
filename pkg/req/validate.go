package req

import (
	"MinersGame/pkg/errs"

	"github.com/go-playground/validator/v10"
)

func IsValid[T any](body T) error {
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			t := e.Tag()
			f := e.Field()
			errMsg := "field:" + f + " is " + t
			return errs.NewCustomErr(400, errMsg)
		}
	}

	return nil

}
