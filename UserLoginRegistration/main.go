package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/crypto/bcrypt"
)

var db = map[string][]byte{}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	errMsg := r.FormValue("errormsg")

	fmt.Fprintf(w, `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
	<h1>if there was any error, here it is: %s</h1>
		<form action="/register" method="POST">
			<input type="email" name="email">
			<input type="password" name="password">
			<input type="submit">
		</form>
	</body>
	</html>`, errMsg)
}

func register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		errorMsg := url.QueryEscape("your method was not post")
		http.Redirect(w, r, "/?errormsg="+errorMsg, http.StatusSeeOther)
		return
	}

	e := r.FormValue("email")
	if e == "" {
		errorMsg := url.QueryEscape("email should not be empty")
		http.Redirect(w, r, "?/errormsg="+errorMsg, http.StatusSeeOther)
		return
	}
	p := r.FormValue("password")
	if p == "" {
		errorMsg := url.QueryEscape("password should not be empty")
		http.Redirect(w, r, "?/errormsg="+errorMsg, http.StatusSeeOther)
		return
	}

	bsp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	if err != nil {
		errorMsg := url.QueryEscape("there was an internal server error")
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	log.Println("password :", p)
	log.Println("bcrypted :", bsp)

	db[e] = bsp

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
