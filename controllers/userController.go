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
