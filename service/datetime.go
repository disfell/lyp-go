package service

import (
	"lyp-go/model"
	"lyp-go/output"
	"strconv"
	"time"
)

type DateTimeServ struct {
}

func (dts *DateTimeServ) CalculateDateDiff(begin time.Time, end time.Time) interface{} {
	swap := false
	if begin.After(end) {
		begin, end = end, begin
		swap = true
	}

	year := end.Year() - begin.Year()
	month := int(end.Month()) - int(begin.Month())
	day := end.Day() - begin.Day()

	// 调整天数
	if day < 0 {
		lastDayOfPreviousMonth := time.Date(begin.Year(), begin.Month(), 0, 0, 0, 0, 0, begin.Location()).Day()
		if begin.Day() > lastDayOfPreviousMonth {
			begin = time.Date(begin.Year(), begin.Month(), lastDayOfPreviousMonth, 0, 0, 0, 0, begin.Location())
		}
		day += lastDayOfPreviousMonth
		month--
	}

	// 调整月份
	if month < 0 {
		month += 12
		year--
	}

	if swap {
		year, month, day = -year, -month, -day
	}
	return map[string]int{
		"year":  year,
		"month": month,
		"day":   day,
	}
}

func (dts *DateTimeServ) ParseDate(date string) time.Time {
	if len(date) != 8 {
		panic(output.Err(model.ErrorCode, "入参格式必须为 yyyyMMdd", ""))
	}

	year, _ := strconv.Atoi(date[0:4])
	month, _ := strconv.Atoi(date[4:6])
	day, _ := strconv.Atoi(date[6:8])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
