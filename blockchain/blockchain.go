package blockchain

import (
	"encoding/hex"
	"fmt"
	"os"
	"runtime"

	"github.com/dgraph-io/badger"
)

const (
	dbPath      = "./tmp/blocks"
	dbFile      = "./tmp/blocks/MANIFEST"
	genesisData = "First Transaction from Genesis"
)

//BlockChain type
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

//Iterator type
type Iterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

//DBexists func
func DBexists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}

//InitBlockChain func
func InitBlockChain(address string) *BlockChain {
	var lastHash []byte

	if DBexists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		cbtx := CoinbaseTx(address, genesisData)
		genesis := Genesis(cbtx)
		fmt.Println("Genesis Created")
		err = txn.Set(genesis.Hash, genesis.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), genesis.Hash)
		lastHash = genesis.Hash
		return err
	})

	Handle(err)
	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

//ContinueBlockChain func
func ContinueBlockChain(address string) *BlockChain {

	if DBexists() == false {
		fmt.Println("No existing Blockchain found, create one!")
		runtime.Goexit()
	}

	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()

		return err

	})

	Handle(err)
	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

//AddBlock method
func (chain *BlockChain) AddBlock(transactions []*Transaction) {
	var lasthHash []byte
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lasthHash, err = item.Value()

		return err
	})
	Handle(err)
	newBlock := CreateBlock(transactions, lasthHash)
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)
		chain.LastHash = newBlock.Hash
		return err
	})
	Handle(err)
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

//FindUnspentTransactions method
func (chain *BlockChain) FindUnspentTransactions(address string) []Transaction {
	var unspentTxs []Transaction        //unspent transactions are transactions that have outputs which not references by other inputs.
	spentTXOs := make(map[string][]int) //spent transaction outputs
	iter := chain.Iterator()
	for {
		block := iter.Next()

		//for each block transaction
		for _, tx := range block.Transactions {
			//convert the transaction id to string
			txIdString := hex.EncodeToString(tx.ID)
		Outputs:
			//each transaction have its own outputs array
			for outIdx, out := range tx.Outputs {
				//have this transactions spent records
				if spentTXOs[txIdString] != nil {
					//the spend out slice
					for _, spentOut := range spentTXOs[txIdString] {
						if spentOut == outIdx {
							continue Outputs //ignore if this outputs been referenced by other inputs
						}
					}
				}
				if out.CanBeUnlocked(address) {
					unspentTxs = append(unspentTxs, *tx)
				}
			}
			if tx.IsCoinbase() == false {
				//each block also have its own inputs, need range too
				for _, in := range tx.Inputs {
					if in.CanUnlock(address) { //it mean already spent if the input can unlock by this address
						inTxIdString := hex.EncodeToString(in.ID)
						spentTXOs[inTxIdString] = append(spentTXOs[inTxIdString], in.Out)
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}
	return unspentTxs
}

//FindUTXO method  find unspent transaction output
func (chain *BlockChain) FindUTXO(address string) []TxOutput {
	var UTXOs []TxOutput
	unspentTransactions := chain.FindUnspentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Outputs {
			if out.CanBeUnlocked(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}
	return UTXOs
}

//FindSpendableOutputs method
func (chain *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	unspentTxs := chain.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txIdString := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Outputs {
			if out.CanBeUnlocked(address) && accumulated < amount {
				accumulated += out.Value
				unspentOuts[txIdString] = append(unspentOuts[txIdString], outIdx)
				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOuts
}
