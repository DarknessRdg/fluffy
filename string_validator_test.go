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

func TestStringValidator_MinLenf(t *testing.T) {
	tests := []struct {
		name            string
		minLen          int
		value           string
		valid           bool
		format          string
		expectedMessage string
		formatFields    []messageFields
	}{
		{
			name:   "When value len is equal to min, Then don't return error",
			minLen: 3,
			value:  "123",
			valid:  true,
		},
		{
			name:   "When value len is greater than min len, Then don't return error",
			minLen: 3,
			value:  "1234",
			valid:  true,
		},
		{
			name:            "When give custom message without any param, Then error message shouldn't change",
			minLen:          3,
			value:           "",
			valid:           false,
			format:          "Message without any extra param.",
			expectedMessage: "Message without any extra param.",
		},
		{
			name:   "When value len is lower than min len, Then return error",
			minLen: 3,
			value:  "12",
			valid:  false,
		},
		{
			name:            "When invalid value, Then format message with all available message fields",
			minLen:          3,
			value:           "12",
			format:          "Type = %s ; Value = %s ; ValueLen = %d ; ExpectedLen = %d",
			expectedMessage: "Type = string ; Value = 12 ; ValueLen = 2 ; ExpectedLen = 3",
			formatFields:    []messageFields{Type, Value, ValueLen, ExpectedLen},
			valid:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewStringValidator().
				MinLenf(tt.minLen, tt.format, tt.formatFields...)

			valid, err := v.Validate(tt.value)

			assert.Equal(t, tt.valid, valid)

			if !tt.valid {
				assert.Equal(t, tt.expectedMessage, err.message)
				assert.Equal(t, tt.value, err.value)
			}
		})
	}
}

func TestStringValidator_MaxLenf(t *testing.T) {
	tests := []struct {
		name            string
		maxLen          int
		value           string
		valid           bool
		format          string
		expectedMessage string
		formatFields    []messageFields
	}{
		{
			name:   "When value len is equal to max len, Then don't return error",
			maxLen: 3,
			value:  "123",
			valid:  true,
		},
		{
			name:   "When value len is lower than max len, Then don't return error",
			maxLen: 3,
			value:  "12",
			valid:  true,
		},
		{
			name:   "When value len is greater than max len, Then return error",
			maxLen: 3,
			value:  "1234",
			valid:  false,
		},
		{
			name:            "When give custom message without any param, Then error message shouldn't change",
			maxLen:          0,
			value:           "1",
			valid:           false,
			format:          "Message without any extra param.",
			expectedMessage: "Message without any extra param.",
		},
		{
			name:            "When invalid value, Then format message with all available message fields",
			maxLen:          0,
			value:           "1",
			format:          "Type = %s ; Value = %s ; ValueLen = %d ; ExpectedLen = %d",
			expectedMessage: "Type = string ; Value = 1 ; ValueLen = 1 ; ExpectedLen = 0",
			formatFields:    []messageFields{Type, Value, ValueLen, ExpectedLen},
			valid:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewStringValidator().
				MaxLenf(tt.maxLen, tt.format, tt.formatFields...)

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

func TestStringValidator_Containsf(t *testing.T) {
	tests := []struct {
		name            string
		contains        string
		format          string
		value           string
		valid           bool
		fields          []messageFields
		expectedMessage string
	}{
		{
			name:     "When value contains the expected substring, Then don't return error",
			contains: "with",
			value:    "message with value",
			valid:    true,
		},
		{
			name:            "When value does not contain expected substring, Then return error",
			contains:        "not contains this",
			value:           "invalid message",
			format:          "Error message",
			expectedMessage: "Error message",
			valid:           false,
		},
		{
			name:     "Substring is always present in any string, Then don't return error",
			contains: "",
			value:    "Any string",
			valid:    true,
		},
		{
			name:            "When does not contains, Then return error message formatted with all fields values available",
			contains:        "Does not contain",
			value:           "invalid string",
			valid:           false,
			format:          "Type = %s , ExpectedToContains = %s , Value = %s",
			expectedMessage: "Type = string , ExpectedToContains = Does not contain , Value = invalid string",
			fields:          []messageFields{Type, ExpectedToContain, Value},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewStringValidator().Containsf(tt.contains, tt.format, tt.fields...)

			valid, err := validator.Validate(tt.value)

			assert.Equal(t, tt.valid, valid)

			if !valid {
				assert.Equal(t, tt.expectedMessage, err.message)
				assert.Equal(t, tt.value, err.value)
			}
		})
	}
}

func TestStringValidator_NotContainsf(t *testing.T) {
	tests := []struct {
		name            string
		notContains     string
		format          string
		value           string
		valid           bool
		fields          []messageFields
		expectedMessage string
	}{
		{
			name:        "When value does not contains the expected substring, Then don't return error",
			notContains: "is not a substring",
			value:       "message with value",
			valid:       true,
		},
		{
			name:            "When value contain a forbidden substring, Then return error",
			notContains:     "forbidden",
			value:           "This message should not contain the forbidden word",
			format:          "Error message",
			expectedMessage: "Error message",
			valid:           false,
		},
		{
			name:            "Substring is always present in any string, Then return error",
			notContains:     "",
			value:           "Any string",
			valid:           false,
			format:          "Error message",
			expectedMessage: "Error message",
		},
		{
			name:            "When contains, Then return error message formatted with all fields values available",
			notContains:     "forbidden",
			value:           "invalid string with forbidden word",
			valid:           false,
			format:          "Type = %s , NotExpectedToContain = %s , Value = %s",
			expectedMessage: "Type = string , NotExpectedToContain = forbidden , Value = invalid string with forbidden word",
			fields:          []messageFields{Type, NotExpectedToContain, Value},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewStringValidator().NotContainsf(tt.notContains, tt.format, tt.fields...)

			valid, err := validator.Validate(tt.value)

			assert.Equal(t, tt.valid, valid)

			if !valid {
				assert.Equal(t, tt.expectedMessage, err.message)
				assert.Equal(t, tt.value, err.value)
			}
		})
	}
}
