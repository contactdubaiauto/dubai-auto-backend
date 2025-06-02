package pkg

import (
	"fmt"
	"reflect"
	"strconv"
)

func BuildParams(v interface{}) (keys []string, values []string, args []interface{}) {
	placeHolder := 1

	if v == nil {
		return nil, nil, nil
	}

	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil, nil, nil
		}
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("Not a struct")
		return nil, nil, nil
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()

		if value == nil {
			fmt.Println("Nil value")
			continue
		}

		if reflect.ValueOf(value).IsZero() {
			fmt.Println("Zero value")
			continue
		}

		jsonTag := field.Tag.Get("json")

		if jsonTag == "" {
			fmt.Println("No json tag")
			return nil, nil, nil
		}

		fmt.Printf("Field: %s, Value: %v\n", jsonTag, value)
		keys = append(keys, jsonTag)
		values = append(values, "$"+strconv.Itoa(placeHolder))
		placeHolder++
		args = append(args, value)
	}
	return keys, values, args
}
