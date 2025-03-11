package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"

)
func TestProcessPayment(t *testing.T){
	//create mock payment getway 

	mockGateway := new(MockPaymentGetway)

	//Deine the expected behavior 
	mockGateway.On("Charge",100.0).Return("Txt_23424",nil)


	mockGateway.On("Charge",0.0).Return("",errors.New("invalid amount"))// Error case

	transactionID ,err := ProcessPayment(mockGateway,100.0)
	assert.Equal(t,"Txt_23424",transactionID)
	assert.NoError(t,err)

	transactionID ,err = ProcessPayment(mockGateway,0.0)
	assert.Equal(t,"",transactionID)
	assert.Error(t,err)

	mockGateway.AssertCalled(t,"Charge",100.0)
	mockGateway.AssertCalled(t,"Charge",0.0)
	mockGateway.AssertExpectations(t)

}