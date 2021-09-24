package main

import (
	"fmt"

	FastJson "github.com/daqnext/fastjson"
)

func main() {
	fj, err := FastJson.NewFromFile("../test2.json")
	if err == nil {
		result, err := fj.GetIntArray("x")
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println(result[0])
		}

		result2, err2 := fj.GetFloat64Array("y")
		if err2 != nil {
			panic(err2.Error())
		} else {
			fmt.Println(result2[0])
		}

		result3, err3 := fj.GetStringArray("a")
		if err3 != nil {
			panic(err3.Error())
		} else {
			fmt.Println(result3[0])
		}
	}
}
