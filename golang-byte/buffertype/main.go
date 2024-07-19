package main

import (
	"bytes"
	"fmt"
)

func main() {

	// Creating a new buffer
	var buf bytes.Buffer

	// Writing data to the buffer
	buf.WriteString("Hello, ")
	buf.WriteByte('W')
	buf.Write([]byte("orld!"))

	// Reading data from the buffer using io.Reader interface
	data := make([]byte, buf.Len())
	_, err := buf.Read(data)
	if err != nil {
		fmt.Println("Error reading data from buffer:", err)
		return
	}
	fmt.Println("Data read from buffer:", string(data))

	// Using io.Writer interface to write data to the buffer
	_, err = fmt.Fprintf(&buf, " How are you?")
	if err != nil {
		fmt.Println("Error writing data to buffer:", err)
		return
	}

	// Outputting the entire content of the buffer
	fmt.Println("Buffer contents:", buf.String())
}
