package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "thang"

	hash, _ := HashPassword(password)
	fmt.Println("Data : ", hash)

	math := CheckPassword(password, hash)
	fmt.Println("Check 1 : ", math)

	fmt.Println("--------------")

	hash2, _ := HashPassword(password)
	fmt.Println("Data2 : ", hash2)

	math2 := CheckPassword(hash, hash2)
	fmt.Println("check 2 = ", math2)

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
