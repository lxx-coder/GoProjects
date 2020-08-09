package httpTest

import (
    "net/http"
    "testing"
)

func TestStart(t *testing.T) {
    Start()
}

func TestServer(t *testing.T) {
    Server()
}

func TestBeegoServer(t *testing.T) {
    BeegoServer()
}

func TestBody(t *testing.T) {
    server := http.Server{
        Addr: "127.0.0.1:8089",
    }

    http.HandleFunc("/body", Body)
    server.ListenAndServe()
}