package main

func blockedReceive() {
	ch := make(chan int)
	<-ch // no sender exists
}

func main() {
	go blockedReceive()
	select {}
}

/**
Why this leaks

	The goroutine blocks immediately.

	No other goroutine can ever unblock it.

	It will never reach return.

Observable signal

	No CPU usage.

	No logs.

	Only visible via goroutine count or dump.
**/
