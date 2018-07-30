package ondemand

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type OnDemand struct {
	BaseURL string
	APIKey  string
	Debug   bool
}

func New(apiKey string, debug bool) (od *OnDemand) {
	od = &OnDemand{
		BaseURL: "http://ondemand.websol.barchart.com/",
		APIKey:  apiKey,
		Debug:   debug,
	}

	return od
}

func (od *OnDemand) Request(api string, params string, result interface{}) ([]byte, error) {
	ep := fmt.Sprintf("%v%v?apikey=%v&%v", od.BaseURL, api, od.APIKey, params)

	if od.Debug == true {
		fmt.Printf("Going After: %v\n", ep)
	}

	body, err := od.get(ep)
	err = json.Unmarshal(body, result)

	//log.Fatal(string(body[:]))

	return body, err
}

func (od *OnDemand) get(url string) (body []byte, err error) {
	client := http.Client{
		Timeout: time.Duration(30 * time.Second),
	}

	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}
