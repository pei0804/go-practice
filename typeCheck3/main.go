package main

import (
	"fmt"

	"github.com/labstack/gommon/log"
)

func main() {
	var name nameType = "ぺい"
	var term termType = "こんちは"
	fmt.Println(greet(name, term))
	fmt.Println(greet(name))
}

type nameType string
type termType string

func greet(args ...interface{}) string {
	var (
		name nameType
		term termType
		ok   bool
	)

	parseNameType := func(index int) {
		if name, ok = args[index].(nameType); !ok {
			fmt.Println(args[index])
			fmt.Println("これはあかんで")
			fmt.Println("-----------")
		}
	}

	parseGreetType := func(index int) {
		if term, ok = args[index].(termType); !ok {
			fmt.Println(args[index])
			fmt.Println("これはあかんで")
			fmt.Println("-----------")
		}
	}

	switch len(args) {
	case 1:
		parseNameType(0)
	case 2:
		parseNameType(0)
		parseGreetType(1)
	default:
		log.Error("少ないか多すぎるで、1から2個やで")
	}
	return fmt.Sprintf("%vです %v", name, term)
}
