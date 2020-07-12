package easy

import (
	"fmt"
	"time"
)

func DayOfYear(date string) int {
	date += " 00:00:00"
	start, _ := time.Parse("2006-01-02 15:04:05", date)
	startUnix := time.Date(start.Year(), 1, 0, 0, 0, 0, 0, start.Location())
	fmt.Println(start.Year())
	return int(start.Unix() - startUnix.Unix()) / 86400
}
