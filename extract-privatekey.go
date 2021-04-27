package main

import (
	"fmt"
	"io/ioutil"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
)

func main() {
	var pass string
	fmt.Println("Type the password to your keyfile.hex:")
	fmt.Scanln(&pass)
	priv, err := LoadPrivateKey(pass)
	if err != nil {
		panic(err)
	}
	convert(priv)
}

//The majority of this code has been taken from https://github.com/vertcoin-project/one-click-miner-vnext, some modifications added for this to be able to work
func LoadPrivateKey(password string) ([]byte, error) {
	filename := "keyFile.hex"
	keyfile, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	if len(keyfile) != 105 {
		return []byte{}, fmt.Errorf("Key length error for %s\n", filename)
	}

	enckey := keyfile[33:]
	// enckey is actually encrypted, get derived key from pass and salt
	// first extract salt
	salt := new([24]byte)      // salt (also nonce for secretbox)
	dk32 := new([32]byte)      // derived key array
	copy(salt[:], enckey[:24]) // first 24 bytes are scrypt salt/box nonce

	dk, err := scrypt.Key([]byte(password), salt[:], 16384, 8, 1, 32) // derive key
	if err != nil {
		return []byte{}, err
	}
	copy(dk32[:], dk[:]) // copy into fixed size array

	// nonce for secretbox is the same as scrypt salt.  Seems fine.  Really.
	priv, worked := secretbox.Open(nil, enckey[24:], salt, dk32)
	if worked != true {
		return []byte{}, fmt.Errorf("Decryption failed for %s\n", filename)
	}

	return priv, nil
}

//This function converts the dectrypted data from the keyfile into a private key in compressed WIF format
func convert(priv []byte) {
	decPriv, _ := btcec.PrivKeyFromBytes(btcec.S256(), priv)

	WIF, err := btcutil.NewWIF(decPriv, &chaincfg.MainNetParams, true)
	if err != nil {
		fmt.Errorf("Error reading the private key from bytes\n")
	}
	fmt.Printf("Private key in compressed WIF format: %s\n", WIF)
}
