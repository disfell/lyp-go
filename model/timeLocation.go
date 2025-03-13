package model

// ======================== 亚洲 ========================
const (
	TIME_LOCATION_SHANGHAI  = "Asia/Shanghai"  // 中国上海
	TIME_LOCATION_TOKYO     = "Asia/Tokyo"     // 日本东京
	TIME_LOCATION_SINGAPORE = "Asia/Singapore" // 新加坡
	TIME_LOCATION_DUBAI     = "Asia/Dubai"     // 阿联酋迪拜（UTC+4）
	TIME_LOCATION_HONG_KONG = "Asia/Hong_Kong" // 中国香港
	TIME_LOCATION_SEOUL     = "Asia/Seoul"     // 韩国首尔
	TIME_LOCATION_BANGKOK   = "Asia/Bangkok"   // 泰国曼谷（UTC+7）
	TIME_LOCATION_KOLKATA   = "Asia/Kolkata"   // 印度加尔各答（UTC+5:30）
)

// ======================== 欧洲 ========================
const (
	TIME_LOCATION_LONDON = "Europe/London" // 英国伦敦（UTC+0/UTC+1 夏令时）
	TIME_LOCATION_PARIS  = "Europe/Paris"  // 法国巴黎
	TIME_LOCATION_BERLIN = "Europe/Berlin" // 德国柏林
	TIME_LOCATION_MOSCOW = "Europe/Moscow" // 俄罗斯莫斯科（UTC+3）
	TIME_LOCATION_ATHENS = "Europe/Athens" // 希腊雅典
)

// ======================== 北美洲 ========================
const (
	TIME_LOCATION_NEW_YORK    = "America/New_York"    // 美国纽约（东部时间）
	TIME_LOCATION_LOS_ANGELES = "America/Los_Angeles" // 美国洛杉矶（太平洋时间）
	TIME_LOCATION_CHICAGO     = "America/Chicago"     // 美国芝加哥（中部时间）
	TIME_LOCATION_TORONTO     = "America/Toronto"     // 加拿大多伦多
	TIME_LOCATION_MEXICO_CITY = "America/Mexico_City" // 墨西哥城
)

// ======================== 南美洲 ========================
const (
	TIME_LOCATION_SAO_PAULO    = "America/Sao_Paulo"              // 巴西圣保罗
	TIME_LOCATION_BUENOS_AIRES = "America/Argentina/Buenos_Aires" // 阿根廷布宜诺斯艾利斯
)

// ======================== 非洲 ========================
const (
	TIME_LOCATION_JOHANNESBURG = "Africa/Johannesburg" // 南非约翰内斯堡
	TIME_LOCATION_CAIRO        = "Africa/Cairo"        // 埃及开罗
	TIME_LOCATION_LAGOS        = "Africa/Lagos"        // 尼日利亚拉各斯（UTC+1）
)

// ======================== 大洋洲 ========================
const (
	TIME_LOCATION_SYDNEY   = "Australia/Sydney" // 澳大利亚悉尼（夏令时）
	TIME_LOCATION_AUCKLAND = "Pacific/Auckland" // 新西兰奥克兰
	TIME_LOCATION_HONOLULU = "Pacific/Honolulu" // 美国夏威夷檀香山（UTC-10）
)

// ======================== 其他标准时区 ========================
const (
	TIME_LOCATION_UTC         = "UTC"       // 协调世界时
	TIME_LOCATION_GMT         = "Etc/GMT"   // 格林尼治标准时间（同 UTC）
	TIME_LOCATION_UTC_PLUS_1  = "Etc/GMT-1" // UTC+1（注意 IANA 时区符号规则）
	TIME_LOCATION_UTC_MINUS_5 = "Etc/GMT+5" // UTC-5（符号与常规相反）
)
