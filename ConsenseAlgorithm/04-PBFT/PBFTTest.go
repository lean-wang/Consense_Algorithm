package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type nodeInfo struct {
	Id string
	Url string
	Writer http.ResponseWriter
}
//存放四个国家的地址
var nodeTable =map[string]string{
	"N0":"localhost:8881",
	"N1":"localhost:8882",
	"N2":"localhost:8883",
	"N3":"localhost:8884",
}


func main() {
	userId:=os.Args[1]
	fmt.Println(userId)
	node:=nodeInfo{userId,nodeTable[userId],nil}
	fmt.Println(node)
	//注册号四个阶段的回调函数
	http.HandleFunc("/req",node.request)
	http.HandleFunc("/prePrepare",node.prePrepare)
	http.HandleFunc("/prepare",node.prepare)
	http.HandleFunc("/commit",node.commit)
	//
	err:=http.ListenAndServe(node.Url,nil)
	if err!=nil{
		fmt.Println("httperr ",err)
		os.Exit(1)
	}

}
func(node *nodeInfo)request(reponsewriter http.ResponseWriter, request  *http.Request)  {
//	设置参数解析
request.ParseForm()
//开始解析参数
if (len(request.Form["WarTime"])>0){
	node.Writer=reponsewriter
	fmt.Println(node.Writer)
 //主机节点激活开始广播其他节点进入下一个阶段
 node.broadcast(request.Form["WarTime"][0],"/prePrepare")
}

}
//
func(node *nodeInfo)prePrepare(reponsewriter http.ResponseWriter, request  *http.Request)  {
	//	设置参数解析
	fmt.Println("进入perperPare阶段")
	request.ParseForm()
	if (len(request.Form["WarTime"])>0) {
		//接受到主节点的信息，开始广播给其他人
		fmt.Println(request.Form["WarTime"][0])
		fmt.Println(node.Id)
		node.broadcast(request.Form["WarTime"][0],"/prepare")
	}
	}
func(node *nodeInfo)prepare(reponsewriter http.ResponseWriter, request  *http.Request) {
	fmt.Println("进入perPare阶段")
	request.ParseForm()
    //
	if (len(request.Form["WarTime"])>0) {
		fmt.Println(request.Form["WarTime"][0])
	}
	if (len(request.Form["nodeId"])>0){
		fmt.Println(request.Form["nodeId"][0])
	}
	node.checkout(request)

}
var checkoutMap=make(map[string]string)
func(node *nodeInfo)checkout( request  *http.Request) {
	fmt.Println("进入检验阶段")
	request.ParseForm()
	if (len(request.Form["nodeId"])>0){
		checkoutMap[request.Form["nodeId"][0]]="OK"
		fmt.Println("id  ",request.Form["nodeId"][0])
	}
	//总结点数
	n:=len(nodeTable)
	//故障节点数量
	m:=n-len(checkoutMap)
	fmt.Println("总结点数",n)
	fmt.Println("故障节点",m)
	if 3*m+1<=n{
		node.broadcast(request.Form["WarTime"][0],"/commit")
	}
}
var flag=true
func(node *nodeInfo)commit(reponsewriter http.ResponseWriter, request  *http.Request) {
	if flag{
	fmt.Println("commit阶段")
	io.WriteString(node.Writer,"ok")
	fmt.Println(node.Writer)
	flag=false
	}

}
//广播

func (node *nodeInfo)broadcast(msg string,path string)  {
	for _,url:=range nodeTable{

		http.Get("http://"+url+path+"?WarTime="+msg+"&nodeId="+node.Id)
		fmt.Println("广播成功")
	}


}