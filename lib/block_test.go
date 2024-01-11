package lib

import (
    "testing"
)

func TestBlock(t *testing.T) {
    data := "data"
    firstBlock := GenerateGenesisBlock()
    block := MineBlock(firstBlock, data)

    t.Run("lastHash matching", func(t *testing.T) {
        firstBlockHash := firstBlock.Hash
        lastHash := block.LastHash

        if firstBlockHash != lastHash {
            t.Errorf("Expected %s, got %s",firstBlockHash, lastHash )
        }


    })

}
