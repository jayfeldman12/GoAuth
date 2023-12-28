package lambdas

import (
	"fmt"
	"reflect"
)

func LogAll(data interface{}) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	if t.Kind() != reflect.Struct {
		fmt.Println("Input is not a struct")
		return
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		// Print out the field name and value, assuming value is a strirng

		fmt.Printf("Field: %s, Value: %v\n", field.Name, value.Interface())
	}
}
