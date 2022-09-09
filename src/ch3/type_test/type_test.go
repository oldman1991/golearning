package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	// 不支持隐形类型转化，不支持别名到元类型的隐私转换
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr +1  go 中不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T, %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string //go中字符串是值类型，其默认的初始化值为空字符串，而不是nil
	t.Log("*" + s + "*")
	t.Log(len(s))
}
