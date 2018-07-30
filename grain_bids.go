package ondemand

import (
	"fmt"
	"strings"
)

// GrainBids https://www.barchart.com/ondemand/api/getGrainBids
type GrainBids struct {
	Results []struct {
		Bids []struct {
			ID                   string      `json:"id"`
			Commodity            string      `json:"commodity"`
			Symbol               string      `json:"symbol"`
			DeliveryStart        string      `json:"delivery_start"`
			DeliveryEnd          string      `json:"delivery_end"`
			Basis                string      `json:"basis"`
			Notes                interface{} `json:"notes"`
			Active               bool        `json:"active"`
			SymRoot              string      `json:"sym_root"`
			CommodityID          string      `json:"commodity_id"`
			CustomerCommodityID  string      `json:"customer_commodity_id"`
			CommodityDisplayName string      `json:"commodity_display_name"`
			Unitvalue            int         `json:"unitvalue"`
			Unitweight           int         `json:"unitweight"`
			DeliveryMonth        string      `json:"deliveryMonth"`
			DeliveryYear         string      `json:"deliveryYear"`
			Basismonth           string      `json:"basismonth"`
			Timestamp            int         `json:"timestamp"`
			AsOf                 string      `json:"as_of"`
			Price                string      `json:"price"`
			Pricecwt             string      `json:"pricecwt"`
			Basiscwt             float64     `json:"basiscwt"`
			Pricetonne           string      `json:"pricetonne"`
			Basistonne           float64     `json:"basistonne"`
			Change               string      `json:"change"`
			Rawchange            float64     `json:"rawchange"`
			Pctchange            string      `json:"pctchange"`
			Cashprice            string      `json:"cashprice"`
			Cashpricetonne       string      `json:"cashpricetonne"`
			DeliverySort         string      `json:"delivery_sort"`
			DeliveryStartRaw     string      `json:"delivery_start_raw"`
			DeliveryEndRaw       string      `json:"delivery_end_raw"`
		} `json:"bids"`
		Distance       interface{} `json:"distance"`
		Company        string      `json:"company"`
		LocationID     int         `json:"locationId"`
		Location       string      `json:"location"`
		FacilityType   string      `json:"facility_type"`
		Address        string      `json:"address"`
		City           string      `json:"city"`
		State          string      `json:"state"`
		Lng            float64     `json:"lng"`
		Lat            float64     `json:"lat"`
		Phone          string      `json:"phone"`
		URL            string      `json:"url"`
		Zip            string      `json:"zip"`
		County         string      `json:"county"`
		CountyCode     string      `json:"county_code"`
		FipsCode       int         `json:"fips_code"`
		BasisTimestamp string      `json:"basisTimestamp"`
	} `json:"results"`
}

// GrainBids https://www.barchart.com/ondemand/api/getGrainBids
func (od *OnDemand) GrainBids(companyName string, fields []string, getAllBids bool, getAllLocations bool) (GrainBids, error) {
	bids := GrainBids{}

	_fields := ""

	if fields != nil {
		_fields = strings.Join(fields, ",")
	}

	_, err := od.Request("getGrainBids.json", fmt.Sprintf("companyName=%s&fields=%v&getAllBids=%t&getAllLocations=%t", companyName, _fields, getAllBids, getAllLocations), &bids)
	return bids, err
}
