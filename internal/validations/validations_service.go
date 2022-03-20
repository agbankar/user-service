package validations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"user-service/internal/app_errors"
	"user-service/respond"
)

type AppValidator interface {
	ApplyStaticValidations(w http.ResponseWriter, r *http.Request, requestDto interface{}) ([]app_errors.AppError, error)
	ApplyBusinessValidations(handlFunc func(requestDto interface{}) []app_errors.AppError) bool
}

//This is empty as it doesnt required external dependencies
type AppValidations struct {
}

func NewValidationService() *AppValidations {
	return &AppValidations{}

}

func (a *AppValidations) ApplyStaticValidations(w http.ResponseWriter, r *http.Request, requestDto interface{}) ([]app_errors.AppError, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respond.With(w, r, http.StatusBadRequest, nil)
	}
	//after reading the body set it back to the request, so that it will not be nil
	reader := ioutil.NopCloser(bytes.NewBuffer(body))
	r.Body = reader

	// unmarshall from request
	err = json.Unmarshal(body, requestDto)
	if err != nil {
		respond.With(w, r, http.StatusBadRequest, nil)
		return nil, nil
	}
	// validate request
	err = V.Struct(requestDto)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Println("JSONResponseFromValidation - expecting err to be validator.ValidationErrors, found",
				reflect.TypeOf(err).Name(), ". Value:", err.Error())
			return nil, nil
		}

		for _, err := range validationErrors {
			code := err.Param()
			rule := LookupValidationRule(code)
			for _, err := range validationErrors {
				code := err.Param()
				if rule == nil {
					log.Printf("ValidationRule not found for code: '%s'\n", code)
				} else {
					element := app_errors.AppError{
						ErrorCode:        rule.Code(),
						ErrorDescription: rule.Description(),
					}
					fmt.Println(element)
					//TODO: respond.WithValidationErrors()
				}
			}
		}
	}
	//respond.WithValidationErrors(w, r, http.StatusBadRequest, err)

	return nil, nil
}
func (a *AppValidations) ApplyBusinessValidations(handlFunc func(requestDto interface{}) []app_errors.AppError) bool {
	return true
}
