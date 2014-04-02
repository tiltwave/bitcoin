package bitcoin

import (
	"io/ioutil"
	"net/http"
)

type BitcoinPrice struct {
	CurBuy  float64 // Last BTC price
	CurSell float64 // Lowest sell order
	Bid     float64 // Highest buy order
	High    float64 // Last 24 hours high price
	Low     float64 // Last 24 hours low price
	Avg     float64 // Last 24 hours avg price
	Time    int64   // Provider time stamp
	Vol     float64 // Volume
	Name    string  // Name of the exchange
}

type Bitcoin interface {
	GetPrice() (BitcoinPrice, error)
}

// Makes an API request
func GetContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
