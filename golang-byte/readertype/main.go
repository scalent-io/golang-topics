package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// Example data
	data := []byte("Hello, World!")

	// Create a bytes.Reader with initial data
	reader := bytes.NewReader(data)

	// Reading data from the reader using io.Reader interface
	buffer := make([]byte, len(data))
	_, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading data:", err)
		return
	}
	fmt.Println("Data read from bytes.Reader:", string(buffer))

	// Seek back to the beginning of the reader
	reader.Seek(0, io.SeekStart)

	// Reading a single byte
	b, err := reader.ReadByte()
	if err != nil {
		fmt.Println("Error reading byte:", err)
		return
	}
	fmt.Println("Single byte read from bytes.Reader:", string(b))

	// Reading a single rune
	r, size, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error reading rune:", err)
		return
	}
	fmt.Printf("Single rune read from bytes.Reader: %c (size: %d)\n", r, size)

	// Reset the reader to start again
	reader.Reset(data)

	// Copying data from reader to writer
	var writer bytes.Buffer
	n, err := io.Copy(&writer, reader)
	if err != nil {

		fmt.Println("Error copying data:", err)
		return
	}
	fmt.Printf("Copied %d bytes from bytes.Reader to buffer: %s\n", n, writer.String())
}
