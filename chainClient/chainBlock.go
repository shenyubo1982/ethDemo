package chainClient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"reflect"
)

type ChainBlock struct {
	blockNum          big.Int //区块高度
	blockHash         string  //区块hash
	blockDifficulty   uint64  //区块难度
	blockTime         uint64  //时间戳
	blockTransactions int     //区块中包含的交易数量
	chainTransaction  []chainTransaction
}

// NewChainBlock 构造器(根据区块高度）
func NewChainBlock(blockNum big.Int) *ChainBlock {
	chainBlockInstance := new(ChainBlock)
	chainBlockInstance.blockNum = blockNum
	return chainBlockInstance
}

//todo 查询区块信息

func (c *ChainBlock) ChainTransaction() []chainTransaction {
	return c.chainTransaction
}

func (c *ChainBlock) setChainTransaction(chainTransaction []chainTransaction) {
	c.chainTransaction = chainTransaction
}

func (c *ChainBlock) BlockNum() big.Int {
	return c.blockNum
}

func (c *ChainBlock) setBlockNum(blockNum big.Int) {
	c.blockNum = blockNum
}

func (c *ChainBlock) BlockHash() string {
	return c.blockHash
}

func (c *ChainBlock) setBlockHash(blockHash string) {
	c.blockHash = blockHash
}

func (c *ChainBlock) BlockDifficulty() uint64 {
	return c.blockDifficulty
}

func (c *ChainBlock) setBlockDifficulty(blockDifficulty uint64) {
	c.blockDifficulty = blockDifficulty
}

func (c *ChainBlock) BlockTime() uint64 {
	return c.blockTime
}

func (c *ChainBlock) setBlockTime(blockTime uint64) {
	c.blockTime = blockTime
}

func (c *ChainBlock) BlockTransactions() int {
	return c.blockTransactions
}

func (c *ChainBlock) setBlockTransactions(blockTransactions int) {
	c.blockTransactions = blockTransactions
}

func getReceiptStatus(client *ethclient.Client, txHex common.Hash, TypeName string) {
	receipt, err := client.TransactionReceipt(context.Background(), txHex)
	if err != nil {
		log.Fatal(err)
	}
	typeOfReceipt := reflect.TypeOf(receipt)
	if reflect.TypeOf(receipt).Name() == TypeName {
		typeOfReceipt.FieldByName(TypeName)
		method, _ := typeOfReceipt.MethodByName("AddInfo")
		fmt.Println(method.Name)

	}

	fmt.Println(receipt.Status) // 1
	fmt.Println(receipt.Logs)   // ...
}
