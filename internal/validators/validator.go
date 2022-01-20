package validators

import "regexp"

type ValidationRule interface {
	Code() string
	Description() string
}

// If a ValidationRule also implements RegexpRule, it tests a regexp
type RegexpRule interface {
	ValidationRule
	Regexp() *regexp.Regexp
}
