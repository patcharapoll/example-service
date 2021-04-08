package validatorrule

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// TokenRule ...
type TokenRule struct{ *CommonRule }

// GetRule ...
func (*TokenRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().String()
		if data != "" {
			matched, err := regexp.Match(`^[a-zA-Z0-9\-_]+?\.[a-zA-Z0-9\-_]+?\.([a-zA-Z0-9\-_]+)?$`, []byte(data))
			if err != nil {
				return false
			}
			return matched
		}
		return true
	}
}

// NewTokenRule ...
func NewTokenRule() *TokenRule {
	return &TokenRule{&CommonRule{
		Field:     "token",
		FieldName: "token",
		Message:   "{0} invalid format!",
	}}
}
