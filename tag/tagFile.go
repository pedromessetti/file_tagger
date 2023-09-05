package tag

import (
	"fmt"
	"os"
	"strings"
)

func TagFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fileSize := stat.Size()

	bufferSize := int(fileSize)
	buffer := make([]byte, bufferSize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	modifiedContent := string(buffer)
	modifiedContent = strings.ReplaceAll(modifiedContent, "user", "[RISK] user")
	modifiedContent = strings.ReplaceAll(modifiedContent, "password", "[RISK] password")

	file.Seek(0, 0)
	_, err = file.WriteString(modifiedContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	if len(modifiedContent) < bufferSize {
		file.Truncate(int64(len(modifiedContent)))
	}

	fmt.Println("File tagged successfully!")
}
