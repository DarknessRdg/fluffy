package fluffy

type IRule[T any] func(value T) (bool, ValidationError[T])

type Validator[T any] struct {
	rules []IRule[T]
}

func New[T any]() *Validator[T] {
	return &Validator[T]{
		rules: make([]IRule[T], 0),
	}
}

func (v *Validator[T]) AddRule(rule IRule[T]) *Validator[T] {
	v.rules = append(v.rules, rule)
	return v
}

func (v *Validator[T]) Validate(value T) (bool, ValidationError[T]) {
	for _, rule := range v.rules {
		valid, err := rule(value)
		if !valid {
			return valid, err
		}
	}

	return true, ValidationError[T]{}
}

func (v *Validator[T]) ValidateAll(value T) (bool, ValidationErrorList[T]) {
	valid := true
	errorsList := ValidationErrorList[T]{}

	for _, rule := range v.rules {
		ruleValid, err := rule(value)
		if !ruleValid {
			errorsList.AddError(err)
			valid = false
		}
	}

	return valid, errorsList
}
