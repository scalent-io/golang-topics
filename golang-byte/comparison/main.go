package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

func main() {
	const GB = 1024 * 1024 * 1024

	data := make([]byte, GB)
	for i := 0; i < GB; i++ {
		data[i] = byte(i % 256) // Filling data with bytes from 0 to 255
	}

	// Measure time to read from bytes.NewReader
	startTimeBytes := time.Now()
	byteReader := bytes.NewReader(data)
	byteBuffer := make([]byte, 1024)
	totalBytes := 0
	for {
		n, err := byteReader.Read(byteBuffer)
		totalBytes += n
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading with bytes.NewReader:", err)
			return
		}
	}
	elapsedBytes := time.Since(startTimeBytes)

	fmt.Printf("Read %d bytes with bytes.NewReader in %v\n", totalBytes, elapsedBytes)

	// Measure time to read from strings.NewReader
	startTimeString := time.Now()
	stringReader := strings.NewReader(string(data))
	stringBuffer := make([]byte, 1024)
	totalBytes = 0
	for {
		n, err := stringReader.Read(stringBuffer)
		totalBytes += n
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading with strings.NewReader:", err)
			return
		}
	}
	elapsedString := time.Since(startTimeString)

	fmt.Printf("Read %d bytes with strings.NewReader in %v\n", totalBytes, elapsedString)
}
