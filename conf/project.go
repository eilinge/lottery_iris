package conf

import "time"

const RunningServiceGroup = false // 是否运行在集群模式
const UserPrizeMax = 3000            // 用户每天最多抽奖次数
const IpPrizeMax = 30000             // 同一个IP每天最多抽奖次数
const IpLimitMax = 300000            // 同一个IP每天最多抽奖次数

const GtypeVirtual = 0   // 虚拟币
const GtypeCodeSame = 1  // 虚拟券，相同的码
const GtypeCodeDiff = 2  // 虚拟券，不同的码
const GtypeGiftSmall = 3 // 实物小奖
const GtypeGiftLarge = 4 // 实物大奖

const SysTimeform = "2006-01-02 15:04:05"
const SysTimeformShort = "2006-01-02"

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

// ObjSalesign 签名密钥
var SignSecret = []byte("0123456789abcdef")

// cookie中的加密验证密钥
var CookieSecret = "hellolottery"
