// Get current Bitcoin exchange rate from Btc-e

package btce

import (
	"../bitcoin"
	"encoding/json"
)

// Gets current Bitcoin price from btc-e.com
type Btce struct {
	apiUrl string
	price  bitcoin.BitcoinPrice
}

// Ticker represents the JSON data returned from the API request
type T struct {
	Ticker map[string]float64
}

// Implements the Bitcion interface GetPrice method
func (b Btce) GetPrice(ch chan bitcoin.BitcoinPrice) {
	var t T
	b.apiUrl = "https://btc-e.com/api/2/btc_usd/ticker"

	// Request the current rate from the exchange
	content, err := bitcoin.GetContent(b.apiUrl)
	if err != nil {
		// Write error to the channel
		price := bitcoin.BitcoinPrice{
			Err:  err,
			Name: "Btc-e",
		}
		ch <- price

		return
	}

	// Decode the JSON data
	json.Unmarshal(content, &t)

	price := bitcoin.BitcoinPrice{
		CurBuy:  t.Ticker["last"],
		CurSell: t.Ticker["sell"],
		High:    t.Ticker["high"],
		Low:     t.Ticker["low"],
		Avg:     t.Ticker["avg"],
		Time:    int64(t.Ticker["updated"]),
		Vol:     t.Ticker["vol"],
		Name:    "Btc-e",
	}

	ch <- price
}
