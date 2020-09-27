package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"mogock.com/ec2crud/pkg/models"
)

type PersonModel struct {
	DB *sql.DB
}

func (m *PersonModel) Insert(cedula, nombre, apellido string) (int64, error) {
	fmt.Println("Insert new Person")
	if cedula == "" || nombre == "" || apellido == ""{
		return -1, errors.New("Please specify all the required fields")
	}

	result , err := m.DB.Exec("insert into person values($1, $2, $3)", cedula, nombre, apellido)
	if err != nil {
		return -1, err
	}

	rowAffected, e := result.RowsAffected()
	if e != nil {
		return -1, e
	}
	return rowAffected, nil
}

func (m *PersonModel) Get(cedula string) (*models.Person, error) {
	if len(cedula) < 9 {
		log.Println("Cedula bad format")
		return nil, errors.New("Invalid Cedula Format")
	}
	row := m.DB.QueryRow("SELECT * FROM person WHERE cedula = $1", cedula)
	p := new(models.Person)
	err := row.Scan(&p.Cedula, &p.Nombre, &p.Apellido)
	if err == sql.ErrNoRows {
		return nil, err
	}else if err != nil {
		return nil, err
	}
	return p, nil
}

func (m *PersonModel) Latest()([]*models.Person, error){
	rows, err := m.DB.Query("SELECT * FROM person")

	if err != nil {
		log.Println("Error trying to get List of Person")
		return nil, err
	}
	defer rows.Close()

	pList := make([]*models.Person, -0)

	for rows.Next() {
		p := new(models.Person)
		err := rows.Scan(&p.Cedula, &p.Nombre, &p.Apellido)
		if err != nil{
			log.Println("error reading result set properties")
			return nil, err
		}
		pList = append(pList, p)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in the result")
		return nil, err
	}

	return pList, nil
}
