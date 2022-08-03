package main

// Storage defines the interface used to maintain the model data
type Storage interface {
	GetItems() ([]Item, error)
	AddItem(Item) error
}

// Memory is a storage implementation that works solely in memory.
type Memory struct {
	items []Item
}

// GetItems retuns a slice of all items
func (m Memory) GetItems() ([]Item, error) {
	return m.items, nil
}

// AddItem appends a new item to the itemlist
func (m *Memory) AddItem(i Item) error {
	m.items = append(m.items, i)
	return nil
}
