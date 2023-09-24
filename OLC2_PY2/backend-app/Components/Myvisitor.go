package Components

import (
	"backend-app/Miceleanos"
	"backend-app/Proceso"
	"backend-app/parser"
	"fmt"

	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

var TypeDate string = ""
var RET_SET bool = false
var RET_FUN bool = false

var RETVALORDEELSE bool = false

// Control de sentencia if
var EnIF bool = false
var EnElseIf bool = false
var EnElse bool = false

type Visitor struct {
	*parser.BaseSintaxVisitor
	EntornoActual *Entorno
	ListErrores   []*Miceleanos.ErrorAnalizador
	TableGlobal   map[string]GlobalSymbol
	ListConsole   []string
	Retorno       bool
	Break         bool
	Continue      bool
}

func NewVisitor(ent *Entorno) *Visitor {
	return &Visitor{
		EntornoActual: ent,
		ListErrores:   []*Miceleanos.ErrorAnalizador{},
		TableGlobal:   make(map[string]GlobalSymbol),
	}
}

func (v *Visitor) AddError(Li int, Co int, In string, De string, Ti string) {
	newError := &Miceleanos.ErrorAnalizador{
		Fila:        Li,
		Columna:     Co,
		Intruccion:  In,
		Descripcion: De,
		Tipo:        Ti,
	}
	v.ListErrores = append(v.ListErrores, newError)
}
func (v *Visitor) AddGlobalSimbol(id string, value GlobalSymbol) {
	_, ok := v.TableGlobal[id]
	if !ok {
		v.TableGlobal[id] = value
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch t := tree.(type) {
	case *parser.InicioContext:
		return v.VisitInicio(t)

	case *parser.ListaContext:
		return v.VisitLista(t)

	case *parser.Lista_procesoContext:
		return v.VisitLista_proceso(t)

	case *parser.DeclaracionContext:
		return v.VisitDeclaracion(t)

	case *parser.AsignacionContext:
		return v.VisitAsignacion(t)

	case *parser.WhileContext:
		return v.VisitWhile(t)

	case *parser.SwitchContext:
		return v.VisitSwitch(t)

	case *parser.PrintContext:
		return v.VisitPrint(t)

	case *parser.IfContext:
		return v.VisitIf(t)
	case *parser.ForContext:
		return v.VisitFor(t)
	case *parser.GuardContext:
		return v.VisitGuard(t)
	case *parser.FuncvectorListContext:
		return v.VisitFuncvectorList(t)

	case *parser.FuncionesContext:
		return v.VisitFunciones(t)

	case *parser.ParametrosContext:
		return v.VisitParametros(t)

	case *parser.FuncLLamadaContext:
		return v.VisitFuncLLamada(t)

	case *parser.ParametrosLLamadaContext:
		return v.VisitParametrosLLamada(t)

	case *parser.RetornarContext:
		return v.VisitRetornar(t)

	case *parser.PrintLLamadaContext:
		return v.VisitPrintLLamada(t)
	case *parser.CasteoContext:
		return v.VisitCasteo(t)

	case *parser.EstructsContext:
		return v.VisitEstructs(t)
		//-------------------------------------------
	case *parser.AsignacionVariableContext:
		return v.VisitAsignacionVariable(t)
	case *parser.AsignacionVectorContext:
		return v.VisitAsignacionVector(t)
	case *parser.Asigtipo1Context:
		return v.VisitAsigtipo1(t)

	case *parser.Asigtipo2Context:
		return v.VisitAsigtipo2(t)

	case *parser.NumberValorContext:
		return v.VisitNumberValor(t)
	case *parser.BoolValorContext:
		return v.VisitBoolValor(t)
	case *parser.StringValorContext:
		return v.VisitStringValor(t)
	case *parser.SinValorContext:
		return v.VisitSinValor(t)
	case *parser.ParentexprContext:
		return v.VisitParentexpr(t)
	case *parser.IdValorContext:
		return v.VisitIdValor(t)

	case *parser.OpAritmeticoContext:
		return v.VisitOpAritmetico(t)
	case *parser.OpComparativoContext:
		return v.VisitOpComparativo(t)
	case *parser.OpRelacionalContext:
		return v.VisitOpRelacional(t)
	case *parser.OpLogicoContext:
		return v.VisitOpLogico(t)
	case *parser.OpLogicoNotContext:
		return v.VisitOpLogicoNot(t)

	case *parser.Elseif_Context:
		return v.VisitElseif_(t)

	case *parser.CasoContext:
		return v.VisitCaso(t)
	case *parser.DefaultContext:
		return v.VisitDefault(t)

	case *parser.Dec_vectorContext:
		return v.VisitDec_vector(t)
	case *parser.VectorCountContext:
		return v.VisitVectorCount(t)
	case *parser.VectorVacioContext:
		return v.VisitVectorVacio(t)
	case *parser.Dec_vector_V_CContext:
		return v.VisitDec_vector_V_C(t)
	case *parser.VectorAsignacionContext:
		return v.VisitVectorAsignacion(t)
	case *parser.BloqueContext:
		return v.VisitBloque(t)

	case *parser.ValorsimpleContext:
		return v.VisitValorsimple(t)

	default:
		return false
	}
}
func (v *Visitor) VisitInicio(ctx *parser.InicioContext) interface{} {
	v.Visit(ctx.Lista())
	return true
}

func (v *Visitor) VisitLista(ctx *parser.ListaContext) interface{} {

	for i := 0; ctx.Funciones(i) != nil; i++ {
		v.Visit(ctx.Funciones(i))
	}

	for i := 0; ctx.Lista_proceso(i) != nil; i++ {
		v.Visit(ctx.Lista_proceso(i))
	}
	return true
}

func (v *Visitor) VisitLista_proceso(ctx *parser.Lista_procesoContext) interface{} {

	if ctx.Declaracion() != nil {
		return v.Visit(ctx.Declaracion())
	}

	if ctx.If_() != nil {
		vart := v.Visit(ctx.If_())
		return vart
	}
	if ctx.Switch_() != nil {
		return v.Visit(ctx.Switch_())
	}
	if ctx.While() != nil {
		return v.Visit(ctx.While())
	}
	if ctx.Asignacion() != nil {
		return v.Visit(ctx.Asignacion())

	}
	if ctx.For_() != nil {
		return v.Visit(ctx.For_())
	}
	if ctx.Guard() != nil {
		return v.Visit(ctx.Guard())
	}
	if ctx.FuncvectorList() != nil {
		return v.Visit(ctx.FuncvectorList())
	}
	if ctx.FuncLLamada() != nil {
		return v.Visit(ctx.FuncLLamada())
	}
	if ctx.Retornar() != nil {
		return v.Visit(ctx.Retornar())
	}

	if ctx.Print_() != nil {
		return v.Visit(ctx.Print_())
	}

	if ctx.Estructs() != nil {
		return v.Visit(ctx.Estructs())
	}
	return nil
}

func (v *Visitor) VisitDeclaracion(ctx *parser.DeclaracionContext) interface{} {

	if ctx.Dec_tipo_valor() != nil {
		v.Visit(ctx.Dec_tipo_valor())
	}
	if ctx.Dec_valor() != nil {
		v.Visit(ctx.Dec_valor())
	}
	if ctx.Dec_vector() != nil {
		v.Visit(ctx.Dec_vector())
	}
	if ctx.Dec_vector_V_C() != nil {
		v.Visit(ctx.Dec_vector_V_C())
	}
	return true
}

// ------------------------------------ DECLARACION DE VARIABLES ----------------------------------------
//------------------------------------------------------------------------------------------------------

func (v *Visitor) VisitAsigtipo1(ctx *parser.Asigtipo1Context) interface{} {
	TypeDate = ctx.Tipo().GetText()
	Line := ctx.Expresion().GetStart().GetLine()
	Col := ctx.Expresion().GetStart().GetColumn()
	Valor := v.Visit(ctx.Expresion())
	if Valor == false {
		return false
	}

	Dec := Proceso.NewDeclaracion(ctx.ID().GetText(), ctx.Tipo().GetText(), Valor)
	CorrectType := Dec.EsTipo()

	if CorrectType {
		CreateSimbol := NewSimbolo(Dec.GetValor(), Dec.GetTipo(), ctx.GetOp().GetText(), false)
		is_exist := v.EntornoActual.EnvAddSimbolo(Dec.GetVariable(), CreateSimbol)
		if is_exist {
			TaGlobal := NewSimbGlobal("Variable", Dec.GetTipo(), v.EntornoActual.GetEntorno(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetLine())
			v.AddGlobalSimbol(Dec.GetVariable(), TaGlobal)
		} else {
			v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetLine(), "DECLARACION", "Ya existe en el entorno actual: "+Dec.GetVariable(), "SEMANTICO")
			return false
		}
	} else {
		v.AddError(Line, Col, "DECLARACION", "Especifico "+Dec.GetTipo()+"()", "SEMANTICO")
		return false
	}
	return true
}

func (v *Visitor) VisitAsigtipo2(ctx *parser.Asigtipo2Context) interface{} {
	Line := ctx.Expresion().GetStart().GetLine()
	Col := ctx.Expresion().GetStart().GetColumn()
	TypeDate = ""
	Dec := Proceso.NewDeclaracion(ctx.ID().GetText(), "", v.Visit(ctx.Expresion()))
	CorrectType := Dec.EsSinTipo()
	if CorrectType {
		CreateSimbol := NewSimbolo(Dec.GetValor(), Dec.GetTipo(), ctx.GetOp().GetText(), false)
		is_exist := v.EntornoActual.EnvAddSimbolo(Dec.GetVariable(), CreateSimbol)
		if is_exist {
			TaGlobal := NewSimbGlobal("Variable", Dec.GetTipo(), v.EntornoActual.GetEntorno(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
			v.AddGlobalSimbol(Dec.GetVariable(), TaGlobal)
		} else {
			v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "DECLARACION", "Ya existe en el entorno actual: "+Dec.GetVariable(), "SEMANTICO")
			return false
		}
	} else {
		v.AddError(Line, Col, "DECLARACION", "Especifico "+Dec.GetTipo()+"()", "SEMANTICO")
		return false
	}
	return true
}

// ------------------------------------ Asignacion de Variables ----------------------------------------

func (v *Visitor) VisitAsignacion(ctx *parser.AsignacionContext) interface{} {

	if ctx.AsignacionVariable() != nil {
		return v.Visit(ctx.AsignacionVariable())
	}
	if ctx.AsignacionVector() != nil {
		return v.Visit(ctx.AsignacionVector())
	}
	return true
}

func (v *Visitor) VisitAsignacionVariable(ctx *parser.AsignacionVariableContext) interface{} {
	Line := ctx.ID().GetSymbol().GetLine()
	Col := ctx.ID().GetSymbol().GetColumn()
	Encontrado := v.EntornoActual.BuscarSimbolo(ctx.ID().GetText())

	if Encontrado.Value != nil {
		if Encontrado.TypeValue == "var" {
			newExpresion := v.Visit(ctx.Expresion())
			NewValor := Proceso.NewConcatenar(Encontrado.Value, newExpresion, ctx.GetOp().GetText()).Reasignacion()
			if Miceleanos.IsChar(NewValor) {
				v.AddError(ctx.GetOp().GetLine(), ctx.GetOp().GetColumn(), "ASIGNACION", "Tipo: "+Encontrado.Type+" No permite: "+ctx.GetOp().GetText(), "SEMANTICO")
				return false
			} else if NewValor == nil {
				v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "ASIGNACION", "Tipo esperado: "+Encontrado.Type, "SEMANTICO")
			} else {
				Encontrado.Value = NewValor
				v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), Encontrado)
			}

		} else {

			v.AddError(Line, Col, "ASIGNACION", "Let es constante no se puede modificar", "SEMANTICO")
			return false
		}

	} else {
		v.AddError(Line, Col, "ASIGNACION", "Variable ["+ctx.ID().GetText()+"] No existe", "SEMANTICO")
		return false
	}
	return true
}

func (v *Visitor) VisitNumberValor(ctx *parser.NumberValorContext) interface{} {

	Entero, error := strconv.ParseInt(ctx.GetText(), 10, 64)
	if error == nil {
		return Entero
	}

	Decimal, error := strconv.ParseFloat(ctx.GetText(), 64)
	if error == nil {
		return Decimal
	}

	return false
}

func (v *Visitor) VisitBoolValor(ctx *parser.BoolValorContext) interface{} {
	boleano, _ := strconv.ParseBool(ctx.GetText())
	return boleano
}

func (v *Visitor) VisitStringValor(ctx *parser.StringValorContext) interface{} {
	value := strings.Trim(ctx.GetText(), "\"")
	return value
}

func (v *Visitor) VisitSinValor(ctx *parser.SinValorContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitValorsimple(ctx *parser.ValorsimpleContext) interface{} {

	valor := v.Visit(ctx.NULO())
	return valor
}

/*
------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------
                                		Operadores
*/

func (v *Visitor) VisitOpAritmetico(ctx *parser.OpAritmeticoContext) interface{} {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	ret := Proceso.NewConcatenar(l, r, op).Aritmeticos()
	return ret
}

func (v *Visitor) VisitParentexpr(ctx *parser.ParentexprContext) interface{} {
	visitValor := v.Visit(ctx.Expresion())
	return visitValor
}

func (v *Visitor) VisitOpComparativo(ctx *parser.OpComparativoContext) interface{} {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	return Proceso.NewConcatenar(l, r, op).Comparativos()
}

func (v *Visitor) VisitOpRelacional(ctx *parser.OpRelacionalContext) interface{} {

	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	return Proceso.NewConcatenar(l, r, op).Relacionales()
}

func (v *Visitor) VisitOpLogico(ctx *parser.OpLogicoContext) interface{} {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	return Proceso.NewConcatenar(l, r, op).Logicos()
}
func (v *Visitor) VisitOpLogicoNot(ctx *parser.OpLogicoNotContext) interface{} {
	l := v.Visit(ctx.Expresion())
	if Miceleanos.Isbool(l) {
		l = !l.(bool)
	}
	return l
}

func (v *Visitor) VisitPrint(ctx *parser.PrintContext) interface{} {
	var Concatenar string = ""
	for i := 0; ctx.Expresion(i) != nil; i++ {
		Concatenar += fmt.Sprint(v.Visit(ctx.Expresion(i)))
	}
	// fmt.Println("CONSOLE: ", Concatenar)
	Concatenar += "\n"
	v.ListConsole = append(v.ListConsole, Concatenar)
	return true
}
func (v *Visitor) VisitIdValor(ctx *parser.IdValorContext) interface{} {
	id := ctx.GetText()
	value := v.EntornoActual.BuscarSimbolo(id)
	AuxPass = id
	if value.Value == nil {
		v.ListErrores = append(v.ListErrores, &Miceleanos.ErrorAnalizador{
			Fila:        ctx.GetStart().GetLine(),
			Columna:     ctx.GetStart().GetColumn(),
			Intruccion:  "Declaracion",
			Descripcion: "No existe la variable [" + id + "]",
			Tipo:        "SEMANTICO",
		})
		return false
	}
	return value.Value
}

func (v *Visitor) VisitIf(ctx *parser.IfContext) interface{} {
	EnIF = false
	EnElseIf = false
	RET_SET = false
	condicion := v.Visit(ctx.Expresion())
	if !Miceleanos.Isbool(condicion) {
		v.ListErrores = append(v.ListErrores, &Miceleanos.ErrorAnalizador{
			Fila:        ctx.Expresion().GetStart().GetLine(),
			Columna:     ctx.Expresion().GetStart().GetColumn(),
			Intruccion:  "IF",
			Descripcion: "La condicion no es booleana",
			Tipo:        "SEMANTICO",
		})
		return false
	}
	if !condicion.(bool) {

		EnIF = true
		v.EntornoActual = NewEntorno(v.EntornoActual, "If")
		for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
			value := v.Visit(ctx.Bloque().Lista_proceso(i))
			if v.Break {
				v.Break = false
				v.EntornoActual = v.EntornoActual.Padre
				return "break"
			} else if v.Continue {
				v.Continue = false
				v.EntornoActual = v.EntornoActual.Padre
				return "continue"
			} else if v.EntornoActual.Retorno {
				v.EntornoActual = v.EntornoActual.Padre
				v.EntornoActual.Retorno = true
				return value
			}

		}
		v.EntornoActual = v.EntornoActual.Padre
	}
	if EnIF {
		return true
	}

	for i := 0; ctx.Elseif_(i) != nil; i++ {
		ResElif := v.Visit(ctx.Elseif_(i))
		if ResElif == "break" {
			return "break"
		} else if ResElif == "continue" {
			return "continue"
		} else if RETVALORDEELSE {
			RETVALORDEELSE = false
			v.EntornoActual.Retorno = true
			return ResElif
		}
	}

	if EnElseIf {
		return true
	}

	if ctx.Else_() != nil {
		if !condicion.(bool) {
			v.EntornoActual = NewEntorno(v.EntornoActual, "ELSE")
			for i := 0; ctx.Else_().Bloque().Lista_proceso(i) != nil; i++ {
				value := v.Visit(ctx.Else_().Bloque().Lista_proceso(i))
				if v.Break {
					v.Break = false
					v.EntornoActual = v.EntornoActual.Padre
					return "break"
				} else if v.Continue {
					v.Continue = false
					v.EntornoActual = v.EntornoActual.Padre
					return "continue"
				} else if v.EntornoActual.Retorno {
					v.EntornoActual = v.EntornoActual.Padre
					return value
				}

			}
			v.EntornoActual = v.EntornoActual.Padre
		}
	}
	return nil
}

func (v *Visitor) VisitElseif_(ctx *parser.Elseif_Context) interface{} {

	condicion := v.Visit(ctx.Expresion())
	if Miceleanos.Isbool(condicion) {

		if condicion.(bool) {
			EnElseIf = true
			v.EntornoActual = NewEntorno(v.EntornoActual, "ElseIf")
			for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
				value := v.Visit(ctx.Bloque().Lista_proceso(i))
				if v.Break {
					v.Break = false
					v.EntornoActual = v.EntornoActual.Padre
					return "break"
				} else if v.Continue {
					v.Continue = false
					v.EntornoActual = v.EntornoActual.Padre
					return "continue"
				} else if v.EntornoActual.Retorno {
					RETVALORDEELSE = true
					v.EntornoActual = v.EntornoActual.Padre
					return value
				}
			}
			v.EntornoActual = v.EntornoActual.Padre
		}
	} else {
		v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "ELSEIF", "La condicion no es booleana", "SEMANTICO")
		return false
	}

	return false
}

func (v *Visitor) VisitSwitch(ctx *parser.SwitchContext) interface{} {
	condicion := v.Visit(ctx.Expresion())
	aux := false
	for i := 0; ctx.Caso(i) != nil; i++ {
		NewSwitch := Proceso.NewSwitch(condicion, v.Visit(ctx.Caso(i).Expresion()))
		if Miceleanos.Isbool(NewSwitch.IsEqual()) {
			if NewSwitch.IsEqual().(bool) {
				aux = true
				v.EntornoActual = NewEntorno(v.EntornoActual, "Switch")
				v.Visit(ctx.Caso(i))
				v.EntornoActual = v.EntornoActual.Padre
				break
			}
		} else {
			NewError := NewSwitch.IsEqual().(Miceleanos.ErrorAnalizador)
			v.AddError(NewError.Fila, NewError.Columna, NewError.Intruccion, NewError.Descripcion, NewError.Tipo)
		}
	}
	if ctx.Default_() != nil {
		if !aux {
			v.EntornoActual = NewEntorno(v.EntornoActual, "Switch")
			v.Visit(ctx.Default_())
			v.EntornoActual = v.EntornoActual.Padre
		}
	}

	return false
}
func (v *Visitor) VisitCaso(ctx *parser.CasoContext) interface{} {

	for i := 0; ctx.Lista_proceso(i) != nil; i++ {
		v.Visit(ctx.Lista_proceso(i))
	}
	return false
}

func (v *Visitor) VisitDefault(ctx *parser.DefaultContext) interface{} {
	for i := 0; ctx.Lista_proceso(i) != nil; i++ {
		v.Visit(ctx.Lista_proceso(i))
	}
	return false
}

//-------------------------------------------------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------   while

func (v *Visitor) VisitWhile(ctx *parser.WhileContext) interface{} {

	condicion := v.Visit(ctx.Expresion())
	if Miceleanos.Isbool(condicion) {
	outerLoop:
		for condicion.(bool) {
			v.EntornoActual = NewEntorno(v.EntornoActual, "While")
			for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
				ResW := v.Visit(ctx.Bloque().Lista_proceso(i))
				if ResW == "break" {
					break outerLoop
				} else if ResW == "continue" {
					condicion = v.Visit(ctx.Expresion())
					continue outerLoop
				}
			}

			condicion = v.Visit(ctx.Expresion())
			v.EntornoActual = v.EntornoActual.Padre
		}
	} else {
		v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "WHILE", "La condicion no es booleana", "SEMANTICO")
	}

	return true
}

//-------------------------------------------------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------   for

func (v *Visitor) VisitFor(ctx *parser.ForContext) interface{} {
	izq := v.Visit(ctx.GetLeft())
	der := interface{}(nil)
	if ctx.GetRight() != nil {
		der = v.Visit(ctx.GetRight())
	}
	var comprobacion bool = true
	if ctx.GUION_BAJO() != nil {
		comprobacion = false
	}
	if ctx.TRESPUNTOS() != nil {
		if Miceleanos.IsInt(izq) && Miceleanos.IsInt(der) {
			if izq.(int64) <= der.(int64) {
				if comprobacion {
					// for tipo 1
					for izq.(int64) <= der.(int64) {

						CreateSymbol := NewSimbolo(izq, "Int", "Let", false)
						v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)
						v.EntornoActual = NewEntorno(v.EntornoActual, "For")
						for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
							newReturn := v.Visit(ctx.Bloque().Lista_proceso(i))
							if newReturn == false {
								v.EntornoActual = v.EntornoActual.Padre
								return false
							}
						}
						izq = izq.(int64) + 1
						v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(izq, "Int", "Let", false))
						v.EntornoActual = v.EntornoActual.Padre
					}

				} else {
					// for tipo 2
					for izq.(int64) <= der.(int64) {
						v.EntornoActual = NewEntorno(v.EntornoActual, "For")
						for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
							v.Visit(ctx.Bloque().Lista_proceso(i))
						}
						izq = izq.(int64) + 1
						v.EntornoActual = v.EntornoActual.Padre
					}

				}

			}
		} else {
			v.AddError(ctx.GetLeft().GetStart().GetLine(), ctx.GetLeft().GetStart().GetColumn(), "FOR", "Tipos valido Int()", "SEMANTICO")
			return false
		}

	} else {
		//buscar si es una cadena o un vector
		Encontrado := v.EntornoActual.BuscarSimbolo(AuxPass.(string))
		if Encontrado.Value != nil {
			Tipo := Encontrado.Type
			if Encontrado.EsVector {
				List := Encontrado.Value.([]interface{})
				v.EntornoActual = NewEntorno(v.EntornoActual, "For")
				if Tipo == "Int" {
					CreateSymbol := NewSimbolo(0, Tipo, "Let", false)
					v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)

					for i := 0; i < len(List); i++ {
						newi := List[i]
						v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, Tipo, "Let", false))
						for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
							v.Visit(ctx.Bloque().Lista_proceso(i))
						}
					}
					v.EntornoActual = v.EntornoActual.Padre
					return true
				} else if Tipo == "Character" {
					CreateSymbol := NewSimbolo("", Tipo, "Let", false)
					v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)

					for i := 0; i < len(List); i++ {
						newi := List[i]
						v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, Tipo, "Let", false))
						for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
							v.Visit(ctx.Bloque().Lista_proceso(i))
						}
					}
					v.EntornoActual = v.EntornoActual.Padre
					return true
				} else if Tipo == "String" {
					createSymbol := NewSimbolo("", Tipo, "Let", false)
					v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), createSymbol)

					for i := 0; i < len(List); i++ {
						newi := List[i]
						v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, Tipo, "Let", false))
						for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
							v.Visit(ctx.Bloque().Lista_proceso(i))
						}
					}
					v.EntornoActual = v.EntornoActual.Padre
					return true
				} else if Tipo == "Float" {
					createSymbol := NewSimbolo("", Tipo, "Let", false)
					v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), createSymbol)

					for i := 0; i < len(List); i++ {
						newi := List[i]
						v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, Tipo, "Let", false))
						for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
							v.Visit(ctx.Bloque().Lista_proceso(i))
						}
					}
					v.EntornoActual = v.EntornoActual.Padre
					return true
				}
			}
		}

		if Miceleanos.IsString(izq) {
			v.EntornoActual = NewEntorno(v.EntornoActual, "For")
			CreateSymbol := NewSimbolo("0", "Character", "Let", false)
			v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)
			for i := 0; i < len(izq.(string)); i++ {
				newi := string(izq.(string)[i])
				v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, "Character", "Let", false))
				for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
					NewReturn := v.Visit(ctx.Bloque().Lista_proceso(i))
					if NewReturn == false {
						v.EntornoActual = v.EntornoActual.Padre
						return false
					}
				}
			}
			v.EntornoActual = v.EntornoActual.Padre
		} else {
			v.AddError(ctx.GetLeft().GetStart().GetLine(), ctx.GetLeft().GetStart().GetColumn(), "FOR", "EL TIPO ES INVALIDO", "SEMANTICO")
			return false
		}

	}

	return true
}

//-------------------------------------------------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------   Guard

func (v *Visitor) VisitGuard(ctx *parser.GuardContext) interface{} {

	condicion := v.Visit(ctx.Expresion())
	if Miceleanos.Isbool(condicion) {
		if !condicion.(bool) {
			v.EntornoActual = NewEntorno(v.EntornoActual, "Guard")
			for i := 0; ctx.Lista_proceso(i) != nil; i++ {
				v.Visit(ctx.Lista_proceso(i))
				if v.Continue {
					v.Continue = false
					v.EntornoActual = v.EntornoActual.Padre
					return "continue"
				}
			}
			v.EntornoActual = v.EntornoActual.Padre

		}
	} else {
		v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "GUARD", "La condicion no es booleana", "SEMANTICO")
		return false
	}

	return true
}

// -------------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------   Vector
func (v *Visitor) VisitDec_vector(ctx *parser.Dec_vectorContext) interface{} {
	NuevoVector := Proceso.NewVector()
	NuevoVector.TipoVariable = ctx.GetOp().GetText()
	NuevoVector.Id = ctx.ID().GetText()
	NuevoVector.TipoDate = ctx.Tipo().GetText()
	for i := 0; ctx.Expresion(i) != nil; i++ {
		valor := v.Visit(ctx.Expresion(i))
		correcto := NuevoVector.VectorAppend(valor)
		if !correcto {
			v.AddError(ctx.Expresion(i).GetStart().GetLine(), ctx.Expresion(i).GetStart().GetColumn(), "VECTOR", "Tipo esperado: "+NuevoVector.TipoDate, "SEMANTICO")
			return false
		}
	}

	CreateSymbol := NewSimbolo(NuevoVector.GetList(), NuevoVector.GetTipo(), NuevoVector.GetTipoVariable(), true)
	Is_correct := v.EntornoActual.EnvAddSimbolo(NuevoVector.GetId(), CreateSymbol)
	if Is_correct {
		TaGlobal := NewSimbGlobal("Vector", NuevoVector.GetTipo(), v.EntornoActual.GetEntorno(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
		v.AddGlobalSimbol(NuevoVector.GetId(), TaGlobal)
	} else {
		v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "Ya existe en el entorno actual: "+NuevoVector.GetId(), "SEMANTICO")
		return false
	}

	return true
}

func (v *Visitor) VisitDec_vector_V_C(ctx *parser.Dec_vector_V_CContext) interface{} {

	if ctx.SubVC().OPEN_SQUARE_BRACKET() == nil && ctx.SubVC().CLOSE_SQUARE_BRACKET() == nil {
		id := ctx.SubVC().ID().GetText()
		Encontrado := v.EntornoActual.BuscarSimbolo(id)
		if Encontrado.Value != nil {
			if Encontrado.Type == ctx.Tipo().GetText() {
				CreateSymbol := NewSimbolo(Encontrado.Value, ctx.Tipo().GetText(), ctx.GetOp().GetText(), true)
				Is_correct := v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)
				if Is_correct {
					TaGlobal := NewSimbGlobal("Vector", Encontrado.Type, v.EntornoActual.GetEntorno(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
					v.AddGlobalSimbol(ctx.ID().GetText(), TaGlobal)
					return true
				} else {
					v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "Ya existe en el entorno actual: "+ctx.ID().GetText(), "SEMANTICO")
					return false
				}
			} else {
				v.AddError(ctx.SubVC().ID().GetSymbol().GetLine(), ctx.SubVC().ID().GetSymbol().GetColumn(), "VECTOR", "Tipo esperado: "+Encontrado.Type, "SEMANTICO")
				return false
			}
		} else {
			v.AddError(ctx.SubVC().ID().GetSymbol().GetLine(), ctx.SubVC().ID().GetSymbol().GetColumn(), "VECTOR", "No existe la variable: "+ctx.SubVC().ID().GetText(), "SEMANTICO")
			return false
		}

	} else {
		NuevoVector := Proceso.NewVector()
		NuevoVector.TipoVariable = ctx.GetOp().GetText()
		NuevoVector.Id = ctx.ID().GetText()
		NuevoVector.TipoDate = ctx.Tipo().GetText()
		CreateSymbol := NewSimbolo(NuevoVector.GetList(), NuevoVector.GetTipo(), NuevoVector.GetTipoVariable(), true)
		Is_correct := v.EntornoActual.EnvAddSimbolo(NuevoVector.GetId(), CreateSymbol)
		if Is_correct {
			TaGlobal := NewSimbGlobal("Vector", NuevoVector.GetTipo(), v.EntornoActual.GetEntorno(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
			v.AddGlobalSimbol(NuevoVector.GetId(), TaGlobal)
		} else {
			v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "Ya existe en el entorno actual: "+NuevoVector.GetId(), "SEMANTICO")
			return false
		}

	}

	return true
}

func (v *Visitor) VisitFuncvectorList(ctx *parser.FuncvectorListContext) interface{} {

	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarSimbolo(Id)
	esTipo := ctx.Tipevector().GetText()
	if Encontrado.EsVector {
		if esTipo == "append" {
			if Encontrado.TypeValue == "var" {
				AppendVector := Proceso.NewVector()
				AppendVector.TipoDate = Encontrado.Type
				AppendVector.List = Encontrado.Value.([]interface{})
				AppendVector.Id = Id
				AppendVector.TipoVariable = Encontrado.TypeValue
				is_Correct := AppendVector.FunctionAppend(v.Visit(ctx.Expresion()))
				if is_Correct {
					NuevoValor := NewSimbolo(AppendVector.GetList(), AppendVector.GetTipo(), AppendVector.GetTipoVariable(), true)
					v.EntornoActual.ActualizarSimbolo(Id, NuevoValor)
				} else {
					v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "VECTOR", "Tipo esperado: "+AppendVector.TipoDate, "SEMANTICO")
					return false
				}

			} else {
				v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.Expresion().GetStart().GetColumn(), "VECTOR", "Es contante no puede asignarse", "SEMANTICO")
				return false
			}
		} else if esTipo == "removeLast" {
			if Encontrado.TypeValue == "var" {
				RemoveVector := Proceso.NewVector()
				RemoveVector.TipoDate = Encontrado.Type
				RemoveVector.List = Encontrado.Value.([]interface{})
				RemoveVector.Id = Id
				RemoveVector.TipoVariable = Encontrado.TypeValue
				is_Correct := RemoveVector.RemoveLast()
				if is_Correct {
					NuevoValor := NewSimbolo(RemoveVector.GetList(), RemoveVector.GetTipo(), RemoveVector.GetTipoVariable(), true)
					v.EntornoActual.ActualizarSimbolo(Id, NuevoValor)
				} else {
					v.AddError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "VECTOR", "El vector no tiene elementos", "SEMANTICO")
					return false
				}

			}
		} else if esTipo == "remove" {
			if Encontrado.TypeValue == "var" {
				RemoveVector := Proceso.NewVector()
				RemoveVector.TipoDate = Encontrado.Type
				RemoveVector.List = Encontrado.Value.([]interface{})
				RemoveVector.Id = Id
				RemoveVector.TipoVariable = Encontrado.TypeValue
				is_Correct := RemoveVector.Remove(v.Visit(ctx.Expresion()))
				if is_Correct == true {
					NuevoValor := NewSimbolo(RemoveVector.GetList(), RemoveVector.GetTipo(), RemoveVector.GetTipoVariable(), true)
					v.EntornoActual.ActualizarSimbolo(Id, NuevoValor)
				} else if is_Correct == false {
					v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "VECTOR", "Tipo esperado: "+RemoveVector.GetTipo(), "SEMANTICO")
					return false

				} else {
					v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "VECTOR", "No tiene elementos o fuera de rango", "SEMANTICO")
					return false
				}

			}
		}
	}

	return true
}

func (v *Visitor) VisitVectorCount(ctx *parser.VectorCountContext) interface{} {
	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarSimbolo(Id)
	CountVector := Proceso.NewVector()
	if Encontrado.EsVector {

		CountVector.TipoDate = Encontrado.Type
		CountVector.List = Encontrado.Value.([]interface{})
		CountVector.Id = Id
		CountVector.TipoVariable = Encontrado.TypeValue
	}
	return CountVector.Count()
}

func (v *Visitor) VisitVectorVacio(ctx *parser.VectorVacioContext) interface{} {
	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarSimbolo(Id)
	CountVector := Proceso.NewVector()
	if Encontrado.EsVector {
		CountVector.TipoDate = Encontrado.Type
		CountVector.List = Encontrado.Value.([]interface{})
		CountVector.Id = Id
		CountVector.TipoVariable = Encontrado.TypeValue
	}
	return CountVector.IsVacio()
}

func (v *Visitor) VisitVectorAsignacion(ctx *parser.VectorAsignacionContext) interface{} {
	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarSimbolo(Id)
	if !SymbolVacio(Encontrado) {
		v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "No existe la variable: "+Id, "SEMANTICO")
		return false
	}

	List := Encontrado.Value.([]interface{})
	if int64(len(List)) < 0 {
		v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "ListaVacia", "SEMANTICO")
		return false
	} else if v.Visit(ctx.Expresion()).(int64) > int64(len(List)) {
		v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "Fuera de rango", "SEMANTICO")
		return false
	}

	return List[v.Visit(ctx.Expresion()).(int64)]
}

func (v *Visitor) VisitAsignacionVector(ctx *parser.AsignacionVectorContext) interface{} {

	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarSimbolo(Id)
	if SymbolVacio(Encontrado) {
		if Encontrado.TypeValue == "var" {
			if ctx.Subasig().NUMBER() != nil {
				i, _ := strconv.ParseInt(ctx.Subasig().NUMBER().GetText(), 10, 64)
				List := Encontrado.Value.([]interface{})

				if i > int64(len(List)) {
					v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "Fuera de rango", "SEMANTICO")
					return false
				}
				if ctx.GetOp().GetText() == "=" {
					List[i] = v.Visit(ctx.Expresion())
					v.EntornoActual.ActualizarSimbolo(Id, NewSimbolo(List, Encontrado.Type, Encontrado.TypeValue, true))

				} else if ctx.GetOp().GetText() == "+=" {
					var aux interface{}
					aux = List[i]
					aux = aux.(int64) + v.Visit(ctx.Expresion()).(int64)
					List[i] = aux
					v.EntornoActual.ActualizarSimbolo(Id, NewSimbolo(List, Encontrado.Type, Encontrado.TypeValue, true))

				}

			} else {
				iter := v.EntornoActual.BuscarSimbolo(ctx.Subasig().ID().GetText())
				i := iter.Value.(int64)

				List := Encontrado.Value.([]interface{})
				if ctx.GetOp().GetText() == "=" {
					List[i] = v.Visit(ctx.Expresion())
					v.EntornoActual.ActualizarSimbolo(Id, NewSimbolo(List, Encontrado.Type, Encontrado.TypeValue, true))

				} else if ctx.GetOp().GetText() == "+=" {
					var aux interface{}
					aux = List[i]
					aux = aux.(int64) + v.Visit(ctx.Expresion()).(int64)
					List[i] = aux
					v.EntornoActual.ActualizarSimbolo(Id, NewSimbolo(List, Encontrado.Type, Encontrado.TypeValue, true))

				}

			}
		} else {
			v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "Es Let no puede cambiar su valor", "SEMANTICO")
			return false
		}
	} else {
		v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "No existe la variable: "+Id, "SEMANTICO")
		return false
	}
	return true
}

// -------------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------   Casteos

func (v *Visitor) VisitCasteo(ctx *parser.CasteoContext) interface{} {

	ValorCasteo := v.Visit(ctx.Expresion())
	Tipo := ctx.Tipo().GetText()
	if Tipo == "Int" {
		switch convertido := ValorCasteo.(type) {
		case string:
			retdecimal := strings.Split(convertido, ".")
			if len(retdecimal) > 0 {
				convertido = retdecimal[0]
			}
			Entero, error := strconv.ParseInt(convertido, 10, 64)
			if error != nil {
				return nil
			}
			return Entero
		case float64:
			Float := int64(convertido)
			return Float

		default:
			return "ninguno"

		}
	} else if Tipo == "Float" {
		switch convertido := ValorCasteo.(type) {
		case string:
			conDecimal := false
			for i := 0; i < len(convertido); i++ {
				if string(convertido[i]) == "." {
					conDecimal = true
					break
				}
			}
			Float, error := strconv.ParseFloat(convertido, 64)
			if error != nil {
				return nil
			}

			if conDecimal {
				return Float
			}

			formattedValue := fmt.Sprintf("%.2f", Float)
			return formattedValue

		default:
			return "ninguno"

		}

	} else if Tipo == "String" {
		switch convertido := ValorCasteo.(type) {
		case int64:
			var cadena string = strconv.FormatInt(convertido, 10)
			return cadena
		case float64:
			var cadena string = strconv.FormatFloat(convertido, 'f', -1, 64)
			return cadena
		case bool:
			var cadena string = strconv.FormatBool(convertido)
			return cadena
		default:
			return "ninguno"

		}
	}

	v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "CASTEO", "No se puede castear a: "+Tipo, "SEMANTICO")
	return "ninguno"
}

func (v *Visitor) VisitEstructs(ctx *parser.EstructsContext) interface{} {
	fmt.Println("Estruct reconocido")
	return true
}
