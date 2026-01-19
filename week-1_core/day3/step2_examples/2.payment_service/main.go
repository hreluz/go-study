package main

import "fmt"

// Payment represents a processed payment.
type Payment struct {
	ID     int
	Amount int
}

// The CONSUMER (ProcessPayment) defines what it needs:
// a store that can SavePayment(Payment) error.
type PaymentStore interface {
	SavePayment(p Payment) error
}

// Concrete implementation 1: database-backed store.
type DatabaseStore struct {
	Dsn string
}

func (d DatabaseStore) SavePayment(p Payment) error {
	fmt.Printf("[DatabaseStore] DSN=%s, saving payment: %+v\n", d.Dsn, p)
	return nil
}

// Concrete implementation 2: in-memory store (useful for tests or dev).
type MemoryStore struct {
	Items []Payment
}

func (m *MemoryStore) SavePayment(p Payment) error {
	m.Items = append(m.Items, p)
	fmt.Printf("[MemoryStore] len=%d, last=%+v\n", len(m.Items), p)
	return nil
}

// ProcessPayment is the CONSUMER.
// It does not care where the payment is stored, only that it can be stored.
func ProcessPayment(store PaymentStore, amount int) error {
	p := Payment{
		ID:     1, // imagine an auto-increment or generated ID
		Amount: amount,
	}

	fmt.Println("Processing payment...")
	if err := store.SavePayment(p); err != nil {
		return fmt.Errorf("failed to save payment: %w", err)
	}

	fmt.Println("Payment processed successfully")
	return nil
}

func main() {
	dbStore := DatabaseStore{Dsn: "postgres://user:pass@localhost/db"}
	memStore := &MemoryStore{}

	fmt.Println("=== Using DatabaseStore ===")
	if err := ProcessPayment(dbStore, 100); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("\n=== Using MemoryStore ===")
	if err := ProcessPayment(memStore, 200); err != nil {
		fmt.Println("error:", err)
	}
}
