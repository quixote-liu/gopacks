package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	object := map[string]interface{}{
		"错误": "错误信息",
	}
	res, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}