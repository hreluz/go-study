package payment

import "day5_package_step_2_avoid_cyclic_dependencies_1_cyclic_dependency/invoice"

func Process(invoiceID string) {
	_ = invoice.Invoice{ID: invoiceID} // artificial reference to create cycle
}
