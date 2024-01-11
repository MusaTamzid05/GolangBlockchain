package main

import  (
    "musa.io/simple_blockchain/lib"
    "fmt"
)


func main() {
    firstBlock := lib.GenerateGenesisBlock()
    fmt.Println(firstBlock)
    block := lib.MakeBlock("timestamp", "data", "hash", "last hash")
    fmt.Println(block.String())

}
