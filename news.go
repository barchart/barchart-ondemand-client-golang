package ondemand

import (
	"fmt"
	"strings"
)

type Headline struct {
	Results []struct {
		NewsID        int         `json:"newsID"`
		Timestamp     string      `json:"timestamp"`
		Source        string      `json:"source"`
		Categories    []string    `json:"categories"`
		SubCategories []string    `json:"subCategories"`
		Headline      string      `json:"headline"`
		IsExternal    bool        `json:"isExternal"`
		Preview       string      `json:"preview"`
		HeadlineURL   interface{} `json:"headlineURL"`
		PdfURL        interface{} `json:"pdfURL"`
		FullText      string      `json:"fullText"`
	}
}

func (od *OnDemand) News(displayType string, sources []string, symbols []string) (Headline, error) {
	headline := Headline{}

	_sources := strings.Join(sources, ",")
	_symbols := ""

	if symbols != nil {
		_symbols = strings.Join(symbols, ",")
	}
	_, err := od.Request("getNews.json", fmt.Sprintf("symbols=%v&sources=%s&displayType=%v", _symbols, _sources, displayType), &headline)
	return headline, err
}
