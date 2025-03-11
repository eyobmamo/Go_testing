package main

import(
	"fmt"
)

func main(){
	gateway := &RealPaymentGateway{}
	transactionID,err := ProcessPayment(gateway,100.0)
	if err != nil {
		fmt.Println("payment failed:",err)
	} else {
		fmt.Println("payment successful! Transaton ID:",transactionID)
	}
}