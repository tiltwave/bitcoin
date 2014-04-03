// Get current Bitcoin price from multiple exchanges
// Author: github.com/tiltwave

package main

import (
	"./bitcoin"
	"./bitfinex"
	"./bitstamp"
	"./btce"
	"./coinbase"
	"fmt"
)

func main() {

	ch := make(chan bitcoin.BitcoinPrice, 4)

	// Get current exchange rate from Bitstamp
	go bitcoin.Bitcoin(bitstamp.Bitstamp{}).GetPrice(ch)

	// Get current exchange rate from Coinbase
	go bitcoin.Bitcoin(coinbase.Coinbase{}).GetPrice(ch)

	// Get current exchange rate from Bitfinex
	go bitcoin.Bitcoin(bitfinex.Bitfinex{}).GetPrice(ch)

	// Get current exchange rate from Btce
	go bitcoin.Bitcoin(btce.Btce{}).GetPrice(ch)

	// Read from the channel and print the price
	ph()
	for i := 0; i < 4; i++ {
		pp(<-ch)
	}
}

// Print the current price
func pp(p bitcoin.BitcoinPrice) {
	if p.Err != nil {
		fmt.Printf("%10s | Error: %s\n", p.Name, p.Err.Error())

		return
	}
	fmt.Printf("%10s |%20.4f| %13.4f| %13.4f| %13.4f| %13.4f| %14.4f|\n",
		p.Name, p.CurBuy, p.CurSell, p.High, p.Low, p.Avg, p.Vol)
}

// Print header
func ph() {
	fmt.Printf("Getting data... (make sure your terminal is wide enough the see each line\n\n")
	fmt.Printf("%10s |%20s| %13s| %13s| %13s| %13s| %14s| \n", "Name", "Buy Price", "Sell Price", "High", "Low", "Avg", "Vol")
	fmt.Printf("%10s |%20s| %13s| %13s| %13s| %13s| %14s| \n", "", "", "", "", "", "", "")
}
