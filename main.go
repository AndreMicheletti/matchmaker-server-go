package main

import (
	"fmt"; "net/http"; "log"
)
	
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

