package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mutex sync.Mutex
var useMutext bool
var counter int
var list []int

func Increment() {
	if useMutext {
		mutex.Lock()         // Adquire o bloqueio do Mutex antes de modificar o recurso compartilhado
		defer mutex.Unlock() // Libera o bloqueio do Mutex após o fim da função
	}
	counter++
	list = append(list, counter)
	fmt.Println("Counter:", counter)
	time.Sleep(time.Millisecond)
}

func TestCounterWithMutex(t *testing.T) {
	// When | Arrange
	useMutext = true
	counter = 0
	list = []int{}

	// prepara o WaitGroup para aguardar a execução de 10 goroutines
	wg := sync.WaitGroup{}
	quantityGoroutines := 10
	wg.Add(quantityGoroutines)

	// When | Act
	for i := 0; i < quantityGoroutines; i++ {
		go func() {
			defer wg.Done()
			Increment()
		}()
	}
	wg.Wait() // Executa todas as goroutines e aguarda o fim de todas elas

	// Then | Assert
	assert.Equal(t, 10, counter)
	assert.Equal(t, 1, list[0])
	assert.Equal(t, 2, list[1])
	assert.Equal(t, 3, list[2])
	assert.Equal(t, 4, list[3])
	assert.Equal(t, 5, list[4])
	assert.Equal(t, 6, list[5])
	assert.Equal(t, 7, list[6])
	assert.Equal(t, 8, list[7])
	assert.Equal(t, 9, list[8])
	assert.Equal(t, 10, list[9])
}

func TestCounterWithoutMutex(t *testing.T) {
	// When | Arrange
	useMutext = false
	counter = 0
	list = []int{}

	// prepara o WaitGroup para aguardar a execução de 10 goroutines
	wg := sync.WaitGroup{}
	quantityGoroutines := 10
	wg.Add(quantityGoroutines)

	// When | Act
	for i := 0; i < quantityGoroutines; i++ {
		go func() {
			defer wg.Done()
			Increment()
		}()
	}
	wg.Wait() // Executa todas as goroutines e aguarda o fim de todas elas

	// Then | Assert
	assert.Equal(t, 10, counter)

	// Como o Mutex não foi usado, aqui vai ter uma inconsistência, pois várias goroutines vão modificar o mesmo recurso
	assert.Equal(t, 1, list[0])
	assert.Equal(t, 2, list[1])
}
