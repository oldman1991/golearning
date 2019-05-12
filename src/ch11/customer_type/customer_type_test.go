package customer_type

import (
	"fmt"
	"testing"
	"time"
)

//自定义类型

type InConv func(op int) int


func timeSpent(inner InConv) InConv{
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op
}


func TestFn(t *testing.T){
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}