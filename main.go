package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 读取某txt文件内容并渲染
	// b, _ := ioutil.ReadFile("./hello.txt") 旧版本的ioutil已经被os.ReadFile取代了
	b, err := os.ReadFile("./hello.txt")
	if err != nil {
		fmt.Printf("txt is error,err:%v\n", err)
	}
	_, _ = fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server faild,err:%v\n", err)
		return
	}
}
