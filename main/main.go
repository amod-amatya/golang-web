package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Welcome to Go Web!")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(response http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Fprintf(response, "Welcome to Go web %q %q", r.URL.Path, resp)
		fmt.Println("Welcome to Go cli")
	}
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Amod"
	fmt.Fprintf(response, "About me....")
	fmt.Println("The Endpoint: ", who)
}
