package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

//Take data from the block

//create counter(nounce) which start at 0

//create a hash of the data plus the counter

//check the hash to see if it meets a set of requirements

//Requirements:
//The first several bytes must contain 0s

//Difficulty of the proof calculation
const Difficulty = 12

//ProofOfWork type
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

//NewProof func
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//	fmt.Printf("%x  | ", target)
	target.Lsh(target, uint(256-Difficulty))
	//	fmt.Println("After Lsh %x\n", target)
	pow := &ProofOfWork{b, target}

	return pow
}

//InitData method
func (pow *ProofOfWork) InitData(nounce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			ToHex(int64(nounce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

//ToHex func
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

//Run pow
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nounce := 0

	for nounce < math.MaxInt64 {
		data := pow.InitData(nounce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nounce++
		}
	}
	fmt.Println()
	return nounce, hash[:]
}

//Validate pow
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nounce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}
