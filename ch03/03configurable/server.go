package main

import "net/http"

func main() {
	sever := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	sever.ListenAndServe()
}
