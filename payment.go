package main

import "errors"

//paymentGateway defines the interface for a payment geteway

type PaymentGateway interface {
	Charge(amount float64) (string,error)


}

type RealPaymentGateway struct {}

func (g *RealPaymentGateway) Charge (amount float64) (string  ,error ){
	if amount <= 0{
		return "",errors.New("invalid amount")
	}
	return "Txn_12345",nil
}

// processPayment processes a payment useing the given PaymentGateway 

func ProcessPayment(geteway PaymentGateway,amount float64)(string ,error) {
	if amount < 0{
		return "",errors.New("invalid amount")
	}
	return geteway.Charge(amount)
}