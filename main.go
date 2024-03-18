package main

import (
	"cryptomasters/go/crypto/api"
	"fmt"
	"time"
)

func main() {
	go getCurrencyData("BTC")
	go getCurrencyData("ETH")
	go getCurrencyData("BCH")
	time.Sleep(time.Second * 5)
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
