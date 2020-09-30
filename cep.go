package brazil

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// NewCEPValidator creates a CEP number validator, containing inner mask and validation functions which can be
// accessed via Mask and Validate methods.
func NewCEPValidator() Validator {
	return &validator{
		maskFunc: func(value string) string {
			value = onlyDigits(value)

			return fmt.Sprintf("%s-%s", value[:5], value[5:])
		},
		validationFunc: func(value string) error {
			value = onlyDigits(value)

			if !regexp.MustCompile(`^\d{8}$`).MatchString(value) {
				return ErrInvalidCEPFormat
			}

			return nil
		},
	}
}

// RandomCEPNumber generates a random valid CEP number
func RandomCEPNumber() string {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)
	return strconv.Itoa(r.Intn(89999999) + 10000000)
}
