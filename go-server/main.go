package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		podName := os.Getenv("POD_NAME")
		ip := getLocalIP()
		fmt.Fprintf(w, "Hello from pod: %s\nIP: %s\n", podName, ip)
	})

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func getLocalIP() string {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "unknown"
}
