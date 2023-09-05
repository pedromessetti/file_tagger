package main

import (
	"os"

	"github.com/pedromessetti/tag_file/gen"
	"github.com/pedromessetti/tag_file/tag"

)

func main() {
	filePath := "users.txt"

	// Check if the file exists, if doesn't, generate it with random users
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		gen.GenerateRandomUsers(filePath, 100000)
	}

	tag.TagFile(filePath)
}
