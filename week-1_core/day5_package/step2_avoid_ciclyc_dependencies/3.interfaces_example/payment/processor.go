package payment

import (
	"fmt"

	"day5_package_step_2_avoid_cyclic_dependencies_3_interfaces_example/domain"
)

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

// This method matches invoice.PaymentProcessor's requirement.
// We do not import invoice here: no cycle.
func (p *Processor) ProcessPayment(id domain.InvoiceID) error {
	fmt.Println("PaymentProcessor: processing payment for", id)
	return nil
}
