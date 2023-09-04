package tag

import (
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

	stat, err := fileToTag.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fileSize := stat.Size()

	bufferSize := int(fileSize)
	buffer := make([]byte, bufferSize)
	_, err = fileToTag.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	modifiedContent := string(buffer)
	modifiedContent = strings.ReplaceAll(modifiedContent, "user", "[RISK] user")
	modifiedContent = strings.ReplaceAll(modifiedContent, "password", "[RISK] password")

	fileToTag.Seek(0, 0)
	_, err = fileToTag.WriteString(modifiedContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	if len(modifiedContent) < bufferSize {
		fileToTag.Truncate(int64(len(modifiedContent)))
	}

	fmt.Println("File tagged successfully!")
}
