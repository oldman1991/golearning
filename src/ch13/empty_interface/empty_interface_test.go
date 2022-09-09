package empty_interface

import (
	"fmt"
	"testing"
)

/*
1.空接口可以表示任何类型
2.通过断言来将空接口转换为定制类型
	v,ok := p.(int) //ok=true时转换成功
*/

func DoSomthing(p interface{}) {
	//if i,ok :=p.(int);ok{
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if j,ok:= p.(string);ok{
	//	fmt.Println("String", j)
	//	return
	//}
	//fmt.Println("Unknow Type")

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")

	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomthing(10)
	DoSomthing("aaa")
}
