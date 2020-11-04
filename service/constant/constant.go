package constant

import "fmt"

var (
	ErrAlreadyExists = fmt.Errorf("Already Exists")
	ErrNotFound      = fmt.Errorf("Not Found")
)
