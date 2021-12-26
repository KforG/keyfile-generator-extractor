# Keyfile.hex privatekey extractor or Keyfile generator
This program was made for [Vertcoin OCM](https://github.com/vertcoin-project/one-click-miner-vnext).
Some users complained about not being able to mine to their own address while using OCM, and would rather have the coins being sent directly into their controlled wallet.

This program will allow the user to either extract the private keys (Compressed WIF format) from a keyfile.hex and import them into their wallet of choice, or generate their own keyfile with a previously generated private key.

Use this tool at your own risk, if you lose your keyfile.hex, and private keys you can NOT recover any coins there may be associated with that key.
You can run this in an offline environment if you choose to do so.

## Private key extractor
Extract the private key from your already generated keyfile.hex.

To use the extractor you need to have a keyfile.hex inside the directory from where you run the program. 

After launching the program it will prompt you for a password, this is the same password your keyfile is encrypted with, without it you can not extract the private key.

Once you have entered the password the program will output your private key in a compressed WIF format which you can then import into your wallet of choice. It will also print a QR code to make it easy to import into wallets like Coinomi.

## Keyfile generator
Generate your own keyfile with an existing private key.

To use this generator you need to have a private key in a WIF format.

After launching the program it will prompt you for the private key and then the desired password. This will be the same password that your new keyfile will be generated in, if you forget the password you can not recover the private key from the keyfile or use the keyfile in any other way.

Shortly after a new keyfile.hex will appear in the same directory as the program was run from, and the program will output a message saying the generation was successful or unsuccessful. If the latter is true DO NOT use the keyfile.

## How to build
I strongly recommend people to build this themselves.

You need to have [Go](https://golang.org/) installed in order to be able to build this.

When Go is installed and the program downloaded you need to download the used packages:
```bash
go mod download
```
Once finished you can build the program:
```bash
go build extract-privatekey.go
```
or:
```bash
go build keyfile-generator.go
```

## Donate
If you want to support the development of this program feel free to donate!

Vertcoin: `Vn9FSJ2WiEWcsuhvC2bGK4ZGaqcfMLZ4VC`