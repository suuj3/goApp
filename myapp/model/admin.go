package model

import postgres "myapp/datastore/postgres"

type Admin struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

const queryInsertAdmin = "INSERT INTO admin(firstname, lastname, email, password) VALUES($1, $2, $3, $4);"

func (adm *Admin) Create() error {
	_, err := postgres.Db.Exec(queryInsertAdmin, adm.FirstName, adm.LastName, adm.Email, adm.Password)
	return err
}

const queryGetAdmin = "SELECT email, password FROM admin WHERE email=$1 and password=$2;"

func (adm *Admin) Get() error {
	return postgres.Db.QueryRow(queryGetAdmin, adm.Email, adm.Password).Scan(&adm.Email, &adm.Password)
}
