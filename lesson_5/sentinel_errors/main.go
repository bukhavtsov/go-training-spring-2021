package sentinel_errors

type MyOwnError string

func (m MyOwnError) Error() string {
	return string(m)
}

type MyStatusError int

const (
	ErrorFirst MyStatusError = iota + 1
	ErrorSecond
)

type ErrorWithStatus struct {
	statusCode MyStatusError
	message    string
}
