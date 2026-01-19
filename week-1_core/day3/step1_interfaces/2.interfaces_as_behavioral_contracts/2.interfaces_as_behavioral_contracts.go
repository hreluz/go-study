package main

import (
	"fmt"
)

// Saver is a behavioral contract.
// It says: "I need something that can Save(string) error".
type Saver interface {
	Save(data string) error
}

// FileStorage is a concrete type (imagine it writes to disk).
type FileStorage struct {
	Path string
}

func (f FileStorage) Save(data string) error {
	// In a real program you would write to a file.
	// Here we just simulate the behavior with a print.
	fmt.Printf("[FileStorage] Saving to %s: %q\n", f.Path, data)
	return nil
}

// MemoryStorage is another concrete type (imagine it stores in memory).
type MemoryStorage struct {
	Items []string
}

func (m *MemoryStorage) Save(data string) error {
	m.Items = append(m.Items, data)
	fmt.Printf("[MemoryStorage] Now I have %d items. Last: %q\n", len(m.Items), data)
	return nil
}

// ProcessAndSave depends only on the Saver behavior.
func ProcessAndSave(s Saver, raw string) error {
	processed := "[processed] " + raw
	return s.Save(processed)
}

func main() {
	fileStore := FileStorage{Path: "/tmp/data.txt"}
	memStore := &MemoryStorage{}

	fmt.Println("--- Using FileStorage ---")
	if err := ProcessAndSave(fileStore, "hello from file"); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("--- Using MemoryStorage ---")
	if err := ProcessAndSave(memStore, "hello from memory"); err != nil {
		fmt.Println("error:", err)
	}
}

/**
	Show that an interface only describes behavior (methods).

	Show that different concrete types can satisfy the same interface.

	Show a function that depends only on the interface, not on concrete types.


	Saver is a description of required behavior.

	FileStorage and *MemoryStorage are concrete implementations.

	ProcessAndSave knows nothing about how saving is done, only that it can be done.
**/
