package yfinance

import (
	"github.com/ujjaldas132/yfinance-go/models/yfModels"
	"log"
	"sync"
	"testing"
	"time"
)

func TestGetStock(t *testing.T) {
	symbol := "bse.ns"
	start := time.Now()
	yf := YFinance{FullTickerSymbol: symbol}
	stock, err := yf.GetHistory(yfModels.RangeFiveDays, yfModels.IntervalOneDay)
	duration := time.Since(start)
	t.Logf("Duration: %v", duration)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(stock)
}

func TestMultipleGetStock(t *testing.T) {
	symbols := []string{"bse.ns", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft"}
	start := time.Now()
	for _, symbol := range symbols {
		k := YFinance{FullTickerSymbol: symbol}
		k.GetHistory(yfModels.RangeFiveDays, yfModels.IntervalOneDay)
	}
	elapsed := time.Since(start)
	t.Logf("without goroutine multistock GetHistory took %s\n", elapsed)
}

func TestGetMutiStockWithGoroutine(t *testing.T) {
	symbols := []string{"bse.ns", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft", "aapl", "msft"}
	start := time.Now()
	var wg sync.WaitGroup
	for _, symbol := range symbols {
		wg.Add(1)
		go func(symbol string) {
			defer wg.Done()
			k := YFinance{FullTickerSymbol: symbol}
			k.GetHistory(yfModels.RangeOneDay, yfModels.IntervalOneMinute)
		}(symbol)
	}
	wg.Wait()
	elapsed := time.Since(start)
	t.Logf("with goroutine multistock GetHistory took %s\n", elapsed)
}
