package customError

import (
	"fmt"
)

func ErrorNoFile() error {
	return fmt.Errorf("NO FILE HAVE BEEN SELECTED")
}
