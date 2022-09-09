package first_class

import (
	"fmt"
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
