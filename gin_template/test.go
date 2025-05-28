package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模版对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	checkErr(err)
	// 利用给定数据渲染模版
	err = tmpl.Execute(w, "windEast")
	checkErr(err)
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	checkErr(err)
}
