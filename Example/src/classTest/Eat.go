package classTest

import "fmt"

//多态特性
//两个结构体继承同一个接口
type Eater interface {
	Eat()
}

type Cat struct {

}

type Dog struct {

}

func (cat *Cat) Eat() {
	fmt.Println("Cat Eat")
}

func (dog *Dog) Eat() {
	fmt.Println("Dog Eat")
}