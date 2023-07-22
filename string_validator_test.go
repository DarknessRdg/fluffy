package fluffy_validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringValidator_Len(t *testing.T) {
	defaultMessage := "Not exact len"

	tests := []struct {
		name     string
		exactLen int
		value    string
		valid    bool
	}{
		{
			name:     "When value len is lower, Then return error",
			exactLen: 3,
			value:    "",
			valid:    false,
		},
		{
			name:     "When value len is greater, Then return error",
			exactLen: 3,
			value:    "1234",
			valid:    false,
		},
		{
			name:     "When value len is equal, Then return valid and no error",
			exactLen: 3,
			value:    "123",
			valid:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewStringValidator()
			v.Len(tt.exactLen)

			valid, err := v.Validate(tt.value)

			assert.Equal(t, tt.valid, valid)

			if !tt.valid {
				assert.Equal(t, defaultMessage, err.message)
				assert.Equal(t, tt.value, err.value)
			}
		})
	}
}

func TestStringValidator_Lenf(t *testing.T) {
	tests := []struct {
		name            string
		exactLen        int
		value           string
		valid           bool
		format          string
		expectedMessage string
		formatFields    []messageFields
	}{
		{
			name:     "When value len is lower, Then return error",
			exactLen: 3,
			value:    "",
			valid:    false,
		},
		{
			name:     "When value len is greater, Then return error",
			exactLen: 3,
			value:    "1234",
			valid:    false,
		},
		{
			name:     "When value len is equal, Then return valid and no error",
			exactLen: 3,
			value:    "123",
			valid:    true,
		},
		{
			name:            "When give custom message without any param, Then error message shouldn't change",
			exactLen:        3,
			value:           "",
			valid:           false,
			format:          "Message without any extra param.",
			expectedMessage: "Message without any extra param.",
		},
		{
			name:            "When give custom message with all params allowed, Then return message with each param in the format respecting params order",
			exactLen:        3,
			value:           "",
			format:          "Type = %s ; ValueLen = %d ; ExpectedLen = %d",
			expectedMessage: "Type = string ; ValueLen = 0 ; ExpectedLen = 3",
			formatFields:    []messageFields{Type, ValueLen, ExpectedLen},
			valid:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewStringValidator()
			v.Lenf(tt.exactLen, tt.format, tt.formatFields...)

			valid, err := v.Validate(tt.value)

			assert.Equal(t, tt.valid, valid)

			if !tt.valid {
				assert.Equal(t, tt.expectedMessage, err.message)
				assert.Equal(t, tt.value, err.value)
			}
		})
	}
}

func TestStringValidator_buildFieldsValues(t *testing.T) {
	allFieldsConfig := map[messageFields]any{
		ExpectedLen: 3,
		Type:        "type",
		ValueLen:    1,
	}

	tests := []struct {
		name     string
		config   map[messageFields]any
		fields   []messageFields
		expected []any
	}{
		{
			name:     "When empty field, Then return empty values",
			config:   allFieldsConfig,
			fields:   []messageFields{},
			expected: []any{},
		},
		{
			name:     "When only add Expected, Then return expected len",
			config:   allFieldsConfig,
			fields:   []messageFields{ExpectedLen},
			expected: []any{3},
		},
		{
			name:     "When only add Type, Then return type",
			config:   allFieldsConfig,
			fields:   []messageFields{ExpectedLen},
			expected: []any{3},
		},
		{
			name:     "When only add ValueLen, Then return value len",
			config:   allFieldsConfig,
			fields:   []messageFields{ValueLen},
			expected: []any{1},
		},
		{
			name:     "When add all fields, Then return all fields respecting the fields order",
			config:   allFieldsConfig,
			fields:   []messageFields{Type, ValueLen, ExpectedLen},
			expected: []any{"type", 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewStringValidator()

			got := v.buildFieldsValues(tt.config, tt.fields...)
			assert.Equal(t, tt.expected, got)
		})
	}
}
