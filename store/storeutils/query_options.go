package storeutils

import (
	"errors"
	"reflect"
)

// A short alias for a conditions map to facilitate writing.
// The key is the variable name in the queried struct.
type M map[string]any

// A short alias for a projection map to facilitate writing.
// The key is the variable name of the queried struct.
// The value is 1 if the variable is wanted, 0 otherwise.
type P map[string]int8

type QueryOptions struct {
	Conditions M
	Projection P
}

func ApplyTagFieldNames[T any](m map[string]T, variable any, tagName string) (map[string]T, error) {
	newM := map[string]T{}
	for k, v := range m {
		field, ok := reflect.ValueOf(variable).Type().FieldByName(k)
		if !ok {
			return nil, errors.New("Field name not found")
		}

		fieldName := field.Tag.Get(tagName)
		newM[fieldName] = v
	}

	return newM, nil
}

func (o *QueryOptions) ApplyTagFieldNames(variable any, tagName string) error {
	conds, err := ApplyTagFieldNames(o.Conditions, variable, tagName)
	if err != nil {
		return err
	}

	proj, err := ApplyTagFieldNames(o.Projection, variable, tagName)
	if err != nil {
		return err
	}

	o.Conditions = conds
	o.Projection = proj

	return nil
}
