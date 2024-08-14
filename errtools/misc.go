package errtools

import "fmt"

type MultipleErrors []error

func (me MultipleErrors) Error() string {
	errStr := ""

	for _, err := range me {
		errStr += err.Error() + "\r\n"
	}

	return fmt.Sprintf("error(s) occured (%d): \r\n%s", len(me), errStr)
}
