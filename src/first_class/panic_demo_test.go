package first_class

import (
	"fmt"
	"testing"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func Test(t *testing.T) {
	//data1 := []*field{{"one"}, {"two"}, {"three"}}
	//for _, v := range data1 {
	//	go (*field).print(v)
	//}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go (*field).print(&v)
	}

	time.Sleep(3 * time.Second)
}
