package models

//Summary struct contains API response
type Summary struct {
	UserID           string        `json:"user_id"`
	From             struct{}      `json:"from"`
	To               struct{}      `json:"to"`
	Projects         []SummaryItem `json:"projects"`
	Languages        []SummaryItem `json:"languages"`
	Editors          []SummaryItem `json:"editors"`
	OperatingSystems []SummaryItem `json:"operating_systems"`
	Machines         []SummaryItem `json:"machines"`
}

//SummaryItem struct contains item details
type SummaryItem struct {
	Key   string `json:"key"`
	Total int    `json:"total"`
}

//ItemsSorter using go sort
type ItemsSorter []SummaryItem

func (a ItemsSorter) Len() int           { return len(a) }
func (a ItemsSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ItemsSorter) Less(i, j int) bool { return a[i].Total > a[j].Total }
