package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	for {
		pwd := getPwd()

		hash := hashAndSalt(pwd)

		pwd2 := getPwd()

		pwdMatch := comparePasswords(hash, pwd2)

		fmt.Println("Password match : ", pwdMatch)
	}

}

func getPwd() []byte {
	fmt.Println("Nhap str")
	var pwd string

	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plaiPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plaiPwd)

	if err != nil {
		log.Println(err)
		return false
	}
	return true

}
