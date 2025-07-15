package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	for {
		resp, err := http.Get("http://go-server-svc:8080")
		if err != nil {
			fmt.Println("Error:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(string(body))
		time.Sleep(5 * time.Second)
	}
}
