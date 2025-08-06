package async

import (
	"sync"
)

// Semaphore простейшая реализация семафора
type Semaphore struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

// NewSemaphore .
func NewSemaphore(buf int) *Semaphore {
	return &Semaphore{
		c:  make(chan struct{}, buf),
		wg: &sync.WaitGroup{},
	}
}

// Go запускает функцию, если это допустимо по буферу
func (s *Semaphore) Go(f func()) {
	s.wg.Add(1)
	go func() {
		s.c <- struct{}{}
		defer s.release()
		f()
	}()
}

func (s *Semaphore) release() {
	<-s.c
	s.wg.Done()
}

// Wait ожидает выполнения всех запущенных функций
func (s *Semaphore) Wait() {
	s.wg.Wait()
}
