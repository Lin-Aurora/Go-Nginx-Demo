package handler

import (
	"Go-Nginx/struct"
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)


// RequestHandler 反向代理请求
func RequestHandler (w http.ResponseWriter,r *http.Request){
	port := _struct.GetSubServer(_struct.ReturnConfig().Sub_Servers)
	serverHost := "http://localhost:" + port
	proxy ,err := url.Parse(serverHost)
	if err != nil{
		log.Println(err)
		return
	}

	transport := http.DefaultTransport
	r.URL.Host = proxy.Host
	r.URL.Scheme = proxy.Scheme
	resp ,err := transport.RoundTrip(r)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for key,value := range resp.Header{
		for _,v := range value{
			w.Header().Add(key,v)
		}
	}
	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(w)

	log.Println("代理服务器被调用：http://localhost:" + port)
}

// NewProxyServer 新建代理服务器
func NewProxyServer (Host string) (*httputil.ReverseProxy, error) {
	s, err := url.Parse(Host)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	server := httputil.NewSingleHostReverseProxy(s)
	//server.ErrorHandler = errorHandler()
	server.ModifyResponse = updateResponseHandler()

	director := server.Director
	server.Director = func(req *http.Request) {
		director(req)
		updateRequestHandler(req)
	}
	return server,err
}

// errorHandler 代理错误处理
func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("当前请求响应出现错误: %v \n", err)
		return
	}
}

// updateRequestHandler 请求头设置
func updateRequestHandler(req *http.Request) {
	req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
	//req.URL.Host
}

// updateResponseHandler 请求响应设置
func updateResponseHandler() func(*http.Response) error {
	return func(resp *http.Response) error {
		return errors.New("当前请求响应无效！")
	}
}
