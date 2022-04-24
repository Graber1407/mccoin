package main

import (
	bc "github.com/Graber1407/mccoin/blockchain"
)

func main() {
	chain := bc.InitBlockChain()

	chain.AddBlock("John sent 1 mccoin to Alice")
	chain.PrintChain()
}
