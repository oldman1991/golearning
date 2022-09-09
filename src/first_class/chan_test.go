package first_class

import (
	"sync"
	"testing"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consumer(ch)
		wg.Done()
	}()

	wg.Wait()
}
