package main

import  (
    "musa.io/simple_blockchain/lib"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "flag"
    "strings"
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
    hostAddr := flag.String("host", "", "http server")
    p2pServer := flag.String("p2pServer", "", "p2p server")
    peerAddrs := flag.String("peerAddrs", "", "addr of peers")

    flag.Parse()

    if *hostAddr == "" {
        fmt.Println("Please give a host addr for http server Ex -host :8080")
        os.Exit(1)
    }


    if *p2pServer == "" {
        fmt.Println("Please give a host addr for p2p server Ex -p2pServer :5001")
        os.Exit(1)
    }


    var err error
    lib.CurrentP2P, err = lib.MakeP2P(*p2pServer)

    if err != nil {
        fmt.Println("Error running P2P Server ", err.Error())
        os.Exit(1)
    }

    go lib.CurrentP2P.StartServer()

    if *peerAddrs != "" {
        peerLists := strings.Split(*peerAddrs, ",")
        go lib.CurrentP2P.AddPeers(peerLists)

    }




    router := gin.Default()
    router.GET("/blocks", GetBlockChainHandler)
    router.POST("/mine", MineBlockHandler)
    router.Run(*hostAddr)



}
