package main

import (
	"os"

	gen "github.com/pedromessetti/tag_file/gen"
	tag "github.com/pedromessetti/tag_file/tag"

)

func main() {
	file := "users.txt"

	// Check if the file exists, if doesn't, generate it with random users
	if _, err := os.Stat(file); os.IsNotExist(err) {
		gen.GenerateRandomUsers(10000)
	}

	tag.TagFile(file)
}
