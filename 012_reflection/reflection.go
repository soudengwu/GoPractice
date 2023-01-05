package reflection

import "reflect"

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name string
	Profile
}

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// if val.Kind() == reflect.Slice {
	// 	for i := 0; i < val.Len(); i++ {
	// 		Walk(val.Index(i).Interface(), fn)
	// 	}
	// 	return
	// }

	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Field(i)
	// 	switch field.Kind() {
	// 	case reflect.String:
	// 		fn(field.String())
	// 	case reflect.Struct:
	// 		Walk(field.Interface(), fn)
	// 	}
	// }

	// switch val.Kind() {
	// case reflect.Struct:
	// 	for i := 0; i < val.NumField(); i++ {
	// 		Walk(val.Field(i).Interface(), fn)
	// 	}
	// case reflect.Slice:
	// 	for i := 0; i < val.Len(); i++ {
	// 		Walk(val.Index(i).Interface(), fn)
	// 	}
	// case reflect.String:
	// 	fn(val.String())
	// }

	numberOfValues := 0
	var getField func(int) reflect.Value

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walkValue(getField(i))
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
