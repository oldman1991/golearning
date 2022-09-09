package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 100)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned resule.")
		retCh <- ret
		fmt.Println("service existed.")
	}()

	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 500):
		t.Error("time out")

	}
	//time.Sleep(time.Second*1)
}
