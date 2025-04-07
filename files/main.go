// using buffered io

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	// Create a file
// 	file, _ := os.Create("example.txt")
// 	defer file.Close()

// 	content := "This is an example file created using Go."
// 	// Write content to the file

// 	// io.WriteString takes two arguments: the file to write to and the string to write
// 	_, err := io.WriteString(file, content+"\n")
// 	if err != nil {
// 		fmt.Println("Error writing to file:", err)
// 		return
// 	}

// 	// Read the file
// 	file, _ = os.Open("example.txt")
// 	defer file.Close()

// 	// Creating a buffer to read the file
// 	buffer := make([]byte, 1024) // Here we create a buffer of 1024 bytes

// 	// Read the content of the file into the buffer
// 	for {
// 		n, err := file.Read(buffer) // It will return either integer or error

// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			fmt.Println("Error reading file:", err)
// 			return
// 		}

// 		// Process the read content
// 		fmt.Print(string(buffer[:n])) // Print the content read from the file

// 	}

// 	fmt.Println("File created successfully")
// }

// Using io util
package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:", string(content))

}
