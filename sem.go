package sembench

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
