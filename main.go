package main

import (
	"fmt"
	"mbanking_project/config"
	"mbanking_project/controllers"
	"mbanking_project/entities"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := config.ConnectToDB()
	defer db.Close()

	inputMenu := 1

	for inputMenu != 0 {
		fmt.Println("-- Selamat Datang di M-Banking Project --")
		fmt.Println("Pilih Menu:\n1. Register\n2. Login\n0. Exit")
		fmt.Println("Masukkan pilihan anda : ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			// code here - REGISTER
			newUser := entities.User{}
			fmt.Println("Masukkan Nama:")
			fmt.Scanln(&newUser.Nama)
			fmt.Println("Masukkan Telepon:")
			fmt.Scanln(&newUser.Telepon)
			fmt.Println("Masukkan Password:")
			fmt.Scanln(&newUser.Password)
			controllers.Register(db, newUser)
		}
		if inputMenu == 2 {
			// code here - LOGIN
			logUser := entities.User{}
			fmt.Println("Masukkan Telepon:")
			fmt.Scanln(&logUser.Telepon)
			fmt.Println("Masukkan Password:")
			fmt.Scanln(&logUser.Password)
			dataLogin, err := controllers.Login(db, logUser)
			fmt.Println(dataLogin)
			if err != nil {
				fmt.Println("login gagal")
			} else {
				loginMenu := 0
				fmt.Println("login berhasil")
				fmt.Println("===============================")
				fmt.Println("----- Menu User -----")
				fmt.Println("Pilih Menu : ")
				fmt.Println("Pilih Menu:\n1. Read Account\n2. Update Account\n3. Delete Account\n4. Top-Up\n5. Transfer\n6. History Top-Up\n7. History Transfer\n8. Read Other Users\n0. Exit")
				fmt.Println("Masukkan pilihan anda : ")
				fmt.Scanln(&loginMenu)
				switch loginMenu {
				case 1:
					// code here - READ ACCOUNT
					var telp, pass string
					fmt.Println("Masukkan Telepon:")
					fmt.Scanln(&telp)
					fmt.Println("Masukkan Password:")
					fmt.Scanln(&pass)
					controllers.ReadData(db, telp, pass)

				case 2:
					// code here - UPDATE ACCOUNT
					var nama, telepon, password string
					fmt.Println("Masukkan Nama:")
					fmt.Scanln(&nama)
					fmt.Println("Masukkan Telepon:")
					fmt.Scanln(&telepon)
					fmt.Println("Masukkan Password:")
					fmt.Scanln(&password)
					_, err := controllers.UpdateUser(db, nama, telepon, password, dataLogin.Id)
					if err != nil {
						fmt.Println("login gagal")
					}
				case 3:
					// code here - DELETE ACCOUNT
					delUser := entities.User{}
					fmt.Println("Masukkan Telepon:")
					fmt.Scanln(&delUser.Telepon)
					fmt.Println("Masukkan Password:")
					fmt.Scanln(&delUser.Password)
					_, err := controllers.DeleteUser(db, delUser)
					if err != nil {
						fmt.Println("delete gagal")
					}
				case 4:
					// code here - TOP-UP
					var nominal int
					fmt.Println("Masukkan Nominal:")
					fmt.Scanln(&nominal)
					saldo, err := controllers.TopUp(db, dataLogin.Id, nominal)
					if err != nil {
						fmt.Println("topup gagal")
					} else {
						fmt.Println("saldo bertambah sebesar : ", saldo)
					}
				case 5:
					// code here - TRANSFER
					var nominal int
					var othertelp string
					fmt.Println("Masukkan Nomor Telepon Tujuan:")
					fmt.Scanln(&othertelp)
					fmt.Println("Masukkan Nominal:")
					fmt.Scanln(&nominal)
					_, err := controllers.Transfer(db, dataLogin.Id, othertelp, nominal)
					if err != nil {
						fmt.Println("topup gagal")
					}

				case 6:
					// code here - HISTORY TOP-UP
					var idUser int
					fmt.Println("Masukkan id anda : ")
					fmt.Scanln(&idUser)
					controllers.HistoryTopup(db, idUser)

				case 7:
					// code here - HISTORY TRANSFER

				case 8:
					// code here - READ OTHER USERS

				}
			}

		}
	}
	fmt.Println("----- Terimakasih telah bertransaksi -----")
}
