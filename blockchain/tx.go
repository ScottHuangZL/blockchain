package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"github.com/ScottHuangZL/blockchain/wallet"
	//"../wallet"
)

//TxOutput type
type TxOutput struct {
	Value      int    //the value in token is signed and locked
	PubKeyHash []byte //a value that is needed to unlock the token which are inside the value field
}

//TxInput type
type TxInput struct {
	ID        []byte //references the transaction that the output is inside it
	Out       int    //index where the output appeared in that transaction
	Signature []byte //quite similar the pub key in output
	PubKey    []byte
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

func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublishKeyHash(in.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-wallet.ChecksumLength]
	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))
	return txo
}
