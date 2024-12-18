// Code generated from GaussdbParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // GaussdbParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GaussdbParser.
type GaussdbParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GaussdbParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by GaussdbParser#plsqlroot.
	VisitPlsqlroot(ctx *PlsqlrootContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmtblock.
	VisitStmtblock(ctx *StmtblockContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmtmulti.
	VisitStmtmulti(ctx *StmtmultiContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#plsqlconsolecommand.
	VisitPlsqlconsolecommand(ctx *PlsqlconsolecommandContext) interface{}

	// Visit a parse tree produced by GaussdbParser#callstmt.
	VisitCallstmt(ctx *CallstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createrolestmt.
	VisitCreaterolestmt(ctx *CreaterolestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_with.
	VisitOpt_with(ctx *Opt_withContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optrolelist.
	VisitOptrolelist(ctx *OptrolelistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alteroptrolelist.
	VisitAlteroptrolelist(ctx *AlteroptrolelistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alteroptroleelem.
	VisitAlteroptroleelem(ctx *AlteroptroleelemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createoptroleelem.
	VisitCreateoptroleelem(ctx *CreateoptroleelemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createuserstmt.
	VisitCreateuserstmt(ctx *CreateuserstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterrolestmt.
	VisitAlterrolestmt(ctx *AlterrolestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_in_database.
	VisitOpt_in_database(ctx *Opt_in_databaseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterrolesetstmt.
	VisitAlterrolesetstmt(ctx *AlterrolesetstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#droprolestmt.
	VisitDroprolestmt(ctx *DroprolestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#creategroupstmt.
	VisitCreategroupstmt(ctx *CreategroupstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altergroupstmt.
	VisitAltergroupstmt(ctx *AltergroupstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#add_drop.
	VisitAdd_drop(ctx *Add_dropContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createschemastmt.
	VisitCreateschemastmt(ctx *CreateschemastmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optschemaname.
	VisitOptschemaname(ctx *OptschemanameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optschemaeltlist.
	VisitOptschemaeltlist(ctx *OptschemaeltlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#schema_stmt.
	VisitSchema_stmt(ctx *Schema_stmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#variablesetstmt.
	VisitVariablesetstmt(ctx *VariablesetstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#set_rest.
	VisitSet_rest(ctx *Set_restContext) interface{}

	// Visit a parse tree produced by GaussdbParser#generic_set.
	VisitGeneric_set(ctx *Generic_setContext) interface{}

	// Visit a parse tree produced by GaussdbParser#set_rest_more.
	VisitSet_rest_more(ctx *Set_rest_moreContext) interface{}

	// Visit a parse tree produced by GaussdbParser#var_name.
	VisitVar_name(ctx *Var_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#var_list.
	VisitVar_list(ctx *Var_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#var_value.
	VisitVar_value(ctx *Var_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#iso_level.
	VisitIso_level(ctx *Iso_levelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_boolean_or_string.
	VisitOpt_boolean_or_string(ctx *Opt_boolean_or_stringContext) interface{}

	// Visit a parse tree produced by GaussdbParser#zone_value.
	VisitZone_value(ctx *Zone_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_encoding.
	VisitOpt_encoding(ctx *Opt_encodingContext) interface{}

	// Visit a parse tree produced by GaussdbParser#nonreservedword_or_sconst.
	VisitNonreservedword_or_sconst(ctx *Nonreservedword_or_sconstContext) interface{}

	// Visit a parse tree produced by GaussdbParser#variableresetstmt.
	VisitVariableresetstmt(ctx *VariableresetstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reset_rest.
	VisitReset_rest(ctx *Reset_restContext) interface{}

	// Visit a parse tree produced by GaussdbParser#generic_reset.
	VisitGeneric_reset(ctx *Generic_resetContext) interface{}

	// Visit a parse tree produced by GaussdbParser#setresetclause.
	VisitSetresetclause(ctx *SetresetclauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#functionsetresetclause.
	VisitFunctionsetresetclause(ctx *FunctionsetresetclauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#variableshowstmt.
	VisitVariableshowstmt(ctx *VariableshowstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constraintssetstmt.
	VisitConstraintssetstmt(ctx *ConstraintssetstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constraints_set_list.
	VisitConstraints_set_list(ctx *Constraints_set_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constraints_set_mode.
	VisitConstraints_set_mode(ctx *Constraints_set_modeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#checkpointstmt.
	VisitCheckpointstmt(ctx *CheckpointstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#discardstmt.
	VisitDiscardstmt(ctx *DiscardstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altertablestmt.
	VisitAltertablestmt(ctx *AltertablestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_table_cmds.
	VisitAlter_table_cmds(ctx *Alter_table_cmdsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#partition_cmd.
	VisitPartition_cmd(ctx *Partition_cmdContext) interface{}

	// Visit a parse tree produced by GaussdbParser#index_partition_cmd.
	VisitIndex_partition_cmd(ctx *Index_partition_cmdContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_table_cmd.
	VisitAlter_table_cmd(ctx *Alter_table_cmdContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_column_default.
	VisitAlter_column_default(ctx *Alter_column_defaultContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_drop_behavior.
	VisitOpt_drop_behavior(ctx *Opt_drop_behaviorContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_collate_clause.
	VisitOpt_collate_clause(ctx *Opt_collate_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_using.
	VisitAlter_using(ctx *Alter_usingContext) interface{}

	// Visit a parse tree produced by GaussdbParser#replica_identity.
	VisitReplica_identity(ctx *Replica_identityContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reloptions.
	VisitReloptions(ctx *ReloptionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_reloptions.
	VisitOpt_reloptions(ctx *Opt_reloptionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reloption_list.
	VisitReloption_list(ctx *Reloption_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reloption_elem.
	VisitReloption_elem(ctx *Reloption_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_identity_column_option_list.
	VisitAlter_identity_column_option_list(ctx *Alter_identity_column_option_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_identity_column_option.
	VisitAlter_identity_column_option(ctx *Alter_identity_column_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#partitionboundspec.
	VisitPartitionboundspec(ctx *PartitionboundspecContext) interface{}

	// Visit a parse tree produced by GaussdbParser#hash_partbound_elem.
	VisitHash_partbound_elem(ctx *Hash_partbound_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#hash_partbound.
	VisitHash_partbound(ctx *Hash_partboundContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altercompositetypestmt.
	VisitAltercompositetypestmt(ctx *AltercompositetypestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_type_cmds.
	VisitAlter_type_cmds(ctx *Alter_type_cmdsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_type_cmd.
	VisitAlter_type_cmd(ctx *Alter_type_cmdContext) interface{}

	// Visit a parse tree produced by GaussdbParser#closeportalstmt.
	VisitCloseportalstmt(ctx *CloseportalstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copystmt.
	VisitCopystmt(ctx *CopystmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_from.
	VisitCopy_from(ctx *Copy_fromContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_program.
	VisitOpt_program(ctx *Opt_programContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_file_name.
	VisitCopy_file_name(ctx *Copy_file_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_options.
	VisitCopy_options(ctx *Copy_optionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_opt_list.
	VisitCopy_opt_list(ctx *Copy_opt_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_opt_item.
	VisitCopy_opt_item(ctx *Copy_opt_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_binary.
	VisitOpt_binary(ctx *Opt_binaryContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_delimiter.
	VisitCopy_delimiter(ctx *Copy_delimiterContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_using.
	VisitOpt_using(ctx *Opt_usingContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_generic_opt_list.
	VisitCopy_generic_opt_list(ctx *Copy_generic_opt_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_generic_opt_elem.
	VisitCopy_generic_opt_elem(ctx *Copy_generic_opt_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_generic_opt_arg.
	VisitCopy_generic_opt_arg(ctx *Copy_generic_opt_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_generic_opt_arg_list.
	VisitCopy_generic_opt_arg_list(ctx *Copy_generic_opt_arg_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#copy_generic_opt_arg_list_item.
	VisitCopy_generic_opt_arg_list_item(ctx *Copy_generic_opt_arg_list_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createstmt.
	VisitCreatestmt(ctx *CreatestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opttemp.
	VisitOpttemp(ctx *OpttempContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opttableelementlist.
	VisitOpttableelementlist(ctx *OpttableelementlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opttypedtableelementlist.
	VisitOpttypedtableelementlist(ctx *OpttypedtableelementlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#tableelementlist.
	VisitTableelementlist(ctx *TableelementlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#typedtableelementlist.
	VisitTypedtableelementlist(ctx *TypedtableelementlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#tableelement.
	VisitTableelement(ctx *TableelementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#typedtableelement.
	VisitTypedtableelement(ctx *TypedtableelementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#columnDef.
	VisitColumnDef(ctx *ColumnDefContext) interface{}

	// Visit a parse tree produced by GaussdbParser#columnOptions.
	VisitColumnOptions(ctx *ColumnOptionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#colquallist.
	VisitColquallist(ctx *ColquallistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#colconstraint.
	VisitColconstraint(ctx *ColconstraintContext) interface{}

	// Visit a parse tree produced by GaussdbParser#colconstraintelem.
	VisitColconstraintelem(ctx *ColconstraintelemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#generated_when.
	VisitGenerated_when(ctx *Generated_whenContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constraintattr.
	VisitConstraintattr(ctx *ConstraintattrContext) interface{}

	// Visit a parse tree produced by GaussdbParser#tablelikeclause.
	VisitTablelikeclause(ctx *TablelikeclauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#tablelikeoptionlist.
	VisitTablelikeoptionlist(ctx *TablelikeoptionlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#tablelikeoption.
	VisitTablelikeoption(ctx *TablelikeoptionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#tableconstraint.
	VisitTableconstraint(ctx *TableconstraintContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constraintelem.
	VisitConstraintelem(ctx *ConstraintelemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#htap_action.
	VisitHtap_action(ctx *Htap_actionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_no_inherit.
	VisitOpt_no_inherit(ctx *Opt_no_inheritContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_column_list.
	VisitOpt_column_list(ctx *Opt_column_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#columnlist.
	VisitColumnlist(ctx *ColumnlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#columnElem.
	VisitColumnElem(ctx *ColumnElemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#column_length.
	VisitColumn_length(ctx *Column_lengthContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_c_include.
	VisitOpt_c_include(ctx *Opt_c_includeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#key_match.
	VisitKey_match(ctx *Key_matchContext) interface{}

	// Visit a parse tree produced by GaussdbParser#exclusionconstraintlist.
	VisitExclusionconstraintlist(ctx *ExclusionconstraintlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#exclusionconstraintelem.
	VisitExclusionconstraintelem(ctx *ExclusionconstraintelemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#exclusionwhereclause.
	VisitExclusionwhereclause(ctx *ExclusionwhereclauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#key_actions.
	VisitKey_actions(ctx *Key_actionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#key_update.
	VisitKey_update(ctx *Key_updateContext) interface{}

	// Visit a parse tree produced by GaussdbParser#key_delete.
	VisitKey_delete(ctx *Key_deleteContext) interface{}

	// Visit a parse tree produced by GaussdbParser#key_action.
	VisitKey_action(ctx *Key_actionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optinherit.
	VisitOptinherit(ctx *OptinheritContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optpartitionspec.
	VisitOptpartitionspec(ctx *OptpartitionspecContext) interface{}

	// Visit a parse tree produced by GaussdbParser#partitionspec.
	VisitPartitionspec(ctx *PartitionspecContext) interface{}

	// Visit a parse tree produced by GaussdbParser#part_params.
	VisitPart_params(ctx *Part_paramsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#part_elem.
	VisitPart_elem(ctx *Part_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#table_access_method_clause.
	VisitTable_access_method_clause(ctx *Table_access_method_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optwith.
	VisitOptwith(ctx *OptwithContext) interface{}

	// Visit a parse tree produced by GaussdbParser#oncommitoption.
	VisitOncommitoption(ctx *OncommitoptionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opttablespace.
	VisitOpttablespace(ctx *OpttablespaceContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optconstablespace.
	VisitOptconstablespace(ctx *OptconstablespaceContext) interface{}

	// Visit a parse tree produced by GaussdbParser#existingindex.
	VisitExistingindex(ctx *ExistingindexContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createstatsstmt.
	VisitCreatestatsstmt(ctx *CreatestatsstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterstatsstmt.
	VisitAlterstatsstmt(ctx *AlterstatsstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createasstmt.
	VisitCreateasstmt(ctx *CreateasstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#create_as_target.
	VisitCreate_as_target(ctx *Create_as_targetContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_with_data.
	VisitOpt_with_data(ctx *Opt_with_dataContext) interface{}

	// Visit a parse tree produced by GaussdbParser#creatematviewstmt.
	VisitCreatematviewstmt(ctx *CreatematviewstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#create_inc_mat_viewstmt.
	VisitCreate_inc_mat_viewstmt(ctx *Create_inc_mat_viewstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#create_mv_target.
	VisitCreate_mv_target(ctx *Create_mv_targetContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optnolog.
	VisitOptnolog(ctx *OptnologContext) interface{}

	// Visit a parse tree produced by GaussdbParser#refreshmatviewstmt.
	VisitRefreshmatviewstmt(ctx *RefreshmatviewstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createseqstmt.
	VisitCreateseqstmt(ctx *CreateseqstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterseqstmt.
	VisitAlterseqstmt(ctx *AlterseqstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optseqoptlist.
	VisitOptseqoptlist(ctx *OptseqoptlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optparenthesizedseqoptlist.
	VisitOptparenthesizedseqoptlist(ctx *OptparenthesizedseqoptlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#seqoptlist.
	VisitSeqoptlist(ctx *SeqoptlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#seqoptelem.
	VisitSeqoptelem(ctx *SeqoptelemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_by.
	VisitOpt_by(ctx *Opt_byContext) interface{}

	// Visit a parse tree produced by GaussdbParser#numericonly.
	VisitNumericonly(ctx *NumericonlyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#numericonly_list.
	VisitNumericonly_list(ctx *Numericonly_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createplangstmt.
	VisitCreateplangstmt(ctx *CreateplangstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_trusted.
	VisitOpt_trusted(ctx *Opt_trustedContext) interface{}

	// Visit a parse tree produced by GaussdbParser#handler_name.
	VisitHandler_name(ctx *Handler_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_inline_handler.
	VisitOpt_inline_handler(ctx *Opt_inline_handlerContext) interface{}

	// Visit a parse tree produced by GaussdbParser#validator_clause.
	VisitValidator_clause(ctx *Validator_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_validator.
	VisitOpt_validator(ctx *Opt_validatorContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_procedural.
	VisitOpt_procedural(ctx *Opt_proceduralContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createtablespacestmt.
	VisitCreatetablespacestmt(ctx *CreatetablespacestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opttablespaceowner.
	VisitOpttablespaceowner(ctx *OpttablespaceownerContext) interface{}

	// Visit a parse tree produced by GaussdbParser#droptablespacestmt.
	VisitDroptablespacestmt(ctx *DroptablespacestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createextensionstmt.
	VisitCreateextensionstmt(ctx *CreateextensionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#create_extension_opt_list.
	VisitCreate_extension_opt_list(ctx *Create_extension_opt_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#create_extension_opt_item.
	VisitCreate_extension_opt_item(ctx *Create_extension_opt_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterextensionstmt.
	VisitAlterextensionstmt(ctx *AlterextensionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_extension_opt_list.
	VisitAlter_extension_opt_list(ctx *Alter_extension_opt_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_extension_opt_item.
	VisitAlter_extension_opt_item(ctx *Alter_extension_opt_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterextensioncontentsstmt.
	VisitAlterextensioncontentsstmt(ctx *AlterextensioncontentsstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createfdwstmt.
	VisitCreatefdwstmt(ctx *CreatefdwstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#fdw_option.
	VisitFdw_option(ctx *Fdw_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#fdw_options.
	VisitFdw_options(ctx *Fdw_optionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_fdw_options.
	VisitOpt_fdw_options(ctx *Opt_fdw_optionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterfdwstmt.
	VisitAlterfdwstmt(ctx *AlterfdwstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#create_generic_options.
	VisitCreate_generic_options(ctx *Create_generic_optionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#generic_option_list.
	VisitGeneric_option_list(ctx *Generic_option_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_generic_options.
	VisitAlter_generic_options(ctx *Alter_generic_optionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_generic_option_list.
	VisitAlter_generic_option_list(ctx *Alter_generic_option_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alter_generic_option_elem.
	VisitAlter_generic_option_elem(ctx *Alter_generic_option_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#generic_option_elem.
	VisitGeneric_option_elem(ctx *Generic_option_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#generic_option_name.
	VisitGeneric_option_name(ctx *Generic_option_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#generic_option_arg.
	VisitGeneric_option_arg(ctx *Generic_option_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createforeignserverstmt.
	VisitCreateforeignserverstmt(ctx *CreateforeignserverstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_type.
	VisitOpt_type(ctx *Opt_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#foreign_server_version.
	VisitForeign_server_version(ctx *Foreign_server_versionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_foreign_server_version.
	VisitOpt_foreign_server_version(ctx *Opt_foreign_server_versionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterforeignserverstmt.
	VisitAlterforeignserverstmt(ctx *AlterforeignserverstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createforeigntablestmt.
	VisitCreateforeigntablestmt(ctx *CreateforeigntablestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#importforeignschemastmt.
	VisitImportforeignschemastmt(ctx *ImportforeignschemastmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#import_qualification_type.
	VisitImport_qualification_type(ctx *Import_qualification_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#import_qualification.
	VisitImport_qualification(ctx *Import_qualificationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createusermappingstmt.
	VisitCreateusermappingstmt(ctx *CreateusermappingstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#auth_ident.
	VisitAuth_ident(ctx *Auth_identContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dropusermappingstmt.
	VisitDropusermappingstmt(ctx *DropusermappingstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterusermappingstmt.
	VisitAlterusermappingstmt(ctx *AlterusermappingstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createpolicystmt.
	VisitCreatepolicystmt(ctx *CreatepolicystmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterpolicystmt.
	VisitAlterpolicystmt(ctx *AlterpolicystmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rowsecurityoptionalexpr.
	VisitRowsecurityoptionalexpr(ctx *RowsecurityoptionalexprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rowsecurityoptionalwithcheck.
	VisitRowsecurityoptionalwithcheck(ctx *RowsecurityoptionalwithcheckContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rowsecuritydefaulttorole.
	VisitRowsecuritydefaulttorole(ctx *RowsecuritydefaulttoroleContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rowsecurityoptionaltorole.
	VisitRowsecurityoptionaltorole(ctx *RowsecurityoptionaltoroleContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rowsecuritydefaultpermissive.
	VisitRowsecuritydefaultpermissive(ctx *RowsecuritydefaultpermissiveContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rowsecuritydefaultforcmd.
	VisitRowsecuritydefaultforcmd(ctx *RowsecuritydefaultforcmdContext) interface{}

	// Visit a parse tree produced by GaussdbParser#row_security_cmd.
	VisitRow_security_cmd(ctx *Row_security_cmdContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createamstmt.
	VisitCreateamstmt(ctx *CreateamstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#am_type.
	VisitAm_type(ctx *Am_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createtrigstmt.
	VisitCreatetrigstmt(ctx *CreatetrigstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggeractiontime.
	VisitTriggeractiontime(ctx *TriggeractiontimeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggerevents.
	VisitTriggerevents(ctx *TriggereventsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggeroneevent.
	VisitTriggeroneevent(ctx *TriggeroneeventContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggerreferencing.
	VisitTriggerreferencing(ctx *TriggerreferencingContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggertransitions.
	VisitTriggertransitions(ctx *TriggertransitionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggertransition.
	VisitTriggertransition(ctx *TriggertransitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#transitionoldornew.
	VisitTransitionoldornew(ctx *TransitionoldornewContext) interface{}

	// Visit a parse tree produced by GaussdbParser#transitionrowortable.
	VisitTransitionrowortable(ctx *TransitionrowortableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#transitionrelname.
	VisitTransitionrelname(ctx *TransitionrelnameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggerforspec.
	VisitTriggerforspec(ctx *TriggerforspecContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggerforopteach.
	VisitTriggerforopteach(ctx *TriggerforopteachContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggerfortype.
	VisitTriggerfortype(ctx *TriggerfortypeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggerwhen.
	VisitTriggerwhen(ctx *TriggerwhenContext) interface{}

	// Visit a parse tree produced by GaussdbParser#function_or_procedure.
	VisitFunction_or_procedure(ctx *Function_or_procedureContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggerfuncargs.
	VisitTriggerfuncargs(ctx *TriggerfuncargsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#triggerfuncarg.
	VisitTriggerfuncarg(ctx *TriggerfuncargContext) interface{}

	// Visit a parse tree produced by GaussdbParser#optconstrfromtable.
	VisitOptconstrfromtable(ctx *OptconstrfromtableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constraintattributespec.
	VisitConstraintattributespec(ctx *ConstraintattributespecContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constraintattributeElem.
	VisitConstraintattributeElem(ctx *ConstraintattributeElemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createeventtrigstmt.
	VisitCreateeventtrigstmt(ctx *CreateeventtrigstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#event_trigger_when_list.
	VisitEvent_trigger_when_list(ctx *Event_trigger_when_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#event_trigger_when_item.
	VisitEvent_trigger_when_item(ctx *Event_trigger_when_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#event_trigger_value_list.
	VisitEvent_trigger_value_list(ctx *Event_trigger_value_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altereventtrigstmt.
	VisitAltereventtrigstmt(ctx *AltereventtrigstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#enable_trigger.
	VisitEnable_trigger(ctx *Enable_triggerContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createassertionstmt.
	VisitCreateassertionstmt(ctx *CreateassertionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#definestmt.
	VisitDefinestmt(ctx *DefinestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#def_list.
	VisitDef_list(ctx *Def_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#def_elem.
	VisitDef_elem(ctx *Def_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#def_arg.
	VisitDef_arg(ctx *Def_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#old_aggr_definition.
	VisitOld_aggr_definition(ctx *Old_aggr_definitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#old_aggr_list.
	VisitOld_aggr_list(ctx *Old_aggr_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#old_aggr_elem.
	VisitOld_aggr_elem(ctx *Old_aggr_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_enum_val_list.
	VisitOpt_enum_val_list(ctx *Opt_enum_val_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#enum_val_list.
	VisitEnum_val_list(ctx *Enum_val_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterenumstmt.
	VisitAlterenumstmt(ctx *AlterenumstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_if_not_exists.
	VisitOpt_if_not_exists(ctx *Opt_if_not_existsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createopclassstmt.
	VisitCreateopclassstmt(ctx *CreateopclassstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opclass_item_list.
	VisitOpclass_item_list(ctx *Opclass_item_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opclass_item.
	VisitOpclass_item(ctx *Opclass_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_default.
	VisitOpt_default(ctx *Opt_defaultContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_opfamily.
	VisitOpt_opfamily(ctx *Opt_opfamilyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opclass_purpose.
	VisitOpclass_purpose(ctx *Opclass_purposeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_recheck.
	VisitOpt_recheck(ctx *Opt_recheckContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createopfamilystmt.
	VisitCreateopfamilystmt(ctx *CreateopfamilystmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alteropfamilystmt.
	VisitAlteropfamilystmt(ctx *AlteropfamilystmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opclass_drop_list.
	VisitOpclass_drop_list(ctx *Opclass_drop_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opclass_drop.
	VisitOpclass_drop(ctx *Opclass_dropContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dropopclassstmt.
	VisitDropopclassstmt(ctx *DropopclassstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dropopfamilystmt.
	VisitDropopfamilystmt(ctx *DropopfamilystmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dropownedstmt.
	VisitDropownedstmt(ctx *DropownedstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reassignownedstmt.
	VisitReassignownedstmt(ctx *ReassignownedstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dropstmt.
	VisitDropstmt(ctx *DropstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#object_type_any_name.
	VisitObject_type_any_name(ctx *Object_type_any_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#object_type_name.
	VisitObject_type_name(ctx *Object_type_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#drop_type_name.
	VisitDrop_type_name(ctx *Drop_type_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#object_type_name_on_any_name.
	VisitObject_type_name_on_any_name(ctx *Object_type_name_on_any_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#any_name_list.
	VisitAny_name_list(ctx *Any_name_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#any_name.
	VisitAny_name(ctx *Any_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#attrs.
	VisitAttrs(ctx *AttrsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#type_name_list.
	VisitType_name_list(ctx *Type_name_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#truncatestmt.
	VisitTruncatestmt(ctx *TruncatestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_restart_seqs.
	VisitOpt_restart_seqs(ctx *Opt_restart_seqsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#commentstmt.
	VisitCommentstmt(ctx *CommentstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#comment_text.
	VisitComment_text(ctx *Comment_textContext) interface{}

	// Visit a parse tree produced by GaussdbParser#seclabelstmt.
	VisitSeclabelstmt(ctx *SeclabelstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_provider.
	VisitOpt_provider(ctx *Opt_providerContext) interface{}

	// Visit a parse tree produced by GaussdbParser#security_label.
	VisitSecurity_label(ctx *Security_labelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cursorstmt.
	VisitCursorstmt(ctx *CursorstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#fetchstmt.
	VisitFetchstmt(ctx *FetchstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#fetch_args.
	VisitFetch_args(ctx *Fetch_argsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#from_in.
	VisitFrom_in(ctx *From_inContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_from_in.
	VisitOpt_from_in(ctx *Opt_from_inContext) interface{}

	// Visit a parse tree produced by GaussdbParser#grantstmt.
	VisitGrantstmt(ctx *GrantstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#revokestmt.
	VisitRevokestmt(ctx *RevokestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#privileges.
	VisitPrivileges(ctx *PrivilegesContext) interface{}

	// Visit a parse tree produced by GaussdbParser#privilege_list.
	VisitPrivilege_list(ctx *Privilege_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#privilege.
	VisitPrivilege(ctx *PrivilegeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#privilege_target.
	VisitPrivilege_target(ctx *Privilege_targetContext) interface{}

	// Visit a parse tree produced by GaussdbParser#grantee_list.
	VisitGrantee_list(ctx *Grantee_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#grantee.
	VisitGrantee(ctx *GranteeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_grant_grant_option.
	VisitOpt_grant_grant_option(ctx *Opt_grant_grant_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#grantrolestmt.
	VisitGrantrolestmt(ctx *GrantrolestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#revokerolestmt.
	VisitRevokerolestmt(ctx *RevokerolestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_grant_admin_option.
	VisitOpt_grant_admin_option(ctx *Opt_grant_admin_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_granted_by.
	VisitOpt_granted_by(ctx *Opt_granted_byContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterdefaultprivilegesstmt.
	VisitAlterdefaultprivilegesstmt(ctx *AlterdefaultprivilegesstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#defacloptionlist.
	VisitDefacloptionlist(ctx *DefacloptionlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#defacloption.
	VisitDefacloption(ctx *DefacloptionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#defaclaction.
	VisitDefaclaction(ctx *DefaclactionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#defacl_privilege_target.
	VisitDefacl_privilege_target(ctx *Defacl_privilege_targetContext) interface{}

	// Visit a parse tree produced by GaussdbParser#indexstmt.
	VisitIndexstmt(ctx *IndexstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_opt_distribute.
	VisitGaussdb_opt_distribute(ctx *Gaussdb_opt_distributeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_unique.
	VisitOpt_unique(ctx *Opt_uniqueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_concurrently.
	VisitOpt_concurrently(ctx *Opt_concurrentlyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_index_name.
	VisitOpt_index_name(ctx *Opt_index_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#access_method_clause.
	VisitAccess_method_clause(ctx *Access_method_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#index_params.
	VisitIndex_params(ctx *Index_paramsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#index_elem_options.
	VisitIndex_elem_options(ctx *Index_elem_optionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#index_elem.
	VisitIndex_elem(ctx *Index_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_include.
	VisitOpt_include(ctx *Opt_includeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#index_including_params.
	VisitIndex_including_params(ctx *Index_including_paramsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_collate.
	VisitOpt_collate(ctx *Opt_collateContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_class.
	VisitOpt_class(ctx *Opt_classContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_asc_desc.
	VisitOpt_asc_desc(ctx *Opt_asc_descContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_nulls_order.
	VisitOpt_nulls_order(ctx *Opt_nulls_orderContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createfunctionstmt.
	VisitCreatefunctionstmt(ctx *CreatefunctionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_or_replace.
	VisitOpt_or_replace(ctx *Opt_or_replaceContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_args.
	VisitFunc_args(ctx *Func_argsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_args_list.
	VisitFunc_args_list(ctx *Func_args_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#function_with_argtypes_list.
	VisitFunction_with_argtypes_list(ctx *Function_with_argtypes_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#function_with_argtypes.
	VisitFunction_with_argtypes(ctx *Function_with_argtypesContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_args_with_defaults.
	VisitFunc_args_with_defaults(ctx *Func_args_with_defaultsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_args_with_defaults_list.
	VisitFunc_args_with_defaults_list(ctx *Func_args_with_defaults_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_arg.
	VisitFunc_arg(ctx *Func_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#arg_class.
	VisitArg_class(ctx *Arg_classContext) interface{}

	// Visit a parse tree produced by GaussdbParser#param_name.
	VisitParam_name(ctx *Param_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_return.
	VisitFunc_return(ctx *Func_returnContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_type.
	VisitFunc_type(ctx *Func_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_arg_with_default.
	VisitFunc_arg_with_default(ctx *Func_arg_with_defaultContext) interface{}

	// Visit a parse tree produced by GaussdbParser#aggr_arg.
	VisitAggr_arg(ctx *Aggr_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#aggr_args.
	VisitAggr_args(ctx *Aggr_argsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#aggr_args_list.
	VisitAggr_args_list(ctx *Aggr_args_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#aggregate_with_argtypes.
	VisitAggregate_with_argtypes(ctx *Aggregate_with_argtypesContext) interface{}

	// Visit a parse tree produced by GaussdbParser#aggregate_with_argtypes_list.
	VisitAggregate_with_argtypes_list(ctx *Aggregate_with_argtypes_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createfunc_opt_list.
	VisitCreatefunc_opt_list(ctx *Createfunc_opt_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#common_func_opt_item.
	VisitCommon_func_opt_item(ctx *Common_func_opt_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createfunc_opt_item.
	VisitCreatefunc_opt_item(ctx *Createfunc_opt_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_as.
	VisitFunc_as(ctx *Func_asContext) interface{}

	// Visit a parse tree produced by GaussdbParser#transform_type_list.
	VisitTransform_type_list(ctx *Transform_type_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_definition.
	VisitOpt_definition(ctx *Opt_definitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#table_func_column.
	VisitTable_func_column(ctx *Table_func_columnContext) interface{}

	// Visit a parse tree produced by GaussdbParser#table_func_column_list.
	VisitTable_func_column_list(ctx *Table_func_column_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterfunctionstmt.
	VisitAlterfunctionstmt(ctx *AlterfunctionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterfunc_opt_list.
	VisitAlterfunc_opt_list(ctx *Alterfunc_opt_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_restrict.
	VisitOpt_restrict(ctx *Opt_restrictContext) interface{}

	// Visit a parse tree produced by GaussdbParser#removefuncstmt.
	VisitRemovefuncstmt(ctx *RemovefuncstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#removeaggrstmt.
	VisitRemoveaggrstmt(ctx *RemoveaggrstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#removeoperstmt.
	VisitRemoveoperstmt(ctx *RemoveoperstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#oper_argtypes.
	VisitOper_argtypes(ctx *Oper_argtypesContext) interface{}

	// Visit a parse tree produced by GaussdbParser#any_operator.
	VisitAny_operator(ctx *Any_operatorContext) interface{}

	// Visit a parse tree produced by GaussdbParser#operator_with_argtypes_list.
	VisitOperator_with_argtypes_list(ctx *Operator_with_argtypes_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#operator_with_argtypes.
	VisitOperator_with_argtypes(ctx *Operator_with_argtypesContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dostmt.
	VisitDostmt(ctx *DostmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dostmt_opt_list.
	VisitDostmt_opt_list(ctx *Dostmt_opt_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dostmt_opt_item.
	VisitDostmt_opt_item(ctx *Dostmt_opt_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createcaststmt.
	VisitCreatecaststmt(ctx *CreatecaststmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cast_context.
	VisitCast_context(ctx *Cast_contextContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dropcaststmt.
	VisitDropcaststmt(ctx *DropcaststmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_if_exists.
	VisitOpt_if_exists(ctx *Opt_if_existsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createtransformstmt.
	VisitCreatetransformstmt(ctx *CreatetransformstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#transform_element_list.
	VisitTransform_element_list(ctx *Transform_element_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#droptransformstmt.
	VisitDroptransformstmt(ctx *DroptransformstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reindexstmt.
	VisitReindexstmt(ctx *ReindexstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reindex_target_type.
	VisitReindex_target_type(ctx *Reindex_target_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reindex_target_multitable.
	VisitReindex_target_multitable(ctx *Reindex_target_multitableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reindex_option_list.
	VisitReindex_option_list(ctx *Reindex_option_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reindex_option_elem.
	VisitReindex_option_elem(ctx *Reindex_option_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altertblspcstmt.
	VisitAltertblspcstmt(ctx *AltertblspcstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#renamestmt.
	VisitRenamestmt(ctx *RenamestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_column.
	VisitOpt_column(ctx *Opt_columnContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_set_data.
	VisitOpt_set_data(ctx *Opt_set_dataContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterobjectdependsstmt.
	VisitAlterobjectdependsstmt(ctx *AlterobjectdependsstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_no.
	VisitOpt_no(ctx *Opt_noContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterobjectschemastmt.
	VisitAlterobjectschemastmt(ctx *AlterobjectschemastmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alteroperatorstmt.
	VisitAlteroperatorstmt(ctx *AlteroperatorstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#operator_def_list.
	VisitOperator_def_list(ctx *Operator_def_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#operator_def_elem.
	VisitOperator_def_elem(ctx *Operator_def_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#operator_def_arg.
	VisitOperator_def_arg(ctx *Operator_def_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altertypestmt.
	VisitAltertypestmt(ctx *AltertypestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterownerstmt.
	VisitAlterownerstmt(ctx *AlterownerstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createpublicationstmt.
	VisitCreatepublicationstmt(ctx *CreatepublicationstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_publication_for_tables.
	VisitOpt_publication_for_tables(ctx *Opt_publication_for_tablesContext) interface{}

	// Visit a parse tree produced by GaussdbParser#publication_for_tables.
	VisitPublication_for_tables(ctx *Publication_for_tablesContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterpublicationstmt.
	VisitAlterpublicationstmt(ctx *AlterpublicationstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createsubscriptionstmt.
	VisitCreatesubscriptionstmt(ctx *CreatesubscriptionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#publication_name_list.
	VisitPublication_name_list(ctx *Publication_name_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#publication_name_item.
	VisitPublication_name_item(ctx *Publication_name_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altersubscriptionstmt.
	VisitAltersubscriptionstmt(ctx *AltersubscriptionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dropsubscriptionstmt.
	VisitDropsubscriptionstmt(ctx *DropsubscriptionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rulestmt.
	VisitRulestmt(ctx *RulestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#ruleactionlist.
	VisitRuleactionlist(ctx *RuleactionlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#ruleactionmulti.
	VisitRuleactionmulti(ctx *RuleactionmultiContext) interface{}

	// Visit a parse tree produced by GaussdbParser#ruleactionstmt.
	VisitRuleactionstmt(ctx *RuleactionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#ruleactionstmtOrEmpty.
	VisitRuleactionstmtOrEmpty(ctx *RuleactionstmtOrEmptyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#event.
	VisitEvent(ctx *EventContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_instead.
	VisitOpt_instead(ctx *Opt_insteadContext) interface{}

	// Visit a parse tree produced by GaussdbParser#notifystmt.
	VisitNotifystmt(ctx *NotifystmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#notify_payload.
	VisitNotify_payload(ctx *Notify_payloadContext) interface{}

	// Visit a parse tree produced by GaussdbParser#listenstmt.
	VisitListenstmt(ctx *ListenstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#unlistenstmt.
	VisitUnlistenstmt(ctx *UnlistenstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#transactionstmt.
	VisitTransactionstmt(ctx *TransactionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_transaction.
	VisitOpt_transaction(ctx *Opt_transactionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#transaction_mode_item.
	VisitTransaction_mode_item(ctx *Transaction_mode_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#transaction_mode_list.
	VisitTransaction_mode_list(ctx *Transaction_mode_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#transaction_mode_list_or_empty.
	VisitTransaction_mode_list_or_empty(ctx *Transaction_mode_list_or_emptyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_transaction_chain.
	VisitOpt_transaction_chain(ctx *Opt_transaction_chainContext) interface{}

	// Visit a parse tree produced by GaussdbParser#viewstmt.
	VisitViewstmt(ctx *ViewstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_check_option.
	VisitOpt_check_option(ctx *Opt_check_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#loadstmt.
	VisitLoadstmt(ctx *LoadstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createdbstmt.
	VisitCreatedbstmt(ctx *CreatedbstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createdb_opt_list.
	VisitCreatedb_opt_list(ctx *Createdb_opt_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createdb_opt_items.
	VisitCreatedb_opt_items(ctx *Createdb_opt_itemsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createdb_opt_item.
	VisitCreatedb_opt_item(ctx *Createdb_opt_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createdb_opt_name.
	VisitCreatedb_opt_name(ctx *Createdb_opt_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_equal.
	VisitOpt_equal(ctx *Opt_equalContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterdatabasestmt.
	VisitAlterdatabasestmt(ctx *AlterdatabasestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterdatabasesetstmt.
	VisitAlterdatabasesetstmt(ctx *AlterdatabasesetstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dropdbstmt.
	VisitDropdbstmt(ctx *DropdbstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#drop_option_list.
	VisitDrop_option_list(ctx *Drop_option_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#drop_option.
	VisitDrop_option(ctx *Drop_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altercollationstmt.
	VisitAltercollationstmt(ctx *AltercollationstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altersystemstmt.
	VisitAltersystemstmt(ctx *AltersystemstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createdomainstmt.
	VisitCreatedomainstmt(ctx *CreatedomainstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alterdomainstmt.
	VisitAlterdomainstmt(ctx *AlterdomainstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_as.
	VisitOpt_as(ctx *Opt_asContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altertsdictionarystmt.
	VisitAltertsdictionarystmt(ctx *AltertsdictionarystmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#altertsconfigurationstmt.
	VisitAltertsconfigurationstmt(ctx *AltertsconfigurationstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#any_with.
	VisitAny_with(ctx *Any_withContext) interface{}

	// Visit a parse tree produced by GaussdbParser#createconversionstmt.
	VisitCreateconversionstmt(ctx *CreateconversionstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#clusterstmt.
	VisitClusterstmt(ctx *ClusterstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cluster_index_specification.
	VisitCluster_index_specification(ctx *Cluster_index_specificationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#vacuumstmt.
	VisitVacuumstmt(ctx *VacuumstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#analyzestmt.
	VisitAnalyzestmt(ctx *AnalyzestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#vac_analyze_option_list.
	VisitVac_analyze_option_list(ctx *Vac_analyze_option_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#analyze_keyword.
	VisitAnalyze_keyword(ctx *Analyze_keywordContext) interface{}

	// Visit a parse tree produced by GaussdbParser#vac_analyze_option_elem.
	VisitVac_analyze_option_elem(ctx *Vac_analyze_option_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#vac_analyze_option_name.
	VisitVac_analyze_option_name(ctx *Vac_analyze_option_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#vac_analyze_option_arg.
	VisitVac_analyze_option_arg(ctx *Vac_analyze_option_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_analyze.
	VisitOpt_analyze(ctx *Opt_analyzeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_verbose.
	VisitOpt_verbose(ctx *Opt_verboseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_full.
	VisitOpt_full(ctx *Opt_fullContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_freeze.
	VisitOpt_freeze(ctx *Opt_freezeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_name_list.
	VisitOpt_name_list(ctx *Opt_name_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#vacuum_relation.
	VisitVacuum_relation(ctx *Vacuum_relationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#vacuum_relation_list.
	VisitVacuum_relation_list(ctx *Vacuum_relation_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_vacuum_relation_list.
	VisitOpt_vacuum_relation_list(ctx *Opt_vacuum_relation_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#explainplanstmt.
	VisitExplainplanstmt(ctx *ExplainplanstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#explainstmt.
	VisitExplainstmt(ctx *ExplainstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#explainablestmt.
	VisitExplainablestmt(ctx *ExplainablestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#explain_option_list.
	VisitExplain_option_list(ctx *Explain_option_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#explain_option_elem.
	VisitExplain_option_elem(ctx *Explain_option_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#explain_option_name.
	VisitExplain_option_name(ctx *Explain_option_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#explain_option_arg.
	VisitExplain_option_arg(ctx *Explain_option_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#preparestmt.
	VisitPreparestmt(ctx *PreparestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#prep_type_clause.
	VisitPrep_type_clause(ctx *Prep_type_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#preparablestmt.
	VisitPreparablestmt(ctx *PreparablestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#executestmt.
	VisitExecutestmt(ctx *ExecutestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#execute_param_clause.
	VisitExecute_param_clause(ctx *Execute_param_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#deallocatestmt.
	VisitDeallocatestmt(ctx *DeallocatestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#insertstmt.
	VisitInsertstmt(ctx *InsertstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#insert_target.
	VisitInsert_target(ctx *Insert_targetContext) interface{}

	// Visit a parse tree produced by GaussdbParser#insert_rest.
	VisitInsert_rest(ctx *Insert_restContext) interface{}

	// Visit a parse tree produced by GaussdbParser#override_kind.
	VisitOverride_kind(ctx *Override_kindContext) interface{}

	// Visit a parse tree produced by GaussdbParser#insert_column_list.
	VisitInsert_column_list(ctx *Insert_column_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#insert_column_item.
	VisitInsert_column_item(ctx *Insert_column_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_on_conflict.
	VisitOpt_on_conflict(ctx *Opt_on_conflictContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_conf_expr.
	VisitOpt_conf_expr(ctx *Opt_conf_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#returning_clause.
	VisitReturning_clause(ctx *Returning_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#mergestmt.
	VisitMergestmt(ctx *MergestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#merge_insert_clause.
	VisitMerge_insert_clause(ctx *Merge_insert_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#merge_update_clause.
	VisitMerge_update_clause(ctx *Merge_update_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#merge_delete_clause.
	VisitMerge_delete_clause(ctx *Merge_delete_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#deletestmt.
	VisitDeletestmt(ctx *DeletestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#using_clause.
	VisitUsing_clause(ctx *Using_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#lockstmt.
	VisitLockstmt(ctx *LockstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#lockbucketsstmt.
	VisitLockbucketsstmt(ctx *LockbucketsstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#markbucketsstmt.
	VisitMarkbucketsstmt(ctx *MarkbucketsstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_bucketlist.
	VisitOpt_bucketlist(ctx *Opt_bucketlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_lock.
	VisitOpt_lock(ctx *Opt_lockContext) interface{}

	// Visit a parse tree produced by GaussdbParser#lock_type.
	VisitLock_type(ctx *Lock_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_nowait.
	VisitOpt_nowait(ctx *Opt_nowaitContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_nowait_or_skip.
	VisitOpt_nowait_or_skip(ctx *Opt_nowait_or_skipContext) interface{}

	// Visit a parse tree produced by GaussdbParser#updatestmt.
	VisitUpdatestmt(ctx *UpdatestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#set_clause_list.
	VisitSet_clause_list(ctx *Set_clause_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#set_clause.
	VisitSet_clause(ctx *Set_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#set_target.
	VisitSet_target(ctx *Set_targetContext) interface{}

	// Visit a parse tree produced by GaussdbParser#set_target_list.
	VisitSet_target_list(ctx *Set_target_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#declarecursorstmt.
	VisitDeclarecursorstmt(ctx *DeclarecursorstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cursor_name.
	VisitCursor_name(ctx *Cursor_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cursor_options.
	VisitCursor_options(ctx *Cursor_optionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_hold.
	VisitOpt_hold(ctx *Opt_holdContext) interface{}

	// Visit a parse tree produced by GaussdbParser#selectstmt.
	VisitSelectstmt(ctx *SelectstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#select_with_parens.
	VisitSelect_with_parens(ctx *Select_with_parensContext) interface{}

	// Visit a parse tree produced by GaussdbParser#select_no_parens.
	VisitSelect_no_parens(ctx *Select_no_parensContext) interface{}

	// Visit a parse tree produced by GaussdbParser#select_clause.
	VisitSelect_clause(ctx *Select_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#simple_select_intersect.
	VisitSimple_select_intersect(ctx *Simple_select_intersectContext) interface{}

	// Visit a parse tree produced by GaussdbParser#simple_select_pramary.
	VisitSimple_select_pramary(ctx *Simple_select_pramaryContext) interface{}

	// Visit a parse tree produced by GaussdbParser#with_clause.
	VisitWith_clause(ctx *With_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cte_list.
	VisitCte_list(ctx *Cte_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#common_table_expr.
	VisitCommon_table_expr(ctx *Common_table_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_materialized.
	VisitOpt_materialized(ctx *Opt_materializedContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_with_clause.
	VisitOpt_with_clause(ctx *Opt_with_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#into_clause.
	VisitInto_clause(ctx *Into_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_strict.
	VisitOpt_strict(ctx *Opt_strictContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opttempTableName.
	VisitOpttempTableName(ctx *OpttempTableNameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_table.
	VisitOpt_table(ctx *Opt_tableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#all_or_distinct.
	VisitAll_or_distinct(ctx *All_or_distinctContext) interface{}

	// Visit a parse tree produced by GaussdbParser#distinct_clause.
	VisitDistinct_clause(ctx *Distinct_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_all_clause.
	VisitOpt_all_clause(ctx *Opt_all_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_sort_clause.
	VisitOpt_sort_clause(ctx *Opt_sort_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#sort_clause.
	VisitSort_clause(ctx *Sort_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#sortby_list.
	VisitSortby_list(ctx *Sortby_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#sortby.
	VisitSortby(ctx *SortbyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#select_limit.
	VisitSelect_limit(ctx *Select_limitContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_select_limit.
	VisitOpt_select_limit(ctx *Opt_select_limitContext) interface{}

	// Visit a parse tree produced by GaussdbParser#limit_clause.
	VisitLimit_clause(ctx *Limit_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#offset_clause.
	VisitOffset_clause(ctx *Offset_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#select_limit_value.
	VisitSelect_limit_value(ctx *Select_limit_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#select_offset_value.
	VisitSelect_offset_value(ctx *Select_offset_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#select_fetch_first_value.
	VisitSelect_fetch_first_value(ctx *Select_fetch_first_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#i_or_f_const.
	VisitI_or_f_const(ctx *I_or_f_constContext) interface{}

	// Visit a parse tree produced by GaussdbParser#row_or_rows.
	VisitRow_or_rows(ctx *Row_or_rowsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#first_or_next.
	VisitFirst_or_next(ctx *First_or_nextContext) interface{}

	// Visit a parse tree produced by GaussdbParser#group_clause.
	VisitGroup_clause(ctx *Group_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#group_by_list.
	VisitGroup_by_list(ctx *Group_by_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#group_by_item.
	VisitGroup_by_item(ctx *Group_by_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#empty_grouping_set.
	VisitEmpty_grouping_set(ctx *Empty_grouping_setContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rollup_clause.
	VisitRollup_clause(ctx *Rollup_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cube_clause.
	VisitCube_clause(ctx *Cube_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#grouping_sets_clause.
	VisitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#having_clause.
	VisitHaving_clause(ctx *Having_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#for_locking_clause.
	VisitFor_locking_clause(ctx *For_locking_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_for_locking_clause.
	VisitOpt_for_locking_clause(ctx *Opt_for_locking_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#for_locking_items.
	VisitFor_locking_items(ctx *For_locking_itemsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#for_locking_item.
	VisitFor_locking_item(ctx *For_locking_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#for_locking_strength.
	VisitFor_locking_strength(ctx *For_locking_strengthContext) interface{}

	// Visit a parse tree produced by GaussdbParser#locked_rels_list.
	VisitLocked_rels_list(ctx *Locked_rels_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#values_clause.
	VisitValues_clause(ctx *Values_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#from_list.
	VisitFrom_list(ctx *From_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#table_ref.
	VisitTable_ref(ctx *Table_refContext) interface{}

	// Visit a parse tree produced by GaussdbParser#joined_table.
	VisitJoined_table(ctx *Joined_tableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#alias_clause.
	VisitAlias_clause(ctx *Alias_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_alias_clause.
	VisitOpt_alias_clause(ctx *Opt_alias_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#table_alias_clause.
	VisitTable_alias_clause(ctx *Table_alias_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_alias_clause.
	VisitFunc_alias_clause(ctx *Func_alias_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#join_type.
	VisitJoin_type(ctx *Join_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#join_qual.
	VisitJoin_qual(ctx *Join_qualContext) interface{}

	// Visit a parse tree produced by GaussdbParser#relation_expr.
	VisitRelation_expr(ctx *Relation_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#relation_expr_list.
	VisitRelation_expr_list(ctx *Relation_expr_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#relation_expr_opt_alias.
	VisitRelation_expr_opt_alias(ctx *Relation_expr_opt_aliasContext) interface{}

	// Visit a parse tree produced by GaussdbParser#tablesample_clause.
	VisitTablesample_clause(ctx *Tablesample_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_repeatable_clause.
	VisitOpt_repeatable_clause(ctx *Opt_repeatable_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_table.
	VisitFunc_table(ctx *Func_tableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rowsfrom_item.
	VisitRowsfrom_item(ctx *Rowsfrom_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rowsfrom_list.
	VisitRowsfrom_list(ctx *Rowsfrom_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_col_def_list.
	VisitOpt_col_def_list(ctx *Opt_col_def_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_ordinality.
	VisitOpt_ordinality(ctx *Opt_ordinalityContext) interface{}

	// Visit a parse tree produced by GaussdbParser#where_clause.
	VisitWhere_clause(ctx *Where_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#where_or_current_clause.
	VisitWhere_or_current_clause(ctx *Where_or_current_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opttablefuncelementlist.
	VisitOpttablefuncelementlist(ctx *OpttablefuncelementlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#tablefuncelementlist.
	VisitTablefuncelementlist(ctx *TablefuncelementlistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#tablefuncelement.
	VisitTablefuncelement(ctx *TablefuncelementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xmltable.
	VisitXmltable(ctx *XmltableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xmltable_column_list.
	VisitXmltable_column_list(ctx *Xmltable_column_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xmltable_column_el.
	VisitXmltable_column_el(ctx *Xmltable_column_elContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xmltable_column_option_list.
	VisitXmltable_column_option_list(ctx *Xmltable_column_option_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xmltable_column_option_el.
	VisitXmltable_column_option_el(ctx *Xmltable_column_option_elContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xml_namespace_list.
	VisitXml_namespace_list(ctx *Xml_namespace_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xml_namespace_el.
	VisitXml_namespace_el(ctx *Xml_namespace_elContext) interface{}

	// Visit a parse tree produced by GaussdbParser#typename.
	VisitTypename(ctx *TypenameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_array_bounds.
	VisitOpt_array_bounds(ctx *Opt_array_boundsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#simpletypename.
	VisitSimpletypename(ctx *SimpletypenameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#consttypename.
	VisitConsttypename(ctx *ConsttypenameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#generictype.
	VisitGenerictype(ctx *GenerictypeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_type_modifiers.
	VisitOpt_type_modifiers(ctx *Opt_type_modifiersContext) interface{}

	// Visit a parse tree produced by GaussdbParser#numeric.
	VisitNumeric(ctx *NumericContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_float.
	VisitOpt_float(ctx *Opt_floatContext) interface{}

	// Visit a parse tree produced by GaussdbParser#bit.
	VisitBit(ctx *BitContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constbit.
	VisitConstbit(ctx *ConstbitContext) interface{}

	// Visit a parse tree produced by GaussdbParser#bitwithlength.
	VisitBitwithlength(ctx *BitwithlengthContext) interface{}

	// Visit a parse tree produced by GaussdbParser#bitwithoutlength.
	VisitBitwithoutlength(ctx *BitwithoutlengthContext) interface{}

	// Visit a parse tree produced by GaussdbParser#character.
	VisitCharacter(ctx *CharacterContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constcharacter.
	VisitConstcharacter(ctx *ConstcharacterContext) interface{}

	// Visit a parse tree produced by GaussdbParser#character_c.
	VisitCharacter_c(ctx *Character_cContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_varying.
	VisitOpt_varying(ctx *Opt_varyingContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constdatetime.
	VisitConstdatetime(ctx *ConstdatetimeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constinterval.
	VisitConstinterval(ctx *ConstintervalContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_timezone.
	VisitOpt_timezone(ctx *Opt_timezoneContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_interval.
	VisitOpt_interval(ctx *Opt_intervalContext) interface{}

	// Visit a parse tree produced by GaussdbParser#interval_second.
	VisitInterval_second(ctx *Interval_secondContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_escape.
	VisitOpt_escape(ctx *Opt_escapeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr.
	VisitA_expr(ctx *A_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_qual.
	VisitA_expr_qual(ctx *A_expr_qualContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_lessless.
	VisitA_expr_lessless(ctx *A_expr_lesslessContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_or.
	VisitA_expr_or(ctx *A_expr_orContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_and.
	VisitA_expr_and(ctx *A_expr_andContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_between.
	VisitA_expr_between(ctx *A_expr_betweenContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_in.
	VisitA_expr_in(ctx *A_expr_inContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_unary_not.
	VisitA_expr_unary_not(ctx *A_expr_unary_notContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_isnull.
	VisitA_expr_isnull(ctx *A_expr_isnullContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_is_not.
	VisitA_expr_is_not(ctx *A_expr_is_notContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_compare.
	VisitA_expr_compare(ctx *A_expr_compareContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_like.
	VisitA_expr_like(ctx *A_expr_likeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_qual_op.
	VisitA_expr_qual_op(ctx *A_expr_qual_opContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_unary_qualop.
	VisitA_expr_unary_qualop(ctx *A_expr_unary_qualopContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_add.
	VisitA_expr_add(ctx *A_expr_addContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_mul.
	VisitA_expr_mul(ctx *A_expr_mulContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_caret.
	VisitA_expr_caret(ctx *A_expr_caretContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_unary_sign.
	VisitA_expr_unary_sign(ctx *A_expr_unary_signContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_at_time_zone.
	VisitA_expr_at_time_zone(ctx *A_expr_at_time_zoneContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_collate.
	VisitA_expr_collate(ctx *A_expr_collateContext) interface{}

	// Visit a parse tree produced by GaussdbParser#a_expr_typecast.
	VisitA_expr_typecast(ctx *A_expr_typecastContext) interface{}

	// Visit a parse tree produced by GaussdbParser#b_expr.
	VisitB_expr(ctx *B_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#c_expr_exists.
	VisitC_expr_exists(ctx *C_expr_existsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#c_expr_expr.
	VisitC_expr_expr(ctx *C_expr_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#c_expr_case.
	VisitC_expr_case(ctx *C_expr_caseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#plsqlvariablename.
	VisitPlsqlvariablename(ctx *PlsqlvariablenameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_application.
	VisitFunc_application(ctx *Func_applicationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_expr.
	VisitFunc_expr(ctx *Func_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_expr_windowless.
	VisitFunc_expr_windowless(ctx *Func_expr_windowlessContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_expr_common_subexpr.
	VisitFunc_expr_common_subexpr(ctx *Func_expr_common_subexprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xml_root_version.
	VisitXml_root_version(ctx *Xml_root_versionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_xml_root_standalone.
	VisitOpt_xml_root_standalone(ctx *Opt_xml_root_standaloneContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xml_attributes.
	VisitXml_attributes(ctx *Xml_attributesContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xml_attribute_list.
	VisitXml_attribute_list(ctx *Xml_attribute_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xml_attribute_el.
	VisitXml_attribute_el(ctx *Xml_attribute_elContext) interface{}

	// Visit a parse tree produced by GaussdbParser#document_or_content.
	VisitDocument_or_content(ctx *Document_or_contentContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xml_whitespace_option.
	VisitXml_whitespace_option(ctx *Xml_whitespace_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xmlexists_argument.
	VisitXmlexists_argument(ctx *Xmlexists_argumentContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xml_passing_mech.
	VisitXml_passing_mech(ctx *Xml_passing_mechContext) interface{}

	// Visit a parse tree produced by GaussdbParser#within_group_clause.
	VisitWithin_group_clause(ctx *Within_group_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#filter_clause.
	VisitFilter_clause(ctx *Filter_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#window_clause.
	VisitWindow_clause(ctx *Window_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#window_definition_list.
	VisitWindow_definition_list(ctx *Window_definition_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#window_definition.
	VisitWindow_definition(ctx *Window_definitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#over_clause.
	VisitOver_clause(ctx *Over_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#window_specification.
	VisitWindow_specification(ctx *Window_specificationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_existing_window_name.
	VisitOpt_existing_window_name(ctx *Opt_existing_window_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_partition_clause.
	VisitOpt_partition_clause(ctx *Opt_partition_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_frame_clause.
	VisitOpt_frame_clause(ctx *Opt_frame_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#frame_extent.
	VisitFrame_extent(ctx *Frame_extentContext) interface{}

	// Visit a parse tree produced by GaussdbParser#frame_bound.
	VisitFrame_bound(ctx *Frame_boundContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_window_exclusion_clause.
	VisitOpt_window_exclusion_clause(ctx *Opt_window_exclusion_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#row.
	VisitRow(ctx *RowContext) interface{}

	// Visit a parse tree produced by GaussdbParser#explicit_row.
	VisitExplicit_row(ctx *Explicit_rowContext) interface{}

	// Visit a parse tree produced by GaussdbParser#implicit_row.
	VisitImplicit_row(ctx *Implicit_rowContext) interface{}

	// Visit a parse tree produced by GaussdbParser#sub_type.
	VisitSub_type(ctx *Sub_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#all_op.
	VisitAll_op(ctx *All_opContext) interface{}

	// Visit a parse tree produced by GaussdbParser#mathop.
	VisitMathop(ctx *MathopContext) interface{}

	// Visit a parse tree produced by GaussdbParser#qual_op.
	VisitQual_op(ctx *Qual_opContext) interface{}

	// Visit a parse tree produced by GaussdbParser#qual_all_op.
	VisitQual_all_op(ctx *Qual_all_opContext) interface{}

	// Visit a parse tree produced by GaussdbParser#subquery_Op.
	VisitSubquery_Op(ctx *Subquery_OpContext) interface{}

	// Visit a parse tree produced by GaussdbParser#expr_list.
	VisitExpr_list(ctx *Expr_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_arg_list.
	VisitFunc_arg_list(ctx *Func_arg_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_arg_expr.
	VisitFunc_arg_expr(ctx *Func_arg_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#type_list.
	VisitType_list(ctx *Type_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#array_expr.
	VisitArray_expr(ctx *Array_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#array_expr_list.
	VisitArray_expr_list(ctx *Array_expr_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#extract_list.
	VisitExtract_list(ctx *Extract_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#extract_arg.
	VisitExtract_arg(ctx *Extract_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#unicode_normal_form.
	VisitUnicode_normal_form(ctx *Unicode_normal_formContext) interface{}

	// Visit a parse tree produced by GaussdbParser#overlay_list.
	VisitOverlay_list(ctx *Overlay_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#position_list.
	VisitPosition_list(ctx *Position_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#substr_list.
	VisitSubstr_list(ctx *Substr_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#trim_list.
	VisitTrim_list(ctx *Trim_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#in_expr_select.
	VisitIn_expr_select(ctx *In_expr_selectContext) interface{}

	// Visit a parse tree produced by GaussdbParser#in_expr_list.
	VisitIn_expr_list(ctx *In_expr_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#case_expr.
	VisitCase_expr(ctx *Case_exprContext) interface{}

	// Visit a parse tree produced by GaussdbParser#when_clause_list.
	VisitWhen_clause_list(ctx *When_clause_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#when_clause.
	VisitWhen_clause(ctx *When_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#case_default.
	VisitCase_default(ctx *Case_defaultContext) interface{}

	// Visit a parse tree produced by GaussdbParser#case_arg.
	VisitCase_arg(ctx *Case_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#columnref.
	VisitColumnref(ctx *ColumnrefContext) interface{}

	// Visit a parse tree produced by GaussdbParser#indirection_el.
	VisitIndirection_el(ctx *Indirection_elContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_slice_bound.
	VisitOpt_slice_bound(ctx *Opt_slice_boundContext) interface{}

	// Visit a parse tree produced by GaussdbParser#indirection.
	VisitIndirection(ctx *IndirectionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_indirection.
	VisitOpt_indirection(ctx *Opt_indirectionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_target_list.
	VisitOpt_target_list(ctx *Opt_target_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#target_list.
	VisitTarget_list(ctx *Target_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#target_label.
	VisitTarget_label(ctx *Target_labelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#target_star.
	VisitTarget_star(ctx *Target_starContext) interface{}

	// Visit a parse tree produced by GaussdbParser#target_alias.
	VisitTarget_alias(ctx *Target_aliasContext) interface{}

	// Visit a parse tree produced by GaussdbParser#qualified_name_list.
	VisitQualified_name_list(ctx *Qualified_name_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#qualified_name.
	VisitQualified_name(ctx *Qualified_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#link_name.
	VisitLink_name(ctx *Link_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#name_list.
	VisitName_list(ctx *Name_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#attr_name.
	VisitAttr_name(ctx *Attr_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#file_name.
	VisitFile_name(ctx *File_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#func_name.
	VisitFunc_name(ctx *Func_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#aexprconst.
	VisitAexprconst(ctx *AexprconstContext) interface{}

	// Visit a parse tree produced by GaussdbParser#xconst.
	VisitXconst(ctx *XconstContext) interface{}

	// Visit a parse tree produced by GaussdbParser#bconst.
	VisitBconst(ctx *BconstContext) interface{}

	// Visit a parse tree produced by GaussdbParser#fconst.
	VisitFconst(ctx *FconstContext) interface{}

	// Visit a parse tree produced by GaussdbParser#iconst.
	VisitIconst(ctx *IconstContext) interface{}

	// Visit a parse tree produced by GaussdbParser#sconst.
	VisitSconst(ctx *SconstContext) interface{}

	// Visit a parse tree produced by GaussdbParser#anysconst.
	VisitAnysconst(ctx *AnysconstContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_uescape.
	VisitOpt_uescape(ctx *Opt_uescapeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#signediconst.
	VisitSignediconst(ctx *SignediconstContext) interface{}

	// Visit a parse tree produced by GaussdbParser#roleid.
	VisitRoleid(ctx *RoleidContext) interface{}

	// Visit a parse tree produced by GaussdbParser#rolespec.
	VisitRolespec(ctx *RolespecContext) interface{}

	// Visit a parse tree produced by GaussdbParser#role_list.
	VisitRole_list(ctx *Role_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#colid.
	VisitColid(ctx *ColidContext) interface{}

	// Visit a parse tree produced by GaussdbParser#table_alias.
	VisitTable_alias(ctx *Table_aliasContext) interface{}

	// Visit a parse tree produced by GaussdbParser#type_function_name.
	VisitType_function_name(ctx *Type_function_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#nonreservedword.
	VisitNonreservedword(ctx *NonreservedwordContext) interface{}

	// Visit a parse tree produced by GaussdbParser#collabel.
	VisitCollabel(ctx *CollabelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by GaussdbParser#plsqlidentifier.
	VisitPlsqlidentifier(ctx *PlsqlidentifierContext) interface{}

	// Visit a parse tree produced by GaussdbParser#unreserved_keyword.
	VisitUnreserved_keyword(ctx *Unreserved_keywordContext) interface{}

	// Visit a parse tree produced by GaussdbParser#col_name_keyword.
	VisitCol_name_keyword(ctx *Col_name_keywordContext) interface{}

	// Visit a parse tree produced by GaussdbParser#type_func_name_keyword.
	VisitType_func_name_keyword(ctx *Type_func_name_keywordContext) interface{}

	// Visit a parse tree produced by GaussdbParser#reserved_keyword.
	VisitReserved_keyword(ctx *Reserved_keywordContext) interface{}

	// Visit a parse tree produced by GaussdbParser#builtin_function_name.
	VisitBuiltin_function_name(ctx *Builtin_function_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#pl_function.
	VisitPl_function(ctx *Pl_functionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#comp_options.
	VisitComp_options(ctx *Comp_optionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#comp_option.
	VisitComp_option(ctx *Comp_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#sharp.
	VisitSharp(ctx *SharpContext) interface{}

	// Visit a parse tree produced by GaussdbParser#option_value.
	VisitOption_value(ctx *Option_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_semi.
	VisitOpt_semi(ctx *Opt_semiContext) interface{}

	// Visit a parse tree produced by GaussdbParser#pl_block.
	VisitPl_block(ctx *Pl_blockContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_sect.
	VisitDecl_sect(ctx *Decl_sectContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_start.
	VisitDecl_start(ctx *Decl_startContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_stmts.
	VisitDecl_stmts(ctx *Decl_stmtsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#label_decl.
	VisitLabel_decl(ctx *Label_declContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_stmt.
	VisitDecl_stmt(ctx *Decl_stmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_statement.
	VisitDecl_statement(ctx *Decl_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_scrollable.
	VisitOpt_scrollable(ctx *Opt_scrollableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_cursor_query.
	VisitDecl_cursor_query(ctx *Decl_cursor_queryContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_cursor_args.
	VisitDecl_cursor_args(ctx *Decl_cursor_argsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_cursor_arglist.
	VisitDecl_cursor_arglist(ctx *Decl_cursor_arglistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_cursor_arg.
	VisitDecl_cursor_arg(ctx *Decl_cursor_argContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_is_for.
	VisitDecl_is_for(ctx *Decl_is_forContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_aliasitem.
	VisitDecl_aliasitem(ctx *Decl_aliasitemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_varname.
	VisitDecl_varname(ctx *Decl_varnameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_const.
	VisitDecl_const(ctx *Decl_constContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_datatype.
	VisitDecl_datatype(ctx *Decl_datatypeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_collate.
	VisitDecl_collate(ctx *Decl_collateContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_notnull.
	VisitDecl_notnull(ctx *Decl_notnullContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_defval.
	VisitDecl_defval(ctx *Decl_defvalContext) interface{}

	// Visit a parse tree produced by GaussdbParser#decl_defkey.
	VisitDecl_defkey(ctx *Decl_defkeyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#assign_operator.
	VisitAssign_operator(ctx *Assign_operatorContext) interface{}

	// Visit a parse tree produced by GaussdbParser#proc_sect.
	VisitProc_sect(ctx *Proc_sectContext) interface{}

	// Visit a parse tree produced by GaussdbParser#proc_stmt.
	VisitProc_stmt(ctx *Proc_stmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_perform.
	VisitStmt_perform(ctx *Stmt_performContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_call.
	VisitStmt_call(ctx *Stmt_callContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_expr_list.
	VisitOpt_expr_list(ctx *Opt_expr_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_assign.
	VisitStmt_assign(ctx *Stmt_assignContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_getdiag.
	VisitStmt_getdiag(ctx *Stmt_getdiagContext) interface{}

	// Visit a parse tree produced by GaussdbParser#getdiag_area_opt.
	VisitGetdiag_area_opt(ctx *Getdiag_area_optContext) interface{}

	// Visit a parse tree produced by GaussdbParser#getdiag_list.
	VisitGetdiag_list(ctx *Getdiag_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#getdiag_list_item.
	VisitGetdiag_list_item(ctx *Getdiag_list_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#getdiag_item.
	VisitGetdiag_item(ctx *Getdiag_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#getdiag_target.
	VisitGetdiag_target(ctx *Getdiag_targetContext) interface{}

	// Visit a parse tree produced by GaussdbParser#assign_var.
	VisitAssign_var(ctx *Assign_varContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_if.
	VisitStmt_if(ctx *Stmt_ifContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_elsifs.
	VisitStmt_elsifs(ctx *Stmt_elsifsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_else.
	VisitStmt_else(ctx *Stmt_elseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_case.
	VisitStmt_case(ctx *Stmt_caseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_expr_until_when.
	VisitOpt_expr_until_when(ctx *Opt_expr_until_whenContext) interface{}

	// Visit a parse tree produced by GaussdbParser#case_when_list.
	VisitCase_when_list(ctx *Case_when_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#case_when.
	VisitCase_when(ctx *Case_whenContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_case_else.
	VisitOpt_case_else(ctx *Opt_case_elseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_loop.
	VisitStmt_loop(ctx *Stmt_loopContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_while.
	VisitStmt_while(ctx *Stmt_whileContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_for.
	VisitStmt_for(ctx *Stmt_forContext) interface{}

	// Visit a parse tree produced by GaussdbParser#for_control.
	VisitFor_control(ctx *For_controlContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_for_using_expression.
	VisitOpt_for_using_expression(ctx *Opt_for_using_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_cursor_parameters.
	VisitOpt_cursor_parameters(ctx *Opt_cursor_parametersContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_reverse.
	VisitOpt_reverse(ctx *Opt_reverseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_by_expression.
	VisitOpt_by_expression(ctx *Opt_by_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#for_variable.
	VisitFor_variable(ctx *For_variableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_foreach_a.
	VisitStmt_foreach_a(ctx *Stmt_foreach_aContext) interface{}

	// Visit a parse tree produced by GaussdbParser#foreach_slice.
	VisitForeach_slice(ctx *Foreach_sliceContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_exit.
	VisitStmt_exit(ctx *Stmt_exitContext) interface{}

	// Visit a parse tree produced by GaussdbParser#exit_type.
	VisitExit_type(ctx *Exit_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_return.
	VisitStmt_return(ctx *Stmt_returnContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_return_result.
	VisitOpt_return_result(ctx *Opt_return_resultContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_raise.
	VisitStmt_raise(ctx *Stmt_raiseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_stmt_raise_level.
	VisitOpt_stmt_raise_level(ctx *Opt_stmt_raise_levelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_raise_list.
	VisitOpt_raise_list(ctx *Opt_raise_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_raise_using.
	VisitOpt_raise_using(ctx *Opt_raise_usingContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_raise_using_elem.
	VisitOpt_raise_using_elem(ctx *Opt_raise_using_elemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_raise_using_elem_list.
	VisitOpt_raise_using_elem_list(ctx *Opt_raise_using_elem_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_assert.
	VisitStmt_assert(ctx *Stmt_assertContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_stmt_assert_message.
	VisitOpt_stmt_assert_message(ctx *Opt_stmt_assert_messageContext) interface{}

	// Visit a parse tree produced by GaussdbParser#loop_body.
	VisitLoop_body(ctx *Loop_bodyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_execsql.
	VisitStmt_execsql(ctx *Stmt_execsqlContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_dynexecute.
	VisitStmt_dynexecute(ctx *Stmt_dynexecuteContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_execute_using.
	VisitOpt_execute_using(ctx *Opt_execute_usingContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_execute_using_list.
	VisitOpt_execute_using_list(ctx *Opt_execute_using_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_execute_into.
	VisitOpt_execute_into(ctx *Opt_execute_intoContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_open.
	VisitStmt_open(ctx *Stmt_openContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_open_bound_list_item.
	VisitOpt_open_bound_list_item(ctx *Opt_open_bound_list_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_open_bound_list.
	VisitOpt_open_bound_list(ctx *Opt_open_bound_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_open_using.
	VisitOpt_open_using(ctx *Opt_open_usingContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_scroll_option.
	VisitOpt_scroll_option(ctx *Opt_scroll_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_scroll_option_no.
	VisitOpt_scroll_option_no(ctx *Opt_scroll_option_noContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_fetch.
	VisitStmt_fetch(ctx *Stmt_fetchContext) interface{}

	// Visit a parse tree produced by GaussdbParser#into_target.
	VisitInto_target(ctx *Into_targetContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_cursor_from.
	VisitOpt_cursor_from(ctx *Opt_cursor_fromContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_fetch_direction.
	VisitOpt_fetch_direction(ctx *Opt_fetch_directionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_move.
	VisitStmt_move(ctx *Stmt_moveContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_close.
	VisitStmt_close(ctx *Stmt_closeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_null.
	VisitStmt_null(ctx *Stmt_nullContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_commit.
	VisitStmt_commit(ctx *Stmt_commitContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_rollback.
	VisitStmt_rollback(ctx *Stmt_rollbackContext) interface{}

	// Visit a parse tree produced by GaussdbParser#plsql_opt_transaction_chain.
	VisitPlsql_opt_transaction_chain(ctx *Plsql_opt_transaction_chainContext) interface{}

	// Visit a parse tree produced by GaussdbParser#stmt_set.
	VisitStmt_set(ctx *Stmt_setContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cursor_variable.
	VisitCursor_variable(ctx *Cursor_variableContext) interface{}

	// Visit a parse tree produced by GaussdbParser#exception_sect.
	VisitException_sect(ctx *Exception_sectContext) interface{}

	// Visit a parse tree produced by GaussdbParser#proc_exceptions.
	VisitProc_exceptions(ctx *Proc_exceptionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#proc_exception.
	VisitProc_exception(ctx *Proc_exceptionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#proc_conditions.
	VisitProc_conditions(ctx *Proc_conditionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#proc_condition.
	VisitProc_condition(ctx *Proc_conditionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_block_label.
	VisitOpt_block_label(ctx *Opt_block_labelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_loop_label.
	VisitOpt_loop_label(ctx *Opt_loop_labelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_label.
	VisitOpt_label(ctx *Opt_labelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_exitcond.
	VisitOpt_exitcond(ctx *Opt_exitcondContext) interface{}

	// Visit a parse tree produced by GaussdbParser#any_identifier.
	VisitAny_identifier(ctx *Any_identifierContext) interface{}

	// Visit a parse tree produced by GaussdbParser#plsql_unreserved_keyword.
	VisitPlsql_unreserved_keyword(ctx *Plsql_unreserved_keywordContext) interface{}

	// Visit a parse tree produced by GaussdbParser#sql_expression.
	VisitSql_expression(ctx *Sql_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#expr_until_then.
	VisitExpr_until_then(ctx *Expr_until_thenContext) interface{}

	// Visit a parse tree produced by GaussdbParser#expr_until_semi.
	VisitExpr_until_semi(ctx *Expr_until_semiContext) interface{}

	// Visit a parse tree produced by GaussdbParser#expr_until_rightbracket.
	VisitExpr_until_rightbracket(ctx *Expr_until_rightbracketContext) interface{}

	// Visit a parse tree produced by GaussdbParser#expr_until_loop.
	VisitExpr_until_loop(ctx *Expr_until_loopContext) interface{}

	// Visit a parse tree produced by GaussdbParser#make_execsql_stmt.
	VisitMake_execsql_stmt(ctx *Make_execsql_stmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opt_returning_clause_into.
	VisitOpt_returning_clause_into(ctx *Opt_returning_clause_intoContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_generic_position.
	VisitGaussdb_create_generic_position(ctx *Gaussdb_create_generic_positionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_foreigntable_extend.
	VisitGaussdb_create_foreigntable_extend(ctx *Gaussdb_create_foreigntable_extendContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_error_table_name.
	VisitGaussdb_error_table_name(ctx *Gaussdb_error_table_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_nodenamelist.
	VisitGaussdb_nodenamelist(ctx *Gaussdb_nodenamelistContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_groupname.
	VisitGaussdb_groupname(ctx *Gaussdb_groupnameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_to_group.
	VisitGaussdb_to_group(ctx *Gaussdb_to_groupContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_limit_clause.
	VisitGaussdb_limit_clause(ctx *Gaussdb_limit_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_sort_clause.
	VisitGaussdb_sort_clause(ctx *Gaussdb_sort_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_from_clause.
	VisitGaussdb_from_clause(ctx *Gaussdb_from_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_insert_when_clause.
	VisitGaussdb_insert_when_clause(ctx *Gaussdb_insert_when_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_truncate_table_partition.
	VisitGaussdb_truncate_table_partition(ctx *Gaussdb_truncate_table_partitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_extended_names.
	VisitGaussdb_partition_extended_names(ctx *Gaussdb_partition_extended_namesContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_name.
	VisitGaussdb_partition_name(ctx *Gaussdb_partition_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_new_partition_name.
	VisitGaussdb_new_partition_name(ctx *Gaussdb_new_partition_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_key_value.
	VisitGaussdb_partition_key_value(ctx *Gaussdb_partition_key_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_PURGE.
	VisitGaussdb_PURGE(ctx *Gaussdb_PURGEContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_or_replace.
	VisitGaussdb_or_replace(ctx *Gaussdb_or_replaceContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_index_name.
	VisitGaussdb_index_name(ctx *Gaussdb_index_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_comment_list.
	VisitGaussdb_comment_list(ctx *Gaussdb_comment_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_comment.
	VisitGaussdb_comment(ctx *Gaussdb_commentContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_alter_index_ops_set.
	VisitGaussdb_alter_index_ops_set(ctx *Gaussdb_alter_index_ops_setContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_rebuild_clause.
	VisitGaussdb_rebuild_clause(ctx *Gaussdb_rebuild_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_alter_index_partitioning.
	VisitGaussdb_alter_index_partitioning(ctx *Gaussdb_alter_index_partitioningContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_modify_index_partition.
	VisitGaussdb_modify_index_partition(ctx *Gaussdb_modify_index_partitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_rename_index_partition.
	VisitGaussdb_rename_index_partition(ctx *Gaussdb_rename_index_partitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_move_index_partition.
	VisitGaussdb_move_index_partition(ctx *Gaussdb_move_index_partitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_online.
	VisitGaussdb_online(ctx *Gaussdb_onlineContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_index_partition_name.
	VisitGaussdb_index_partition_name(ctx *Gaussdb_index_partition_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_value.
	VisitGaussdb_partition_value(ctx *Gaussdb_partition_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_index_partition_tablespace.
	VisitGaussdb_index_partition_tablespace(ctx *Gaussdb_index_partition_tablespaceContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_index_partition_list.
	VisitGaussdb_index_partition_list(ctx *Gaussdb_index_partition_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_index_partition.
	VisitGaussdb_index_partition(ctx *Gaussdb_index_partitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_value_list.
	VisitGaussdb_partition_value_list(ctx *Gaussdb_partition_value_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_index_subpartition_list.
	VisitGaussdb_index_subpartition_list(ctx *Gaussdb_index_subpartition_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_index_subppartition.
	VisitGaussdb_index_subppartition(ctx *Gaussdb_index_subppartitionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_connect_clause.
	VisitGaussdb_connect_clause(ctx *Gaussdb_connect_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_engine_list.
	VisitGaussdb_engine_list(ctx *Gaussdb_engine_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_engine_item.
	VisitGaussdb_engine_item(ctx *Gaussdb_engine_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_table_option_list.
	VisitGaussdb_table_option_list(ctx *Gaussdb_table_option_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_table_option.
	VisitGaussdb_table_option(ctx *Gaussdb_table_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_htap_option.
	VisitGaussdb_htap_option(ctx *Gaussdb_htap_optionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_distribute.
	VisitGaussdb_distribute(ctx *Gaussdb_distributeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_slice_less_than_item_list.
	VisitGaussdb_slice_less_than_item_list(ctx *Gaussdb_slice_less_than_item_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_slice_less_than_item.
	VisitGaussdb_slice_less_than_item(ctx *Gaussdb_slice_less_than_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_slice_start_end_item_list.
	VisitGaussdb_slice_start_end_item_list(ctx *Gaussdb_slice_start_end_item_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_slice_start_end_item.
	VisitGaussdb_slice_start_end_item(ctx *Gaussdb_slice_start_end_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_slice_values_item_list.
	VisitGaussdb_slice_values_item_list(ctx *Gaussdb_slice_values_item_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_slice_values_item_item.
	VisitGaussdb_slice_values_item_item(ctx *Gaussdb_slice_values_item_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_ilm_add_policy.
	VisitGaussdb_ilm_add_policy(ctx *Gaussdb_ilm_add_policyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#label_name.
	VisitLabel_name(ctx *Label_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#exception_name.
	VisitException_name(ctx *Exception_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#implementation_type_name.
	VisitImplementation_type_name(ctx *Implementation_type_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#parameter_name.
	VisitParameter_name(ctx *Parameter_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#numeric_negative.
	VisitNumeric_negative(ctx *Numeric_negativeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#plsql_body.
	VisitPlsql_body(ctx *Plsql_bodyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#exception_handler.
	VisitException_handler(ctx *Exception_handlerContext) interface{}

	// Visit a parse tree produced by GaussdbParser#seq_of_statements.
	VisitSeq_of_statements(ctx *Seq_of_statementsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#label_declaration.
	VisitLabel_declaration(ctx *Label_declarationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#seq_of_declare_specs.
	VisitSeq_of_declare_specs(ctx *Seq_of_declare_specsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#declare_spec.
	VisitDeclare_spec(ctx *Declare_specContext) interface{}

	// Visit a parse tree produced by GaussdbParser#pragma_declaration.
	VisitPragma_declaration(ctx *Pragma_declarationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#exception_declaration.
	VisitException_declaration(ctx *Exception_declarationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#default_value_part.
	VisitDefault_value_part(ctx *Default_value_partContext) interface{}

	// Visit a parse tree produced by GaussdbParser#variable_declaration.
	VisitVariable_declaration(ctx *Variable_declarationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#subtype_declaration.
	VisitSubtype_declaration(ctx *Subtype_declarationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cursor_declaration.
	VisitCursor_declaration(ctx *Cursor_declarationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#parameter_spec.
	VisitParameter_spec(ctx *Parameter_specContext) interface{}

	// Visit a parse tree produced by GaussdbParser#ref_cursor_type_def.
	VisitRef_cursor_type_def(ctx *Ref_cursor_type_defContext) interface{}

	// Visit a parse tree produced by GaussdbParser#type_declaration.
	VisitType_declaration(ctx *Type_declarationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#table_type_def.
	VisitTable_type_def(ctx *Table_type_defContext) interface{}

	// Visit a parse tree produced by GaussdbParser#table_indexed_by_part.
	VisitTable_indexed_by_part(ctx *Table_indexed_by_partContext) interface{}

	// Visit a parse tree produced by GaussdbParser#varray_type_def.
	VisitVarray_type_def(ctx *Varray_type_defContext) interface{}

	// Visit a parse tree produced by GaussdbParser#record_type_def.
	VisitRecord_type_def(ctx *Record_type_defContext) interface{}

	// Visit a parse tree produced by GaussdbParser#field_spec.
	VisitField_spec(ctx *Field_specContext) interface{}

	// Visit a parse tree produced by GaussdbParser#procedure_spec.
	VisitProcedure_spec(ctx *Procedure_specContext) interface{}

	// Visit a parse tree produced by GaussdbParser#accessible_by_clause.
	VisitAccessible_by_clause(ctx *Accessible_by_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#accessor.
	VisitAccessor(ctx *AccessorContext) interface{}

	// Visit a parse tree produced by GaussdbParser#function_spec.
	VisitFunction_spec(ctx *Function_specContext) interface{}

	// Visit a parse tree produced by GaussdbParser#parallel_enable_clause.
	VisitParallel_enable_clause(ctx *Parallel_enable_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#partition_by_clause.
	VisitPartition_by_clause(ctx *Partition_by_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#streaming_clause.
	VisitStreaming_clause(ctx *Streaming_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#column_list.
	VisitColumn_list(ctx *Column_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#paren_column_list.
	VisitParen_column_list(ctx *Paren_column_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#procedure_body.
	VisitProcedure_body(ctx *Procedure_bodyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#function_body.
	VisitFunction_body(ctx *Function_bodyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#result_cache_clause.
	VisitResult_cache_clause(ctx *Result_cache_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#relies_on_part.
	VisitRelies_on_part(ctx *Relies_on_partContext) interface{}

	// Visit a parse tree produced by GaussdbParser#invoker_rights_clause.
	VisitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#plsql_statement.
	VisitPlsql_statement(ctx *Plsql_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#sql_statement.
	VisitSql_statement(ctx *Sql_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cursor_manipulation_statements.
	VisitCursor_manipulation_statements(ctx *Cursor_manipulation_statementsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#close_statement.
	VisitClose_statement(ctx *Close_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#open_statement.
	VisitOpen_statement(ctx *Open_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#fetch_statement.
	VisitFetch_statement(ctx *Fetch_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#variable_or_collection.
	VisitVariable_or_collection(ctx *Variable_or_collectionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#collection_expression.
	VisitCollection_expression(ctx *Collection_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#open_for_statement.
	VisitOpen_for_statement(ctx *Open_for_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#collection_method_call.
	VisitCollection_method_call(ctx *Collection_method_callContext) interface{}

	// Visit a parse tree produced by GaussdbParser#execute_immediate.
	VisitExecute_immediate(ctx *Execute_immediateContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dynamic_returning_clause.
	VisitDynamic_returning_clause(ctx *Dynamic_returning_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#raise_statement.
	VisitRaise_statement(ctx *Raise_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#pipe_row_statement.
	VisitPipe_row_statement(ctx *Pipe_row_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#forall_statement.
	VisitForall_statement(ctx *Forall_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#bounds_clause.
	VisitBounds_clause(ctx *Bounds_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#between_bound.
	VisitBetween_bound(ctx *Between_boundContext) interface{}

	// Visit a parse tree produced by GaussdbParser#data_manipulation_language_statements.
	VisitData_manipulation_language_statements(ctx *Data_manipulation_language_statementsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#null_statement.
	VisitNull_statement(ctx *Null_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#goto_statement.
	VisitGoto_statement(ctx *Goto_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#elsif_part.
	VisitElsif_part(ctx *Elsif_partContext) interface{}

	// Visit a parse tree produced by GaussdbParser#else_part.
	VisitElse_part(ctx *Else_partContext) interface{}

	// Visit a parse tree produced by GaussdbParser#loop_statement.
	VisitLoop_statement(ctx *Loop_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cursor_loop_param.
	VisitCursor_loop_param(ctx *Cursor_loop_paramContext) interface{}

	// Visit a parse tree produced by GaussdbParser#lower_bound.
	VisitLower_bound(ctx *Lower_boundContext) interface{}

	// Visit a parse tree produced by GaussdbParser#upper_bound.
	VisitUpper_bound(ctx *Upper_boundContext) interface{}

	// Visit a parse tree produced by GaussdbParser#exit_statement.
	VisitExit_statement(ctx *Exit_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GaussdbParser#assignment_statement.
	VisitAssignment_statement(ctx *Assignment_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#general_element.
	VisitGeneral_element(ctx *General_elementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#cursor_expression.
	VisitCursor_expression(ctx *Cursor_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#logical_expression.
	VisitLogical_expression(ctx *Logical_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#unary_logical_expression.
	VisitUnary_logical_expression(ctx *Unary_logical_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#unary_logical_operation.
	VisitUnary_logical_operation(ctx *Unary_logical_operationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#logical_operation.
	VisitLogical_operation(ctx *Logical_operationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#multiset_expression.
	VisitMultiset_expression(ctx *Multiset_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#concatenation.
	VisitConcatenation(ctx *ConcatenationContext) interface{}

	// Visit a parse tree produced by GaussdbParser#relational_expression.
	VisitRelational_expression(ctx *Relational_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#compound_expression.
	VisitCompound_expression(ctx *Compound_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#relational_operator.
	VisitRelational_operator(ctx *Relational_operatorContext) interface{}

	// Visit a parse tree produced by GaussdbParser#in_elements.
	VisitIn_elements(ctx *In_elementsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#between_elements.
	VisitBetween_elements(ctx *Between_elementsContext) interface{}

	// Visit a parse tree produced by GaussdbParser#interval_expression.
	VisitInterval_expression(ctx *Interval_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#model_expression.
	VisitModel_expression(ctx *Model_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#model_expression_element.
	VisitModel_expression_element(ctx *Model_expression_elementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#single_column_for_loop.
	VisitSingle_column_for_loop(ctx *Single_column_for_loopContext) interface{}

	// Visit a parse tree produced by GaussdbParser#multi_column_for_loop.
	VisitMulti_column_for_loop(ctx *Multi_column_for_loopContext) interface{}

	// Visit a parse tree produced by GaussdbParser#unary_expression.
	VisitUnary_expression(ctx *Unary_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#implicit_cursor_expression.
	VisitImplicit_cursor_expression(ctx *Implicit_cursor_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#case_statement.
	VisitCase_statement(ctx *Case_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#simple_case_statement.
	VisitSimple_case_statement(ctx *Simple_case_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#simple_case_when_part.
	VisitSimple_case_when_part(ctx *Simple_case_when_partContext) interface{}

	// Visit a parse tree produced by GaussdbParser#searched_case_statement.
	VisitSearched_case_statement(ctx *Searched_case_statementContext) interface{}

	// Visit a parse tree produced by GaussdbParser#searched_case_when_part.
	VisitSearched_case_when_part(ctx *Searched_case_when_partContext) interface{}

	// Visit a parse tree produced by GaussdbParser#case_else_part.
	VisitCase_else_part(ctx *Case_else_partContext) interface{}

	// Visit a parse tree produced by GaussdbParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by GaussdbParser#outer_join_sign.
	VisitOuter_join_sign(ctx *Outer_join_signContext) interface{}

	// Visit a parse tree produced by GaussdbParser#quantified_expression.
	VisitQuantified_expression(ctx *Quantified_expressionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#standard_function.
	VisitStandard_function(ctx *Standard_functionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by GaussdbParser#trigger_anonymous_block.
	VisitTrigger_anonymous_block(ctx *Trigger_anonymous_blockContext) interface{}

	// Visit a parse tree produced by GaussdbParser#anonymous_block.
	VisitAnonymous_block(ctx *Anonymous_blockContext) interface{}

	// Visit a parse tree produced by GaussdbParser#package_name.
	VisitPackage_name(ctx *Package_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#create_package.
	VisitCreate_package(ctx *Create_packageContext) interface{}

	// Visit a parse tree produced by GaussdbParser#create_package_body.
	VisitCreate_package_body(ctx *Create_package_bodyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#package_obj_spec.
	VisitPackage_obj_spec(ctx *Package_obj_specContext) interface{}

	// Visit a parse tree produced by GaussdbParser#package_obj_body.
	VisitPackage_obj_body(ctx *Package_obj_bodyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_cleanconnection.
	VisitGaussdb_cleanconnection(ctx *Gaussdb_cleanconnectionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#policy_name.
	VisitPolicy_name(ctx *Policy_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_createpolicystmt.
	VisitGaussdb_createpolicystmt(ctx *Gaussdb_createpolicystmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_audit_clause_list.
	VisitGaussdb_audit_clause_list(ctx *Gaussdb_audit_clause_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#privilege_access_audit_clause.
	VisitPrivilege_access_audit_clause(ctx *Privilege_access_audit_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#ddl_clause.
	VisitDdl_clause(ctx *Ddl_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#resource_label_name.
	VisitResource_label_name(ctx *Resource_label_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#dml_clause.
	VisitDml_clause(ctx *Dml_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#filter_group_clause.
	VisitFilter_group_clause(ctx *Filter_group_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#filter_type_list.
	VisitFilter_type_list(ctx *Filter_type_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#filter_type.
	VisitFilter_type(ctx *Filter_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#filter_value_list.
	VisitFilter_value_list(ctx *Filter_value_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#filter_value.
	VisitFilter_value(ctx *Filter_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_createbarrierstmt.
	VisitGaussdb_createbarrierstmt(ctx *Gaussdb_createbarrierstmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_client_master_key.
	VisitGaussdb_create_client_master_key(ctx *Gaussdb_create_client_master_keyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#client_master_key_name.
	VisitClient_master_key_name(ctx *Client_master_key_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#client_master_key_list.
	VisitClient_master_key_list(ctx *Client_master_key_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#client_master_key.
	VisitClient_master_key(ctx *Client_master_keyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_column_encryption_key.
	VisitGaussdb_create_column_encryption_key(ctx *Gaussdb_create_column_encryption_keyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#column_encryption_key_name.
	VisitColumn_encryption_key_name(ctx *Column_encryption_key_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#column_encryption_key_list.
	VisitColumn_encryption_key_list(ctx *Column_encryption_key_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#column_encryption_key.
	VisitColumn_encryption_key(ctx *Column_encryption_keyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_database_link.
	VisitGaussdb_create_database_link(ctx *Gaussdb_create_database_linkContext) interface{}

	// Visit a parse tree produced by GaussdbParser#user_object_name.
	VisitUser_object_name(ctx *User_object_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#password_value.
	VisitPassword_value(ctx *Password_valueContext) interface{}

	// Visit a parse tree produced by GaussdbParser#database_link_using.
	VisitDatabase_link_using(ctx *Database_link_usingContext) interface{}

	// Visit a parse tree produced by GaussdbParser#database_link_using_opt.
	VisitDatabase_link_using_opt(ctx *Database_link_using_optContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_directory.
	VisitGaussdb_create_directory(ctx *Gaussdb_create_directoryContext) interface{}

	// Visit a parse tree produced by GaussdbParser#directory_name.
	VisitDirectory_name(ctx *Directory_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_masking_policy.
	VisitGaussdb_create_masking_policy(ctx *Gaussdb_create_masking_policyContext) interface{}

	// Visit a parse tree produced by GaussdbParser#masking_clause_list.
	VisitMasking_clause_list(ctx *Masking_clause_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#masking_clause.
	VisitMasking_clause(ctx *Masking_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#masking_function.
	VisitMasking_function(ctx *Masking_functionContext) interface{}

	// Visit a parse tree produced by GaussdbParser#policy_filter_clause.
	VisitPolicy_filter_clause(ctx *Policy_filter_clauseContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_node.
	VisitGaussdb_create_node(ctx *Gaussdb_create_nodeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_nodename.
	VisitGaussdb_nodename(ctx *Gaussdb_nodenameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_node_group.
	VisitGaussdb_create_node_group(ctx *Gaussdb_create_node_groupContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_bucketnumber.
	VisitGaussdb_bucketnumber(ctx *Gaussdb_bucketnumberContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_resource_pool.
	VisitGaussdb_create_resource_pool(ctx *Gaussdb_create_resource_poolContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_resource_pool_name.
	VisitGaussdb_resource_pool_name(ctx *Gaussdb_resource_pool_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_resource_label.
	VisitGaussdb_create_resource_label(ctx *Gaussdb_create_resource_labelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#label_item_list.
	VisitLabel_item_list(ctx *Label_item_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#label_item.
	VisitLabel_item(ctx *Label_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#resource_path.
	VisitResource_path(ctx *Resource_pathContext) interface{}

	// Visit a parse tree produced by GaussdbParser#resource_type.
	VisitResource_type(ctx *Resource_typeContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_security_label.
	VisitGaussdb_create_security_label(ctx *Gaussdb_create_security_labelContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_synonym.
	VisitGaussdb_create_synonym(ctx *Gaussdb_create_synonymContext) interface{}

	// Visit a parse tree produced by GaussdbParser#synonym_name.
	VisitSynonym_name(ctx *Synonym_nameContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_less_than_item_list.
	VisitGaussdb_partition_less_than_item_list(ctx *Gaussdb_partition_less_than_item_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_less_than_item.
	VisitGaussdb_partition_less_than_item(ctx *Gaussdb_partition_less_than_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_start_end_item_list.
	VisitGaussdb_partition_start_end_item_list(ctx *Gaussdb_partition_start_end_item_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_start_end_item.
	VisitGaussdb_partition_start_end_item(ctx *Gaussdb_partition_start_end_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#opttablespace_list.
	VisitOpttablespace_list(ctx *Opttablespace_listContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_hash_item.
	VisitGaussdb_partition_hash_item(ctx *Gaussdb_partition_hash_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_partition_list_item.
	VisitGaussdb_partition_list_item(ctx *Gaussdb_partition_list_itemContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_create_weak_password_dic.
	VisitGaussdb_create_weak_password_dic(ctx *Gaussdb_create_weak_password_dicContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_purge_stmt.
	VisitGaussdb_purge_stmt(ctx *Gaussdb_purge_stmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_replacestmt.
	VisitGaussdb_replacestmt(ctx *Gaussdb_replacestmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_shutdown_stmt.
	VisitGaussdb_shutdown_stmt(ctx *Gaussdb_shutdown_stmtContext) interface{}

	// Visit a parse tree produced by GaussdbParser#gaussdb_show_events_stmt.
	VisitGaussdb_show_events_stmt(ctx *Gaussdb_show_events_stmtContext) interface{}
}
