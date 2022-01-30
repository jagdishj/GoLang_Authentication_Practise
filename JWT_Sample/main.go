package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID string
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token as expired")
	}
	if u.SessionID == "" {
		return fmt.Errorf("Invalid session id")
	}
	return nil
}

func main() {
	//base64 encoding, use only with https, try to avoid using it as decoding is simple.
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:password")))

	//never save passwords store only one-way encryption hash values of the password.
	//hash on the client & server

	//setting key value

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

	fmt.Println("Generating JWT key")
	generateNewKey()

	claims := &jwt.StandardClaims{
		ExpiresAt: 150000,
		Issuer:    "test",
	}

	var uc = &UserClaims{
		StandardClaims: *claims,
		SessionID:      "SESSION ID",
	}

	tok, err := createToken(uc)
	if err != nil {
		fmt.Println("error in creating token ", err)
	}
	fmt.Println("Token : ", tok)

	token, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

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
	h := hmac.New(sha512.New, dbkeys[currentKey].key)

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

func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(dbkeys[currentKey].key)
	if err != nil {
		return "", fmt.Errorf("Error in createtoken function - signing token: ", err)
	}
	return signedToken, nil
}

type mykey_struct struct {
	key     []byte
	created time.Time
}

var currentKey = ""
var dbkeys = map[string]mykey_struct{} // you can write a logic to get the keys from database.

func generateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return fmt.Errorf("Error in generating a new key ", err)
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("Error in generating uuid ", err)
	}

	dbkeys[uid.String()] = mykey_struct{
		key:     newKey,
		created: time.Now(),
	}
	currentKey = uid.String()
	fmt.Println("CurrentKey ", currentKey)
	return nil
}

func parseToken(signedtoken string) (*UserClaims, error) {

	t, err := jwt.ParseWithClaims(signedtoken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodES512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}

		presentkey, ok := t.Header["PresentKey"].(string)
		if !ok {
			return nil, fmt.Errorf("Invalid Key Id")
		}

		k, ok := dbkeys[presentkey]
		if !ok {
			return nil, fmt.Errorf("Invalid Key.Id")
		}

		return k, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error in parseToken :", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return t.Claims.(*UserClaims), nil
}
