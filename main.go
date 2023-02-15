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

	fmt.Println("Pilih Menu:\n1. Register\n2. Login\n3. Read Account\n4. Update Account\n5. Delete Account\n6. Top-Up\n7. Transfer\n8. History Tup-Up\n9. History Transfer\n10. Other User\n0. Exit")
	fmt.Println("Input pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		// code here - REGISTER

	case 2:
		// code here - LOGIN
		logUser := entities.User{}

		fmt.Println("Masukkan Telepon:")
		fmt.Scanln(&logUser.Telepon)
		fmt.Println("Masukkan Password:")
		fmt.Scanln(&logUser.Password)
		dataLogin, err := controllers.Login(db, logUser)
		if err != nil {
			fmt.Println("login gagal")
		} else {
			fmt.Println("login berhasil")
		}
		fmt.Println(dataLogin)
	case 3:
		// code here - READ ACCOUNT

	case 4:
		// code here - UPDATE ACCOUNT

	case 5:
		// code here - DELETE ACCOUNT

	case 6:
		// code here - TOP-UP

	case 7:
		// code here - TRANSFER

	case 8:
		// code here - HISTORY TOP-UP

	case 9:
		// code here - HISTORY TRANSFER

	case 10:
		// code here - OTHER USER

	case 0:
		// code here - EXIT
	}
}
