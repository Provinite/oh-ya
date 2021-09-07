package serverutils

import "fmt"

type CompoundError struct {
	Errors []error
}

func (ce CompoundError) Error() string {
	result := ""
	for _, err := range ce.Errors {
		result = fmt.Sprintf("%s\n%s", result, err.Error())
	}
	return result
}
