package fluffy

import "fmt"

type ValidationError[T any] struct {
	message string
	value   T
}

func (v ValidationError[T]) Error() string {
	return v.message
}

type ValidationErrorList[T any] struct {
	errors []ValidationError[T]
}

func (v ValidationErrorList[T]) Error() string {
	messages := make([]string, len(v.errors))

	for _, err := range v.errors {
		messages = append(messages, err.Error())
	}

	return fmt.Sprintf("ValidationErrorList: %s", messages)
}

func (v ValidationErrorList[T]) AddError(err ValidationError[T]) {
	v.errors = append(v.errors, err)
}
