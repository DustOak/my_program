package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 百度显示 redis协议 共有一下几个格式
// 1. 简单字符串 Simple Strings, 以 "+"加号 开头
//      格式：+ 字符串 \r\n

// 2. 错误 Errors, 以"-"减号 开头
// 　　格式：- 错误前缀 错误信息 \r\n

// 3. 整数型 Integer， 以 ":" 冒号开头
// 　　格式：: 数字 \r\n
//
//      eg: ":1000\r\n"
//
// 4. 大字符串类型 Bulk Strings, 以 "$"美元符号开头，长度限制512M
// 　　格式：$ 字符串的长度 \r\n 字符串 \r\n
//
// 5. 数组类型 Arrays，以 "*"星号开头
// 　　格式：* 数组元素个数 \r\n 其他所有类型 (结尾不需要\r\n)

//
//           "*2\r\n$2\r\nfoo\r\n$3\r\nbar\r\n"      数组包含2个元素，分别是字符串foo和bar
//
// 　　　　"*3\r\n:1\r\n:2\r\n:3\r\n"       数组包含3个整数：1、2、3
//
//           "*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$6\r\nfoobar\r\n"  包含混合类型的数组
//
//
//           "*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Foo\r\n-Bar\r\n"   数组嵌套，外层数组包含2个数组，整理后如下：

func GetSimpleStringsOrErrorStrings(str string) string {
	//简单字符和异常字符 前缀和后缀 均为1字节 和两字节 所以 只截取1-len(str)-2 长度字节
	return str[1 : len(str)-2]
}

func GetIntegerStrings(str string) int {
	//整形类型字符 基本截取方式与简单字符相同 只是多了转换为整形
	temp, err := strconv.Atoi(str[1 : len(str)-2])
	if err != nil {
		log.Fatal("整形数据异常")
		os.Exit(0)
	}
	return temp
}

func GetBigString(str string) string {
	//大字符串先将字符串使用\r\n分割 然后返回下标为1开始的字符串 (并不知道 长度是做什么的)
	return strings.Join(strings.Split(str, "\r\n")[1:], "")
}

func GetArrayString(str string) []interface{} {
	//首先切割字符串 判断长度
	array := strings.Split(str, "\r\n")
	//创建切片
	as := make([]interface{}, 0)
	//循环判断属于那种字符串(简单 异常 整形 大字符)
	for i := 1; i < len(array)-1; i++ {
		switch array[i][0] {
		//当数组嵌套时 递归调用
		case '*':
			temp, err := strconv.Atoi(fmt.Sprintf("%c", array[i][1]))
			if err != nil {
				log.Fatal("整形数据异常")
				os.Exit(0)
			}
			a := make([]string, 0)
			for j := 0; j <= temp; j++ {
				a = append(a, fmt.Sprintf("%s\r\n", array[i+j]))
			}
			as = append(as, GetArrayString(strings.Join(a, "")))
			i += temp
			//简单和异常 整形处理方式相同
		case '+':
			fallthrough
		case '-':
			as = append(as, GetSimpleStringsOrErrorStrings(fmt.Sprintf("%s\r\n", array[i])))
		case ':':
			as = append(as, GetIntegerStrings(fmt.Sprintf("%s\r\n", array[i])))
			//大字符串将长度与字符串传送给相应函数 用于切割
		case '$':
			as = append(as, GetBigString(fmt.Sprintf("%s\r\n%s\r\n", array[i], array[i+1])))
			i += 1
		default:
			log.Fatal("协议参数格式不正确")
			os.Exit(0)
		}
	}
	return as
}
func main() {
	str := "+OK\r\n"
	switch str[0] {
	case '+':
		fallthrough
	case '-':
		fmt.Println(GetSimpleStringsOrErrorStrings(str))
	case '$':
		fmt.Println(GetBigString(str))
	case ':':
		fmt.Println(GetIntegerStrings(str))
	case '*':
		fmt.Println(GetArrayString(str))
	default:
		log.Fatal("协议参数格式不正确")
		os.Exit(0)
	}
}
