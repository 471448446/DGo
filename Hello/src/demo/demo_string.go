package demo

import (
	"fmt"
	"strconv"
	"strings"
)

const H = "H"

func RunString() {
	DemoTitle("String")

	hello := "Hello Go,Hello Wold"
	//fmt.Printf()
	//fmt.Springs()

	fmt.Println(fmt.Sprintf("string \"%s\" has prefix %s?", hello, H),
		strings.HasPrefix(hello, H))
	fmt.Println(fmt.Sprintf("string \"%s\" has suffix %s?", hello, H),
		strings.HasSuffix(hello, H))
	fmt.Printf("string \"%s\" contain %s? %t\n", hello, H,
		strings.Contains(hello, H))
	fmt.Printf("string \"%s\" count %s = %d\n", hello, H,
		strings.Count(hello, H))
	fmt.Printf("string \"%s\" %s, index =  %d ,lastIndex = %d\n", hello, H,
		strings.Index(hello, H), strings.LastIndex(hello, H))
	// 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，
	// 如果 n = -1 则替换所有字符串 old 为字符串 new
	fmt.Printf("string \"%s\" replace %s = %s,origin = %s\n", hello, H,
		strings.Replace(hello, H, H+H, -1), hello)
	//用于重复 count 次字符串 s 并返回一个新的字符串
	fmt.Printf("string \"%s\" repeat 2 = %s\n", hello, strings.Repeat(hello, 2))
	fmt.Printf("string \"%s\" up = %s,low = %s,origin = %s \n",
		hello, strings.ToUpper(hello), strings.ToLower(hello), hello)
	//Trim = TrimLeft
	fmt.Printf("string \"%s\" trimSpace = %s , trim = %s ,trimL = %s ,trimR = %s\n", hello,
		strings.TrimSpace(hello, ), strings.Trim(hello, "H"),
		strings.TrimLeft(hello, "H"), strings.TrimRight(hello, "H"))
	// 利用空白符号分割字符串
	var spit = strings.Split(hello, ",")
	fmt.Printf("Fields() = %s, spilt = %s size = %d \n", strings.Fields(hello), spit, len(spit))
	num, error := strconv.Atoi("2019.")
	fmt.Println(num, error)

}
