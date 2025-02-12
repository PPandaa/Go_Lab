package tool

import (
	"GoLab/guard"
	"strings"
	"time"
)

func ConvertStringToTime(timeString string) time.Time {
	timeFormat, _ := time.Parse(time.RFC3339, timeString)

	return timeFormat
}

func ConvertStringToTimeByLayout(layout string, target_string string, location_string string) time.Time {
	// Parse the time layout := "2006-01-02 15:04" target_string := "2025-02-10 21:00"
	// fmt.Println(layout, target_string, location_string)
	location, _ := time.LoadLocation(location_string)
	target_time, err := time.ParseInLocation(layout, target_string, location)

	if err != nil {
		guard.Logger.Error("Error parsing time - " + err.Error())
	}
	return target_time
}

func CovertAbbrToNumberString(target_string string) string {
	// Month abbreviation to number mapping
	month_map := map[string]string{
		"JAN": "01",
		"FEB": "02",
		"MAR": "03",
		"APR": "04",
		"MAY": "05",
		"JUN": "06",
		"JUL": "07",
		"AUG": "08",
		"SEP": "09",
		"OCT": "10",
		"NOV": "11",
		"DEC": "12",
	}

	// Convert abbreviation to corresponding number
	return month_map[strings.ToUpper(target_string)]
}
