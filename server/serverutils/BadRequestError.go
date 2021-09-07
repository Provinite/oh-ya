package serverutils

type BadRequestError struct {
}

func (e BadRequestError) Error() string {
	return "Bad request"
}
