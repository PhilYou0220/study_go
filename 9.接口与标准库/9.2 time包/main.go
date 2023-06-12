package main

import (
	"fmt"
	"reflect"
	"time"
)

// timezoneDemo 时区示例
func timezoneDemo() {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
	if err != nil {
		fmt.Println("load America/New_York location failed", err)
		return
	}
	fmt.Println()
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	fmt.Println(timesAreEqual)
}

// timestampDemo 时间戳
func timestampDemo() {
	now := time.Now()        // 获取当前时间
	timestamp := now.Unix()  // 秒级时间戳
	milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
	micro := now.UnixMicro() // 微秒时间戳 Go1.17+
	nano := now.UnixNano()   // 纳秒时间戳
	fmt.Println(timestamp, milli, micro, nano)
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式

	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
func main() {

	//1获取当前时间 时间类型为 自定义类型 type Duration int64 单位是纳秒 1000*1000*1000是秒
	now := time.Now()             //2023-06-09 20:27:59.0141575 +0800 CST m=+0.004994201
	month := time.Now().Month()   // June
	day := time.Now().Day()       //9
	hour := time.Now().Hour()     //20
	minute := time.Now().Minute() //32
	second := time.Now().Second() //32
	fmt.Println(now, month, day, hour, minute, second)

	//2locatime和time zone
	timezoneDemo()

	//3 unixtime
	timestampDemo() //时间戳都是UTC时间

	//4 add 增加时间
	later := now.Add(time.Hour + time.Second*10) // 当前时间加1小时后的时间
	fmt.Println(later)

	//5 Sub求两个时间之间的差值
	//func (t Time) Sub(u Time) Duration
	duration := later.Sub(now)                     //later-now
	fmt.Println(duration)                          //1h0m10s
	fmt.Println(reflect.TypeOf(duration))          //time.Duration
	var interval int64 = 3600 * 1000 * 1000 * 1000 //单位是纳秒 1000*1000*1000是秒
	if duration > time.Duration(interval) {
		fmt.Println("这是真的")
		fmt.Println(duration - time.Duration(interval)) //1h0m9.9999964s
	}

	//6定时器选了chan再说

	//7(时间格式转格式化字符串)
	fmt.Println(now.Format("2006-01-02 15:04:05")) //2023-06-12 23:27:56

	//时间戳转时间格式
	timestampDemo2(1686587075)

	//格式化字符串转时间格式
	//time.ParseInLocation
	//时间格式转时间戳
	//fmt.Println(timeObj.Unix())
}