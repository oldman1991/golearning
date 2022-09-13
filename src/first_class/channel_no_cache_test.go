package first_class

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type signal struct{}

func worker() {
	println("worker is working...")
	time.Sleep(time.Second)
}

func spawn01(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func TestChannelNoCache(t *testing.T) {
	println("start a worker...")
	c := spawn01(worker)
	<-c
	fmt.Println("worker work done...")
}

func worker2(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d: works done...\n", i)
}

type signal11 struct {
}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal11) <-chan signal11 {
	c := make(chan signal11)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d: start to worker...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}
	go func() {
		wg.Wait()
		c <- signal11{}
	}()
	return c
}

func TestChannelNoCacheOneToN(t *testing.T) {
	fmt.Println("Start a group of workses...")
	groupSignal := make(chan signal11)
	c := spawnGroup(worker2, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	//close(groupSignal)
	for i := 0; i < 5; i++ {
		groupSignal <- signal11{}
	}
	<-c
	fmt.Println("the group of workers work done!")

}

type counter struct {
	c chan int
	i int
}

func NewCounter() *counter {
	cter := &counter{
		c: make(chan int),
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counter) Incres() int {
	return <-cter.c
}

func TestChannelNoCacheLock(t *testing.T) {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Incres()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
