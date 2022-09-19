package main

import "net/http"

//http框架原理梳理
//1.输出hello world
func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!\n"))
}

func main() {
	//2.注册路由处理函数
	http.HandleFunc("/sayHello", sayHello)
	//curl http://127.0.0.1:8081/sayHello

	//2.1文件浏览请求
	fileHandler := http.FileServer(http.Dir("./video"))
	http.Handle("/video/", http.StripPrefix("/video/", fileHandler))
	//浏览器 http://127.0.0.1:8081/video/test.mp4

	//3.启动web服务
	http.ListenAndServe(":8081", nil)
}
