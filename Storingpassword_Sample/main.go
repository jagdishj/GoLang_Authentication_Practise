package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	//base64 encoding, use only with https, try to avoid using it as decoding is simple.
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:password")))

	//never save passwords store only one-way encryption hash values of the password.
	//hash on the client & server
	pass := "goLang1234"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		fmt.Println("Not Logged in :", err)
	}

	fmt.Println("Logged in")
}

func hashPassword(password string) ([]byte, error) {
	fmt.Println("Password Hashing Started")
	//max cost theoretically takes long time to encrypt and decrypt the password
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: ", err)
	}
	fmt.Println("Hashed Password: ", string(bs))
	return bs, nil
}

func comparePassword(password string, hashedpass []byte) error {
	fmt.Println("compare hashing password started")
	err := bcrypt.CompareHashAndPassword(hashedpass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid Password:", err)
	}
	fmt.Printf("Password Matched")
	return nil
}
