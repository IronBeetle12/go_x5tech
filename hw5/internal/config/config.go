package config

import "os"

const defaultHTTPAddr = ":8080"

func HTTPAddr() string {
	if addr := os.Getenv("HTTP_ADDR"); addr != "" {
		return addr
	}
	return defaultHTTPAddr
}
