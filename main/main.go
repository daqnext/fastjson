package main

import (
	"fmt"

	FastJson "github.com/daqnext/fastjson"
)

func main() {

	fj, _ := FastJson.NewFromString("{}")

	str := `avasf1er李龙固额""[][""""\\\\asdf 包*&*^)`
	fj.SetString(str, "xxx")

	result, _ := fj.GetString("xxx")
	fmt.Println(result)

	err := fj.ClearFileAndOutput("./test2.json")
	if err != nil {
		fmt.Println(err)
		//t.Error("something wrong", err)
	}

}
