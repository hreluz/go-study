package main

import "fmt"

// The CONSUMER (SendWelcome) defines what it needs:
// "I need something that can Notify(userID, message)".
type Notifier interface {
	Notify(userID int, message string) error
}

// Concrete implementation 1: email notifications.
type EmailNotifier struct{}

func (EmailNotifier) Notify(userID int, message string) error {
	fmt.Printf("[EmailNotifier] To user %d: %s\n", userID, message)
	return nil
}

// Concrete implementation 2: SMS notifications.
type SMSNotifier struct{}

func (SMSNotifier) Notify(userID int, message string) error {
	fmt.Printf("[SMSNotifier] To user %d: %s\n", userID, message)
	return nil
}

// This function is the CONSUMER.
// It depends only on the behavior described by Notifier.
func SendWelcome(n Notifier, userID int) error {
	return n.Notify(userID, "Welcome to the system!")
}

func main() {
	email := EmailNotifier{}
	sms := SMSNotifier{}

	fmt.Println("Using EmailNotifier:")
	if err := SendWelcome(email, 42); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("\nUsing SMSNotifier:")
	if err := SendWelcome(sms, 99); err != nil {
		fmt.Println("error:", err)
	}
}

/**
	Key point:
		Notifier is defined next to SendWelcome, where the need appears.
		Email and SMS implementations do not define the interface; they just satisfy it.
**/
