package validations

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"user-service/internal/app_errors"
	"user-service/respond"
)

type AppValidator interface {
	ApplyStaticValidations(w http.ResponseWriter, r *http.Request, requestDto interface{}, validateReq bool) bool
	//This should be abstract , its the responsibility of the caller to pass the business validation function
	//TODO: Implement inheritance
	ApplyBusinessValidations(handlFunc func(requestDto interface{}) []app_errors.AppError) bool
}

//This is empty as it doesnt required external dependencies
type AppValidations struct {
}

func NewValidationService() *AppValidations {
	return &AppValidations{}

}

func (a *AppValidations) ApplyStaticValidations(w http.ResponseWriter, r *http.Request, requestDto interface{}, validateReq bool) bool {
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
		return false
	}
	// validate request
	if validateReq {
		err = v.Struct(requestDto)
		if err != nil {
			respond.With(w, r, http.StatusBadRequest, nil)
			return false
		}
	}
	return true
}
func (a *AppValidations) ApplyBusinessValidations(handlFunc func(requestDto interface{}) []app_errors.AppError) bool {
	return true
}
