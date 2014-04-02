// Gets current Bitcoin exchange rate
package main

import (
	"./bitcoin"
	"./bitfinex"
	"./bitstamp"
	"./btce"
	"./coinbase"
	"fmt"
	"time"
)

// Print the current price
func pt(p bitcoin.BitcoinPrice) {
	var t time.Time
	if p.Time > 0 {
		t = time.Unix(p.Time, 0)
	}
	fmt.Printf("\n%s (%v)\nCur: %14.4f\nSell:%14.4f\nHigh:%14.4f\nLow: %14.4f\nAvg: %14.4f\nVol: %14.4f\n\n",
		p.Name, t, p.CurBuy, p.CurSell, p.High, p.Low, p.Avg, p.Vol)
}
func main() {

	// Get current exchange rate from Bitstamp
	b := bitcoin.Bitcoin(bitstamp.Bitstamp{})
	p, err := b.GetPrice()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	pt(p)

	// Get current exchange rate from Btce
	b = bitcoin.Bitcoin(btce.Btce{})
	p, err = b.GetPrice()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	pt(p)

	// Get current exchange rate from  Coinbase
	b = bitcoin.Bitcoin(coinbase.Coinbase{})
	p, err = b.GetPrice()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	pt(p)

	// Get current exchange rate from Bitfinex
	b = bitcoin.Bitcoin(bitfinex.Bitfinex{})
	p, err = b.GetPrice()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	pt(p)
}
