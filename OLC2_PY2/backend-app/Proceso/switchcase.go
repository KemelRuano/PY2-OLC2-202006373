package Proceso

import (
	"backend-app/Miceleanos"
)

type Switch struct {
	valorSwitch interface{}
	cases       interface{}
	linea       int
	columna     int
}

func NewSwitch(valorSwitch interface{}, cases interface{}) *Switch {
	return &Switch{
		valorSwitch: valorSwitch,
		cases:       cases,
	}
}

func (s *Switch) IsEqual() interface{} {
	switch a := s.valorSwitch.(type) {
	case int64:
		if Miceleanos.IsInt(s.cases) {
			return a == s.cases
		} else {

			newError := Miceleanos.ErrorAnalizador{
				Fila:        s.linea,
				Columna:     s.columna,
				Intruccion:  "Switch",
				Descripcion: "Los Tipos de datos no coinciden",
				Tipo:        "Semantico",
			}
			return newError
		}
	case float64:
		if Miceleanos.IsFloat(s.cases) {
			return a == s.cases
		} else {
			newError := Miceleanos.ErrorAnalizador{
				Fila:        s.linea,
				Columna:     s.columna,
				Intruccion:  "Switch",
				Descripcion: "Los Tipos de datos no coinciden",
				Tipo:        "Semantico",
			}
			return newError
		}
	case bool:
		if Miceleanos.Isbool(s.cases) {
			return a == s.cases
		} else {
			newError := Miceleanos.ErrorAnalizador{
				Fila:        s.linea,
				Columna:     s.columna,
				Intruccion:  "Switch",
				Descripcion: "Los Tipos de datos no coinciden",
				Tipo:        "Semantico",
			}
			return newError
		}

	case string:
		if Miceleanos.IsString(s.cases) {
			return a == s.cases
		} else {

			newError := Miceleanos.ErrorAnalizador{
				Fila:        s.linea,
				Columna:     s.columna,
				Intruccion:  "Switch",
				Descripcion: "Los Tipos de datos no coinciden",
				Tipo:        "Semantico",
			}
			return newError
		}
	}

	return 1
}
