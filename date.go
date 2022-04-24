package forum

import (
	"strconv"
	"time"
)

func TimeSince(date1 string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	year1, _ := strconv.Atoi(date1[0:4])
	year2, _ := strconv.Atoi(t[0:4])
	month1, _ := strconv.Atoi(date1[5:7])
	month2, _ := strconv.Atoi(t[5:7])
	day1, _ := strconv.Atoi(date1[8:10])
	day2, _ := strconv.Atoi(t[8:10])
	hour1, _ := strconv.Atoi(date1[11:13])
	hour2, _ := strconv.Atoi(t[11:13])
	min1, _ := strconv.Atoi(date1[14:16])
	min2, _ := strconv.Atoi(t[14:16])
	sec1, _ := strconv.Atoi(date1[17:19])
	sec2, _ := strconv.Atoi(t[17:19])

	year := int(year2) - int(year1)
	month := int(month2) - int(month1)
	day := int(day2) - int(day1)
	hour := int(hour2) - int(hour1)
	min := int(min2) - int(min1)
	sec := int(sec2) - int(sec1)

	if year > 0 {
		return strconv.Itoa(year) + "y"
	} else if month > 0 {
		return strconv.Itoa(month) + "m"
	} else if day > 0 {
		return strconv.Itoa(day) + "d"
	} else if hour > 0 {
		return strconv.Itoa(hour) + "h"
	} else if min > 0 {
		return strconv.Itoa(min) + "min"
	} else if sec > 0 {
		return strconv.Itoa(sec) + "s"
	} else {
		return "now"
	}

}
