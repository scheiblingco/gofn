// Error tools contains some commonly used error types
package errtools

type InvalidKeyError string
type InvalidFieldError string
type MissingValueError string
type InvalidTypeError string

func (e InvalidKeyError) Error() string {
	return "invalid key: " + string(e)
}

func (e InvalidFieldError) Error() string {
	return "invalid field: " + string(e)
}

func (e MissingValueError) Error() string {
	return "missing value for field: " + string(e)
}

func (e InvalidTypeError) Error() string {
	return "invalid type for field: " + string(e)
}
