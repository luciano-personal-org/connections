package connections_exception

import (
	exceptioncore "github.com/luciano-personal-org/exception"
)

var (
	CedroConnectionError     = exceptioncore.NewTradingError("E001", "Unable to connect to Cedro to obtain Market Data")
	BufSizeParameterError    = exceptioncore.NewTradingError("E002", "Error converting buffer size parameter, which is an integer")
	SendingMessageError      = exceptioncore.NewTradingError("E003", "Error sending message over the socket")
	ReceivingMessageError    = exceptioncore.NewTradingError("E004", "Error receiving message from the socket")
	SimulationParameterError = exceptioncore.NewTradingError("E005", "Error processing simulation parameter, which is a boolean")
)
