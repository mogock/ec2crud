package main

import (
	"net/http"
	"mogock.com/ec2crud/cmd/web/handler"
)

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/person/insert", handler.InsertPerson)
	http.HandleFunc("/person/getByCedula", handler.GetPerson)
	http.HandleFunc("/person/all", handler.GetAllPerson)
	http.ListenAndServe(":3000", nil)
}