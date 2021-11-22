package tool

import "time"

func ConvertStringToTime(timeString string) time.Time {

	timeFormat, _ := time.Parse(time.RFC3339, timeString)

	return timeFormat

}
