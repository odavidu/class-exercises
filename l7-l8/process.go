package main

type Process struct {
	clock []int
	pid   int
	N     int
}

func (p *Process) Start(N int, PID int) {
	p.clock = make([]int, N)
	p.pid = PID
	p.N = N
}

func (p *Process) Internal() {
	p.clock[p.pid]++
}

func (p *Process) Send() []int {
	p.Internal()
	return p.clock
}

func (p *Process) Receive(ts []int) {
	for i := 0; i < p.N; i++ {
		p.clock[i] = max(ts[i], p.clock[i])
	}

	p.Internal()
}

func Compare(ts1 []int, ts2 []int) int {
	var comparison = 0
	for key, val := range ts1 {
		if val > ts2[key] {
			if comparison == -1 {
				return 0
			}
			comparison = 1
			continue
		} else if val < ts2[key] {
			if comparison == 1 {
				return 0
			}
			comparison = -1
			continue
		} else {
			continue
		}
	}
	return comparison
}
