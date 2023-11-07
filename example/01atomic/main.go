package main

import "net/http"

// 不使用框架的一个web后台

func SayHello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	case "POST":
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/sayhello", SayHello)
	http.ListenAndServe("127.0.0.1:4001", nil)
}
