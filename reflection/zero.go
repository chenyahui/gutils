package reflection

import "reflect"

func IsZero(x interface{}) bool {
	if x == nil {
		return true
	}

	reflectValue := reflect.ValueOf(x)
	if reflectValue.Kind() == reflect.Slice {
		return reflectValue.Len() == 0
	}

	return reflectValue.IsZero()
}
