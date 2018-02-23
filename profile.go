package ondemand

import (
	"fmt"
	"strings"
)

type Profile struct {
	Results []struct {
		Symbol               string   `json:"symbol"`
		SymbolCode           string   `json:"symbolCode"`
		Exchange             string   `json:"exchange"`
		ExchangeName         string   `json:"exchangeName"`
		SicSector            string   `json:"sicSector"`
		Industry             string   `json:"industry"`
		SubIndustry          string   `json:"subIndustry"`
		IndexMembership      []string `json:"indexMembership"`
		BusinessSummary      string   `json:"businessSummary"`
		Ceo                  string   `json:"ceo"`
		QtrOneEarnings       float64  `json:"qtrOneEarnings"`
		QtrOneEarningsDate   string   `json:"qtrOneEarningsDate"`
		QtrTwoEarnings       float64  `json:"qtrTwoEarnings"`
		QtrTwoEarningsDate   string   `json:"qtrTwoEarningsDate"`
		QtrThreeEarnings     float64  `json:"qtrThreeEarnings"`
		QtrThreeEarningsDate string   `json:"qtrThreeEarningsDate"`
		QtrFourEarnings      float64  `json:"qtrFourEarnings"`
		QtrFourEarningsDate  string   `json:"qtrFourEarningsDate"`
		PeRatio              string   `json:"peRatio"`
		EpsGrowth            string   `json:"epsGrowth"`
		RecentEarnings       string   `json:"recentEarnings"`
		AnnualEPS            string   `json:"annualEPS"`
		Address              string   `json:"address"`
		City                 string   `json:"city"`
		State                string   `json:"state"`
		Country              string   `json:"country"`
		ZipCode              string   `json:"zipCode"`
		PhoneNumber          string   `json:"phoneNumber"`
		InstrumentType       string   `json:"instrumentType"`
	}
}

//

func (od *OnDemand) Profile(symbols []string, fields []string) (Profile, error) {
	profile := Profile{}
	_symbols := strings.ToUpper(strings.Join(symbols, ","))
	_fields := ""

	if fields != nil {
		_fields = strings.Join(fields, ",")
	}

	_, err := od.Request("getProfile.json", fmt.Sprintf("symbols=%v&fields=%s", _symbols, _fields), &profile)
	return profile, err
}
