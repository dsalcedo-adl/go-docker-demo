package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type test_struct struct {
	Test string
}

func main() {
	// Set the flags for the logging package to give us the filename in the logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting server...")
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
//		_, _ = fmt.Fprintln(w, `Hello, visitor!`)
		decoder := json.NewDecoder(r.Body)
		var t test_struct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		var jsonData []byte
		jsonData, err = json.Marshal(t)
		if err != nil {
			log.Println(err)
		}
		//log.Println(t)
		//_, _ = fmt.Fprintln(w, t)
		log.Println(string(jsonData))
		_, _ = fmt.Fprintln(w, string(jsonData))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
