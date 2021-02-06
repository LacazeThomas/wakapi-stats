package utils

import (
	"os"
	"sort"

	"github.com/pkg/errors"
	"github.com/wcharczuk/go-chart"

	"github.com/lacazethomas/wakapi-stats/models"
)

//CreateStatsDiagram using SummaryItems and image name
func CreateStatsDiagram(summaryItems []models.SummaryItem, imageName string) (err error) {
	sort.Sort(models.ItemsSorter(summaryItems))
	var chartItems []chart.Value

	maxLen := 7
	if len(summaryItems) < maxLen {
		maxLen = len(summaryItems)
	}

	for i := 0; i < maxLen; i++ {
		val := chart.Value{Value: float64(summaryItems[i].Total), Label: summaryItems[i].Key, Style: chart.Style{
			FontColor: chart.ColorWhite,
		}}
		chartItems = append(chartItems, val)

	}

	return createPieFromChartValues(chartItems, imageName)
}

func createPieFromChartValues(chartItems []chart.Value, imageName string) (err error) {
	chart.DefaultBackgroundColor = chart.ColorTransparent
	chart.DefaultCanvasColor = chart.ColorTransparent

	graph := chart.PieChart{
		Width:  512,
		Height: 512,
		DPI:    float64(75),

		Values: chartItems,
	}

	f, err := os.Create(imageName)
	if err != nil {
		return errors.Wrap(err, "error creating image")
	}
	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		return errors.Wrap(err, "error rendering image from graph")
	}

	return
}
