# Go Mastery Plan — 30-Day Senior-Level Roadmap

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)
![Status](https://img.shields.io/badge/Status-Active-brightgreen)
![License](https://img.shields.io/badge/License-MIT-blue)

A structured and intensive 30-day program designed to elevate a Go developer from intermediate to senior level. The curriculum emphasizes idiomatic Go, concurrency correctness, runtime internals, performance engineering, and scalable architecture patterns.

## 📚 Table of Contents
1. Overall Goal
2. Week 1 — Core Language Mastery
3. Week 2 — Concurrency, Memory, and the Runtime
4. Week 3 — Performance, Tooling, and Production
5. Week 4 — Architecture, Testing, and Expert Patterns
6. Capstone Project
7. Repository Structure
8. How to Use This Repository
9. License

## Overall Goal

By completing this 30-day mastery plan, you will be able to:
- Write idiomatic, minimal, and maintainable Go.
- Reason about concurrency safety, memory behavior, and runtime performance.
- Understand the Go scheduler (M:P:G), GC phases, and memory model.
- Design large, modular Go services with proper boundaries.
- Benchmark, profile, and optimize real systems.
- Pass senior-level Go interviews and lead Go codebases.

# Week 1 — Core Language Mastery
**Focus:** semantics, composition, interfaces, errors, idioms.

### Topics:
- Zero values; value vs pointer semantics
- Struct embedding vs inheritance
- Accept interfaces, return structs
- Method sets & interface satisfaction
- Error handling best practices (errors.Is, errors.As, wrapping)
- Package design & visibility rules
- Avoid abusing init()
- Context propagation everywhere

### Required Reading:
- Effective Go
- Go Code Review Comments
- errors and context source code

### Exercises:
- Refactor an OO-heavy codebase into idiomatic Go
- Replace exception-like logic with explicit error returns
- Remove unnecessary interfaces & observe clarity improvement

# Week 2 — Concurrency, Memory, and the Runtime
**Focus:** correctness over cleverness.

### Topics:
- Goroutine lifecycle & leaks
- Channels and when not to use them
- Worker pools; fan-in/fan-out
- sync.Mutex, RWMutex, Once, Map
- Go memory model
- Scheduler internals (M:P:G)
- GC phases & tuning
- Escape analysis

### Required Reading:
- The Go Memory Model
- Concurrency in Go
- runtime/proc.go and runtime/mgc.go

### Exercises:
- Build a leak-free worker pool
- Introduce races intentionally; fix them
- Benchmark mutex vs channel solutions
- Tune GC for a memory-heavy service

# Week 3 — Performance, Tooling, and Production
**Focus:** operating Go in real environments.

### Topics:
- pprof profiling
- go tool trace
- Benchmarking with testing.B
- Allocation minimization
- sync.Pool usage patterns
- HTTP internals
- Observability and logging
- Build tags, modules, dependency hygiene

### Required Reading:
- Profiling Go Programs
- net/http and context source
- Uber Go Style Guide

### Exercises:
- Reduce allocations by 50% in a hot path
- Build a high-throughput HTTP API with graceful shutdown
- Diagnose a CPU bottleneck with pprof

# Week 4 — Architecture, Testing, and Expert Patterns
**Focus:** Go architecture at scale.

### Topics:
- Package-oriented architecture
- Domain modeling
- Table-driven tests and fuzzing
- Integration vs unit testing
- Dependency injection without frameworks
- Config management & feature flags
- API versioning & backward compatibility

### Required Reading:
- Standard Go Project Layout
- testing and httptest source
- Practical Go — Dave Cheney

# Capstone Project

Build a production-grade Go service with:
- Clean package boundaries
- Context-aware handlers
- Concurrency-safe internals
- Profiling and benchmarking artifacts
- Graceful shutdown
- Structured logging
- A full test suite

# Repository Structure

Suggested layout:

.
├── week01_core/
│   ├── exercises/
│   └── notes.md
├── week02_concurrency/
│   ├── workerpool/
│   ├── races/
│   └── notes.md
├── week03_performance/
│   ├── profiling/
│   ├── http_api/
│   └── notes.md
├── week04_architecture/
│   ├── testing/
│   ├── domain/
│   └── notes.md
├── capstone/
├── README.md
└── LICENSE

# How to Use This Repository

1. Follow the plan weekly.
2. Keep exercises in each week's folder.
3. Document findings and benchmarks.
4. Build the final project in /capstone.
5. Share progress publicly.

# License

MIT License (recommended).
