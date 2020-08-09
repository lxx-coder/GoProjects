package httpTest

import (
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := Get("http://127.0.0.1:18443/hello")
	if err != nil {
		t.Error("Get response failed:", err)
	}

	//f, err := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE, 0644)
	//if err != nil {
	//	t.Error("open file failed:", err)
	//}
	//defer f.Close()
	//
	//f.WriteString(resp)
	t.Log(resp)
}

func TestHeaderGet(t *testing.T) {
	resp, err := HeaderGet("http://127.0.0.1:8080/hello")
	if err != nil {
		t.Error("Get response failed；", err)
	}

	t.Log(resp)
}