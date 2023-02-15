package controllers

import (
	"database/sql"
	"log"
	"mbanking_project/entities"
)

func Login(db *sql.DB, logUser entities.User) (entities.User, error) {
	loginQry, err := db.Prepare("SELECT id FROM users WHERE phone = ? AND password = ?")
	if err != nil {
		log.Fatal("error prepare insert", err.Error())
	}

	row := loginQry.QueryRow(logUser.Telepon, logUser.Password)

	if row.Err() != nil {
		log.Println("login query ", row.Err().Error())
	}
	res := entities.User{}
	err = row.Scan(&res.Id)

	if err != nil {
		log.Println("after login query ", err.Error())
		return res, err
	}
	// res.Telepon = phone

	return res, nil
}
