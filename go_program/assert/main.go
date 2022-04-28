package main

import (
	"fmt"

	"github.com/google/uuid"
)

type ip struct {
	SubnetID string `json:"subnet_id"`
}

func main() {
	// demo2()
	demo3()
}

func demo2() {
	var m interface{}

	m = "string"

	fmt.Println(m.(bool))
}

func demo1() {
	m := make(map[string]interface{})

	m["ips"] = []ip{
		{SubnetID: uuid.NewString()},
		{SubnetID: uuid.NewString()},
	}

	ips, ok := m["ips"].([]interface{})
	if !ok {
		panic("find ips failed")
	}
	fmt.Println("ips=", ips)
}

func demo3() {
	// var jsonData string = `{"ports":["a","b","c"]}`
	payload := map[string]interface{}{
		"ports": [3]interface{}{"a", "b", "c"},
	}
	// if err := json.Unmarshal([]byte(jsonData), &payload); err != nil {
	// 	panic(err)
	// }

	fmt.Printf("%T\n", payload["ports"])

	ss, ok := payload["ports"].([3]string)
	if !ok {
		fmt.Println("no []string")
		return
	}
	fmt.Println("is []string")
	fmt.Println("ss=", ss)
}
