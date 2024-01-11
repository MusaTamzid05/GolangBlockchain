package lib

import (
    "time"
    "strconv"
    "crypto/sha256"
    "fmt"

)


type Block struct {
    Timestamp string
    Data string
    Hash string
    LastHash string
}

func (b Block) String() string  {
    str := "\n\tBlock\n"
    str += "Timestamp : " + b.Timestamp + "\n"
    str += "Data : " + b.Data+ "\n"
    str += "Hash : " + b.Hash+ "\n"
    str += "Last Hash: " + b.LastHash+ "\n"

    return str

}

func MakeBlock(data, lastHash string) Block {
    timestamp := strconv.Itoa(int(time.Now().Unix()))
    hash := GenerateHash(timestamp, data, lastHash)

    return Block{
        Timestamp : timestamp,
        Data: data,
        Hash: hash,
        LastHash: lastHash,

    }
}

func GenerateGenesisBlock() Block {
    return MakeBlock("first data", "123")
}


func MineBlock(lastBlock Block, data string) Block {
    lastHash := lastBlock.Hash
    return MakeBlock(data, lastHash)
}

func GenerateHash(timestamp, data, lastHash string) string {
    hashData := timestamp + data +  lastHash

    hashGenerator := sha256.New()
    hashGenerator.Write([]byte(hashData))

    result := hashGenerator.Sum(nil)
    return fmt.Sprintf("%x", result)


}
