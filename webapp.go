// Author: Anthony Hein
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/info", info)
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func info(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `This is a Golang server that simply servers beautiful HTML files capturing my class notes that were exported from Typora, a markdown editor.`)
}
