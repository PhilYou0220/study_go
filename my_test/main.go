//package main
//
//import (
//	"fmt"
//	"time"
//)

//import (
//	"fmt"
//	"reflect"
//	"time"
//)
//
////type Todo struct {
////Title  string
////Status bool
////}
//
//func main() {
//	//var todo Todo
//	//todo =Todo{}
//	//fmt.Println(todo)
//	//auth_param := "bear 23123 132"
//	//part := strings.Split(auth_param, " ")
//	//fmt.Println(part)
//	//for k,v:= range part{
//	//	fmt.Println(k,v)
//	//}
//	fmt.Println(time.Now()) //2023-05-13 12:30:44.5024001 +0800 CST m=+0.004532601
//	fmt.Println(reflect.TypeOf(time.Now()))
//	fmt.Println(time.Now().Add(time.Second*10).Unix()) //1683952254 转时间戳
//}
package main

import (
"fmt"
"time"
"encoding/json"
"code.shomes.cn/pkg/mqcli/rabbit"
)
// 水质设备上下线及
// 检查库 last_cam_online_log
// water_quality
func main() {
	mq := rabbit.New(&rabbit.Options{
		IP: "106.75.154.221",
		Port: 5672,
		User: "epuser",
		Password: "epuser@123",
	})
	createOne(mq)
	// c1(mq)
}

func c1(mq *rabbit.RabbitMQ) {
	m := map[string]string{
		"MN": "CDHT5101410206",
		"w11019-Flag": "N&&",
		"w11019-Avg": "0",
	}
	buf, _ := json.Marshal(&m)
	fmt.Println("send", string(buf))
	mq.Produce("sensor").QeueuName("water_sensor").Publish(buf, "water.warn.event")

}

func createZero(mq *rabbit.RabbitMQ) {
	for i := 0; i < 4; i++ {
		m := map[string]string{
			"MN": "CDHT5101410206",
			"w11019-Flag": "N&&",
			"w11019-Avg": "0",
			"DataTime": time.Now().Add(time.Hour * time.Duration(-1*i)).Format("20060102150400"),
		}
		buf, _ := json.Marshal(&m)
		fmt.Println("send", string(buf))
		mq.Produce("sensor").QeueuName("water_sensor").Publish(buf, "water.warn.event")
		time.Sleep(time.Second * 2)
	}
}
func createOne(mq *rabbit.RabbitMQ) {
	for i := 1; i >0; i-- {
		//  氨氮 高猛酸钾  氨氮>2 高猛酸钾大于15 就是劣v
		m := map[string]string{
			"MN": "CDHT5101080185",
			"w11019-Flag": "", //
			"w11019-Avg": "11.19", //高猛酸钾和cod
			"DataTime": time.Now().Add(time.Hour * time.Duration(-1*i)).Format("20060102150400"),
		}
		buf, _ := json.Marshal(&m)
		fmt.Println("send", string(buf))
		mq.Produce("sensor").QeueuName("water_sensor").Publish(buf, "water.warn.event")

		//4类水质 氨氮
		m = map[string]string{
			"MN": "CDHT5101080185",
			"w21003-Flag": "0.1", //未知
			"w21003-Avg": "1." + fmt.Sprint(i), //氨氮
			//"w21003-Avg": "0.1",//氨氮
			"DataTime": time.Now().Add(time.Hour * time.Duration(-1*i)).Format("20060102150400"),
		}
		buf, _ = json.Marshal(&m)
		fmt.Println("send", string(buf))
		mq.Produce("sensor").QeueuName("water_sensor").Publish(buf, "water.warn.event")

		//4类水质 高猛酸钾
		//m := map[string]string{
		//	"MN": "CDHT5101080185",
		//	"w11019-Flag": "2", //未知
		//	"w11019-Avg": "9." + fmt.Sprint(i), //高猛酸钾
		//	"DataTime": time.Now().Add(time.Hour * time.Duration(-1*i)).Format("20060102150400"),
		//}
		//buf, _ := json.Marshal(&m)
		//fmt.Println("send", string(buf))
		//mq.Produce("sensor").QeueuName("water_sensor").Publish(buf, "water.warn.event")

		//6类水质
		//m := map[string]string{
		//	"MN": "CDHT5101080185",
		//	"w21003-Flag": "0.3",
		//	"w21003-Avg": "9." + fmt.Sprint(i), //氨氮
		//	"DataTime": time.Now().Add(time.Hour * time.Duration(-1*i)).Format("20060102150400"),
		//}
		//buf, _ := json.Marshal(&m)
		//fmt.Println("send", string(buf))
		//mq.Produce("sensor").QeueuName("water_sensor").Publish(buf, "water.warn.event")

		time.Sleep(time.Second * 2)

	}
}