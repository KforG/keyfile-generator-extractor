package main

import (
	"fmt"
	"io/ioutil"
	"syscall"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	LoadPrivateKey()
}

func getPass() (pass []byte) {
	fmt.Println("Type the password to your keyfile.hex:")
	pass, err := terminal.ReadPassword(int(syscall.Stdin)) //Hides password from the terminal
	if err != nil {
		fmt.Println("Something went wrong getting the password")
	}
	return pass
}

//Decryption part of this code has been taken from https://github.com/vertcoin-project/one-click-miner-vnext
func LoadPrivateKey() {
	filename := "keyFile.hex"
	password := getPass()

	keyfile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not find a %s in the same folder as the extractor\n", filename)
		panic(err)
	}
	if len(keyfile) != 105 {
		fmt.Printf("Key length error for %s. Your file may be incorrectly made or corrupted!\n", filename)
		panic(err)
	}

	enckey := keyfile[33:]
	// enckey is actually encrypted, get derived key from pass and salt
	// first extract salt
	salt := new([24]byte)      // salt (also nonce for secretbox)
	dk32 := new([32]byte)      // derived key array
	copy(salt[:], enckey[:24]) // first 24 bytes are scrypt salt/box nonce

	dk, err := scrypt.Key([]byte(password), salt[:], 16384, 8, 1, 32) // derive key
	if err != nil {
		fmt.Println(err)
	}
	copy(dk32[:], dk[:]) // copy into fixed size array

	// nonce for secretbox is the same as scrypt salt.  Seems fine.  Really.
	priv, worked := secretbox.Open(nil, enckey[24:], salt, dk32)
	if worked != true {
		fmt.Printf("Decryption failed for %s! Make sure to use the correct password\n", filename)
		main()
	} else {
		convert(priv)
		fmt.Println("Press enter to exit")
		fmt.Scanln()

	}
}

//This function converts the dectrypted data from the keyfile into a private key in a compressed WIF format
func convert(priv []byte) {
	decPriv, _ := btcec.PrivKeyFromBytes(btcec.S256(), priv)
	WIF, err := btcutil.NewWIF(decPriv, &chaincfg.MainNetParams, true)
	if err != nil {
		fmt.Printf("Error reading the private key from bytes\n")
		panic(err)
	}
	fmt.Printf("Private key in compressed WIF format: %s\n", WIF)
	ShowQR(WIF)
}

func ShowQR(WIF *btcutil.WIF) {
	content := fmt.Sprintf("%s", WIF)
	qr := qrcodeTerminal.New()
	qr.Get(content).Print()
}
