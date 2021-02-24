package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Block struct {
	//区块高度
	Index int
	//时间戳
	Timestamp string
	//数据
	Data string
	//前一个区块的hash
	Prehash string
	// 本区块的hash
	Hash string
   //添加该区块的节点
     NodeAdress string

}
var Blockchain []Block
//添加区块的节点结构体
type Node struct {
	//币龄
	CoinAge int
//	持币数量
 	CoinNum int
	//节点地址
	NodeAdress string
}
//随即开始抽取记账节点的数组
var  tempNode []Node
func CreateNode(){
	var newNode1=Node{10,100,"01"}
	var newNode2=Node{20,99,"02"}
	var newNode3=Node{30,40,"03"}
	var newNode4=Node{40,90,"04"}
	var newNode5=Node{20,88,"05"}
	appendTempNode(newNode1)
	appendTempNode(newNode5)
	appendTempNode(newNode2)
	appendTempNode(newNode3)
	appendTempNode(newNode4)
}
func appendTempNode(node Node){
	var i=0
	for  i=0;i<node.CoinAge*node.CoinNum;i++{
		tempNode=append(tempNode,node)
	}
}
//pos将区块由谁来添加到区块链的节点下标
func PoS() int{
rand.Seed(time.Now().UnixNano())
index:=rand.Intn(len(tempNode))
return index
}
//创世块
func GenesisBlock()Block{
	var genesisBlock=Block{0,time.Now().String(),"genesisblock",
		"0","","01"}

	genesisBlock.Hash=hex.EncodeToString(claculateHash(&genesisBlock))
	Blockchain=append(Blockchain,genesisBlock)
	return genesisBlock

}
//创建下一个区块
func CreateNextBlock(oldBlock Block,data string)Block{
	var newBlock=Block{oldBlock.Index+1,time.Now().String(),data,
		oldBlock.Hash,"",""}
	index:=PoS()
	newBlock.NodeAdress=tempNode[index].NodeAdress
	newBlock.Hash=hex.EncodeToString(claculateHash(&newBlock))
	Blockchain=append(Blockchain,newBlock)
	return newBlock
}

//计算区块的hash值
func claculateHash(block *Block)[]byte{
	record:=strconv.Itoa(block.Index)+block.NodeAdress+
		block.Data+block.Timestamp+block.Prehash
	h:=sha256.New()
	h.Write([]byte(record))
	hashed:=h.Sum(nil)
	return hashed

}


func printBlockchain(){
	for i:=0;i<len(Blockchain);i++{
		block:=Blockchain[i]
		fmt.Println("-------------------------------")
		fmt.Printf("区块链高度:%d,时间戳:%s,交易信息:%s,前一区块哈希值:%s,区块哈希：%s，" +
			"节点地址：%s\n",block.Index,block.Timestamp,block.Data,block.Prehash,
				block.Hash,block.NodeAdress)
		fmt.Println("-----------------------------")
	}
}
func main() {
	CreateNode()
genblock:= GenesisBlock()
temp:=genblock
	for i:=0;i<5 ;i++  {
		str:="我是第"+strconv.Itoa(i)+"个区块"
		temp=CreateNextBlock(temp,str)
	}
printBlockchain()
}
