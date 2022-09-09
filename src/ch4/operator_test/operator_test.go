package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c:=[...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a==c)
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestOperator(t *testing.T) {
	// 按位置0
	// 1、如果右侧是0，则左侧数保持不变
	//
	//2、如果右侧是1，则左侧数一定清零
	a := 1
	b := 0
	t.Log(a &^ b)
	t.Log(b &^ a)

	c := 7 // 00000111
	d := 3 // 00000011
	t.Log(c &^ d)
	t.Log(d &^ c)
}
