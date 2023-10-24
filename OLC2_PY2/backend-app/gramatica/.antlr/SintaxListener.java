// Generated from c:/Users/LENOVO/Desktop/PY2-OLC2-202006373/OLC2_PY2/backend-app/gramatica/Sintax.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SintaxParser}.
 */
public interface SintaxListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SintaxParser#inicio}.
	 * @param ctx the parse tree
	 */
	void enterInicio(SintaxParser.InicioContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#inicio}.
	 * @param ctx the parse tree
	 */
	void exitInicio(SintaxParser.InicioContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#lista}.
	 * @param ctx the parse tree
	 */
	void enterLista(SintaxParser.ListaContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#lista}.
	 * @param ctx the parse tree
	 */
	void exitLista(SintaxParser.ListaContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#lista_proceso}.
	 * @param ctx the parse tree
	 */
	void enterLista_proceso(SintaxParser.Lista_procesoContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#lista_proceso}.
	 * @param ctx the parse tree
	 */
	void exitLista_proceso(SintaxParser.Lista_procesoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#comentarios}.
	 * @param ctx the parse tree
	 */
	void enterComentarios(SintaxParser.ComentariosContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#comentarios}.
	 * @param ctx the parse tree
	 */
	void exitComentarios(SintaxParser.ComentariosContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#retornar}.
	 * @param ctx the parse tree
	 */
	void enterRetornar(SintaxParser.RetornarContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#retornar}.
	 * @param ctx the parse tree
	 */
	void exitRetornar(SintaxParser.RetornarContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracion(SintaxParser.DeclaracionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracion(SintaxParser.DeclaracionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclaracionTipo}
	 * labeled alternative in {@link SintaxParser#dec_tipo_valor}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracionTipo(SintaxParser.DeclaracionTipoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclaracionTipo}
	 * labeled alternative in {@link SintaxParser#dec_tipo_valor}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracionTipo(SintaxParser.DeclaracionTipoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclaracionTipoImplicito}
	 * labeled alternative in {@link SintaxParser#dec_valor}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracionTipoImplicito(SintaxParser.DeclaracionTipoImplicitoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclaracionTipoImplicito}
	 * labeled alternative in {@link SintaxParser#dec_valor}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracionTipoImplicito(SintaxParser.DeclaracionTipoImplicitoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion(SintaxParser.AsignacionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion(SintaxParser.AsignacionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#asignacionVariable}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionVariable(SintaxParser.AsignacionVariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#asignacionVariable}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionVariable(SintaxParser.AsignacionVariableContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#asignacionVector}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionVector(SintaxParser.AsignacionVectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#asignacionVector}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionVector(SintaxParser.AsignacionVectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#subasig}.
	 * @param ctx the parse tree
	 */
	void enterSubasig(SintaxParser.SubasigContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#subasig}.
	 * @param ctx the parse tree
	 */
	void exitSubasig(SintaxParser.SubasigContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#print}.
	 * @param ctx the parse tree
	 */
	void enterPrint(SintaxParser.PrintContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#print}.
	 * @param ctx the parse tree
	 */
	void exitPrint(SintaxParser.PrintContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#if}.
	 * @param ctx the parse tree
	 */
	void enterIf(SintaxParser.IfContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#if}.
	 * @param ctx the parse tree
	 */
	void exitIf(SintaxParser.IfContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#elseif_}.
	 * @param ctx the parse tree
	 */
	void enterElseif_(SintaxParser.Elseif_Context ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#elseif_}.
	 * @param ctx the parse tree
	 */
	void exitElseif_(SintaxParser.Elseif_Context ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#else}.
	 * @param ctx the parse tree
	 */
	void enterElse(SintaxParser.ElseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#else}.
	 * @param ctx the parse tree
	 */
	void exitElse(SintaxParser.ElseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#switch}.
	 * @param ctx the parse tree
	 */
	void enterSwitch(SintaxParser.SwitchContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#switch}.
	 * @param ctx the parse tree
	 */
	void exitSwitch(SintaxParser.SwitchContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#caso}.
	 * @param ctx the parse tree
	 */
	void enterCaso(SintaxParser.CasoContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#caso}.
	 * @param ctx the parse tree
	 */
	void exitCaso(SintaxParser.CasoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#default}.
	 * @param ctx the parse tree
	 */
	void enterDefault(SintaxParser.DefaultContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#default}.
	 * @param ctx the parse tree
	 */
	void exitDefault(SintaxParser.DefaultContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#while}.
	 * @param ctx the parse tree
	 */
	void enterWhile(SintaxParser.WhileContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#while}.
	 * @param ctx the parse tree
	 */
	void exitWhile(SintaxParser.WhileContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#for}.
	 * @param ctx the parse tree
	 */
	void enterFor(SintaxParser.ForContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#for}.
	 * @param ctx the parse tree
	 */
	void exitFor(SintaxParser.ForContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#guard}.
	 * @param ctx the parse tree
	 */
	void enterGuard(SintaxParser.GuardContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#guard}.
	 * @param ctx the parse tree
	 */
	void exitGuard(SintaxParser.GuardContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#dec_vector}.
	 * @param ctx the parse tree
	 */
	void enterDec_vector(SintaxParser.Dec_vectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#dec_vector}.
	 * @param ctx the parse tree
	 */
	void exitDec_vector(SintaxParser.Dec_vectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#dec_vector_V_C}.
	 * @param ctx the parse tree
	 */
	void enterDec_vector_V_C(SintaxParser.Dec_vector_V_CContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#dec_vector_V_C}.
	 * @param ctx the parse tree
	 */
	void exitDec_vector_V_C(SintaxParser.Dec_vector_V_CContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#subVC}.
	 * @param ctx the parse tree
	 */
	void enterSubVC(SintaxParser.SubVCContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#subVC}.
	 * @param ctx the parse tree
	 */
	void exitSubVC(SintaxParser.SubVCContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#funcvectorList}.
	 * @param ctx the parse tree
	 */
	void enterFuncvectorList(SintaxParser.FuncvectorListContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#funcvectorList}.
	 * @param ctx the parse tree
	 */
	void exitFuncvectorList(SintaxParser.FuncvectorListContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#tipevector}.
	 * @param ctx the parse tree
	 */
	void enterTipevector(SintaxParser.TipevectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#tipevector}.
	 * @param ctx the parse tree
	 */
	void exitTipevector(SintaxParser.TipevectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#bloque}.
	 * @param ctx the parse tree
	 */
	void enterBloque(SintaxParser.BloqueContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#bloque}.
	 * @param ctx the parse tree
	 */
	void exitBloque(SintaxParser.BloqueContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#control}.
	 * @param ctx the parse tree
	 */
	void enterControl(SintaxParser.ControlContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#control}.
	 * @param ctx the parse tree
	 */
	void exitControl(SintaxParser.ControlContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo(SintaxParser.TipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo(SintaxParser.TipoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorCount}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterVectorCount(SintaxParser.VectorCountContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorCount}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitVectorCount(SintaxParser.VectorCountContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OpLogico}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterOpLogico(SintaxParser.OpLogicoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OpLogico}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitOpLogico(SintaxParser.OpLogicoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SinValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterSinValor(SintaxParser.SinValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SinValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitSinValor(SintaxParser.SinValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorAsignacion}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterVectorAsignacion(SintaxParser.VectorAsignacionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorAsignacion}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitVectorAsignacion(SintaxParser.VectorAsignacionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OpAritmetico}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterOpAritmetico(SintaxParser.OpAritmeticoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OpAritmetico}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitOpAritmetico(SintaxParser.OpAritmeticoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NumberValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterNumberValor(SintaxParser.NumberValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NumberValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitNumberValor(SintaxParser.NumberValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorsimple}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterValorsimple(SintaxParser.ValorsimpleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorsimple}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitValorsimple(SintaxParser.ValorsimpleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Casteo}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterCasteo(SintaxParser.CasteoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Casteo}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitCasteo(SintaxParser.CasteoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorVacio}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterVectorVacio(SintaxParser.VectorVacioContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorVacio}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitVectorVacio(SintaxParser.VectorVacioContext ctx);
	/**
	 * Enter a parse tree produced by the {@code parentexpr}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterParentexpr(SintaxParser.ParentexprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code parentexpr}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitParentexpr(SintaxParser.ParentexprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code selfCALL}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterSelfCALL(SintaxParser.SelfCALLContext ctx);
	/**
	 * Exit a parse tree produced by the {@code selfCALL}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitSelfCALL(SintaxParser.SelfCALLContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StringValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterStringValor(SintaxParser.StringValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StringValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitStringValor(SintaxParser.StringValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OpRelacional}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterOpRelacional(SintaxParser.OpRelacionalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OpRelacional}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitOpRelacional(SintaxParser.OpRelacionalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BoolValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterBoolValor(SintaxParser.BoolValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BoolValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitBoolValor(SintaxParser.BoolValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code printLLamada}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterPrintLLamada(SintaxParser.PrintLLamadaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code printLLamada}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitPrintLLamada(SintaxParser.PrintLLamadaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OpLogicoNot}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterOpLogicoNot(SintaxParser.OpLogicoNotContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OpLogicoNot}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitOpLogicoNot(SintaxParser.OpLogicoNotContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterIdValor(SintaxParser.IdValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdValor}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitIdValor(SintaxParser.IdValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OpComparativo}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterOpComparativo(SintaxParser.OpComparativoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OpComparativo}
	 * labeled alternative in {@link SintaxParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitOpComparativo(SintaxParser.OpComparativoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#funciones}.
	 * @param ctx the parse tree
	 */
	void enterFunciones(SintaxParser.FuncionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#funciones}.
	 * @param ctx the parse tree
	 */
	void exitFunciones(SintaxParser.FuncionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#parametros}.
	 * @param ctx the parse tree
	 */
	void enterParametros(SintaxParser.ParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#parametros}.
	 * @param ctx the parse tree
	 */
	void exitParametros(SintaxParser.ParametrosContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#existeExInt}.
	 * @param ctx the parse tree
	 */
	void enterExisteExInt(SintaxParser.ExisteExIntContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#existeExInt}.
	 * @param ctx the parse tree
	 */
	void exitExisteExInt(SintaxParser.ExisteExIntContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#tipofuncion}.
	 * @param ctx the parse tree
	 */
	void enterTipofuncion(SintaxParser.TipofuncionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#tipofuncion}.
	 * @param ctx the parse tree
	 */
	void exitTipofuncion(SintaxParser.TipofuncionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#tipoinout}.
	 * @param ctx the parse tree
	 */
	void enterTipoinout(SintaxParser.TipoinoutContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#tipoinout}.
	 * @param ctx the parse tree
	 */
	void exitTipoinout(SintaxParser.TipoinoutContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#funcLLamada}.
	 * @param ctx the parse tree
	 */
	void enterFuncLLamada(SintaxParser.FuncLLamadaContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#funcLLamada}.
	 * @param ctx the parse tree
	 */
	void exitFuncLLamada(SintaxParser.FuncLLamadaContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#parametrosLLamada}.
	 * @param ctx the parse tree
	 */
	void enterParametrosLLamada(SintaxParser.ParametrosLLamadaContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#parametrosLLamada}.
	 * @param ctx the parse tree
	 */
	void exitParametrosLLamada(SintaxParser.ParametrosLLamadaContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#estructs}.
	 * @param ctx the parse tree
	 */
	void enterEstructs(SintaxParser.EstructsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#estructs}.
	 * @param ctx the parse tree
	 */
	void exitEstructs(SintaxParser.EstructsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#lista_struct}.
	 * @param ctx the parse tree
	 */
	void enterLista_struct(SintaxParser.Lista_structContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#lista_struct}.
	 * @param ctx the parse tree
	 */
	void exitLista_struct(SintaxParser.Lista_structContext ctx);
	/**
	 * Enter a parse tree produced by {@link SintaxParser#funcionestruct}.
	 * @param ctx the parse tree
	 */
	void enterFuncionestruct(SintaxParser.FuncionestructContext ctx);
	/**
	 * Exit a parse tree produced by {@link SintaxParser#funcionestruct}.
	 * @param ctx the parse tree
	 */
	void exitFuncionestruct(SintaxParser.FuncionestructContext ctx);
}