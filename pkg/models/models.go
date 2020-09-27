package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")

type Person struct {
	Cedula string `json:cedula`
	Nombre string `json:nombre`
	Apellido string `json:apellido`
}

type StatusMessage struct {
	Code int16 `json:code`
	Message string `json:message`
	Error error
}

