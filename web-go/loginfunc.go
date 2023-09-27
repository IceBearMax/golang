package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		timestamp := strconv.Itoa(time.Now().Nanosecond())
		hashWr := md5.New()
		hashWr.Write([]byte(timestamp))
		token := fmt.Sprintf("%x", hashWr.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {

		} else {

		}
		fmt.Println("username length", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))

	}

}
