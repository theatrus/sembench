package sembench

import (
	"sync"
)

type Semaphore interface {
	Acquire()
	Release()
}

type channelSemaphore struct {
	channel chan int
}

func (s *channelSemaphore) Acquire() {
	<-s.channel
}

func (s *channelSemaphore) Release() {
	s.channel <- 0
}

func NewChannelSemaphore(max int, initial int) Semaphore {
	sem := channelSemaphore{channel: make(chan int, max)}
	for i := 0; i < initial; i++ {
		sem.channel <- 0
	}
	return &sem
}

type condSemaphore struct {
	m *sync.Mutex
	c *sync.Cond
	max int
	count int
}

func NewCondSemaphore(max int, initial int) Semaphore {
	m := &sync.Mutex{}
	c := sync.NewCond(m)
	
	sem := condSemaphore{max: max, count: initial, m: m, c: c}

	return &sem
}

func (sem *condSemaphore) Acquire() {
	sem.m.Lock()
	if sem.count > 0 {
		sem.count--
		sem.m.Unlock()
		return
	}
	for {
		sem.c.Wait()
		if sem.count > 0 {
			sem.count--
			sem.m.Unlock()
			return
		}
	}
}

func (sem *condSemaphore) Release() {
	sem.m.Lock()
	lastCount := sem.count
	if lastCount <= sem.max {
		sem.c.Signal()
	}
	sem.count++
	sem.m.Unlock()
}
