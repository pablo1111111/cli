package p

import "reflect"

func StrP(v string) *string {
	return &v
}

func PStr(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func BoolP(v bool) *bool {
	return &v
}

func PBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// IsNil asserts whether the underlying type is nil, which `interface{} == nil` does not
func IsNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}
