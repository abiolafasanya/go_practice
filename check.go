package main

import (
	"fmt"	
	"crypto/sha256"
	"crypto/md5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func verifyPassword(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}


func main() {

	pass := "bee"

	hashPass := sha256.Sum256([]byte(pass))
	hashMd := md5.Sum([]byte(pass))

	password := "abiola"
	hash, _ := HashPassword(password)

	check := verifyPassword(password, hash)

	fmt.Println(password)
	fmt.Println("Password:  ", hash)
	fmt.Println("Verify Password:  ", check)

	fmt.Printf("Pass: %v\n",hashPass)
	fmt.Printf("Pass: %v\n",hashMd)

	fmt.Println("Hello learning")
}