package first_class

import (
	"errors"
	"fmt"
	"testing"
)

func spawn(f func() error) <-chan error {
	c := make(chan error)
	go func() {
		c <- f()
	}()
	return c
}

func TestGo(t *testing.T) {
	c := spawn(func() error {
		//time.Sleep(2 * time.Second)
		return errors.New("timeout")
	})
	fmt.Println(<-c)
}
