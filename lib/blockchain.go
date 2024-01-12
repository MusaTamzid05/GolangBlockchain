package lib

import (
    "fmt"
)

type BlockChain struct {
    Blocks []Block
}

func MakeBlockChain() BlockChain {
    blockChain := BlockChain{}
    blockChain.Blocks = append(blockChain.Blocks, GenerateGenesisBlock() )
    return blockChain
}

func (b *BlockChain) Add(data string) Block  {
    lastBlock := b.Blocks[len(b.Blocks) - 1]
    newBlock  := MineBlock(lastBlock, data)

    b.Blocks = append(b.Blocks, newBlock)

    return newBlock

}

func (b BlockChain) Show() {
    for _, block := range b.Blocks {
        fmt.Println(block.String())
    }

}


