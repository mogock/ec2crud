package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")

type Person struct {
	Cedula string
	Nombre string
	Apellido string
}

