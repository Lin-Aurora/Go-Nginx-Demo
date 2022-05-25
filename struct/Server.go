package _struct

import (
	yml "github.com/kylelemons/go-gypsy/yaml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var index = 0

// Server 服务器结构体
type Server struct {
	// 反向代理服务器端口
	Server_Proxy_Port string `yaml:"server_Proxy_Port"`
	// 默认服务器
	Server_Default string `yaml:"server_Default"`
	// 可选服务器集
	Sub_Servers []string `yaml:"sub_Servers"`
}

// ReadConfig  配置读取
func (s *Server) ReadConfig() (*yml.File,error){
	//file, err := ioutil.ReadFile("./config.yml")
	file, err := yml.ReadFile("./config.yml")
	if err != nil {
		log.Print(err.Error())
		return nil,err
	}
	return file,nil
}

// ReturnConfig 配置返回
func ReturnConfig() *Server{
	s := new(Server)
	data, _ := ioutil.ReadFile("./config.yml")
	err := yaml.Unmarshal(data, &s)
	if err != nil {
		log.Println(err)
	}
	//d, _ := json.Marshal(s)
	return s
}

// GetSubServer 获取子服务器IP （这里使用轮询策略）
func GetSubServer(Sub_Server []string)  string{
	index = index % len(Sub_Server)
	host := Sub_Server[index]
	index++
	return host
}

// Run 服务器启动
func  Run(s string){
	log.Println("服务器正在运行 IP地址为：http://localhost:" + s)
	mux := http.NewServeMux()
	mux.HandleFunc("/",func (w http.ResponseWriter,r *http.Request) {
		w.Write([]byte(s))
	})
	server := &http.Server{
		Addr: ":" + s,
		ReadTimeout: time.Second * 5,
		WriteTimeout: time.Second * 5,
		Handler: mux,
	}
	go func(){
		if err := server.ListenAndServe();err != nil{
			log.Fatal("服务器启动失败：:",err)
		}
	}()
}