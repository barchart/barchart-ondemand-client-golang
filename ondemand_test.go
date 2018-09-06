package ondemand

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var key = os.Getenv("ONDEMAND_KEY")
var od = New(key, true)

func TestInit(t *testing.T) {

	if len(key) == 0 {
		log.Fatal("Set ONDEMAND_KEY env var")
	}

	// get a quote
	//
	quote, err := od.Quote([]string{"AAPL", "ESH8"}, []string{"impliedVolatility"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quote.Results[1].ImpliedVolatility)
	fmt.Println(quote.Results[0].ImpliedVolatility)

	t.Log(err, quote)

	// get an instrument profile
	//
	profile, err := od.Profile([]string{"NEM"}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profile.Results[0].SymbolCode)

	// get some News for AAPL
	//
	news, err := od.News("headline", []string{"AP"}, []string{"AAPL"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(news.Results[0].Headline)

	// get grain bids
	//
	bids, _ := od.GrainBids("Heartland Coop", []string{"county", "fips_code"}, true, true)
	fmt.Println("Bids", bids.Results[0].Bids[0].Basis)

	// get grain instruments
	//
	insts, _ := od.GrainInstruments("instrumentType=bidBasis&instrumentRoot=ZC&zipCode=50152&deliveryMonth=201812")
	fmt.Println("Grain Insts", insts)

	// get financial highlights
	//
	fh, _ := od.FinancialHighlights([]string{"GE", "TSLA"}, []string{"beta"})
	fmt.Println("Financial Highlights", fh)
}
