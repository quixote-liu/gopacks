package main

import "fmt"

func main() {
	var inter interface{}
	var mnil map[string]interface{}
	inter = mnil
	res, ok := inter.(map[string]interface{})
	fmt.Println("ok:", ok)
	fmt.Println("res:", res)
	fmt.Println("res==nil:", res == nil)

	var inter2 interface{} = 2
	res2 := inter2.(string)
	fmt.Printf("res2: %v\n", res2)
}
