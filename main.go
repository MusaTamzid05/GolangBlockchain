package main

import  (
    "musa.io/simple_blockchain/lib"
    "fmt"
)


func main() {
    blockChain := lib.MakeBlockChain()
    blockChain.Add("data2")
    blockChain.Add("data3")


    blockChain.Show()
    bc2 := lib.MakeBlockChain()
    bc2.Add("data2")
    bc2.Add("data3")
    bc2.Add("data4")


    fmt.Println(blockChain.Replace(bc2))

    blockChain.Show()



}
