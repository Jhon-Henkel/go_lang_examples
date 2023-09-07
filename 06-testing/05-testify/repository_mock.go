package tax

import "github.com/stretchr/testify/mock"

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) SaveTax(amount float64) error {
	args := m.Called(amount)
	return args.Error(0)
}