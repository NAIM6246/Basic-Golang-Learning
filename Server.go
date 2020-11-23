package main

import (
	"fmt"
	"net/http"
)

func handler(a http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/create":
		fmt.Fprintf(a, "Your request is accepted.You will be notified soon")
	case "/post":
		fmt.Fprintf(a, "Your message is posted")
	case "/update":
		fmt.Fprintf(a, "Data is updated")
	case "/delete":
		fmt.Fprintf(a, "Your account has been deleted")
	default:
		fmt.Fprintf(a, "Page not found!")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Satrating the server")
}
