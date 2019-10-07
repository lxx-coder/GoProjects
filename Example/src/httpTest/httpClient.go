package httpTest

import (
	"errors"
	"io/ioutil"
	"net/http"
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
