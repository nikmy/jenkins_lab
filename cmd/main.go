package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func formatResponse(addr string) []byte {
	const responseFmt = `<html><body><p>Your IP address is %s</p></body></html>`
	return []byte(fmt.Sprintf(responseFmt, addr))
}

func main() {
	http.HandleFunc("/get_ip", func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write(formatResponse(readUserIP(req)))
		if err != nil {
			slog.Error("cannot send response: %s", err)
		}
	})
	err := http.ListenAndServe("", nil)
	if err != nil {
		slog.Error("run server: %s", err)
	}
}
