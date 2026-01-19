package main

import "fmt"

// No interface yet.

// First implementation
type DatabaseStore struct{}

func (DatabaseStore) SaveUser(name string) error {
	fmt.Println("DB:", name)
	return nil
}

// Second implementation appears later
type MemoryStore struct {
	users []string
}

func (m *MemoryStore) SaveUser(name string) error {
	m.users = append(m.users, name)
	fmt.Println("MEM:", name)
	return nil
}

// Now a function needs to work with *either* implementation.
// ✔ NOW we introduce the interface.
type UserSaver interface {
	SaveUser(name string) error
}

func ProcessUser(s UserSaver, name string) {
	fmt.Println("Processing:", name)
	s.SaveUser(name)
}

func main() {
	db := DatabaseStore{}
	mem := &MemoryStore{}

	ProcessUser(db, "Héctor with DB")
	ProcessUser(mem, "Héctor with MEM")
}
