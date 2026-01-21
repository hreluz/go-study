package main

import "day5_package_step_2_avoid_cyclic_dependencies_1_cyclic_dependency/invoice"

func main() {
	inv := invoice.Invoice{ID: "inv-1"}
	inv.Pay()
}
