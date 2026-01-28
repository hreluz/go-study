# Design Note — Transfer Service Refactor (Day 7)

## Packages
- main: wiring only (HTTP routes + service construction). No domain rules should live here.

## Domain model (values vs pointers)
- Account is a value type.
  - Rationale: prevents aliasing and shared mutable state.
  - Each request works on a copy and persists changes explicitly via SaveAccount.
  - This avoids data races where multiple requests mutate the same *Account outside locks.

## Context
- context.Context is owned by the caller and passed explicitly into Transfer(ctx, ...).
- Context is not stored in structs (service or domain) because it is request-scoped.

## Errors
- Transfer returns error (not bool/string).
  - Rationale: callers can classify errors with errors.Is, and we can wrap with context.
- Sentinel errors exist only for branching:
  - ErrInvalidAmount, ErrNotFound, ErrInsufficientFunds.
- Cancellation and deadlines return ctx.Err() for standard classification.

## Interfaces
- Removed Clock, Logger, DB interfaces because there was only one implementation.
- Testability is preserved by injection of:
  - Now func() time.Time (instead of Clock)
  - *log.Logger (instead of Logger)
- Interfaces will be introduced only when there are real alternate implementations
  (e.g., multiple DB backends) with an actual consumer need.

## HTTP boundary
- No panic in handlers.
- Handlers translate domain errors into HTTP status codes and safe messages.
- Internal error details are logged; clients get stable semantics.

