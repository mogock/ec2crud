package handler

import (
	"database/sql"
	json2 "encoding/json"
	//xml2 "encoding/xml"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"mogock.com/ec2crud/pkg/models/postgres"
	"mogock.com/ec2crud/pkg/models"
)

var db *sql.DB

//Stard Database Connection
func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://gosta:86Chota34..Ene@gosta.c3de3fnfq0uc.us-east-1.rds.amazonaws.com:5432/gosta?sslmode=disable")	//Se usa igual si las variable fueron creadas
	if err != nil {
		log.Fatal(err)
	}

	//Connection Pool Setting
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Server is running"))
}

func InsertPerson(w http.ResponseWriter, request *http.Request) {
	var p models.Person
	err := json2.NewDecoder(request.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn := &postgres.PersonModel{DB: db}
	cp, err := conn.Insert(p.Cedula, p.Nombre, p.Apellido)
	if err != nil {
		log.Println("And error happens")
		return
	}
	fmt.Fprintf(w, "Person: %+v", cp)
}

func GetPerson (w http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	cedula := request.FormValue("cedula")
	if cedula == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	pModel := &postgres.PersonModel{DB: db}
	person, error := pModel.Get(cedula);
	if error != nil {
		http.Error(w, "failed", 500)
		return
	}

	jsonRosponse, err := json2.Marshal(person)
	if err != nil {
		http.Error(w, "failed", 500)
		return
	}
	w.Write(jsonRosponse)
}

func GetAllPerson(w http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	pModel := &postgres.PersonModel{DB: db}
	persons, error := pModel.Latest();
	if error != nil {
		http.Error(w, "failed", 500)
		return
	}
	jsonRosponse, err := json2.Marshal(persons)
	if err != nil {
		http.Error(w, "failed", 500)
		return
	}
	w.Write(jsonRosponse)
}