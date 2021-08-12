package limiter

import (
	"reflect"
	"sync"
)

// Limiter holds information for controler
type Limiter struct {
	limiter chan int
	group   sync.WaitGroup
}

//New creates a controler with the specified number of resources
func New(max int) (l *Limiter) {
	l = new(Limiter)
	l.limiter = make(chan int, max)

	for len(l.limiter) < max {
		l.limiter <- 1
	}

	return
}

// Execute consumes one resource to run the function, f must be a function and ini must have same arguments of f (types and amount) otherwise it panics
func (l *Limiter) Execute(f interface{}, in ...interface{}) {
	l.group.Add(1)
	<-l.limiter

	go func() {
		fun := reflect.ValueOf(f)
		args := make([]reflect.Value, len(in))

		for i, val := range in {
			args[i] = reflect.ValueOf(val)
		}

		fun.Call(args)

		l.limiter <- 1
		l.group.Done()
	}()
}

// Wait waits for all processes to finish to continue
func (l *Limiter) Wait() {
	l.group.Wait()
}
