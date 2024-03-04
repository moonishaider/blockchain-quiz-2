// main.go
package main

import (
	"fmt"
	
)

type Block struct {
	ID   int
	Data string
}

var BlockDatabase []Block

func DisplayAllBlocks() {
	fmt.Println("All Blocks:")
	for _, block := range BlockDatabase {
		fmt.Printf("ID: %d, Data: %s\n", block.ID, block.Data)
	}
	fmt.Println()
}

func NewBlock(data string) {
	newID := len(BlockDatabase) + 1
	newBlock := Block{ID: newID, Data: data}
	BlockDatabase = append(BlockDatabase, newBlock)
	fmt.Printf("New Block Created - ID: %d, Data: %s\n\n", newID, data)
}


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
	
	NewBlock("Block 1 Data")
	NewBlock("Block 2 Data")
	DisplayAllBlocks()

	ModifyBlock(1, "Updated Block 1 Data")
	DisplayAllBlocks()
}

