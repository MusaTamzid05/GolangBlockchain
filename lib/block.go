package lib

import (
    "time"
    "crypto/sha256"
    "fmt"

)


type Block struct {
    Timestamp int
    Data string
    Hash string
    LastHash string
    Nonce int
    Difficulty int
}

func (b Block) String() string  {
    str := "\n\tBlock\n"
    str += fmt.Sprintf("%s%d\n", str, b.Timestamp)
    str += "Data : " + b.Data+ "\n"
    str += "Hash : " + b.Hash+ "\n"
    str += "Last Hash: " + b.LastHash+ "\n"
    str = fmt.Sprintf("%s\nNonce: %d", str, b.Nonce)
    str = fmt.Sprintf("%s\nDifficulty: %d", str, b.Difficulty)

    return str

}

func MakeBlock(timestamp int, data, lastHash string, nonce, difficulty int, firstBlock bool) Block {

    // 000

    hash := ""

    if firstBlock {
        hash = "First Hash"

    } else {
        //timestamp = strconv.Itoa(int(time.Now().Unix()))
        hash = GenerateHash(timestamp, data, lastHash, nonce, difficulty)

    }
    

    return Block{
        Timestamp : timestamp,
        Data: data,
        Hash: hash,
        LastHash: lastHash,
        Nonce: nonce,
        Difficulty: difficulty,

    }
}

func GenerateGenesisBlock() Block {
    return MakeBlock(0, "Genesis Block", "Last Hash", 0, DEFAULT_DIFFICULTY, true)
}


func MineBlock(lastBlock Block, data string) Block {

    // Find out the difficulty
    // if Find out the time to find difficulty
    // if difficulty > min_time = difficulty - 1
    // else difficulty += 1
    // generate the block with this difficulty

    lastHash := lastBlock.Hash
    timestamp := int(time.Now().Unix())
    difficulty := lastBlock.Difficulty

    firstZeroes := ""
    nonce := 0

    for i := 0; i < difficulty; i += 1 {
        firstZeroes += "0"
    }


    hashSloved := false

    for hashSloved == false {
        hash := GenerateHash(timestamp, data, lastHash, nonce, difficulty)

        if hash[:difficulty] == firstZeroes {
            hashSloved = true
            continue
        }

        nonce += 1
    }

    mininingTime := int(time.Now().Unix()) - timestamp

    if mininingTime > MINING_TIME_SECONDS {
        difficulty -= 1
        fmt.Println("Difficulty decrease")
    } else {

        difficulty += 1
        fmt.Println("Difficulty increase")
    }

    firstZeroes = ""
    nonce = 0

    for i := 0; i < difficulty; i += 1 {
        firstZeroes += "0"
    }

    var block Block

    blockGenerated := false

    for blockGenerated == false {
        block = MakeBlock(timestamp, data, lastHash, nonce, difficulty,  false)

        hash := block.Hash

        if hash[:difficulty] == firstZeroes {
            blockGenerated = true
            continue
            
        }

        nonce += 1
    }



    return block
}

func GenerateHash(timestamp int, data, lastHash string, nonce, difficulty int) string {
    //hashData := timestamp + data +  lastHash
    hashData := fmt.Sprintf("%d%s%s%d%d", timestamp, data, lastHash, nonce, difficulty)

    hashGenerator := sha256.New()
    hashGenerator.Write([]byte(hashData))

    result := hashGenerator.Sum(nil)
    return fmt.Sprintf("%x", result)


}
