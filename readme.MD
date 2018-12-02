# blockchain

##Golang Block Chain

**Need Go version 1.11 at least since we use go mod function,**

**Better with go1.11.3 since it can handle golang.org/x related lib smoothly by replace as github.com/golang** 

**20181130 : Add the basic blockchain**

**20181201 : Add wallet function**

**For example:**

For tag 0.2:
- blockchain createblockchain -address "ScottHuang"
- blockchain printchain
- blockchain getbalance -address "ScottHuang"
- blockchain send -from "ScottHuang" -to "CMY" -amount 50
- blockchain send -from "ScottHuang" -to "HJH" -amount 49
- blockchain printchain
- blockchain getbalance -address "ScottHuang"

For tag0.3 add below very basic wallet, still not integrate into block chain, will do it later:
- blockchain createwallet  (You can try run multiple times)
- blockchain listaddresses (List all addresses which created before)
