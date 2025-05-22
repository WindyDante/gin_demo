package main

import (
	"fmt"
	"net/http"
)

var val = 0

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<h1>Hello World</h1>")
	val += 1
	fmt.Printf("%d", val)
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server faild,err:%v\n", err)
		return
	}
}
