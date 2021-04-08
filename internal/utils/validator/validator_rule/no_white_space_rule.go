package validatorrule

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// NoWhiteSpaceRule ..
type NoWhiteSpaceRule struct{ *CommonRule }

// GetRule ...
func (*NoWhiteSpaceRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().String()
		if data != "" {
			matched, err := regexp.Match(`^[\S]+$`, []byte(data))
			if err != nil {
				return false
			}
			return matched
		}
		return true
	}
}

// NewNoWhiteSpaceRule ...
func NewNoWhiteSpaceRule() *NoWhiteSpaceRule {
	return &NoWhiteSpaceRule{&CommonRule{
		Field:     "no_white_space",
		FieldName: "no_white_space",
		Message:   "{0} invalid format!",
	}}
}
