// Get current Bitcoin exchange rate from Coinbase

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

// Implements the Bitcion interface GetPrice method
func (b Coinbase) GetPrice(ch chan bitcoin.BitcoinPrice) {
	var t T
	b.apiUrl = "https://coinbase.com/api/v1/prices/buy"

	// Request the current rate from the exchange
	content, err := bitcoin.GetContent(b.apiUrl)
	if err != nil {
		// Write error to the channel
		price := bitcoin.BitcoinPrice{
			Err:  err,
			Name: "Coinbase",
		}
		ch <- price

		return
	}

	// Decode the JSON data
	json.Unmarshal(content, &t)
	curBuy, _ := strconv.ParseFloat(t.Amount, 64)

	b.apiUrl = "https://coinbase.com/api/v1/prices/sell"
	content, err = bitcoin.GetContent(b.apiUrl)
	if err != nil {
		// Write error to the channel
		price := bitcoin.BitcoinPrice{
			Err:  err,
			Name: "Coinbase",
		}
		ch <- price

		return
	}

	json.Unmarshal(content, &t)
	curSell, _ := strconv.ParseFloat(t.Amount, 64)

	price := bitcoin.BitcoinPrice{
		CurBuy:  curBuy,
		CurSell: curSell,
		Name:    "Coinbase",
	}

	ch <- price
}
