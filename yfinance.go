package yfinace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ujjaldas132/yfinance-go/models"
	"github.com/ujjaldas132/yfinance-go/models/yfModels"
	"net/http"
	"text/template"
)

const (
	yahooFinanceAPI = "https://query1.finance.yahoo.com/v8/finance/chart/{{.Symbol}}?range={{.Range}}&includePrePost=false&interval={{.Interval}}&corsDomain=finance.yahoo.com&.tsrc=finance"
)

var (
	// DebugLogging enables verbose output from the yf package when set to true.
	DebugLogging = false
)

type YFinance struct {
	// fullTickerSymbol is the full ticker symbol for the stock, e.g. "BSE.NS"
	fullTickerSymbol string
}

// Get Stock price history data
func (yf YFinance) GetHistory(dateRange string, interval string) ([]models.HistoryData, error) {
	stock, err := getStock(yf.fullTickerSymbol, dateRange, interval)
	if err != nil {
		fmt.Println(err)
	}

	length := len(stock.Timestamp)

	var response []models.HistoryData

	for i := 0; i < length; i++ {
		response = append(response, models.HistoryData{
			Datetime: stock.Timestamp[i],
			Open:     stock.Indicators.Quote[0].Open[i],
			High:     stock.Indicators.Quote[0].High[i],
			Low:      stock.Indicators.Quote[0].Low[i],
			Close:    stock.Indicators.Quote[0].Close[i],
			Volume:   stock.Indicators.Quote[0].Volume[i],
			AdjClose: stock.Indicators.Adjclose[0].Adjclose[i],
		})
	}
	return response, nil

}

// GetStock returns stock details for a particular symbol from the Yahoo Finance API.
func getStock(fullTickerSymbol, dateRange, interval string) (*yfModels.Stock, error) {
	resp, err := doRequest(fullTickerSymbol, dateRange, interval)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected StatusCode: %d", resp.StatusCode)
	}

	var target struct {
		Chart struct {
			Result []yfModels.Stock
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, err
	}

	if len(target.Chart.Result) != 1 {
		return nil, fmt.Errorf("Unexpected response, %d results returned, API change must have occurred", len(target.Chart.Result))
	}

	return &target.Chart.Result[0], nil
}

func doRequest(symbol, dateRange, interval string) (*http.Response, error) {
	tmpl, err := template.New("YF-API").Parse(yahooFinanceAPI)
	if err != nil {
		return nil, err
	}

	p := struct {
		Symbol   string
		Range    string
		Interval string
	}{
		Symbol:   symbol,
		Range:    dateRange,
		Interval: interval,
	}
	var url bytes.Buffer
	if err := tmpl.Execute(&url, p); err != nil {
		return nil, err
	}

	debug(url.String())
	return http.Get(url.String())
}

func debug(str string) {
	if !DebugLogging {
		return
	}

	fmt.Println(str)
}
