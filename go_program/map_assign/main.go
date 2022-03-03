package main

import "fmt"

func main() {
	t := map[string]interface{}{}

	m := reserveFields(t)
	fmt.Println("len m=", len(m))
	fmt.Println("m=", m)

	var t2 map[string]interface{}
	fmt.Println("len t2", len(t2))
	fmt.Println("testing=", t2["testing"])
}

func reserveFields(t map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	if t == nil {
		return m
	}

	m["id"] = t["id"]
	m["network_id"] = t["network_id"]
	m["fixed_ips"] = t["fixed_ips"]
	m["device_id"] = t["device_id"]
	m["device_owner"] = t["device_owner"]

	return m
}
