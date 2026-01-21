package subscription

import "fmt"

// Public API (visible to other packages)
type Service struct {
	plan string
}

// NewService is an exported constructor.
func NewService(plan string) *Service {
	return &Service{plan: plan}
}

// Activate is an exported behavior: part of the package API.
func (s *Service) Activate(userID string) {
	// Internal logic hidden behind the package boundary.
	fmt.Printf("Activating %s on plan %s\n", userID, s.plan)
}
