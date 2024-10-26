package connections

import (
	"context"
	"fmt"
	"net"
	"strconv"

	exceptioncore "github.com/luciano-personal-org/exception"

	"github.com/luciano-personal-org/config"
	"github.com/luciano-personal-org/connections/connections_exception"
)

// Initialize the custom error
var custom_error exceptioncore.TradingError

// Initialize the subscription
func InitSubscription(ctx context.Context, conn net.Conn, subscription string) {
	conn.Write([]byte(subscription + "\n"))
}

// ConnectToCedroServer is an exported function that connects to the Cedro server.
func ConnectToCedroServer(ctx context.Context, endpoint string, simulation bool, configuration config.Config) (net.Conn, error) {

	// Connect to the Cedro server
	conn, err := net.Dial("tcp4", endpoint)
	// Check if there was an error connecting to the Cedro server
	if err != nil {
		custom_error = connections_exception.CedroConnectionError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to connect into Quotes Server on ConnectToCedroServer")
		exceptioncore.DoPanic(custom_error)
	}

	if simulation {
		conn.Write([]byte("\n"))
	}

	// Get the buffer size from the configuration
	bufferSizeStr := configuration.Get("connector_processor.buffer_size")
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	// Check if there was an error converting the buffer size
	if err != nil {
		custom_error = connections_exception.BufSizeParameterError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to convert the buffer size parameter on ConnectToCedroServer")
		exceptioncore.DoPanic(custom_error)
	}

	// Create a buffer
	buffer := make([]byte, bufferSize)
	// Read the buffer
	_, err = Readbuffer(ctx, conn, buffer)
	// Check if there was an error reading the buffer to receive a message
	if err != nil {
		custom_error = connections_exception.ReceivingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to receive a message from Quotes Server on ConnectToCedroServer")
		exceptioncore.DoPanic(custom_error)
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
	// Check if there was an error sending the message
	if err != nil {
		custom_error = connections_exception.SendingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to send a message with Software Key to Quotes Server on SoftwareKey")
		exceptioncore.DoPanic(custom_error)
	}

	// Get the buffer size from the configuration
	bufferSizeStr := configuration.Get("connector_processor.buffer_size")
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	// Check if there was an error converting the buffer size
	if err != nil {
		custom_error = connections_exception.BufSizeParameterError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to convert the buffer size parameter on SoftwareKey")
		exceptioncore.DoPanic(custom_error)
	}

	// Create a buffer
	buffer := make([]byte, bufferSize)
	// Read the buffer
	_, err = Readbuffer(ctx, conn, buffer)
	// Check if there was an error reading the buffer to receive a message
	if err != nil {
		custom_error = connections_exception.ReceivingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to receive a message of Software Key sent to Quotes Server on SoftwareKey")
		exceptioncore.DoPanic(custom_error)
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
	// Check if there was an error sending the message
	if err != nil {
		custom_error = connections_exception.SendingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to send a message with User Login to Quotes Server on UserLogin")
		exceptioncore.DoPanic(custom_error)
		return err
	}

	// Get the buffer size from the configuration
	bufferSizeStr := configuration.Get("connector_processor.buffer_size")
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	// Check if there was an error converting the buffer size
	if err != nil {
		custom_error = connections_exception.BufSizeParameterError
		custom_error.SetOriginalError(err)
		exceptioncore.DoPanic(custom_error)
	}

	// Create a buffer
	buffer := make([]byte, bufferSize)
	// Read the buffer
	_, err = Readbuffer(ctx, conn, buffer)
	// Check if there was an error reading the buffer to receive a message
	if err != nil {
		custom_error = connections_exception.ReceivingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to receive a message of User Login sent to Quotes Server on UserLogin")
		exceptioncore.DoPanic(custom_error)
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
	// Check if there was an error sending the message
	if err != nil {
		custom_error = connections_exception.SendingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to send a message with User Password to Quotes Server on UserPassword")
		exceptioncore.DoPanic(custom_error)
	}

	// Get the buffer size from the configuration
	bufferSizeStr := configuration.Get("connector_processor.buffer_size")
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	// Check if there was an error converting the buffer size
	if err != nil {
		custom_error = connections_exception.BufSizeParameterError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to convert the buffer size parameter on UserPassword")
		exceptioncore.DoPanic(custom_error)
	}

	// Create a buffer
	buffer := make([]byte, bufferSize)
	// Read the buffer
	_, err = Readbuffer(ctx, conn, buffer)
	// Check if there was an error reading the buffer to receive a message
	if err != nil {
		custom_error = connections_exception.ReceivingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to receive a message of User Password sent to Quotes Server on UserPassword")
		exceptioncore.DoPanic(custom_error)
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
	// Check if there was an error converting the simulation parameter
	if err != nil {
		custom_error = connections_exception.SimulationParameterError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to convert the Simulation parameter on NewCedroConnector")
		exceptioncore.DoPanic(custom_error)
	}

	// Connect to the Cedro server
	conn, err := ConnectToCedroServer(ctx, endpoint, simulation, configuration)

	// Send the software key to the Cedro server
	err = SoftwareKey(ctx, softwarekey, conn, configuration)
	// Check if there was an error sending the message
	if err != nil {
		custom_error = connections_exception.SendingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to finish Software Key to Quotes Server on NewCedroConnector")
		exceptioncore.DoPanic(custom_error)
	}

	// Send the username to the Cedro server
	err = UserLogin(ctx, username, conn, configuration)
	// Check if there was an error sending the message
	if err != nil {
		custom_error = connections_exception.SendingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to finish User Login to Quotes Server on NewCedroConnector")
		exceptioncore.DoPanic(custom_error)
	}

	// Send the password to the Cedro server
	err = UserPassword(ctx, password, conn, configuration)
	// Check if there was an error sending the message
	if err != nil {
		custom_error = connections_exception.SendingMessageError
		custom_error.SetOriginalError(err)
		custom_error.SetDetails("When trying to finish User Password to Quotes Server on NewCedroConnector")
		exceptioncore.DoPanic(custom_error)
	}

	return conn, err
}
