package conf

import "time"

const SysTimeform = "2016-01-02 15:04:05"
const SysTimeformShort = "2016-01-02"

var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

var SignSecret = []byte("0123456789abcdef")

var CookieSeceret = "hellolottery"
