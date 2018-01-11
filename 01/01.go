package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	ID      int    `json:"id, string"`
}
type helloRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	http.HandleFunc("/hello", helloHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	logMessage := fmt.Sprintf("the req of type %s went to %s path", r.Method, r.URL)
	log.Printf("Logger says:\n%s", logMessage)
	/*
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	*/
	var reqObj helloRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqObj)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := helloResponse{
		Message: "Thank you for the initiation :). Hello " + reqObj.Name,
		Author:  "Salar Hafezi",
		Date:    "2018-01-11",
		ID:      1371,
	}
	/*
		data, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Fprint(w, "Hello World\n")
		}
		fmt.Fprintf(w, string(data))
	*/
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
