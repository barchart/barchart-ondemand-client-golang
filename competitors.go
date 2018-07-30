package ondemand

import (
	"fmt"
	"strings"
)

// Competitors https://www.barchart.com/ondemand/api/getCompetitors
type Competitors struct {
	Results []struct {
		Symbol             string  `json:"symbol"`
		Name               string  `json:"name"`
		MarketCap          int64   `json:"marketCap"`
		FiftyTwoWkHigh     float64 `json:"fiftyTwoWkHigh"`
		FiftyTwoWkHighDate string  `json:"fiftyTwoWkHighDate"`
		FiftyTwoWkLow      float64 `json:"fiftyTwoWkLow"`
		FiftyTwoWkLowDate  string  `json:"fiftyTwoWkLowDate"`
	}
}

// Competitors https://www.barchart.com/ondemand/api/getCompetitors
func (od *OnDemand) Competitors(symbol string, fields []string) (Competitors, error) {
	competitors := Competitors{}

	_symbol := strings.ToUpper(symbol)
	_fields := ""

	if fields != nil {
		_fields = strings.Join(fields, ",")
	}

	_, err := od.Request("getCompetitors.json", fmt.Sprintf("symbol=%v&fields=%s", _symbol, _fields), &competitors)
	return competitors, err
}
