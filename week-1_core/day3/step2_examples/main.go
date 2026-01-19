package main

import "fmt"

// The CONSUMER (PrintDailyReport) defines what it needs:
// a generator that can Generate() string.
type ReportGenerator interface {
	Generate() string
}

// Concrete implementation 1: a real report generator.
type DailyReport struct {
	Day string
}

func (d DailyReport) Generate() string {
	return fmt.Sprintf("Daily report for %s: all systems operational.", d.Day)
}

// Concrete implementation 2: a fake implementation (good for tests).
type FakeReport struct {
	Content string
}

func (f FakeReport) Generate() string {
	return f.Content
}

// PrintDailyReport is the CONSUMER.
// It only depends on the behavior described by ReportGenerator.
func PrintDailyReport(g ReportGenerator) {
	fmt.Println("=== REPORT BEGIN ===")
	fmt.Println(g.Generate())
	fmt.Println("=== REPORT END ===")
}

func main() {
	real := DailyReport{Day: "2026-01-19"}
	fake := FakeReport{Content: "FAKE REPORT for testing output formatting."}

	fmt.Println("Using real DailyReport:")
	PrintDailyReport(real)

	fmt.Println("\nUsing FakeReport (e.g., in tests):")
	PrintDailyReport(fake)
}
