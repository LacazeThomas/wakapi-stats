package models

import (
	"time"
)

//Summary struct contains API response
type Summary struct {
	Data struct {
		Username              string        `json:"username"`
		UserID                string        `json:"user_id"`
		Start                 time.Time     `json:"start"`
		End                   time.Time     `json:"end"`
		TotalSeconds          int           `json:"total_seconds"`
		DailyAverage          float64       `json:"daily_average"`
		DaysIncludingHolidays int           `json:"days_including_holidays"`
		Editors               []SummaryItem `json:"editors"`
		Languages             []SummaryItem `json:"languages"`
		Machines              []SummaryItem `json:"machines"`
		Projects              []SummaryItem `json:"projects"`
		OperatingSystems      []SummaryItem `json:"operating_systems"`
	} `json:"data"`
}

//SummaryItem struct contains item details
type SummaryItem struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      int     `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds int     `json:"total_seconds"`
}

//ItemsSorter using go sort
type ItemsSorter []ColorSummaryItem

func (a ItemsSorter) Len() int           { return len(a) }
func (a ItemsSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ItemsSorter) Less(i, j int) bool { return a[i].TotalSeconds > a[j].TotalSeconds }
