package httpTest

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"reflect"
	"runtime"
	"sort"
)

const TOKEN = "weixin"

func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func myHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("request:%v\n", request.Header)
	response.Write([]byte("Hello World"))
}

func doWeixinConnect(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	signature := req.Form["signature"]
	timestamp := req.Form["timestamp"]
	nonce := req.Form["nonce"]
	echostr := req.Form["echostr"]

	arr := []string{TOKEN, timestamp[0], nonce[0]}
	sort.Strings(arr)
	arrStr := arr[0] + arr[1] + arr[2]
	arrSha1 := str2sha1(arrStr)
	if arrSha1 == signature[0] {
		fmt.Fprintf(resp, echostr[0])
	}

}

func Start() {
	http.HandleFunc("/hello", myHandler)
	http.HandleFunc("/weixin", doWeixinConnect)

	//该方法用于在指定的 TCP 网络地址 addr 进行监听，然后调用服务端处理程序来处理传入的连接请求
	//该方法有两个参数：第一个参数 addr 即监听地址；第二个参数表示服务端处理程序，通常为nil
	//当参2为nil时，服务端调用 http.DefaultServeMux 进行处理
	fmt.Println("Starting the server")
	http.ListenAndServe("127.0.0.1:8080",nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func Server() {
	server := http.Server{
		Addr:              ":8080",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	fmt.Printf("server start...")
	server.ListenAndServe()
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct {

}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func Server1() {
	server := http.Server{
		Addr:              "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	world := WorldHandler{}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function call - " + name)
		h(writer, request)
	}
}

func LogServer() {
	server := http.Server{
		Addr:              "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}

func hello1(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s\n", p.ByName("name"))
}

func RouteGet() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello1)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler:mux,
	}

	server.ListenAndServe()
}

type MainController struct {
	beego.Controller
}

func (controller *MainController) Get() {
	controller.Ctx.WriteString("hello world")
}

func BeegoServer() {
	beego.Router("/", &MainController{})
	beego.Run()
}