package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func main() {
	router := mux.NewRouter()
	h := Handler{}
	router.Handle("/", h)
	http.ListenAndServe(":8080", router)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context())
	fmt.Println(r.Body)
	fmt.Println(r.URL.Query())
	pathParamsMap := mux.Vars(r)
	fmt.Println(pathParamsMap)
	w.Write([]byte("Hello World"))
}
