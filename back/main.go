package main

import (
	"back/config"
	"back/routers"
	// "fmt"
	// "golang.org/x/crypto/bcrypt"
)

func main() {
	config.ConnectDB()

	router := routers.SetUpRouter()

	router.Run(":8080")

	// password := "123456"
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	fmt.Println("Error hashing password:", err)
	// 	return
	// }
	// fmt.Println("Hashed Password:", string(hashedPassword))

	// err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	// if err != nil {
	// 	fmt.Println("Failed to verify password:", err)
	// } else {
	// 	fmt.Println("Password verification successful!")
	// }

}
