package first_class

import (
	"fmt"
	"testing"
)

func MyAppend(sl []int, elems ...int) []int {
	fmt.Printf("%T\n", elems)
	if len(elems) == 0 {
		println("no elems to append")
		return sl
	}
	sl = append(sl, elems...)
	return sl
}

func TestMyAppend(t *testing.T) {
	sl := []int{1, 2, 3}
	sl = MyAppend(sl)
	sl = MyAppend(sl, 4, 5, 6)
	fmt.Println(sl)
}
