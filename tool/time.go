package tool

import (
	"GoLab/guard"
	"time"
)

func ConvertStringToTime(timeString string) time.Time {

	timeFormat, _ := time.Parse(time.RFC3339, timeString)

	return timeFormat

}

func ConvertStringToTimeByLayout(layout string, target_string string) time.Time {
	// Parse the time layout := "2006-01-02 15:04" target_string := "2025-02-10 21:00"
	location, _ := time.LoadLocation(time.Local.String())
	target_time, err := time.ParseInLocation(layout, target_string, location)

	if err != nil {
		guard.Logger.Error("Error parsing time - " + err.Error())
	}
	return target_time
}
