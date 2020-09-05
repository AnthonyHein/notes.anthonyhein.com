// Author: Anthony Hein
package main

import (
    "io"
    "net/http"
)

func main() {
    http.HandleFunc("/info", info)
    http.Handle("/", http.FileServer(http.Dir("./notes")))
    http.ListenAndServe(":443", nil)
}

func info(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, `This is a Golang server that simply servers beautiful HTML files capturing my class notes that were exported from Typora, a markdown editor.`)
}
