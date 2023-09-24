package Proceso

import (
	"backend-app/Miceleanos"
)

type Vector struct {
	TipoVariable string
	Id           string
	List         []interface{}
	TipoDate     string
}

func NewVector() *Vector {
	return &Vector{}
}
func (v *Vector) VectorAppend(valor interface{}) bool {
	if v.TipoDate == "Int" {
		if Miceleanos.IsInt(valor) {
			v.List = append(v.List, valor.(int64))
			return true
		}
	} else if v.TipoDate == "Float" {
		if Miceleanos.IsFloat(valor) {
			v.List = append(v.List, valor.(float64))
			return true
		}
	} else if v.TipoDate == "Bool" {
		if Miceleanos.Isbool(valor) {
			v.List = append(v.List, valor.(bool))
			return true
		}
	} else if v.TipoDate == "String" {
		if Miceleanos.IsString(valor) {
			v.List = append(v.List, valor.(string))
			return true
		}
	} else if v.TipoDate == "Character" {
		if Miceleanos.IsString(valor) && len(valor.(string)) < 2 {
			v.List = append(v.List, valor.(string))
			return true
		}
	}

	return false
}

func (v *Vector) GetList() interface{} {
	return v.List
}

func (v *Vector) GetTipo() string {
	return v.TipoDate
}
func (v *Vector) GetId() string {
	return v.Id
}
func (v *Vector) GetTipoVariable() string {
	return v.TipoVariable
}

func (v *Vector) FunctionAppend(valor interface{}) bool {
	if v.TipoVariable == "var" {
		if v.TipoDate == "Int" {
			if Miceleanos.IsInt(valor) {
				v.List = append(v.List, valor.(int64))
				return true
			}
		} else if v.TipoDate == "Float" {
			if Miceleanos.IsFloat(valor) {
				v.List = append(v.List, valor.(float64))
				return true
			}
		} else if v.TipoDate == "Bool" {
			if Miceleanos.Isbool(valor) {
				v.List = append(v.List, valor.(bool))
				return true
			}
		} else if v.TipoDate == "String" {
			if Miceleanos.IsString(valor) {
				v.List = append(v.List, valor.(string))
				return true
			}
		} else if v.TipoDate == "Character" {
			if Miceleanos.IsString(valor) && len(valor.(string)) < 2 {
				v.List = append(v.List, valor.(string))
				return true
			}
		}
	}
	return false
}

func (v *Vector) RemoveLast() bool {
	if len(v.List) > 0 {
		v.List = v.List[:len(v.List)-1]
		return true
	}
	return false
}

func (v *Vector) Remove(buscar interface{}) interface{} {

	if !Miceleanos.IsInt(buscar) {
		return false
	}

	if buscar.(int64) >= 0 && buscar.(int64) < int64(len(v.List)) {
		v.List = append(v.List[:buscar.(int64)], v.List[buscar.(int64)+1:]...)
		return true
	}
	return nil
}

func (v *Vector) IsVacio() bool {

	return len(v.List) == 0
}
func (v *Vector) Count() int64 {

	return int64(len(v.List))
}
