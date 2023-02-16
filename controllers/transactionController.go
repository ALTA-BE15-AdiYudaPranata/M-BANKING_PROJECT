package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"mbanking_project/entities"
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

func Transfer(db *sql.DB, idUser int, phoneOther string, nominal int) (bool, error) {
	// proses untuk mendapatkan nomer telepon
	querySaldo2, errSaldo2 := db.Prepare("SELECT id FROM users WHERE phone = ?")
	if errSaldo2 != nil {
		log.Fatal("error query select", errSaldo2.Error())
	}

	row2 := querySaldo2.QueryRow(phoneOther)

	if row2.Err() != nil {
		log.Println("select query ", row2.Err().Error())
	}
	res2 := entities.User{}
	err2 := row2.Scan(&res2.Id)

	if err2 != nil {
		log.Println("after select query ", err2.Error())
		return false, err2
	}

	// proses transfer

	transferQry, errTransfer := db.Prepare("INSERT INTO Transfer (user_id_pengirim, user_id_penerima, value) VALUES (?,?,?)")
	if errTransfer != nil {
		log.Fatal("error query insert", errTransfer.Error())
	}

	result, err := transferQry.Exec(idUser, res2.Id, nominal)
	if err != nil {
		log.Fatal("error exec insert", err.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("proses transfer berhasil")
		} else {
			fmt.Println("proses gagal")
		}
	}

	// update saldo pengirim
	querySaldo11, errSaldo11 := db.Prepare("UPDATE Users set saldo = saldo - ? where id = ? ")
	if errSaldo11 != nil {
		log.Fatal("error query update", errSaldo11.Error())
	}

	statement, err := querySaldo11.Exec(nominal, idUser)
	if err != nil {
		log.Fatal("error exec update", err.Error())
	} else {
		row, _ := statement.RowsAffected()
		if row > 0 {
			fmt.Println("")
		} else {
			fmt.Println("proses gagal")
		}
	}

	// update saldo penerima
	querySaldo21, errSaldo21 := db.Prepare("UPDATE Users set saldo = saldo + ? where id = ? ")
	if errSaldo21 != nil {
		log.Fatal("error query update", errSaldo21.Error())
	}

	statement1, err := querySaldo21.Exec(nominal, res2.Id)
	if err != nil {
		log.Fatal("error exec update", err.Error())
	} else {
		row, _ := statement1.RowsAffected()
		if row > 0 {
			fmt.Println("")
		} else {
			fmt.Println("proses gagal")
		}
		return false, err
	}
	return true, nil
}

func HistoryTopup(db *sql.DB, idUser int) {
	rows, errSelect := db.Query("SELECT value, created_at FROM TopUp WHERE user_id = ?", idUser)
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allHistory []entities.TopUp
	for rows.Next() {
		var datarow entities.TopUp
		errScan := rows.Scan(&datarow.Value, &datarow.Created)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())

		}
		allHistory = append(allHistory, datarow)
		fmt.Println("Nominal:", datarow.Value, "Tgl TopUp:", datarow.Created)
	}

}
