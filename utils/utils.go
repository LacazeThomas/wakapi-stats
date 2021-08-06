package utils

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

//Check error with details informations
func Check(details string, err error) {
	if err != nil {
		zap.S().Errorf("Error %s %s\n", details, err.Error())
	} else {
		zap.S().Infof("Success %s", details)
	}
}

func SecondToFormattedTime(timeInSecond int) string {
	d := time.Duration(timeInSecond * 1000000000).Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	hFormat := fmt.Sprintf("%dh", h)
	mFormat := ""

	if h == 0 {
		hFormat = ""
	}

	if m < 10 && m > 1 && h > 1 {
		mFormat = fmt.Sprintf("0%dmins", m)
	} else if m == 1 && h == 0 {
		mFormat = fmt.Sprintf("%dmin", m)
	} else if m > 1 && h == 0 {
		mFormat = fmt.Sprintf("%dmins", m)
	} else if (h == 1 && m == 1) || (h > 0 && m < 10 && m > 0) {
		mFormat = fmt.Sprintf("0%d", m)
	} else if m > 1 && h > 0 {
		mFormat = fmt.Sprintf("%d", m)
	}

	return fmt.Sprintf("%s%s", hFormat, mFormat)
}
