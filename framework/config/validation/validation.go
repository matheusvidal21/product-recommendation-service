package validation

import (
	"encoding/json"
	"errors"
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/matheusvidal21/product-recommendation-service/framework/config/rest_err"
	"strings"
)

var (
	Validate     = validator.New()
	transl       ut.Translator
	validActions = []string{
		"view", "addtocart", "purchase", "removefromcart", "wishlist",
		"search", "rate", "review", "click", "share",
	}
)

func init() {
	en := en.New()
	uni := ut.New(en, en)
	transl, _ = uni.GetTranslator("en")
	en_translation.RegisterDefaultTranslations(Validate, transl)
	Validate.RegisterValidation("validAction", validateAction)
}

func ValidateStructError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		errorsCauses := []rest_err.Cause{}
		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Cause{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	}
	return rest_err.NewBadRequestError("Error trying to convert fields")
}

func validateAction(fl validator.FieldLevel) bool {
	action := fl.Field().String()
	for _, a := range validActions {
		if strings.ToLower(a) == strings.ToLower(action) {
			return true
		}
	}
	return false
}
