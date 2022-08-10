package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port, exist := os.LookupEnv("PORT")

	if !exist {
		port = "8000"
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Println("Listing for" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
