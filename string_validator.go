package fluffy_validator

import "fmt"

type stringMessageFields string

const (
	Type        = stringMessageFields("type")
	ValueLen    = "ValueLne"
	ExpectedLen = "ExpectedLne"
)

type StringValidator struct {
	*Validator[string]
}

func NewStringValidator() *StringValidator {
	return &StringValidator{
		Validator: NewValidator[string](),
	}
}

// Len validates given string has exact len as expected.
// It is the same as `len(value) == expectedLen`
func (v *StringValidator) Len(exactLen int) *StringValidator {
	message := "Not exact len"
	v.Lenf(exactLen, message)

	return v
}

// Lenf is the same as `len` validator, with addition of custom message when it fails.
func (v *StringValidator) Lenf(exactLen int, format string, fields ...stringMessageFields) *StringValidator {
	isValid := func(value string) bool {
		return len(value) == exactLen
	}

	config := v.getDefaultStringMessagesConfig()
	config[ExpectedLen] = exactLen

	v.AddRule(func(value string) (bool, ValidationError[string]) {
		config[ValueLen] = len(value)
		err := ValidationError[string]{}

		valid := isValid(value)

		if !valid {
			errMessageFields := v.buildFieldsValues(config, fields...)

			err = ValidationError[string]{
				message: fmt.Sprintf(format, errMessageFields...),
				value:   value,
			}
		}

		return valid, err
	})

	return v
}

func (v *StringValidator) getDefaultStringMessagesConfig() map[stringMessageFields]any {
	config := make(map[stringMessageFields]any)

	config[Type] = "string"
	return config
}

func (v *StringValidator) buildFieldsValues(config map[stringMessageFields]any, fields ...stringMessageFields) []any {
	values := make([]any, 0, len(fields))

	for _, field := range fields {
		value, ok := config[field]

		if ok {
			values = append(values, value)
		}
	}
	return values
}
