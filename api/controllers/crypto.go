package controller

import (
	"demo/api/helpers"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo/v4"
)

type Crypto struct {
	Rank      string  `json:"rank"`
	Name      string  `json:"name"`
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	MarketCap float64 `json:"marketcap"`
	Volume    float64 `json:"volume"`
	H1        float64 `json:"h1"`
	H24       float64 `json:"h24"`
	D7        float64 `json:"d7"`
}

// endpoint for retriving data of crypto from coingecko.com
func RetrieveCryptos(c echo.Context) error {

	var cryptos []Crypto = []Crypto{}

	resp, err := http.Get("https://www.coingecko.com/")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	// doc := soup.HTMLParse(resp)
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		crypto := getCrypto(s)
		cryptos = append(cryptos, crypto)
	})

	return c.JSONPretty(http.StatusOK, cryptos, "	")
}

// private method to collect crypto data from html
func getCrypto(row *goquery.Selection) Crypto {

	td1 := row.Find("td:nth-child(1)")
	td2 := td1.Next()
	td3 := td2.Next()
	td4 := td3.Next()
	td5 := td4.Next()
	td6 := td5.Next()
	td7 := td6.Next()
	td8 := td7.Next()
	td9 := td8.Next()

	// TODO: opportunity to add error exception
	rank := td2.Text()
	name := td3.Find("div a span:nth-child(1)").Text()
	symbol := td3.Find("div a span:nth-child(2)").Text()
	price := helpers.GetAttr(td4, "data-sort")
	marketcap := helpers.GetAttr(td9, "data-sort")
	volume := helpers.GetAttr(td8, "data-sort")
	h1 := helpers.GetAttr(td5, "data-sort")
	h24 := helpers.GetAttr(td6, "data-sort")
	d7 := helpers.GetAttr(td7, "data-sort")

	return Crypto{
		Rank:      helpers.RemoveSpace(rank),
		Name:      helpers.RemoveSpace(name),
		Symbol:    helpers.RemoveSpace(symbol),
		Price:     helpers.ParseFloat(price),
		MarketCap: helpers.ParseFloat(marketcap),
		Volume:    helpers.ParseFloat(volume),
		H1:        helpers.ParseFloat(h1),
		H24:       helpers.ParseFloat(h24),
		D7:        helpers.ParseFloat(d7),
	}
}
