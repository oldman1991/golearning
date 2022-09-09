package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThenHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fibList := []int{1, 1}

	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

//func GetFibonacci(n int) ([]int,error){
//	if n<0 || n>100 {
//		return nil, errors.New("n should be in [2,100]")
//	}
//	fibList := []int{1,1}
//	for i:=2; i<n ;i++{
//		fibList = append(fibList, fibList[i-2] + fibList[i-1])
//
//	}
//	return fibList, nil
//}

// 嵌套错误处理，避免这种写法
func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

// 使用这一种错误处理机制
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1000); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
