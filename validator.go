package brazil

type Validator interface {
	Validate(value string) error
}

type Maskable interface {
	Mask(value string) string
}

type MaskableValidator interface {
	Maskable
	Validator
}

type validator struct {
	maskFunc       func(value string) string
	validationFunc func(value string) error
}

// Validate checks if the input value is valid according to the validator rules and returns an error if it isn't.
func (v *validator) Validate(value string) error {
	return v.validationFunc(value)
}

// Mask applies the validator mask to a value. If the value is not valid, the masking may not be completely correct and the method may panic.
func (v *validator) Mask(value string) string {
	return v.maskFunc(value)
}
