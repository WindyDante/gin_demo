package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		// 生成时间戳
		currentTime := time.Now().Unix()
		// 新建md5
		h := md5.New()
		// 将时间戳转为10进制字符串写入md5
		_, err := io.WriteString(h, strconv.FormatInt(currentTime, 10))
		if err != nil {
			log.Fatalln("io.WriteString err:", err)
		}
		// 将md5转为16进制字符串
		token := fmt.Sprintf("%x", h.Sum(nil))

		// 解析文件
		t, _ := template.ParseFiles("upload.gtpl")
		// 写入token到文件中
		err = t.Execute(w, token)
		if err != nil {
			log.Fatalln("t.Execute err:", err)
		}
	} else {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			return
		}
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			log.Fatalln("r.FormFile err:", err)
		}
		defer file.Close()
		_, err = fmt.Fprintf(w, "%v", handler.Header)
		if err != nil {
			return
		}
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatalln("os.OpenFile err:", err)
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	// 拦截器的方法
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	// 监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("ListenAndServe err:", err)
	}
}
