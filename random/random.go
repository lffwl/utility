package random

import (
	"math/rand"
	"strings"
)

var Dics = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
}

// RandAllString 生成随机字符串([a~zA~Z0~9])
func RandAllString(lenNum int) string {
	str := strings.Builder{}
	length := len(Dics)
	for i := 0; i < lenNum; i++ {
		l := Dics[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

// RandNumString  生成随机数字字符串([0~9])
func RandNumString(lenNum int) string {
	str := strings.Builder{}
	length := 10
	for i := 0; i < lenNum; i++ {
		str.WriteString(Dics[52+rand.Intn(length)])
	}
	return str.String()
}

// RandString  生成随机字符串(a~zA~Z])
func RandString(lenNum int) string {
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		str.WriteString(Dics[rand.Intn(length)])
	}
	return str.String()
}
