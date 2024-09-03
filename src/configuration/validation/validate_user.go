package validation

import (
	"encoding/json"
	"errors"
	"github.com/MateusGuedess/crudzinho-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate   = validator.New()
	translator ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		universalTranslator := ut.New(en, en)
		translator, _ = universalTranslator.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, translator)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	//Error in the type of the value in the unmarshal moment
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	//Verify if error is in the type of the value
	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("invalid field type")
		//Verify if the error is invalid values
	} else if errors.As(validation_err, &jsonValidationError) {
		var errorsCauses []rest_err.Causes

		for _, err := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: err.Translate(translator),
				Field:   err.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)

	}
	//return error if none of the conditions above solve
	return rest_err.NewBadRequestError("Error trying to convert fields")

}
