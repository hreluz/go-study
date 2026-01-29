package main

func spinForever() {
	for {
		// no exit condition
	}
}

func main() {
	go spinForever()
	select {}
}

/**
	The loop has no conditional exit.

	Control flow can never reach return.

	The goroutine is actively running forever.
**/
