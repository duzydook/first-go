package helpers

import (
	"strconv"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func RemoveSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

func ParseFloat(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return f
}

func GetAttr(s *goquery.Selection, selector string) string {
	text, ok := s.Attr(selector)
	if ok {
		return text
	} else {
		return ""
	}
}
