package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"mbanking_project/entities"
)

func Register(db *sql.DB, newUser entities.User) {
	var query = "INSERT INTO users(name, phone, password) VALUES(?,?,?)"
	registerQry, err := db.Prepare(query)
	if err != nil {
		log.Fatal("error prepare insert", err.Error())
	}

	result, err := registerQry.Exec(newUser.Nama, newUser.Telepon, newUser.Password)
	if err != nil {
		log.Fatal("error exec insert", err.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("proses berhasil dijalankan")
		} else {
			fmt.Println("proses gagal")
		}
	}
}

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

func UpdateUser(db *sql.DB, nama string, telepon string, password string, id int) (entities.User, error) {
	updateQry, errUpdate := db.Prepare("UPDATE users SET name = ?, phone = ?, password = ? WHERE id = ?")
	if errUpdate != nil {
		log.Fatal("error query select", errUpdate.Error())
	}

	result, err := updateQry.Exec(nama, telepon, password, id)
	if err != nil {
		log.Fatal("error exec insert", err.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("proses berhasil dijalankan")
		} else {
			fmt.Println("proses gagal")
		}
	}
	res := entities.User{}
	// err = result.Scan(&res.Id)
	if err != nil {
		log.Println("after update query ", err.Error())
		return res, err
	}
	return res, nil
}
