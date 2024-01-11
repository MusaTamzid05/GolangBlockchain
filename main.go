package main

import  (
    "musa.io/simple_blockchain/lib"
    "fmt"
)


func main() {
    block := lib.MakeBlock("timestamp", "data", "hash", "last hash")
    fmt.Println(block.String())

}
