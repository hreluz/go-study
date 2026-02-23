type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func (s *SafeMap) Get(k string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[k]
}

func (s *SafeMap) Set(k string, v int) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

/**
What Happens Under Heavy Read Concurrency?

	Suppose:

		1,000 goroutines frequently read from the map.

		Writes are rare.

	With a mutex:

		Only one goroutine can access the map at a time.

		Even reads block each other.

		This becomes a bottleneck.

	The mutex serializes everything.
**/