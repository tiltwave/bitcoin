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

type T struct {
	Ticker map[string]float64
}

// Implements GetPrice
func (b Btce) GetPrice() (bitcoin.BitcoinPrice, error) {
	var t T
	b.apiUrl = "https://btc-e.com/api/2/btc_usd/ticker"
	content, err := bitcoin.GetContent(b.apiUrl)
	if err != nil {
		return bitcoin.BitcoinPrice{}, err
	}

	json.Unmarshal(content, &t)

	return bitcoin.BitcoinPrice{
		CurBuy:  t.Ticker["last"],
		CurSell: t.Ticker["sell"],
		High:    t.Ticker["high"],
		Low:     t.Ticker["low"],
		Avg:     t.Ticker["avg"],
		Time:    int64(t.Ticker["updated"]),
		Vol:     t.Ticker["vol"],
		Name:    "Btc-e",
	}, nil
}
