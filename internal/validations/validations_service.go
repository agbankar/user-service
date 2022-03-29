package validations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"io/ioutil"
	"net/http"
	"user-service/internal/app_errors"
	"user-service/respond"
)

type AppValidator interface {
	ApplyStaticValidations(w http.ResponseWriter, r *http.Request, requestDto interface{}) bool
	ApplyBusinessValidations(handlFunc func(requestDto interface{}) []app_errors.AppError) bool
}

// AppValidations This is empty as it doesn't requires external dependencies
type AppValidations struct {
}

func NewValidationService() *AppValidations {
	return &AppValidations{}

}

func (a *AppValidations) ApplyStaticValidations(w http.ResponseWriter, r *http.Request, requestDto interface{}) bool {
	var AppErrors []app_errors.AppError

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("validayion_service::ApplyStaticValidations , io.ReadAll(r.Body) thrown errorr", err)
		respond.WithAppError(w, "App Error :: Please contact support team")
		return false
	}
	//After reading the body set it back to the request, so that it will not be nil
	reader := ioutil.NopCloser(bytes.NewBuffer(body))
	r.Body = reader
	// unmarshall from request
	err = json.Unmarshal(body, requestDto)
	if err != nil {
		fmt.Println("validayion_service::ApplyStaticValidations ,json.Unmarshal(body), requestDto) thrown errorr", err)
		respond.WithAppError(w, "App Error :: Please contact support team")
		return false
	}
	// validate request
	err = V.Struct(requestDto)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			fmt.Println("validayion_service::ApplyStaticValidations ,V.Struct(requestDto), requestDto) thrown error", err)
			return false
		}
		for _, err := range validationErrors {
			code := err.Param()
			rule := LookupValidationRule(code)
			for _, err := range validationErrors {
				code := err.Param()
				if rule == nil {
					fmt.Printf("ValidationRule not found for code: '%s'\n", code)
				} else {
					e := app_errors.AppError{
						ErrorCode:        rule.Code(),
						ErrorDescription: rule.Description(),
					}
					AppErrors = append(AppErrors, e)
				}
			}
		}
		if len(AppErrors) != 0 {
			respond.WithValidationErrors(w, http.StatusBadRequest, AppErrors)
			return false
		}
	}
	return true
}

//TODO:
func (a *AppValidations) ApplyBusinessValidations(handleFunc func(requestDto interface{}) []app_errors.AppError) bool {
	return true
}
