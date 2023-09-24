// Code generated from C:\Users\LENOVO\Desktop\OLC2_PY1\backend-app\grammar\Sintax.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Sintax
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SintaxParser.
type SintaxVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SintaxParser#inicio.
	VisitInicio(ctx *InicioContext) interface{}

	// Visit a parse tree produced by SintaxParser#lista.
	VisitLista(ctx *ListaContext) interface{}

	// Visit a parse tree produced by SintaxParser#lista_proceso.
	VisitLista_proceso(ctx *Lista_procesoContext) interface{}

	// Visit a parse tree produced by SintaxParser#retornar.
	VisitRetornar(ctx *RetornarContext) interface{}

	// Visit a parse tree produced by SintaxParser#declaracion.
	VisitDeclaracion(ctx *DeclaracionContext) interface{}

	// Visit a parse tree produced by SintaxParser#asigtipo1.
	VisitAsigtipo1(ctx *Asigtipo1Context) interface{}

	// Visit a parse tree produced by SintaxParser#asigtipo2.
	VisitAsigtipo2(ctx *Asigtipo2Context) interface{}

	// Visit a parse tree produced by SintaxParser#asignacion.
	VisitAsignacion(ctx *AsignacionContext) interface{}

	// Visit a parse tree produced by SintaxParser#asignacionVariable.
	VisitAsignacionVariable(ctx *AsignacionVariableContext) interface{}

	// Visit a parse tree produced by SintaxParser#asignacionVector.
	VisitAsignacionVector(ctx *AsignacionVectorContext) interface{}

	// Visit a parse tree produced by SintaxParser#subasig.
	VisitSubasig(ctx *SubasigContext) interface{}

	// Visit a parse tree produced by SintaxParser#print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by SintaxParser#if.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by SintaxParser#elseif_.
	VisitElseif_(ctx *Elseif_Context) interface{}

	// Visit a parse tree produced by SintaxParser#else.
	VisitElse(ctx *ElseContext) interface{}

	// Visit a parse tree produced by SintaxParser#switch.
	VisitSwitch(ctx *SwitchContext) interface{}

	// Visit a parse tree produced by SintaxParser#caso.
	VisitCaso(ctx *CasoContext) interface{}

	// Visit a parse tree produced by SintaxParser#default.
	VisitDefault(ctx *DefaultContext) interface{}

	// Visit a parse tree produced by SintaxParser#while.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by SintaxParser#for.
	VisitFor(ctx *ForContext) interface{}

	// Visit a parse tree produced by SintaxParser#guard.
	VisitGuard(ctx *GuardContext) interface{}

	// Visit a parse tree produced by SintaxParser#dec_vector.
	VisitDec_vector(ctx *Dec_vectorContext) interface{}

	// Visit a parse tree produced by SintaxParser#dec_vector_V_C.
	VisitDec_vector_V_C(ctx *Dec_vector_V_CContext) interface{}

	// Visit a parse tree produced by SintaxParser#subVC.
	VisitSubVC(ctx *SubVCContext) interface{}

	// Visit a parse tree produced by SintaxParser#funcvectorList.
	VisitFuncvectorList(ctx *FuncvectorListContext) interface{}

	// Visit a parse tree produced by SintaxParser#tipevector.
	VisitTipevector(ctx *TipevectorContext) interface{}

	// Visit a parse tree produced by SintaxParser#bloque.
	VisitBloque(ctx *BloqueContext) interface{}

	// Visit a parse tree produced by SintaxParser#control.
	VisitControl(ctx *ControlContext) interface{}

	// Visit a parse tree produced by SintaxParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by SintaxParser#VectorCount.
	VisitVectorCount(ctx *VectorCountContext) interface{}

	// Visit a parse tree produced by SintaxParser#OpLogico.
	VisitOpLogico(ctx *OpLogicoContext) interface{}

	// Visit a parse tree produced by SintaxParser#SinValor.
	VisitSinValor(ctx *SinValorContext) interface{}

	// Visit a parse tree produced by SintaxParser#VectorAsignacion.
	VisitVectorAsignacion(ctx *VectorAsignacionContext) interface{}

	// Visit a parse tree produced by SintaxParser#OpAritmetico.
	VisitOpAritmetico(ctx *OpAritmeticoContext) interface{}

	// Visit a parse tree produced by SintaxParser#NumberValor.
	VisitNumberValor(ctx *NumberValorContext) interface{}

	// Visit a parse tree produced by SintaxParser#valorsimple.
	VisitValorsimple(ctx *ValorsimpleContext) interface{}

	// Visit a parse tree produced by SintaxParser#Casteo.
	VisitCasteo(ctx *CasteoContext) interface{}

	// Visit a parse tree produced by SintaxParser#VectorVacio.
	VisitVectorVacio(ctx *VectorVacioContext) interface{}

	// Visit a parse tree produced by SintaxParser#parentexpr.
	VisitParentexpr(ctx *ParentexprContext) interface{}

	// Visit a parse tree produced by SintaxParser#selfCALL.
	VisitSelfCALL(ctx *SelfCALLContext) interface{}

	// Visit a parse tree produced by SintaxParser#StringValor.
	VisitStringValor(ctx *StringValorContext) interface{}

	// Visit a parse tree produced by SintaxParser#OpRelacional.
	VisitOpRelacional(ctx *OpRelacionalContext) interface{}

	// Visit a parse tree produced by SintaxParser#BoolValor.
	VisitBoolValor(ctx *BoolValorContext) interface{}

	// Visit a parse tree produced by SintaxParser#printLLamada.
	VisitPrintLLamada(ctx *PrintLLamadaContext) interface{}

	// Visit a parse tree produced by SintaxParser#OpLogicoNot.
	VisitOpLogicoNot(ctx *OpLogicoNotContext) interface{}

	// Visit a parse tree produced by SintaxParser#IdValor.
	VisitIdValor(ctx *IdValorContext) interface{}

	// Visit a parse tree produced by SintaxParser#OpComparativo.
	VisitOpComparativo(ctx *OpComparativoContext) interface{}

	// Visit a parse tree produced by SintaxParser#funciones.
	VisitFunciones(ctx *FuncionesContext) interface{}

	// Visit a parse tree produced by SintaxParser#parametros.
	VisitParametros(ctx *ParametrosContext) interface{}

	// Visit a parse tree produced by SintaxParser#existeExInt.
	VisitExisteExInt(ctx *ExisteExIntContext) interface{}

	// Visit a parse tree produced by SintaxParser#tipofuncion.
	VisitTipofuncion(ctx *TipofuncionContext) interface{}

	// Visit a parse tree produced by SintaxParser#tipoinout.
	VisitTipoinout(ctx *TipoinoutContext) interface{}

	// Visit a parse tree produced by SintaxParser#funcLLamada.
	VisitFuncLLamada(ctx *FuncLLamadaContext) interface{}

	// Visit a parse tree produced by SintaxParser#parametrosLLamada.
	VisitParametrosLLamada(ctx *ParametrosLLamadaContext) interface{}

	// Visit a parse tree produced by SintaxParser#estructs.
	VisitEstructs(ctx *EstructsContext) interface{}

	// Visit a parse tree produced by SintaxParser#lista_struct.
	VisitLista_struct(ctx *Lista_structContext) interface{}

	// Visit a parse tree produced by SintaxParser#funcionestruct.
	VisitFuncionestruct(ctx *FuncionestructContext) interface{}
}
