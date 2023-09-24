package Proceso

import "backend-app/Miceleanos"

type Declaracion struct {
	Variable string
	Tipo     string
	Valor    interface{}
}

func NewDeclaracion(Variable string, Tipo string, Valor interface{}) *Declaracion {
	return &Declaracion{
		Variable: Variable,
		Tipo:     Tipo,
		Valor:    Valor,
	}
}

func (d *Declaracion) EsTipo() bool {
	switch d.Tipo {
	case "Int":
		if Miceleanos.IsNull(d.Valor) {
			d.Valor = nil
			return true
		} else if Miceleanos.IsInt(d.Valor) {
			return true
		}

	case "Float":
		if Miceleanos.IsNull(d.Valor) {
			d.Valor = nil
			return true
		} else if Miceleanos.IsFloat(d.Valor) {
			return true
		}

	case "Bool":
		if Miceleanos.IsNull(d.Valor) {
			d.Valor = nil
			return true
		} else if Miceleanos.Isbool(d.Valor) {
			return true
		}

	case "String":
		if Miceleanos.IsNull(d.Valor) {
			d.Valor = nil
			return true
		} else if Miceleanos.IsString(d.Valor) {
			return true
		}
	case "Character":
		if Miceleanos.IsNull(d.Valor) {
			d.Valor = nil
			return true
		} else if Miceleanos.IsString(d.Valor) && len(d.Valor.(string)) < 2 {
			return true
		}
	}
	return false
}

func (d *Declaracion) EsSinTipo() bool {
	switch (d.Valor).(type) {
	case int64:
		d.Tipo = "Int"

		return true
	case float64:
		d.Tipo = "Float"
		return true
	case bool:
		d.Tipo = "Bool"
		return true
	case string:
		d.Tipo = "String"
		return true

	}
	return false
}

func (d *Declaracion) GetValor() interface{} {
	return d.Valor
}

func (d *Declaracion) GetTipo() string {
	return d.Tipo
}

func (d *Declaracion) GetVariable() string {
	return d.Variable
}
