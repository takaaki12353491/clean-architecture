package errs

type HTTPError interface {
	error
	StatusCode() int
}

type httpError struct {
	statusCode int
	msg        string
}

func NewHTTPError(statuscode int, msg string) HTTPError {
	return &httpError{
		statusCode: statuscode,
		msg:        msg,
	}
}

func (err *httpError) Error() string {
	return err.msg
}

func (err *httpError) StatusCode() int {
	return err.statusCode
}
