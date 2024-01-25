package lib

import (
    "time"
    //"strconv"
    "crypto/sha256"
    "fmt"

)


type Block struct {
    Timestamp int
    Data string
    Hash string
    LastHash string
    Nonce int
}

func (b Block) String() string  {
    str := "\n\tBlock\n"
    str += fmt.Sprintf("%s%d\n", str, b.Timestamp)
    str += "Data : " + b.Data+ "\n"
    str += "Hash : " + b.Hash+ "\n"
    str += "Last Hash: " + b.LastHash+ "\n"

    return str

}

func MakeBlock(timestamp int, data, lastHash string, nonce int, firstBlock bool) Block {

    // 000

    hash := ""

    if firstBlock {
        hash = "First Hash"

    } else {
        //timestamp = strconv.Itoa(int(time.Now().Unix()))
        hash = GenerateHash(timestamp, data, lastHash, nonce)

    }

    
    

    return Block{
        Timestamp : timestamp,
        Data: data,
        Hash: hash,
        LastHash: lastHash,
        Nonce: nonce,

    }
}

func GenerateGenesisBlock() Block {
    return MakeBlock(0, "Genesis Block", "Last Hash", 0, true)
}


func MineBlock(lastBlock Block, data string) Block {
    lastHash := lastBlock.Hash
    timestamp := int(time.Now().Unix())

    firstZeroes := ""
    nonce := 0

    for i := 0; i < DIFFICULTY; i += 1 {
        firstZeroes += "0"
    }


    sloved := false

    for sloved == false {
        hash := GenerateHash(timestamp, data, lastHash, nonce)

        if hash[:DIFFICULTY] == firstZeroes {
            sloved = true
            continue
        }

        nonce += 1

    }

    return MakeBlock(timestamp, data, lastHash, nonce,  false)
}

func GenerateHash(timestamp int, data, lastHash string, nonce int) string {
    //hashData := timestamp + data +  lastHash
    hashData := fmt.Sprintf("%d%s%s%d", timestamp, data, lastHash, nonce)

    hashGenerator := sha256.New()
    hashGenerator.Write([]byte(hashData))

    result := hashGenerator.Sum(nil)
    return fmt.Sprintf("%x", result)


}
