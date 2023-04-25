package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	h := Handler{}
	router.Handle("/handler", h)
	router.HandleFunc("/handlerfunc", GetHandlerFuncResponse())
	http.ListenAndServe(":8080", router)
}

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context())
	fmt.Println(r.Body)
	fmt.Println(r.URL.Query())
	pathParamsMap := mux.Vars(r)
	fmt.Println(pathParamsMap)
	w.Write([]byte("Hello World"))
}

func GetHandlerFuncResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Context())
		fmt.Println(r.Body)
		fmt.Println(r.URL.Query())
		pathParamsMap := mux.Vars(r)
		fmt.Println(pathParamsMap)
		w.Write([]byte("Hello World"))
	}
}
