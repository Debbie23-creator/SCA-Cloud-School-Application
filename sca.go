package main

import (
	"fmt"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SCA Cloud School Applicaton, this is my first assessment %s", time.Now())
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
