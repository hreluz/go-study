package main

func blockedSelect(ch <-chan int) {
	select {
	case <-ch:
		// never happens
	}
}

func main() {
	ch := make(chan int)
	go blockedSelect(ch)
	select {}
}
