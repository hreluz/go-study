package invoice

import "day5_package_step_2_avoid_cyclic_dependencies_1_cyclic_dependency/payment"

type Invoice struct {
	ID string
}

func (i *Invoice) Pay() {
	payment.Process(i.ID)
}
