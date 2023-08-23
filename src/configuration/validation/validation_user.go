package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/intwone/golang-api/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		en := en.New()
		unicode := ut.New(en, en)
		transl, _ = unicode.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("invalid field type")
	}

	if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Cause{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("some field are invalids", errorsCauses)
	}

	return rest_err.NewBadRequestError("error trying to convert fields")
}
