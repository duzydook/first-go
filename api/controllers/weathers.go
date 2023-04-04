package controller

import (
	"demo/api/helpers"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo/v4"
)

type Weather struct {
	Capital     string `json:"capital"`
	Time        string `json:"time"`
	Condition   string `json:"condition"`
	Temperature string `json:"temperature"`
}

func loadWeatherDataFromSource() []Weather {
	var weathers []Weather

	resp, err := http.Get("https://www.timeanddate.com/weather/?low=c")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	doc.Find("table.zebra tbody tr").Each(func(i int, s *goquery.Selection) {
		td1 := s.Find("td:nth-child(1)")

		var weather Weather
		var td5 *goquery.Selection

		if td1.Find("a").Length() != 0 {
			weather, td5 = getWeather(td1)
			weathers = append(weathers, weather)
		}

		if td5.Find("a").Length() != 0 {
			weather, _ = getWeather(td5)
			weathers = append(weathers, weather)
		}
	})

	return weathers
}

// endpoint for retriving data of weather from timeanddate.com
func RetrieveWeathers(c echo.Context) error {
	weathers := loadWeatherDataFromSource()

	return c.JSON(http.StatusOK, weathers)
}

// endpoint of filtering weathers by capital
func SearchWeathers(c echo.Context) error {
	weathers := loadWeatherDataFromSource()

	capital := c.QueryParam("capital")
	var filteredWeathers = []Weather{}

	for _, weather := range weathers {
		if strings.Contains(strings.ToLower(weather.Capital), strings.ToLower(capital)) {
			filteredWeathers = append(filteredWeathers, weather)
		}
	}

	return c.JSON(http.StatusOK, filteredWeathers)
}

// private method to collect Weather data from html
func getWeather(td1 *goquery.Selection) (Weather, *goquery.Selection) {

	td2 := td1.Next()
	td3 := td2.Next()
	td4 := td3.Next()

	capital := td1.Find("a").Text()
	time := td2.Text()
	condition := helpers.GetAttr(td3.Find("img"), "alt")
	temperature := td4.Text()

	return Weather{
		Capital:     capital,
		Time:        time,
		Condition:   condition,
		Temperature: temperature,
	}, td4.Next()
}
