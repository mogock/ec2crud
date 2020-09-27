package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"mogock.com/ec2crud/pkg/models"
)

type PersonModel struct {
	DB *sql.DB
}

func (m *PersonModel) Insert(cedula, nombre, apellido string) (int, error) {
	fmt.Println("Insert new Person")
	return 0, nil
}

func (m *PersonModel) Get(cedula int) (*models.Person, error) {
	fmt.Println("Get Person")
	return nil, nil
}

func (m *PersonModel) Latest()([]*models.Person, error){
	fmt.Println("List of Person")

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
			log.Println("Is not working Propertly")
			return nil, err
		}
		pList = append(pList, p)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error Trying to read the result")
		return nil, err
	}

	return pList, nil
}


