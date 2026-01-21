package main

import (
	"day5_package_step_2_avoid_cyclic_dependencies_4_runnable_example/internal/subscription"
)

func main() {
	svc := subscription.NewService("pro")
	svc.Activate("user-123")
	svc.Activate("") // invalid user
}
