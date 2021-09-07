package serverutils

type NotFoundError struct{}

func (nfe NotFoundError) Error() string {
	return "Not found"
}
