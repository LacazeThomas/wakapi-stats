package utils

import (
	"bytes"
	"sort"

	"github.com/pkg/errors"
	"github.com/wcharczuk/go-chart"

	"github.com/lacazethomas/wakapi-stats/models"
)

//CreateStatsDiagram using SummaryItems and image name
func CreateStatsDiagram(summaryItems []models.SummaryItem) ([]byte, error) {
	if len(summaryItems) == 0 {
		return nil, errors.New("This item is empty")
	}

	sort.Sort(models.ItemsSorter(summaryItems))
	var chartItems []chart.Value

	maxLen := 7
	if len(summaryItems) < maxLen {
		maxLen = len(summaryItems)
	}

	for i := 0; i < maxLen; i++ {
		val := chart.Value{Value: float64(summaryItems[i].TotalSeconds), Label: summaryItems[i].Name, Style: chart.Style{
			FontColor: chart.ColorWhite,
		}}
		chartItems = append(chartItems, val)

	}

	return createPieFromChartValues(chartItems)
}

func createPieFromChartValues(chartItems []chart.Value) ([]byte, error) {
	chart.DefaultBackgroundColor = chart.ColorTransparent
	chart.DefaultCanvasColor = chart.ColorTransparent

	graph := chart.PieChart{
		Width:  512,
		Height: 512,
		DPI:    float64(75),

		Values: chartItems,
	}

	b := bytes.NewBuffer([]byte{})

	err := graph.Render(chart.PNG, b)
	if err != nil {
		return nil, errors.Wrap(err, "error rendering image from graph")
	}

	return b.Bytes(), nil
}
