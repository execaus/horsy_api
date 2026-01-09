package pkg

import "github.com/aarondl/opt/null"

func ConvertNullableVarToPtr[T any](value null.Val[T]) *T {
	if value.IsValue() {
		v := value.GetOrZero()
		return &v
	}

	return nil
}
