package tag

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TagFile(file string) {
	fileToTag, err := os.OpenFile(file, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileToTag.Close()

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(fileToTag)

	// Slice to store the modified lines
	var modifiedLines []string

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "user") || strings.Contains(line, "password") {
			line = "[RISK] " + line
		}
		modifiedLines = append(modifiedLines, line)
		// fmt.Println(modifiedLines)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fileToTag.Truncate(0)
	fileToTag.Seek(0, 0)

	for _, line := range modifiedLines {
		fmt.Fprintln(fileToTag, line)
	}
	fmt.Println("File tagged successfully!")
}
