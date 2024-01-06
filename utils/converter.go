package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ConvertStringDurationIntoInteger(durationString string) (int, error) {
	arrayDuration := strings.Split(durationString, " ")

	duration, err := strconv.Atoi(arrayDuration[0])

	if err != nil {
		return 0, err
	}

	unit := arrayDuration[1]

	if unit != "sec" && unit != "min" {
		return 0, errors.New("only sec and min unit accepted")
	}

	if unit == "min" {
		duration *= 60
	}

	return duration, nil
}
