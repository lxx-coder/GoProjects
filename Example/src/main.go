package main

import (
	"classTest"
	"fmt"
	"goroutineTest"
)

func eatTest(eater classTest.Eater){
	eater.Eat()
}

func main() {
	fmt.Println("Hello world")

	goroutineTest.Routine()
	person := classTest.Person{"Zhangsan"}
	person.SetName("Lisi")
	person.GetInfo()

	women := classTest.Women{classTest.Person{"wangwu"},"女"}
	fmt.Println(women.Name)
	fmt.Println(women.Sex)

	//多态特性
	cat := classTest.Cat{}
	dog := classTest.Dog{}

	eatTest(&cat)
	eatTest(&dog)

}