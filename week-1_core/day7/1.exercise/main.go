package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	ErrNotFound          = errors.New("account not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidAmount     = errors.New("invalid amount")
)

// --- Over-exported domain types (intentionally) ---

type UserID string

type Account struct {
	ID      UserID
	Balance int
}

// Pointer receivers everywhere (intentionally)
func (a *Account) Deposit(amount int) {
	a.Balance += amount
}
func (a *Account) Withdraw(amount int) bool { // bool return (intentionally weak)
	if a.Balance < amount {
		return false
	}
	a.Balance -= amount
	return true
}

type InMemoryDB struct {
	mu       sync.Mutex
	accounts map[UserID]Account
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{accounts: map[UserID]Account{}}
}

func (db *InMemoryDB) GetAccount(id UserID) (Account, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	a, ok := db.accounts[id]
	return a, ok
}

func (db *InMemoryDB) SaveAccount(a Account) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.accounts[a.ID] = a
}

// --- Global init + hidden side effects (intentionally) ---

var (
	globalDB     *InMemoryDB
	globalLogger *log.Logger
)

func init() {
	globalDB = NewInMemoryDB()
	globalLogger = log.Default()

	// random initial state + magic env var (intentionally)
	seed := time.Now().UnixNano()
	if v := os.Getenv("SEED"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			seed = n
		}
	}
	rand.Seed(seed)

	for i := 1; i <= 3; i++ {
		id := UserID(fmt.Sprintf("u%d", i))
		globalDB.SaveAccount(Account{
			ID:      id,
			Balance: 100 + rand.Intn(200),
		})
	}
}

// --- Service (intentionally mixing responsibilities) ---

type TransferService struct {
	DB     *InMemoryDB
	Log    *log.Logger
	Now    func() time.Time
	Mode   string // random config
	Unused int    // dead weight
}

func NewTransferService() *TransferService {
	return &TransferService{
		DB:   globalDB,
		Log:  globalLogger,
		Now:  time.Now,
		Mode: "prod",
	}
}

type TransferRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

type TransferResponse struct {
	OK        bool   `json:"ok"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// Transfer returns (bool, string) (intentionally weak)
func (s *TransferService) Transfer(ctx context.Context, from, to UserID, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("%w: must be positive", ErrInvalidAmount)
	}

	select {
	case <-time.After(10 * time.Millisecond): // fake latency
	case <-ctx.Done():
		return ctx.Err()
	}

	fromAcc, ok := s.DB.GetAccount(from)
	if !ok {
		return fmt.Errorf("%w: from=%s", ErrNotFound, from)
	}
	toAcc, ok := s.DB.GetAccount(to)
	if !ok {
		return fmt.Errorf("%w: to=%s", ErrNotFound, to)
	}

	if ok := fromAcc.Withdraw(amount); !ok {
		return fmt.Errorf("%w: from=%s amount=%d balance=%d", ErrInsufficientFunds, from, amount, fromAcc.Balance)
	}
	toAcc.Deposit(amount)

	s.DB.SaveAccount(fromAcc)
	s.DB.SaveAccount(toAcc)

	s.Log.Println(fmt.Sprintf("transfer: %s -> %s amount=%d", from, to, amount))
	return nil
}

func (s *TransferService) HandleTransfer(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid JSON payload",
			http.StatusBadRequest,
		)
		return
	}

	err := s.Transfer(r.Context(), UserID(req.From), UserID(req.To), req.Amount)

	status := http.StatusOK
	msg := "ok"

	switch {
	case err == nil:
		// keep OK
	case errors.Is(err, ErrInvalidAmount):
		status = http.StatusBadRequest
		msg = "invalid amount"
	case errors.Is(err, ErrNotFound):
		status = http.StatusNotFound
		msg = "account not found"
	case errors.Is(err, ErrInsufficientFunds):
		status = http.StatusConflict
		msg = "insufficient funds"
	case errors.Is(err, context.Canceled):
		status = 499 // common non-standard client-cancel code used by nginx
		msg = "request cancelled"
	case errors.Is(err, context.DeadlineExceeded):
		status = http.StatusGatewayTimeout
		msg = "deadline exceeded"
	default:
		status = http.StatusInternalServerError
		msg = "internal error"
		// log internal detail, do not leak it to clients
		s.Log.Println(err.Error())
	}

	resp := TransferResponse{
		OK:        err == nil,
		Message:   msg,
		Timestamp: s.Now().Format(time.RFC3339Nano),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *TransferService) HandleAccounts(w http.ResponseWriter, r *http.Request) {
	type view struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	}
	var out []view
	for _, id := range []UserID{"u1", "u2", "u3"} {
		if a, ok := s.DB.GetAccount(id); ok {
			out = append(out, view{ID: string(a.ID), Balance: a.Balance})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(out)
}

func main() {
	svc := NewTransferService()

	mux := http.NewServeMux()
	mux.HandleFunc("/transfer", svc.HandleTransfer)
	mux.HandleFunc("/accounts", svc.HandleAccounts)

	// no graceful shutdown, no context wiring (intentionally)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
