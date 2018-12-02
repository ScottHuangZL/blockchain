package blockchain

import (
	"encoding/hex"
	"fmt"
	"log"
)



//Transaction type
type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

//CoinbaseTx func
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}
	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}

	tx.SetID()

	return &tx
}


//NewTransaction func
func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)
	if acc < amount {
		log.Panic("Error: not enough founds")
	}

	for txId, outs := range validOutputs {
		txIdString, err := hex.DecodeString(txId)
		Handle(err)

		//each outs also is an array
		for _, out := range outs {
			input := TxInput{txIdString, out, from}
			//create inputs which contains all the unspent outputs in the transaction
			inputs = append(inputs, input)
		}
	}

	//this is the outputs to toAddress
	outputs = append(outputs, TxOutput{amount, to})


	if acc > amount {
		//deduct the from's amount, this another outputs
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
