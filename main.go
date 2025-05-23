package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // 解析参数
	if err != nil {
		fmt.Println("ParseForm err:", err)
	}
	for k, v := range r.Form {
		// 遍历get请求的参数
		fmt.Println("key:", k)
		// value是一个切片,通过strings方法组成一个字符串返回
		fmt.Println("val:", strings.Join(v, " "))
	}
	_, err = fmt.Fprintf(w, "Hello EastWind!")
	if err != nil {
		return
	}
}

func main() {
	// 拦截器的方法
	http.HandleFunc("/", sayHelloName)
	// 监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("ListenAndServe err:", err)
	}
}
