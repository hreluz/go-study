package payment

import (
	"day5_package_step_2_avoid_cyclic_dependencies_2_refactor/domain"
	"fmt"
)

func Process(id domain.InvoiceID) {
	fmt.Println("Processing payment for:", id)
}
