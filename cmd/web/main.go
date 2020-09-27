package main

import (
	"net/http"
	"mogock.com/ec2crud/cmd/web/handler"
)

func main() {
	http.HandleFunc("/", handler.Home)
	http.ListenAndServe(":3000", nil)
}