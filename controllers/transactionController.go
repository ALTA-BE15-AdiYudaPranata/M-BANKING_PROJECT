package controllers

import (
	"database/sql"
	"fmt"
	"log"
)

func TopUp(db *sql.DB, id int, nominal int) (saldo int, status error) {
	topupQry, errTopup := db.Prepare("INSERT INTO TopUp (user_id, value) VALUES (?,?)")
	if errTopup != nil {
		log.Fatal("error query insert", errTopup.Error())
	}

	result, err := topupQry.Exec(id, nominal)
	saldo = saldo + nominal
	if err != nil {
		log.Fatal("error exec insert", err.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			rss, err := db.Prepare("UPDATE Users SET saldo = saldo + ? WHERE id = ?")
			if err != nil {
				log.Fatal("error query insert", err.Error())
			}
			result, err := rss.Exec(saldo, id)
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

		} else {
			fmt.Println("proses gagal")
		}
	}

	if err != nil {
		log.Println("after insert query ", err.Error())
		return saldo, err
	}
	return saldo, nil
}
