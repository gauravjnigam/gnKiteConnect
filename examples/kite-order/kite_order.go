package main

import (
	"fmt"
	"gnKiteConnect/examples/setup"

	kiteconnect "github.com/zerodhatech/gokiteconnect"
)

func main() {

	kc := setup.SetupKiteConnection()
	limitPrice := 160.0
	triggerPrice := limitPrice - 0.10 // TODO : Gaurav - configurable?

	createSellSLOrder(kc, limitPrice, triggerPrice, "NIFTY22AUG17600PE")

}

func createSellSLOrder(kc *kiteconnect.Client, limitPrice float64, triggerPrice float64, tradingsymbol string) string {
	orderParam := kiteconnect.OrderParams{
		Exchange:        kiteconnect.ExchangeNFO,
		Tradingsymbol:   tradingsymbol,
		TransactionType: kiteconnect.TransactionTypeSell,
		OrderType:       kiteconnect.OrderTypeSL,
		Product:         kiteconnect.ProductNRML,
		Quantity:        50,
		Price:           limitPrice,
		TriggerPrice:    triggerPrice,
		Tag:             "gnAlgoSELLSLOrder",
	}

	fmt.Printf("SELL Order param : %v\n", orderParam)

	fmt.Printf("Creating SELL ORDER (orderPrice : %f, triggerPrice : %f) : %v", limitPrice, triggerPrice, orderParam)
	orderResp, err := kc.PlaceOrder(kiteconnect.VarietyAMO, orderParam)
	if err != nil {
		fmt.Printf("Error while placing SELL order - %v\n", err)
		return ""
	} else {
		fmt.Printf("SELL Order created  : %v", orderResp)
	}

	fmt.Printf("SELL Order Resp : %v\n", orderResp)
	return orderResp.OrderID

}
