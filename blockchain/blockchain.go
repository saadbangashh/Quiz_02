package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	ID        int
	Timestamp string
	Data      string
	Hash      string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock(0, "Genesis Block")
}

func NewBlock(id int, data string) *Block {
	block := &Block{
		ID:        id,
		Timestamp: time.Now().String(),
		Data:      data,
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() string {
	record := fmt.Sprintf("%d%s%s", b.ID, b.Timestamp, b.Data)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (bc *Blockchain) AddBlock(data string) { // Modify block for clarity
	newBlock := NewBlock(len(bc.Blocks), data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) DisplayAllBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Block ID: %d\nTimestamp: %s\nData: %s\nHash: %s\n\n", block.ID, block.Timestamp, block.Data, block.Hash)
	}
}
