// Get current Bitcoin exchange rate from Bitfinex
package bitfinex

import (
	"../bitcoin"
	"encoding/json"
	"strconv"
)

type Bitfinex struct {
	apiUrl string
	price  bitcoin.BitcoinPrice
}

// Ticker represents the JSON data returned from the API request
type T struct {
	Mid        string
	Bid        string
	Ask        string
	Last_price string
	timestamp  string
}

// T2 represents the JSON data returned from the API request
type T2 struct {
	Low    string
	High   string
	Volume string
}

func (b Bitfinex) GetPrice() (bitcoin.BitcoinPrice, error) {
	var t T
	var t2 T2

	b.apiUrl = "https://api.bitfinex.com/v1/ticker/btcusd"

	// Request the current rate from the exchange
	content, err := bitcoin.GetContent(b.apiUrl)
	if err != nil {
		return bitcoin.BitcoinPrice{}, err
	}

	// Decode the JSON data
	json.Unmarshal(content, &t)
	curBuy, _ := strconv.ParseFloat(t.Last_price, 64)
	curSell, _ := strconv.ParseFloat(t.Ask, 64)
	mid, _ := strconv.ParseFloat(t.Mid, 64)

	b.apiUrl = "https://api.bitfinex.com/v1/today/btcusd"
	content, err = bitcoin.GetContent(b.apiUrl)
	if err != nil {
		return bitcoin.BitcoinPrice{}, err
	}

	json.Unmarshal(content, &t2)
	high, _ := strconv.ParseFloat(t2.High, 64)
	vol, _ := strconv.ParseFloat(t2.Volume, 64)
	low, _ := strconv.ParseFloat(t2.Low, 64)

	return bitcoin.BitcoinPrice{
		CurBuy:  curBuy,
		CurSell: curSell,
		High:    high,
		Low:     low,
		Avg:     mid,
		Vol:     vol,
		Name:    "Bitfinex",
	}, nil

}
