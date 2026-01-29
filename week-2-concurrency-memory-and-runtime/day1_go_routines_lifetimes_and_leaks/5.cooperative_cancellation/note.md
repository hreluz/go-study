Cancellation is a signal, not an action

Important distinctions:

    ❌ Cancellation does not stop execution

    ❌ Cancellation does not interrupt code

    ❌ Cancellation does not force a return

    ✅ Cancellation is observed

    ✅ The goroutine decides what to do

    ✅ Correct response is: clean up → return