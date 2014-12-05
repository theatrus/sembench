package sembench

import (
	"testing"
)

func TestChannelSemaphore(t *testing.T) {
	maxAllowed := 1
	initial := 1
	sem := NewChannelSemaphore(maxAllowed, initial)
	sem.Acquire()
	sem.Release()
	
}

func semBench(sem Semaphore, maxAllowed int) {
	sem.Acquire()
	sem.Acquire()

	sem.Acquire()
	sem.Release()
	sem.Acquire()
	sem.Release()
	
	sem.Release()
	sem.Release()
}

func BenchmarkChannelSemaphore(b *testing.B) {
	maxAllowed := 3000
	initial := 3000

	sem := NewChannelSemaphore(maxAllowed, initial)
	for i := 0; i < b.N; i++ {
		semBench(sem, maxAllowed)
	}	
}


func BenchmarkCondSemaphore(b *testing.B) {
	maxAllowed := 3000
	initial := 3000

	sem := NewCondSemaphore(maxAllowed, initial)
	for i := 0; i < b.N; i++ {
		semBench(sem, maxAllowed)
	}
}
