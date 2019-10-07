package httpTest

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := Get("http://www.baidu.com")
	if err != nil {
		t.Error("Get response failed:", err)
	}

	f, err := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		t.Error("open file failed:", err)
	}
	defer f.Close()

	f.WriteString(resp)
	t.Log(resp)
}
