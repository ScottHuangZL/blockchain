package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

//Block type
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nounce   int
}

// //DeriveHash method
// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

//CreateBlock func
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nounce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nounce = nounce
	return block
}

//Genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//Serialize method
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}

//Deserialize func
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)

	return &block
}

//Handle error func
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
