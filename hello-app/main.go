package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
)

var err error

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api is called to add")
	uuid := uuid.New()
	fmt.Fprintf(w, "http2 server api is called", uuid)
	ioutil.WriteFile("/app/hello.txt", []byte("hello added data"), 0644)
}
func ReadData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api is called to read")
	b, err := ioutil.ReadFile("/app/hello.txt")
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(string(b))
}

func main() {
	fmt.Println("server started on 9091")
	http.HandleFunc("/add", rootHandler)
	http.HandleFunc("/read", ReadData)

	log.Fatal(http.ListenAndServe(":9091", nil))
}
