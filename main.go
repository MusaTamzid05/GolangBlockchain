package main

import  (
    "musa.io/simple_blockchain/lib"
    "fmt"
)


func main() {
    firstBlock := lib.GenerateGenesisBlock()
    block := lib.MineBlock(firstBlock, "some data")

    fmt.Println(block.String())

}
