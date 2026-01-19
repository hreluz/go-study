package main

import "fmt"

// ❌ BAD: Created without a real consumer, 1:1 with the concrete type
type UserStore interface {
	SaveUser(name string) error
}

type DatabaseStore struct{}

func (DatabaseStore) SaveUser(name string) error {
	fmt.Println("Saving user to database:", name)
	return nil
}

// ❌ BAD: Constructor returns interface
func NewUserStore() UserStore {
	return DatabaseStore{}
}

func main() {
	store := NewUserStore()
	store.SaveUser("Héctor")
}

/**

Why this is harmful:

    UserStore has only one implementation.

    No other code depends on this abstraction.

    If DatabaseStore later has new useful methods, callers cannot access them.

    The interface adds no value and hides behavior.
**/
