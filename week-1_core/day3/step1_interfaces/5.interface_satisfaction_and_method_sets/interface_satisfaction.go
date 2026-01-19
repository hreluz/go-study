package main

import "fmt"

// Interface requiring one method.
type Printer interface {
	Print()
}

type Document struct {
	Text string
}

// Method with VALUE receiver: available to T and *T
func (d Document) Print() {
	fmt.Println("Document:", d.Text)
}

// Method with POINTER receiver: available only to *T
func (d *Document) Clear() {
	d.Text = ""
}

func AcceptPrinter(p Printer) {
	p.Print()
}

func main() {
	// Value
	docValue := Document{Text: "Value receiver example"}

	// Pointer
	docPointer := &Document{Text: "Pointer receiver example"}

	// Both satisfy the Printer interface, because Print() is defined on T.
	AcceptPrinter(docValue)
	AcceptPrinter(docPointer)

	// But only pointers have access to Clear().
	// The following line WOULD NOT COMPILE:
	//
	// docValue.Clear()
	//
	// Because Clear() has a pointer receiver.

	docPointer.Clear()
	fmt.Println("After Clear():", docPointer.Text)
}
