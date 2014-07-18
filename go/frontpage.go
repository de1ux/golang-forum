package server

import (
    "io/ioutil"
    "net/http"
    "fmt"
)

func FrontPageHandler(writer http.ResponseWriter, request *http.Request) {
    contents, _ := ioutil.ReadFile("frontpage.html")
    fmt.Fprintf(writer, string(contents))
}