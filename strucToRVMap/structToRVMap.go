package strucToRVMap

import (
	"reflect"
)

// StrucToRVMap blabla
func StrucToRVMap(data interface{}) (m map[string]*reflect.Value) {
	m = make(map[string]*reflect.Value)
	v := reflect.ValueOf(data)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		typefield := v.Type().Field(i)
		Valuefield := v.Field(i)

		if Valuefield.CanSet() == false {
			continue
		}
		// log.Println("typefield.Name:", typefield.Name)
		// log.Println("typefield.PkgPath:", typefield.PkgPath)

		m[typefield.Name] = &Valuefield
	}

	return
}
