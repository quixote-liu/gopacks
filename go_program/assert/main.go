package main

import (
	"fmt"

	"github.com/google/uuid"
)

type ip struct {
	SubnetID string `json:"subnet_id"`
}

func main() {
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
