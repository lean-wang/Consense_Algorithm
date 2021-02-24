package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//dpos 给五个节点，轮流投票

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
//创世块
func GenesisBlock()Block{
	var genesisBlock=Block{0,time.Now().String(),"genesisblock",
		"0","",""}
	SortTempNode()
	index:=DPoS(genesisBlock.Index)
	genesisBlock.NodeAdress=tempNode[index].NodeAdress
	genesisBlock.Hash=hex.EncodeToString(claculateHash(&genesisBlock))
	Blockchain=append(Blockchain,genesisBlock)
	return genesisBlock

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
func CreateNextBlock(oldBlock Block,data string)Block{
	var newBlock=Block{oldBlock.Index+1,time.Now().String(),data,
		oldBlock.Hash,"",""}
	SortTempNode()
	index:=DPoS(newBlock.Index)
	newBlock.NodeAdress=tempNode[index].NodeAdress
	newBlock.Hash=hex.EncodeToString(claculateHash(&newBlock))
	Blockchain=append(Blockchain,newBlock)
	return newBlock
}
func DPoS(index int)int{
	index=index%5
	return index
}
//添加区块的节点结构体
type Node struct {
	//持票数
	Token int
	//节点地址
	NodeAdress string
}
//用记录区块信息的节点的数组
var  tempNode []*Node

//初始化五个记账节点
func CreateNode(){
	var newNode1=Node{0,"01"}
	var newNode2=Node{0,"02"}
	var newNode3=Node{0,"03"}
	var newNode4=Node{0,"04"}
	var newNode5=Node{0,"05"}
	appendTempNode(&newNode1)
	appendTempNode(&newNode5)
	appendTempNode(&newNode2)
	appendTempNode(&newNode3)
	appendTempNode(&newNode4)
}

func appendTempNode(node *Node){
		node.Token=Voted()
		tempNode=append(tempNode,node)
}


//投票函数
func Voted()int{
	rand.Seed(time.Now().UnixNano())
	shake:=rand.Intn(10)
	return shake

}
//对于五个节点一句token多少来进行排序

func SortTempNode(){
	for i:=0;i<len(tempNode);i++{
		for j:=0;j<len(tempNode)-1;j++{
			if tempNode[j].Token<tempNode[j+1].Token{
				tempNode[j],tempNode[j+1]=tempNode[j+1],tempNode[j]
			}
		}

	}

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
func printTempNode(){
	for i:=0;i<len(tempNode) ;i++  {
		node:=tempNode[i]
		fmt.Println("===============")
		fmt.Printf("节点持票数量%d个，节点地址为%s\n",node.Token,node.NodeAdress)
		fmt.Println("==============")

	}
}
func main() {
CreateNode()
temp:=GenesisBlock()
	for i:=0;i<7 ;i++  {
		str:="我是第"+strconv.Itoa(i+1)+"个区块"
		temp=CreateNextBlock(temp,str)

	}
	printTempNode()
	printBlockchain()
}
