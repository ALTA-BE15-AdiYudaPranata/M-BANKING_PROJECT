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

func DeleteUser(db *sql.DB, delUser entities.User) (bool, error) {
	deleteQry, errDelete := db.Prepare("DELETE FROM users WHERE phone = ? AND password = ?")
	if errDelete != nil {
		log.Fatal("error query select", errDelete.Error())
	}

	result, err := deleteQry.Exec(delUser.Telepon, delUser.Password)
	if err != nil {
		log.Fatal("error exec delete", err.Error())

	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("akun telah dihapus dari sistem")
		} else {
			fmt.Println("proses gagal")
		}
	}
	return true, nil

}

func ReadData(db *sql.DB, telp string, pass string) {
	rows, errSelect := db.Query("SELECT id, name, phone, saldo, created_at FROM Users WHERE phone = ? AND password = ?", telp, pass)
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	for rows.Next() {
		var datarow entities.User
		errScan := rows.Scan(&datarow.Id, &datarow.Nama, &datarow.Telepon, &datarow.Saldo, &datarow.Created)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		fmt.Println("Id:", datarow.Id, "Nama:", datarow.Nama, "Telepon:", datarow.Telepon, "Saldo:", datarow.Saldo, "Tgl Pembuatan Akun:", datarow.Created)
	}

}
