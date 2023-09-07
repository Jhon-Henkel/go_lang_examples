package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &RepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Once()
	repository.On("SaveTax", 0.0).Return(errors.New("Invalid value"))
	repository.On("SaveTax", mock.Anything).Return(errors.New("Invalid value"))

	err := CalculateTaxAndSave(1000, repository)
	assert.Nil(t, err, "Error should be nil")

	err = CalculateTaxAndSave(0, repository)
	assert.NotNil(t, err, "Error should not be nil")
	assert.Error(t, err, "Error should be returned")
	assert.Equal(t, "Invalid value", err.Error(), "Error message should be 'Invalid value'")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}