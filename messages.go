package fluffy

// messageFields are the fields available to provide a custom error message formatted with the validations constraints.
// For instance, to add the value validated to the error message.
type messageFields string

const (
	// Type is a message field that returns the element data type. For instance: int8, string.
	Type = messageFields("Type")
	// Value is the actual value being validated
	Value = messageFields("Value")
	// ValueLen is the len() of the value. It's only present in len validations.
	ValueLen = messageFields("ValueLen")
	// ExpectedLen is the len constraint expected but was not equals to value len. It's only present in len validations.
	ExpectedLen = messageFields("ExpectedLen")
	// ExpectedToContain is the value the was expected to contain. It's only available in `contains` validations.
	ExpectedToContain = messageFields("ExpectedToContain")
	// NotExpectedToContain is the value that was not expect to be present.
	// It's only available in `NotContains` validations
	NotExpectedToContain = messageFields("NotExpectedToContain")
)
