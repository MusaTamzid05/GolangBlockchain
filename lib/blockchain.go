package lib

type BlockChain struct {
    Blocks []Block
}

func MakeBlockChain() BlockChain {
    blockChain := BlockChain{}
    blockChain.Blocks = append(blockChain.Blocks, GenerateGenesisBlock() )

    return blockChain

}


