package easy

import (
	"math"
	"time"
)

func DaysBetweenDates(date1 string, date2 string) int {
	time1, _:=time.ParseInLocation("2006-01-02 15:04:05", date1 + " 15:04:05", time.Local)
	time2, _:=time.ParseInLocation("2006-01-02 15:04:05", date2 + " 15:04:05", time.Local)
	diff := math.Abs(float64((time1.Unix() - time2.Unix())/86400))

	return int(diff)
}