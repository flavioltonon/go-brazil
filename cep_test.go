package brazil_test

import (
	"testing"

	. "github.com/flavioltonon/go-brazil"

	"github.com/stretchr/testify/assert"
)

func TestCEP(t *testing.T) {
	c := NewCEPValidator()

	t.Run("When parsing a CEP with an valid number, no errors should be returned and cep should contain a value", func(t *testing.T) {
		assert.NoError(t, c.Validate("12345-678"))
	})

	t.Run("When parsing a CEP with an valid masked number, no errors should be returned and cep should contain a value", func(t *testing.T) {
		assert.NoError(t, c.Validate("12345678"))
	})

	t.Run("When parsing a CEP with an invalid input number, an ErrInvalidCEPFormat should be returned and cep should be nil", func(t *testing.T) {
		assert.Error(t, c.Validate("123456789"), ErrInvalidCEPFormat)
	})

	t.Run("When a random CEP number is generated, it should be successfully validated", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			value := RandomCEPNumber()
			assert.NoError(t, c.Validate(value))
		}
	})
}
