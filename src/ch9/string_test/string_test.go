package string_test

import "testing"

func TestString(t *testing.T){
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	// s[1] = '3' //String是不可变的byte slice
	t.Log(s)
	t.Log(len(s))
}