package yfinance

import (
	"github.com/ujjaldas132/yfinance-go/models/yfModels"
	"log"
	"testing"
)

func TestGetStock(t *testing.T) {

	symbol := "bse.ns"
	yf := YFinance{FullTickerSymbol: symbol}

	stock, err := yf.GetHistory(yfModels.RangeFiveDays, yfModels.IntervalOneDay)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(stock)

}
