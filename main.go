package main

import  (
    "musa.io/simple_blockchain/lib"
    "fmt"
)


func main() {
    blockChain := lib.MakeBlockChain()
    blockChain.Add("data2")
    blockChain.Add("data3")
    blockChain.Add("data4")


    blockChain.Show()

    fmt.Println(blockChain.IsValid())



}
