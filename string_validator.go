package fluffy_validator

import (
	"fmt"
	"github.com/DarknessRdg/fluffy-validator/internal/utils/maps"
	"strings"
)

type StringValidator struct {
	*Validator[string]
}

func NewStringValidator() *StringValidator {
	return &StringValidator{
		Validator: New[string](),
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
// Message fields options are:
//   - Type
//   - Value
//   - ValueLen
//   - ExpectedLen
func (v *StringValidator) Lenf(exactLen int, format string, fields ...messageFields) *StringValidator {
	config := v.getDefaultStringMessagesConfig()
	config[ExpectedLen] = exactLen

	isValid := func(value string) bool {
		config[ValueLen] = len(value)
		return len(value) == exactLen
	}

	return v.addRule(isValid, format, config, fields...)
}

// Containsf validated given string contains a required substring.
// It's the same as `strings.Contains(value, substring)`
// Message fields options are:
//   - Type
//   - Value
//   - ExpectedToContain
func (v *StringValidator) Containsf(contains string, format string, fields ...messageFields) *StringValidator {
	config := v.getDefaultStringMessagesConfig()
	config[ExpectedToContain] = contains

	isValid := func(value string) bool {
		return strings.Contains(value, contains)
	}

	return v.addRule(isValid, format, config, fields...)
}

func (v *StringValidator) getDefaultStringMessagesConfig() map[messageFields]any {
	return map[messageFields]any{
		Type: "string",
	}
}

func (v *StringValidator) buildFieldsValues(config map[messageFields]any, fields ...messageFields) []any {
	return maps.FilterValuesInKeys(config, fields...)
}

func (v *StringValidator) addRule(
	isValid func(string) bool,
	format string,
	config map[messageFields]any,
	fields ...messageFields,
) *StringValidator {
	v.Validator.AddRule(func(value string) (bool, ValidationError[string]) {
		config[Value] = value

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
