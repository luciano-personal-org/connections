package connections_exception

import (
	exceptioncore "github.com/luciano-personal-org/exception"
)

var (
	CedroConnectionError     = exceptioncore.NewTradingError("CON001", "Unable to connect to Cedro to obtain Market Data")
	BufSizeParameterError    = exceptioncore.NewTradingError("CON002", "Error converting buffer size parameter, which is an integer")
	SendingMessageError      = exceptioncore.NewTradingError("CON003", "Error sending message over the socket")
	ReceivingMessageError    = exceptioncore.NewTradingError("CON004", "Error receiving message from the socket")
	SimulationParameterError = exceptioncore.NewTradingError("CON005", "Error processing simulation parameter, which is a boolean")
)
