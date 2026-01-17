package main

import "fmt"

type Identity struct {
	Name string
}

type Species struct {
	Kind string
}

type Dog struct {
	Identity
	Species
	AnimalDescriber
}

type Cat struct {
	Identity
	Species
	AnimalDescriber
}

type AnimalDescriber struct{}

func (AnimalDescriber) Describe(id Identity, sp Species) string {
	return "Animal: " + id.Name + ", Kind: " + sp.Kind
}

type Logger struct{}

func (Logger) Info(msg string) {
	fmt.Println("[INFO]", msg)
}

type AuditLogger struct{}

func (AuditLogger) Info(msg string) {
	fmt.Println("[AUDIT]", msg)
}

type PaymentService struct {
	AppLogger   Logger
	AuditLogger AuditLogger
}

func main() {
	dog := Dog{
		Identity{Name: "Fido"},
		Species{Kind: "Mammal"},
		AnimalDescriber{},
	}

	cat := Cat{
		Identity:        Identity{Name: "Misu"},
		Species:         Species{Kind: "Mammal"},
		AnimalDescriber: AnimalDescriber{},
	}

	fmt.Println(dog.Describe(dog.Identity, dog.Species))
	fmt.Println(cat.Describe(dog.Identity, dog.Species))

	ps := PaymentService{}
	ps.AppLogger.Info("payment created")
	_ = PaymentService{}
}
