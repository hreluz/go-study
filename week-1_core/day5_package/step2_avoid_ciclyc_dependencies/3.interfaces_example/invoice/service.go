package invoice

import (
	"fmt"

	"day5_package_step_2_avoid_cyclic_dependencies_3_interfaces_example/domain"
)

// PaymentProcessor is the behavior invoice needs.
// Note: it lives in the consumer package (invoice).
type PaymentProcessor interface {
	ProcessPayment(domain.InvoiceID) error
}

type Service struct {
	processor PaymentProcessor
}

func NewService(processor PaymentProcessor) *Service {
	return &Service{processor: processor}
}

func (s *Service) Pay(id domain.InvoiceID) error {
	fmt.Println("InvoiceService: requesting payment for", id)
	return s.processor.ProcessPayment(id)
}
