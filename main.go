package main

import  (
    "musa.io/simple_blockchain/lib"
    "fmt"
)


func main() {
    blockChain := lib.MakeBlockChain()
    fmt.Println(blockChain.Blocks[0].String())


}
