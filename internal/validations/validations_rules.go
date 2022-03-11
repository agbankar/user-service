package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var v *validator.Validate
var validationRuleMap = map[string]ValidationRule{}

type ValidationRule interface {
	Code() string
	Description() string
}
type RegexpRule interface {
	ValidationRule
	Regexp() *regexp.Regexp
}

func RegisterValidationRule(r ValidationRule) {
	validationRuleMap[r.Code()] = r
}

func init() {
	v = validator.New()
	v.RegisterValidation("code", ValidateErrorCode)
}

func ValidateErrorCode(f validator.FieldLevel) bool {
	fmt.Println("Inside validations")
	code := f.Param()
	fmt.Println()
	fmt.Println(code)
	rule, ok := validationRuleMap[code]
	if ok {
		fmt.Println(rule)
	}
	return true
}
func ValidateRequest(r interface{}) {
	v.Struct(r)

}
