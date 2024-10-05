package connections

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/luciano-personal-org/config"
)

// Initialize the subscription
func InitSubscription(ctx context.Context, conn net.Conn, subscription string) {
	conn.Write([]byte(subscription + "\n"))
}

// ConnectToCedroServer is an exported function that connects to the Cedro server.
func ConnectToCedroServer(ctx context.Context, endpoint string, simulation bool, configuration config.Config) (net.Conn, error) {

	// Connect to the Cedro server
	conn, err := net.Dial("tcp4", endpoint)
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		os.Exit(1)
	}

	if simulation {
		conn.Write([]byte("\n"))
	}

	// Get the buffer size from the configuration
	bufferSizeStr := configuration.Get("connector_processor.buffer_size")
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	if err != nil {
		fmt.Println("Error converting buffer size", err.Error())
		return nil, err
	}

	// Create a buffer
	buffer := make([]byte, bufferSize)
	// Read the buffer
	_, err = Readbuffer(ctx, conn, buffer)
	if err != nil {
		fmt.Println("Error reading", err.Error())
		os.Exit(1)
	}

	// Print the buffer
	fmt.Print(string(buffer))

	// Empty Buffer
	buffer = nil

	// Return the connection
	return conn, err
}

// SoftwareKey is an exported function that sends the software key to the Cedro server.
func SoftwareKey(ctx context.Context, softwarekey string, conn net.Conn, configuration config.Config) (err error) {

	// Prepare the message with the software key
	// message := "SoftwareKey:" + softwarekey + "\n"
	message := "\n"
	fmt.Print(message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message", err.Error())
		os.Exit(1)

	}

	// Get the buffer size from the configuration
	bufferSizeStr := configuration.Get("connector_processor.buffer_size")
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	if err != nil {
		fmt.Println("Error converting buffer size", err.Error())
		return err
	}

	// Create a buffer
	buffer := make([]byte, bufferSize)
	// Read the buffer
	_, err = Readbuffer(ctx, conn, buffer)
	if err != nil {
		fmt.Println("Error reading", err.Error())
		os.Exit(1)
	}

	// Print the buffer
	fmt.Print(string(buffer))

	// Empty Buffer
	buffer = nil

	return nil
}

// UserLogin is an exported function that sends the username to the Cedro server.
func UserLogin(ctx context.Context, username string, conn net.Conn, configuration config.Config) (err error) {

	message := username + "\n"
	fmt.Print(message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message", err.Error())
		os.Exit(1)

	}

	// Get the buffer size from the configuration
	bufferSizeStr := configuration.Get("connector_processor.buffer_size")
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	if err != nil {
		fmt.Println("Error converting buffer size", err.Error())
		return err
	}

	// Create a buffer
	buffer := make([]byte, bufferSize)
	// Read the buffer
	_, err = Readbuffer(ctx, conn, buffer)
	if err != nil {
		fmt.Println("Error reading", err.Error())
		os.Exit(1)
	}

	// Print the buffer
	fmt.Print(string(buffer))

	// Empty Buffer
	buffer = nil

	return nil
}

// UserPassword is an exported function that sends the password to the Cedro server.
func UserPassword(ctx context.Context, password string, conn net.Conn, configuration config.Config) (err error) {

	message := password + "\n"
	fmt.Print(message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message", err.Error())
		os.Exit(1)

	}

	// Get the buffer size from the configuration
	bufferSizeStr := configuration.Get("connector_processor.buffer_size")
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	if err != nil {
		fmt.Println("Error converting buffer size", err.Error())
		return err
	}

	// Create a buffer
	buffer := make([]byte, bufferSize)
	// Read the buffer
	_, err = Readbuffer(ctx, conn, buffer)
	if err != nil {
		fmt.Println("Error reading", err.Error())
		os.Exit(1)
	}

	// Print the buffer
	fmt.Print(string(buffer))

	// Empty Buffer
	buffer = nil

	return nil
}

// NewCedroConnector is an exported function that creates a new CedroConnector.
func NewCedroConnector(ctx context.Context, configuration config.Config) (net.Conn, error) {

	// Get the software key from the configuration
	endpoint := configuration.Get("connector_processor.hostport")
	// Get the software key from the configuration
	softwarekey := configuration.Get("connector_processor.softwarekey")
	// // Get the username from the configuration
	username := configuration.Get("connector_processor.username")
	// // Get the password from the configuration
	password := configuration.Get("connector_processor.password")

	// Get the simulation value from the configuration
	simulation, err := strconv.ParseBool(configuration.Get("connector_processor.simulation"))
	if err != nil {
		fmt.Println("Error parsing simulation value", err.Error())
		os.Exit(1)
	}

	// Connect to the Cedro server
	conn, err := ConnectToCedroServer(ctx, endpoint, simulation, configuration)

	// Send the software key to the Cedro server
	err = SoftwareKey(ctx, softwarekey, conn, configuration)
	if err != nil {
		fmt.Println("Error sending message", err.Error())
		os.Exit(1)
	}

	// Send the username to the Cedro server
	err = UserLogin(ctx, username, conn, configuration)
	if err != nil {
		fmt.Println("Error sending message", err.Error())
		os.Exit(1)
	}

	err = UserPassword(ctx, password, conn, configuration)
	if err != nil {
		fmt.Println("Error sending message", err.Error())
		os.Exit(1)
	}

	return conn, err
}
