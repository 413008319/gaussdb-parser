// Code generated from GaussdbParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // GaussdbParser
import "github.com/antlr4-go/antlr/v4"

type BaseGaussdbParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGaussdbParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPlsqlroot(ctx *PlsqlrootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmtblock(ctx *StmtblockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmtmulti(ctx *StmtmultiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPlsqlconsolecommand(ctx *PlsqlconsolecommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCallstmt(ctx *CallstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreaterolestmt(ctx *CreaterolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_with(ctx *Opt_withContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptrolelist(ctx *OptrolelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlteroptrolelist(ctx *AlteroptrolelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlteroptroleelem(ctx *AlteroptroleelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateoptroleelem(ctx *CreateoptroleelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateuserstmt(ctx *CreateuserstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterrolestmt(ctx *AlterrolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_in_database(ctx *Opt_in_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterrolesetstmt(ctx *AlterrolesetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDroprolestmt(ctx *DroprolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreategroupstmt(ctx *CreategroupstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltergroupstmt(ctx *AltergroupstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAdd_drop(ctx *Add_dropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateschemastmt(ctx *CreateschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptschemaname(ctx *OptschemanameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptschemaeltlist(ctx *OptschemaeltlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSchema_stmt(ctx *Schema_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVariablesetstmt(ctx *VariablesetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSet_rest(ctx *Set_restContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGeneric_set(ctx *Generic_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSet_rest_more(ctx *Set_rest_moreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVar_name(ctx *Var_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVar_list(ctx *Var_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVar_value(ctx *Var_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIso_level(ctx *Iso_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_boolean_or_string(ctx *Opt_boolean_or_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitZone_value(ctx *Zone_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_encoding(ctx *Opt_encodingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitNonreservedword_or_sconst(ctx *Nonreservedword_or_sconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVariableresetstmt(ctx *VariableresetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReset_rest(ctx *Reset_restContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGeneric_reset(ctx *Generic_resetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSetresetclause(ctx *SetresetclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunctionsetresetclause(ctx *FunctionsetresetclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVariableshowstmt(ctx *VariableshowstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstraintssetstmt(ctx *ConstraintssetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstraints_set_list(ctx *Constraints_set_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstraints_set_mode(ctx *Constraints_set_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCheckpointstmt(ctx *CheckpointstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDiscardstmt(ctx *DiscardstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltertablestmt(ctx *AltertablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_table_cmds(ctx *Alter_table_cmdsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPartition_cmd(ctx *Partition_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIndex_partition_cmd(ctx *Index_partition_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_table_cmd(ctx *Alter_table_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_column_default(ctx *Alter_column_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_drop_behavior(ctx *Opt_drop_behaviorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_collate_clause(ctx *Opt_collate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_using(ctx *Alter_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReplica_identity(ctx *Replica_identityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReloptions(ctx *ReloptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_reloptions(ctx *Opt_reloptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReloption_list(ctx *Reloption_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReloption_elem(ctx *Reloption_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_identity_column_option_list(ctx *Alter_identity_column_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_identity_column_option(ctx *Alter_identity_column_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPartitionboundspec(ctx *PartitionboundspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitHash_partbound_elem(ctx *Hash_partbound_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitHash_partbound(ctx *Hash_partboundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltercompositetypestmt(ctx *AltercompositetypestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_type_cmds(ctx *Alter_type_cmdsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_type_cmd(ctx *Alter_type_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCloseportalstmt(ctx *CloseportalstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopystmt(ctx *CopystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_from(ctx *Copy_fromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_program(ctx *Opt_programContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_file_name(ctx *Copy_file_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_options(ctx *Copy_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_opt_list(ctx *Copy_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_opt_item(ctx *Copy_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_binary(ctx *Opt_binaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_delimiter(ctx *Copy_delimiterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_using(ctx *Opt_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_generic_opt_list(ctx *Copy_generic_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_generic_opt_elem(ctx *Copy_generic_opt_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_generic_opt_arg(ctx *Copy_generic_opt_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_generic_opt_arg_list(ctx *Copy_generic_opt_arg_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCopy_generic_opt_arg_list_item(ctx *Copy_generic_opt_arg_list_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatestmt(ctx *CreatestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpttemp(ctx *OpttempContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpttableelementlist(ctx *OpttableelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpttypedtableelementlist(ctx *OpttypedtableelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTableelementlist(ctx *TableelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTypedtableelementlist(ctx *TypedtableelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTableelement(ctx *TableelementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTypedtableelement(ctx *TypedtableelementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumnDef(ctx *ColumnDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumnOptions(ctx *ColumnOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColquallist(ctx *ColquallistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColconstraint(ctx *ColconstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColconstraintelem(ctx *ColconstraintelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGenerated_when(ctx *Generated_whenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstraintattr(ctx *ConstraintattrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTablelikeclause(ctx *TablelikeclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTablelikeoptionlist(ctx *TablelikeoptionlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTablelikeoption(ctx *TablelikeoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTableconstraint(ctx *TableconstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstraintelem(ctx *ConstraintelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitHtap_action(ctx *Htap_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_no_inherit(ctx *Opt_no_inheritContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_column_list(ctx *Opt_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumnlist(ctx *ColumnlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumnElem(ctx *ColumnElemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumn_length(ctx *Column_lengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_c_include(ctx *Opt_c_includeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitKey_match(ctx *Key_matchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExclusionconstraintlist(ctx *ExclusionconstraintlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExclusionconstraintelem(ctx *ExclusionconstraintelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExclusionwhereclause(ctx *ExclusionwhereclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitKey_actions(ctx *Key_actionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitKey_update(ctx *Key_updateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitKey_delete(ctx *Key_deleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitKey_action(ctx *Key_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptinherit(ctx *OptinheritContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptpartitionspec(ctx *OptpartitionspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPartitionspec(ctx *PartitionspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPart_params(ctx *Part_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPart_elem(ctx *Part_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTable_access_method_clause(ctx *Table_access_method_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptwith(ctx *OptwithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOncommitoption(ctx *OncommitoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpttablespace(ctx *OpttablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptconstablespace(ctx *OptconstablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExistingindex(ctx *ExistingindexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatestatsstmt(ctx *CreatestatsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterstatsstmt(ctx *AlterstatsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateasstmt(ctx *CreateasstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreate_as_target(ctx *Create_as_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_with_data(ctx *Opt_with_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatematviewstmt(ctx *CreatematviewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreate_inc_mat_viewstmt(ctx *Create_inc_mat_viewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreate_mv_target(ctx *Create_mv_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptnolog(ctx *OptnologContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRefreshmatviewstmt(ctx *RefreshmatviewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateseqstmt(ctx *CreateseqstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterseqstmt(ctx *AlterseqstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptseqoptlist(ctx *OptseqoptlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptparenthesizedseqoptlist(ctx *OptparenthesizedseqoptlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSeqoptlist(ctx *SeqoptlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSeqoptelem(ctx *SeqoptelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_by(ctx *Opt_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitNumericonly(ctx *NumericonlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitNumericonly_list(ctx *Numericonly_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateplangstmt(ctx *CreateplangstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_trusted(ctx *Opt_trustedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitHandler_name(ctx *Handler_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_inline_handler(ctx *Opt_inline_handlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitValidator_clause(ctx *Validator_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_validator(ctx *Opt_validatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_procedural(ctx *Opt_proceduralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatetablespacestmt(ctx *CreatetablespacestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpttablespaceowner(ctx *OpttablespaceownerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDroptablespacestmt(ctx *DroptablespacestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateextensionstmt(ctx *CreateextensionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreate_extension_opt_list(ctx *Create_extension_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreate_extension_opt_item(ctx *Create_extension_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterextensionstmt(ctx *AlterextensionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_extension_opt_list(ctx *Alter_extension_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_extension_opt_item(ctx *Alter_extension_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterextensioncontentsstmt(ctx *AlterextensioncontentsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatefdwstmt(ctx *CreatefdwstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFdw_option(ctx *Fdw_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFdw_options(ctx *Fdw_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_fdw_options(ctx *Opt_fdw_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterfdwstmt(ctx *AlterfdwstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreate_generic_options(ctx *Create_generic_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGeneric_option_list(ctx *Generic_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_generic_options(ctx *Alter_generic_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_generic_option_list(ctx *Alter_generic_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_generic_option_elem(ctx *Alter_generic_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGeneric_option_elem(ctx *Generic_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGeneric_option_name(ctx *Generic_option_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGeneric_option_arg(ctx *Generic_option_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateforeignserverstmt(ctx *CreateforeignserverstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_type(ctx *Opt_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitForeign_server_version(ctx *Foreign_server_versionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_foreign_server_version(ctx *Opt_foreign_server_versionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterforeignserverstmt(ctx *AlterforeignserverstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateforeigntablestmt(ctx *CreateforeigntablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitImportforeignschemastmt(ctx *ImportforeignschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitImport_qualification_type(ctx *Import_qualification_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitImport_qualification(ctx *Import_qualificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateusermappingstmt(ctx *CreateusermappingstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAuth_ident(ctx *Auth_identContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDropusermappingstmt(ctx *DropusermappingstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterusermappingstmt(ctx *AlterusermappingstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatepolicystmt(ctx *CreatepolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterpolicystmt(ctx *AlterpolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRowsecurityoptionalexpr(ctx *RowsecurityoptionalexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRowsecurityoptionalwithcheck(ctx *RowsecurityoptionalwithcheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRowsecuritydefaulttorole(ctx *RowsecuritydefaulttoroleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRowsecurityoptionaltorole(ctx *RowsecurityoptionaltoroleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRowsecuritydefaultpermissive(ctx *RowsecuritydefaultpermissiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRowsecuritydefaultforcmd(ctx *RowsecuritydefaultforcmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRow_security_cmd(ctx *Row_security_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateamstmt(ctx *CreateamstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAm_type(ctx *Am_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatetrigstmt(ctx *CreatetrigstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggeractiontime(ctx *TriggeractiontimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggerevents(ctx *TriggereventsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggeroneevent(ctx *TriggeroneeventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggerreferencing(ctx *TriggerreferencingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggertransitions(ctx *TriggertransitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggertransition(ctx *TriggertransitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTransitionoldornew(ctx *TransitionoldornewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTransitionrowortable(ctx *TransitionrowortableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTransitionrelname(ctx *TransitionrelnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggerforspec(ctx *TriggerforspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggerforopteach(ctx *TriggerforopteachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggerfortype(ctx *TriggerfortypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggerwhen(ctx *TriggerwhenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunction_or_procedure(ctx *Function_or_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggerfuncargs(ctx *TriggerfuncargsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTriggerfuncarg(ctx *TriggerfuncargContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOptconstrfromtable(ctx *OptconstrfromtableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstraintattributespec(ctx *ConstraintattributespecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstraintattributeElem(ctx *ConstraintattributeElemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateeventtrigstmt(ctx *CreateeventtrigstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitEvent_trigger_when_list(ctx *Event_trigger_when_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitEvent_trigger_when_item(ctx *Event_trigger_when_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitEvent_trigger_value_list(ctx *Event_trigger_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltereventtrigstmt(ctx *AltereventtrigstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitEnable_trigger(ctx *Enable_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateassertionstmt(ctx *CreateassertionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDefinestmt(ctx *DefinestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDef_list(ctx *Def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDef_elem(ctx *Def_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDef_arg(ctx *Def_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOld_aggr_definition(ctx *Old_aggr_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOld_aggr_list(ctx *Old_aggr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOld_aggr_elem(ctx *Old_aggr_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_enum_val_list(ctx *Opt_enum_val_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitEnum_val_list(ctx *Enum_val_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterenumstmt(ctx *AlterenumstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_if_not_exists(ctx *Opt_if_not_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateopclassstmt(ctx *CreateopclassstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpclass_item_list(ctx *Opclass_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpclass_item(ctx *Opclass_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_default(ctx *Opt_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_opfamily(ctx *Opt_opfamilyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpclass_purpose(ctx *Opclass_purposeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_recheck(ctx *Opt_recheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateopfamilystmt(ctx *CreateopfamilystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlteropfamilystmt(ctx *AlteropfamilystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpclass_drop_list(ctx *Opclass_drop_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpclass_drop(ctx *Opclass_dropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDropopclassstmt(ctx *DropopclassstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDropopfamilystmt(ctx *DropopfamilystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDropownedstmt(ctx *DropownedstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReassignownedstmt(ctx *ReassignownedstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDropstmt(ctx *DropstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitObject_type_any_name(ctx *Object_type_any_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitObject_type_name(ctx *Object_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDrop_type_name(ctx *Drop_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitObject_type_name_on_any_name(ctx *Object_type_name_on_any_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAny_name_list(ctx *Any_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAny_name(ctx *Any_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAttrs(ctx *AttrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitType_name_list(ctx *Type_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTruncatestmt(ctx *TruncatestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_restart_seqs(ctx *Opt_restart_seqsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCommentstmt(ctx *CommentstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitComment_text(ctx *Comment_textContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSeclabelstmt(ctx *SeclabelstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_provider(ctx *Opt_providerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSecurity_label(ctx *Security_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCursorstmt(ctx *CursorstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFetchstmt(ctx *FetchstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFetch_args(ctx *Fetch_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFrom_in(ctx *From_inContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_from_in(ctx *Opt_from_inContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGrantstmt(ctx *GrantstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRevokestmt(ctx *RevokestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPrivileges(ctx *PrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPrivilege_list(ctx *Privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPrivilege(ctx *PrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPrivilege_target(ctx *Privilege_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGrantee_list(ctx *Grantee_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGrantee(ctx *GranteeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_grant_grant_option(ctx *Opt_grant_grant_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGrantrolestmt(ctx *GrantrolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRevokerolestmt(ctx *RevokerolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_grant_admin_option(ctx *Opt_grant_admin_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_granted_by(ctx *Opt_granted_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterdefaultprivilegesstmt(ctx *AlterdefaultprivilegesstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDefacloptionlist(ctx *DefacloptionlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDefacloption(ctx *DefacloptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDefaclaction(ctx *DefaclactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDefacl_privilege_target(ctx *Defacl_privilege_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIndexstmt(ctx *IndexstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_opt_distribute(ctx *Gaussdb_opt_distributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_unique(ctx *Opt_uniqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_concurrently(ctx *Opt_concurrentlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_index_name(ctx *Opt_index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAccess_method_clause(ctx *Access_method_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIndex_params(ctx *Index_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIndex_elem_options(ctx *Index_elem_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIndex_elem(ctx *Index_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_include(ctx *Opt_includeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIndex_including_params(ctx *Index_including_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_collate(ctx *Opt_collateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_class(ctx *Opt_classContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_asc_desc(ctx *Opt_asc_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_nulls_order(ctx *Opt_nulls_orderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatefunctionstmt(ctx *CreatefunctionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_or_replace(ctx *Opt_or_replaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_args(ctx *Func_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_args_list(ctx *Func_args_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunction_with_argtypes_list(ctx *Function_with_argtypes_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunction_with_argtypes(ctx *Function_with_argtypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_args_with_defaults(ctx *Func_args_with_defaultsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_args_with_defaults_list(ctx *Func_args_with_defaults_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_arg(ctx *Func_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitArg_class(ctx *Arg_classContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitParam_name(ctx *Param_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_return(ctx *Func_returnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_type(ctx *Func_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_arg_with_default(ctx *Func_arg_with_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAggr_arg(ctx *Aggr_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAggr_args(ctx *Aggr_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAggr_args_list(ctx *Aggr_args_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAggregate_with_argtypes(ctx *Aggregate_with_argtypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAggregate_with_argtypes_list(ctx *Aggregate_with_argtypes_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatefunc_opt_list(ctx *Createfunc_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCommon_func_opt_item(ctx *Common_func_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatefunc_opt_item(ctx *Createfunc_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_as(ctx *Func_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTransform_type_list(ctx *Transform_type_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_definition(ctx *Opt_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTable_func_column(ctx *Table_func_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTable_func_column_list(ctx *Table_func_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterfunctionstmt(ctx *AlterfunctionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterfunc_opt_list(ctx *Alterfunc_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_restrict(ctx *Opt_restrictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRemovefuncstmt(ctx *RemovefuncstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRemoveaggrstmt(ctx *RemoveaggrstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRemoveoperstmt(ctx *RemoveoperstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOper_argtypes(ctx *Oper_argtypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAny_operator(ctx *Any_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOperator_with_argtypes_list(ctx *Operator_with_argtypes_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOperator_with_argtypes(ctx *Operator_with_argtypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDostmt(ctx *DostmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDostmt_opt_list(ctx *Dostmt_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDostmt_opt_item(ctx *Dostmt_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatecaststmt(ctx *CreatecaststmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCast_context(ctx *Cast_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDropcaststmt(ctx *DropcaststmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_if_exists(ctx *Opt_if_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatetransformstmt(ctx *CreatetransformstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTransform_element_list(ctx *Transform_element_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDroptransformstmt(ctx *DroptransformstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReindexstmt(ctx *ReindexstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReindex_target_type(ctx *Reindex_target_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReindex_target_multitable(ctx *Reindex_target_multitableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReindex_option_list(ctx *Reindex_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReindex_option_elem(ctx *Reindex_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltertblspcstmt(ctx *AltertblspcstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRenamestmt(ctx *RenamestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_column(ctx *Opt_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_set_data(ctx *Opt_set_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterobjectdependsstmt(ctx *AlterobjectdependsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_no(ctx *Opt_noContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterobjectschemastmt(ctx *AlterobjectschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlteroperatorstmt(ctx *AlteroperatorstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOperator_def_list(ctx *Operator_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOperator_def_elem(ctx *Operator_def_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOperator_def_arg(ctx *Operator_def_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltertypestmt(ctx *AltertypestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterownerstmt(ctx *AlterownerstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatepublicationstmt(ctx *CreatepublicationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_publication_for_tables(ctx *Opt_publication_for_tablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPublication_for_tables(ctx *Publication_for_tablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterpublicationstmt(ctx *AlterpublicationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatesubscriptionstmt(ctx *CreatesubscriptionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPublication_name_list(ctx *Publication_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPublication_name_item(ctx *Publication_name_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltersubscriptionstmt(ctx *AltersubscriptionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDropsubscriptionstmt(ctx *DropsubscriptionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRulestmt(ctx *RulestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRuleactionlist(ctx *RuleactionlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRuleactionmulti(ctx *RuleactionmultiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRuleactionstmt(ctx *RuleactionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRuleactionstmtOrEmpty(ctx *RuleactionstmtOrEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitEvent(ctx *EventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_instead(ctx *Opt_insteadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitNotifystmt(ctx *NotifystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitNotify_payload(ctx *Notify_payloadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitListenstmt(ctx *ListenstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUnlistenstmt(ctx *UnlistenstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTransactionstmt(ctx *TransactionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_transaction(ctx *Opt_transactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTransaction_mode_item(ctx *Transaction_mode_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTransaction_mode_list(ctx *Transaction_mode_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTransaction_mode_list_or_empty(ctx *Transaction_mode_list_or_emptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_transaction_chain(ctx *Opt_transaction_chainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitViewstmt(ctx *ViewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_check_option(ctx *Opt_check_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLoadstmt(ctx *LoadstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatedbstmt(ctx *CreatedbstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatedb_opt_list(ctx *Createdb_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatedb_opt_items(ctx *Createdb_opt_itemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatedb_opt_item(ctx *Createdb_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatedb_opt_name(ctx *Createdb_opt_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_equal(ctx *Opt_equalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterdatabasestmt(ctx *AlterdatabasestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterdatabasesetstmt(ctx *AlterdatabasesetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDropdbstmt(ctx *DropdbstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDrop_option_list(ctx *Drop_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDrop_option(ctx *Drop_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltercollationstmt(ctx *AltercollationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltersystemstmt(ctx *AltersystemstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreatedomainstmt(ctx *CreatedomainstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlterdomainstmt(ctx *AlterdomainstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_as(ctx *Opt_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltertsdictionarystmt(ctx *AltertsdictionarystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAltertsconfigurationstmt(ctx *AltertsconfigurationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAny_with(ctx *Any_withContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreateconversionstmt(ctx *CreateconversionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitClusterstmt(ctx *ClusterstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCluster_index_specification(ctx *Cluster_index_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVacuumstmt(ctx *VacuumstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAnalyzestmt(ctx *AnalyzestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVac_analyze_option_list(ctx *Vac_analyze_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAnalyze_keyword(ctx *Analyze_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVac_analyze_option_elem(ctx *Vac_analyze_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVac_analyze_option_name(ctx *Vac_analyze_option_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVac_analyze_option_arg(ctx *Vac_analyze_option_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_analyze(ctx *Opt_analyzeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_verbose(ctx *Opt_verboseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_full(ctx *Opt_fullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_freeze(ctx *Opt_freezeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_name_list(ctx *Opt_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVacuum_relation(ctx *Vacuum_relationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVacuum_relation_list(ctx *Vacuum_relation_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_vacuum_relation_list(ctx *Opt_vacuum_relation_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExplainplanstmt(ctx *ExplainplanstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExplainstmt(ctx *ExplainstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExplainablestmt(ctx *ExplainablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExplain_option_list(ctx *Explain_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExplain_option_elem(ctx *Explain_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExplain_option_name(ctx *Explain_option_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExplain_option_arg(ctx *Explain_option_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPreparestmt(ctx *PreparestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPrep_type_clause(ctx *Prep_type_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPreparablestmt(ctx *PreparablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExecutestmt(ctx *ExecutestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExecute_param_clause(ctx *Execute_param_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDeallocatestmt(ctx *DeallocatestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInsertstmt(ctx *InsertstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInsert_target(ctx *Insert_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInsert_rest(ctx *Insert_restContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOverride_kind(ctx *Override_kindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInsert_column_list(ctx *Insert_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInsert_column_item(ctx *Insert_column_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_on_conflict(ctx *Opt_on_conflictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_conf_expr(ctx *Opt_conf_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReturning_clause(ctx *Returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMergestmt(ctx *MergestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMerge_insert_clause(ctx *Merge_insert_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMerge_update_clause(ctx *Merge_update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMerge_delete_clause(ctx *Merge_delete_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDeletestmt(ctx *DeletestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUsing_clause(ctx *Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLockstmt(ctx *LockstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLockbucketsstmt(ctx *LockbucketsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMarkbucketsstmt(ctx *MarkbucketsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_bucketlist(ctx *Opt_bucketlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_lock(ctx *Opt_lockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLock_type(ctx *Lock_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_nowait(ctx *Opt_nowaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_nowait_or_skip(ctx *Opt_nowait_or_skipContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUpdatestmt(ctx *UpdatestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSet_clause_list(ctx *Set_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSet_clause(ctx *Set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSet_target(ctx *Set_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSet_target_list(ctx *Set_target_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDeclarecursorstmt(ctx *DeclarecursorstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCursor_name(ctx *Cursor_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCursor_options(ctx *Cursor_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_hold(ctx *Opt_holdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSelectstmt(ctx *SelectstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSelect_with_parens(ctx *Select_with_parensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSelect_no_parens(ctx *Select_no_parensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSelect_clause(ctx *Select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSimple_select_intersect(ctx *Simple_select_intersectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSimple_select_pramary(ctx *Simple_select_pramaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWith_clause(ctx *With_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCte_list(ctx *Cte_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCommon_table_expr(ctx *Common_table_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_materialized(ctx *Opt_materializedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_with_clause(ctx *Opt_with_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInto_clause(ctx *Into_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_strict(ctx *Opt_strictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpttempTableName(ctx *OpttempTableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_table(ctx *Opt_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAll_or_distinct(ctx *All_or_distinctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDistinct_clause(ctx *Distinct_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_all_clause(ctx *Opt_all_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_sort_clause(ctx *Opt_sort_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSort_clause(ctx *Sort_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSortby_list(ctx *Sortby_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSortby(ctx *SortbyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSelect_limit(ctx *Select_limitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_select_limit(ctx *Opt_select_limitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLimit_clause(ctx *Limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOffset_clause(ctx *Offset_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSelect_limit_value(ctx *Select_limit_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSelect_offset_value(ctx *Select_offset_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSelect_fetch_first_value(ctx *Select_fetch_first_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitI_or_f_const(ctx *I_or_f_constContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRow_or_rows(ctx *Row_or_rowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFirst_or_next(ctx *First_or_nextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGroup_clause(ctx *Group_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGroup_by_list(ctx *Group_by_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGroup_by_item(ctx *Group_by_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitEmpty_grouping_set(ctx *Empty_grouping_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRollup_clause(ctx *Rollup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCube_clause(ctx *Cube_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitHaving_clause(ctx *Having_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFor_locking_clause(ctx *For_locking_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_for_locking_clause(ctx *Opt_for_locking_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFor_locking_items(ctx *For_locking_itemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFor_locking_item(ctx *For_locking_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFor_locking_strength(ctx *For_locking_strengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLocked_rels_list(ctx *Locked_rels_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitValues_clause(ctx *Values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFrom_list(ctx *From_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTable_ref(ctx *Table_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitJoined_table(ctx *Joined_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlias_clause(ctx *Alias_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_alias_clause(ctx *Opt_alias_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTable_alias_clause(ctx *Table_alias_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_alias_clause(ctx *Func_alias_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitJoin_type(ctx *Join_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitJoin_qual(ctx *Join_qualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRelation_expr(ctx *Relation_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRelation_expr_list(ctx *Relation_expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRelation_expr_opt_alias(ctx *Relation_expr_opt_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTablesample_clause(ctx *Tablesample_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_repeatable_clause(ctx *Opt_repeatable_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_table(ctx *Func_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRowsfrom_item(ctx *Rowsfrom_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRowsfrom_list(ctx *Rowsfrom_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_col_def_list(ctx *Opt_col_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_ordinality(ctx *Opt_ordinalityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWhere_clause(ctx *Where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWhere_or_current_clause(ctx *Where_or_current_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpttablefuncelementlist(ctx *OpttablefuncelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTablefuncelementlist(ctx *TablefuncelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTablefuncelement(ctx *TablefuncelementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXmltable(ctx *XmltableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXmltable_column_list(ctx *Xmltable_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXmltable_column_el(ctx *Xmltable_column_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXmltable_column_option_list(ctx *Xmltable_column_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXmltable_column_option_el(ctx *Xmltable_column_option_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXml_namespace_list(ctx *Xml_namespace_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXml_namespace_el(ctx *Xml_namespace_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTypename(ctx *TypenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_array_bounds(ctx *Opt_array_boundsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSimpletypename(ctx *SimpletypenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConsttypename(ctx *ConsttypenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGenerictype(ctx *GenerictypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_type_modifiers(ctx *Opt_type_modifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitNumeric(ctx *NumericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_float(ctx *Opt_floatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBit(ctx *BitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstbit(ctx *ConstbitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBitwithlength(ctx *BitwithlengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBitwithoutlength(ctx *BitwithoutlengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCharacter(ctx *CharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstcharacter(ctx *ConstcharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCharacter_c(ctx *Character_cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_varying(ctx *Opt_varyingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstdatetime(ctx *ConstdatetimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstinterval(ctx *ConstintervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_timezone(ctx *Opt_timezoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_interval(ctx *Opt_intervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInterval_second(ctx *Interval_secondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_escape(ctx *Opt_escapeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr(ctx *A_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_qual(ctx *A_expr_qualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_lessless(ctx *A_expr_lesslessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_or(ctx *A_expr_orContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_and(ctx *A_expr_andContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_between(ctx *A_expr_betweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_in(ctx *A_expr_inContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_unary_not(ctx *A_expr_unary_notContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_isnull(ctx *A_expr_isnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_is_not(ctx *A_expr_is_notContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_compare(ctx *A_expr_compareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_like(ctx *A_expr_likeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_qual_op(ctx *A_expr_qual_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_unary_qualop(ctx *A_expr_unary_qualopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_add(ctx *A_expr_addContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_mul(ctx *A_expr_mulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_caret(ctx *A_expr_caretContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_unary_sign(ctx *A_expr_unary_signContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_at_time_zone(ctx *A_expr_at_time_zoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_collate(ctx *A_expr_collateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitA_expr_typecast(ctx *A_expr_typecastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitB_expr(ctx *B_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitC_expr_exists(ctx *C_expr_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitC_expr_expr(ctx *C_expr_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitC_expr_case(ctx *C_expr_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPlsqlvariablename(ctx *PlsqlvariablenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_application(ctx *Func_applicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_expr(ctx *Func_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_expr_windowless(ctx *Func_expr_windowlessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_expr_common_subexpr(ctx *Func_expr_common_subexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXml_root_version(ctx *Xml_root_versionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_xml_root_standalone(ctx *Opt_xml_root_standaloneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXml_attributes(ctx *Xml_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXml_attribute_list(ctx *Xml_attribute_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXml_attribute_el(ctx *Xml_attribute_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDocument_or_content(ctx *Document_or_contentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXml_whitespace_option(ctx *Xml_whitespace_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXmlexists_argument(ctx *Xmlexists_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXml_passing_mech(ctx *Xml_passing_mechContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWithin_group_clause(ctx *Within_group_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFilter_clause(ctx *Filter_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWindow_clause(ctx *Window_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWindow_definition_list(ctx *Window_definition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWindow_definition(ctx *Window_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOver_clause(ctx *Over_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWindow_specification(ctx *Window_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_existing_window_name(ctx *Opt_existing_window_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_partition_clause(ctx *Opt_partition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_frame_clause(ctx *Opt_frame_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFrame_extent(ctx *Frame_extentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFrame_bound(ctx *Frame_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_window_exclusion_clause(ctx *Opt_window_exclusion_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRow(ctx *RowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExplicit_row(ctx *Explicit_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitImplicit_row(ctx *Implicit_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSub_type(ctx *Sub_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAll_op(ctx *All_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMathop(ctx *MathopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitQual_op(ctx *Qual_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitQual_all_op(ctx *Qual_all_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSubquery_Op(ctx *Subquery_OpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExpr_list(ctx *Expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_arg_list(ctx *Func_arg_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_arg_expr(ctx *Func_arg_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitType_list(ctx *Type_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitArray_expr(ctx *Array_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitArray_expr_list(ctx *Array_expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExtract_list(ctx *Extract_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExtract_arg(ctx *Extract_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUnicode_normal_form(ctx *Unicode_normal_formContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOverlay_list(ctx *Overlay_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPosition_list(ctx *Position_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSubstr_list(ctx *Substr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTrim_list(ctx *Trim_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIn_expr_select(ctx *In_expr_selectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIn_expr_list(ctx *In_expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCase_expr(ctx *Case_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWhen_clause_list(ctx *When_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitWhen_clause(ctx *When_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCase_default(ctx *Case_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCase_arg(ctx *Case_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumnref(ctx *ColumnrefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIndirection_el(ctx *Indirection_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_slice_bound(ctx *Opt_slice_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIndirection(ctx *IndirectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_indirection(ctx *Opt_indirectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_target_list(ctx *Opt_target_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTarget_list(ctx *Target_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTarget_label(ctx *Target_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTarget_star(ctx *Target_starContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTarget_alias(ctx *Target_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitQualified_name_list(ctx *Qualified_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitQualified_name(ctx *Qualified_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLink_name(ctx *Link_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAt_link_name(ctx *At_link_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitName_list(ctx *Name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAttr_name(ctx *Attr_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFile_name(ctx *File_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunc_name(ctx *Func_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAexprconst(ctx *AexprconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitXconst(ctx *XconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBconst(ctx *BconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFconst(ctx *FconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIconst(ctx *IconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSconst(ctx *SconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAnysconst(ctx *AnysconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_uescape(ctx *Opt_uescapeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSignediconst(ctx *SignediconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRoleid(ctx *RoleidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRolespec(ctx *RolespecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRole_list(ctx *Role_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColid(ctx *ColidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTable_alias(ctx *Table_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitType_function_name(ctx *Type_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitNonreservedword(ctx *NonreservedwordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCollabel(ctx *CollabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPlsqlidentifier(ctx *PlsqlidentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUnreserved_keyword(ctx *Unreserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCol_name_keyword(ctx *Col_name_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitType_func_name_keyword(ctx *Type_func_name_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReserved_keyword(ctx *Reserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBuiltin_function_name(ctx *Builtin_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPl_function(ctx *Pl_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitComp_options(ctx *Comp_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitComp_option(ctx *Comp_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSharp(ctx *SharpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOption_value(ctx *Option_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_semi(ctx *Opt_semiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPl_block(ctx *Pl_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_sect(ctx *Decl_sectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_start(ctx *Decl_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_stmts(ctx *Decl_stmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLabel_decl(ctx *Label_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_stmt(ctx *Decl_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_statement(ctx *Decl_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_scrollable(ctx *Opt_scrollableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_cursor_query(ctx *Decl_cursor_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_cursor_args(ctx *Decl_cursor_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_cursor_arglist(ctx *Decl_cursor_arglistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_cursor_arg(ctx *Decl_cursor_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_is_for(ctx *Decl_is_forContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_aliasitem(ctx *Decl_aliasitemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_varname(ctx *Decl_varnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_const(ctx *Decl_constContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_datatype(ctx *Decl_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_collate(ctx *Decl_collateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_notnull(ctx *Decl_notnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_defval(ctx *Decl_defvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDecl_defkey(ctx *Decl_defkeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAssign_operator(ctx *Assign_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitProc_sect(ctx *Proc_sectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitProc_stmt(ctx *Proc_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_perform(ctx *Stmt_performContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_call(ctx *Stmt_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_expr_list(ctx *Opt_expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_assign(ctx *Stmt_assignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_getdiag(ctx *Stmt_getdiagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGetdiag_area_opt(ctx *Getdiag_area_optContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGetdiag_list(ctx *Getdiag_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGetdiag_list_item(ctx *Getdiag_list_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGetdiag_item(ctx *Getdiag_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGetdiag_target(ctx *Getdiag_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAssign_var(ctx *Assign_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_if(ctx *Stmt_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_elsifs(ctx *Stmt_elsifsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_else(ctx *Stmt_elseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_case(ctx *Stmt_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_expr_until_when(ctx *Opt_expr_until_whenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCase_when_list(ctx *Case_when_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCase_when(ctx *Case_whenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_case_else(ctx *Opt_case_elseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_loop(ctx *Stmt_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_while(ctx *Stmt_whileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_for(ctx *Stmt_forContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFor_control(ctx *For_controlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_for_using_expression(ctx *Opt_for_using_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_cursor_parameters(ctx *Opt_cursor_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_reverse(ctx *Opt_reverseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_by_expression(ctx *Opt_by_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFor_variable(ctx *For_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_foreach_a(ctx *Stmt_foreach_aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitForeach_slice(ctx *Foreach_sliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_exit(ctx *Stmt_exitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExit_type(ctx *Exit_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_return(ctx *Stmt_returnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_return_result(ctx *Opt_return_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_raise(ctx *Stmt_raiseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_stmt_raise_level(ctx *Opt_stmt_raise_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_raise_list(ctx *Opt_raise_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_raise_using(ctx *Opt_raise_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_raise_using_elem(ctx *Opt_raise_using_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_raise_using_elem_list(ctx *Opt_raise_using_elem_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_assert(ctx *Stmt_assertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_stmt_assert_message(ctx *Opt_stmt_assert_messageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLoop_body(ctx *Loop_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_execsql(ctx *Stmt_execsqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_dynexecute(ctx *Stmt_dynexecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_execute_using(ctx *Opt_execute_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_execute_using_list(ctx *Opt_execute_using_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_execute_into(ctx *Opt_execute_intoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_open(ctx *Stmt_openContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_open_bound_list_item(ctx *Opt_open_bound_list_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_open_bound_list(ctx *Opt_open_bound_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_open_using(ctx *Opt_open_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_scroll_option(ctx *Opt_scroll_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_scroll_option_no(ctx *Opt_scroll_option_noContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_fetch(ctx *Stmt_fetchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInto_target(ctx *Into_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_cursor_from(ctx *Opt_cursor_fromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_fetch_direction(ctx *Opt_fetch_directionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_move(ctx *Stmt_moveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_close(ctx *Stmt_closeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_null(ctx *Stmt_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_commit(ctx *Stmt_commitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_rollback(ctx *Stmt_rollbackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPlsql_opt_transaction_chain(ctx *Plsql_opt_transaction_chainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStmt_set(ctx *Stmt_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCursor_variable(ctx *Cursor_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitException_sect(ctx *Exception_sectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitProc_exceptions(ctx *Proc_exceptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitProc_exception(ctx *Proc_exceptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitProc_conditions(ctx *Proc_conditionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitProc_condition(ctx *Proc_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_block_label(ctx *Opt_block_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_loop_label(ctx *Opt_loop_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_label(ctx *Opt_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_exitcond(ctx *Opt_exitcondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAny_identifier(ctx *Any_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPlsql_unreserved_keyword(ctx *Plsql_unreserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSql_expression(ctx *Sql_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExpr_until_then(ctx *Expr_until_thenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExpr_until_semi(ctx *Expr_until_semiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExpr_until_rightbracket(ctx *Expr_until_rightbracketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExpr_until_loop(ctx *Expr_until_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMake_execsql_stmt(ctx *Make_execsql_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpt_returning_clause_into(ctx *Opt_returning_clause_intoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_generic_position(ctx *Gaussdb_create_generic_positionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_foreigntable_extend(ctx *Gaussdb_create_foreigntable_extendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_error_table_name(ctx *Gaussdb_error_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_nodenamelist(ctx *Gaussdb_nodenamelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_groupname(ctx *Gaussdb_groupnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_to_group(ctx *Gaussdb_to_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_limit_clause(ctx *Gaussdb_limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_sort_clause(ctx *Gaussdb_sort_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_from_clause(ctx *Gaussdb_from_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_insert_when_clause(ctx *Gaussdb_insert_when_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_truncate_table_partition(ctx *Gaussdb_truncate_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_extended_names(ctx *Gaussdb_partition_extended_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_name(ctx *Gaussdb_partition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_new_partition_name(ctx *Gaussdb_new_partition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_key_value(ctx *Gaussdb_partition_key_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_PURGE(ctx *Gaussdb_PURGEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_or_replace(ctx *Gaussdb_or_replaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_index_name(ctx *Gaussdb_index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_comment_list(ctx *Gaussdb_comment_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_comment(ctx *Gaussdb_commentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_alter_index_ops_set(ctx *Gaussdb_alter_index_ops_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_rebuild_clause(ctx *Gaussdb_rebuild_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_alter_index_partitioning(ctx *Gaussdb_alter_index_partitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_modify_index_partition(ctx *Gaussdb_modify_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_rename_index_partition(ctx *Gaussdb_rename_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_move_index_partition(ctx *Gaussdb_move_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_online(ctx *Gaussdb_onlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_index_partition_name(ctx *Gaussdb_index_partition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_value(ctx *Gaussdb_partition_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_index_partition_tablespace(ctx *Gaussdb_index_partition_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_index_partition_list(ctx *Gaussdb_index_partition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_index_partition(ctx *Gaussdb_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_value_list(ctx *Gaussdb_partition_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_index_subpartition_list(ctx *Gaussdb_index_subpartition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_index_subppartition(ctx *Gaussdb_index_subppartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_connect_clause(ctx *Gaussdb_connect_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_engine_list(ctx *Gaussdb_engine_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_engine_item(ctx *Gaussdb_engine_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_table_option_list(ctx *Gaussdb_table_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_table_option(ctx *Gaussdb_table_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_htap_option(ctx *Gaussdb_htap_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_distribute(ctx *Gaussdb_distributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_slice_less_than_item_list(ctx *Gaussdb_slice_less_than_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_slice_less_than_item(ctx *Gaussdb_slice_less_than_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_slice_start_end_item_list(ctx *Gaussdb_slice_start_end_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_slice_start_end_item(ctx *Gaussdb_slice_start_end_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_slice_values_item_list(ctx *Gaussdb_slice_values_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_slice_values_item_item(ctx *Gaussdb_slice_values_item_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_ilm_add_policy(ctx *Gaussdb_ilm_add_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLabel_name(ctx *Label_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitException_name(ctx *Exception_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitImplementation_type_name(ctx *Implementation_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitParameter_name(ctx *Parameter_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitNumeric_negative(ctx *Numeric_negativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPlsql_body(ctx *Plsql_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitException_handler(ctx *Exception_handlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSeq_of_statements(ctx *Seq_of_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLabel_declaration(ctx *Label_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSeq_of_declare_specs(ctx *Seq_of_declare_specsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDeclare_spec(ctx *Declare_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPragma_declaration(ctx *Pragma_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitException_declaration(ctx *Exception_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDefault_value_part(ctx *Default_value_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVariable_declaration(ctx *Variable_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSubtype_declaration(ctx *Subtype_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCursor_declaration(ctx *Cursor_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitParameter_spec(ctx *Parameter_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRef_cursor_type_def(ctx *Ref_cursor_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitType_declaration(ctx *Type_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTable_type_def(ctx *Table_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTable_indexed_by_part(ctx *Table_indexed_by_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVarray_type_def(ctx *Varray_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRecord_type_def(ctx *Record_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitField_spec(ctx *Field_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitProcedure_spec(ctx *Procedure_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAccessible_by_clause(ctx *Accessible_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAccessor(ctx *AccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunction_spec(ctx *Function_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitParallel_enable_clause(ctx *Parallel_enable_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPartition_by_clause(ctx *Partition_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStreaming_clause(ctx *Streaming_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumn_list(ctx *Column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitParen_column_list(ctx *Paren_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitProcedure_body(ctx *Procedure_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFunction_body(ctx *Function_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitResult_cache_clause(ctx *Result_cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRelies_on_part(ctx *Relies_on_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPlsql_statement(ctx *Plsql_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSql_statement(ctx *Sql_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCursor_manipulation_statements(ctx *Cursor_manipulation_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitClose_statement(ctx *Close_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpen_statement(ctx *Open_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFetch_statement(ctx *Fetch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitVariable_or_collection(ctx *Variable_or_collectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCollection_expression(ctx *Collection_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpen_for_statement(ctx *Open_for_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCollection_method_call(ctx *Collection_method_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExecute_immediate(ctx *Execute_immediateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDynamic_returning_clause(ctx *Dynamic_returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRaise_statement(ctx *Raise_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPipe_row_statement(ctx *Pipe_row_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitForall_statement(ctx *Forall_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBounds_clause(ctx *Bounds_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBetween_bound(ctx *Between_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitData_manipulation_language_statements(ctx *Data_manipulation_language_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitNull_statement(ctx *Null_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGoto_statement(ctx *Goto_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitElsif_part(ctx *Elsif_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitElse_part(ctx *Else_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLoop_statement(ctx *Loop_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCursor_loop_param(ctx *Cursor_loop_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLower_bound(ctx *Lower_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUpper_bound(ctx *Upper_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExit_statement(ctx *Exit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAssignment_statement(ctx *Assignment_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGeneral_element(ctx *General_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitContinue_statement(ctx *Continue_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCursor_expression(ctx *Cursor_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLogical_expression(ctx *Logical_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUnary_logical_expression(ctx *Unary_logical_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUnary_logical_operation(ctx *Unary_logical_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLogical_operation(ctx *Logical_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMultiset_expression(ctx *Multiset_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConcatenation(ctx *ConcatenationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRelational_expression(ctx *Relational_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCompound_expression(ctx *Compound_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitRelational_operator(ctx *Relational_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitIn_elements(ctx *In_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitBetween_elements(ctx *Between_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitInterval_expression(ctx *Interval_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitModel_expression(ctx *Model_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitModel_expression_element(ctx *Model_expression_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSingle_column_for_loop(ctx *Single_column_for_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMulti_column_for_loop(ctx *Multi_column_for_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUnary_expression(ctx *Unary_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitImplicit_cursor_expression(ctx *Implicit_cursor_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCase_statement(ctx *Case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSimple_case_statement(ctx *Simple_case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSimple_case_when_part(ctx *Simple_case_when_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSearched_case_statement(ctx *Searched_case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSearched_case_when_part(ctx *Searched_case_when_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCase_else_part(ctx *Case_else_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOuter_join_sign(ctx *Outer_join_signContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitQuantified_expression(ctx *Quantified_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitStandard_function(ctx *Standard_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitTrigger_anonymous_block(ctx *Trigger_anonymous_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAnonymous_block(ctx *Anonymous_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPackage_name(ctx *Package_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitAlter_package(ctx *Alter_packageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreate_package(ctx *Create_packageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitCreate_package_body(ctx *Create_package_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPackage_obj_spec(ctx *Package_obj_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPackage_obj_body(ctx *Package_obj_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_cleanconnection(ctx *Gaussdb_cleanconnectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPolicy_name(ctx *Policy_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_createpolicystmt(ctx *Gaussdb_createpolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_audit_clause_list(ctx *Gaussdb_audit_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPrivilege_access_audit_clause(ctx *Privilege_access_audit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDdl_clause(ctx *Ddl_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitResource_label_name(ctx *Resource_label_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDml_clause(ctx *Dml_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFilter_group_clause(ctx *Filter_group_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFilter_type_list(ctx *Filter_type_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFilter_type(ctx *Filter_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFilter_value_list(ctx *Filter_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitFilter_value(ctx *Filter_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_createbarrierstmt(ctx *Gaussdb_createbarrierstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_client_master_key(ctx *Gaussdb_create_client_master_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitClient_master_key_name(ctx *Client_master_key_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitClient_master_key_list(ctx *Client_master_key_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitClient_master_key(ctx *Client_master_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_column_encryption_key(ctx *Gaussdb_create_column_encryption_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumn_encryption_key_name(ctx *Column_encryption_key_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumn_encryption_key_list(ctx *Column_encryption_key_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitColumn_encryption_key(ctx *Column_encryption_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_database_link(ctx *Gaussdb_create_database_linkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitUser_object_name(ctx *User_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPassword_value(ctx *Password_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDatabase_link_using(ctx *Database_link_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDatabase_link_using_opt(ctx *Database_link_using_optContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_alter_database_link(ctx *Gaussdb_alter_database_linkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_directory(ctx *Gaussdb_create_directoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitDirectory_name(ctx *Directory_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_masking_policy(ctx *Gaussdb_create_masking_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMasking_clause_list(ctx *Masking_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMasking_clause(ctx *Masking_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitMasking_function(ctx *Masking_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitPolicy_filter_clause(ctx *Policy_filter_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_node(ctx *Gaussdb_create_nodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_nodename(ctx *Gaussdb_nodenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_node_group(ctx *Gaussdb_create_node_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_bucketnumber(ctx *Gaussdb_bucketnumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_resource_pool(ctx *Gaussdb_create_resource_poolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_resource_pool_name(ctx *Gaussdb_resource_pool_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_resource_label(ctx *Gaussdb_create_resource_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLabel_item_list(ctx *Label_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitLabel_item(ctx *Label_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitResource_path(ctx *Resource_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitResource_type(ctx *Resource_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_security_label(ctx *Gaussdb_create_security_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_synonym(ctx *Gaussdb_create_synonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitSynonym_name(ctx *Synonym_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_less_than_item_list(ctx *Gaussdb_partition_less_than_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_less_than_item(ctx *Gaussdb_partition_less_than_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_start_end_item_list(ctx *Gaussdb_partition_start_end_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_start_end_item(ctx *Gaussdb_partition_start_end_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitOpttablespace_list(ctx *Opttablespace_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_hash_item(ctx *Gaussdb_partition_hash_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_partition_list_item(ctx *Gaussdb_partition_list_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_create_weak_password_dic(ctx *Gaussdb_create_weak_password_dicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_purge_stmt(ctx *Gaussdb_purge_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_replacestmt(ctx *Gaussdb_replacestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_shutdown_stmt(ctx *Gaussdb_shutdown_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGaussdbParserVisitor) VisitGaussdb_show_events_stmt(ctx *Gaussdb_show_events_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}
