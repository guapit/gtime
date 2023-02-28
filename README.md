# Gtime-Golang模块



<p align="center">
<a href="#"><img src="https://img.shields.io/badge/Module-gtime-critical.svg"/></a>
<a href="#"><img src="https://img.shields.io/badge/Language-Golang-blue"/></a>
    <a href="#"><img src="https://img.shields.io/badge/Version-0.1.0-f1c232"/></a>
<img src="https://img.shields.io/badge/Author-guapit-ff69b4"/>
<a href="https://www.github.com/guapit"><img src="https://img.shields.io/badge/Github-guapit-success"/></a>
<a href="https://www.gitee.com/guapit"><img src="https://img.shields.io/badge/Gitee-guapit-yellowgreen"/></a>
<a href="#"><img src="https://img.shields.io/badge/E--mail-guapit%40qq.com-yellowgreen"/></a>
</p><br>

根据Python版的gtime时间,按照设计模式的设计规范,方便前后端对接和一致性,开发了这个时间模块,如果在使用中遇到问题,请随时联系我

## 接口文档

比如获取当前时间:

`dt := gtime.Now()`

| 接口                           | 说明                                                     |
| ------------------------------ | -------------------------------------------------------- |
| dt.ToString()                  | 转换成`2023-02-28 18:47:43`格式字符串                    |
| dt.Date()                      | 转换成`2023-02-28格式字符串                              |
| dt.Time()                      | 转换成`18:47:43`格式字符串                               |
| dt.Unix()                      | 得到从`1970-01-01`到现在的秒数                           |
| dt.Micro()                     | 得到从`1970-01-01`到现在的微秒数                         |
| dt.MicorOnly()                 | 只得到微秒时间戳: `.666666`                              |
| dt.Nano()                      | 得到从`1970-01-01`到现在的纳秒数                         |
| dt.NanoOnly()                  | 只得到微秒时间戳: `.666666000`                           |
| dt.Year()                      | 获取年份                                                 |
| dt.Month()                     | 获取月份                                                 |
| dt.Day()                       | 获取天                                                   |
| dt.Hour()                      | 获取小时                                                 |
| dt.Minute()                    | 获取分钟                                                 |
| dt.Second()                    | 获取秒数                                                 |
| dt.Days()                      | 获取从当年01月01日到现在的天数                           |
| dt.Week()                      | 获取当前时间的星期数, 星期日: 0 -> 星期六: 6             |
| dt.WeekEn()                    | 获取星期的英文单词,星期日: `Sunday`                      |
| dt.WeekEns()                   | 获取星期的英文单词简写,星期日: `Sun`                     |
| dt.WeekCn()                    | 获取星期的中文单词,`星期日`                              |
| dt.WeekCns()                   | 获取星期的中文数,`日`                                    |
| dt.WeekISO()                   | 获取ISO星期整数                                          |
| dt.LocalName()                 | 获取当前时区名称                                         |
| dt.ZoneName()                  | 获取当前时区名称(规范名称)                               |
| dt.Zone()                      | 获取时差(秒数)                                           |
| dt.ToCN                        | 转换成中文日期时间: `2023年02月28日 16:14:08`            |
| dt.ToDateCN()                  | 转换成中文只有日期: `2023年02月28日`                     |
| dt.ToTimeCN()                  | 转换成中文只有时间: `16时14分08秒`                       |
| 计算时间                       |                                                          |
| dt.Add(2400)                   | 默认为秒数, 如果用加法增加`2400秒`, 如果用减法 `-2400秒` |
| dt.Sub(other GTime)            | 计算两个时间对象的差值,返回秒数                          |
| dt.SubMicro(dt2)               | 计算两个时间对象的差值,返回微秒数                        |
| dt.SubNano(dt2)                | 计算两个时间对象的差值,返回纳秒数                        |
| 比较运算                       |                                                          |
| dt.Equal(other GTime)          | 比较2个时间对象是否相等                                  |
| dt.After(other GTime)          | 比较当前时间是否在指定的时间之后                         |
| dt.Before(other GTime)         | 比较当前时间是否在指定的时间之前                         |
| 定时器                         |                                                          |
| tick := Ticker(间隔秒数)       | 设置一个定时器,关闭方法: `tick.Stop()`                   |
|                                |                                                          |
| 格式化转换成时间               |                                                          |
| FromString(字符串, 格式化方式) | 根据字符串返回时间类型                                   |
| FromUnix(秒数)                 | 根据秒数返回时间类型                                     |
|                                |                                                          |



## 使用方法

```go
func main() {
	dt := gtime.Now()
	fmt.Printf("当前完整时间: %v\n", dt.ToString())
	fmt.Printf("当前日期: %v\n", dt.Date())
	fmt.Printf("当前时间: %v\n", dt.Time())
	fmt.Printf("当前总秒数: %v\n", dt.Unix())
	fmt.Printf("当前年份: %v\n", dt.Year())
	fmt.Printf("当前月份: %v\n", dt.Month())
	fmt.Printf("当前天: %v\n", dt.Day())
	fmt.Printf("当前小时: %v\n", dt.Hour())
	fmt.Printf("当前分钟: %v\n", dt.Minute())
	fmt.Printf("当前秒数: %v\n", dt.Second())
	fmt.Printf("当前微秒: %v\n", dt.Micro())
	fmt.Printf("当前微秒(只返回微秒): %v\n", dt.MicorOnly())
	fmt.Printf("当前纳秒: %v\n", dt.Nano())
	fmt.Printf("开年到现在的天数: %v\n", dt.Days())
	fmt.Printf("当前微秒(只返回纳秒): %v\n", dt.NanoOnly())
	fmt.Printf("当前周期: %v\n", dt.Week())
	fmt.Printf("当前周期(整型): %v\n", int(dt.Week()))
	fmt.Printf("当前周期(英文全拼): %v\n", dt.WeekEn())
	fmt.Printf("当前周期(英文简写): %v\n", dt.WeekEns())
	fmt.Printf("当前周期(中文全拼): %v\n", dt.WeekCn())
	fmt.Printf("当前周期(中文简写): %v\n", dt.WeekCns())
	fmt.Printf("当前周期(ISO): %v\n", dt.WeekISO())
	fmt.Printf("当前时区名称: %v\n", dt.LocalName())
	fmt.Printf("当前时区名称(规范名称): %v\n", dt.ZoneName())
	fmt.Printf("当前时区偏差(秒数): %v\n", dt.Zone())
	dt = gtime.FromString("2023-02-28 16:10:30", gtime.TIMEFORMAT1)
	fmt.Printf("字符串格式化后的时间: %v\n", dt.ToString())
	dt = gtime.FromUnix(1677572048)
	fmt.Printf("秒数格式化后的时间: %v\n", dt.ToString())
	fmt.Printf("中文日期时间: %v\n", dt.ToCN())
	fmt.Printf("中文日期: %v\n", dt.ToDateCN())
	fmt.Printf("中文时间: %v\n", dt.ToTimeCN())
	dt = gtime.Now()
	dt2 := dt.Add(2400)
	fmt.Printf("原始时间: %v\n", dt.ToString())
	fmt.Printf("当前时间增加2400秒后: %v\n", dt2.ToString())
	fmt.Printf("差值(秒): %v\n", dt.Sub(dt2))
	fmt.Printf("差值(微秒): %v\n", dt.SubMicro(dt2))
	fmt.Printf("差值(纳秒): %v\n", dt.SubNano(dt2))

}
```

