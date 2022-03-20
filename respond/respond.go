package respond

import (
	"encoding/json"
	"net/http"
	"user-service/internal/app_errors"
)

func With(w http.ResponseWriter, r *http.Request, httpStatus int, bodyData interface{}) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(bodyData)
}

func WithValidationErrors(w http.ResponseWriter, r *http.Request, httpStatus int, Errors []app_errors.AppError) {

	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode("bodyData")
}
