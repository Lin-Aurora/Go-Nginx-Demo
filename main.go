package main

import (
	"Go-Nginx/handler"
	"Go-Nginx/struct"
	"log"
	"net/http"
)

var (
	Server_Proxy_Port string
	Server_Default string
	Sub_Server []string


)

func init() {
	// 初始化配置服务器参数
	s := _struct.ReturnConfig()
	Server_Proxy_Port = s.Server_Proxy_Port
	Server_Default = s.Server_Default
	Sub_Server = s.Sub_Servers

}




func main() {
	// 启动服务器
	for _ =  range Sub_Server{
		_struct.Run(_struct.GetSubServer(Sub_Server))
	}

	// 启动代理服务器
	http.HandleFunc("/", handler.RequestHandler)
	log.Println("代理服务器正在运行 IP地址为：http://localhost:" + Server_Proxy_Port)
	log.Fatal(http.ListenAndServe(":" + Server_Proxy_Port, nil))

	//// 阻塞协程
	//ch := make(chan int)
	//<- ch
}
