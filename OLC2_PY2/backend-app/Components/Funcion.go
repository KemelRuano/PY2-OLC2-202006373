package Components

import (
	"backend-app/Miceleanos"
	"backend-app/parser"
	"reflect"
)

var RETURNFUNCION bool = false
var AuxPass interface{} = nil

type Parametros struct {
	Variable string
	Tipo     string
	EsInEx   string
	Inout    bool
	EsVector bool
}
type Function struct {
	ID                  string
	Parametros          []Parametros
	Tipo                string
	BloqueInstrucciones interface{}
}
type FuncionLlamada struct {
	ID       string
	isInout  bool
	Valor    interface{}
	ValorRef string
}

func (v *Visitor) VisitFunciones(ctx *parser.FuncionesContext) interface{} {
	Line := ctx.ID().GetSymbol().GetLine()
	Column := ctx.ID().GetSymbol().GetColumn()
	Id := ctx.ID().GetText()
	ListParametros := v.Visit(ctx.Parametros()).([]Parametros)
	Tipo := ""
	if ctx.FLECHA() != nil {
		Tipo = ctx.Tipo().GetText()
	} else {
		Tipo = "void"
	}
	Instrucciones := ctx.Bloque()
	NuevaFuncion := Function{ID: Id, Parametros: ListParametros, Tipo: Tipo, BloqueInstrucciones: Instrucciones}
	VerificacionExistencia := v.EntornoActual.AddFuncion(Id, NuevaFuncion)
	if VerificacionExistencia {
		TaGlobal := NewSimbGlobal("Funcion", Tipo, v.EntornoActual.GetEntorno(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetLine())
		v.AddGlobalSimbol(Id, TaGlobal)
	} else {
		v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetLine(), "Funcion", "Ya existe en el entorno actual: "+Id, "SEMANTICO")
		return false
	}
	if !VerificacionExistencia {
		v.AddError(Line, Column, "FUNCION", "La funcion "+Id+" ya existe en el entorno actual", "Semantico")
		return false
	}
	return true
}

func (v *Visitor) VisitParametros(ctx *parser.ParametrosContext) interface{} {
	NuevoParametro := Parametros{}
	ListParametros := []Parametros{}
	for i := 0; ctx.ID(i) != nil; i++ {
		NuevoParametro.Variable = ctx.ID(i).GetText()
		if ctx.ExisteExInt(i) != nil {
			if ctx.ExisteExInt(i).GetText() != "_" {
				NuevoParametro.EsInEx = ctx.ExisteExInt(i).GetText()
			} else {
				NuevoParametro.EsInEx = "_"
			}
		} else {
			NuevoParametro.EsInEx = "none"
		}
		if ctx.Tipofuncion(i).OPEN_SQUARE_BRACKET() != nil {
			NuevoParametro.EsVector = true
		} else {
			NuevoParametro.EsVector = false
		}

		NuevoParametro.Tipo = ctx.Tipofuncion(i).Tipo().GetText()
		if ctx.Tipoinout(i) != nil {
			NuevoParametro.Inout = true
		} else {
			NuevoParametro.Inout = false
		}
		ListParametros = append(ListParametros, NuevoParametro)
	}
	return ListParametros
}

func (v *Visitor) VisitFuncLLamada(ctx *parser.FuncLLamadaContext) interface{} {
	RETURNFUNCION = false
	var Variable interface{}
	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarFuncion(Id)
	if FuncionVacio(Encontrado) {
		v.EntornoActual = NewEntorno(v.EntornoActual, Id)
		Listcall := v.Visit(ctx.ParametrosLLamada()).([]FuncionLlamada)
		ListParametros := Encontrado.Parametros

		if len(Listcall) != len(ListParametros) {
			Line := ctx.ID().GetSymbol().GetLine()
			Column := ctx.ID().GetSymbol().GetColumn()
			v.AddError(Line, Column, "LLAMADA FUNCION", "La cantidad de parametros no coincide con la funcion", "SEMANTICO")
			return false
		}

		for i := 0; i < len(Listcall); i++ {

			if ListParametros[i].EsInEx == "none" && Listcall[i].ID != "_" {
				if ListParametros[i].Variable == Listcall[i].ID {
					if ListParametros[i].EsVector {
						if ListParametros[i].Tipo == "Int" {
							if ListEnteros(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						} else if ListParametros[i].Tipo == "Float" {
							if ListFloat(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						} else if ListParametros[i].Tipo == "String" {
							if ListString(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						} else if ListParametros[i].Tipo == "Character" {
							if ListCharacter(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						} else if ListParametros[i].Tipo == "Bool" {
							if ListBool(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						}

					} else {
						if ValidarTipo(Listcall[i].Valor, ListParametros[i].Tipo) {
							NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", false)
							v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
						} else {
							Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
							Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
							v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
							return false
						}
					}

				} else {
					line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
					col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
					v.AddError(line, col, "LLAMADA FUNCION", "Los paremtros son incorrectos", "SEMANTICO")
					return false
				}
				// ------------------------------------- Tipo 1 -------------------------------------
				//-----------------------------------------------------------------------------------

			} else if ListParametros[i].EsInEx != "_" && Listcall[i].ID != "_" {
				if ListParametros[i].EsInEx == Listcall[i].ID {
					if ListParametros[i].EsVector {
						if ListParametros[i].Tipo == "Int" {
							if ListEnteros(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						} else if ListParametros[i].Tipo == "Float" {
							if ListFloat(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						} else if ListParametros[i].Tipo == "String" {
							if ListString(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						} else if ListParametros[i].Tipo == "Character" {
							if ListCharacter(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						} else if ListParametros[i].Tipo == "Bool" {
							if ListBool(Listcall[i].Valor.([]interface{})) {
								NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
								v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
							} else {
								Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
								Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
								v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
								return false
							}
						}

					} else {
						if ValidarTipo(Listcall[i].Valor, ListParametros[i].Tipo) {
							NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", false)
							v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)

						} else {
							Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
							Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
							v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
							return false
						}

					}
				} else {
					Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
					Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
					v.AddError(Line, Col, "LLAMADA FUNCION", "El parametro no existe en la funcion", "SEMANITCO")
					return false
				}

				// ------------------------------------- Tipo 2 -------------------------------------
				//-----------------------------------------------------------------------------------

			} else if ListParametros[i].EsInEx == "_" && Listcall[i].ID == "_" {
				if ListParametros[i].EsVector {
					if ListParametros[i].Tipo == "Int" {
						if ListEnteros(Listcall[i].Valor.([]interface{})) {
							NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
							v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
						} else {
							Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
							Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
							v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
							return false
						}
					} else if ListParametros[i].Tipo == "Float" {
						if ListFloat(Listcall[i].Valor.([]interface{})) {
							NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
							v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
						} else {
							Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
							Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
							v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
							return false
						}
					} else if ListParametros[i].Tipo == "String" {
						if ListString(Listcall[i].Valor.([]interface{})) {
							NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
							v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
						} else {
							Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
							Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
							v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
							return false
						}
					} else if ListParametros[i].Tipo == "Character" {
						if ListCharacter(Listcall[i].Valor.([]interface{})) {
							NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
							v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
						} else {
							Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
							Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
							v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
							return false
						}
					} else if ListParametros[i].Tipo == "Bool" {
						if ListBool(Listcall[i].Valor.([]interface{})) {
							NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", true)
							v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
						} else {
							Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
							Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
							v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANTICO")
							return false
						}
					}

				} else {
					if Listcall[i].ValorRef == "none" {
						if ValidarTipo(Listcall[i].Valor, ListParametros[i].Tipo) {
							NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", false)
							v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
						} else {
							Line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
							Col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
							v.AddError(Line, Col, "LLAMADA FUNCION", "El tipo es distinto al de la funcion", "SEMANITCO")
							return false
						}
					} else {
						NuevoSymbol := NewSimbolo(Listcall[i].Valor, ListParametros[i].Tipo, "var", false)
						v.EntornoActual.EnvAddSimbolo(ListParametros[i].Variable, NuevoSymbol)
					}

				}

				// ------------------------------------- Tipo 3 -------------------------------------
				//-----------------------------------------------------------------------------------

			} else {
				line := ctx.ParametrosLLamada().Expresion(i).GetStart().GetLine()
				col := ctx.ParametrosLLamada().Expresion(i).GetStart().GetColumn()
				v.AddError(line, col, "LLAMADA FUNCION", "La llamada no necesita parametro", "SEMANITCO")
				return false
			}

		}

		Recorrer := Encontrado.BloqueInstrucciones.(*parser.BloqueContext)
		for i := 0; Recorrer.Lista_proceso(i) != nil; i++ {
			Variable = v.Visit(Recorrer.Lista_proceso(i))
			if v.EntornoActual.Retorno {
				v.EntornoActual.Retorno = false
				break
			}
		}

		for i := 0; i < len(Listcall); i++ {
			if ListParametros[i].Inout {
				if Listcall[i].isInout {
					SearchSymbol := v.EntornoActual.BuscarSimbolo(ListParametros[i].Variable)
					SearchNew := v.EntornoActual.BuscarSimbolo(Listcall[i].ValorRef)
					SearchNew.Value = SearchSymbol.Value
					v.EntornoActual.ActualizarSimbolo(Listcall[i].ValorRef, SearchNew)
				}
			}

		}
		v.EntornoActual = v.EntornoActual.Padre
	} else {
		Line := ctx.ID().GetSymbol().GetLine()
		Column := ctx.ID().GetSymbol().GetColumn()
		v.AddError(Line, Column, "FUNCION", "La funcion "+Id+" no existe en el entorno actual", "Semantico")
		return false
	}

	return Variable
}

func (v *Visitor) VisitParametrosLLamada(ctx *parser.ParametrosLLamadaContext) interface{} {
	NuevoParametrocall := FuncionLlamada{}
	ListVariable := []FuncionLlamada{}
	Tipo := 0
	for i := 0; ctx.ID(i) != nil; i++ {
		if ctx.ID(i) != nil {
			NuevoParametrocall.ID = ctx.ID(i).GetText()
		}
		if ctx.REFERENCIA(i) != nil {
			NuevoParametrocall.isInout = true
		}
		NuevoParametrocall.Valor = v.Visit(ctx.Expresion(i))
		if AuxPass != nil {
			NuevoParametrocall.ValorRef = AuxPass.(string)
		} else {
			NuevoParametrocall.ValorRef = "none"
		}
		ListVariable = append(ListVariable, NuevoParametrocall)
		Tipo = 1
	}
	if Tipo == 0 {
		for i := 0; ctx.Expresion(i) != nil; i++ {
			NuevoParametrocall.ID = "_"
			if ctx.REFERENCIA(i) != nil {
				NuevoParametrocall.isInout = true
			} else {
				NuevoParametrocall.isInout = false
			}
			NuevoParametrocall.Valor = v.Visit(ctx.Expresion(i))
			if AuxPass != nil {
				NuevoParametrocall.ValorRef = AuxPass.(string)
			} else {
				NuevoParametrocall.ValorRef = "none"
			}
			AuxPass = nil
			ListVariable = append(ListVariable, NuevoParametrocall)
		}
	}
	return ListVariable
}

func (v *Visitor) VisitRetornar(ctx *parser.RetornarContext) interface{} {
	// v.Retorno = true
	// RET_SET = true
	// RETURNFUNCION = true
	var valor interface{}
	if ctx.BREAK() != nil {
		v.Break = true
		valor = "break"
	} else if ctx.RETURN() != nil {
		if ctx.Expresion() != nil {
			v.EntornoActual.Retorno = true
			valor = v.Visit(ctx.Expresion())

			return valor
		}

	} else if ctx.CONTINUE() != nil {
		valor = "continue"
	}
	return valor
}

func (v *Visitor) VisitPrintLLamada(ctx *parser.PrintLLamadaContext) interface{} {
	return v.Visit(ctx.FuncLLamada())
}

func ValidarTipo(valor interface{}, tipo string) bool {
	if tipo == "Int" {
		if Miceleanos.IsInt(valor) {
			return true
		}
	} else if tipo == "Float" {
		if Miceleanos.IsFloat(valor) {
			return true
		}
	} else if tipo == "String" {
		if Miceleanos.IsString(valor) {
			return true
		}
	} else if tipo == "Bool" {
		if Miceleanos.Isbool(valor) {
			return true
		}
	} else if tipo == "Character" {
		if Miceleanos.IsChar(valor) {
			return true
		}
	}
	return false
}

func ListEnteros(lista []interface{}) bool {
	for _, elemento := range lista {
		if reflect.TypeOf(elemento).Kind() != reflect.Int64 {
			return false
		}
	}
	return true
}

func ListFloat(lista []interface{}) bool {
	for _, elemento := range lista {
		if reflect.TypeOf(elemento).Kind() != reflect.Float64 {
			return false
		}
	}
	return true
}

func ListString(lista []interface{}) bool {
	for _, elemento := range lista {
		if reflect.TypeOf(elemento).Kind() != reflect.String {
			return false
		}
	}
	return true
}
func ListCharacter(lista []interface{}) bool {
	for _, elemento := range lista {
		if reflect.TypeOf(elemento).Kind() != reflect.String || len(elemento.(string)) > 1 {
			return false
		}
	}
	return true
}

func ListBool(lista []interface{}) bool {
	for _, elemento := range lista {
		if reflect.TypeOf(elemento).Kind() != reflect.Bool {
			return false
		}
	}
	return true
}
