package main

import "fmt"

// Item stores the data of indivual items
type Item struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// String prints the item as human readable string
func (i Item) String() string {
	return fmt.Sprintf("text: %s, done: %t", i.Text, i.Done)
}
