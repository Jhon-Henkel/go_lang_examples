package tax

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 10.0, tax, "Tax should be 10.0")

	tax, err = CalculateTax(0)
	assert.NotNil(t, err, "Error should not be nil")
	assert.Error(t, err, "Error should be returned")
	assert.Equal(t, 0.0, tax, "Tax should be 0.0")
	assert.Equal(t, "amount must be greater than 0", err.Error(), "Error message should be 'Invalid value'")
}