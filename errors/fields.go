package main

type InvalidFieldError string
type MissingValueError string

func (e InvalidFieldError) Error() string {
	return "invalid field: " + string(e)
}

func (e MissingValueError) Error() string {
	return "missing value for field: " + string(e)
}
