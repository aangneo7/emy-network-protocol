package main

import (
	"fmt"
	"github.com/emy-network/core/node"
)

func main() {
	fmt.Println("Starting Emy Network Node...")
	
	// Initialize Emy Protocol Config
	config := node.DefaultConfig()
	
	// Start the decentralized engine
	emyNode := node.New(config)
	emyNode.Start()
}
