package blockchain

type Chain struct {
	genesis_block *block
	curr_block    *block
}

func CreateBlockChain() *Chain {
	var genesisBlock = CreateGenesisBlock()
	var b = Chain{
		genesis_block: genesisBlock,
		curr_block:    genesisBlock,
	}
	return &b
}

func (b *Chain) LastBlock() *block {
	return b.curr_block
}

func (b *Chain) AddBlock(data SHAable) error {
	var lastBlock = b.LastBlock()
	var prevHash, err = lastBlock.Hash()
	if err != nil {
		return err
	}
	var index = lastBlock.Header.Index + 1
	var block = CreateBlock(index, prevHash, data, lastBlock)
	b.curr_block = block
	return nil
}

func (b *Chain) GenesisBlock() *block {
	return b.genesis_block
}

func (b *Chain) Validate() bool {
	var currBlock = b.LastBlock()
	var validation = true
	for currBlock.PrevBlock != nil {
		validation = validation && currBlock.Validate(currBlock.PrevBlock)
		currBlock = currBlock.PrevBlock
	}
	return validation
}
