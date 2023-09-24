package Miceleanos

import (
	"reflect"
)

func IsInt(Valor interface{}) bool {
	return reflect.TypeOf(Valor).Kind() == reflect.Int64
}

func Isbool(Valor interface{}) bool {
	return reflect.TypeOf(Valor).Kind() == reflect.Bool
}

func IsString(Valor interface{}) bool {
	return reflect.TypeOf(Valor).Kind() == reflect.String
}

func IsFloat(Valor interface{}) bool {
	return reflect.TypeOf(Valor).Kind() == reflect.Float64
}

func IsNull(Valor interface{}) bool {
	if strValue, ok := Valor.(string); ok {
		return strValue == "?"
	}
	return false
}

func IsChar(Valor interface{}) bool {
	_, ok := Valor.(rune)
	return ok
}
