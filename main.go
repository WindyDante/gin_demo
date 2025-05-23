package main

import (
	"fmt"
	"html/template"
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

func login(w http.ResponseWriter, r *http.Request) {
	// 获取请求的方法
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		// 解析gtpl模版
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 解析表单参数
		err := r.ParseForm()
		if err != nil {
			log.Fatalln("ParseForm err:", err)
		}
		// 循环遍历表单
		for k, v := range r.Form {
			// 如果存在表单携带了多个参数，会返回一个切片
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, " "))
		}
	}
}

func main() {
	// 拦截器的方法
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	// 监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("ListenAndServe err:", err)
	}
}
