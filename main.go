package main

import (
	"log"
	"net/http"
)

var storage Storage

// main start ups the server
// It starts up a http server listening to localhost:8080
// And sets up the handler functions
func main() {
	storage = &Memory{}

	fillStorageWithSampleData(storage)

	http.HandleFunc("/additem", AddItem)
	http.HandleFunc("/", GetAllItems)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

// fillStorageWithSampleData generates sample data
func fillStorageWithSampleData(s Storage) {
	err := s.AddItem(Item{
		Text: "Buy Some Bread",
		Done: false,
	})
	if err != nil {
		panic("Failed to add Item")
	}
	err = s.AddItem(Item{
		Text: "Call Mom",
		Done: false,
	})
	if err != nil {
		panic("Failed to add Item")
	}
}
