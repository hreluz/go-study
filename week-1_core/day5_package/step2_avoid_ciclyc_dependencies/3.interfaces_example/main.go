package main

import (
	"day5_package_step_2_avoid_cyclic_dependencies_3_interfaces_example/domain"
	"day5_package_step_2_avoid_cyclic_dependencies_3_interfaces_example/invoice"
	"day5_package_step_2_avoid_cyclic_dependencies_3_interfaces_example/payment"
)

func main() {
	id := domain.InvoiceID("inv-1")

	processor := payment.NewProcessor()

	// This assignment ensures at compile time that *Processor
	// satisfies invoice.PaymentProcessor.
	var p invoice.PaymentProcessor = processor

	svc := invoice.NewService(p)
	_ = svc.Pay(id)
}
