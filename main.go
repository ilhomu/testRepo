package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Define the file name and content
	fileName := "example.txt"
	fileContent := []byte("Hello, Go File Handling!")

	// 1. Create and write to a file
	err := ioutil.WriteFile(fileName, fileContent, 0644) // 0644 sets file permissions
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Successfully wrote to", fileName)

	// 2. Read from a file
	readContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("Content read from", fileName + ":", string(readContent))

	// 3. Append to a file (using os.OpenFile)
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for appending:", err)
		return
	}
	defer f.Close() // Ensure the file is closed

	appendContent := []byte("\nAppending some more text.")
	_, err = f.Write(appendContent)
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("Successfully appended to", fileName)

	// 4. Read the file again to see appended content
	readContentAfterAppend, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading from file after append:", err)
		return
	}
	fmt.Println("Content after appending:", string(readContentAfterAppend))

	// 5. Delete the file
	err = os.Remove(fileName)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("Successfully deleted", fileName)
}
