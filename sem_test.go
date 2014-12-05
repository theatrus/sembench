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

func BenchmarkChannelSemaphore(b *testing.B) {
	maxAllowed := 3000
	initial := 3000

	sem := NewChannelSemaphore(maxAllowed, initial)

	for i := 0; i < maxAllowed; i++ {
		sem.Acquire()
	}
	for i := 0; i < b.N; i++ {
		sem.Release()
		sem.Acquire()
	}
}
