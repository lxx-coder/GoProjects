package classTest

import "fmt"

/*
Go语言中的面向对象
Go语言并没有提供类class关键字，只提供了struct结构体类型
struct类型可以包含变量和方法
*/

type Person struct {
	Name string
}

//公有属性和私有属性的区分机制就是方法或属性书否是首字母大写，
// 如果大写就是公有的，小写就是私有变量
func (person *Person) SetName(name string){
	person.Name = name
}

func (person *Person) GetInfo() {
	fmt.Println(person.Name)
}
