package validations

import (
	"net/http"
	"regexp"
)

type ErrorCode interface {
	Code() string
	Description() string
	HTTPStatus() int
}

type errorCode struct {
	code        string
	description string
	httpStatus  int
}

//Code function
func (e errorCode) Code() string {
	return e.code
}

//Description function
func (e errorCode) Description() string {
	return e.description
}

//HTTPStatus function
func (e errorCode) HTTPStatus() int {
	return e.httpStatus
}

// RegexpErrorCode is ErrorCode implementation for regular expressions
type RegexpErrorCode struct {
	errorCode
	regexp *regexp.Regexp
}

// Regexp gets the regular expression for the validation
func (e RegexpErrorCode) Regexp() *regexp.Regexp {
	return e.regexp
}

// NewRegexpCode creates an ErrorCode for validating regexp input
func NewRegexpCode(code string, description string, pattern string) RegexpErrorCode {
	rtn := RegexpErrorCode{}
	rtn.code = code
	rtn.description = description
	rtn.httpStatus = http.StatusBadRequest
	rtn.regexp = regexp.MustCompile(pattern)
	RegisterValidationRule(rtn)
	return rtn
}

var (
	CODE101 = NewRegexpCode("CODE101", "Name  is required. AlphaNumeric Max 50 characters.", "^[ a-zA-Z0-9]{1,50}$")
	CODE102 = NewRegexpCode("CODE102", "Id  is required. Numeric only Max 3 characters.", "^[0-9]{1,3}$")
)
