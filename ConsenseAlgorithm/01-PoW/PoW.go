package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
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
	//随机值
	Nonce int
	//难度值
	Difficulty int
}
//区块链，可以用linklist来代替。
var Blockchain []Block
//生成创世区块
func genesisBlock()Block{
 		var genesisBlock=Block{0,time.Now().String(),"genesisblock",
 			"0","",0,4}

 		genesisBlock.Hash=hex.EncodeToString(claculateHash(&genesisBlock))
 		Blockchain=append(Blockchain,genesisBlock)
 		return genesisBlock

}
//计算区块的hash值
func claculateHash(block *Block)[]byte{
	record:=strconv.Itoa(block.Index)+strconv.Itoa(block.Nonce)+strconv.Itoa(block.Difficulty)+
		    block.Data+block.Timestamp+block.Prehash
	h:=sha256.New()
	h.Write([]byte(record))
	hashed:=h.Sum(nil)
	return hashed

}

func (Block)GeneraNextBlock(oldBlock Block,data string)Block{
	var newBlock Block=Block{oldBlock.Index+1,
		time.Now().String(),
		data,oldBlock.Hash,"",0,4}
	for{
		hash:=hex.EncodeToString(claculateHash(&newBlock))
		if PoW(hash,newBlock.Nonce){
			newBlock.Hash=hash
			fmt.Println("挖矿成功")
			Blockchain=append(Blockchain,newBlock)
			return newBlock
		}else{
			newBlock.Nonce++
		}

	}


}
//sha256(block+nonce)=oooo 目标难度值0000
func PoW(data string,nonce int)bool{
	record:=data+strconv.Itoa(nonce)
	h:=sha256.New()
	h.Write([]byte(record))
	hashed:=h.Sum(nil)
	hashed1:=hex.EncodeToString(hashed)
	fmt.Println("挖矿中",hashed1)
	return  strings.HasPrefix(hashed1,"0000")
}

func main(){
var gen=genesisBlock()
gen.GeneraNextBlock(gen,"第二个")
}
