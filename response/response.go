package response

type Response interface {
	StatusCode() int
	getBody() ([]byte, error)
	Error() string
	getData() interface{}
}
