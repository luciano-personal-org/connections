// Package helpers contains helper functions, for example, the function to read the buffer from the connection.
package connections

import (
	"context"
	"net"
)

// Readbuffer reads the buffer from the connection.
func Readbuffer(ctx context.Context, conn net.Conn, buffer []byte) (record string, err error) {

	// Read the buffer
	_, err = conn.Read(buffer)

	// Check if there is an error
	if err != nil {
		return "", err
	}

	// Return the record
	record = string(buffer)

	// Return the record
	return record, nil
}
