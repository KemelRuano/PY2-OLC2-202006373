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

// Variables de manejo en C3D

// manejo funciones nativas
var Print_String = false
var CopiarHeap = false
var Control_Count = false

var TypeBoolC3D int = 0
var ControlString = false

// while
var LabelPassSI = ""
var LabelPassNo = ""

// manejo de operadores AND y OR
var Es_AND_NOT = 0
var LabelAN = ""

var ControlCiclos = false

var TypeDate string = ""
var RET_SET bool = false
var RET_FUN bool = false

var RETVALORDEELSE bool = false

type Visitor struct {
	*parser.BaseSintaxVisitor
	EntornoActual *Entorno
	ListErrores   []*Miceleanos.ErrorAnalizador
	TableGlobal   map[string]GlobalSymbol
	ListConsole   []string
	Traductor     *Traduction
	Retorno       bool
	Break         bool
	Continue      bool
}

func NewVisitor(ent *Entorno) *Visitor {
	return &Visitor{
		EntornoActual: ent,
		ListErrores:   []*Miceleanos.ErrorAnalizador{},
		TableGlobal:   make(map[string]GlobalSymbol),
		ListConsole:   []string{},
		Traductor:     NewTraduction(),
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
	case *antlr.ErrorNodeImpl:
		return "ERROR PERRRROOOOOOO  " + t.GetText()
	default:
		return tree.Accept(v)
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

	if Print_String {
		v.Traductor.List_funciones = append(v.Traductor.List_funciones, v.Traductor.PrimitiveString())
	}
	if CopiarHeap {
		v.Traductor.List_funciones = append(v.Traductor.List_funciones, v.Traductor.PrimiteConcatString())
	}
	if Control_Count {
		v.Traductor.List_funciones = append(v.Traductor.List_funciones, v.Traductor.VectorCount())
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
	if ctx.Comentarios() != nil {
		return v.Visit(ctx.Comentarios())
	}
	return nil
}

func (v *Visitor) VisitDeclaracion(ctx *parser.DeclaracionContext) interface{} {

	if ctx.Dec_tipo_valor() != nil {
		return v.Visit(ctx.Dec_tipo_valor())
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

func (v *Visitor) VisitDeclaracionTipo(ctx *parser.DeclaracionTipoContext) interface{} {
	TypeDate = ctx.Tipo().GetText()
	LineValor := ctx.Expresion().GetStart().GetLine()
	ColValor := ctx.Expresion().GetStart().GetColumn()
	LineId := ctx.ID().GetSymbol().GetLine()
	ColId := ctx.ID().GetSymbol().GetColumn()
	Valor := v.Visit(ctx.Expresion())

	// validar errores entre operaciones
	if Valor.(Value).Valor == "-1" {
		v.AddError(LineValor, ColValor, "DECLARACION", "Error en la operacion", "SEMANTICO")
		return false
	}

	Dec := Proceso.NewDeclaracion(ctx.ID().GetText(), ctx.Tipo().GetText(), Valor.(Value).Valor)
	CorrectType := Dec.EsTipo()
	if CorrectType {
		CreateSimbol := NewSimbolo(Dec.GetValor(), Dec.GetTipo(), ctx.GetOp().GetText(), false)
		is_exist := v.EntornoActual.EnvAddSimbolo(Dec.GetVariable(), CreateSimbol)
		if is_exist {
			TaGlobal := NewSimbGlobal("Variable", Dec.GetTipo(), v.EntornoActual.GetEntorno(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetLine())
			v.AddGlobalSimbol(Dec.GetVariable(), TaGlobal)

			// C3D El id se agrega a lista de temporales
			if Dec.GetTipo() == "Int" || Dec.GetTipo() == "Float" {
				NewStack := v.Traductor.GetIndex()
				v.Traductor.AddComentario("\t//Declaracion Numerica")
				if Valor.(Value).Label1 != nil {
					v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(Valor.(Value).Temp))
				} else {
					v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(Valor.(Value).Valor))
				}
				v.Traductor.Br()
				v.EntornoActual.ActualizarStack(Dec.GetVariable(), NewStack)

			} else if Dec.GetTipo() == "Bool" {
				NewLabel1 := ""
				NewLabel3 := ""
				Index := v.Traductor.GetIndex()
				v.Traductor.AddComentario("\t//Declaracion Bool")
				switch TypeBoolC3D {
				case 2:
					NewLabel1 = Valor.(Value).Label1.(string)
					NewLabel3 = Valor.(Value).Label2.(string)
				default:
					NewLabel1 = v.Traductor.NewLabel()
					NewLabel3 = v.Traductor.NewLabel()
					if Dec.GetValor().(bool) {
						v.Traductor.AddGoto(NewLabel1)
					} else {
						v.Traductor.AddGoto(NewLabel3)
					}
				}

				NewLabel2 := v.Traductor.NewLabel()
				v.Traductor.AddLabel(NewLabel1) // Lx:
				v.Traductor.SetStack(fmt.Sprint(Index), "1")
				v.Traductor.AddGoto(NewLabel2)  // goto Lx;
				v.Traductor.AddLabel(NewLabel3) // Lx:
				v.Traductor.SetStack(fmt.Sprint(Index), "0")
				v.Traductor.AddGoto(NewLabel2)  // Lx:
				v.Traductor.AddLabel(NewLabel2) // Lx:
				v.EntornoActual.ActualizarStack(Dec.GetVariable(), Index)
			} else if Dec.GetTipo() == "String" {
				NewStack := v.Traductor.GetIndex()
				v.Traductor.AddComentario("\t//Declaracion String")
				v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(Valor.(Value).Temp))
				v.Traductor.Br()
				v.EntornoActual.ActualizarStack(Dec.GetVariable(), NewStack)
			} else if Dec.GetTipo() == "Character" {
				NewStack := v.Traductor.GetIndex()
				v.Traductor.AddComentario("\t//Declaracion Character")
				v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(Valor.(Value).Temp))
				v.Traductor.Br()
				v.EntornoActual.ActualizarStack(Dec.GetVariable(), NewStack)
			}

		} else {
			v.AddError(LineId, ColId, "DECLARACION", "Ya existe en el entorno actual: "+Dec.GetVariable(), "SEMANTICO")
			return false
		}
	} else {
		v.AddError(LineValor, ColValor, "DECLARACION", "Especifico "+Dec.GetTipo()+"()", "SEMANTICO")
		return false
	}
	return "Kemel Ruano"
}

func (v *Visitor) VisitDeclaracionTipoImplicito(ctx *parser.DeclaracionTipoImplicitoContext) interface{} {

	Line := ctx.Expresion().GetStart().GetLine()
	Col := ctx.Expresion().GetStart().GetColumn()
	TypeDate = ""
	Valor := v.Visit(ctx.Expresion())
	Dec := Proceso.NewDeclaracion(ctx.ID().GetText(), "", Valor.(Value).Valor)
	CorrectType := Dec.EsSinTipo()
	if CorrectType {
		CreateSimbol := NewSimbolo(Dec.GetValor(), Dec.GetTipo(), ctx.GetOp().GetText(), false)
		is_exist := v.EntornoActual.EnvAddSimbolo(Dec.GetVariable(), CreateSimbol)
		if is_exist {
			TaGlobal := NewSimbGlobal("Variable", Dec.GetTipo(), v.EntornoActual.GetEntorno(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
			v.AddGlobalSimbol(Dec.GetVariable(), TaGlobal)

			// CODIGO EN C3D
			if Dec.GetTipo() == "Int" || Dec.GetTipo() == "Float" {
				NewStack := v.Traductor.GetIndex()
				v.Traductor.AddComentario("\t//Declaracion Numerica")
				if Valor.(Value).Label1 != nil {
					v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(Valor.(Value).Temp))
				} else {
					v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(Valor.(Value).Valor))
				}
				v.Traductor.Br()
				v.EntornoActual.ActualizarStack(Dec.GetVariable(), NewStack)

			} else if Dec.GetTipo() == "Bool" {
				NewLabel1 := ""
				NewLabel3 := ""
				Index := v.Traductor.GetIndex()
				v.Traductor.AddComentario("\t//Declaracion Bool")
				switch TypeBoolC3D {
				case 2:
					NewLabel1 = Valor.(Value).Label1.(string)
					NewLabel3 = Valor.(Value).Label2.(string)
				default:
					NewLabel1 = v.Traductor.NewLabel()
					NewLabel3 = v.Traductor.NewLabel()
					if Dec.GetValor().(bool) {
						v.Traductor.AddGoto(NewLabel1)
					} else {
						v.Traductor.AddGoto(NewLabel3)
					}
				}

				NewLabel2 := v.Traductor.NewLabel()
				v.Traductor.AddLabel(NewLabel1) // Lx:
				v.Traductor.SetStack(fmt.Sprint(Index), "1")
				v.Traductor.AddGoto(NewLabel2)  // goto Lx;
				v.Traductor.AddLabel(NewLabel3) // Lx:
				v.Traductor.SetStack(fmt.Sprint(Index), "0")
				v.Traductor.AddGoto(NewLabel2)  // Lx:
				v.Traductor.AddLabel(NewLabel2) // Lx:
				v.EntornoActual.ActualizarStack(Dec.GetVariable(), Index)
			} else if Dec.GetTipo() == "String" {
				NewStack := v.Traductor.GetIndex()
				v.Traductor.AddComentario("\t//Declaracion String")
				v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(Valor.(Value).Temp))
				v.Traductor.Br()
				v.EntornoActual.ActualizarStack(Dec.GetVariable(), NewStack)
			} else if Dec.GetTipo() == "Character" {
				NewStack := v.Traductor.GetIndex()
				v.Traductor.AddComentario("\t//Declaracion Character")
				v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(Valor.(Value).Temp))
				v.Traductor.Br()
				v.EntornoActual.ActualizarStack(Dec.GetVariable(), NewStack)
			}

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
			var newExpresion interface{}

			// iniciando Traduccion en C3D
			if Encontrado.Type == "Int" || Encontrado.Type == "Float" {
				newExpresion = v.Visit(ctx.Expresion())
				if ctx.GetOp().GetText() == "=" {
					v.Traductor.AddComentario("\t//Asignacion =")
					v.Traductor.SetStack(fmt.Sprint(Encontrado.Stack), fmt.Sprint(newExpresion.(Value).Temp))
				} else if ctx.GetOp().GetText() == "+=" {
					v.Traductor.AddComentario("\t//Asignacion +=")
					NewTemp := v.Traductor.NewTemp()
					v.Traductor.GetStack(NewTemp, fmt.Sprint(Encontrado.Stack))
					NewTemp2 := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTemp2, NewTemp, newExpresion.(Value).Temp, "+")
					v.Traductor.SetStack(fmt.Sprint(Encontrado.Stack), fmt.Sprint(NewTemp2))
				} else if ctx.GetOp().GetText() == "-=" {
					v.Traductor.AddComentario("\t//Asignacion -=")
					NewTemp := v.Traductor.NewTemp()
					v.Traductor.GetStack(NewTemp, fmt.Sprint(Encontrado.Stack))
					NewTemp2 := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTemp2, NewTemp, newExpresion.(Value).Temp, "-")
					v.Traductor.SetStack(fmt.Sprint(Encontrado.Stack), fmt.Sprint(NewTemp2))
				}
			} else if Encontrado.Type == "Bool" {
				newExpresion = v.Visit(ctx.Expresion())
				if ctx.GetOp().GetText() == "=" {
					v.Traductor.AddComentario("\t//Asignacion Bool =")
					if newExpresion.(Value).Valor.(bool) {
						v.Traductor.SetStack(fmt.Sprint(Encontrado.Stack), "1")
					} else {
						v.Traductor.SetStack(fmt.Sprint(Encontrado.Stack), "0")
					}

				}
			} else if Encontrado.Type == "String" {
				if ctx.GetOp().GetText() == "=" {
					v.Traductor.AddComentario("\t//Asignacion String o Char =")
					newExpresion = v.Visit(ctx.Expresion())
					v.Traductor.SetStack(fmt.Sprint(Encontrado.Stack), fmt.Sprint(newExpresion.(Value).Temp))
				} else if ctx.GetOp().GetText() == "+=" {
					CopiarHeap = true
					ControlString = true
					newExpresion = v.Visit(ctx.Expresion())
					NewString := newExpresion.(Value).Valor.(string)
					NewTemp := v.Traductor.NewTemp()
					v.Traductor.Br()
					v.Traductor.AddComentario("\t//Concatenando String")
					v.Traductor.Br()
					v.Traductor.Igual(NewTemp, "H", nil, nil)
					byteArray := []byte(NewString)
					for i := 0; i < len(byteArray); i++ {
						v.Traductor.SetHeap("(int)H", strconv.Itoa(int(byteArray[i])))
						v.Traductor.Contador("H", "H", "+", "1")
					}
					v.Traductor.SetHeap("(int)H", "-1")
					v.Traductor.Contador("H", "H", "+", "1")
					v.Traductor.Br()

					NewTemp2 := v.Traductor.NewTemp()
					v.Traductor.GetStack(NewTemp2, fmt.Sprint(Encontrado.Stack))
					NewTemp3 := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTemp3, "H", nil, nil)
					v.Traductor.Br()
					v.Traductor.Br()
					NewTemp4 := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTemp4, "P", v.EntornoActual.Size, "+")
					v.Traductor.Igual(NewTemp4, NewTemp4, "1", "+")
					v.Traductor.SetStack("(int) "+fmt.Sprint(NewTemp4), NewTemp2)
					v.Traductor.Igual("P", "P", v.EntornoActual.Size, "+")
					v.Traductor.callVoid("CopiarHeap")
					v.Traductor.Igual("P", "P", v.EntornoActual.Size, "-")
					v.Traductor.Br()

					v.Traductor.Igual(NewTemp4, "P", v.EntornoActual.Size, "+")
					v.Traductor.Igual(NewTemp4, NewTemp4, "1", "+")
					v.Traductor.SetStack("(int) "+fmt.Sprint(NewTemp4), NewTemp)
					v.Traductor.Igual("P", "P", v.EntornoActual.Size, "+")
					v.Traductor.callVoid("ConcatenarString")
					v.Traductor.Igual("P", "P", v.EntornoActual.Size, "-")
					v.Traductor.Br()

					v.Traductor.SetHeap("(int)H", "-1")
					v.Traductor.Contador("H", "H", "+", "1")
					v.Traductor.Br()

					v.Traductor.SetStack(fmt.Sprint(Encontrado.Stack), NewTemp3)
					v.Traductor.Br()
					v.Traductor.AddComentario("\t//Fin Concatenando String")

					ControlString = false

				}

			}
			// fin de traduccion en C3D

			NewValor := Proceso.NewConcatenar(Encontrado.Value, newExpresion.(Value).Valor, ctx.GetOp().GetText()).Reasignacion()
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
		return NewValue(Entero, nil, nil, Entero)
	}

	Decimal, error := strconv.ParseFloat(ctx.GetText(), 64)
	if error == nil {
		return NewValue(Decimal, nil, nil, Decimal)
	}

	return false
}

func (v *Visitor) VisitBoolValor(ctx *parser.BoolValorContext) interface{} {
	boleano, _ := strconv.ParseBool(ctx.GetText())
	return NewValue(nil, nil, nil, boleano)
}

func (v *Visitor) VisitStringValor(ctx *parser.StringValorContext) interface{} {
	cadena := strings.Trim(ctx.GetText(), "\"")
	NewTemp := ""
	if !ControlString {
		NewTemp = v.Traductor.NewTemp()
		v.Traductor.Br()
		v.Traductor.AddComentario("\t//Ingresando String")
		v.Traductor.Br()
		v.Traductor.Igual(NewTemp, "H", nil, nil)
		byteArray := []byte(cadena)
		for i := 0; i < len(byteArray); i++ {
			v.Traductor.SetHeap("(int)H", strconv.Itoa(int(byteArray[i])))
			v.Traductor.Contador("H", "H", "+", "1")
		}
		v.Traductor.SetHeap("(int)H", "-1")
		v.Traductor.Contador("H", "H", "+", "1")
		v.Traductor.Br()
	}

	return NewValue(NewTemp, nil, nil, cadena)
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

	// iniciando traduciendo a C3D
	NewTemp := v.Traductor.NewTemp()
	if op == "+" {
		v.Traductor.Igual(NewTemp, l.(Value).Temp, r.(Value).Temp, op)
	} else if op == "-" {
		v.Traductor.Igual(NewTemp, l.(Value).Temp, r.(Value).Temp, op)
	} else if op == "*" {
		v.Traductor.Igual(NewTemp, l.(Value).Temp, r.(Value).Temp, op)
	} else if op == "/" {
		NewLabel := v.Traductor.NewLabel()
		v.Traductor.AddIf(fmt.Sprint(r.(Value).Temp), "0", "!=", NewLabel)
		v.Traductor.AddPrint("77", "c")
		v.Traductor.AddPrint("97", "c")
		v.Traductor.AddPrint("116", "c")
		v.Traductor.AddPrint("104", "c")
		v.Traductor.AddPrint("69", "c")
		v.Traductor.AddPrint("114", "c")
		v.Traductor.AddPrint("114", "c")
		v.Traductor.AddPrint("111", "c")
		v.Traductor.AddPrint("114", "c")
		NewLabel2 := v.Traductor.NewLabel()
		v.Traductor.Igual(NewTemp, "0", "", "")
		v.Traductor.AddGoto(NewLabel2)
		v.Traductor.AddLabel(NewLabel)
		v.Traductor.Igual(NewTemp, l.(Value).Temp, r.(Value).Temp, op)
		v.Traductor.AddLabel(NewLabel2)
	} else if op == "%" {
		v.Traductor.Igual(NewTemp, "(int) "+fmt.Sprint(l.(Value).Temp), r.(Value).Temp, op)
	}
	// fin traduccion a C3D

	ret := Proceso.NewConcatenar(l.(Value).Valor, r.(Value).Valor, op).Aritmeticos()
	return NewValue(NewTemp, true, nil, ret)
}

func (v *Visitor) VisitParentexpr(ctx *parser.ParentexprContext) interface{} {
	visitValor := v.Visit(ctx.Expresion())
	return visitValor
}

func (v *Visitor) VisitOpComparativo(ctx *parser.OpComparativoContext) interface{} {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()

	// inicio traducion a C3D
	NewLabel1 := ""
	NewLabel2 := ""

	switch Es_AND_NOT {
	case 1:
		NewLabel1 = v.Traductor.NewLabel()
		NewLabel2 = LabelAN
		v.Traductor.AddIf(fmt.Sprint(l.(Value).Temp), fmt.Sprint(r.(Value).Temp), op, NewLabel1)
		v.Traductor.AddGoto(NewLabel2)

	case 2:
		NewLabel1 = LabelAN
		NewLabel2 = v.Traductor.NewLabel()
		v.Traductor.AddIf(fmt.Sprint(l.(Value).Temp), fmt.Sprint(r.(Value).Temp), op, NewLabel1)
		v.Traductor.AddGoto(NewLabel2)

	default:
		NewLabel1 = v.Traductor.NewLabel()
		NewLabel2 = v.Traductor.NewLabel()
		v.Traductor.AddIf(fmt.Sprint(l.(Value).Temp), fmt.Sprint(r.(Value).Temp), op, NewLabel1)
		v.Traductor.AddGoto(NewLabel2)

	}
	// fin traduccion a C3D

	TypeBoolC3D = 2
	ret := Proceso.NewConcatenar(l.(Value).Valor, r.(Value).Valor, op).Comparativos()
	return NewValue("", NewLabel1, NewLabel2, ret)
}

func (v *Visitor) VisitOpRelacional(ctx *parser.OpRelacionalContext) interface{} {

	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()

	// inicio traducion a C3D
	NewLabel1 := ""
	NewLabel2 := ""

	switch Es_AND_NOT {
	case 1:
		NewLabel1 = v.Traductor.NewLabel()
		NewLabel2 = LabelAN
		v.Traductor.AddIf(fmt.Sprint(l.(Value).Temp), fmt.Sprint(r.(Value).Temp), op, NewLabel1)
		v.Traductor.AddGoto(NewLabel2)

	case 2:
		NewLabel1 = LabelAN
		NewLabel2 = v.Traductor.NewLabel()
		v.Traductor.AddIf(fmt.Sprint(l.(Value).Temp), fmt.Sprint(r.(Value).Temp), op, NewLabel1)
		v.Traductor.AddGoto(NewLabel2)

	default:
		NewLabel1 = v.Traductor.NewLabel()
		NewLabel2 = v.Traductor.NewLabel()
		v.Traductor.AddIf(fmt.Sprint(l.(Value).Temp), fmt.Sprint(r.(Value).Temp), op, NewLabel1)
		v.Traductor.AddGoto(NewLabel2)

	}
	// fin traduccion a C3D

	TypeBoolC3D = 2
	ret := Proceso.NewConcatenar(l.(Value).Valor, r.(Value).Valor, op).Relacionales()

	return NewValue(nil, NewLabel1, NewLabel2, ret)
}

func (v *Visitor) VisitOpLogico(ctx *parser.OpLogicoContext) interface{} {
	var l interface{}
	var r interface{}
	op := ctx.GetOp().GetText()

	// Inicio Traduccion C3D
	Label1 := ""
	Label2 := ""
	if op == "&&" {
		v.Traductor.AddComentario("\t//Operador Logico AND")
		l = v.Visit(ctx.GetLeft())
		v.Traductor.AddLabel(l.(Value).Label1.(string))
		Es_AND_NOT = 1
		LabelAN = l.(Value).Label2.(string)
		r = v.Visit(ctx.GetRight())
		Label1 = r.(Value).Label1.(string)
		Label2 = r.(Value).Label2.(string)
		Es_AND_NOT = 0
		LabelAN = ""
	} else if op == "||" {
		v.Traductor.AddComentario("\t//Operador Logico OR")
		l = v.Visit(ctx.GetLeft())
		v.Traductor.AddLabel(l.(Value).Label2.(string))
		Es_AND_NOT = 2
		LabelAN = l.(Value).Label1.(string)
		r = v.Visit(ctx.GetRight())
		Label1 = r.(Value).Label1.(string)
		Label2 = r.(Value).Label2.(string)
		Es_AND_NOT = 0
		LabelAN = ""
	}
	// Fin Traduccion C3D

	ret := Proceso.NewConcatenar(l.(Value).Valor, r.(Value).Valor, op).Logicos()
	return NewValue(nil, Label1, Label2, ret)
}
func (v *Visitor) VisitOpLogicoNot(ctx *parser.OpLogicoNotContext) interface{} {
	l := v.Visit(ctx.Expresion())
	if Miceleanos.Isbool(l) {
		l = !l.(bool)
	}
	return l
}

// Print
func (v *Visitor) VisitPrint(ctx *parser.PrintContext) interface{} {

	var Concatenar string = ""
	for i := 0; ctx.Expresion(i) != nil; i++ {
		Print := v.Visit(ctx.Expresion(i))
		Valor := Print.(Value).Valor
		Temp := Print.(Value).Temp

		if Miceleanos.IsString(Valor) {
			if len(Valor.(string)) < 2 {

				v.Traductor.AddComentario("\t//Imprimir Character")
				NewTemp := v.Traductor.NewTemp()
				v.Traductor.GetHeap(NewTemp, "(int)"+fmt.Sprint(Temp))
				v.Traductor.AddPrint("(char) "+NewTemp, "c")
				v.Traductor.AddPrint("10", "c")
				v.Traductor.Br()
			} else {
				v.Traductor.AddComentario("\t//Imprimir String")
				NewTemp := v.Traductor.NewTemp()
				NewTemp2 := v.Traductor.NewTemp()
				v.Traductor.Igual(NewTemp, "P", v.EntornoActual.Size, "+")
				v.Traductor.Contador(NewTemp, NewTemp, "+", "1")
				v.Traductor.SetStack("(int)"+NewTemp, fmt.Sprint(Temp))
				v.Traductor.Contador("P", "P", "+", fmt.Sprint(v.EntornoActual.Size))
				v.Traductor.callVoid("ImprimirString")
				v.Traductor.GetStack(NewTemp2, "(int)P")
				v.Traductor.Contador("P", "P", "-", fmt.Sprint(v.EntornoActual.Size))
				v.Traductor.AddPrint("10", "c")
				v.Traductor.Br()
				Print_String = true
			}

		} else if Miceleanos.IsInt(Valor) {
			// C3D Imprimir entero
			v.Traductor.AddComentario("\t//Imprimir Entero")
			v.Traductor.AddPrint("(int) "+fmt.Sprint(Temp), "d")
			v.Traductor.AddPrint("10", "c")
			v.Traductor.Br()

		} else if Miceleanos.IsFloat(Valor) {

			// C3D Imprimir float
			v.Traductor.AddComentario("\t//Imprimir Float")
			v.Traductor.AddPrint("(float) "+fmt.Sprint(Temp), "f")
			v.Traductor.AddPrint("10", "c")
			v.Traductor.Br()

		} else if Miceleanos.Isbool(Valor) {
			// C3D Imprimir bool
			v.Traductor.AddComentario("\t//Imprimir Bool")
			NewLabel1 := ""
			NewLabel2 := ""
			switch TypeBoolC3D {
			case 1:
				NewLabel1 = v.Traductor.NewLabel()
				NewLabel2 = v.Traductor.NewLabel()
				v.Traductor.AddIf(fmt.Sprint(Temp), "1", "==", NewLabel1)
				v.Traductor.AddGoto(NewLabel2)
			case 2:
				NewLabel1 = fmt.Sprint(Print.(Value).Label1)
				NewLabel2 = fmt.Sprint(Print.(Value).Label2)
			default:
				NewLabel1 = v.Traductor.NewLabel()
				NewLabel2 = v.Traductor.NewLabel()
				if Valor.(bool) {
					v.Traductor.AddGoto(NewLabel1)
				} else {
					v.Traductor.AddGoto(NewLabel2)
				}

			}
			NewLabel3 := v.Traductor.NewLabel()
			v.Traductor.AddLabel(NewLabel1)         // Lx:
			v.Traductor.AddPrint("(char) 116", "c") // printf("%c", 116); // t
			v.Traductor.AddPrint("(char) 114", "c") // printf("%c", 114); // r
			v.Traductor.AddPrint("(char) 117", "c") // printf("%c", 117); // u
			v.Traductor.AddPrint("(char) 101", "c") // printf("%c", 101); // e
			v.Traductor.AddGoto(NewLabel3)          // goto Lx;
			v.Traductor.AddLabel(NewLabel2)         // Lx:
			v.Traductor.AddPrint("(char) 102", "c") // printf("%c", 102); // f
			v.Traductor.AddPrint("(char) 97", "c")  // printf("%c", 97); // a
			v.Traductor.AddPrint("(char) 108", "c") // printf("%c", 108); // l
			v.Traductor.AddPrint("(char) 115", "c") // printf("%c", 115); // s
			v.Traductor.AddPrint("(char) 101", "c") // printf("%c", 101); // e
			v.Traductor.AddLabel(NewLabel3)         // Lx:
			v.Traductor.AddPrint("10", "c")         // printf("%c", 10); // \n

		} else {
			SearchSymbol := v.EntornoActual.BuscarSimbolo(AuxPass.(string))
			if SearchSymbol.EsVector {
				v.Traductor.AddComentario("\t//-------------------------")
				v.Traductor.AddComentario("\t//Imprimir Vector")
				NewTemp := v.Traductor.NewTemp()
				LabelSI := v.Traductor.NewLabel()
				LabelNO := v.Traductor.NewLabel()
				v.Traductor.AddPrint("91", "c")
				v.Traductor.AddPrint("32", "c")

				v.Traductor.AddLabel(LabelNO)
				v.Traductor.GetHeap(NewTemp, "(int)"+fmt.Sprint(Temp))
				v.Traductor.AddIf(NewTemp, "-1", "==", LabelSI)
				v.Traductor.AddPrint("32", "c")
				if SearchSymbol.Type == "Int" {
					v.Traductor.AddPrint("(int) "+NewTemp, "d")
				} else if SearchSymbol.Type == "Float" {
					v.Traductor.AddPrint("(float) "+NewTemp, "f")
				} else if SearchSymbol.Type == "String" {
					LabelSI2 := v.Traductor.NewLabel()
					LabelNO2 := v.Traductor.NewLabel()
					NewTemp2 := v.Traductor.NewTemp()
					v.Traductor.AddPrint("91", "c")
					v.Traductor.AddLabel(LabelNO2)
					v.Traductor.GetHeap(NewTemp2, "(int)"+fmt.Sprint(NewTemp))
					v.Traductor.AddIf(NewTemp2, "-1", "==", LabelSI2)
					v.Traductor.AddPrint("(char) "+NewTemp2, "c")
					v.Traductor.Contador(NewTemp, NewTemp, "+", "1")
					v.Traductor.AddGoto(LabelNO2)
					v.Traductor.AddLabel(LabelSI2)
					v.Traductor.AddPrint("93", "c")
					v.Traductor.AddPrint("32", "c")
				} else if SearchSymbol.Type == "Character" {
					v.Traductor.AddPrint("(char) "+NewTemp, "c")
				}
				v.Traductor.Contador(fmt.Sprint(Temp), fmt.Sprint(Temp), "+", "1")
				v.Traductor.AddGoto(LabelNO)
				v.Traductor.AddLabel(LabelSI)
				v.Traductor.AddPrint("32", "c")
				v.Traductor.AddPrint("93", "c")
				v.Traductor.AddPrint("10", "c")
				AuxPass = nil

			}

		}
		Concatenar += fmt.Sprint(Valor)
	}
	// fmt.Println("CONSOLE: ", Concatenar)
	Concatenar += "\n"
	v.ListConsole = append(v.ListConsole, Concatenar)
	return true
}
func (v *Visitor) VisitIdValor(ctx *parser.IdValorContext) interface{} {
	id := ctx.GetText()
	value := v.EntornoActual.BuscarSimbolo(id)
	fmt.Println(value.EsVector, value.Type, value.Stack, value.Value, value.TypeValue)
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
	//Codigo en 3 direcciones
	Temp := v.Traductor.NewTemp()
	if value.Type == "Int" || value.Type == "Float" {
		v.Traductor.AddComentario("\t//llamando una variable")
		v.Traductor.GetStack(Temp, fmt.Sprint(value.Stack))
		v.Traductor.Br()

	} else if value.Type == "Bool" {
		v.Traductor.AddComentario("\t//llamando una variable")
		v.Traductor.GetStack(Temp, fmt.Sprint(value.Stack))
		TypeBoolC3D = 1
	} else if value.Type == "String" {
		v.Traductor.AddComentario("\t//llamando una variable")
		v.Traductor.GetStack(Temp, fmt.Sprint(value.Stack))
		v.Traductor.Br()
	} else if value.Type == "Character" {
		v.Traductor.AddComentario("\t//llamando una variable")
		v.Traductor.GetStack(Temp, fmt.Sprint(value.Stack))
		v.Traductor.Br()
	}
	return NewValue(Temp, nil, nil, value.Value)
}

func (v *Visitor) VisitIf(ctx *parser.IfContext) interface{} {
	RET_SET = false
	v.Traductor.AddComentario("//------------------------------------")
	v.Traductor.AddComentario("//Setencia IF")
	Valor := v.Visit(ctx.Expresion())
	condicion := Valor.(Value).Valor
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

	// inicio traduccion a C3D
	LabelIf := Valor.(Value).Label1
	LabelElse := ""
	LabelDefault := ""
	LabelAux := ""
	if LabelIf == nil && Valor.(Value).Label2 == nil {
		LabelIf = v.Traductor.NewLabel()
		if condicion.(bool) {
			v.Traductor.AddGoto(fmt.Sprint(LabelIf))
			v.Traductor.AddLabel(fmt.Sprint(LabelIf))
			LabelDefault = v.Traductor.NewLabel()
			fmt.Println(LabelDefault)
		} else {

			LabelElse = v.Traductor.NewLabel()
			v.Traductor.AddGoto(fmt.Sprint(LabelElse))
			v.Traductor.AddLabel(fmt.Sprint(LabelIf))
			LabelDefault = v.Traductor.NewLabel()
			fmt.Println(LabelDefault)
		}
	} else {
		v.Traductor.AddLabel(fmt.Sprint(LabelIf))
		LabelDefault = v.Traductor.NewLabel()
		LabelAux = Valor.(Value).Label2.(string)
		fmt.Println(LabelDefault)
	}

	// fin traduccion a C3D
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

	for i := 0; ctx.Elseif_(i) != nil; i++ {
		if i == 0 {
			v.Traductor.AddGoto(fmt.Sprint(LabelDefault))
			v.Traductor.AddLabel(LabelAux)
		} else {
			v.Traductor.AddLabel(LabelAN)
		}
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
		v.Traductor.AddGoto(fmt.Sprint(LabelDefault))
	}
	// inicio traduccion a C3D
	// if !v.EntornoActual.Elseif {
	// 	v.Traductor.AddGoto(fmt.Sprint(LabelAux))
	// }

	if ctx.Else_() != nil {
		if v.EntornoActual.Elseif {
			v.Traductor.AddComentario("//------------------------------------")
			v.Traductor.AddComentario("//Setencia ELSE")
			v.Traductor.AddLabel(LabelAN)
		} else {
			v.Traductor.AddGoto(fmt.Sprint(LabelDefault))
			v.Traductor.AddComentario("//------------------------------------")
			v.Traductor.AddComentario("//Setencia ELSE")
			v.Traductor.AddLabel(LabelAux)
		}
		v.EntornoActual.Else = true
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
	v.Traductor.AddComentario("//------------------------------------")
	if !v.EntornoActual.Elseif && !v.EntornoActual.Else {
		// Solo hay setencia if sin else ni elseif
		LabelDefault = LabelAux
		v.Traductor.AddLabel(fmt.Sprint(LabelDefault))
	} else if v.EntornoActual.Elseif && !v.EntornoActual.Else {
		// Solo hay setencia if con elseif
		v.Traductor.AddLabel(LabelAN)
		v.Traductor.AddLabel(fmt.Sprint(LabelDefault))
		LabelAN = ""
	} else if !v.EntornoActual.Elseif && v.EntornoActual.Else {
		// Solo hay setencia if con else
		v.Traductor.AddLabel(fmt.Sprint(LabelDefault))
	} else if v.EntornoActual.Elseif && v.EntornoActual.Else {
		// Tiene elseif y else
		v.Traductor.AddLabel(fmt.Sprint(LabelDefault))
	}

	return nil
}

func (v *Visitor) VisitElseif_(ctx *parser.Elseif_Context) interface{} {
	// inicio traduccion a C3D
	v.EntornoActual.Elseif = true
	v.Traductor.AddComentario("//-------------------------------------")
	v.Traductor.AddComentario("//Setencia ELSEIF")
	v.Traductor.Br()
	condicion := v.Visit(ctx.Expresion())
	LabelAN = ""
	LabelSI := condicion.(Value).Label1
	LabelAN = condicion.(Value).Label2.(string)
	v.Traductor.AddLabel(fmt.Sprint(LabelSI))
	if Miceleanos.Isbool(condicion.(Value).Valor) {
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

	} else {
		v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "ELSEIF", "La condicion no es booleana", "SEMANTICO")
		return false
	}

	return false
}

func (v *Visitor) VisitSwitch(ctx *parser.SwitchContext) interface{} {
	condicion := v.Visit(ctx.Expresion())

	v.Traductor.AddComentario("//-------------------------------------")
	v.Traductor.AddComentario("//Setencia SWITCH")
	v.Traductor.Br()
	NewLabel3 := v.Traductor.NewLabel()
	for i := 0; ctx.Caso(i) != nil; i++ {
		// codigo en 3 direcciones
		TypecCase := v.Visit(ctx.Caso(i).Expresion())
		NewLabel := v.Traductor.NewLabel()
		NewLabel2 := v.Traductor.NewLabel()
		v.Traductor.AddIf(fmt.Sprint(condicion.(Value).Temp), fmt.Sprint(TypecCase.(Value).Temp), "==", NewLabel)
		v.Traductor.AddGoto(NewLabel2)
		v.Traductor.AddLabel(NewLabel)
		v.EntornoActual = NewEntorno(v.EntornoActual, "Switch")
		v.Visit(ctx.Caso(i))
		v.EntornoActual = v.EntornoActual.Padre
		v.Traductor.AddGoto(NewLabel3)

		v.Traductor.AddLabel(NewLabel2)

	}
	if ctx.Default_() != nil {
		v.EntornoActual = NewEntorno(v.EntornoActual, "Switch")
		v.Visit(ctx.Default_())
		v.EntornoActual = v.EntornoActual.Padre
	}
	v.Traductor.AddLabel(NewLabel3)

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

	v.Traductor.AddComentario("//-------------------------------------")
	v.Traductor.AddComentario("//Setencia WHILE")
	v.Traductor.Br()
	NewLabel := v.Traductor.NewLabel()
	v.Traductor.AddLabel(NewLabel)
	Valor := v.Visit(ctx.Expresion())
	LabelPassSI = NewLabel
	LabelPassNo = Valor.(Value).Label2.(string)
	condicion := Valor.(Value).Valor
	if Miceleanos.Isbool(condicion.(bool)) {
		v.Traductor.AddLabel(Valor.(Value).Label1.(string))
		v.EntornoActual = NewEntorno(v.EntornoActual, "While")
		for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
			v.Visit(ctx.Bloque().Lista_proceso(i))

		}
		v.EntornoActual = v.EntornoActual.Padre
		v.Traductor.AddGoto(NewLabel)
		v.Traductor.AddLabel(Valor.(Value).Label2.(string))
	} else {
		v.AddError(ctx.Expresion().GetStart().GetLine(), ctx.Expresion().GetStart().GetColumn(), "WHILE", "La condicion no es booleana", "SEMANTICO")
	}

	return true
}

//-------------------------------------------------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------   for

func (v *Visitor) VisitFor(ctx *parser.ForContext) interface{} {
	izquierda := v.Visit(ctx.GetLeft())
	derecha := interface{}(nil)
	if ctx.GetRight() != nil {
		derecha = v.Visit(ctx.GetRight())
	}
	var comprobacion bool = true
	if ctx.GUION_BAJO() != nil {
		comprobacion = false
	}
	v.Traductor.AddComentario("//-------------------------------------")
	v.Traductor.AddComentario("//Setencia FOR")
	if ctx.TRESPUNTOS() != nil {
		ValorIz := izquierda.(Value).Valor
		TemporalIz := izquierda.(Value).Temp
		TemporalDer := derecha.(Value).Temp
		if comprobacion {
			// fot tipo 1  [ valor in range ]

			CreateSymbol := NewSimbolo(ValorIz, "Int", "Let", false)
			v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)

			// Inicio Codigo en 3 direcciones
			v.Traductor.AddComentario("\t//Declaracion variable de control")
			NewStack := v.Traductor.GetIndex()
			v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(TemporalIz))
			v.EntornoActual.ActualizarStack(ctx.ID().GetText(), NewStack)
			LabelCiclo := v.Traductor.NewLabel()
			LabelSi := v.Traductor.NewLabel()
			LabelNo := v.Traductor.NewLabel()
			v.Traductor.AddLabel(LabelCiclo)
			NewTemp := v.Traductor.NewTemp()
			v.Traductor.GetStack(NewTemp, fmt.Sprint(NewStack))
			v.Traductor.AddIf(NewTemp, fmt.Sprint(TemporalDer), "<=", LabelSi)
			v.Traductor.AddGoto(LabelNo)
			v.Traductor.AddLabel(LabelSi)
			v.EntornoActual = NewEntorno(v.EntornoActual, "For")
			for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
				newReturn := v.Visit(ctx.Bloque().Lista_proceso(i))
				if newReturn == false {
					v.EntornoActual = v.EntornoActual.Padre
					return false
				}
			}
			v.EntornoActual = v.EntornoActual.Padre
			NewTemp2 := v.Traductor.NewTemp()
			v.Traductor.GetStack(NewTemp2, fmt.Sprint(NewStack))
			v.Traductor.Contador(NewTemp2, NewTemp2, "+", "1")
			v.Traductor.SetStack(fmt.Sprint(NewStack), NewTemp2)
			v.Traductor.AddGoto(LabelCiclo)
			v.Traductor.AddLabel(LabelNo)
			// fin codigo en 3 direcciones

		} else {
			// for tipo 2  [   _ ... range  ]

			// inicio codigo en 3 direcciones
			LabelCiclo := v.Traductor.NewLabel()
			LabelSi := v.Traductor.NewLabel()
			LabelNo := v.Traductor.NewLabel()
			NewTemp := v.Traductor.NewTemp()
			v.Traductor.Igual(NewTemp, fmt.Sprint(ValorIz), nil, nil)
			v.Traductor.AddLabel(LabelCiclo)
			v.Traductor.AddIf(NewTemp, fmt.Sprint(TemporalDer), "<=", LabelSi)
			v.Traductor.AddGoto(LabelNo)
			v.Traductor.AddLabel(LabelSi)
			v.EntornoActual = NewEntorno(v.EntornoActual, "For")
			for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
				v.Visit(ctx.Bloque().Lista_proceso(i))
			}
			v.EntornoActual = v.EntornoActual.Padre
			v.Traductor.Contador(NewTemp, NewTemp, "+", "1")
			v.Traductor.AddGoto(LabelCiclo)
			v.Traductor.AddLabel(LabelNo)
			// fin codigo en 3 direcciones
		}
	} else {

		// for  tipo 3  [ valor  in  "string" ]
		if AuxPass == nil {
			AuxPass = ""
		}

		SearchSymbol := v.EntornoActual.BuscarSimbolo(AuxPass.(string))
		fmt.Println(SearchSymbol)

		if SearchSymbol.Value != nil || AuxPass == "" {
			ValorIz := izquierda.(Value).Valor
			TemporalIz := izquierda.(Value).Temp

			if !SearchSymbol.EsVector {
				if Miceleanos.IsString(ValorIz) || SearchSymbol.Type == "String" {
					CreateSymbol := NewSimbolo("0", "Character", "Let", false)
					v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)

					v.Traductor.AddComentario("\t//Declaracion variable de control")
					NewStack := v.Traductor.GetIndex()
					v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(TemporalIz))
					v.EntornoActual.ActualizarStack(ctx.ID().GetText(), NewStack)

					NewTempSize := v.Traductor.NewTemp()
					LenString := len(ValorIz.(string))
					v.Traductor.Igual(NewTempSize, fmt.Sprint(TemporalIz), fmt.Sprint(LenString), "+")
					LabelCiclo := v.Traductor.NewLabel()
					LabelSi := v.Traductor.NewLabel()
					LabelNo := v.Traductor.NewLabel()
					v.Traductor.AddLabel(LabelCiclo)
					NewTemp := v.Traductor.NewTemp()
					v.Traductor.GetStack(NewTemp, fmt.Sprint(NewStack))
					v.Traductor.AddIf(NewTemp, fmt.Sprint(NewTempSize), "<", LabelSi)
					v.Traductor.AddGoto(LabelNo)
					v.Traductor.AddLabel(LabelSi)
					v.EntornoActual = NewEntorno(v.EntornoActual, "For")
					for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
						v.Visit(ctx.Bloque().Lista_proceso(i))
					}
					v.EntornoActual = v.EntornoActual.Padre
					NewTemp2 := v.Traductor.NewTemp()
					v.Traductor.GetStack(NewTemp2, fmt.Sprint(NewStack))
					v.Traductor.Contador(NewTemp2, NewTemp2, "+", "1")
					v.Traductor.SetStack(fmt.Sprint(NewStack), NewTemp2)
					v.Traductor.AddGoto(LabelCiclo)
					v.Traductor.AddLabel(LabelNo)
					AuxPass = nil

				} else {
					v.AddError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "FOR", "No es de tipo String", "SEMANTICO")
					return false
				}

			} else {
				fmt.Println("VECTOR")
			}

		} else {
			v.AddError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "FOR", "La variable no existe", "SEMANTICO")
			return false
		}
	}

	// } else {
	// 	//buscar si es una cadena o un vector
	// 	Encontrado := v.EntornoActual.BuscarSimbolo(AuxPass.(string))
	// 	if Encontrado.Value != nil {
	// 		Tipo := Encontrado.Type
	// 		if Encontrado.EsVector {
	// 			List := Encontrado.Value.([]interface{})
	// 			v.EntornoActual = NewEntorno(v.EntornoActual, "For")
	// 			if Tipo == "Int" {
	// 				CreateSymbol := NewSimbolo(0, Tipo, "Let", false)
	// 				v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)

	// 				for i := 0; i < len(List); i++ {
	// 					newi := List[i]
	// 					v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, Tipo, "Let", false))
	// 					for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
	// 						v.Visit(ctx.Bloque().Lista_proceso(i))
	// 					}
	// 				}
	// 				v.EntornoActual = v.EntornoActual.Padre
	// 				return true
	// 			} else if Tipo == "Character" {
	// 				CreateSymbol := NewSimbolo("", Tipo, "Let", false)
	// 				v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)

	// 				for i := 0; i < len(List); i++ {
	// 					newi := List[i]
	// 					v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, Tipo, "Let", false))
	// 					for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
	// 						v.Visit(ctx.Bloque().Lista_proceso(i))
	// 					}
	// 				}
	// 				v.EntornoActual = v.EntornoActual.Padre
	// 				return true
	// 			} else if Tipo == "String" {
	// 				createSymbol := NewSimbolo("", Tipo, "Let", false)
	// 				v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), createSymbol)

	// 				for i := 0; i < len(List); i++ {
	// 					newi := List[i]
	// 					v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, Tipo, "Let", false))
	// 					for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
	// 						v.Visit(ctx.Bloque().Lista_proceso(i))
	// 					}
	// 				}
	// 				v.EntornoActual = v.EntornoActual.Padre
	// 				return true
	// 			} else if Tipo == "Float" {
	// 				createSymbol := NewSimbolo("", Tipo, "Let", false)
	// 				v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), createSymbol)

	// 				for i := 0; i < len(List); i++ {
	// 					newi := List[i]
	// 					v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, Tipo, "Let", false))
	// 					for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
	// 						v.Visit(ctx.Bloque().Lista_proceso(i))
	// 					}
	// 				}
	// 				v.EntornoActual = v.EntornoActual.Padre
	// 				return true
	// 			}
	// 		}
	// 	}

	// 	if Miceleanos.IsString(izq) {
	// 		v.EntornoActual = NewEntorno(v.EntornoActual, "For")
	// 		CreateSymbol := NewSimbolo("0", "Character", "Let", false)
	// 		v.EntornoActual.EnvAddSimbolo(ctx.ID().GetText(), CreateSymbol)
	// 		for i := 0; i < len(izq.(string)); i++ {
	// 			newi := string(izq.(string)[i])
	// 			v.EntornoActual.ActualizarSimbolo(ctx.ID().GetText(), NewSimbolo(newi, "Character", "Let", false))
	// 			for i := 0; ctx.Bloque().Lista_proceso(i) != nil; i++ {
	// 				NewReturn := v.Visit(ctx.Bloque().Lista_proceso(i))
	// 				if NewReturn == false {
	// 					v.EntornoActual = v.EntornoActual.Padre
	// 					return false
	// 				}
	// 			}
	// 		}
	// 		v.EntornoActual = v.EntornoActual.Padre
	// 	} else {
	// 		v.AddError(ctx.GetLeft().GetStart().GetLine(), ctx.GetLeft().GetStart().GetColumn(), "FOR", "EL TIPO ES INVALIDO", "SEMANTICO")
	// 		return false
	// 	}

	// }

	return true
}

//-------------------------------------------------------------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------   Guard

func (v *Visitor) VisitGuard(ctx *parser.GuardContext) interface{} {
	v.Traductor.AddComentario("//-------------------------------------")
	v.Traductor.AddComentario("//Setencia GUARD")
	condicion := v.Visit(ctx.Expresion())
	BoolSi := condicion.(Value).Label1
	BoolNo := condicion.(Value).Label2
	fmt.Println(BoolSi, BoolNo)
	if Miceleanos.Isbool(condicion.(Value).Valor) {
		v.Traductor.AddLabel(BoolNo.(string))
		v.EntornoActual = NewEntorno(v.EntornoActual, "Guard")
		for i := 0; ctx.Lista_proceso(i) != nil; i++ {
			value := v.Visit(ctx.Lista_proceso(i))
			if value == "continue" {
				v.Traductor.Br()
				v.Traductor.AddComentario("\t//-----------------------")
				v.Traductor.AddComentario("\t//control CONTINUE")
				v.Traductor.AddGoto(LabelPassSI)
				v.Traductor.Br()
				LabelPassSI = ""
			} else if value == "break" {
				v.Traductor.Br()
				v.Traductor.AddComentario("\t//-----------------------")
				v.Traductor.AddComentario("\t//control BREAK")
				v.Traductor.AddGoto(LabelPassNo)
				v.Traductor.Br()
				LabelPassNo = ""
			}

		}
		v.EntornoActual = v.EntornoActual.Padre
		v.Traductor.AddLabel(BoolSi.(string))

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

	v.Traductor.AddComentario("//-------------------------------------")
	v.Traductor.AddComentario("//Declaracion de Vector")
	NewTemp := v.Traductor.NewTemp()
	if NuevoVector.TipoDate != "String" {
		v.Traductor.Igual(NewTemp, "H", nil, nil)
	}
	var ListTempString = make([]string, 0)
	for i := 0; ctx.Expresion(i) != nil; i++ {

		var valor interface{}
		if NuevoVector.TipoDate == "Int" || NuevoVector.TipoDate == "Float" {
			valor = v.Visit(ctx.Expresion(i))
			valor = valor.(Value).Valor
			v.Traductor.SetHeap("(int)H", fmt.Sprint(valor))
			v.Traductor.Contador("H", "H", "+", "1")
		} else if NuevoVector.TipoDate == "Character" {
			ControlString = true
			valor = v.Visit(ctx.Expresion(i))
			valor = valor.(Value).Valor
			v.Traductor.SetHeap("(int)H", strconv.Itoa(int(valor.(string)[0])))
			v.Traductor.Contador("H", "H", "+", "1")
			ControlString = false
		} else if NuevoVector.TipoDate == "String" {
			valor = v.Visit(ctx.Expresion(i))
			Temp := valor.(Value).Temp.(string)
			valor = valor.(Value).Valor
			ListTempString = append(ListTempString, Temp)
		}
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

		if ctx.Tipo().GetText() == "String" {
			v.Traductor.Igual(NewTemp, "H", nil, nil)
			for i := 0; i < len(ListTempString); i++ {
				fmt.Println(ListTempString[i])
				v.Traductor.SetHeap("(int)H", ListTempString[i])
				v.Traductor.Contador("H", "H", "+", "1")
			}
		}
		v.Traductor.SetHeap("(int)H", fmt.Sprint(-1))
		v.Traductor.Contador("H", "H", "+", "1")

		NewStack := v.Traductor.GetIndex()
		v.Traductor.Br()
		v.Traductor.SetStack(fmt.Sprint(NewStack), fmt.Sprint(NewTemp))
		v.Traductor.Br()
		v.EntornoActual.ActualizarStack(NuevoVector.GetId(), NewStack)

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
					CopiarHeap = true
					v.Traductor.AddComentario("//----------------------------")
					v.Traductor.AddComentario("//Declaracion vector")
					v.Traductor.Br()
					NewTemp := v.Traductor.NewTemp()
					v.Traductor.GetStack(NewTemp, fmt.Sprint(Encontrado.Stack))
					v.Traductor.Br()
					NewList := v.Traductor.NewTemp()
					v.Traductor.Igual(NewList, "H", nil, nil)
					v.Traductor.Br()
					Newcont := v.Traductor.NewTemp()
					v.Traductor.Igual(Newcont, "P", fmt.Sprint(v.EntornoActual.Size), "+")
					v.Traductor.Igual(Newcont, Newcont, "1", "+")
					v.Traductor.SetStack("(int)"+Newcont, NewTemp)
					v.Traductor.Contador("P", "P", "+", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.callVoid("CopiarHeap")
					v.Traductor.Contador("P", "P", "-", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.SetHeap("(int) H", "-1")
					v.Traductor.Contador("H", "H", "+", "1")
					v.Traductor.Br()
					Stack := v.Traductor.GetIndex()
					v.Traductor.SetStack(fmt.Sprint(Stack), NewList)
					v.Traductor.Br()

					v.EntornoActual.ActualizarStack(ctx.ID().GetText(), Stack)

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

			v.Traductor.AddComentario("//-----------------------")
			v.Traductor.AddComentario("//Declaracion de Vector")
			v.Traductor.Br()
			Stack := v.Traductor.NewTemp()
			v.Traductor.Igual(Stack, "H", nil, nil)
			v.Traductor.Br()
			v.Traductor.SetHeap("(int) H", fmt.Sprint(-1))
			v.Traductor.Contador("H", "H", "+", "1")
			Index := v.Traductor.GetIndex()
			v.Traductor.Br()
			v.Traductor.SetStack(fmt.Sprint(Index), Stack)
			v.EntornoActual.ActualizarStack(NuevoVector.Id, Index)

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
				NewVal := v.Visit(ctx.Expresion())

				is_Correct := AppendVector.FunctionAppend(NewVal.(Value).Valor)
				if is_Correct {
					NuevoValor := NewSimbolo(AppendVector.GetList(), AppendVector.GetTipo(), AppendVector.GetTipoVariable(), true)
					v.EntornoActual.ActualizarSimbolo(Id, NuevoValor)
					v.Traductor.AddComentario("\t//-------------------")
					v.Traductor.AddComentario("\t//Vector Append")

					CopiarHeap = true
					v.Traductor.Br()
					NewTempStack := v.Traductor.NewTemp()
					v.Traductor.GetStack(NewTempStack, fmt.Sprint(Encontrado.Stack))
					v.Traductor.Br()
					NewHeap := v.Traductor.NewTemp()
					v.Traductor.Igual(NewHeap, "H", nil, nil)
					v.Traductor.Br()
					NewTemp := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTemp, "P", v.EntornoActual.Size, "+")
					v.Traductor.Contador(NewTemp, NewTemp, "+", "1")
					v.Traductor.SetStack("(int)"+NewTemp, NewTempStack)
					v.Traductor.Contador("P", "P", "+", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.callVoid("CopiarHeap")
					v.Traductor.Contador("P", "P", "-", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.Br()
					v.Traductor.SetHeap("(int)H", fmt.Sprint(NewVal.(Value).Temp))
					v.Traductor.Contador("H", "H", "+", "1")
					v.Traductor.SetHeap("(int)H", fmt.Sprint(-1))
					v.Traductor.Contador("H", "H", "+", "1")
					v.Traductor.Br()
					v.Traductor.SetStack(fmt.Sprint(Encontrado.Stack), NewHeap)
					v.Traductor.Br()

					v.Traductor.AddComentario("\t//Fin Vector Append")
					v.Traductor.AddComentario("\t//-------------------")
					v.EntornoActual.ActualizarStack(Id, Encontrado.Stack)

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

					v.Traductor.AddComentario("\t//-------------------")
					v.Traductor.AddComentario("\t//Vector RemoveLast")
					v.Traductor.Br()
					NewTemp := v.Traductor.NewTemp()
					v.Traductor.GetStack(NewTemp, fmt.Sprint(Encontrado.Stack))
					NewTempPos := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTempPos, "0", nil, nil)
					LabelSI := v.Traductor.NewLabel()
					LabelNO := v.Traductor.NewLabel()
					HeapVal := v.Traductor.NewTemp()
					v.Traductor.AddLabel(LabelNO)
					v.Traductor.GetHeap(HeapVal, "(int)"+NewTemp)
					v.Traductor.AddIf(HeapVal, fmt.Sprint(-1), "==", LabelSI)
					v.Traductor.Contador(NewTempPos, NewTempPos, "+", "1")
					v.Traductor.Contador(NewTemp, NewTemp, "+", "1")
					v.Traductor.AddGoto(LabelNO)
					v.Traductor.AddLabel(LabelSI)
					LabelVacioSI := v.Traductor.NewLabel()
					LabelVacioNO := v.Traductor.NewLabel()
					v.Traductor.AddIf(NewTempPos, fmt.Sprint(0), "!=", LabelVacioSI)
					v.Traductor.AddPrint("118", "c")
					v.Traductor.AddPrint("97", "c")
					v.Traductor.AddPrint("99", "c")
					v.Traductor.AddPrint("105", "c")
					v.Traductor.AddPrint("111", "c")
					v.Traductor.AddPrint("32", "c")
					v.Traductor.AddGoto(LabelVacioNO)
					v.Traductor.AddLabel(LabelVacioSI)
					v.Traductor.Contador(NewTemp, NewTemp, "-", "1")
					v.Traductor.SetHeap("(int)"+NewTemp, fmt.Sprint(-1))
					v.Traductor.AddLabel(LabelVacioNO)

					v.Traductor.AddComentario("\t//Fin Vector RemoveLast")
					v.Traductor.AddComentario("\t//-------------------")
					v.EntornoActual.ActualizarStack(Id, Encontrado.Stack)

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
				Newval := v.Visit(ctx.Expresion())
				is_Correct := RemoveVector.Remove(Newval.(Value).Valor)
				if is_Correct == true {
					NuevoValor := NewSimbolo(RemoveVector.GetList(), RemoveVector.GetTipo(), RemoveVector.GetTipoVariable(), true)
					v.EntornoActual.ActualizarSimbolo(Id, NuevoValor)
					Control_Count = true
					v.Traductor.AddComentario("\t//-------------------")
					v.Traductor.AddComentario("\t//vector removeAt")

					v.Traductor.Br()
					Stack := v.Traductor.NewTemp()
					v.Traductor.GetStack(Stack, fmt.Sprint(Encontrado.Stack))
					Pos := v.Traductor.NewTemp()
					ListaNueva := v.Traductor.NewTemp()
					v.Traductor.Igual(ListaNueva, "H", nil, nil)
					v.Traductor.Igual(Pos, "0", nil, nil)

					NetTemp := v.Traductor.NewTemp()
					v.Traductor.Igual(NetTemp, "P", v.EntornoActual.Size, "+")
					v.Traductor.Contador(NetTemp, NetTemp, "+", "1")
					v.Traductor.SetStack("(int)"+NetTemp, Stack)
					v.Traductor.Contador("P", "P", "+", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.callVoid("VectorCount")
					Size := v.Traductor.NewTemp()
					v.Traductor.GetStack(Size, "(int)P")
					v.Traductor.Contador("P", "P", "-", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.Br()
					LabelError := v.Traductor.NewLabel()
					LabelNoError := v.Traductor.NewLabel()
					v.Traductor.Contador(Size, Size, "-", "1")
					v.Traductor.AddIf(fmt.Sprint(Newval.(Value).Temp), Size, ">", LabelError)
					v.Traductor.AddIf(fmt.Sprint(Newval.(Value).Temp), "0", "<", LabelError)
					v.Traductor.Br()
					Seguir := v.Traductor.NewLabel()
					NoSeguir := v.Traductor.NewLabel()
					v.Traductor.AddLabel(Seguir)
					HeapVal := v.Traductor.NewTemp()
					v.Traductor.GetHeap(HeapVal, "(int)"+Stack)
					v.Traductor.AddIf(HeapVal, fmt.Sprint(-1), "==", NoSeguir)
					v.Traductor.Br()

					IndiceSI := v.Traductor.NewLabel()
					v.Traductor.AddIf(Pos, fmt.Sprint(Newval.(Value).Temp), "==", IndiceSI)
					v.Traductor.Br()
					v.Traductor.SetHeap("(int)H", HeapVal)
					v.Traductor.Contador("H", "H", "+", "1")
					v.Traductor.AddLabel(IndiceSI)
					v.Traductor.Br()

					v.Traductor.Contador(Pos, Pos, "+", "1")
					v.Traductor.Contador(Stack, Stack, "+", "1")
					v.Traductor.AddGoto(Seguir)
					v.Traductor.Br()
					v.Traductor.AddLabel(NoSeguir)
					v.Traductor.SetHeap("(int)H", fmt.Sprint(-1))
					v.Traductor.Contador("H", "H", "+", "1")
					v.Traductor.SetStack(fmt.Sprint(Encontrado.Stack), ListaNueva)
					v.Traductor.Br()

					v.Traductor.AddGoto(LabelNoError)
					v.Traductor.AddLabel(LabelError)
					v.Traductor.AddPrint("66", "c")
					v.Traductor.AddPrint("111", "c")
					v.Traductor.AddPrint("117", "c")
					v.Traductor.AddPrint("110", "c")
					v.Traductor.AddPrint("100", "c")
					v.Traductor.AddPrint("115", "c")
					v.Traductor.AddPrint("69", "c")
					v.Traductor.AddPrint("114", "c")
					v.Traductor.AddPrint("114", "c")
					v.Traductor.AddPrint("111", "c")
					v.Traductor.AddPrint("114", "c")
					v.Traductor.AddPrint("10", "c")
					v.Traductor.AddLabel(LabelNoError)

					v.Traductor.AddComentario("\t//Fin vector removeAt")
					v.Traductor.AddComentario("\t//-------------------")
					v.EntornoActual.ActualizarStack(Id, Encontrado.Stack)

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
	Control_Count = true
	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarSimbolo(Id)
	CountVector := Proceso.NewVector()
	if Encontrado.EsVector {

		CountVector.TipoDate = Encontrado.Type
		CountVector.List = Encontrado.Value.([]interface{})
		CountVector.Id = Id
		CountVector.TipoVariable = Encontrado.TypeValue
	}

	v.Traductor.AddComentario("//-------------------------------------")
	v.Traductor.AddComentario("//Vector Count")
	v.Traductor.Br()
	NewTemp := v.Traductor.NewTemp()
	v.Traductor.GetStack(NewTemp, fmt.Sprint(Encontrado.Stack))
	NewTemp2 := v.Traductor.NewTemp()
	v.Traductor.Igual(NewTemp2, "P", v.EntornoActual.Size, "+")
	v.Traductor.Contador(NewTemp2, NewTemp2, "+", "1")
	v.Traductor.SetStack("(int)"+fmt.Sprint(NewTemp2), NewTemp)
	v.Traductor.Contador("P", "P", "+", fmt.Sprint(v.EntornoActual.Size))
	v.Traductor.callVoid("VectorCount")
	NewTemp3 := v.Traductor.NewTemp()
	v.Traductor.GetStack(fmt.Sprint(NewTemp3), "(int) P")
	v.Traductor.Contador("P", "P", "-", fmt.Sprint(v.EntornoActual.Size))

	return NewValue(NewTemp3, nil, nil, CountVector.Count())
}

func (v *Visitor) VisitVectorVacio(ctx *parser.VectorVacioContext) interface{} {
	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarSimbolo(Id)
	CountVector := Proceso.NewVector()
	LabelVacio := ""
	LabelNoVacio := ""
	if Encontrado.EsVector {
		TypeBoolC3D = 2
		Control_Count = true
		CountVector.TipoDate = Encontrado.Type
		CountVector.List = Encontrado.Value.([]interface{})
		CountVector.Id = Id
		CountVector.TipoVariable = Encontrado.TypeValue
		v.Traductor.AddComentario("//------------------")
		v.Traductor.AddComentario("//Vector EsVacio")
		v.Traductor.Br()
		NewTemp := v.Traductor.NewTemp()
		v.Traductor.GetStack(NewTemp, fmt.Sprint(Encontrado.Stack))
		v.Traductor.Br()
		NewTemp2 := v.Traductor.NewTemp()
		v.Traductor.Igual(NewTemp2, "P", v.EntornoActual.Size, "+")
		v.Traductor.Contador(NewTemp2, NewTemp2, "+", "1")
		v.Traductor.SetStack("(int)"+fmt.Sprint(NewTemp2), NewTemp)
		v.Traductor.Contador("P", "P", "+", fmt.Sprint(v.EntornoActual.Size))
		v.Traductor.callVoid("VectorCount")
		NewTemp3 := v.Traductor.NewTemp()
		v.Traductor.GetStack(fmt.Sprint(NewTemp3), "(int) P")
		v.Traductor.Contador("P", "P", "-", fmt.Sprint(v.EntornoActual.Size))

		LabelVacio = v.Traductor.NewLabel()
		LabelNoVacio = v.Traductor.NewLabel()
		v.Traductor.AddIf(NewTemp3, fmt.Sprint(0), "==", LabelVacio)
		v.Traductor.AddGoto(LabelNoVacio)

		v.Traductor.AddComentario("//Fin Vector EsVacio")
		v.Traductor.AddComentario("//------------------")

	}
	return NewValue(nil, LabelVacio, LabelNoVacio, CountVector.IsVacio())
}

func (v *Visitor) VisitVectorAsignacion(ctx *parser.VectorAsignacionContext) interface{} {
	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarSimbolo(Id)
	if !SymbolVacio(Encontrado) {
		v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "No existe la variable: "+Id, "SEMANTICO")
		return false
	}

	List := Encontrado.Value.([]interface{})
	NewV := v.Visit(ctx.Expresion())
	TEMP := ""
	if int64(len(List)) < 0 {
		v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "ListaVacia", "SEMANTICO")
		return false
	} else if NewV.(Value).Valor.(int64) > int64(len(List)) {
		v.AddError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "VECTOR", "Fuera de rango", "SEMANTICO")
		return false
	}
	Control_Count = true
	v.Traductor.AddComentario("\t//--------------------")
	v.Traductor.AddComentario("\t//GetVector")
	v.Traductor.Br()
	Stack := v.Traductor.NewTemp()
	v.Traductor.GetStack(Stack, fmt.Sprint(Encontrado.Stack))
	v.Traductor.Br()
	Index := v.Traductor.NewTemp()
	v.Traductor.Igual(Index, NewV.(Value).Temp, nil, nil)
	v.Traductor.Br()
	NewTemp := v.Traductor.NewTemp()
	v.Traductor.Igual(NewTemp, "P", fmt.Sprint(v.EntornoActual.Size), "+")
	v.Traductor.Contador(NewTemp, NewTemp, "+", "1")
	v.Traductor.SetStack("(int)"+NewTemp, Stack)
	v.Traductor.Contador("P", "P", "+", fmt.Sprint(v.EntornoActual.Size))
	v.Traductor.callVoid("VectorCount")
	Size := v.Traductor.NewTemp()
	v.Traductor.GetStack(Size, "(int)P")
	v.Traductor.Contador("P", "P", "-", fmt.Sprint(v.EntornoActual.Size))
	v.Traductor.Br()
	MsgError := v.Traductor.NewLabel()
	Fin := v.Traductor.NewLabel()
	NoError := v.Traductor.NewLabel()
	v.Traductor.Contador(Size, Size, "-", "1")
	v.Traductor.AddIf(Index, Size, ">", MsgError)
	v.Traductor.AddIf(Index, "0", "<", MsgError)
	v.Traductor.AddGoto(NoError)
	v.Traductor.AddLabel(MsgError)
	v.Traductor.AddPrint("66", "c")
	v.Traductor.AddPrint("111", "c")
	v.Traductor.AddPrint("117", "c")
	v.Traductor.AddPrint("110", "c")
	v.Traductor.AddPrint("100", "c")
	v.Traductor.AddPrint("115", "c")
	v.Traductor.AddPrint("69", "c")
	v.Traductor.AddPrint("114", "c")
	v.Traductor.AddPrint("114", "c")
	v.Traductor.AddPrint("111", "c")
	v.Traductor.AddPrint("114", "c")
	v.Traductor.AddPrint("10", "c")
	v.Traductor.AddGoto(Fin)
	v.Traductor.AddLabel(NoError)
	v.Traductor.Br()
	NewTempGet := v.Traductor.NewTemp()
	v.Traductor.Igual(NewTempGet, Index, Stack, "+")
	TEMP = v.Traductor.NewTemp()
	v.Traductor.GetHeap(TEMP, "(int)"+NewTempGet)

	v.Traductor.AddLabel(Fin)
	v.Traductor.AddComentario("\t// Fin GetVector")
	v.Traductor.AddComentario("\t//-----------------------")
	v.Traductor.Br()

	return NewValue(TEMP, nil, nil, List[NewV.(Value).Valor.(int64)])
}

func (v *Visitor) VisitAsignacionVector(ctx *parser.AsignacionVectorContext) interface{} {

	Id := ctx.ID().GetText()
	Encontrado := v.EntornoActual.BuscarSimbolo(Id)
	if SymbolVacio(Encontrado) {
		if Encontrado.TypeValue == "var" {
			if ctx.Subasig().NUMBER() != nil {
				fmt.Println("ASIGNACION VECTOR")
				i, _ := strconv.ParseInt(ctx.Subasig().NUMBER().GetText(), 10, 64)
				List := Encontrado.Value.([]interface{})

				v.Traductor.AddComentario("\t//--------------------")
				v.Traductor.AddComentario("\t//Asignacion Vector")
				v.Traductor.Br()
				ValorIndex := v.Traductor.NewTemp()
				v.Traductor.Igual(ValorIndex, fmt.Sprint(i), nil, nil)
				v.Traductor.Br()
				Stack := v.Traductor.NewTemp()
				v.Traductor.GetStack(Stack, fmt.Sprint(Encontrado.Stack))
				if ctx.GetOp().GetText() == "=" {
					Control_Count = true
					NewVal := v.Visit(ctx.Expresion())

					v.Traductor.Br()
					NewTemp := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTemp, "P", fmt.Sprint(v.EntornoActual.Size), "+")
					v.Traductor.Contador(NewTemp, NewTemp, "+", "1")
					v.Traductor.SetStack("(int)"+NewTemp, Stack)
					v.Traductor.Contador("P", "P", "+", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.callVoid("VectorCount")
					Size := v.Traductor.NewTemp()
					v.Traductor.GetStack(Size, "(int)P")
					v.Traductor.Contador("P", "P", "-", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.Br()

					MsgError := v.Traductor.NewLabel()
					Fin := v.Traductor.NewLabel()
					NoError := v.Traductor.NewLabel()
					v.Traductor.Contador(Size, Size, "-", "1")
					v.Traductor.AddIf(ValorIndex, Size, ">", MsgError)
					v.Traductor.AddIf(ValorIndex, "0", "<", MsgError)
					v.Traductor.AddGoto(NoError)
					v.Traductor.AddLabel(MsgError)
					v.Traductor.AddPrint("66", "c")
					v.Traductor.AddPrint("111", "c")
					v.Traductor.AddPrint("117", "c")
					v.Traductor.AddPrint("110", "c")
					v.Traductor.AddPrint("100", "c")
					v.Traductor.AddPrint("115", "c")
					v.Traductor.AddPrint("69", "c")
					v.Traductor.AddPrint("114", "c")
					v.Traductor.AddPrint("114", "c")
					v.Traductor.AddPrint("111", "c")
					v.Traductor.AddPrint("114", "c")
					v.Traductor.AddPrint("10", "c")
					v.Traductor.AddGoto(Fin)
					v.Traductor.AddLabel(NoError)
					v.Traductor.Br()
					NewTempSet := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTempSet, ValorIndex, Stack, "+")
					v.Traductor.SetHeap("(int)"+NewTempSet, fmt.Sprint(NewVal.(Value).Temp))
					v.Traductor.AddLabel(Fin)
					v.Traductor.AddComentario("\t// Fin Asignacion Vector")
					v.Traductor.AddComentario("\t//-----------------------")
					v.Traductor.Br()

					List[i] = NewVal.(Value).Valor
					v.EntornoActual.ActualizarSimbolo(Id, NewSimbolo(List, Encontrado.Type, Encontrado.TypeValue, true))
					v.EntornoActual.ActualizarStack(Id, Encontrado.Stack)

				}

			} else {
				iter := v.EntornoActual.BuscarSimbolo(ctx.Subasig().ID().GetText())
				i := iter.Value.(int64)
				List := Encontrado.Value.([]interface{})
				if ctx.GetOp().GetText() == "=" {
					NewVal := v.Visit(ctx.Expresion())
					Control_Count = true

					v.Traductor.AddComentario("\t//--------------------")
					v.Traductor.AddComentario("\t//Asignacion Vector")
					v.Traductor.Br()
					ValorIndex := v.Traductor.NewTemp()
					v.Traductor.Igual(ValorIndex, fmt.Sprint(i), nil, nil)
					v.Traductor.Br()
					Stack := v.Traductor.NewTemp()
					v.Traductor.GetStack(Stack, fmt.Sprint(Encontrado.Stack))
					v.Traductor.Br()
					NewTemp := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTemp, "P", fmt.Sprint(v.EntornoActual.Size), "+")
					v.Traductor.Contador(NewTemp, NewTemp, "+", "1")
					v.Traductor.SetStack("(int)"+NewTemp, Stack)
					v.Traductor.Contador("P", "P", "+", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.callVoid("VectorCount")
					Size := v.Traductor.NewTemp()
					v.Traductor.GetStack(Size, "(int)P")
					v.Traductor.Contador("P", "P", "-", fmt.Sprint(v.EntornoActual.Size))
					v.Traductor.Br()

					MsgError := v.Traductor.NewLabel()
					Fin := v.Traductor.NewLabel()
					NoError := v.Traductor.NewLabel()
					v.Traductor.Contador(Size, Size, "-", "1")
					v.Traductor.AddIf(ValorIndex, Size, ">", MsgError)
					v.Traductor.AddIf(ValorIndex, "0", "<", MsgError)
					v.Traductor.AddGoto(NoError)
					v.Traductor.AddLabel(MsgError)
					v.Traductor.AddPrint("66", "c")
					v.Traductor.AddPrint("111", "c")
					v.Traductor.AddPrint("117", "c")
					v.Traductor.AddPrint("110", "c")
					v.Traductor.AddPrint("100", "c")
					v.Traductor.AddPrint("115", "c")
					v.Traductor.AddPrint("69", "c")
					v.Traductor.AddPrint("114", "c")
					v.Traductor.AddPrint("114", "c")
					v.Traductor.AddPrint("111", "c")
					v.Traductor.AddPrint("114", "c")
					v.Traductor.AddPrint("10", "c")
					v.Traductor.AddGoto(Fin)
					v.Traductor.AddLabel(NoError)
					v.Traductor.Br()
					NewTempSet := v.Traductor.NewTemp()
					v.Traductor.Igual(NewTempSet, ValorIndex, Stack, "+")
					v.Traductor.SetHeap("(int)"+NewTempSet, fmt.Sprint(NewVal.(Value).Temp))
					v.Traductor.AddLabel(Fin)
					v.Traductor.AddComentario("\t// Fin Asignacion Vector")
					v.Traductor.AddComentario("\t//-----------------------")
					v.Traductor.Br()

					List[i] = NewVal.(Value).Valor
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

func (v *Visitor) VisitComentarios(ctx *parser.ComentariosContext) interface{} {
	if ctx.COMMENT() != nil {
		v.Traductor.List_Process = append(v.Traductor.List_Process, ctx.COMMENT().GetText()+"\n")
	}
	if ctx.COMMENT_MULT() != nil {
		v.Traductor.List_Process = append(v.Traductor.List_Process, ctx.COMMENT_MULT().GetText()+"\n")
	}
	return true
}
