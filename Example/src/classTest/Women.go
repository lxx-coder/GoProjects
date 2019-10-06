package classTest

//继承特性
//采用的是匿名组合的方式，Women结构体中包含匿名字段Person，那么Person中的属性也属于

type Women struct {
	Person
	Sex string
}
