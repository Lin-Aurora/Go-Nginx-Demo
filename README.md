# Go-Nginx-Demo
## 简易Golang语言实现的类Nginx中间件功能Demo

目前实现功能有：端口监听转发

### 使用：
```
项目启动：./Go-Nginx
```
### 项目启动后访问http://127.0.0.1:2202如图所示：   （代理服务器端口和服务器端口皆可通过config.yml修改）
![image](https://user-images.githubusercontent.com/64384229/170268403-0742d860-8f49-4595-94f9-63f4fc4ab720.png)

### 同时集成了压力测试go-stress-testing，可提供压力测试
```
压力测试：./go-stress-testing -c 并发数 -n 每个并发执行数 http://127.0.0.1:2202
```
### 部分测试条件：并发数为10 执行数为1000 的结果如图所示
![image](https://user-images.githubusercontent.com/64384229/170268692-94480c3e-d0d1-421b-94cb-b715a1ab5d87.png)
![image](https://user-images.githubusercontent.com/64384229/170268711-554107d4-e947-4080-af24-27b040f05922.png)
