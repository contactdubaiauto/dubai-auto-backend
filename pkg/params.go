package pkg

import (
	"reflect"
	"strconv"
	"strings"
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
		return nil, nil, nil
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()

		if value == nil {
			continue
		}

		if reflect.ValueOf(value).IsZero() {
			continue
		}

		jsonTag := field.Tag.Get("json")

		if jsonTag == "" {
			return nil, nil, nil
		}

		keys = append(keys, jsonTag)
		values = append(values, "$"+strconv.Itoa(placeHolder))
		placeHolder++
		args = append(args, value)
	}
	return keys, values, args
}

func QueryParamToArray(param string) []string {
	if param == "" {
		return nil
	}
	parts := strings.Split(param, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func QueryParamToIntArray(param string) ([]int, error) {
	if param == "" {
		return nil, nil
	}
	parts := strings.Split(param, ",")
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}
