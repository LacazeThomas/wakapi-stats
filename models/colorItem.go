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

	color, _ := jsonparser.GetString(languagesColor, s.Name, "color")
	return ColorSummaryItem{*s, color}
}

func ColorSummaryItems(s []SummaryItem, file string) (colorSummaryItems []ColorSummaryItem, err error) {
	initColor(file)

	for _, element := range s {
		colorSummaryItems = append(colorSummaryItems, element.colorSummaryItem())
	}
	return
}

func initColor(file string) {
	if languagesColor == nil {
		languagesColor, _ = ioutil.ReadFile(file)
	}
}
