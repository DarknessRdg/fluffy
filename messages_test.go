package fluffy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageFieldsAreUnique(t *testing.T) {
	allMessageFields := []messageFields{
		Type, Value, ValueLen, ExpectedLen, ExpectedToContain, NotExpectedToContain,
	}

	set := make(map[messageFields]bool)

	for _, field := range allMessageFields {
		set[field] = true
	}

	assert.Len(
		t,
		set,
		len(allMessageFields),
		"There are message fields duplicated. Contains %d message fields, but only %d are unique",
		len(allMessageFields),
		len(set),
	)
}
