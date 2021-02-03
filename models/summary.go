package models

//Summary struct contains API response
type Summary struct {
	UserID           string         `json:"user_id"`
	From             struct{}       `json:"from"`
	To               struct{}       `json:"to"`
	Projects         []SummaryItems `json:"projects"`
	Languages        []SummaryItems `json:"languages"`
	Editors          []SummaryItems `json:"editors"`
	OperatingSystems []SummaryItems `json:"operating_systems"`
	Machines         []SummaryItems `json:"machines"`
}

//SummaryItems struct contains item details
type SummaryItems struct {
	Key   string `json:"key"`
	Total int    `json:"total"`
}

//ItemsSorter using go sort
type ItemsSorter []SummaryItems

func (a ItemsSorter) Len() int           { return len(a) }
func (a ItemsSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ItemsSorter) Less(i, j int) bool { return a[i].Total > a[j].Total }
