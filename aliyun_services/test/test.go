package main

import (
	"fmt"
	"io/ioutil"
	"sync/atomic"

	"net/http"
)

func main() {
	t := int64(0)
	atomic.AddInt64(&t, 1)
	fmt.Println(t)
	fmt.Println(atomic.LoadInt64(&t))
	atomic.StoreInt64(&t, 0)
	fmt.Println(t)
}

func get_url() {
	t := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NTRmNWFkOTMyNWM0YjVkYmUxNTJmZmVkNDg4YWUyZCIsImN1c3RvbUpzb24iOiJ7XCJjbGllbnRJZFwiOlwiMjVkelgzdmJZcWt0Vnh5WFwiLFwiZG9tYWluSWRcIjpcImJqMjlcIixcInNjb3BlXCI6W1wiRFJJVkUuQUxMXCIsXCJTSEFSRS5BTExcIixcIkZJTEUuQUxMXCIsXCJVU0VSLkFMTFwiLFwiVklFVy5BTExcIixcIlNUT1JBR0UuQUxMXCIsXCJTVE9SQUdFRklMRS5MSVNUXCIsXCJCQVRDSFwiLFwiT0FVVEguQUxMXCIsXCJJTUFHRS5BTExcIixcIklOVklURS5BTExcIixcIkFDQ09VTlQuQUxMXCIsXCJTWU5DTUFQUElORy5MSVNUXCJdLFwicm9sZVwiOlwidXNlclwiLFwicmVmXCI6XCJodHRwczovL3d3dy5hbGl5dW5kcml2ZS5jb20vXCIsXCJkZXZpY2VfaWRcIjpcIjBkZDg1MDNjMjViZTQxYjI5ZWMxMDBjNjhkYzQ5ZjAzXCJ9IiwiZXhwIjoxNjUyNjE2ODQ2LCJpYXQiOjE2NTI2MDk1ODZ9.YiBavMK_rg-OIvrSiD-sD71lukhLfZLthOkJ9Zd_FRTeM41Jn7-skVBVqJneugwqm67lhJCB3_g2SHijTM6WI6_BQ5uyuFfb_XMt3z88xoWjnyUb6XHKWAZifabRRJEhEAGAZsLD3SGwAWHGjc1j4nHInLAqyGnzXeDOcAlC2Uc"
	headers := map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39",
		"Referer":         "https://www.aliyundrive.com/",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Host":            "api.aliyundrive.com",
	}
	headers["Content-Type"] = "application/json; charset=utf-8"
	headers["Authorization"] = "Bearer " + t

	c := http.Client{}
	r, err := http.NewRequest("POST", "https://api.aliyundrive.com/v2/file/get_download_url", nil)
	if err != nil {
		panic(err)
	}
	for i, v := range headers {
		r.Header.Set(i, v)
	}
	resp, err := c.Do(r)
	re, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s", re)

}
