package invoice

import (
	"day5_package_step_2_avoid_cyclic_dependencies_2_refactor/domain"
	"fmt"
)

type Invoice struct {
	ID domain.InvoiceID
}

func (i *Invoice) Print() {
	fmt.Println("Invoice:", i.ID)
}
