INFINITE LOOP LEAK
- Actively runnable
- Control flow cycles forever
- CPU may be high or low
- Easy to spot if careless

BLOCKED LEAK
- Control flow frozen
- Waiting on something impossible
- CPU near zero
- Hard to diagnose


Core identification rule
    If you cannot point to who makes the goroutine return, it is a leak candidate.  