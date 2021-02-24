package main

import (
	"fmt"
	"net/http"
)
//127.0.0.1:8881/test?WarTime=90
func handel(reponsewriter http.ResponseWriter, request  *http.Request){
	//设置允许解析参数
	request.ParseForm()
	//从参数价键值对获取相应的值
	fmt.Println(request.Form["WarTime"][0])
	fmt.Println(request)
	buff:=[]byte("hello")
	reponsewriter.Write(buff)

}
//bs
func main() {
	fmt.Println("服务端程序已启动")
	//http服务器配置第一步我们需要将url与注册回调函数
http.HandleFunc("/test",handel)
 //第二步我们需要将服务器地址与端口进行配置监听
http.ListenAndServe("127.0.0.1:8881",nil)


}