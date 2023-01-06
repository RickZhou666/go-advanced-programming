package chap1

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var total1 uint64

func worker1(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total1, i)
	}

	fmt.Println("Total value is: ", total1)
}

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker1(&wg)
	go worker1(&wg)

	fmt.Println("Total value is: ", total1)

	wg.Wait()
}
