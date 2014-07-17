package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/board/{resourceId}", BoardHandler)
	r.HandleFunc("/boards/", BoardListHandler)

	r.HandleFunc("/topic/{resourceId}", TopicHandler)
	r.HandleFunc("/topics/", TopicListHandler)

	r.HandleFunc("/post/{resourceId}", PostHandler)
	r.HandleFunc("/posts/", PostListHandler)

	http.Handle("/", r)
}
