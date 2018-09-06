package ondemand

import (
	"fmt"
	"strings"
)

// FinancialHighlights Response: https://www.barchart.com/ondemand/api/getFinancialHighlights
type FinancialHighlights struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Results []struct {
		Symbol                    string      `json:"symbol"`
		MarketCapitalization      int         `json:"marketCapitalization"`
		InsiderShareholders       interface{} `json:"insiderShareholders"`
		AnnualRevenue             int         `json:"annualRevenue"`
		TtmRevenue                interface{} `json:"ttmRevenue"`
		SharesOutstanding         int         `json:"sharesOutstanding"`
		InstitutionalShareholders float64     `json:"institutionalShareholders"`
		AnnualNetIncome           int         `json:"annualNetIncome"`
		TtmNetIncome              int         `json:"ttmNetIncome"`
		TtmNetProfitMargin        float64     `json:"ttmNetProfitMargin"`
		OneYearReturn             float64     `json:"oneYearReturn"`
		ThreeYearReturn           float64     `json:"threeYearReturn"`
		FiveYearReturn            float64     `json:"fiveYearReturn"`
		FiveYearRevenueGrowth     float64     `json:"fiveYearRevenueGrowth"`
		FiveYearEarningsGrowth    float64     `json:"fiveYearEarningsGrowth"`
		FiveYearDividendGrowth    float64     `json:"fiveYearDividendGrowth"`
		LastQtrEPS                float64     `json:"lastQtrEPS"`
		AnnualEPS                 float64     `json:"annualEPS"`
		TtmEPS                    float64     `json:"ttmEPS"`
		AnnualDividendRate        float64     `json:"annualDividendRate"`
		AnnualDividendYield       float64     `json:"annualDividendYield"`
		TwelveMonthEPSChg         float64     `json:"twelveMonthEPSChg"`
		PeRatio                   float64     `json:"peRatio"`
		RecentEarnings            float64     `json:"recentEarnings"`
		RecentDividends           float64     `json:"recentDividends"`
		RecentSplit               string      `json:"recentSplit"`
		Beta                      float64     `json:"beta"`
		WeightAlpha               float64     `json:"weightAlpha"`
	} `json:"results"`
}

// FinancialHighlights https://www.barchart.com/ondemand/api/getFinancialHighlights
func (od *OnDemand) FinancialHighlights(symbols []string, fields []string) (FinancialHighlights, error) {
	fhighlights := FinancialHighlights{}

	_symbols := strings.ToUpper(strings.Join(symbols, ","))
	_fields := "peRatio"

	if fields != nil {
		_fields = strings.Join(fields, ",")
	}

	_, err := od.Request("getFinancialHighlights.json", fmt.Sprintf("symbols=%v&fields=%s", _symbols, _fields), &fhighlights)
	return fhighlights, err
}
