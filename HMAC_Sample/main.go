package main

import (
	"encoding/hex"
	"fmt"
	"io"

	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"

	"golang.org/x/crypto/bcrypt"
)

var key = []byte{}

func main() {
	//base64 encoding, use only with https, try to avoid using it as decoding is simple.
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:password")))

	//never save passwords store only one-way encryption hash values of the password.
	//hash on the client & server

	//setting key value
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	ExploringHash()
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

	enc, enc_err := signmessage([]byte("Hello"))
	if enc_err != nil {
		fmt.Println("Error in Signing a Message : ", enc_err)
	}

	fmt.Println("Checking sign")
	flag, dec_err := checkSign([]byte("Hello"), enc)
	if dec_err != nil {
		fmt.Println("Error in Checking Sign : ", dec_err)
	}

	fmt.Println("Check Sign Result:", flag)
}

func ExploringHash() {

	fmt.Println("SHA256 hash : ", sha256.New())

	fmt.Println("SHA512 hash : ", sha512.New())

	fmt.Println("SHA512 Sum(nil) :", sha512.New().Sum(nil))

	fmt.Println("Length of SHA512 Sum(nil) : ", len(sha512.New().Sum(nil))*8)

	h := sha256.New()
	io.WriteString(h, "hello")
	s := h.Sum(nil)
	fmt.Println("sha256 encoding hello : ", hex.EncodeToString(s))

	h = sha512.New()
	io.WriteString(h, "hello")
	s = h.Sum(nil)
	fmt.Println("sha512 encoding hello : ", hex.EncodeToString(s))
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

func signmessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)

	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in sign Message while hashing message:", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSign(msg, sign []byte) (bool, error) {
	newSig, err := signmessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSign while getting signature of message :", err)
	}

	same := hmac.Equal(newSig, sign)
	return same, nil
}
