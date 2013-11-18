package main 

import (
		"fmt"
		"net/http"
		"html/template"
		"strings"
		"log")

func sayhelloName(rep http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form{
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, " "))
	}
	fmt.Fprintf(rep, "Hello astaxie!")
}

func login(rep http.ResponseWriter, r * http.Request) {
	fmt.Println("method: ", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.gtpl")
		t.Execute(rep, nil)
	} else {
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}