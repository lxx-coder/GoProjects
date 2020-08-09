package main

import (
	"bufio"
	"classTest"
	"fmt"
	"github.com/astaxie/beego"
	"goroutineTest"
	"httpTest"
	"io"
	"os"
	"reflect"
	"runtime"
)

type Rect struct {
	x, y float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func NewRect(x,y,width,height float64) *Rect {
	fmt.Println("New Rect")
	return &Rect{x,y,width,height}
}

func eatTest(eater classTest.Eater){
	eater.Eat()
}


func main() {

	fmt.Println("num of cpu: ", runtime.NumCPU())
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(bufio.ScanWords)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}


	s := "截取中文"
	//试试这样能不能截取?
	res := []rune(s)
	fmt.Println(string(res[:2]))

	fmt.Println(s[:6])
	res1 := []byte(s)
	fmt.Println(string(res1[:6]))

	x :=[]int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)

	var students = make(map[string]string)
	students["001"] = "zhangsan"
	stuType := reflect.TypeOf(students)
	stuValue := reflect.ValueOf(students)
	stuType.Kind()
	fmt.Println("type:", stuType)
	fmt.Println("value", stuValue)

	//log.Trace.Println("I have something standard to say")
	//log.Info.Println("Special Information")
	//log.Trace.Println("There is something you need to know about")
	//log.Error.Println("Something has failed")

	fmt.Println("numCPU: ", runtime.NumCPU())
	url := "http://www.baidu.com"
	httpTest.HttpGet(url)

	reader := bufio.NewReader(os.Stdin)
	var str []string
	for {
		result, err := reader.ReadString('\n')
		if err != nil  || io.EOF == err{
			return
		}
		str= append(str, result)
	}

	fmt.Println(str)
	//strSplit := strings.Split(result, "")

	//httpTest.Start()

	//rect := new(Rect{0,0,100,200})
	//rect := &Rect{}
	//fmt.Printf("%v\n",rect)
	//name := new(interface{})
	//name["name"] = "name"
	//fmt.Printf("type: %T, size:%d", *name, unsafe.Sizeof(*name))

	//db := mysqlTest.Connect()
	//cust, err := mysqlTest.QueryOne(db)
	//if err != nil {
	//	fmt.Printf("query error: %v\n", err)
	//}
	//fmt.Printf("query result: %v\n", cust)

	//var v1 interface{}="string"
	//fmt.Printf("type:%T, size:%d", v1, unsafe.Sizeof(v1))
	//myslice := []int{1,2,3,4,5}
	//myslice1 := make([]int,5,10)
	//fmt.Printf("size: %d, cap: %d", len(myslice), cap(myslice))
	//fmt.Printf("size: %d, cap: %d", len(myslice1), cap(myslice1))
	//var a int32 = 777
	//fmt.Printf("type:%T, size:%d", a, unsafe.Sizeof(a))

	beego.Run()
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