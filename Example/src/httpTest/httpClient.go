package httpTest

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Get(url string) (body string, err error) {
	if len(url) == 0 {
		return "", errors.New("url is empty")
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("response error")
	}
	return string(b), nil
	//fmt.Println(resp.Body)
}

func HttpGet(url string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	out, err := os.Create("out.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	//buffer := bytes.NewBuffer(make([]byte, 1024))
	_, err = io.Copy(out, r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Header: %v\n", r.Header)
	fmt.Printf("StatusCode: %v\n", r.StatusCode)
}
func HeaderGet(url string) (string, error) {
	req, err := http.NewRequest("Get", url,nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}