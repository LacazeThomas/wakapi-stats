package models

import (
	"io/ioutil"

	"github.com/buger/jsonparser"
)

var languagesColor []byte

//ColorSummaryItem is SummaryItem with color in hex
type ColorSummaryItem struct {
	SummaryItem
	Color string
}

type ColorLanguage struct {
	Color interface{} `json:"color"`
	URL   string      `json:"url"`
}

func (s *SummaryItem) colorSummaryItem() ColorSummaryItem {
	initColor()

	color, _ := jsonparser.GetString(languagesColor, s.Name, "color")
	return ColorSummaryItem{*s, color}
}

func ColorSummaryItems(s []SummaryItem) (c []ColorSummaryItem) {
	for _, element := range s {
		c = append(c, element.colorSummaryItem())
	}
	return
}

func initColor() {
	if languagesColor == nil {
		languagesColor, _ = ioutil.ReadFile("colors.json")
	}
}
