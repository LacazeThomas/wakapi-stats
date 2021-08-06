package utils

import (
	"bytes"
	"sort"

	"github.com/pkg/errors"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"

	"github.com/lacazethomas/wakapi-stats/models"
)

//CreateStatsDiagram using SummaryItems and image name
func CreateStatsDiagram(c []models.ColorSummaryItem, timeSpent bool) ([]byte, error) {
	if len(c) == 0 {
		return nil, errors.New("This item is empty")
	}

	sort.Sort(models.ItemsSorter(c))
	var chartItems []chart.Value

	maxLen := 7
	if len(c) < maxLen {
		maxLen = len(c)
	}

	totalTimeSpent := models.TotalTimeSpent(c)

	for i := 0; i < maxLen; i++ {
		if c[i].Name != "unknown" {

			var style = chart.Style{FontColor: chart.ColorWhite}

			if c[i].Color != "" {
				style = chart.Style{
					FontColor: chart.ColorWhite,
					FillColor: drawing.ColorFromHex(c[i].Color[1:]),
				}
			}

			label := c[i].Name
			if timeSpent {
				if totalTimeSpent/100*2 < c[i].TotalSeconds {
					label = c[i].Name + " " + SecondToFormattedTime(c[i].TotalSeconds)
				}
			}

			val := chart.Value{Value: float64(c[i].TotalSeconds), Label: label, Style: style}
			chartItems = append(chartItems, val)
		}
	}

	return createPieFromChartValues(chartItems)
}

func createPieFromChartValues(chartItems []chart.Value) ([]byte, error) {
	chart.DefaultBackgroundColor = chart.ColorTransparent
	chart.DefaultCanvasColor = chart.ColorTransparent

	graph := chart.PieChart{
		Width:  1024,
		Height: 1024,
		DPI:    float64(75),

		Values: chartItems,
	}

	b := bytes.NewBuffer([]byte{})

	err := graph.Render(chart.SVG, b)
	if err != nil {
		return nil, errors.Wrap(err, "error rendering image from graph")
	}

	return b.Bytes(), nil
}
