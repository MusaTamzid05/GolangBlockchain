package main

import  (
    "musa.io/simple_blockchain/lib"
    //"fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

var chain lib.BlockChain

func init() {
    chain = lib.MakeBlockChain()
}



func GetBlockChainHandler(c *gin.Context) {
    blocks := chain.Blocks
    c.IndentedJSON(http.StatusOK, blocks)



}

func main() {
    /*
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
    */


    router := gin.Default()
    router.GET("/blocks", GetBlockChainHandler)
    router.Run()



}
