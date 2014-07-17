package server

import (
	"appengine"
	"fmt"
	"net/http"
    "encoding/json"
)

func BoardHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func BoardListHandler(writer http.ResponseWriter, request *http.Request) {
	context := appengine.NewContext(request)

    if request.Method == "GET" {
        boards := GetBoards(context)
        resource, _ := json.Marshal(boards)
        fmt.Fprintf(writer, string(resource))
    } else if request.Method == "POST" {
        board := CreateBoard(request.Data)
        fmt.Fprintf(writer, fmt.Sprintf("Got POST request %#v\n", request))
    } else {
        http.Error(writer, "Bad request to ListHandler", 500)
    }

}

func TopicHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func TopicListHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func PostHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func PostListHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}
