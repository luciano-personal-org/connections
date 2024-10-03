// Package helpers contains helper functions, for example, the function to read the buffer from the connection.
package connections

import (
	"context"
	"fmt"
	"net"
)

// Buffer Size to capture the data
// const bufferSize = 4096

// Readbuffer reads the buffer from the connection.
func Readbuffer(ctx context.Context, conn net.Conn, buffer []byte) (record string, err error) {

	// Create a buffer
	// TODO: Move to Configuration
	// TODO: Checkout for memory leak
	// buffer := make([]byte, bufferSize)

	// Read the buffer
	_, err = conn.Read(buffer)

	// Check if there is an error
	if err != nil {
		fmt.Println("Error reading", err.Error())
		return "", err
	}

	// Return the record
	record = string(buffer)

	// Return the record
	return record, nil
}
