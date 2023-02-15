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

				case 2:
					// code here - UPDATE ACCOUNT

				case 3:
					// code here - DELETE ACCOUNT

				case 4:
					// code here - TOP-UP

				case 5:
					// code here - TRANSFER

				case 6:
					// code here - HISTORY TOP-UP

				case 7:
					// code here - HISTORY TRANSFER

				case 8:
					// code here - READ OTHER USERS

				}
			}

		}
	}

}

// 		case 3:
// 			// code here - READ ACCOUNT

// 		case 4:
// 			// code here - UPDATE ACCOUNT

// 		case 5:
// 			// code here - DELETE ACCOUNT

// 		case 6:
// 			// code here - TOP-UP

// 		case 7:
// 			// code here - TRANSFER

// 		case 8:
// 			// code here - HISTORY TOP-UP

// 		case 9:
// 			// code here - HISTORY TRANSFER

// 		case 10:
// 			// code here - OTHER USER

// 		case 0:
// 			// code here - EXIT
// 		}
// 	}
// }
