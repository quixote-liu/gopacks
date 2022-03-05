package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	data := []map[string]interface{}{
		{"name": "liu", "age": 15, "email": "123@foxmail.com"},
		{"name": "liu", "age": 15, "email": "123@foxmail.com"},
		{"name": "liu", "age": 15, "email": "123@foxmail.com"},
		{"name": "liu", "age": 15, "email": "123@foxmail.com"},
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	var payload interface{}
	if err := json.Unmarshal(dataBytes, &payload); err != nil {
		panic(err)
	}

	resources, ok := payload.([]interface{})
	if !ok {
		fmt.Println("this is not []interface{}")
		return
	}

	for _, r := range resources {
		_, ok := r.(map[string]interface{})
		if !ok {
			fmt.Println("this is not map[string]interface{}")
			return
		}
	}
	fmt.Println("END")
}

func GinTest() {
	r := gin.New()

	r.GET("/hello", func(c *gin.Context) {
		data := []map[string]interface{}{
			{"name": "liu", "age": 15, "email": "123@foxmail.com"},
			{"name": "liu", "age": 15, "email": "123@foxmail.com"},
			{"name": "liu", "age": 15, "email": "123@foxmail.com"},
			{"name": "liu", "age": 15, "email": "123@foxmail.com"},
		}

		c.JSON(http.StatusOK, data)
	})

	r.Run(":8989")
}
