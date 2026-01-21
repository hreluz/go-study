package main

import (
	"day5_package_app/internal/subscription"
)

func main() {
	// We only see what subscription exposes:
	//  - type Service
	//  - function NewService
	//  - method Activate
	svc := subscription.NewService("pro")
	svc.Activate("user-123")
}
