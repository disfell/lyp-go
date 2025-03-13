package model

const (
	TimeFormatISO       = "2006-01-02" // ISO 日期（如 2025-03-13）
	TimeFormatSlashDate = "2006/01/02" // 斜杠分隔日期
	TimeFormatDotDate   = "2006.01.02" // 点分隔日期

	TimeFormat24Hour = "15:04:05"    // 24 小时制时间
	TimeFormat12Hour = "03:04:05 PM" // 12 小时制时间

	TimeFormatDateTimeMs = "2006-01-02 15:04:05.000"       // 含毫秒
	TimeFormatDateTimeUs = "2006-01-02 15:04:05.000000"    // 含微秒
	TimeFormatDateTimeNs = "2006-01-02 15:04:05.000000000" // 含纳秒

	TimeFormatDateTimeTZ         = "2006-01-02 15:04:05 MST"   // 带时区缩写
	TimeFormatISOWithTZ          = "2006-01-02T15:04:05-07:00" // ISO 8601 带时区偏移
	TimeFormatISO860_UTC_WITH_MS = "2006-01-02T15:04:05.000Z"
	TimeFormatCompactDateTime    = "20060102_150405" // 紧凑格式（如 20250313_1420）
)
