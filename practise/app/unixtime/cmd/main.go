package main

import (
	// 将项目放在 GOPATH/src/ 目录下面
	"DGo/practise/library/log"
	"errors"
	"os"
	"strconv"
	"time"
)

const ErrorParse = "error parse time"

var TAG = " unixApp"
//https://gobyexample.com/time-formatting-parsing
func main() {
	input := os.Args[1]

	loc, _ := time.LoadLocation("Asia/Chongqing")
	log.Info(TAG, "Now In BJ : ", time.Now().In(loc).Format("2006-01-02 15:04:05"))

	//input := "1555971410000"
	//input := "1555971410"
	//input := "2019-04-22 22:16:50"

	if s, e := formatTimeStamp(input); nil == e {
		log.Info(TAG, input, " => ", s)
		return
	}
	if s, e := toTimeStampSecond(input); nil != e {
		log.Info(TAG, e)
	} else {
		log.Info(TAG, input, " => ", s)
		// test
		/*tim, _ := formatTimeStamp(strconv.FormatInt(s, 10))
		log.Info(TAG, s, " __ ", tim)*/
	}
}

// 输入字符串，转化成秒
func toTimeStampSecond(s string) (int64, error) {
	loc, _ := time.LoadLocation("Asia/Chongqing")
	t, e := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	if nil != e {
		return 0, errors.New(ErrorParse)
	}
	//t.Unix() -> 秒
	//t.UnixNano() -> 纳秒
	//t.UnixNano()/1e6 -> 毫秒
	//t.UnixNano()/1e9 -> 纳秒转换成秒
	return t.Unix(), nil
}

// 输入数字，格式化成字符串
func formatTimeStamp(s string) (string, error) {
	//s := "1555913810362"
	//参数1 数字的字符串形式
	//参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	//参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	timeOfSecond, e := strconv.ParseInt(s, 10, 64)
	if nil != e {
		return "", errors.New(ErrorParse)
	}

	switch expr := len(s); expr {
	case 10:
	case 13:
		timeOfSecond /= 1e3
	case 19:
		timeOfSecond /= 1e9
	default:
		return "", errors.New(ErrorParse)
	}

	t1 := time.Unix(timeOfSecond, 0)
	loc, _ := time.LoadLocation("Asia/Chongqing")
	format := t1.In(loc).Format("2006-01-02 15:04:05")
	return format, nil
}
