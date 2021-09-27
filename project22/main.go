package main

import (
	"net/http"
	"time"
	"log"
	"fmt"
)

func root(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1,0,0)
	cookie := http.Cookie{
		Name:		"username",
		Value:		"hoge",
		Expires: 	expiration,
	}
	http.SetCookie(w, &cookie)
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	fmt.Fprint(w, cookie)
}

func readCookies(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, cookie.Name)
	}
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/readcookie", readCookie)
	http.HandleFunc("/readcookies", readCookies)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}	
}