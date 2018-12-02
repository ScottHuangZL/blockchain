package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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
