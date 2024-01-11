package lib


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

func MakeBlock(timestamp, data, hash, lastHash string) Block {
    return Block{
        Timestamp : timestamp,
        Data: data,
        Hash: hash,
        LastHash: lastHash,

    }
}

func GenerateGenesisBlock() Block {
    return MakeBlock("Genesis Block", "first data", "genesis-hash", "123")

}
