# Keyfile.hex privatekey extractor or Keyfile generator
This program was made for [Vertcoin OCM](https://github.com/vertcoin-project/one-click-miner-vnext).
Some users complained about not being able to mine to their own address while using OCM, and would rather have the coins being sent directly into their controlled wallet.
This program will allow the user to either extract the private keys (Compressed WIF format) from a keyfile.hex and import them into their wallet of choice, or generate their own keyfile with a previously generated private key.

Use this tool at your own risk, if you lose your keyfile.hex, and private keys you can NOT recover any coins there may be associated with that key.
You can run this in an offline environment if you choose to do so.

## Private key extractor
Extract the private key from your already generated keyfile.hex.
To use the extractor you need to have the keyfile.hex inside the folder from where you run the program. 
Once launching the program it will prompt you for a password, this is the same password your keyfile is encrypted with, without it you can not extract the private key.
Once you have entered the password the program will output your private key in a compressed WIF format which you can then import into your wallet of choice.
