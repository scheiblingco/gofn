package errtools

type BodyNotAcceptedError string
type BodyConsumedError string

func (e BodyNotAcceptedError) Error() string {
	return "body not accepted: " + string(e)
}

func (e BodyConsumedError) Error() string {
	return "body already consumed: " + string(e)
}
