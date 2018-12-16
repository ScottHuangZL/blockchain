#Readme
#Golang Block Chain (A simple implementation)

**Need Go version 1.11 at least since we use go mod function,**

**Better with go1.12 beta since it can handle golang.org/x related lib smoothly by replace as github.com/golang** 

**20181130 : Add the basic blockchain**

**20181201 : Add wallet function**

**For example:**
- How to build:
    - go mod tidy
    - go mod vendor
    - go build -mod=vendor


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


For tag 0.6 Adding the network
- open 3 command window, execute below 3 command
    - set NODE_ID=3000
    - set NODE_ID=4000
    - set NODE_ID=5000
- In each command window, also run 
    - blockchain createwallet , it will show address, such as 1BpZRTTK4kGSShLTnJDyjwHGgy9ntYaYxn
- in the NODE_ID=3000 window, run: 
    - blockchain createblockchain -address 1BpZRTTK4kGSShLTnJDyjwHGgy9ntYaYxn
    - blockchain send -from 1PmiqcScaNkxNyc5xoo2zkKvTkikvagaE8 -to 1Kvgu9HWVeeCw6xXrV1tcAgtWTcajWnYb3 -amount 10 -mine
    - blockchain startnode
- in the NODE_ID=4000 window, also run: blockchain startnode
- in the NODE_ID=5000 window, run: blockchain send -from 1Kvgu9HWVeeCw6xXrV1tcAgtWTcajWnYb3 -to 1Afi42GPnzkgLuAvTtiZt4LdF2nceGkSCt -amount 1
- in the NODE_ID=5000 window, run: blockchain startnode -miner 1Kvgu9HWVeeCw6xXrV1tcAgtWTcajWnYb3

Note: I get ideas from Tensor-Programing, thanks.