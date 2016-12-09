package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello from Flynn on port %s from container %s\n", port, os.Getenv("HOSTNAME"))
	})
	fmt.Println("app listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
