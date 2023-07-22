# fluffy

Fluffy is a golang fluent validation.

```go
import "github.com/DarknessRdg/fluffy"

validateString := flufly.NewStringValidator().
  Contains("required").
  Len(64)


valid, err := validateString.Validate("failed string)
```
