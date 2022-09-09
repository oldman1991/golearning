package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on somethins else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("retuen result")
		retCh <- ret
		fmt.Println("service existed.")
	}()

	return retCh
}

func BenchmarkAsyncService(t *testing.B) {
	rech := AsyncService()
	otherTask()

	fmt.Println(<-rech)
}
