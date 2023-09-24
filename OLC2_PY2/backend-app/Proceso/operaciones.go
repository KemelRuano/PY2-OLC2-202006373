package Proceso

import (
	"backend-app/Miceleanos"
	"fmt"
)

type Concatenar struct {
	derecha       interface{}
	izquierda     interface{}
	TipoOperacion string
}

func NewConcatenar(derecha interface{}, izquierda interface{}, ti string) *Concatenar {
	return &Concatenar{
		derecha:       derecha,
		izquierda:     izquierda,
		TipoOperacion: ti,
	}
}

func (c *Concatenar) Aritmeticos() interface{} {
	if Miceleanos.IsInt(c.derecha) && Miceleanos.IsInt(c.izquierda) {
		switch c.TipoOperacion {
		case "+":
			return c.derecha.(int64) + c.izquierda.(int64)
		case "-":
			return c.derecha.(int64) - c.izquierda.(int64)
		case "*":
			return c.derecha.(int64) * c.izquierda.(int64)
		case "/":
			return c.derecha.(int64) / c.izquierda.(int64)
		case "%":
			return c.derecha.(int64) % c.izquierda.(int64)
		}
	} else if Miceleanos.IsInt(c.derecha) && Miceleanos.IsFloat(c.izquierda) {
		switch c.TipoOperacion {
		case "+":
			return float64(c.derecha.(int64)) + c.izquierda.(float64)
		case "-":
			return float64(c.derecha.(int64)) - c.izquierda.(float64)
		case "*":
			return float64(c.derecha.(int64)) * c.izquierda.(float64)
		case "/":
			return float64(c.derecha.(int64)) / c.izquierda.(float64)
		}
	} else if Miceleanos.IsFloat(c.derecha) && Miceleanos.IsInt(c.izquierda) {

		switch c.TipoOperacion {
		case "+":
			return c.derecha.(float64) + float64(c.izquierda.(int64))
		case "-":
			return c.derecha.(float64) - float64(c.izquierda.(int64))
		case "*":
			return c.derecha.(float64) * float64(c.izquierda.(int64))
		case "/":
			return c.derecha.(float64) / float64(c.izquierda.(int64))
		}
	} else if Miceleanos.IsFloat(c.derecha) && Miceleanos.IsFloat(c.izquierda) {

		switch c.TipoOperacion {
		case "+":
			return c.derecha.(float64) + c.izquierda.(float64)
		case "-":
			return c.derecha.(float64) - c.izquierda.(float64)
		case "*":
			return c.derecha.(float64) * c.izquierda.(float64)
		case "/":
			return c.derecha.(float64) / c.izquierda.(float64)
		}
	} else if Miceleanos.IsString(c.derecha) && Miceleanos.IsString(c.izquierda) {
		switch c.TipoOperacion {
		case "+":
			return c.derecha.(string) + c.izquierda.(string)
		}
	}

	return nil
}

func (c *Concatenar) Comparativos() interface{} {

	if Miceleanos.IsInt(c.derecha) && Miceleanos.IsInt(c.izquierda) {

		switch c.TipoOperacion {
		case "==":
			return c.derecha.(int64) == c.izquierda.(int64)
		case "!=":
			return c.derecha.(int64) != c.izquierda.(int64)
		}
	} else if Miceleanos.IsFloat(c.derecha) && Miceleanos.IsFloat(c.izquierda) {

		switch c.TipoOperacion {
		case "==":
			return c.derecha.(float64) == c.izquierda.(float64)
		case "!=":
			return c.derecha.(float64) != c.izquierda.(float64)
		}
	} else if Miceleanos.IsString(c.derecha) && Miceleanos.IsString(c.izquierda) {

		switch c.TipoOperacion {
		case "==":
			return c.derecha.(string) == c.izquierda.(string)
		case "!=":
			return c.derecha.(string) != c.izquierda.(string)
		}
	} else if Miceleanos.Isbool(c.derecha) && Miceleanos.Isbool(c.izquierda) {

		switch c.TipoOperacion {
		case "==":
			return c.derecha.(bool) == c.izquierda.(bool)
		case "!=":
			return c.derecha.(bool) != c.izquierda.(bool)
		}
	}
	return nil
}

func (c *Concatenar) Relacionales() interface{} {

	if Miceleanos.IsInt(c.derecha) && Miceleanos.IsInt(c.izquierda) {
		switch c.TipoOperacion {
		case ">":
			return c.derecha.(int64) > c.izquierda.(int64)
		case "<":
			return c.derecha.(int64) < c.izquierda.(int64)
		case ">=":
			return c.derecha.(int64) >= c.izquierda.(int64)
		case "<=":
			return c.derecha.(int64) <= c.izquierda.(int64)
		}
	} else if Miceleanos.IsFloat(c.derecha) && Miceleanos.IsFloat(c.izquierda) {
		switch c.TipoOperacion {
		case ">":
			return c.derecha.(float64) > c.izquierda.(float64)
		case "<":
			return c.derecha.(float64) < c.izquierda.(float64)
		case ">=":
			return c.derecha.(float64) >= c.izquierda.(float64)
		case "<=":
			return c.derecha.(float64) <= c.izquierda.(float64)
		}
	} else if Miceleanos.IsString(c.derecha) && Miceleanos.IsString(c.izquierda) {
		if len(c.derecha.(string)) < 2 && len(c.izquierda.(string)) < 2 {
			switch c.TipoOperacion {
			case ">":
				return c.derecha.(string) > c.izquierda.(string)
			case "<":
				return c.derecha.(string) < c.izquierda.(string)
			case ">=":
				return c.derecha.(string) >= c.izquierda.(string)
			case "<=":
				fmt.Print("esL: ", c.derecha.(string) <= c.izquierda.(string))
				return c.derecha.(string) <= c.izquierda.(string)
			}
		} else {
			switch c.TipoOperacion {
			case ">":
				return c.derecha.(string) > c.izquierda.(string)
			case "<":
				return c.derecha.(string) < c.izquierda.(string)
			case ">=":
				return c.derecha.(string) >= c.izquierda.(string)
			case "<=":
				return c.derecha.(string) <= c.izquierda.(string)
			}

		}

	}

	return nil
}

func (c *Concatenar) Logicos() interface{} {

	if Miceleanos.Isbool(c.derecha) && Miceleanos.Isbool(c.izquierda) {
		switch c.TipoOperacion {
		case "&&":
			return c.derecha.(bool) && c.izquierda.(bool)
		case "||":
			return c.derecha.(bool) || c.izquierda.(bool)
		}
	}
	return nil
}

func (c *Concatenar) Reasignacion() interface{} {

	if Miceleanos.IsInt(c.derecha) && Miceleanos.IsInt(c.izquierda) {
		switch c.TipoOperacion {
		case "=":
			return c.izquierda.(int64)
		case "+=":
			return c.derecha.(int64) + c.izquierda.(int64)
		case "-=":
			return c.derecha.(int64) - c.izquierda.(int64)
		default:
			return '1'
		}
	} else if Miceleanos.IsFloat(c.derecha) && Miceleanos.IsFloat(c.izquierda) {
		switch c.TipoOperacion {
		case "=":
			return c.izquierda.(float64)
		case "+=":
			return c.derecha.(float64) + c.izquierda.(float64)
		case "-=":
			return c.derecha.(float64) - c.izquierda.(float64)
		default:
			return '1'
		}
	} else if Miceleanos.IsString(c.derecha) && Miceleanos.IsString(c.izquierda) {
		switch c.TipoOperacion {
		case "=":
			return c.izquierda.(string)
		case "+=":
			return c.derecha.(string) + c.izquierda.(string)
		default:
			return '1'
		}
	} else if Miceleanos.Isbool(c.derecha) && Miceleanos.Isbool(c.izquierda) {
		switch c.TipoOperacion {
		case "=":
			return c.izquierda.(bool)
		default:
			return '1'
		}
	}

	return nil
}
