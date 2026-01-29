package main

func blockedLeak() {
	ch := make(chan int)
	<-ch // waits forever; no sender exists
}

func main() {
	go blockedLeak()
	select {} // keep process alive
}

