#Readme
#Golang Block Chain (A simple implementation)

**Need Go version 1.11 at least since we use go mod function,**

**Better with go1.12 beta since it can handle golang.org/x related lib smoothly by replace as github.com/golang** 

**20181130 : Add the basic blockchain**

**20181201 : Add wallet function**

**For example:**

For tag 0.21:
- blockchain createblockchain -address "ScottHuang"
- blockchain printchain
- blockchain getbalance -address "ScottHuang"
- blockchain send -from "ScottHuang" -to "CMY" -amount 50
- blockchain send -from "ScottHuang" -to "HJH" -amount 49
- blockchain printchain
- blockchain getbalance -address "ScottHuang"

For tag0.31 add below very basic wallet, still not integrate into block chain, will do it later:
- blockchain createwallet  (You can try run multiple times)
- blockchain listaddresses (List all addresses which created before)

For tag0.41 add signature, you can delete tmp/wallets.data & tmp/blocks/*.* firstly
- blockchain createwallet (You can run this multiple times)
- blockchain listaddresses (List all addresses which created before)
- blockchain send -from "The address list in above" -to "another address list in above" -amount 50
- blockchain printchain
- blockchain getbalance -address "The address list in above"

For tag 0.5 Adding the unspent transaction outputs(UTXO) persistence layer
- You can delete ./tmp/block/*.* & ./tmp/wallet.data to clear the chain and wallet
- blockchain createwallet (You can run this multiple times)
- blockchain createblockchain -address 1G41xLzHMtKnozJqaL8RBbd44MqFVEETTD(This is the address which you just created at above)
- blockchain send -from 189iX92pxDWaFiii7FLpcWfFsm2adkdvfA -to 174Xcjx6ePrzVVj28mHJbQPh6AkQ89ehpq -amount 20
- blockchain listaddresses
- blockchain printchain
- You can try all blockchain commands by yourself

Note: I get ideas from TensorFlow, thanks.