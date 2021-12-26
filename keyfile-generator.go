package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
)

func main() {
	err := generate()
	if err != nil {
		fmt.Println("There was an error generating the keyfile.. exiting", err)
		panic(err)
	}
	successful := KeyFileValid()
	if !successful {
		fmt.Println("Keyfile has wrong length and is NOT valid!")
	} else {
		fmt.Println("Keyfile was successfully generated")
	}
}

func generate() error {
	filename := "keyfile.hex"

	// Get priv key
	var privKey string
	fmt.Println("Input your privatekey:")
	fmt.Scanln(&privKey)

	// Decode priv key into an array of bytes
	privWIF, err := btcutil.DecodeWIF(privKey)
	if err != nil {
		fmt.Println("error decoding wif")
		return err
	}
	// We need to do a bit of conversion between types
	privECDSA := privWIF.PrivKey.ToECDSA()
	priv32 := privECDSA.D.Bytes()

	// Get password for keyfile
	var pass string
	fmt.Println("Create a password. If you loose the privatekey and forget the password you won't be able to recover the coins!")
	fmt.Println("Password:")
	fmt.Scanln(&pass)

	// Derive pubkey
	_, pub := btcec.PrivKeyFromBytes(btcec.S256(), priv32[:])

	salt := new([24]byte) // salt for scrypt / nonce for secretbox
	dk32 := new([32]byte) // derived key from scrypt

	//get 24 random bytes for scrypt salt (and secretbox nonce)
	_, err = rand.Read(salt[:])
	if err != nil {
		return err
	}
	// next use the pass and salt to make a 32-byte derived key
	dk, err := scrypt.Key([]byte(pass), salt[:], 16384, 8, 1, 32)
	if err != nil {
		return err
	}
	copy(dk32[:], dk[:])

	enckey := append(salt[:], secretbox.Seal(nil, priv32[:], salt, dk32)...)
	return ioutil.WriteFile(filename, append(pub.SerializeCompressed(), enckey...), 0600)
}

func loadPublicKey() []byte {
	filename := "keyFile.hex"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading keyfile: %s", err.Error())
		return []byte{}
	}
	if len(b) != 105 {
		fmt.Printf("Error: Keyfile is not valid! Keyfile had wrong length. Expected 129, got %d", len(b))
		return []byte{}
	}
	ret := make([]byte, 33)
	copy(ret, b[:33])
	b = nil
	return ret
}

// KeyFileValid returns true if there is a valid initialized keyfile available
func KeyFileValid() bool {
	return len(loadPublicKey()) == 33
}
