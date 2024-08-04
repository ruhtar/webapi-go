package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/ruhtar/webapi-go/src/configuration/rest_err"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

// global variable
var (
	Validate = validator.New()
	transl   ut.Translator
)

// The init function is automatically called when the package is initialized,
// before the main function or any other code in the package is executed.
func init() {
	engine := binding.Validator.Engine()    //any
	val, ok := engine.(*validator.Validate) //casting

	if ok {
		en := en.New()
		unt := ut.New(en, en)
		trans, _ := unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, trans)
	}
}

func ValidateError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
