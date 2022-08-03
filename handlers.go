package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// AddItem is the handler function to add new items
func AddItem(w http.ResponseWriter, r *http.Request) {
	log.Println("Add Items called")

	decoder := json.NewDecoder(r.Body)
	var newItem Item
	err := decoder.Decode(&newItem)

	if err != nil {
		log.Printf("AddItem: Failed to decode new item. Body: %s. Error %v", r.Body, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = storage.AddItem(newItem)
	if err != nil {
		log.Printf("AddItem: Failed to add new item. Error %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode("success")
	log.Println("Added new item successfully")
}

// AddGetAllItems is the handler function to return all items
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllItems called")
	items, err := storage.GetItems()
	if err != nil {
		log.Fatal("GetAllItems: Failed to load items:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
	log.Println("Returned all items successfully")
}
