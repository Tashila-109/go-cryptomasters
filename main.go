package main

import "cryptomasters/go/crypto/api"

func main() {
	rate, err := api.GetRate("BTC")

	print(rate, err)
}
