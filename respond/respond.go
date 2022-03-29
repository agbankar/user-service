package respond

import (
	"encoding/json"
	"net/http"
	"user-service/internal/app_errors"
)

func With(w http.ResponseWriter, httpStatus int, bodyData interface{}) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(bodyData)
}

func WithValidationErrors(w http.ResponseWriter, httpStatus int, Errors []app_errors.AppError) {

	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(Errors)
}

func WithErrors(w http.ResponseWriter, httpStatus int, Error string) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(Error)
}

func WithAppError(w http.ResponseWriter, Errors string) {

	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Errors)
}
