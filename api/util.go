package api

import (
	"reflect"
)

// structToMap converts struct to map without parent-child nesting;
// the key is in the format of parent.child; value is the child value.
func structToMap(s interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})
	sv := reflect.ValueOf(s)
	st := reflect.TypeOf(s)

	for i := 0; i < st.NumField(); i++ {
		field := sv.Field(i)
		tag := st.Field(i).Tag.Get("json")

		switch field.Kind() {
		case reflect.Struct:
			m2 := structToMap(field.Interface())
			for k, v := range m2 {
				t := tag + "." + k
				m[t] = v
			}
		case reflect.String:
			m[tag] = field
		}
	}

	return m
}
