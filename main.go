package main

import (
	"QUIZ_02/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("First block after Genesis")
	bc.AddBlock("Second block after Genesis")
	bc.AddBlock("Third block after Genesis")

	bc.DisplayAllBlocks()
}
