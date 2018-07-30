# A Golang Client for Barchart OnDemand

A simple (work in progress) Go client for [Barchart OnDemand](http://barchartondemand.com/api.php)

## Install

```
go get github.com/barchart/barchart-ondemand-client-golang
```

### Get Market Data for Apple and Exelon

```go
import (
    "fmt"
    od "github.com/barchart/barchart-ondemand-client-golang"
)

var ondemand = od.New("YOUR_API_KEY", false)

quote, err := ondemand.Quote([]string{"AAPL", "EXC"}, []string{"bid", "ask"})

if err != nil {
    fmt.Sprintf("Symbol: %v Last: %v", q.Symbol, q.LastPrice)
}
```

### Generic API call by name

```go
var result interface{}
ondemand.Request("getQuote.json", "symbols=AAPL", &result)

fmt.Println("Result:", result)
```

### Get AP News Headlines for Apple

```go
news, err := ondemand.News("headline", []string{"AP"}, []string{"AAPL"})

if err != nil {
    log.Fatal(err)
}

fmt.Println(news.Results[0].Headline)
```

## Development

```
# building

> go build

# run tests

> SET ONDEMAND_KEY=YOUR_API_KEY
> go test -timeout 30s -run ^TestInit$
```