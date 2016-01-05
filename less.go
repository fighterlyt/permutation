package permutation

import (
    "reflect"
    "errors"
)

type Less func(left, right interface{}) bool

func getLessFunctionByValueType(value reflect.Value) (Less, error) {
    switch value.Type().Elem().Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return lessInt, nil
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return lessUint, nil
    case reflect.Float32, reflect.Float64:
        return lessFloat, nil
    case reflect.String:
        return lessString, nil
    default:
        return nil, errors.New("the element type of slice is not ordered, you must provide a function\n")
    }
}

func lessUint(left, right interface{}) bool {
	return reflect.ValueOf(left).Uint() < reflect.ValueOf(right).Uint()
}

func lessInt(left, right interface{}) bool {
	return reflect.ValueOf(left).Int() < reflect.ValueOf(right).Int()
}

func lessFloat(left, right interface{}) bool {
	return reflect.ValueOf(left).Float() < reflect.ValueOf(right).Float()
}

func lessString(left, right interface{}) bool {
	return reflect.ValueOf(left).Interface().(string) < reflect.ValueOf(right).Interface().(string)
}