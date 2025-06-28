package async

import (
	"sync"
)

// Group удобная надстройка над sync.WaitGroup
type Group struct {
	wg *sync.WaitGroup
}

// NewGroup .
func NewGroup() *Group {
	return &Group{
		wg: &sync.WaitGroup{},
	}
}

// Go запускает горутину
func (g *Group) Go(f func()) {
	g.wg.Add(1)
	go func() {
		f()
	}()
}

// Wait ожидает выполнения всех запущенных функций
func (g *Group) Wait() {
	g.wg.Wait()
}
