package main

import  (
    "musa.io/simple_blockchain/lib"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

type RequestData struct {
    Data string `json:"data"`
}


func init() {
    lib.CurrentBlockchain = lib.MakeBlockChain()
}



func GetBlockChainHandler(c *gin.Context) {
    blocks := lib.CurrentBlockchain.Blocks
    c.IndentedJSON(http.StatusOK, blocks)
}

func MineBlockHandler(c *gin.Context) {
    var requestData RequestData
    err := c.BindJSON(&requestData)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H {
            "success" : false,
            "message" : err.Error(),

        })

        return
    }

    lib.CurrentBlockchain.Add(requestData.Data)

    c.JSON(http.StatusCreated, gin.H {
        "success" : true,
        "message" : "",

    })
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

    var err error
    lib.CurrentP2P, err = lib.MakeP2P(":5001")

    if err != nil {
        fmt.Println("Error running P2P Server ", err.Error())
        os.Exit(1)
    }

    go lib.CurrentP2P.StartServer()



    router := gin.Default()
    router.GET("/blocks", GetBlockChainHandler)
    router.POST("/mine", MineBlockHandler)
    router.Run()



}
