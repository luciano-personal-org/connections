package connections_exception

import (
	exceptioncore "github.com/luciano-personal-org/exception"
)

var (
	ErrCedroConnection = exceptioncore.NewTradingError("E001", "Unable to connect to Cedro to obtain Market Data")
)
