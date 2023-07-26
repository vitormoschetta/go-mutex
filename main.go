package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu        sync.Mutex
	useMutext bool
	value     int
}

func (c *Counter) Increment() {
	if c.useMutext {
		c.mu.Lock()         // Adquire o bloqueio do Mutex antes de modificar o recurso compartilhado
		defer c.mu.Unlock() // Libera o bloqueio do Mutex após o fim da função
	}
	c.value++
	fmt.Println(c.value)
	time.Sleep(time.Millisecond)
}

func main() {
	var counter Counter
	counter.useMutext = false // Altere para true para usar Mutex
	wg := sync.WaitGroup{}
	quantity := 10
	wg.Add(quantity)
	for i := 0; i < quantity; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter.value)
}
