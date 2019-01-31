package main

import (
	"crypto/sha256"
	"os"
	"time"
)
import (
	"bytes"
	"fmt"
	"strconv"
)

/*
区块结构
*/
type Block struct {
	TimeStamp     int64  //时间戳
	Data          []byte //当前区块 ，存放比特币、账单信息等
	PrevBlockHash []byte //上一个区块加密的hash
	Hash          []byte //当前区块的hash
}

func (this *Block) SetHash() {
	//将本区块的TimeStamp+Data+PrevBlockHash --->Hash
	//将时间戳由整型转换为二进制
	timestamp := []byte(strconv.FormatInt(this.TimeStamp, 10))
	//将三个二进制进行拼接
	headers := bytes.Join([][]byte{this.PrevBlockHash, this.Data, timestamp}, []byte{}) //[]byte{}代表无缝连接

	//将拼接之后的headers进行SHA256加密
	hash := sha256.Sum256(headers)
	this.Hash = hash[:]
}

/*
新建一个区块的api
*/
func NewBlock(data string, prevblockhash []byte) *Block {
	//生成一个区块
	block := Block{}
	//给当前的区块赋值（创建时间，data，前驱hash）
	block.TimeStamp = time.Now().Unix()
	block.Data = []byte(data)
	block.PrevBlockHash = prevblockhash
	//给当前区块进行hash加密
	block.SetHash()
	return &block
}

/*
定义一个区块链的结构
*/
type BlockChain struct {
	Blocks []*Block //有序的区块链
}

//新建一个创世块
func NewGenesisBlock() *Block {
	genesisBlock := Block{}
	genesisBlock.Data = []byte("Genesis block")
	genesisBlock.PrevBlockHash = []byte{}

	return &genesisBlock
}

//新建一个区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

//将区块添加到区块链中
func (this *BlockChain) AddBlock(data string) {
	//得到新添加的区块的前驱区块
	prevBlock := this.Blocks[len(this.Blocks)-1]
	//根据data 创建一个新的区块
	newBlock := NewBlock(data, prevBlock.Hash)
	//依照前驱区块和新区块，添加到区块链blocks中
	this.Blocks = append(this.Blocks, newBlock)
}
func main() {
	//新建一个区块链
	blockchain := NewBlockChain()

	var cmd string
	for {
		fmt.Println("输入'1'添加信息数据到区块链中")
		fmt.Println("输入'2'遍历当前的区块链都有哪些区块信息")
		fmt.Println("按其他按键就是退出")
		fmt.Scanf("%s", &cmd)
		switch cmd {
		case "1":
			//添加一个区块
			inter := []byte{}
			fmt.Println("请输入区块链的行为数据（要添加保存的数据）")
			os.Stdin.Read(inter)
			//fmt.Scan()
			//fmt.Println("i = ", i)
			//fmt.Printf("len(inter) = %d", len(inter))
			blockchain.AddBlock(string(inter))
		case "2":
			//遍历整个区块链
			for i, block := range blockchain.Blocks {
				fmt.Println("=================")
				fmt.Println("第", i, "个区块的信息：")
				fmt.Printf("prevHash = %x\n", block.PrevBlockHash)
				fmt.Printf("Data = %s\n", block.Data)
				fmt.Printf("Hash = %x\n", block.Hash)
				fmt.Println("=================")
			}
		default:
			//退出程序
			fmt.Println("您已经退出程序")
			return
		}
	}
}
