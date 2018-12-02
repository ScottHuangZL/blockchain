package wallet

import (
	"github.com/ScottHuangZL/blockchain/blockchain"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4
	version = byte(0x00)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	blockchain.Handle(err)

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func MakeWallet() *Wallet{
	private, public := NewKeyPair()
	wallet := Wallet{private, public}


	return &wallet
}

func PublishKeyHash(pubKey []byte) []byte{
	pubHash := sha256.Sum256(pubKey)
	hasher := ripemd160.New()
	_,err:=hasher.Write(pubHash[:])
	blockchain.Handle(err)
	publicRipMD := hasher.Sum(nil)
	return publicRipMD
}

func Checksum(payload []byte) []byte{
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}


func (w Wallet) Address() []byte{
	pubHash :=PublishKeyHash(w.PublicKey)

	versionedHash := append([]byte{version}, pubHash...)
	checkSum := Checksum(versionedHash)
	fullHash := append(versionedHash,checkSum...)
	address := Base58Encode(fullHash)

	fmt.Printf("Pub key: %x\n", w.PublicKey)
	fmt.Printf("Pub hash: %x\n", pubHash)
	fmt.Printf("Address: %x\n",address)

	return address
}