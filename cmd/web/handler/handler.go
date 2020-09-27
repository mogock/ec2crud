package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)
import "mogock.com/ec2crud/pkg/models/postgres"

/*
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	person *postgres.PersonModel
}
 */

//var person = *postgres.PersonModel

var db *sql.DB

//Stard Database Connection
func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://gosta:86Chota34..Ene@gosta.c3de3fnfq0uc.us-east-1.rds.amazonaws.com:5432/gosta?sslmode=disable")	//Se usa igual si las variable fueron creadas
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err = db.Ping(); err != nil { //Muy Interensante
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	//p := &postgres.PersonModel{DB: db}
	//content, error := p.Latest();
	//cp, err := p.Get("22600051399")
	/*cp, err := p.Insert("52300051875", "Rosa", "Hernandez")
	if err != nil {
		log.Println("And error happens")
		return
	}
	fmt.Println(cp)
	*/
	w.Write([]byte("This is my home page", ))
}
