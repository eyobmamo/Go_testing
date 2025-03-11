package main 

import "github.com/stretchr/testify/mock"

type MockPaymentGetway struct {
	mock.Mock
}

// charge is a mock method

func (m *MockPaymentGetway) Charge(amount float64)(string,error){
	args := m.Called(amount)
	return args.String(0),args.Error(1)
}
