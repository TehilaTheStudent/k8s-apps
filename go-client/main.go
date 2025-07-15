package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	serverURL := os.Getenv("SERVER_URL")
	if serverURL == "" {
		serverURL = "http://go-server-svc:8080"
	}
	errorSleep := 2 * time.Second
	if val := os.Getenv("ERROR_SLEEP"); val != "" {
		if d, err := time.ParseDuration(val); err == nil {
			errorSleep = d
		}
	}
	loopSleep := 5 * time.Second
	if val := os.Getenv("LOOP_SLEEP"); val != "" {
		if d, err := time.ParseDuration(val); err == nil {
			loopSleep = d
		}
	}
	for {
		resp, err := http.Get(serverURL)
		if err != nil {
			fmt.Println("Error:", err)
			time.Sleep(errorSleep)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(string(body))
		time.Sleep(loopSleep)
	}
}
