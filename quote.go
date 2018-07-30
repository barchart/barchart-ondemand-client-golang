package ondemand

import (
	"fmt"
	"strings"
)

// Quote https://www.barchart.com/ondemand/api/getQuote
type Quote struct {
	Results []struct {
		Symbol                 string      `json:"symbol"`
		Name                   string      `json:"name"`
		DayCode                string      `json:"dayCode"`
		ServerTimestamp        string      `json:"serverTimestamp"`
		Mode                   string      `json:"mode"`
		LastPrice              float64     `json:"lastPrice"`
		TradeSize              int         `json:"tradeSize"`
		TradeTimestamp         string      `json:"tradeTimestamp"`
		NetChange              float64     `json:"netChange"`
		PercentChange          float64     `json:"percentChange"`
		Tick                   string      `json:"tick"`
		PreviousLastPrice      float64     `json:"previousLastPrice"`
		PreviousTimestamp      string      `json:"previousTimestamp"`
		Bid                    float64     `json:"bid"`
		BidSize                int         `json:"bidSize"`
		Ask                    float64     `json:"ask"`
		AskSize                int         `json:"askSize"`
		UnitCode               string      `json:"unitCode"`
		Open                   float64     `json:"open"`
		High                   float64     `json:"high"`
		Low                    float64     `json:"low"`
		Close                  float64     `json:"close"`
		NumTrades              int         `json:"numTrades"`
		DollarVolume           float64     `json:"dollarVolume"`
		Flag                   string      `json:"flag"`
		PreviousOpen           float64     `json:"previousOpen"`
		PreviousHigh           float64     `json:"previousHigh"`
		PreviousLow            float64     `json:"previousLow"`
		PreviousClose          float64     `json:"previousClose"`
		PreviousSettlement     float64     `json:"previousSettlement"`
		Volume                 int         `json:"volume"`
		PreviousVolume         int         `json:"previousVolume"`
		OpenInterest           interface{} `json:"openInterest"`
		FiftyTwoWkHigh         float64     `json:"fiftyTwoWkHigh"`
		FiftyTwoWkHighDate     string      `json:"fiftyTwoWkHighDate"`
		FiftyTwoWkLow          float64     `json:"fiftyTwoWkLow"`
		FiftyTwoWkLowDate      string      `json:"fiftyTwoWkLowDate"`
		AvgVolume              int         `json:"avgVolume"`
		SharesOutstanding      int         `json:"sharesOutstanding"`
		DividendRateAnnual     string      `json:"dividendRateAnnual"`
		DividendYieldAnnual    string      `json:"dividendYieldAnnual"`
		ExDividendDate         string      `json:"exDividendDate"`
		ImpliedVolatility      interface{} `json:"impliedVolatility"`
		TwentyDayAvgVol        int         `json:"twentyDayAvgVol"`
		Month                  interface{} `json:"month"`
		Year                   interface{} `json:"year"`
		ExpirationDate         interface{} `json:"expirationDate"`
		LastTradingDay         interface{} `json:"lastTradingDay"`
		TwelveMnthPct          float64     `json:"twelveMnthPct"`
		AverageWeeklyVolume    int         `json:"averageWeeklyVolume"`
		AverageMonthlyVolume   int         `json:"averageMonthlyVolume"`
		AverageQuarterlyVolume int         `json:"averageQuarterlyVolume"`
		ExchangeMargin         interface{} `json:"exchangeMargin"`
		OneMonthHigh           float64     `json:"oneMonthHigh"`
		OneMonthHighDate       string      `json:"oneMonthHighDate"`
		OneMonthLow            float64     `json:"oneMonthLow"`
		OneMonthLowDate        string      `json:"oneMonthLowDate"`
		ThreeMonthHigh         float64     `json:"threeMonthHigh"`
		ThreeMonthHighDate     string      `json:"threeMonthHighDate"`
		ThreeMonthLow          float64     `json:"threeMonthLow"`
		ThreeMonthLowDate      string      `json:"threeMonthLowDate"`
		SixMonthHigh           float64     `json:"sixMonthHigh"`
		SixMonthHighDate       string      `json:"sixMonthHighDate"`
		SixMonthLow            float64     `json:"sixMonthLow"`
		SixMonthLowDate        string      `json:"sixMonthLowDate"`
		Exchange               string      `json:"exchange"`
	} `json:"results"`
}

// Quote https://www.barchart.com/ondemand/api/getQuote
func (od *OnDemand) Quote(symbols []string, fields []string) (Quote, error) {
	quote := Quote{}

	_symbols := strings.ToUpper(strings.Join(symbols, ","))
	_fields := "histVol100"

	if fields != nil {
		_fields = strings.Join(fields, ",")
	}

	_, err := od.Request("getQuote.json", fmt.Sprintf("symbols=%v&fields=%s", _symbols, _fields), &quote)
	return quote, err
}
