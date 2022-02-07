package errors

import "fmt"

var (
	ErrInvalidConfig = fmt.Errorf("error the config file is not valid")
)
