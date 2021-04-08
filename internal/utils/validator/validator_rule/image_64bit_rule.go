package validatorrule

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Image64BitRule ...
type Image64BitRule struct{ *CommonRule }

// GetRule ...
func (*Image64BitRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().String()
		if data != "" {
			coI := strings.Index(data, ",")
			if coI == -1 {
				return false
			}
			rawImage := data[coI+1:]
			matched, err := regexp.Match(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`, []byte(rawImage))
			if err != nil {
				return false
			}
			return matched
		}
		return true
	}
}

// NewImage64BitRule ...
func NewImage64BitRule() *Image64BitRule {
	return &Image64BitRule{&CommonRule{
		Field:     "image_64_bit",
		FieldName: "Image 64 Bit",
		Message:   "{0} invalid format!",
	}}
}
