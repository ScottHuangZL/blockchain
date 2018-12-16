package blockchain

import "github.com/dgraph-io/badger"

//Iterator type
type Iterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

//Iterator method for blockchain
func (chain *BlockChain) Iterator() *Iterator {
	iter := &Iterator{chain.LastHash, chain.Database}

	return iter
}

//Next for iterator
func (iter *Iterator) Next() *Block {
	var block *Block
	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)
		encodeBlock, err := item.Value()
		block = Deserialize(encodeBlock)

		return err
	})
	Handle(err)
	iter.CurrentHash = block.PrevHash
	return block
}
