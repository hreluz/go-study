package subscription

import (
	"fmt"

	"day5_package_step_2_avoid_cyclic_dependencies_4_runnable_example/internal/billing"
)

type Service struct {
	plan    string
	billing billing.Client
}

// NewService is the public constructor exposed to other packages
// inside this module.
func NewService(plan string) *Service {
	return &Service{
		plan:    plan,
		billing: billing.NewClient(),
	}
}

func (s *Service) Activate(userID string) {
	if !isValidUser(userID) {
		fmt.Println("subscription: invalid user id")
		return
	}

	fmt.Printf("subscription: activating user %s on plan %s\n", userID, s.plan)
	s.billing.Charge(userID, 10.0)
}
