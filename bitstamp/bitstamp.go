package bitstamp

import (
	"../bitcoin"
	"encoding/json"
	"strconv"
)

type Bitstamp struct {
	apiUrl string
	price  bitcoin.BitcoinPrice
}

type T struct {
	High      string
	Last      string
	Timestamp string
	Bid       string
	Vwap      string
	Volume    string
	Low       string
	Ask       string
}

func (b Bitstamp) GetPrice() (bitcoin.BitcoinPrice, error) {
	var t T
	b.apiUrl = "https://www.bitstamp.net/api/ticker/"
	content, err := bitcoin.GetContent(b.apiUrl)
	if err != nil {
		return bitcoin.BitcoinPrice{}, err
	}

	json.Unmarshal(content, &t)
	high, _ := strconv.ParseFloat(t.High, 64)
	last, _ := strconv.ParseFloat(t.Last, 64)
	time, _ := strconv.ParseInt(t.Timestamp, 10, 64)
	bid, _ := strconv.ParseFloat(t.Bid, 64)
	vwap, _ := strconv.ParseFloat(t.Vwap, 64)
	vol, _ := strconv.ParseFloat(t.Volume, 64)
	low, _ := strconv.ParseFloat(t.Low, 64)
	ask, _ := strconv.ParseFloat(t.Ask, 64)

	return bitcoin.BitcoinPrice{
		CurBuy:  last,
		CurSell: ask,
		Bid:     bid,
		High:    high,
		Low:     low,
		Avg:     vwap,
		Time:    time,
		Vol:     vol,
		Name:    "Bitstamp",
	}, nil

}
