// Code generated from C:\Users\LENOVO\Desktop\PY2-OLC2-202006373\OLC2_PY2\backend-app\gramatica\Sintax.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Sintax
import "github.com/antlr4-go/antlr/v4"

type BaseSintaxVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSintaxVisitor) VisitInicio(ctx *InicioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitLista(ctx *ListaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitLista_proceso(ctx *Lista_procesoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitComentarios(ctx *ComentariosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitRetornar(ctx *RetornarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitDeclaracion(ctx *DeclaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitDeclaracionTipo(ctx *DeclaracionTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitDeclaracionTipoImplicito(ctx *DeclaracionTipoImplicitoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitAsignacion(ctx *AsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitAsignacionVariable(ctx *AsignacionVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitAsignacionVector(ctx *AsignacionVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitSubasig(ctx *SubasigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitPrint(ctx *PrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitIf(ctx *IfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitElseif_(ctx *Elseif_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitElse(ctx *ElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitSwitch(ctx *SwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitCaso(ctx *CasoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitDefault(ctx *DefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitWhile(ctx *WhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitFor(ctx *ForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitGuard(ctx *GuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitDec_vector(ctx *Dec_vectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitDec_vector_V_C(ctx *Dec_vector_V_CContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitSubVC(ctx *SubVCContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitFuncvectorList(ctx *FuncvectorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitTipevector(ctx *TipevectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitBloque(ctx *BloqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitControl(ctx *ControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitVectorCount(ctx *VectorCountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitOpLogico(ctx *OpLogicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitSinValor(ctx *SinValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitVectorAsignacion(ctx *VectorAsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitOpAritmetico(ctx *OpAritmeticoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitNumberValor(ctx *NumberValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitValorsimple(ctx *ValorsimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitCasteo(ctx *CasteoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitVectorVacio(ctx *VectorVacioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitParentexpr(ctx *ParentexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitSelfCALL(ctx *SelfCALLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitStringValor(ctx *StringValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitOpRelacional(ctx *OpRelacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitBoolValor(ctx *BoolValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitPrintLLamada(ctx *PrintLLamadaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitOpLogicoNot(ctx *OpLogicoNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitIdValor(ctx *IdValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitOpComparativo(ctx *OpComparativoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitFunciones(ctx *FuncionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitParametros(ctx *ParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitExisteExInt(ctx *ExisteExIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitTipofuncion(ctx *TipofuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitTipoinout(ctx *TipoinoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitFuncLLamada(ctx *FuncLLamadaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitParametrosLLamada(ctx *ParametrosLLamadaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitEstructs(ctx *EstructsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitLista_struct(ctx *Lista_structContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSintaxVisitor) VisitFuncionestruct(ctx *FuncionestructContext) interface{} {
	return v.VisitChildren(ctx)
}
