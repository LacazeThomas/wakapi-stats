package utils

import "go.uber.org/zap"

//Check error with details informations
func Check(details string, err error) {
	if err != nil {
		zap.S().Errorf("Error %s %s\n", details, err.Error())
	} else {
		zap.S().Infof("Success %s", details)
	}
}
