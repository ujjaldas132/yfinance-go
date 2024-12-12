package yfModels

type Stock struct {
	Meta struct {
		Currency             string `json:"currency"`
		Symbol               string `json:"symbol"`
		ExchangeName         string `json:"exchangeName"`
		InstrumentType       string `json:"instrumentType"`
		FirstTradeDate       int64  `json:"firstTradeDate"`
		GmtOffset            int64  `json:"gmtoffset"`
		Timezone             string `json:"timezone"`
		ExchangeTimezoneName string `json:"exchangeTimezoneName"`
		CurrentTradingPeriod struct {
			Pre     TradingPeriod `json:"Pre"`
			Regular TradingPeriod `json:"Regular"`
			Post    TradingPeriod `json:"Post"`
		}
		DataGranularity string   `json:"dataGranularity"`
		ValidRanges     []string `json:"validRanges"`
	}
	Timestamp  []int64 `json:"Timestamp"`
	Indicators struct {
		Quote []struct {
			Low    []float64 `json:"Low"`
			Volume []float64 `json:"Volume"`
			High   []float64 `json:"High"`
			Open   []float64 `json:"Open"`
			Close  []float64 `json:"Close"`
		}
		Unadjclose []struct {
			Unadjclose []float64 `json:"Unadjclose"`
		}
		Adjclose []struct {
			Adjclose []float64 `json:"Adjclose"`
		}
	}
}

type TradingPeriod struct {
	Timezone  string `json:"timezone"`
	End       int64  `json:"end"`
	Start     int64  `json:"start"`
	GmtOffset int64  `json:"gmtoffset"`
}
