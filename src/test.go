package main

import (
	"log"
	"net/http"
	"reflect"
)

func main() {
	http.HandleFunc("/", getvaluehandler)
	http.ListenAndServe("127.0.0.1:8800", nil)
}
func getvaluehandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	value := query.Get("test_key")
	log.Println(reflect.TypeOf(value))
	log.Println(value)
}
