package gtime

import (
	"fmt"
	"time"
)

const (
	TIMEFORMAT1      = "2006-01-02 15:04:05"
	TIMEFORMAT2      = "2006-01-02 15:04:05.000000"
	TIMEFORMAT3      = "2006/01/02 15:04:05"
	TIMEFORMAT4      = "2006/01/02 15:04:05.000000"
	TIMEFORMATDate   = "2006-01-02"
	TIMEFORMATDate2  = "2006/01/02"
	TIMEFORMATTIME   = "15:04:05"
	TIMEFORMATTIME2  = "15:04:05.000000"
	TIMEFORMATCN     = "2006年01月02日 15:04:05"
	TIMEFORMATCNDATE = "2006年01月02日"
	TIMEFORMATCNTIME = "15时04分05秒"
)

type GTime struct {
	_time time.Time
}

func Now() GTime {
	return GTime{_time: time.Now()}
}

// 获取当前时间的秒数
func (g GTime) Unix() int64 {
	return g._time.Unix()
}

// 获取当前时间的纳秒
func (g GTime) UnixNano() int64 {
	return g._time.UnixNano()
}

// 获取当前时间的日期部分 yyyy-MM-HH
func (g GTime) Date() string {
	return g._time.Format(TIMEFORMATDate)
}

// 获取当前时间的时间部分 hh:mm:ss
func (g GTime) Time() string {
	return g._time.Format(TIMEFORMATTIME)
}

// 获取当前时间的年份
func (g GTime) Year() int {
	return g._time.Year()
}

// 返回当前时间的月份
func (g GTime) Month() int {
	return (int)(g._time.Month())
}

// 返回当前时间的天数
func (g GTime) Day() int {
	return g._time.Day()
}

// 返回当前时间的小时数
func (g GTime) Hour() int {
	return g._time.Hour()
}

// 返回当前时间的分钟数
func (g GTime) Minute() int {
	return g._time.Minute()
}

// 返回当前时间的秒数
func (g GTime) Second() int {
	return g._time.Second()
}

// 返回1970-01-01到当前时间的微秒
func (g GTime) Micro() int64 {
	return g._time.UnixMicro()
}

// 只返回当前时间的微秒
func (g GTime) MicorOnly() int64 {
	return g._time.UnixMicro() - g._time.Unix()*1000000
}

// 返回1970-01-01到当前时间的纳秒
func (g GTime) Nano() int64 {
	return g._time.UnixNano()
}

// 只返回当前时间的纳秒
func (g GTime) NanoOnly() int64 {
	return g._time.UnixNano() - g._time.Unix()*1000000000
}

// 返回当前时间的星期整数
func (g GTime) Week() int {
	return (int)(g._time.Weekday())
}

// 返回当前年份到现在的天数
func (g GTime) Days() int {
	return g._time.YearDay()
}

// 返回星期的英文字符串
func (g GTime) WeekEn() string {
	val := g._time.Weekday()
	var val_str string
	switch val {
	case 0:
		val_str = "Sunday"
	case 1:
		val_str = "Monday"
	case 2:
		val_str = "Tuesday"
	case 3:
		val_str = "Wednesday"
	case 4:
		val_str = "Thursday"
	case 5:
		val_str = "Friday"
	default:
		val_str = "Saturday"
	}
	return val_str
}

// 返回星期的英文简写字符串
func (g GTime) WeekEns() string {
	val := g._time.Weekday()
	var val_str string
	switch val {
	case 0:
		val_str = "Sun"
	case 1:
		val_str = "Mon"
	case 2:
		val_str = "Tue"
	case 3:
		val_str = "Wed"
	case 4:
		val_str = "Thu"
	case 5:
		val_str = "Fri"
	default:
		val_str = "Sat"
	}
	return val_str
}

// 返回星期的中文字符串
func (g GTime) WeekCn() string {
	val := g._time.Weekday()
	var val_str string
	switch val {
	case 0:
		val_str = "星期日"
	case 1:
		val_str = "星期一"
	case 2:
		val_str = "星期二"
	case 3:
		val_str = "星期三"
	case 4:
		val_str = "星期四"
	case 5:
		val_str = "星期五"
	default:
		val_str = "星期六"
	}
	return val_str
}

// 返回星期的中文字符串
func (g GTime) WeekCns() string {
	val := g._time.Weekday()
	var val_str string
	switch val {
	case 0:
		val_str = "日"
	case 1:
		val_str = "一"
	case 2:
		val_str = "二"
	case 3:
		val_str = "三"
	case 4:
		val_str = "四"
	case 5:
		val_str = "五"
	default:
		val_str = "六"
	}
	return val_str
}

// 返回ISO标准的周期整数
func (g GTime) WeekISO() int {
	_, week := g._time.ISOWeek()
	return week
}

func (g GTime) ToString() string {
	return g._time.Format("2006-01-02 15:04:05")
}

// 返回当前时区
func (g GTime) UTC() string {

	return g._time.UTC().Format(TIMEFORMAT1)
}

// 返回当前时区名称(规范名称)
func (g GTime) ZoneName() string {
	name, _ := g._time.Zone()
	return name
}

// 返回时区差的秒数
func (g GTime) Zone() int {
	_, sec := g._time.Zone()
	return sec
}

// 返回当前时区名称
func (g GTime) LocalName() string {
	return g._time.Location().String()
}

// 返回当前标准时间
func (g GTime) Local() time.Time {
	return g._time.Local()
}

// 返回指定格式
func (g GTime) ToStr(format string) string {
	return g._time.Format(format)
}

// 返回中文日期时间
func (g GTime) ToCN() string {
	return g._time.Format(TIMEFORMATCN)
}

// 返回中文日期
func (g GTime) ToDateCN() string {
	return g._time.Format(TIMEFORMATCNDATE)
}

// 返回中文时间
func (g GTime) ToTimeCN() string {
	return g._time.Format(TIMEFORMATCNTIME)
}

/*
时间计算
*/
// 增加时间多少小时(如果参数为负数就相减),并返回新的时间对象
func (g GTime) AddHour(hours time.Duration) GTime {
	d, _ := time.ParseDuration("1h")
	return GTime{
		_time: g._time.Add(hours * d),
	}
}

// 增加时间多少分钟(如果参数为负数就相减),并返回新的时间对象
func (g GTime) AddMinute(minute time.Duration) GTime {
	m, _ := time.ParseDuration("1m")
	return GTime{
		_time: g._time.Add(minute * m),
	}
}

// 增加时间多少秒数(如果参数为负数就相减),并返回新的时间对象
func (g GTime) Add(secound time.Duration) GTime {
	s, _ := time.ParseDuration("1s")
	return GTime{
		_time: g._time.Add(secound * s),
	}
}

// 两个时间的相差的秒数
func (g GTime) Sub(other GTime) int64 {
	return other._time.Unix() - g._time.Unix()
}

// 两个时间相差的微秒数
func (g GTime) SubMicro(other GTime) int64 {
	return other._time.UnixMicro() - g._time.UnixMicro()
}

// 两个时间相差的纳秒数
func (g GTime) SubNano(other GTime) int64 {
	return other._time.UnixNano() - g._time.UnixNano()
}

// 根据字符串返回时间类型
func FromString(datetime string, format string) GTime {
	dt, err := time.ParseInLocation(format, datetime, time.Local)
	if err != nil {
		fmt.Printf("字符串格式有误,无法转换")

	}
	return GTime{
		_time: dt,
	}
}

// 比较2个时间是否相同
func (g GTime) Equal(other GTime) bool {
	return g._time.Equal(other._time)
}

// 比较当前时间是否在指定的时间之后
func (g GTime) After(other GTime) bool {
	return g._time.After(other._time)
}

// 比较当前时间是否在指定的时间之前
func (g GTime) Before(other GTime) bool {
	return g._time.Before(other._time)
}

// 定时器
func Ticker(second time.Duration) *time.Ticker {
	return (*time.Ticker)(time.NewTimer(time.Second * second))
}

// 根据秒数返回时间类型
func FromUnix(unix int64) GTime {
	dt := time.Unix(unix, 0)
	return GTime{
		_time: dt,
	}

}

// 自定义日期时间
func FromDateTime(year int, month int, day int, hour int, minute int, second int, nano int, loc *time.Location) GTime {
	dt := time.Date(year, (time.Month)(month), day, hour, minute, second, nano, loc)
	return GTime{
		_time: dt,
	}
}

// 返回当前时区的符号
func LoadZone(zonename string) *time.Location {
	local, err := time.LoadLocation(zonename)
	if err != nil {
		fmt.Print("time location error")
		return time.Local
	}

	return local
}
