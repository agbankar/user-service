package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
)

var V *validator.Validate
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
	V = validator.New()
	V.RegisterValidation("code", ValidateErrorCode)
}

func ValidateErrorCode(f validator.FieldLevel) bool {
	fmt.Println("Inside validations")
	code := f.Param()
	fmt.Println()
	fmt.Println(code)
	rule, ok := validationRuleMap[code]
	if r, ok := rule.(RegexpRule); ok {
		var value string
		//Logic to handle integer input fields
		if f.Field().Type().Name() == "int" {
			value = strconv.Itoa(int(f.Field().Int()))
		} else {
			value = f.Field().String()
		}
		return r.Regexp().MatchString(value)

	}
	if !ok {
		fmt.Println("Validators - a DTO was annotated with an unregistered ValidationRule code: ", code)
		fmt.Println(rule)
		return false
	}
	return true
}
func LookupValidationRule(code string) ValidationRule {
	return validationRuleMap[code]
}

//Not in use for now
