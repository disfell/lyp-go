package tests

import (
	"lyp-go/model"
	"testing"
	"time"
)

func TestStr2time(t *testing.T) {
	// 示例1：解析无时区字符串为 CST
	timeStr := "2025-03-12 06:09:00"
	loc, _ := time.LoadLocation(model.TIME_LOCATION_SHANGHAI)
	t1, _ := time.ParseInLocation(time.DateTime, timeStr, loc)
	t.Logf("直接解析为 CST: %v", t1)

	// 示例2：转换 UTC 时间字符串为 CST
	utcStr := "2025-03-12T06:09:00.000Z"
	t2, _ := time.Parse(model.TimeFormatISO860UTCWithMs, utcStr)
	t.Logf("UTC 时间:%v", t2)
	t.Logf("转换后 CST:%v", t2.In(loc))
}

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	t.Logf("now.Local() = %v", now.Local())
	t.Logf("now.Format() = %v", now.Format(model.TimeFormatDateTimeMs))
}
