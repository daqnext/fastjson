package main

import (
	"fmt"

	FastJson "github.com/daqnext/fastjson"
)

func main() {
	fj, err := FastJson.NewFromFile("../test.json")
	if err == nil {
		result, _ := fj.GetString("company", "name")
		fmt.Println(result)
	}
}
