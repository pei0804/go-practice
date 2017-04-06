package main

import (
	"fmt"
)

const (
	a stringType = "A文字列だぜ"
	b stringType = "B文字列だぜ"
	c intType    = 1
)

type stringType string
type intType int

func main() {
	typeCheck(a)
	typeCheck(b)
	typeCheck(c)
}

func typeCheck(stringpoi interface{}) {
	var (
		stringdayo stringType
		ok         bool
	)
	if stringdayo, ok = stringpoi.(stringType); !ok {
		fmt.Println(c)
		fmt.Println("こいつはStringTypeじゃねえ！")
		return
	}
	fmt.Println(stringdayo)
	fmt.Println("ええやん")
	fmt.Println("------------")
}
