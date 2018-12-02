package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

//TxOutput type
type TxOutput struct {
	Value  int //the value in token is signed and locked
	PubKey string //a value that is needed to unlock the token which are inside the value field
}

//TxInput type
type TxInput struct {
	ID  []byte //references the transaction that the output is inside it
	Out int    //index where the output appeared in that transaction
	Sig string //quite similar the pub key in output
}

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

//SetID method
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]

}

//IsCoinbase method
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

//CanUnlock method
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

//CanBeUnlocked method
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
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
