// main.go
package main

import (
	"fmt"
	
)

// Block represents a data structure for your blocks
type Block struct {
	ID   int
	Data string
}

// BlockDatabase is a simple in-memory database for blocks
var BlockDatabase []Block

// DisplayAllBlocks prints all the blocks in the database
func DisplayAllBlocks() {
	fmt.Println("All Blocks:")
	for _, block := range BlockDatabase {
		fmt.Printf("ID: %d, Data: %s\n", block.ID, block.Data)
	}
	fmt.Println()
}

// NewBlock creates a new block and adds it to the database
func NewBlock(data string) {
	newID := len(BlockDatabase) + 1
	newBlock := Block{ID: newID, Data: data}
	BlockDatabase = append(BlockDatabase, newBlock)
	fmt.Printf("New Block Created - ID: %d, Data: %s\n\n", newID, data)
}

// ModifyBlock updates the data of an existing block in the database
func ModifyBlock(id int, newData string) {
	for i := range BlockDatabase {
		if BlockDatabase[i].ID == id {
			BlockDatabase[i].Data = newData
			fmt.Printf("Block Modified - ID: %d, New Data: %s\n\n", id, newData)
			return
		}
	}
	fmt.Printf("Block with ID %d not found\n\n", id)
}

func main() {
	// Example usage
	NewBlock("Block 1 Data")
	NewBlock("Block 2 Data")
	DisplayAllBlocks()

	ModifyBlock(1, "Updated Block 1 Data")
	DisplayAllBlocks()
}

