package maps

func FilterValuesInKeys[Key comparable, Value any](mapInstance map[Key]Value, keysToFilter ...Key) []Value {
	values := make([]Value, 0, len(keysToFilter))

	for _, key := range keysToFilter {
		value, ok := mapInstance[key]

		if ok {
			values = append(values, value)
		}
	}

	return values
}
