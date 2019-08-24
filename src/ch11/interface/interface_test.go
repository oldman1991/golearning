package interface_test

import (
	"fmt"
	"testing"
)

//type Programmer interface {
//	WriteHelloWorld() string
//}
//
//type GoProgrammer struct {
//
//}
//
//func (g *GoProgrammer) WriteHelloWorld() string{
//	return "fmt.Println(\"Hello World\")"
//}
//
//func TestClient(t *testing.T){
//	var p Programmer
//	p = new(GoProgrammer)
//	t.Log(p.WriteHelloWorld())
//}

// go接口与其他主要编程语言的差异
//1. 接口为非入侵性，实现不依赖于接口定义
//2. 所以接口定义可以包含在接口使用者包内


type Code string;

type Progarmmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct{

}

func (p *GoProgrammer) WriteHelloWorld() Code  {
	return "fmt.Println(\" Hello World!\")"
}


type JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code{
	return "System.out.Println(\"Hello World!\")"
}


func WriteFirstProgram(p Progarmmer){
	fmt.Printf("%T %v\n",p,p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T)  {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
}