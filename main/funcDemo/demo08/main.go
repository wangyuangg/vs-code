package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 	/* str := "hello北"
	// 	fmt.Println("str的长度", len(str))
	// 	str2 := "hello北京"
	// 	str3 := []rune(str2)
	// 	for i := 0; i < len(str3); i++ {
	// 		fmt.Printf("%c\n", str3[i])
	// 	}
	//  */
	n, err := strconv.Atoi("123")//字符串转整型
	if err != nil {
		fmt.Println("转换失败", err)
	} else {
		fmt.Println("转换成功", n)
	}

	var bytes = []byte("hello")//字节数组
	fmt.Println(bytes)

	str1 := string([]byte{97, 98, 99})//字符串转字节数组
	fmt.Println(str1)
	b := strings.Contains("abc", "b")//判断是否包含
	fmt.Println(b)
	i := strings.Count("abc", "a")//统计字符串中某个字符出现的次数
	fmt.Println(i)
	b2 := strings.EqualFold("abc", "Abc")//比较两个字符串是否相等，忽略大小写
	fmt.Println(b2)
	fmt.Printf("strings.Index(\"NTL_abc\", \"abc\"): %v\n", strings.Index("NTL_abc", "abc"))
	//返回第一个匹配的字符串的索引
	fmt.Printf("strings.Replace(\"NTL_abc\", \"abc\", \"ABC\", 1): %v\n", strings.Replace("NTL_abc_abc", "abc", "ABC", -1))
	//替换字符串
	s := strings.Split("NTL_abc_abc", "_")
	//按照指定的分隔符将字符串分割成字符串切片
	for _, v := range s {
		fmt.Println(v)
	}
	s2 := strings.ToUpper("NTL_abc_abc")//将字符串转换为大写
	fmt.Println(s2)
	s3 := strings.Trim("! hello! ", " !")//去除字符串两边的空格
	fmt.Println(s3)
	b3 := strings.HasPrefix("ftp://192.168.10.1", "ftp")//判断字符串是否以指定的前缀开头
	fmt.Println(b3)
}
