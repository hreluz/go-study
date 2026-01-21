package main

import (
	"day5_package_step_2_avoid_cyclic_dependencies_2_refactor/domain"
	"day5_package_step_2_avoid_cyclic_dependencies_2_refactor/invoice"
	"day5_package_step_2_avoid_cyclic_dependencies_2_refactor/payment"
)

func main() {
	id := domain.InvoiceID("inv-1")

	inv := invoice.Invoice{ID: id}
	inv.Print()

	payment.Process(id)
}
