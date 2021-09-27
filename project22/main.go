package main

import (
	"net/http"
	"time"
	"log"
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

func main() {
	http.HandleFunc("/", root)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}	
}