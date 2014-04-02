package coinbase

import (
	"../bitcoin"
	"encoding/json"
	"strconv"
)

type Coinbase struct {
	apiUrl string
	price  bitcoin.BitcoinPrice
}

type T struct {
	Subtotal map[string]float64
	Fees     map[string]float64
	Total    map[string]float64
	Amount   string
	Currency string
}

func (b Coinbase) GetPrice() (bitcoin.BitcoinPrice, error) {
	var t T
	b.apiUrl = "https://coinbase.com/api/v1/prices/buy"

	// Request the current rate from the exchange
	content, err := bitcoin.GetContent(b.apiUrl)
	if err != nil {
		return bitcoin.BitcoinPrice{}, err
	}

	// Decode the JSON data
	json.Unmarshal(content, &t)
	curBuy, _ := strconv.ParseFloat(t.Amount, 64)

	b.apiUrl = "https://coinbase.com/api/v1/prices/sell"
	content, err = bitcoin.GetContent(b.apiUrl)
	if err != nil {
		return bitcoin.BitcoinPrice{}, err
	}

	json.Unmarshal(content, &t)
	curSell, _ := strconv.ParseFloat(t.Amount, 64)

	return bitcoin.BitcoinPrice{
		CurBuy:  curBuy,
		CurSell: curSell,
		Name:    "Coinbase",
	}, nil

}
