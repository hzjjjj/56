package main

//根据王吉舟讲易经原理编写的网站，有htmx功能，金辉
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var pl = fmt.Println
var diagrams8 = []string{"乾", "兑", "离", "震", "巽", "坎", "艮", "坤"}

/*
	var diagrams64 = []string{"乾", "坤", "屯", "蒙", "需", "讼", "师", "比", "小畜", "履", "泰",

"否", "同人", "大有", "谦", "豫", "随", "蛊", "临", "观", "噬嗑", "贲", "剥", "复", "无妄", "大畜",
"颐", "大过", "坎", "离", "咸", "恒", "遁", "大壮", "晋", "明夷", "家人", "睽", "蹇", "解", "损", "益",
"夬", "姤", "萃", "升", "困", "井", "革", "鼎", "震", "艮", "渐", "归妹", "丰", "旅", "巽", "兑",
"涣", "节", "中孚", "小过", "既济", "未济"}
*/
var links = []string{"103", "113", "140", "153", "184", "109", "172", "127", "183", "200",
	"189", "144", "167", "187", "170", "185", "141", "177", "169",
	"148", "190", "263", "197", "174", "173", "195", "196", "192",
	"171", "180", "256", "143", "112", "255", "176", "182", "198",
	"212", "194", "147", "108", "244", "257", "106", "188", "168",
	"179", "111", "159", "181", "149", "164", "145", "107", "193",
	"150", "126", "146", "175", "152", "186", "110", "142", "105"}

var diagrams10 = []string{"乾", "履", "同人", "无妄", "姤", "讼", "遁", "否", "夬", "兑",
	"革", "随", "大过", "困", "咸", "萃", "大有", "睽", "离",
	"噬嗑", "鼎", "未济", "旅", "晋", "大壮", "归妹", "丰", "震",
	"恒", "解", "小过", "豫", "小畜", "中孚", "家人", "益", "巽",
	"涣", "渐", "观", "需", "节", "既济", "屯", "井", "坎",
	"蹇", "比", "大畜", "损", "贲", "颐", "蛊", "蒙", "艮",
	"剥", "泰", "临", "明夷", "复", "升", "师", "谦", "坤"}
var fengshui = []string{"天", "泽", "火", "雷", "风", "水", "山", "地"}
var binary = []string{"000", "001", "010", "011", "100", "101", "110", "111"}

func numToDiagrams(num int) string {
	return (diagrams8[num])
}

func numToDiagrams64(num int64) string {
	return diagrams10[num]
}

func numToLink(num int64) string {
	return links[num]
}

func numToFengshui(num int) string {
	return (fengshui[num])
}
func numToBinary(num int) string {
	return (binary[num])
}

func binaryToNum(binary string) int64 {
	numInput, _ := strconv.ParseInt(binary, 2, 0)
	return numInput
}

func main() {
	//	网站首页 index.html  请求方法GET POST
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
		//pl(time.Now().Format("01-02 15:04:05.000"))
		pl("welcome!")
	}

	//输入0>乾卦
	h2 := func(w http.ResponseWriter, r *http.Request) {
		numInput := r.PostFormValue("num")
		num, _ := strconv.Atoi(numInput)
		words1 := numToDiagrams(num)
		words2 := numToFengshui(num)
		words3 := numToBinary(num)
		htmlStr := fmt.Sprintf("<p> %d => %s 为%s Binary is %s</p>", num, words1, words2, words3)
		//使用fmt.Sprintf函数将信息格式化为字符串
		tmpl, _ := template.New("t").Parse(htmlStr)
		pl("收到POST:", num, "-", words1)
		tmpl.Execute(w, nil)
	}

	//高岛易占：输入3个数字，出来一个卦象和动爻变卦
	h3 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 200) //实现前端页面加载效果
		numInput1 := r.PostFormValue("num1")
		numInput2 := r.PostFormValue("num2")
		numInput3 := r.PostFormValue("num3")
		num1, _ := strconv.Atoi(numInput1)
		num2, _ := strconv.Atoi(numInput2)
		num3, _ := strconv.Atoi(numInput3)
		words1 := numToFengshui(num1 - 1)
		words2 := numToFengshui(num2 - 1)
		words3 := numToBinary(num1-1) + numToBinary(num2-1)
		numOutput := binaryToNum(words3)
		anser := numToDiagrams64(numOutput)
		words4 := numToLink(numOutput)
		now := time.Now()
		htmlStr := fmt.Sprintf("<p><h3>本卦解答： %s </h3> 上%d下%d => %s%s (%s)十进制%d，动爻是下起第%d爻</p><p>The divination result is <a href='https://www.zhouyi.cc/zhouyi/yijing64/4%s.html' target='_blank'> %s </a> .</p>", anser, num1, num2, words1, words2, words3, numOutput+1, num3, words4, anser)
		//使用fmt.Sprintf函数将信息格式化为字符串
		tmpl, _ := template.New("t").Parse(htmlStr)
		pl(now.Format("01-02 15:04:05"), "收到POST:", num1, num2, num3, "-", anser)
		tmpl.Execute(w, nil)
	}

	pl("web start")

	http.HandleFunc("/", h1)
	http.HandleFunc("/please", h2)
	http.HandleFunc("/divination", h3)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	//online change to 80,线上部署是改为：80
	log.Fatal(http.ListenAndServe(":8000", nil))

}
