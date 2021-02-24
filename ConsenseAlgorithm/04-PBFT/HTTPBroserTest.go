package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	res,err:=http.Get("http://127.0.0.1:8881/test")
	if err!=nil{
		fmt.Println("错误",err)
		os.Exit(1)
	}
	defer res.Body.Close()
	fmt.Println(res)
	fmt.Println(res.Body)
	body,_:=ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

}