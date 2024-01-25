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



func (b BlockChain) IsValid() bool  {
    if b.Blocks[0] != GenerateGenesisBlock() {
        return false
    }

    for i := 1; i < len(b.Blocks); i += 1 {
        currentBlock := b.Blocks[i]
        parentBlock := b.Blocks[i - 1]

        if parentBlock.Hash != currentBlock.LastHash {
            return false
        }


        currentHash := GenerateHash(
            currentBlock.Timestamp,
            currentBlock.Data,
            currentBlock.LastHash,
            currentBlock.Nonce,
        )

        if currentHash != currentBlock.Hash {
            return false

        }

    }



    return true
}

func (b *BlockChain) Replace(newBlockChain BlockChain) bool {
    if len(b.Blocks) >= len(newBlockChain.Blocks) {
        return false
    }

    if newBlockChain.IsValid() == false {
        return false
    }

    b.Blocks = newBlockChain.Blocks

    return true
}
