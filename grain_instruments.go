package ondemand

import (
	"fmt"
)

// GrainInstruments https://www.barchart.com/ondemand/api/getGrainInstruments
type GrainInstruments struct {
	Results []struct {
		Symbol            string `json:"symbol"`
		SymbolDescription string `json:"symbolDescription"`
		Commodity         string `json:"commodity"`
		UnderlyingFuture  string `json:"underlyingFuture"`
		DeliveryEnd       string `json:"deliveryEnd"`
		State             string `json:"state"`
		County            string `json:"county"`
		IndexGroup        string `json:"indexGroup"`
		Country           string `json:"country"`
		Region            string `json:"region"`
		CountyFipsCode    string `json:"countyFipsCode"`
		DistrictCode      string `json:"districtCode"`
		StateFipsCode     string `json:"stateFipsCode"`
	} `json:"results"`
}

// GrainInstruments https://www.barchart.com/ondemand/api/getGrainInstruments
func (od *OnDemand) GrainInstruments(params string) (GrainInstruments, error) {
	instruments := GrainInstruments{}

	_, err := od.Request("getGrainInstruments.json", fmt.Sprintf("%s", params), &instruments)
	return instruments, err
}
