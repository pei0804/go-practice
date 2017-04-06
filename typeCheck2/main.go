package main

import "fmt"

func main() {
	var atai1 intType = 3
	var atai2 intType = 1
	fmt.Println(add(atai1, atai2))
}

type intType int

func add(args ...interface{}) intType {
	var ans intType
	for _, v := range args {
		ans += v.(intType)
	}
	return ans
}
