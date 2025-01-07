// Code generated from GaussdbParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // GaussdbParser
import "github.com/antlr4-go/antlr/v4"

// BaseGaussdbParserListener is a complete listener for a parse tree produced by GaussdbParser.
type BaseGaussdbParserListener struct{}

var _ GaussdbParserListener = &BaseGaussdbParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGaussdbParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGaussdbParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGaussdbParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGaussdbParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseGaussdbParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseGaussdbParserListener) ExitRoot(ctx *RootContext) {}

// EnterPlsqlroot is called when production plsqlroot is entered.
func (s *BaseGaussdbParserListener) EnterPlsqlroot(ctx *PlsqlrootContext) {}

// ExitPlsqlroot is called when production plsqlroot is exited.
func (s *BaseGaussdbParserListener) ExitPlsqlroot(ctx *PlsqlrootContext) {}

// EnterStmtblock is called when production stmtblock is entered.
func (s *BaseGaussdbParserListener) EnterStmtblock(ctx *StmtblockContext) {}

// ExitStmtblock is called when production stmtblock is exited.
func (s *BaseGaussdbParserListener) ExitStmtblock(ctx *StmtblockContext) {}

// EnterStmtmulti is called when production stmtmulti is entered.
func (s *BaseGaussdbParserListener) EnterStmtmulti(ctx *StmtmultiContext) {}

// ExitStmtmulti is called when production stmtmulti is exited.
func (s *BaseGaussdbParserListener) ExitStmtmulti(ctx *StmtmultiContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseGaussdbParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseGaussdbParserListener) ExitStmt(ctx *StmtContext) {}

// EnterPlsqlconsolecommand is called when production plsqlconsolecommand is entered.
func (s *BaseGaussdbParserListener) EnterPlsqlconsolecommand(ctx *PlsqlconsolecommandContext) {}

// ExitPlsqlconsolecommand is called when production plsqlconsolecommand is exited.
func (s *BaseGaussdbParserListener) ExitPlsqlconsolecommand(ctx *PlsqlconsolecommandContext) {}

// EnterCallstmt is called when production callstmt is entered.
func (s *BaseGaussdbParserListener) EnterCallstmt(ctx *CallstmtContext) {}

// ExitCallstmt is called when production callstmt is exited.
func (s *BaseGaussdbParserListener) ExitCallstmt(ctx *CallstmtContext) {}

// EnterCreaterolestmt is called when production createrolestmt is entered.
func (s *BaseGaussdbParserListener) EnterCreaterolestmt(ctx *CreaterolestmtContext) {}

// ExitCreaterolestmt is called when production createrolestmt is exited.
func (s *BaseGaussdbParserListener) ExitCreaterolestmt(ctx *CreaterolestmtContext) {}

// EnterOpt_with is called when production opt_with is entered.
func (s *BaseGaussdbParserListener) EnterOpt_with(ctx *Opt_withContext) {}

// ExitOpt_with is called when production opt_with is exited.
func (s *BaseGaussdbParserListener) ExitOpt_with(ctx *Opt_withContext) {}

// EnterOptrolelist is called when production optrolelist is entered.
func (s *BaseGaussdbParserListener) EnterOptrolelist(ctx *OptrolelistContext) {}

// ExitOptrolelist is called when production optrolelist is exited.
func (s *BaseGaussdbParserListener) ExitOptrolelist(ctx *OptrolelistContext) {}

// EnterAlteroptrolelist is called when production alteroptrolelist is entered.
func (s *BaseGaussdbParserListener) EnterAlteroptrolelist(ctx *AlteroptrolelistContext) {}

// ExitAlteroptrolelist is called when production alteroptrolelist is exited.
func (s *BaseGaussdbParserListener) ExitAlteroptrolelist(ctx *AlteroptrolelistContext) {}

// EnterAlteroptroleelem is called when production alteroptroleelem is entered.
func (s *BaseGaussdbParserListener) EnterAlteroptroleelem(ctx *AlteroptroleelemContext) {}

// ExitAlteroptroleelem is called when production alteroptroleelem is exited.
func (s *BaseGaussdbParserListener) ExitAlteroptroleelem(ctx *AlteroptroleelemContext) {}

// EnterCreateoptroleelem is called when production createoptroleelem is entered.
func (s *BaseGaussdbParserListener) EnterCreateoptroleelem(ctx *CreateoptroleelemContext) {}

// ExitCreateoptroleelem is called when production createoptroleelem is exited.
func (s *BaseGaussdbParserListener) ExitCreateoptroleelem(ctx *CreateoptroleelemContext) {}

// EnterCreateuserstmt is called when production createuserstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateuserstmt(ctx *CreateuserstmtContext) {}

// ExitCreateuserstmt is called when production createuserstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateuserstmt(ctx *CreateuserstmtContext) {}

// EnterAlterrolestmt is called when production alterrolestmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterrolestmt(ctx *AlterrolestmtContext) {}

// ExitAlterrolestmt is called when production alterrolestmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterrolestmt(ctx *AlterrolestmtContext) {}

// EnterOpt_in_database is called when production opt_in_database is entered.
func (s *BaseGaussdbParserListener) EnterOpt_in_database(ctx *Opt_in_databaseContext) {}

// ExitOpt_in_database is called when production opt_in_database is exited.
func (s *BaseGaussdbParserListener) ExitOpt_in_database(ctx *Opt_in_databaseContext) {}

// EnterAlterrolesetstmt is called when production alterrolesetstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterrolesetstmt(ctx *AlterrolesetstmtContext) {}

// ExitAlterrolesetstmt is called when production alterrolesetstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterrolesetstmt(ctx *AlterrolesetstmtContext) {}

// EnterDroprolestmt is called when production droprolestmt is entered.
func (s *BaseGaussdbParserListener) EnterDroprolestmt(ctx *DroprolestmtContext) {}

// ExitDroprolestmt is called when production droprolestmt is exited.
func (s *BaseGaussdbParserListener) ExitDroprolestmt(ctx *DroprolestmtContext) {}

// EnterCreategroupstmt is called when production creategroupstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreategroupstmt(ctx *CreategroupstmtContext) {}

// ExitCreategroupstmt is called when production creategroupstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreategroupstmt(ctx *CreategroupstmtContext) {}

// EnterAltergroupstmt is called when production altergroupstmt is entered.
func (s *BaseGaussdbParserListener) EnterAltergroupstmt(ctx *AltergroupstmtContext) {}

// ExitAltergroupstmt is called when production altergroupstmt is exited.
func (s *BaseGaussdbParserListener) ExitAltergroupstmt(ctx *AltergroupstmtContext) {}

// EnterAdd_drop is called when production add_drop is entered.
func (s *BaseGaussdbParserListener) EnterAdd_drop(ctx *Add_dropContext) {}

// ExitAdd_drop is called when production add_drop is exited.
func (s *BaseGaussdbParserListener) ExitAdd_drop(ctx *Add_dropContext) {}

// EnterCreateschemastmt is called when production createschemastmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateschemastmt(ctx *CreateschemastmtContext) {}

// ExitCreateschemastmt is called when production createschemastmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateschemastmt(ctx *CreateschemastmtContext) {}

// EnterOptschemaname is called when production optschemaname is entered.
func (s *BaseGaussdbParserListener) EnterOptschemaname(ctx *OptschemanameContext) {}

// ExitOptschemaname is called when production optschemaname is exited.
func (s *BaseGaussdbParserListener) ExitOptschemaname(ctx *OptschemanameContext) {}

// EnterOptschemaeltlist is called when production optschemaeltlist is entered.
func (s *BaseGaussdbParserListener) EnterOptschemaeltlist(ctx *OptschemaeltlistContext) {}

// ExitOptschemaeltlist is called when production optschemaeltlist is exited.
func (s *BaseGaussdbParserListener) ExitOptschemaeltlist(ctx *OptschemaeltlistContext) {}

// EnterSchema_stmt is called when production schema_stmt is entered.
func (s *BaseGaussdbParserListener) EnterSchema_stmt(ctx *Schema_stmtContext) {}

// ExitSchema_stmt is called when production schema_stmt is exited.
func (s *BaseGaussdbParserListener) ExitSchema_stmt(ctx *Schema_stmtContext) {}

// EnterVariablesetstmt is called when production variablesetstmt is entered.
func (s *BaseGaussdbParserListener) EnterVariablesetstmt(ctx *VariablesetstmtContext) {}

// ExitVariablesetstmt is called when production variablesetstmt is exited.
func (s *BaseGaussdbParserListener) ExitVariablesetstmt(ctx *VariablesetstmtContext) {}

// EnterSet_rest is called when production set_rest is entered.
func (s *BaseGaussdbParserListener) EnterSet_rest(ctx *Set_restContext) {}

// ExitSet_rest is called when production set_rest is exited.
func (s *BaseGaussdbParserListener) ExitSet_rest(ctx *Set_restContext) {}

// EnterGeneric_set is called when production generic_set is entered.
func (s *BaseGaussdbParserListener) EnterGeneric_set(ctx *Generic_setContext) {}

// ExitGeneric_set is called when production generic_set is exited.
func (s *BaseGaussdbParserListener) ExitGeneric_set(ctx *Generic_setContext) {}

// EnterSet_rest_more is called when production set_rest_more is entered.
func (s *BaseGaussdbParserListener) EnterSet_rest_more(ctx *Set_rest_moreContext) {}

// ExitSet_rest_more is called when production set_rest_more is exited.
func (s *BaseGaussdbParserListener) ExitSet_rest_more(ctx *Set_rest_moreContext) {}

// EnterVar_name is called when production var_name is entered.
func (s *BaseGaussdbParserListener) EnterVar_name(ctx *Var_nameContext) {}

// ExitVar_name is called when production var_name is exited.
func (s *BaseGaussdbParserListener) ExitVar_name(ctx *Var_nameContext) {}

// EnterVar_list is called when production var_list is entered.
func (s *BaseGaussdbParserListener) EnterVar_list(ctx *Var_listContext) {}

// ExitVar_list is called when production var_list is exited.
func (s *BaseGaussdbParserListener) ExitVar_list(ctx *Var_listContext) {}

// EnterVar_value is called when production var_value is entered.
func (s *BaseGaussdbParserListener) EnterVar_value(ctx *Var_valueContext) {}

// ExitVar_value is called when production var_value is exited.
func (s *BaseGaussdbParserListener) ExitVar_value(ctx *Var_valueContext) {}

// EnterIso_level is called when production iso_level is entered.
func (s *BaseGaussdbParserListener) EnterIso_level(ctx *Iso_levelContext) {}

// ExitIso_level is called when production iso_level is exited.
func (s *BaseGaussdbParserListener) ExitIso_level(ctx *Iso_levelContext) {}

// EnterOpt_boolean_or_string is called when production opt_boolean_or_string is entered.
func (s *BaseGaussdbParserListener) EnterOpt_boolean_or_string(ctx *Opt_boolean_or_stringContext) {}

// ExitOpt_boolean_or_string is called when production opt_boolean_or_string is exited.
func (s *BaseGaussdbParserListener) ExitOpt_boolean_or_string(ctx *Opt_boolean_or_stringContext) {}

// EnterZone_value is called when production zone_value is entered.
func (s *BaseGaussdbParserListener) EnterZone_value(ctx *Zone_valueContext) {}

// ExitZone_value is called when production zone_value is exited.
func (s *BaseGaussdbParserListener) ExitZone_value(ctx *Zone_valueContext) {}

// EnterOpt_encoding is called when production opt_encoding is entered.
func (s *BaseGaussdbParserListener) EnterOpt_encoding(ctx *Opt_encodingContext) {}

// ExitOpt_encoding is called when production opt_encoding is exited.
func (s *BaseGaussdbParserListener) ExitOpt_encoding(ctx *Opt_encodingContext) {}

// EnterNonreservedword_or_sconst is called when production nonreservedword_or_sconst is entered.
func (s *BaseGaussdbParserListener) EnterNonreservedword_or_sconst(ctx *Nonreservedword_or_sconstContext) {
}

// ExitNonreservedword_or_sconst is called when production nonreservedword_or_sconst is exited.
func (s *BaseGaussdbParserListener) ExitNonreservedword_or_sconst(ctx *Nonreservedword_or_sconstContext) {
}

// EnterVariableresetstmt is called when production variableresetstmt is entered.
func (s *BaseGaussdbParserListener) EnterVariableresetstmt(ctx *VariableresetstmtContext) {}

// ExitVariableresetstmt is called when production variableresetstmt is exited.
func (s *BaseGaussdbParserListener) ExitVariableresetstmt(ctx *VariableresetstmtContext) {}

// EnterReset_rest is called when production reset_rest is entered.
func (s *BaseGaussdbParserListener) EnterReset_rest(ctx *Reset_restContext) {}

// ExitReset_rest is called when production reset_rest is exited.
func (s *BaseGaussdbParserListener) ExitReset_rest(ctx *Reset_restContext) {}

// EnterGeneric_reset is called when production generic_reset is entered.
func (s *BaseGaussdbParserListener) EnterGeneric_reset(ctx *Generic_resetContext) {}

// ExitGeneric_reset is called when production generic_reset is exited.
func (s *BaseGaussdbParserListener) ExitGeneric_reset(ctx *Generic_resetContext) {}

// EnterSetresetclause is called when production setresetclause is entered.
func (s *BaseGaussdbParserListener) EnterSetresetclause(ctx *SetresetclauseContext) {}

// ExitSetresetclause is called when production setresetclause is exited.
func (s *BaseGaussdbParserListener) ExitSetresetclause(ctx *SetresetclauseContext) {}

// EnterFunctionsetresetclause is called when production functionsetresetclause is entered.
func (s *BaseGaussdbParserListener) EnterFunctionsetresetclause(ctx *FunctionsetresetclauseContext) {}

// ExitFunctionsetresetclause is called when production functionsetresetclause is exited.
func (s *BaseGaussdbParserListener) ExitFunctionsetresetclause(ctx *FunctionsetresetclauseContext) {}

// EnterVariableshowstmt is called when production variableshowstmt is entered.
func (s *BaseGaussdbParserListener) EnterVariableshowstmt(ctx *VariableshowstmtContext) {}

// ExitVariableshowstmt is called when production variableshowstmt is exited.
func (s *BaseGaussdbParserListener) ExitVariableshowstmt(ctx *VariableshowstmtContext) {}

// EnterConstraintssetstmt is called when production constraintssetstmt is entered.
func (s *BaseGaussdbParserListener) EnterConstraintssetstmt(ctx *ConstraintssetstmtContext) {}

// ExitConstraintssetstmt is called when production constraintssetstmt is exited.
func (s *BaseGaussdbParserListener) ExitConstraintssetstmt(ctx *ConstraintssetstmtContext) {}

// EnterConstraints_set_list is called when production constraints_set_list is entered.
func (s *BaseGaussdbParserListener) EnterConstraints_set_list(ctx *Constraints_set_listContext) {}

// ExitConstraints_set_list is called when production constraints_set_list is exited.
func (s *BaseGaussdbParserListener) ExitConstraints_set_list(ctx *Constraints_set_listContext) {}

// EnterConstraints_set_mode is called when production constraints_set_mode is entered.
func (s *BaseGaussdbParserListener) EnterConstraints_set_mode(ctx *Constraints_set_modeContext) {}

// ExitConstraints_set_mode is called when production constraints_set_mode is exited.
func (s *BaseGaussdbParserListener) ExitConstraints_set_mode(ctx *Constraints_set_modeContext) {}

// EnterCheckpointstmt is called when production checkpointstmt is entered.
func (s *BaseGaussdbParserListener) EnterCheckpointstmt(ctx *CheckpointstmtContext) {}

// ExitCheckpointstmt is called when production checkpointstmt is exited.
func (s *BaseGaussdbParserListener) ExitCheckpointstmt(ctx *CheckpointstmtContext) {}

// EnterDiscardstmt is called when production discardstmt is entered.
func (s *BaseGaussdbParserListener) EnterDiscardstmt(ctx *DiscardstmtContext) {}

// ExitDiscardstmt is called when production discardstmt is exited.
func (s *BaseGaussdbParserListener) ExitDiscardstmt(ctx *DiscardstmtContext) {}

// EnterAltertablestmt is called when production altertablestmt is entered.
func (s *BaseGaussdbParserListener) EnterAltertablestmt(ctx *AltertablestmtContext) {}

// ExitAltertablestmt is called when production altertablestmt is exited.
func (s *BaseGaussdbParserListener) ExitAltertablestmt(ctx *AltertablestmtContext) {}

// EnterAlter_table_cmds is called when production alter_table_cmds is entered.
func (s *BaseGaussdbParserListener) EnterAlter_table_cmds(ctx *Alter_table_cmdsContext) {}

// ExitAlter_table_cmds is called when production alter_table_cmds is exited.
func (s *BaseGaussdbParserListener) ExitAlter_table_cmds(ctx *Alter_table_cmdsContext) {}

// EnterPartition_cmd is called when production partition_cmd is entered.
func (s *BaseGaussdbParserListener) EnterPartition_cmd(ctx *Partition_cmdContext) {}

// ExitPartition_cmd is called when production partition_cmd is exited.
func (s *BaseGaussdbParserListener) ExitPartition_cmd(ctx *Partition_cmdContext) {}

// EnterIndex_partition_cmd is called when production index_partition_cmd is entered.
func (s *BaseGaussdbParserListener) EnterIndex_partition_cmd(ctx *Index_partition_cmdContext) {}

// ExitIndex_partition_cmd is called when production index_partition_cmd is exited.
func (s *BaseGaussdbParserListener) ExitIndex_partition_cmd(ctx *Index_partition_cmdContext) {}

// EnterAlter_table_cmd is called when production alter_table_cmd is entered.
func (s *BaseGaussdbParserListener) EnterAlter_table_cmd(ctx *Alter_table_cmdContext) {}

// ExitAlter_table_cmd is called when production alter_table_cmd is exited.
func (s *BaseGaussdbParserListener) ExitAlter_table_cmd(ctx *Alter_table_cmdContext) {}

// EnterAlter_column_default is called when production alter_column_default is entered.
func (s *BaseGaussdbParserListener) EnterAlter_column_default(ctx *Alter_column_defaultContext) {}

// ExitAlter_column_default is called when production alter_column_default is exited.
func (s *BaseGaussdbParserListener) ExitAlter_column_default(ctx *Alter_column_defaultContext) {}

// EnterOpt_drop_behavior is called when production opt_drop_behavior is entered.
func (s *BaseGaussdbParserListener) EnterOpt_drop_behavior(ctx *Opt_drop_behaviorContext) {}

// ExitOpt_drop_behavior is called when production opt_drop_behavior is exited.
func (s *BaseGaussdbParserListener) ExitOpt_drop_behavior(ctx *Opt_drop_behaviorContext) {}

// EnterOpt_collate_clause is called when production opt_collate_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_collate_clause(ctx *Opt_collate_clauseContext) {}

// ExitOpt_collate_clause is called when production opt_collate_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_collate_clause(ctx *Opt_collate_clauseContext) {}

// EnterAlter_using is called when production alter_using is entered.
func (s *BaseGaussdbParserListener) EnterAlter_using(ctx *Alter_usingContext) {}

// ExitAlter_using is called when production alter_using is exited.
func (s *BaseGaussdbParserListener) ExitAlter_using(ctx *Alter_usingContext) {}

// EnterReplica_identity is called when production replica_identity is entered.
func (s *BaseGaussdbParserListener) EnterReplica_identity(ctx *Replica_identityContext) {}

// ExitReplica_identity is called when production replica_identity is exited.
func (s *BaseGaussdbParserListener) ExitReplica_identity(ctx *Replica_identityContext) {}

// EnterReloptions is called when production reloptions is entered.
func (s *BaseGaussdbParserListener) EnterReloptions(ctx *ReloptionsContext) {}

// ExitReloptions is called when production reloptions is exited.
func (s *BaseGaussdbParserListener) ExitReloptions(ctx *ReloptionsContext) {}

// EnterOpt_reloptions is called when production opt_reloptions is entered.
func (s *BaseGaussdbParserListener) EnterOpt_reloptions(ctx *Opt_reloptionsContext) {}

// ExitOpt_reloptions is called when production opt_reloptions is exited.
func (s *BaseGaussdbParserListener) ExitOpt_reloptions(ctx *Opt_reloptionsContext) {}

// EnterReloption_list is called when production reloption_list is entered.
func (s *BaseGaussdbParserListener) EnterReloption_list(ctx *Reloption_listContext) {}

// ExitReloption_list is called when production reloption_list is exited.
func (s *BaseGaussdbParserListener) ExitReloption_list(ctx *Reloption_listContext) {}

// EnterReloption_elem is called when production reloption_elem is entered.
func (s *BaseGaussdbParserListener) EnterReloption_elem(ctx *Reloption_elemContext) {}

// ExitReloption_elem is called when production reloption_elem is exited.
func (s *BaseGaussdbParserListener) ExitReloption_elem(ctx *Reloption_elemContext) {}

// EnterAlter_identity_column_option_list is called when production alter_identity_column_option_list is entered.
func (s *BaseGaussdbParserListener) EnterAlter_identity_column_option_list(ctx *Alter_identity_column_option_listContext) {
}

// ExitAlter_identity_column_option_list is called when production alter_identity_column_option_list is exited.
func (s *BaseGaussdbParserListener) ExitAlter_identity_column_option_list(ctx *Alter_identity_column_option_listContext) {
}

// EnterAlter_identity_column_option is called when production alter_identity_column_option is entered.
func (s *BaseGaussdbParserListener) EnterAlter_identity_column_option(ctx *Alter_identity_column_optionContext) {
}

// ExitAlter_identity_column_option is called when production alter_identity_column_option is exited.
func (s *BaseGaussdbParserListener) ExitAlter_identity_column_option(ctx *Alter_identity_column_optionContext) {
}

// EnterPartitionboundspec is called when production partitionboundspec is entered.
func (s *BaseGaussdbParserListener) EnterPartitionboundspec(ctx *PartitionboundspecContext) {}

// ExitPartitionboundspec is called when production partitionboundspec is exited.
func (s *BaseGaussdbParserListener) ExitPartitionboundspec(ctx *PartitionboundspecContext) {}

// EnterHash_partbound_elem is called when production hash_partbound_elem is entered.
func (s *BaseGaussdbParserListener) EnterHash_partbound_elem(ctx *Hash_partbound_elemContext) {}

// ExitHash_partbound_elem is called when production hash_partbound_elem is exited.
func (s *BaseGaussdbParserListener) ExitHash_partbound_elem(ctx *Hash_partbound_elemContext) {}

// EnterHash_partbound is called when production hash_partbound is entered.
func (s *BaseGaussdbParserListener) EnterHash_partbound(ctx *Hash_partboundContext) {}

// ExitHash_partbound is called when production hash_partbound is exited.
func (s *BaseGaussdbParserListener) ExitHash_partbound(ctx *Hash_partboundContext) {}

// EnterAltercompositetypestmt is called when production altercompositetypestmt is entered.
func (s *BaseGaussdbParserListener) EnterAltercompositetypestmt(ctx *AltercompositetypestmtContext) {}

// ExitAltercompositetypestmt is called when production altercompositetypestmt is exited.
func (s *BaseGaussdbParserListener) ExitAltercompositetypestmt(ctx *AltercompositetypestmtContext) {}

// EnterAlter_type_cmds is called when production alter_type_cmds is entered.
func (s *BaseGaussdbParserListener) EnterAlter_type_cmds(ctx *Alter_type_cmdsContext) {}

// ExitAlter_type_cmds is called when production alter_type_cmds is exited.
func (s *BaseGaussdbParserListener) ExitAlter_type_cmds(ctx *Alter_type_cmdsContext) {}

// EnterAlter_type_cmd is called when production alter_type_cmd is entered.
func (s *BaseGaussdbParserListener) EnterAlter_type_cmd(ctx *Alter_type_cmdContext) {}

// ExitAlter_type_cmd is called when production alter_type_cmd is exited.
func (s *BaseGaussdbParserListener) ExitAlter_type_cmd(ctx *Alter_type_cmdContext) {}

// EnterCloseportalstmt is called when production closeportalstmt is entered.
func (s *BaseGaussdbParserListener) EnterCloseportalstmt(ctx *CloseportalstmtContext) {}

// ExitCloseportalstmt is called when production closeportalstmt is exited.
func (s *BaseGaussdbParserListener) ExitCloseportalstmt(ctx *CloseportalstmtContext) {}

// EnterCopystmt is called when production copystmt is entered.
func (s *BaseGaussdbParserListener) EnterCopystmt(ctx *CopystmtContext) {}

// ExitCopystmt is called when production copystmt is exited.
func (s *BaseGaussdbParserListener) ExitCopystmt(ctx *CopystmtContext) {}

// EnterCopy_from is called when production copy_from is entered.
func (s *BaseGaussdbParserListener) EnterCopy_from(ctx *Copy_fromContext) {}

// ExitCopy_from is called when production copy_from is exited.
func (s *BaseGaussdbParserListener) ExitCopy_from(ctx *Copy_fromContext) {}

// EnterOpt_program is called when production opt_program is entered.
func (s *BaseGaussdbParserListener) EnterOpt_program(ctx *Opt_programContext) {}

// ExitOpt_program is called when production opt_program is exited.
func (s *BaseGaussdbParserListener) ExitOpt_program(ctx *Opt_programContext) {}

// EnterCopy_file_name is called when production copy_file_name is entered.
func (s *BaseGaussdbParserListener) EnterCopy_file_name(ctx *Copy_file_nameContext) {}

// ExitCopy_file_name is called when production copy_file_name is exited.
func (s *BaseGaussdbParserListener) ExitCopy_file_name(ctx *Copy_file_nameContext) {}

// EnterCopy_options is called when production copy_options is entered.
func (s *BaseGaussdbParserListener) EnterCopy_options(ctx *Copy_optionsContext) {}

// ExitCopy_options is called when production copy_options is exited.
func (s *BaseGaussdbParserListener) ExitCopy_options(ctx *Copy_optionsContext) {}

// EnterCopy_opt_list is called when production copy_opt_list is entered.
func (s *BaseGaussdbParserListener) EnterCopy_opt_list(ctx *Copy_opt_listContext) {}

// ExitCopy_opt_list is called when production copy_opt_list is exited.
func (s *BaseGaussdbParserListener) ExitCopy_opt_list(ctx *Copy_opt_listContext) {}

// EnterCopy_opt_item is called when production copy_opt_item is entered.
func (s *BaseGaussdbParserListener) EnterCopy_opt_item(ctx *Copy_opt_itemContext) {}

// ExitCopy_opt_item is called when production copy_opt_item is exited.
func (s *BaseGaussdbParserListener) ExitCopy_opt_item(ctx *Copy_opt_itemContext) {}

// EnterOpt_binary is called when production opt_binary is entered.
func (s *BaseGaussdbParserListener) EnterOpt_binary(ctx *Opt_binaryContext) {}

// ExitOpt_binary is called when production opt_binary is exited.
func (s *BaseGaussdbParserListener) ExitOpt_binary(ctx *Opt_binaryContext) {}

// EnterCopy_delimiter is called when production copy_delimiter is entered.
func (s *BaseGaussdbParserListener) EnterCopy_delimiter(ctx *Copy_delimiterContext) {}

// ExitCopy_delimiter is called when production copy_delimiter is exited.
func (s *BaseGaussdbParserListener) ExitCopy_delimiter(ctx *Copy_delimiterContext) {}

// EnterOpt_using is called when production opt_using is entered.
func (s *BaseGaussdbParserListener) EnterOpt_using(ctx *Opt_usingContext) {}

// ExitOpt_using is called when production opt_using is exited.
func (s *BaseGaussdbParserListener) ExitOpt_using(ctx *Opt_usingContext) {}

// EnterCopy_generic_opt_list is called when production copy_generic_opt_list is entered.
func (s *BaseGaussdbParserListener) EnterCopy_generic_opt_list(ctx *Copy_generic_opt_listContext) {}

// ExitCopy_generic_opt_list is called when production copy_generic_opt_list is exited.
func (s *BaseGaussdbParserListener) ExitCopy_generic_opt_list(ctx *Copy_generic_opt_listContext) {}

// EnterCopy_generic_opt_elem is called when production copy_generic_opt_elem is entered.
func (s *BaseGaussdbParserListener) EnterCopy_generic_opt_elem(ctx *Copy_generic_opt_elemContext) {}

// ExitCopy_generic_opt_elem is called when production copy_generic_opt_elem is exited.
func (s *BaseGaussdbParserListener) ExitCopy_generic_opt_elem(ctx *Copy_generic_opt_elemContext) {}

// EnterCopy_generic_opt_arg is called when production copy_generic_opt_arg is entered.
func (s *BaseGaussdbParserListener) EnterCopy_generic_opt_arg(ctx *Copy_generic_opt_argContext) {}

// ExitCopy_generic_opt_arg is called when production copy_generic_opt_arg is exited.
func (s *BaseGaussdbParserListener) ExitCopy_generic_opt_arg(ctx *Copy_generic_opt_argContext) {}

// EnterCopy_generic_opt_arg_list is called when production copy_generic_opt_arg_list is entered.
func (s *BaseGaussdbParserListener) EnterCopy_generic_opt_arg_list(ctx *Copy_generic_opt_arg_listContext) {
}

// ExitCopy_generic_opt_arg_list is called when production copy_generic_opt_arg_list is exited.
func (s *BaseGaussdbParserListener) ExitCopy_generic_opt_arg_list(ctx *Copy_generic_opt_arg_listContext) {
}

// EnterCopy_generic_opt_arg_list_item is called when production copy_generic_opt_arg_list_item is entered.
func (s *BaseGaussdbParserListener) EnterCopy_generic_opt_arg_list_item(ctx *Copy_generic_opt_arg_list_itemContext) {
}

// ExitCopy_generic_opt_arg_list_item is called when production copy_generic_opt_arg_list_item is exited.
func (s *BaseGaussdbParserListener) ExitCopy_generic_opt_arg_list_item(ctx *Copy_generic_opt_arg_list_itemContext) {
}

// EnterCreatestmt is called when production createstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatestmt(ctx *CreatestmtContext) {}

// ExitCreatestmt is called when production createstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatestmt(ctx *CreatestmtContext) {}

// EnterOpttemp is called when production opttemp is entered.
func (s *BaseGaussdbParserListener) EnterOpttemp(ctx *OpttempContext) {}

// ExitOpttemp is called when production opttemp is exited.
func (s *BaseGaussdbParserListener) ExitOpttemp(ctx *OpttempContext) {}

// EnterOpttableelementlist is called when production opttableelementlist is entered.
func (s *BaseGaussdbParserListener) EnterOpttableelementlist(ctx *OpttableelementlistContext) {}

// ExitOpttableelementlist is called when production opttableelementlist is exited.
func (s *BaseGaussdbParserListener) ExitOpttableelementlist(ctx *OpttableelementlistContext) {}

// EnterOpttypedtableelementlist is called when production opttypedtableelementlist is entered.
func (s *BaseGaussdbParserListener) EnterOpttypedtableelementlist(ctx *OpttypedtableelementlistContext) {
}

// ExitOpttypedtableelementlist is called when production opttypedtableelementlist is exited.
func (s *BaseGaussdbParserListener) ExitOpttypedtableelementlist(ctx *OpttypedtableelementlistContext) {
}

// EnterTableelementlist is called when production tableelementlist is entered.
func (s *BaseGaussdbParserListener) EnterTableelementlist(ctx *TableelementlistContext) {}

// ExitTableelementlist is called when production tableelementlist is exited.
func (s *BaseGaussdbParserListener) ExitTableelementlist(ctx *TableelementlistContext) {}

// EnterTypedtableelementlist is called when production typedtableelementlist is entered.
func (s *BaseGaussdbParserListener) EnterTypedtableelementlist(ctx *TypedtableelementlistContext) {}

// ExitTypedtableelementlist is called when production typedtableelementlist is exited.
func (s *BaseGaussdbParserListener) ExitTypedtableelementlist(ctx *TypedtableelementlistContext) {}

// EnterTableelement is called when production tableelement is entered.
func (s *BaseGaussdbParserListener) EnterTableelement(ctx *TableelementContext) {}

// ExitTableelement is called when production tableelement is exited.
func (s *BaseGaussdbParserListener) ExitTableelement(ctx *TableelementContext) {}

// EnterTypedtableelement is called when production typedtableelement is entered.
func (s *BaseGaussdbParserListener) EnterTypedtableelement(ctx *TypedtableelementContext) {}

// ExitTypedtableelement is called when production typedtableelement is exited.
func (s *BaseGaussdbParserListener) ExitTypedtableelement(ctx *TypedtableelementContext) {}

// EnterColumnDef is called when production columnDef is entered.
func (s *BaseGaussdbParserListener) EnterColumnDef(ctx *ColumnDefContext) {}

// ExitColumnDef is called when production columnDef is exited.
func (s *BaseGaussdbParserListener) ExitColumnDef(ctx *ColumnDefContext) {}

// EnterColumnOptions is called when production columnOptions is entered.
func (s *BaseGaussdbParserListener) EnterColumnOptions(ctx *ColumnOptionsContext) {}

// ExitColumnOptions is called when production columnOptions is exited.
func (s *BaseGaussdbParserListener) ExitColumnOptions(ctx *ColumnOptionsContext) {}

// EnterColquallist is called when production colquallist is entered.
func (s *BaseGaussdbParserListener) EnterColquallist(ctx *ColquallistContext) {}

// ExitColquallist is called when production colquallist is exited.
func (s *BaseGaussdbParserListener) ExitColquallist(ctx *ColquallistContext) {}

// EnterColconstraint is called when production colconstraint is entered.
func (s *BaseGaussdbParserListener) EnterColconstraint(ctx *ColconstraintContext) {}

// ExitColconstraint is called when production colconstraint is exited.
func (s *BaseGaussdbParserListener) ExitColconstraint(ctx *ColconstraintContext) {}

// EnterColconstraintelem is called when production colconstraintelem is entered.
func (s *BaseGaussdbParserListener) EnterColconstraintelem(ctx *ColconstraintelemContext) {}

// ExitColconstraintelem is called when production colconstraintelem is exited.
func (s *BaseGaussdbParserListener) ExitColconstraintelem(ctx *ColconstraintelemContext) {}

// EnterGenerated_when is called when production generated_when is entered.
func (s *BaseGaussdbParserListener) EnterGenerated_when(ctx *Generated_whenContext) {}

// ExitGenerated_when is called when production generated_when is exited.
func (s *BaseGaussdbParserListener) ExitGenerated_when(ctx *Generated_whenContext) {}

// EnterConstraintattr is called when production constraintattr is entered.
func (s *BaseGaussdbParserListener) EnterConstraintattr(ctx *ConstraintattrContext) {}

// ExitConstraintattr is called when production constraintattr is exited.
func (s *BaseGaussdbParserListener) ExitConstraintattr(ctx *ConstraintattrContext) {}

// EnterTablelikeclause is called when production tablelikeclause is entered.
func (s *BaseGaussdbParserListener) EnterTablelikeclause(ctx *TablelikeclauseContext) {}

// ExitTablelikeclause is called when production tablelikeclause is exited.
func (s *BaseGaussdbParserListener) ExitTablelikeclause(ctx *TablelikeclauseContext) {}

// EnterTablelikeoptionlist is called when production tablelikeoptionlist is entered.
func (s *BaseGaussdbParserListener) EnterTablelikeoptionlist(ctx *TablelikeoptionlistContext) {}

// ExitTablelikeoptionlist is called when production tablelikeoptionlist is exited.
func (s *BaseGaussdbParserListener) ExitTablelikeoptionlist(ctx *TablelikeoptionlistContext) {}

// EnterTablelikeoption is called when production tablelikeoption is entered.
func (s *BaseGaussdbParserListener) EnterTablelikeoption(ctx *TablelikeoptionContext) {}

// ExitTablelikeoption is called when production tablelikeoption is exited.
func (s *BaseGaussdbParserListener) ExitTablelikeoption(ctx *TablelikeoptionContext) {}

// EnterTableconstraint is called when production tableconstraint is entered.
func (s *BaseGaussdbParserListener) EnterTableconstraint(ctx *TableconstraintContext) {}

// ExitTableconstraint is called when production tableconstraint is exited.
func (s *BaseGaussdbParserListener) ExitTableconstraint(ctx *TableconstraintContext) {}

// EnterConstraintelem is called when production constraintelem is entered.
func (s *BaseGaussdbParserListener) EnterConstraintelem(ctx *ConstraintelemContext) {}

// ExitConstraintelem is called when production constraintelem is exited.
func (s *BaseGaussdbParserListener) ExitConstraintelem(ctx *ConstraintelemContext) {}

// EnterHtap_action is called when production htap_action is entered.
func (s *BaseGaussdbParserListener) EnterHtap_action(ctx *Htap_actionContext) {}

// ExitHtap_action is called when production htap_action is exited.
func (s *BaseGaussdbParserListener) ExitHtap_action(ctx *Htap_actionContext) {}

// EnterOpt_no_inherit is called when production opt_no_inherit is entered.
func (s *BaseGaussdbParserListener) EnterOpt_no_inherit(ctx *Opt_no_inheritContext) {}

// ExitOpt_no_inherit is called when production opt_no_inherit is exited.
func (s *BaseGaussdbParserListener) ExitOpt_no_inherit(ctx *Opt_no_inheritContext) {}

// EnterOpt_column_list is called when production opt_column_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_column_list(ctx *Opt_column_listContext) {}

// ExitOpt_column_list is called when production opt_column_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_column_list(ctx *Opt_column_listContext) {}

// EnterColumnlist is called when production columnlist is entered.
func (s *BaseGaussdbParserListener) EnterColumnlist(ctx *ColumnlistContext) {}

// ExitColumnlist is called when production columnlist is exited.
func (s *BaseGaussdbParserListener) ExitColumnlist(ctx *ColumnlistContext) {}

// EnterColumnElem is called when production columnElem is entered.
func (s *BaseGaussdbParserListener) EnterColumnElem(ctx *ColumnElemContext) {}

// ExitColumnElem is called when production columnElem is exited.
func (s *BaseGaussdbParserListener) ExitColumnElem(ctx *ColumnElemContext) {}

// EnterColumn_length is called when production column_length is entered.
func (s *BaseGaussdbParserListener) EnterColumn_length(ctx *Column_lengthContext) {}

// ExitColumn_length is called when production column_length is exited.
func (s *BaseGaussdbParserListener) ExitColumn_length(ctx *Column_lengthContext) {}

// EnterOpt_c_include is called when production opt_c_include is entered.
func (s *BaseGaussdbParserListener) EnterOpt_c_include(ctx *Opt_c_includeContext) {}

// ExitOpt_c_include is called when production opt_c_include is exited.
func (s *BaseGaussdbParserListener) ExitOpt_c_include(ctx *Opt_c_includeContext) {}

// EnterKey_match is called when production key_match is entered.
func (s *BaseGaussdbParserListener) EnterKey_match(ctx *Key_matchContext) {}

// ExitKey_match is called when production key_match is exited.
func (s *BaseGaussdbParserListener) ExitKey_match(ctx *Key_matchContext) {}

// EnterExclusionconstraintlist is called when production exclusionconstraintlist is entered.
func (s *BaseGaussdbParserListener) EnterExclusionconstraintlist(ctx *ExclusionconstraintlistContext) {
}

// ExitExclusionconstraintlist is called when production exclusionconstraintlist is exited.
func (s *BaseGaussdbParserListener) ExitExclusionconstraintlist(ctx *ExclusionconstraintlistContext) {
}

// EnterExclusionconstraintelem is called when production exclusionconstraintelem is entered.
func (s *BaseGaussdbParserListener) EnterExclusionconstraintelem(ctx *ExclusionconstraintelemContext) {
}

// ExitExclusionconstraintelem is called when production exclusionconstraintelem is exited.
func (s *BaseGaussdbParserListener) ExitExclusionconstraintelem(ctx *ExclusionconstraintelemContext) {
}

// EnterExclusionwhereclause is called when production exclusionwhereclause is entered.
func (s *BaseGaussdbParserListener) EnterExclusionwhereclause(ctx *ExclusionwhereclauseContext) {}

// ExitExclusionwhereclause is called when production exclusionwhereclause is exited.
func (s *BaseGaussdbParserListener) ExitExclusionwhereclause(ctx *ExclusionwhereclauseContext) {}

// EnterKey_actions is called when production key_actions is entered.
func (s *BaseGaussdbParserListener) EnterKey_actions(ctx *Key_actionsContext) {}

// ExitKey_actions is called when production key_actions is exited.
func (s *BaseGaussdbParserListener) ExitKey_actions(ctx *Key_actionsContext) {}

// EnterKey_update is called when production key_update is entered.
func (s *BaseGaussdbParserListener) EnterKey_update(ctx *Key_updateContext) {}

// ExitKey_update is called when production key_update is exited.
func (s *BaseGaussdbParserListener) ExitKey_update(ctx *Key_updateContext) {}

// EnterKey_delete is called when production key_delete is entered.
func (s *BaseGaussdbParserListener) EnterKey_delete(ctx *Key_deleteContext) {}

// ExitKey_delete is called when production key_delete is exited.
func (s *BaseGaussdbParserListener) ExitKey_delete(ctx *Key_deleteContext) {}

// EnterKey_action is called when production key_action is entered.
func (s *BaseGaussdbParserListener) EnterKey_action(ctx *Key_actionContext) {}

// ExitKey_action is called when production key_action is exited.
func (s *BaseGaussdbParserListener) ExitKey_action(ctx *Key_actionContext) {}

// EnterOptinherit is called when production optinherit is entered.
func (s *BaseGaussdbParserListener) EnterOptinherit(ctx *OptinheritContext) {}

// ExitOptinherit is called when production optinherit is exited.
func (s *BaseGaussdbParserListener) ExitOptinherit(ctx *OptinheritContext) {}

// EnterOptpartitionspec is called when production optpartitionspec is entered.
func (s *BaseGaussdbParserListener) EnterOptpartitionspec(ctx *OptpartitionspecContext) {}

// ExitOptpartitionspec is called when production optpartitionspec is exited.
func (s *BaseGaussdbParserListener) ExitOptpartitionspec(ctx *OptpartitionspecContext) {}

// EnterPartitionspec is called when production partitionspec is entered.
func (s *BaseGaussdbParserListener) EnterPartitionspec(ctx *PartitionspecContext) {}

// ExitPartitionspec is called when production partitionspec is exited.
func (s *BaseGaussdbParserListener) ExitPartitionspec(ctx *PartitionspecContext) {}

// EnterPart_params is called when production part_params is entered.
func (s *BaseGaussdbParserListener) EnterPart_params(ctx *Part_paramsContext) {}

// ExitPart_params is called when production part_params is exited.
func (s *BaseGaussdbParserListener) ExitPart_params(ctx *Part_paramsContext) {}

// EnterPart_elem is called when production part_elem is entered.
func (s *BaseGaussdbParserListener) EnterPart_elem(ctx *Part_elemContext) {}

// ExitPart_elem is called when production part_elem is exited.
func (s *BaseGaussdbParserListener) ExitPart_elem(ctx *Part_elemContext) {}

// EnterTable_access_method_clause is called when production table_access_method_clause is entered.
func (s *BaseGaussdbParserListener) EnterTable_access_method_clause(ctx *Table_access_method_clauseContext) {
}

// ExitTable_access_method_clause is called when production table_access_method_clause is exited.
func (s *BaseGaussdbParserListener) ExitTable_access_method_clause(ctx *Table_access_method_clauseContext) {
}

// EnterOptwith is called when production optwith is entered.
func (s *BaseGaussdbParserListener) EnterOptwith(ctx *OptwithContext) {}

// ExitOptwith is called when production optwith is exited.
func (s *BaseGaussdbParserListener) ExitOptwith(ctx *OptwithContext) {}

// EnterOncommitoption is called when production oncommitoption is entered.
func (s *BaseGaussdbParserListener) EnterOncommitoption(ctx *OncommitoptionContext) {}

// ExitOncommitoption is called when production oncommitoption is exited.
func (s *BaseGaussdbParserListener) ExitOncommitoption(ctx *OncommitoptionContext) {}

// EnterOpttablespace is called when production opttablespace is entered.
func (s *BaseGaussdbParserListener) EnterOpttablespace(ctx *OpttablespaceContext) {}

// ExitOpttablespace is called when production opttablespace is exited.
func (s *BaseGaussdbParserListener) ExitOpttablespace(ctx *OpttablespaceContext) {}

// EnterOptconstablespace is called when production optconstablespace is entered.
func (s *BaseGaussdbParserListener) EnterOptconstablespace(ctx *OptconstablespaceContext) {}

// ExitOptconstablespace is called when production optconstablespace is exited.
func (s *BaseGaussdbParserListener) ExitOptconstablespace(ctx *OptconstablespaceContext) {}

// EnterExistingindex is called when production existingindex is entered.
func (s *BaseGaussdbParserListener) EnterExistingindex(ctx *ExistingindexContext) {}

// ExitExistingindex is called when production existingindex is exited.
func (s *BaseGaussdbParserListener) ExitExistingindex(ctx *ExistingindexContext) {}

// EnterCreatestatsstmt is called when production createstatsstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatestatsstmt(ctx *CreatestatsstmtContext) {}

// ExitCreatestatsstmt is called when production createstatsstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatestatsstmt(ctx *CreatestatsstmtContext) {}

// EnterAlterstatsstmt is called when production alterstatsstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterstatsstmt(ctx *AlterstatsstmtContext) {}

// ExitAlterstatsstmt is called when production alterstatsstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterstatsstmt(ctx *AlterstatsstmtContext) {}

// EnterCreateasstmt is called when production createasstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateasstmt(ctx *CreateasstmtContext) {}

// ExitCreateasstmt is called when production createasstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateasstmt(ctx *CreateasstmtContext) {}

// EnterCreate_as_target is called when production create_as_target is entered.
func (s *BaseGaussdbParserListener) EnterCreate_as_target(ctx *Create_as_targetContext) {}

// ExitCreate_as_target is called when production create_as_target is exited.
func (s *BaseGaussdbParserListener) ExitCreate_as_target(ctx *Create_as_targetContext) {}

// EnterOpt_with_data is called when production opt_with_data is entered.
func (s *BaseGaussdbParserListener) EnterOpt_with_data(ctx *Opt_with_dataContext) {}

// ExitOpt_with_data is called when production opt_with_data is exited.
func (s *BaseGaussdbParserListener) ExitOpt_with_data(ctx *Opt_with_dataContext) {}

// EnterCreatematviewstmt is called when production creatematviewstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatematviewstmt(ctx *CreatematviewstmtContext) {}

// ExitCreatematviewstmt is called when production creatematviewstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatematviewstmt(ctx *CreatematviewstmtContext) {}

// EnterCreate_inc_mat_viewstmt is called when production create_inc_mat_viewstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreate_inc_mat_viewstmt(ctx *Create_inc_mat_viewstmtContext) {
}

// ExitCreate_inc_mat_viewstmt is called when production create_inc_mat_viewstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreate_inc_mat_viewstmt(ctx *Create_inc_mat_viewstmtContext) {
}

// EnterCreate_mv_target is called when production create_mv_target is entered.
func (s *BaseGaussdbParserListener) EnterCreate_mv_target(ctx *Create_mv_targetContext) {}

// ExitCreate_mv_target is called when production create_mv_target is exited.
func (s *BaseGaussdbParserListener) ExitCreate_mv_target(ctx *Create_mv_targetContext) {}

// EnterOptnolog is called when production optnolog is entered.
func (s *BaseGaussdbParserListener) EnterOptnolog(ctx *OptnologContext) {}

// ExitOptnolog is called when production optnolog is exited.
func (s *BaseGaussdbParserListener) ExitOptnolog(ctx *OptnologContext) {}

// EnterRefreshmatviewstmt is called when production refreshmatviewstmt is entered.
func (s *BaseGaussdbParserListener) EnterRefreshmatviewstmt(ctx *RefreshmatviewstmtContext) {}

// ExitRefreshmatviewstmt is called when production refreshmatviewstmt is exited.
func (s *BaseGaussdbParserListener) ExitRefreshmatviewstmt(ctx *RefreshmatviewstmtContext) {}

// EnterCreateseqstmt is called when production createseqstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateseqstmt(ctx *CreateseqstmtContext) {}

// ExitCreateseqstmt is called when production createseqstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateseqstmt(ctx *CreateseqstmtContext) {}

// EnterAlterseqstmt is called when production alterseqstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterseqstmt(ctx *AlterseqstmtContext) {}

// ExitAlterseqstmt is called when production alterseqstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterseqstmt(ctx *AlterseqstmtContext) {}

// EnterOptseqoptlist is called when production optseqoptlist is entered.
func (s *BaseGaussdbParserListener) EnterOptseqoptlist(ctx *OptseqoptlistContext) {}

// ExitOptseqoptlist is called when production optseqoptlist is exited.
func (s *BaseGaussdbParserListener) ExitOptseqoptlist(ctx *OptseqoptlistContext) {}

// EnterOptparenthesizedseqoptlist is called when production optparenthesizedseqoptlist is entered.
func (s *BaseGaussdbParserListener) EnterOptparenthesizedseqoptlist(ctx *OptparenthesizedseqoptlistContext) {
}

// ExitOptparenthesizedseqoptlist is called when production optparenthesizedseqoptlist is exited.
func (s *BaseGaussdbParserListener) ExitOptparenthesizedseqoptlist(ctx *OptparenthesizedseqoptlistContext) {
}

// EnterSeqoptlist is called when production seqoptlist is entered.
func (s *BaseGaussdbParserListener) EnterSeqoptlist(ctx *SeqoptlistContext) {}

// ExitSeqoptlist is called when production seqoptlist is exited.
func (s *BaseGaussdbParserListener) ExitSeqoptlist(ctx *SeqoptlistContext) {}

// EnterSeqoptelem is called when production seqoptelem is entered.
func (s *BaseGaussdbParserListener) EnterSeqoptelem(ctx *SeqoptelemContext) {}

// ExitSeqoptelem is called when production seqoptelem is exited.
func (s *BaseGaussdbParserListener) ExitSeqoptelem(ctx *SeqoptelemContext) {}

// EnterOpt_by is called when production opt_by is entered.
func (s *BaseGaussdbParserListener) EnterOpt_by(ctx *Opt_byContext) {}

// ExitOpt_by is called when production opt_by is exited.
func (s *BaseGaussdbParserListener) ExitOpt_by(ctx *Opt_byContext) {}

// EnterNumericonly is called when production numericonly is entered.
func (s *BaseGaussdbParserListener) EnterNumericonly(ctx *NumericonlyContext) {}

// ExitNumericonly is called when production numericonly is exited.
func (s *BaseGaussdbParserListener) ExitNumericonly(ctx *NumericonlyContext) {}

// EnterNumericonly_list is called when production numericonly_list is entered.
func (s *BaseGaussdbParserListener) EnterNumericonly_list(ctx *Numericonly_listContext) {}

// ExitNumericonly_list is called when production numericonly_list is exited.
func (s *BaseGaussdbParserListener) ExitNumericonly_list(ctx *Numericonly_listContext) {}

// EnterCreateplangstmt is called when production createplangstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateplangstmt(ctx *CreateplangstmtContext) {}

// ExitCreateplangstmt is called when production createplangstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateplangstmt(ctx *CreateplangstmtContext) {}

// EnterOpt_trusted is called when production opt_trusted is entered.
func (s *BaseGaussdbParserListener) EnterOpt_trusted(ctx *Opt_trustedContext) {}

// ExitOpt_trusted is called when production opt_trusted is exited.
func (s *BaseGaussdbParserListener) ExitOpt_trusted(ctx *Opt_trustedContext) {}

// EnterHandler_name is called when production handler_name is entered.
func (s *BaseGaussdbParserListener) EnterHandler_name(ctx *Handler_nameContext) {}

// ExitHandler_name is called when production handler_name is exited.
func (s *BaseGaussdbParserListener) ExitHandler_name(ctx *Handler_nameContext) {}

// EnterOpt_inline_handler is called when production opt_inline_handler is entered.
func (s *BaseGaussdbParserListener) EnterOpt_inline_handler(ctx *Opt_inline_handlerContext) {}

// ExitOpt_inline_handler is called when production opt_inline_handler is exited.
func (s *BaseGaussdbParserListener) ExitOpt_inline_handler(ctx *Opt_inline_handlerContext) {}

// EnterValidator_clause is called when production validator_clause is entered.
func (s *BaseGaussdbParserListener) EnterValidator_clause(ctx *Validator_clauseContext) {}

// ExitValidator_clause is called when production validator_clause is exited.
func (s *BaseGaussdbParserListener) ExitValidator_clause(ctx *Validator_clauseContext) {}

// EnterOpt_validator is called when production opt_validator is entered.
func (s *BaseGaussdbParserListener) EnterOpt_validator(ctx *Opt_validatorContext) {}

// ExitOpt_validator is called when production opt_validator is exited.
func (s *BaseGaussdbParserListener) ExitOpt_validator(ctx *Opt_validatorContext) {}

// EnterOpt_procedural is called when production opt_procedural is entered.
func (s *BaseGaussdbParserListener) EnterOpt_procedural(ctx *Opt_proceduralContext) {}

// ExitOpt_procedural is called when production opt_procedural is exited.
func (s *BaseGaussdbParserListener) ExitOpt_procedural(ctx *Opt_proceduralContext) {}

// EnterCreatetablespacestmt is called when production createtablespacestmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatetablespacestmt(ctx *CreatetablespacestmtContext) {}

// ExitCreatetablespacestmt is called when production createtablespacestmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatetablespacestmt(ctx *CreatetablespacestmtContext) {}

// EnterOpttablespaceowner is called when production opttablespaceowner is entered.
func (s *BaseGaussdbParserListener) EnterOpttablespaceowner(ctx *OpttablespaceownerContext) {}

// ExitOpttablespaceowner is called when production opttablespaceowner is exited.
func (s *BaseGaussdbParserListener) ExitOpttablespaceowner(ctx *OpttablespaceownerContext) {}

// EnterDroptablespacestmt is called when production droptablespacestmt is entered.
func (s *BaseGaussdbParserListener) EnterDroptablespacestmt(ctx *DroptablespacestmtContext) {}

// ExitDroptablespacestmt is called when production droptablespacestmt is exited.
func (s *BaseGaussdbParserListener) ExitDroptablespacestmt(ctx *DroptablespacestmtContext) {}

// EnterCreateextensionstmt is called when production createextensionstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateextensionstmt(ctx *CreateextensionstmtContext) {}

// ExitCreateextensionstmt is called when production createextensionstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateextensionstmt(ctx *CreateextensionstmtContext) {}

// EnterCreate_extension_opt_list is called when production create_extension_opt_list is entered.
func (s *BaseGaussdbParserListener) EnterCreate_extension_opt_list(ctx *Create_extension_opt_listContext) {
}

// ExitCreate_extension_opt_list is called when production create_extension_opt_list is exited.
func (s *BaseGaussdbParserListener) ExitCreate_extension_opt_list(ctx *Create_extension_opt_listContext) {
}

// EnterCreate_extension_opt_item is called when production create_extension_opt_item is entered.
func (s *BaseGaussdbParserListener) EnterCreate_extension_opt_item(ctx *Create_extension_opt_itemContext) {
}

// ExitCreate_extension_opt_item is called when production create_extension_opt_item is exited.
func (s *BaseGaussdbParserListener) ExitCreate_extension_opt_item(ctx *Create_extension_opt_itemContext) {
}

// EnterAlterextensionstmt is called when production alterextensionstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterextensionstmt(ctx *AlterextensionstmtContext) {}

// ExitAlterextensionstmt is called when production alterextensionstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterextensionstmt(ctx *AlterextensionstmtContext) {}

// EnterAlter_extension_opt_list is called when production alter_extension_opt_list is entered.
func (s *BaseGaussdbParserListener) EnterAlter_extension_opt_list(ctx *Alter_extension_opt_listContext) {
}

// ExitAlter_extension_opt_list is called when production alter_extension_opt_list is exited.
func (s *BaseGaussdbParserListener) ExitAlter_extension_opt_list(ctx *Alter_extension_opt_listContext) {
}

// EnterAlter_extension_opt_item is called when production alter_extension_opt_item is entered.
func (s *BaseGaussdbParserListener) EnterAlter_extension_opt_item(ctx *Alter_extension_opt_itemContext) {
}

// ExitAlter_extension_opt_item is called when production alter_extension_opt_item is exited.
func (s *BaseGaussdbParserListener) ExitAlter_extension_opt_item(ctx *Alter_extension_opt_itemContext) {
}

// EnterAlterextensioncontentsstmt is called when production alterextensioncontentsstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterextensioncontentsstmt(ctx *AlterextensioncontentsstmtContext) {
}

// ExitAlterextensioncontentsstmt is called when production alterextensioncontentsstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterextensioncontentsstmt(ctx *AlterextensioncontentsstmtContext) {
}

// EnterCreatefdwstmt is called when production createfdwstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatefdwstmt(ctx *CreatefdwstmtContext) {}

// ExitCreatefdwstmt is called when production createfdwstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatefdwstmt(ctx *CreatefdwstmtContext) {}

// EnterFdw_option is called when production fdw_option is entered.
func (s *BaseGaussdbParserListener) EnterFdw_option(ctx *Fdw_optionContext) {}

// ExitFdw_option is called when production fdw_option is exited.
func (s *BaseGaussdbParserListener) ExitFdw_option(ctx *Fdw_optionContext) {}

// EnterFdw_options is called when production fdw_options is entered.
func (s *BaseGaussdbParserListener) EnterFdw_options(ctx *Fdw_optionsContext) {}

// ExitFdw_options is called when production fdw_options is exited.
func (s *BaseGaussdbParserListener) ExitFdw_options(ctx *Fdw_optionsContext) {}

// EnterOpt_fdw_options is called when production opt_fdw_options is entered.
func (s *BaseGaussdbParserListener) EnterOpt_fdw_options(ctx *Opt_fdw_optionsContext) {}

// ExitOpt_fdw_options is called when production opt_fdw_options is exited.
func (s *BaseGaussdbParserListener) ExitOpt_fdw_options(ctx *Opt_fdw_optionsContext) {}

// EnterAlterfdwstmt is called when production alterfdwstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterfdwstmt(ctx *AlterfdwstmtContext) {}

// ExitAlterfdwstmt is called when production alterfdwstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterfdwstmt(ctx *AlterfdwstmtContext) {}

// EnterCreate_generic_options is called when production create_generic_options is entered.
func (s *BaseGaussdbParserListener) EnterCreate_generic_options(ctx *Create_generic_optionsContext) {}

// ExitCreate_generic_options is called when production create_generic_options is exited.
func (s *BaseGaussdbParserListener) ExitCreate_generic_options(ctx *Create_generic_optionsContext) {}

// EnterGeneric_option_list is called when production generic_option_list is entered.
func (s *BaseGaussdbParserListener) EnterGeneric_option_list(ctx *Generic_option_listContext) {}

// ExitGeneric_option_list is called when production generic_option_list is exited.
func (s *BaseGaussdbParserListener) ExitGeneric_option_list(ctx *Generic_option_listContext) {}

// EnterAlter_generic_options is called when production alter_generic_options is entered.
func (s *BaseGaussdbParserListener) EnterAlter_generic_options(ctx *Alter_generic_optionsContext) {}

// ExitAlter_generic_options is called when production alter_generic_options is exited.
func (s *BaseGaussdbParserListener) ExitAlter_generic_options(ctx *Alter_generic_optionsContext) {}

// EnterAlter_generic_option_list is called when production alter_generic_option_list is entered.
func (s *BaseGaussdbParserListener) EnterAlter_generic_option_list(ctx *Alter_generic_option_listContext) {
}

// ExitAlter_generic_option_list is called when production alter_generic_option_list is exited.
func (s *BaseGaussdbParserListener) ExitAlter_generic_option_list(ctx *Alter_generic_option_listContext) {
}

// EnterAlter_generic_option_elem is called when production alter_generic_option_elem is entered.
func (s *BaseGaussdbParserListener) EnterAlter_generic_option_elem(ctx *Alter_generic_option_elemContext) {
}

// ExitAlter_generic_option_elem is called when production alter_generic_option_elem is exited.
func (s *BaseGaussdbParserListener) ExitAlter_generic_option_elem(ctx *Alter_generic_option_elemContext) {
}

// EnterGeneric_option_elem is called when production generic_option_elem is entered.
func (s *BaseGaussdbParserListener) EnterGeneric_option_elem(ctx *Generic_option_elemContext) {}

// ExitGeneric_option_elem is called when production generic_option_elem is exited.
func (s *BaseGaussdbParserListener) ExitGeneric_option_elem(ctx *Generic_option_elemContext) {}

// EnterGeneric_option_name is called when production generic_option_name is entered.
func (s *BaseGaussdbParserListener) EnterGeneric_option_name(ctx *Generic_option_nameContext) {}

// ExitGeneric_option_name is called when production generic_option_name is exited.
func (s *BaseGaussdbParserListener) ExitGeneric_option_name(ctx *Generic_option_nameContext) {}

// EnterGeneric_option_arg is called when production generic_option_arg is entered.
func (s *BaseGaussdbParserListener) EnterGeneric_option_arg(ctx *Generic_option_argContext) {}

// ExitGeneric_option_arg is called when production generic_option_arg is exited.
func (s *BaseGaussdbParserListener) ExitGeneric_option_arg(ctx *Generic_option_argContext) {}

// EnterCreateforeignserverstmt is called when production createforeignserverstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateforeignserverstmt(ctx *CreateforeignserverstmtContext) {
}

// ExitCreateforeignserverstmt is called when production createforeignserverstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateforeignserverstmt(ctx *CreateforeignserverstmtContext) {
}

// EnterOpt_type is called when production opt_type is entered.
func (s *BaseGaussdbParserListener) EnterOpt_type(ctx *Opt_typeContext) {}

// ExitOpt_type is called when production opt_type is exited.
func (s *BaseGaussdbParserListener) ExitOpt_type(ctx *Opt_typeContext) {}

// EnterForeign_server_version is called when production foreign_server_version is entered.
func (s *BaseGaussdbParserListener) EnterForeign_server_version(ctx *Foreign_server_versionContext) {}

// ExitForeign_server_version is called when production foreign_server_version is exited.
func (s *BaseGaussdbParserListener) ExitForeign_server_version(ctx *Foreign_server_versionContext) {}

// EnterOpt_foreign_server_version is called when production opt_foreign_server_version is entered.
func (s *BaseGaussdbParserListener) EnterOpt_foreign_server_version(ctx *Opt_foreign_server_versionContext) {
}

// ExitOpt_foreign_server_version is called when production opt_foreign_server_version is exited.
func (s *BaseGaussdbParserListener) ExitOpt_foreign_server_version(ctx *Opt_foreign_server_versionContext) {
}

// EnterAlterforeignserverstmt is called when production alterforeignserverstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterforeignserverstmt(ctx *AlterforeignserverstmtContext) {}

// ExitAlterforeignserverstmt is called when production alterforeignserverstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterforeignserverstmt(ctx *AlterforeignserverstmtContext) {}

// EnterCreateforeigntablestmt is called when production createforeigntablestmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateforeigntablestmt(ctx *CreateforeigntablestmtContext) {}

// ExitCreateforeigntablestmt is called when production createforeigntablestmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateforeigntablestmt(ctx *CreateforeigntablestmtContext) {}

// EnterImportforeignschemastmt is called when production importforeignschemastmt is entered.
func (s *BaseGaussdbParserListener) EnterImportforeignschemastmt(ctx *ImportforeignschemastmtContext) {
}

// ExitImportforeignschemastmt is called when production importforeignschemastmt is exited.
func (s *BaseGaussdbParserListener) ExitImportforeignschemastmt(ctx *ImportforeignschemastmtContext) {
}

// EnterImport_qualification_type is called when production import_qualification_type is entered.
func (s *BaseGaussdbParserListener) EnterImport_qualification_type(ctx *Import_qualification_typeContext) {
}

// ExitImport_qualification_type is called when production import_qualification_type is exited.
func (s *BaseGaussdbParserListener) ExitImport_qualification_type(ctx *Import_qualification_typeContext) {
}

// EnterImport_qualification is called when production import_qualification is entered.
func (s *BaseGaussdbParserListener) EnterImport_qualification(ctx *Import_qualificationContext) {}

// ExitImport_qualification is called when production import_qualification is exited.
func (s *BaseGaussdbParserListener) ExitImport_qualification(ctx *Import_qualificationContext) {}

// EnterCreateusermappingstmt is called when production createusermappingstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateusermappingstmt(ctx *CreateusermappingstmtContext) {}

// ExitCreateusermappingstmt is called when production createusermappingstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateusermappingstmt(ctx *CreateusermappingstmtContext) {}

// EnterAuth_ident is called when production auth_ident is entered.
func (s *BaseGaussdbParserListener) EnterAuth_ident(ctx *Auth_identContext) {}

// ExitAuth_ident is called when production auth_ident is exited.
func (s *BaseGaussdbParserListener) ExitAuth_ident(ctx *Auth_identContext) {}

// EnterDropusermappingstmt is called when production dropusermappingstmt is entered.
func (s *BaseGaussdbParserListener) EnterDropusermappingstmt(ctx *DropusermappingstmtContext) {}

// ExitDropusermappingstmt is called when production dropusermappingstmt is exited.
func (s *BaseGaussdbParserListener) ExitDropusermappingstmt(ctx *DropusermappingstmtContext) {}

// EnterAlterusermappingstmt is called when production alterusermappingstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterusermappingstmt(ctx *AlterusermappingstmtContext) {}

// ExitAlterusermappingstmt is called when production alterusermappingstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterusermappingstmt(ctx *AlterusermappingstmtContext) {}

// EnterCreatepolicystmt is called when production createpolicystmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatepolicystmt(ctx *CreatepolicystmtContext) {}

// ExitCreatepolicystmt is called when production createpolicystmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatepolicystmt(ctx *CreatepolicystmtContext) {}

// EnterAlterpolicystmt is called when production alterpolicystmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterpolicystmt(ctx *AlterpolicystmtContext) {}

// ExitAlterpolicystmt is called when production alterpolicystmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterpolicystmt(ctx *AlterpolicystmtContext) {}

// EnterRowsecurityoptionalexpr is called when production rowsecurityoptionalexpr is entered.
func (s *BaseGaussdbParserListener) EnterRowsecurityoptionalexpr(ctx *RowsecurityoptionalexprContext) {
}

// ExitRowsecurityoptionalexpr is called when production rowsecurityoptionalexpr is exited.
func (s *BaseGaussdbParserListener) ExitRowsecurityoptionalexpr(ctx *RowsecurityoptionalexprContext) {
}

// EnterRowsecurityoptionalwithcheck is called when production rowsecurityoptionalwithcheck is entered.
func (s *BaseGaussdbParserListener) EnterRowsecurityoptionalwithcheck(ctx *RowsecurityoptionalwithcheckContext) {
}

// ExitRowsecurityoptionalwithcheck is called when production rowsecurityoptionalwithcheck is exited.
func (s *BaseGaussdbParserListener) ExitRowsecurityoptionalwithcheck(ctx *RowsecurityoptionalwithcheckContext) {
}

// EnterRowsecuritydefaulttorole is called when production rowsecuritydefaulttorole is entered.
func (s *BaseGaussdbParserListener) EnterRowsecuritydefaulttorole(ctx *RowsecuritydefaulttoroleContext) {
}

// ExitRowsecuritydefaulttorole is called when production rowsecuritydefaulttorole is exited.
func (s *BaseGaussdbParserListener) ExitRowsecuritydefaulttorole(ctx *RowsecuritydefaulttoroleContext) {
}

// EnterRowsecurityoptionaltorole is called when production rowsecurityoptionaltorole is entered.
func (s *BaseGaussdbParserListener) EnterRowsecurityoptionaltorole(ctx *RowsecurityoptionaltoroleContext) {
}

// ExitRowsecurityoptionaltorole is called when production rowsecurityoptionaltorole is exited.
func (s *BaseGaussdbParserListener) ExitRowsecurityoptionaltorole(ctx *RowsecurityoptionaltoroleContext) {
}

// EnterRowsecuritydefaultpermissive is called when production rowsecuritydefaultpermissive is entered.
func (s *BaseGaussdbParserListener) EnterRowsecuritydefaultpermissive(ctx *RowsecuritydefaultpermissiveContext) {
}

// ExitRowsecuritydefaultpermissive is called when production rowsecuritydefaultpermissive is exited.
func (s *BaseGaussdbParserListener) ExitRowsecuritydefaultpermissive(ctx *RowsecuritydefaultpermissiveContext) {
}

// EnterRowsecuritydefaultforcmd is called when production rowsecuritydefaultforcmd is entered.
func (s *BaseGaussdbParserListener) EnterRowsecuritydefaultforcmd(ctx *RowsecuritydefaultforcmdContext) {
}

// ExitRowsecuritydefaultforcmd is called when production rowsecuritydefaultforcmd is exited.
func (s *BaseGaussdbParserListener) ExitRowsecuritydefaultforcmd(ctx *RowsecuritydefaultforcmdContext) {
}

// EnterRow_security_cmd is called when production row_security_cmd is entered.
func (s *BaseGaussdbParserListener) EnterRow_security_cmd(ctx *Row_security_cmdContext) {}

// ExitRow_security_cmd is called when production row_security_cmd is exited.
func (s *BaseGaussdbParserListener) ExitRow_security_cmd(ctx *Row_security_cmdContext) {}

// EnterCreateamstmt is called when production createamstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateamstmt(ctx *CreateamstmtContext) {}

// ExitCreateamstmt is called when production createamstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateamstmt(ctx *CreateamstmtContext) {}

// EnterAm_type is called when production am_type is entered.
func (s *BaseGaussdbParserListener) EnterAm_type(ctx *Am_typeContext) {}

// ExitAm_type is called when production am_type is exited.
func (s *BaseGaussdbParserListener) ExitAm_type(ctx *Am_typeContext) {}

// EnterCreatetrigstmt is called when production createtrigstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatetrigstmt(ctx *CreatetrigstmtContext) {}

// ExitCreatetrigstmt is called when production createtrigstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatetrigstmt(ctx *CreatetrigstmtContext) {}

// EnterTriggeractiontime is called when production triggeractiontime is entered.
func (s *BaseGaussdbParserListener) EnterTriggeractiontime(ctx *TriggeractiontimeContext) {}

// ExitTriggeractiontime is called when production triggeractiontime is exited.
func (s *BaseGaussdbParserListener) ExitTriggeractiontime(ctx *TriggeractiontimeContext) {}

// EnterTriggerevents is called when production triggerevents is entered.
func (s *BaseGaussdbParserListener) EnterTriggerevents(ctx *TriggereventsContext) {}

// ExitTriggerevents is called when production triggerevents is exited.
func (s *BaseGaussdbParserListener) ExitTriggerevents(ctx *TriggereventsContext) {}

// EnterTriggeroneevent is called when production triggeroneevent is entered.
func (s *BaseGaussdbParserListener) EnterTriggeroneevent(ctx *TriggeroneeventContext) {}

// ExitTriggeroneevent is called when production triggeroneevent is exited.
func (s *BaseGaussdbParserListener) ExitTriggeroneevent(ctx *TriggeroneeventContext) {}

// EnterTriggerreferencing is called when production triggerreferencing is entered.
func (s *BaseGaussdbParserListener) EnterTriggerreferencing(ctx *TriggerreferencingContext) {}

// ExitTriggerreferencing is called when production triggerreferencing is exited.
func (s *BaseGaussdbParserListener) ExitTriggerreferencing(ctx *TriggerreferencingContext) {}

// EnterTriggertransitions is called when production triggertransitions is entered.
func (s *BaseGaussdbParserListener) EnterTriggertransitions(ctx *TriggertransitionsContext) {}

// ExitTriggertransitions is called when production triggertransitions is exited.
func (s *BaseGaussdbParserListener) ExitTriggertransitions(ctx *TriggertransitionsContext) {}

// EnterTriggertransition is called when production triggertransition is entered.
func (s *BaseGaussdbParserListener) EnterTriggertransition(ctx *TriggertransitionContext) {}

// ExitTriggertransition is called when production triggertransition is exited.
func (s *BaseGaussdbParserListener) ExitTriggertransition(ctx *TriggertransitionContext) {}

// EnterTransitionoldornew is called when production transitionoldornew is entered.
func (s *BaseGaussdbParserListener) EnterTransitionoldornew(ctx *TransitionoldornewContext) {}

// ExitTransitionoldornew is called when production transitionoldornew is exited.
func (s *BaseGaussdbParserListener) ExitTransitionoldornew(ctx *TransitionoldornewContext) {}

// EnterTransitionrowortable is called when production transitionrowortable is entered.
func (s *BaseGaussdbParserListener) EnterTransitionrowortable(ctx *TransitionrowortableContext) {}

// ExitTransitionrowortable is called when production transitionrowortable is exited.
func (s *BaseGaussdbParserListener) ExitTransitionrowortable(ctx *TransitionrowortableContext) {}

// EnterTransitionrelname is called when production transitionrelname is entered.
func (s *BaseGaussdbParserListener) EnterTransitionrelname(ctx *TransitionrelnameContext) {}

// ExitTransitionrelname is called when production transitionrelname is exited.
func (s *BaseGaussdbParserListener) ExitTransitionrelname(ctx *TransitionrelnameContext) {}

// EnterTriggerforspec is called when production triggerforspec is entered.
func (s *BaseGaussdbParserListener) EnterTriggerforspec(ctx *TriggerforspecContext) {}

// ExitTriggerforspec is called when production triggerforspec is exited.
func (s *BaseGaussdbParserListener) ExitTriggerforspec(ctx *TriggerforspecContext) {}

// EnterTriggerforopteach is called when production triggerforopteach is entered.
func (s *BaseGaussdbParserListener) EnterTriggerforopteach(ctx *TriggerforopteachContext) {}

// ExitTriggerforopteach is called when production triggerforopteach is exited.
func (s *BaseGaussdbParserListener) ExitTriggerforopteach(ctx *TriggerforopteachContext) {}

// EnterTriggerfortype is called when production triggerfortype is entered.
func (s *BaseGaussdbParserListener) EnterTriggerfortype(ctx *TriggerfortypeContext) {}

// ExitTriggerfortype is called when production triggerfortype is exited.
func (s *BaseGaussdbParserListener) ExitTriggerfortype(ctx *TriggerfortypeContext) {}

// EnterTriggerwhen is called when production triggerwhen is entered.
func (s *BaseGaussdbParserListener) EnterTriggerwhen(ctx *TriggerwhenContext) {}

// ExitTriggerwhen is called when production triggerwhen is exited.
func (s *BaseGaussdbParserListener) ExitTriggerwhen(ctx *TriggerwhenContext) {}

// EnterFunction_or_procedure is called when production function_or_procedure is entered.
func (s *BaseGaussdbParserListener) EnterFunction_or_procedure(ctx *Function_or_procedureContext) {}

// ExitFunction_or_procedure is called when production function_or_procedure is exited.
func (s *BaseGaussdbParserListener) ExitFunction_or_procedure(ctx *Function_or_procedureContext) {}

// EnterTriggerfuncargs is called when production triggerfuncargs is entered.
func (s *BaseGaussdbParserListener) EnterTriggerfuncargs(ctx *TriggerfuncargsContext) {}

// ExitTriggerfuncargs is called when production triggerfuncargs is exited.
func (s *BaseGaussdbParserListener) ExitTriggerfuncargs(ctx *TriggerfuncargsContext) {}

// EnterTriggerfuncarg is called when production triggerfuncarg is entered.
func (s *BaseGaussdbParserListener) EnterTriggerfuncarg(ctx *TriggerfuncargContext) {}

// ExitTriggerfuncarg is called when production triggerfuncarg is exited.
func (s *BaseGaussdbParserListener) ExitTriggerfuncarg(ctx *TriggerfuncargContext) {}

// EnterOptconstrfromtable is called when production optconstrfromtable is entered.
func (s *BaseGaussdbParserListener) EnterOptconstrfromtable(ctx *OptconstrfromtableContext) {}

// ExitOptconstrfromtable is called when production optconstrfromtable is exited.
func (s *BaseGaussdbParserListener) ExitOptconstrfromtable(ctx *OptconstrfromtableContext) {}

// EnterConstraintattributespec is called when production constraintattributespec is entered.
func (s *BaseGaussdbParserListener) EnterConstraintattributespec(ctx *ConstraintattributespecContext) {
}

// ExitConstraintattributespec is called when production constraintattributespec is exited.
func (s *BaseGaussdbParserListener) ExitConstraintattributespec(ctx *ConstraintattributespecContext) {
}

// EnterConstraintattributeElem is called when production constraintattributeElem is entered.
func (s *BaseGaussdbParserListener) EnterConstraintattributeElem(ctx *ConstraintattributeElemContext) {
}

// ExitConstraintattributeElem is called when production constraintattributeElem is exited.
func (s *BaseGaussdbParserListener) ExitConstraintattributeElem(ctx *ConstraintattributeElemContext) {
}

// EnterCreateeventtrigstmt is called when production createeventtrigstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateeventtrigstmt(ctx *CreateeventtrigstmtContext) {}

// ExitCreateeventtrigstmt is called when production createeventtrigstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateeventtrigstmt(ctx *CreateeventtrigstmtContext) {}

// EnterEvent_trigger_when_list is called when production event_trigger_when_list is entered.
func (s *BaseGaussdbParserListener) EnterEvent_trigger_when_list(ctx *Event_trigger_when_listContext) {
}

// ExitEvent_trigger_when_list is called when production event_trigger_when_list is exited.
func (s *BaseGaussdbParserListener) ExitEvent_trigger_when_list(ctx *Event_trigger_when_listContext) {
}

// EnterEvent_trigger_when_item is called when production event_trigger_when_item is entered.
func (s *BaseGaussdbParserListener) EnterEvent_trigger_when_item(ctx *Event_trigger_when_itemContext) {
}

// ExitEvent_trigger_when_item is called when production event_trigger_when_item is exited.
func (s *BaseGaussdbParserListener) ExitEvent_trigger_when_item(ctx *Event_trigger_when_itemContext) {
}

// EnterEvent_trigger_value_list is called when production event_trigger_value_list is entered.
func (s *BaseGaussdbParserListener) EnterEvent_trigger_value_list(ctx *Event_trigger_value_listContext) {
}

// ExitEvent_trigger_value_list is called when production event_trigger_value_list is exited.
func (s *BaseGaussdbParserListener) ExitEvent_trigger_value_list(ctx *Event_trigger_value_listContext) {
}

// EnterAltereventtrigstmt is called when production altereventtrigstmt is entered.
func (s *BaseGaussdbParserListener) EnterAltereventtrigstmt(ctx *AltereventtrigstmtContext) {}

// ExitAltereventtrigstmt is called when production altereventtrigstmt is exited.
func (s *BaseGaussdbParserListener) ExitAltereventtrigstmt(ctx *AltereventtrigstmtContext) {}

// EnterEnable_trigger is called when production enable_trigger is entered.
func (s *BaseGaussdbParserListener) EnterEnable_trigger(ctx *Enable_triggerContext) {}

// ExitEnable_trigger is called when production enable_trigger is exited.
func (s *BaseGaussdbParserListener) ExitEnable_trigger(ctx *Enable_triggerContext) {}

// EnterCreateassertionstmt is called when production createassertionstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateassertionstmt(ctx *CreateassertionstmtContext) {}

// ExitCreateassertionstmt is called when production createassertionstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateassertionstmt(ctx *CreateassertionstmtContext) {}

// EnterDefinestmt is called when production definestmt is entered.
func (s *BaseGaussdbParserListener) EnterDefinestmt(ctx *DefinestmtContext) {}

// ExitDefinestmt is called when production definestmt is exited.
func (s *BaseGaussdbParserListener) ExitDefinestmt(ctx *DefinestmtContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseGaussdbParserListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseGaussdbParserListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterDef_list is called when production def_list is entered.
func (s *BaseGaussdbParserListener) EnterDef_list(ctx *Def_listContext) {}

// ExitDef_list is called when production def_list is exited.
func (s *BaseGaussdbParserListener) ExitDef_list(ctx *Def_listContext) {}

// EnterDef_elem is called when production def_elem is entered.
func (s *BaseGaussdbParserListener) EnterDef_elem(ctx *Def_elemContext) {}

// ExitDef_elem is called when production def_elem is exited.
func (s *BaseGaussdbParserListener) ExitDef_elem(ctx *Def_elemContext) {}

// EnterDef_arg is called when production def_arg is entered.
func (s *BaseGaussdbParserListener) EnterDef_arg(ctx *Def_argContext) {}

// ExitDef_arg is called when production def_arg is exited.
func (s *BaseGaussdbParserListener) ExitDef_arg(ctx *Def_argContext) {}

// EnterOld_aggr_definition is called when production old_aggr_definition is entered.
func (s *BaseGaussdbParserListener) EnterOld_aggr_definition(ctx *Old_aggr_definitionContext) {}

// ExitOld_aggr_definition is called when production old_aggr_definition is exited.
func (s *BaseGaussdbParserListener) ExitOld_aggr_definition(ctx *Old_aggr_definitionContext) {}

// EnterOld_aggr_list is called when production old_aggr_list is entered.
func (s *BaseGaussdbParserListener) EnterOld_aggr_list(ctx *Old_aggr_listContext) {}

// ExitOld_aggr_list is called when production old_aggr_list is exited.
func (s *BaseGaussdbParserListener) ExitOld_aggr_list(ctx *Old_aggr_listContext) {}

// EnterOld_aggr_elem is called when production old_aggr_elem is entered.
func (s *BaseGaussdbParserListener) EnterOld_aggr_elem(ctx *Old_aggr_elemContext) {}

// ExitOld_aggr_elem is called when production old_aggr_elem is exited.
func (s *BaseGaussdbParserListener) ExitOld_aggr_elem(ctx *Old_aggr_elemContext) {}

// EnterOpt_enum_val_list is called when production opt_enum_val_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_enum_val_list(ctx *Opt_enum_val_listContext) {}

// ExitOpt_enum_val_list is called when production opt_enum_val_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_enum_val_list(ctx *Opt_enum_val_listContext) {}

// EnterEnum_val_list is called when production enum_val_list is entered.
func (s *BaseGaussdbParserListener) EnterEnum_val_list(ctx *Enum_val_listContext) {}

// ExitEnum_val_list is called when production enum_val_list is exited.
func (s *BaseGaussdbParserListener) ExitEnum_val_list(ctx *Enum_val_listContext) {}

// EnterAlterenumstmt is called when production alterenumstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterenumstmt(ctx *AlterenumstmtContext) {}

// ExitAlterenumstmt is called when production alterenumstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterenumstmt(ctx *AlterenumstmtContext) {}

// EnterOpt_if_not_exists is called when production opt_if_not_exists is entered.
func (s *BaseGaussdbParserListener) EnterOpt_if_not_exists(ctx *Opt_if_not_existsContext) {}

// ExitOpt_if_not_exists is called when production opt_if_not_exists is exited.
func (s *BaseGaussdbParserListener) ExitOpt_if_not_exists(ctx *Opt_if_not_existsContext) {}

// EnterCreateopclassstmt is called when production createopclassstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateopclassstmt(ctx *CreateopclassstmtContext) {}

// ExitCreateopclassstmt is called when production createopclassstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateopclassstmt(ctx *CreateopclassstmtContext) {}

// EnterOpclass_item_list is called when production opclass_item_list is entered.
func (s *BaseGaussdbParserListener) EnterOpclass_item_list(ctx *Opclass_item_listContext) {}

// ExitOpclass_item_list is called when production opclass_item_list is exited.
func (s *BaseGaussdbParserListener) ExitOpclass_item_list(ctx *Opclass_item_listContext) {}

// EnterOpclass_item is called when production opclass_item is entered.
func (s *BaseGaussdbParserListener) EnterOpclass_item(ctx *Opclass_itemContext) {}

// ExitOpclass_item is called when production opclass_item is exited.
func (s *BaseGaussdbParserListener) ExitOpclass_item(ctx *Opclass_itemContext) {}

// EnterOpt_default is called when production opt_default is entered.
func (s *BaseGaussdbParserListener) EnterOpt_default(ctx *Opt_defaultContext) {}

// ExitOpt_default is called when production opt_default is exited.
func (s *BaseGaussdbParserListener) ExitOpt_default(ctx *Opt_defaultContext) {}

// EnterOpt_opfamily is called when production opt_opfamily is entered.
func (s *BaseGaussdbParserListener) EnterOpt_opfamily(ctx *Opt_opfamilyContext) {}

// ExitOpt_opfamily is called when production opt_opfamily is exited.
func (s *BaseGaussdbParserListener) ExitOpt_opfamily(ctx *Opt_opfamilyContext) {}

// EnterOpclass_purpose is called when production opclass_purpose is entered.
func (s *BaseGaussdbParserListener) EnterOpclass_purpose(ctx *Opclass_purposeContext) {}

// ExitOpclass_purpose is called when production opclass_purpose is exited.
func (s *BaseGaussdbParserListener) ExitOpclass_purpose(ctx *Opclass_purposeContext) {}

// EnterOpt_recheck is called when production opt_recheck is entered.
func (s *BaseGaussdbParserListener) EnterOpt_recheck(ctx *Opt_recheckContext) {}

// ExitOpt_recheck is called when production opt_recheck is exited.
func (s *BaseGaussdbParserListener) ExitOpt_recheck(ctx *Opt_recheckContext) {}

// EnterCreateopfamilystmt is called when production createopfamilystmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateopfamilystmt(ctx *CreateopfamilystmtContext) {}

// ExitCreateopfamilystmt is called when production createopfamilystmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateopfamilystmt(ctx *CreateopfamilystmtContext) {}

// EnterAlteropfamilystmt is called when production alteropfamilystmt is entered.
func (s *BaseGaussdbParserListener) EnterAlteropfamilystmt(ctx *AlteropfamilystmtContext) {}

// ExitAlteropfamilystmt is called when production alteropfamilystmt is exited.
func (s *BaseGaussdbParserListener) ExitAlteropfamilystmt(ctx *AlteropfamilystmtContext) {}

// EnterOpclass_drop_list is called when production opclass_drop_list is entered.
func (s *BaseGaussdbParserListener) EnterOpclass_drop_list(ctx *Opclass_drop_listContext) {}

// ExitOpclass_drop_list is called when production opclass_drop_list is exited.
func (s *BaseGaussdbParserListener) ExitOpclass_drop_list(ctx *Opclass_drop_listContext) {}

// EnterOpclass_drop is called when production opclass_drop is entered.
func (s *BaseGaussdbParserListener) EnterOpclass_drop(ctx *Opclass_dropContext) {}

// ExitOpclass_drop is called when production opclass_drop is exited.
func (s *BaseGaussdbParserListener) ExitOpclass_drop(ctx *Opclass_dropContext) {}

// EnterDropopclassstmt is called when production dropopclassstmt is entered.
func (s *BaseGaussdbParserListener) EnterDropopclassstmt(ctx *DropopclassstmtContext) {}

// ExitDropopclassstmt is called when production dropopclassstmt is exited.
func (s *BaseGaussdbParserListener) ExitDropopclassstmt(ctx *DropopclassstmtContext) {}

// EnterDropopfamilystmt is called when production dropopfamilystmt is entered.
func (s *BaseGaussdbParserListener) EnterDropopfamilystmt(ctx *DropopfamilystmtContext) {}

// ExitDropopfamilystmt is called when production dropopfamilystmt is exited.
func (s *BaseGaussdbParserListener) ExitDropopfamilystmt(ctx *DropopfamilystmtContext) {}

// EnterDropownedstmt is called when production dropownedstmt is entered.
func (s *BaseGaussdbParserListener) EnterDropownedstmt(ctx *DropownedstmtContext) {}

// ExitDropownedstmt is called when production dropownedstmt is exited.
func (s *BaseGaussdbParserListener) ExitDropownedstmt(ctx *DropownedstmtContext) {}

// EnterReassignownedstmt is called when production reassignownedstmt is entered.
func (s *BaseGaussdbParserListener) EnterReassignownedstmt(ctx *ReassignownedstmtContext) {}

// ExitReassignownedstmt is called when production reassignownedstmt is exited.
func (s *BaseGaussdbParserListener) ExitReassignownedstmt(ctx *ReassignownedstmtContext) {}

// EnterDropstmt is called when production dropstmt is entered.
func (s *BaseGaussdbParserListener) EnterDropstmt(ctx *DropstmtContext) {}

// ExitDropstmt is called when production dropstmt is exited.
func (s *BaseGaussdbParserListener) ExitDropstmt(ctx *DropstmtContext) {}

// EnterObject_type_any_name is called when production object_type_any_name is entered.
func (s *BaseGaussdbParserListener) EnterObject_type_any_name(ctx *Object_type_any_nameContext) {}

// ExitObject_type_any_name is called when production object_type_any_name is exited.
func (s *BaseGaussdbParserListener) ExitObject_type_any_name(ctx *Object_type_any_nameContext) {}

// EnterObject_type_name is called when production object_type_name is entered.
func (s *BaseGaussdbParserListener) EnterObject_type_name(ctx *Object_type_nameContext) {}

// ExitObject_type_name is called when production object_type_name is exited.
func (s *BaseGaussdbParserListener) ExitObject_type_name(ctx *Object_type_nameContext) {}

// EnterDrop_type_name is called when production drop_type_name is entered.
func (s *BaseGaussdbParserListener) EnterDrop_type_name(ctx *Drop_type_nameContext) {}

// ExitDrop_type_name is called when production drop_type_name is exited.
func (s *BaseGaussdbParserListener) ExitDrop_type_name(ctx *Drop_type_nameContext) {}

// EnterObject_type_name_on_any_name is called when production object_type_name_on_any_name is entered.
func (s *BaseGaussdbParserListener) EnterObject_type_name_on_any_name(ctx *Object_type_name_on_any_nameContext) {
}

// ExitObject_type_name_on_any_name is called when production object_type_name_on_any_name is exited.
func (s *BaseGaussdbParserListener) ExitObject_type_name_on_any_name(ctx *Object_type_name_on_any_nameContext) {
}

// EnterAny_name_list is called when production any_name_list is entered.
func (s *BaseGaussdbParserListener) EnterAny_name_list(ctx *Any_name_listContext) {}

// ExitAny_name_list is called when production any_name_list is exited.
func (s *BaseGaussdbParserListener) ExitAny_name_list(ctx *Any_name_listContext) {}

// EnterAny_name is called when production any_name is entered.
func (s *BaseGaussdbParserListener) EnterAny_name(ctx *Any_nameContext) {}

// ExitAny_name is called when production any_name is exited.
func (s *BaseGaussdbParserListener) ExitAny_name(ctx *Any_nameContext) {}

// EnterAttrs is called when production attrs is entered.
func (s *BaseGaussdbParserListener) EnterAttrs(ctx *AttrsContext) {}

// ExitAttrs is called when production attrs is exited.
func (s *BaseGaussdbParserListener) ExitAttrs(ctx *AttrsContext) {}

// EnterType_name_list is called when production type_name_list is entered.
func (s *BaseGaussdbParserListener) EnterType_name_list(ctx *Type_name_listContext) {}

// ExitType_name_list is called when production type_name_list is exited.
func (s *BaseGaussdbParserListener) ExitType_name_list(ctx *Type_name_listContext) {}

// EnterTruncatestmt is called when production truncatestmt is entered.
func (s *BaseGaussdbParserListener) EnterTruncatestmt(ctx *TruncatestmtContext) {}

// ExitTruncatestmt is called when production truncatestmt is exited.
func (s *BaseGaussdbParserListener) ExitTruncatestmt(ctx *TruncatestmtContext) {}

// EnterOpt_restart_seqs is called when production opt_restart_seqs is entered.
func (s *BaseGaussdbParserListener) EnterOpt_restart_seqs(ctx *Opt_restart_seqsContext) {}

// ExitOpt_restart_seqs is called when production opt_restart_seqs is exited.
func (s *BaseGaussdbParserListener) ExitOpt_restart_seqs(ctx *Opt_restart_seqsContext) {}

// EnterCommentstmt is called when production commentstmt is entered.
func (s *BaseGaussdbParserListener) EnterCommentstmt(ctx *CommentstmtContext) {}

// ExitCommentstmt is called when production commentstmt is exited.
func (s *BaseGaussdbParserListener) ExitCommentstmt(ctx *CommentstmtContext) {}

// EnterComment_text is called when production comment_text is entered.
func (s *BaseGaussdbParserListener) EnterComment_text(ctx *Comment_textContext) {}

// ExitComment_text is called when production comment_text is exited.
func (s *BaseGaussdbParserListener) ExitComment_text(ctx *Comment_textContext) {}

// EnterSeclabelstmt is called when production seclabelstmt is entered.
func (s *BaseGaussdbParserListener) EnterSeclabelstmt(ctx *SeclabelstmtContext) {}

// ExitSeclabelstmt is called when production seclabelstmt is exited.
func (s *BaseGaussdbParserListener) ExitSeclabelstmt(ctx *SeclabelstmtContext) {}

// EnterOpt_provider is called when production opt_provider is entered.
func (s *BaseGaussdbParserListener) EnterOpt_provider(ctx *Opt_providerContext) {}

// ExitOpt_provider is called when production opt_provider is exited.
func (s *BaseGaussdbParserListener) ExitOpt_provider(ctx *Opt_providerContext) {}

// EnterSecurity_label is called when production security_label is entered.
func (s *BaseGaussdbParserListener) EnterSecurity_label(ctx *Security_labelContext) {}

// ExitSecurity_label is called when production security_label is exited.
func (s *BaseGaussdbParserListener) ExitSecurity_label(ctx *Security_labelContext) {}

// EnterCursorstmt is called when production cursorstmt is entered.
func (s *BaseGaussdbParserListener) EnterCursorstmt(ctx *CursorstmtContext) {}

// ExitCursorstmt is called when production cursorstmt is exited.
func (s *BaseGaussdbParserListener) ExitCursorstmt(ctx *CursorstmtContext) {}

// EnterFetchstmt is called when production fetchstmt is entered.
func (s *BaseGaussdbParserListener) EnterFetchstmt(ctx *FetchstmtContext) {}

// ExitFetchstmt is called when production fetchstmt is exited.
func (s *BaseGaussdbParserListener) ExitFetchstmt(ctx *FetchstmtContext) {}

// EnterFetch_args is called when production fetch_args is entered.
func (s *BaseGaussdbParserListener) EnterFetch_args(ctx *Fetch_argsContext) {}

// ExitFetch_args is called when production fetch_args is exited.
func (s *BaseGaussdbParserListener) ExitFetch_args(ctx *Fetch_argsContext) {}

// EnterFrom_in is called when production from_in is entered.
func (s *BaseGaussdbParserListener) EnterFrom_in(ctx *From_inContext) {}

// ExitFrom_in is called when production from_in is exited.
func (s *BaseGaussdbParserListener) ExitFrom_in(ctx *From_inContext) {}

// EnterOpt_from_in is called when production opt_from_in is entered.
func (s *BaseGaussdbParserListener) EnterOpt_from_in(ctx *Opt_from_inContext) {}

// ExitOpt_from_in is called when production opt_from_in is exited.
func (s *BaseGaussdbParserListener) ExitOpt_from_in(ctx *Opt_from_inContext) {}

// EnterGrantstmt is called when production grantstmt is entered.
func (s *BaseGaussdbParserListener) EnterGrantstmt(ctx *GrantstmtContext) {}

// ExitGrantstmt is called when production grantstmt is exited.
func (s *BaseGaussdbParserListener) ExitGrantstmt(ctx *GrantstmtContext) {}

// EnterRevokestmt is called when production revokestmt is entered.
func (s *BaseGaussdbParserListener) EnterRevokestmt(ctx *RevokestmtContext) {}

// ExitRevokestmt is called when production revokestmt is exited.
func (s *BaseGaussdbParserListener) ExitRevokestmt(ctx *RevokestmtContext) {}

// EnterPrivileges is called when production privileges is entered.
func (s *BaseGaussdbParserListener) EnterPrivileges(ctx *PrivilegesContext) {}

// ExitPrivileges is called when production privileges is exited.
func (s *BaseGaussdbParserListener) ExitPrivileges(ctx *PrivilegesContext) {}

// EnterPrivilege_list is called when production privilege_list is entered.
func (s *BaseGaussdbParserListener) EnterPrivilege_list(ctx *Privilege_listContext) {}

// ExitPrivilege_list is called when production privilege_list is exited.
func (s *BaseGaussdbParserListener) ExitPrivilege_list(ctx *Privilege_listContext) {}

// EnterPrivilege is called when production privilege is entered.
func (s *BaseGaussdbParserListener) EnterPrivilege(ctx *PrivilegeContext) {}

// ExitPrivilege is called when production privilege is exited.
func (s *BaseGaussdbParserListener) ExitPrivilege(ctx *PrivilegeContext) {}

// EnterPrivilege_target is called when production privilege_target is entered.
func (s *BaseGaussdbParserListener) EnterPrivilege_target(ctx *Privilege_targetContext) {}

// ExitPrivilege_target is called when production privilege_target is exited.
func (s *BaseGaussdbParserListener) ExitPrivilege_target(ctx *Privilege_targetContext) {}

// EnterGrantee_list is called when production grantee_list is entered.
func (s *BaseGaussdbParserListener) EnterGrantee_list(ctx *Grantee_listContext) {}

// ExitGrantee_list is called when production grantee_list is exited.
func (s *BaseGaussdbParserListener) ExitGrantee_list(ctx *Grantee_listContext) {}

// EnterGrantee is called when production grantee is entered.
func (s *BaseGaussdbParserListener) EnterGrantee(ctx *GranteeContext) {}

// ExitGrantee is called when production grantee is exited.
func (s *BaseGaussdbParserListener) ExitGrantee(ctx *GranteeContext) {}

// EnterOpt_grant_grant_option is called when production opt_grant_grant_option is entered.
func (s *BaseGaussdbParserListener) EnterOpt_grant_grant_option(ctx *Opt_grant_grant_optionContext) {}

// ExitOpt_grant_grant_option is called when production opt_grant_grant_option is exited.
func (s *BaseGaussdbParserListener) ExitOpt_grant_grant_option(ctx *Opt_grant_grant_optionContext) {}

// EnterGrantrolestmt is called when production grantrolestmt is entered.
func (s *BaseGaussdbParserListener) EnterGrantrolestmt(ctx *GrantrolestmtContext) {}

// ExitGrantrolestmt is called when production grantrolestmt is exited.
func (s *BaseGaussdbParserListener) ExitGrantrolestmt(ctx *GrantrolestmtContext) {}

// EnterRevokerolestmt is called when production revokerolestmt is entered.
func (s *BaseGaussdbParserListener) EnterRevokerolestmt(ctx *RevokerolestmtContext) {}

// ExitRevokerolestmt is called when production revokerolestmt is exited.
func (s *BaseGaussdbParserListener) ExitRevokerolestmt(ctx *RevokerolestmtContext) {}

// EnterOpt_grant_admin_option is called when production opt_grant_admin_option is entered.
func (s *BaseGaussdbParserListener) EnterOpt_grant_admin_option(ctx *Opt_grant_admin_optionContext) {}

// ExitOpt_grant_admin_option is called when production opt_grant_admin_option is exited.
func (s *BaseGaussdbParserListener) ExitOpt_grant_admin_option(ctx *Opt_grant_admin_optionContext) {}

// EnterOpt_granted_by is called when production opt_granted_by is entered.
func (s *BaseGaussdbParserListener) EnterOpt_granted_by(ctx *Opt_granted_byContext) {}

// ExitOpt_granted_by is called when production opt_granted_by is exited.
func (s *BaseGaussdbParserListener) ExitOpt_granted_by(ctx *Opt_granted_byContext) {}

// EnterAlterdefaultprivilegesstmt is called when production alterdefaultprivilegesstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterdefaultprivilegesstmt(ctx *AlterdefaultprivilegesstmtContext) {
}

// ExitAlterdefaultprivilegesstmt is called when production alterdefaultprivilegesstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterdefaultprivilegesstmt(ctx *AlterdefaultprivilegesstmtContext) {
}

// EnterDefacloptionlist is called when production defacloptionlist is entered.
func (s *BaseGaussdbParserListener) EnterDefacloptionlist(ctx *DefacloptionlistContext) {}

// ExitDefacloptionlist is called when production defacloptionlist is exited.
func (s *BaseGaussdbParserListener) ExitDefacloptionlist(ctx *DefacloptionlistContext) {}

// EnterDefacloption is called when production defacloption is entered.
func (s *BaseGaussdbParserListener) EnterDefacloption(ctx *DefacloptionContext) {}

// ExitDefacloption is called when production defacloption is exited.
func (s *BaseGaussdbParserListener) ExitDefacloption(ctx *DefacloptionContext) {}

// EnterDefaclaction is called when production defaclaction is entered.
func (s *BaseGaussdbParserListener) EnterDefaclaction(ctx *DefaclactionContext) {}

// ExitDefaclaction is called when production defaclaction is exited.
func (s *BaseGaussdbParserListener) ExitDefaclaction(ctx *DefaclactionContext) {}

// EnterDefacl_privilege_target is called when production defacl_privilege_target is entered.
func (s *BaseGaussdbParserListener) EnterDefacl_privilege_target(ctx *Defacl_privilege_targetContext) {
}

// ExitDefacl_privilege_target is called when production defacl_privilege_target is exited.
func (s *BaseGaussdbParserListener) ExitDefacl_privilege_target(ctx *Defacl_privilege_targetContext) {
}

// EnterIndexstmt is called when production indexstmt is entered.
func (s *BaseGaussdbParserListener) EnterIndexstmt(ctx *IndexstmtContext) {}

// ExitIndexstmt is called when production indexstmt is exited.
func (s *BaseGaussdbParserListener) ExitIndexstmt(ctx *IndexstmtContext) {}

// EnterGaussdb_opt_distribute is called when production gaussdb_opt_distribute is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_opt_distribute(ctx *Gaussdb_opt_distributeContext) {}

// ExitGaussdb_opt_distribute is called when production gaussdb_opt_distribute is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_opt_distribute(ctx *Gaussdb_opt_distributeContext) {}

// EnterOpt_unique is called when production opt_unique is entered.
func (s *BaseGaussdbParserListener) EnterOpt_unique(ctx *Opt_uniqueContext) {}

// ExitOpt_unique is called when production opt_unique is exited.
func (s *BaseGaussdbParserListener) ExitOpt_unique(ctx *Opt_uniqueContext) {}

// EnterOpt_concurrently is called when production opt_concurrently is entered.
func (s *BaseGaussdbParserListener) EnterOpt_concurrently(ctx *Opt_concurrentlyContext) {}

// ExitOpt_concurrently is called when production opt_concurrently is exited.
func (s *BaseGaussdbParserListener) ExitOpt_concurrently(ctx *Opt_concurrentlyContext) {}

// EnterOpt_index_name is called when production opt_index_name is entered.
func (s *BaseGaussdbParserListener) EnterOpt_index_name(ctx *Opt_index_nameContext) {}

// ExitOpt_index_name is called when production opt_index_name is exited.
func (s *BaseGaussdbParserListener) ExitOpt_index_name(ctx *Opt_index_nameContext) {}

// EnterAccess_method_clause is called when production access_method_clause is entered.
func (s *BaseGaussdbParserListener) EnterAccess_method_clause(ctx *Access_method_clauseContext) {}

// ExitAccess_method_clause is called when production access_method_clause is exited.
func (s *BaseGaussdbParserListener) ExitAccess_method_clause(ctx *Access_method_clauseContext) {}

// EnterIndex_params is called when production index_params is entered.
func (s *BaseGaussdbParserListener) EnterIndex_params(ctx *Index_paramsContext) {}

// ExitIndex_params is called when production index_params is exited.
func (s *BaseGaussdbParserListener) ExitIndex_params(ctx *Index_paramsContext) {}

// EnterIndex_elem_options is called when production index_elem_options is entered.
func (s *BaseGaussdbParserListener) EnterIndex_elem_options(ctx *Index_elem_optionsContext) {}

// ExitIndex_elem_options is called when production index_elem_options is exited.
func (s *BaseGaussdbParserListener) ExitIndex_elem_options(ctx *Index_elem_optionsContext) {}

// EnterIndex_elem is called when production index_elem is entered.
func (s *BaseGaussdbParserListener) EnterIndex_elem(ctx *Index_elemContext) {}

// ExitIndex_elem is called when production index_elem is exited.
func (s *BaseGaussdbParserListener) ExitIndex_elem(ctx *Index_elemContext) {}

// EnterOpt_include is called when production opt_include is entered.
func (s *BaseGaussdbParserListener) EnterOpt_include(ctx *Opt_includeContext) {}

// ExitOpt_include is called when production opt_include is exited.
func (s *BaseGaussdbParserListener) ExitOpt_include(ctx *Opt_includeContext) {}

// EnterIndex_including_params is called when production index_including_params is entered.
func (s *BaseGaussdbParserListener) EnterIndex_including_params(ctx *Index_including_paramsContext) {}

// ExitIndex_including_params is called when production index_including_params is exited.
func (s *BaseGaussdbParserListener) ExitIndex_including_params(ctx *Index_including_paramsContext) {}

// EnterOpt_collate is called when production opt_collate is entered.
func (s *BaseGaussdbParserListener) EnterOpt_collate(ctx *Opt_collateContext) {}

// ExitOpt_collate is called when production opt_collate is exited.
func (s *BaseGaussdbParserListener) ExitOpt_collate(ctx *Opt_collateContext) {}

// EnterOpt_class is called when production opt_class is entered.
func (s *BaseGaussdbParserListener) EnterOpt_class(ctx *Opt_classContext) {}

// ExitOpt_class is called when production opt_class is exited.
func (s *BaseGaussdbParserListener) ExitOpt_class(ctx *Opt_classContext) {}

// EnterOpt_asc_desc is called when production opt_asc_desc is entered.
func (s *BaseGaussdbParserListener) EnterOpt_asc_desc(ctx *Opt_asc_descContext) {}

// ExitOpt_asc_desc is called when production opt_asc_desc is exited.
func (s *BaseGaussdbParserListener) ExitOpt_asc_desc(ctx *Opt_asc_descContext) {}

// EnterOpt_nulls_order is called when production opt_nulls_order is entered.
func (s *BaseGaussdbParserListener) EnterOpt_nulls_order(ctx *Opt_nulls_orderContext) {}

// ExitOpt_nulls_order is called when production opt_nulls_order is exited.
func (s *BaseGaussdbParserListener) ExitOpt_nulls_order(ctx *Opt_nulls_orderContext) {}

// EnterCreatefunctionstmt is called when production createfunctionstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatefunctionstmt(ctx *CreatefunctionstmtContext) {}

// ExitCreatefunctionstmt is called when production createfunctionstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatefunctionstmt(ctx *CreatefunctionstmtContext) {}

// EnterOpt_or_replace is called when production opt_or_replace is entered.
func (s *BaseGaussdbParserListener) EnterOpt_or_replace(ctx *Opt_or_replaceContext) {}

// ExitOpt_or_replace is called when production opt_or_replace is exited.
func (s *BaseGaussdbParserListener) ExitOpt_or_replace(ctx *Opt_or_replaceContext) {}

// EnterFunc_args is called when production func_args is entered.
func (s *BaseGaussdbParserListener) EnterFunc_args(ctx *Func_argsContext) {}

// ExitFunc_args is called when production func_args is exited.
func (s *BaseGaussdbParserListener) ExitFunc_args(ctx *Func_argsContext) {}

// EnterFunc_args_list is called when production func_args_list is entered.
func (s *BaseGaussdbParserListener) EnterFunc_args_list(ctx *Func_args_listContext) {}

// ExitFunc_args_list is called when production func_args_list is exited.
func (s *BaseGaussdbParserListener) ExitFunc_args_list(ctx *Func_args_listContext) {}

// EnterFunction_with_argtypes_list is called when production function_with_argtypes_list is entered.
func (s *BaseGaussdbParserListener) EnterFunction_with_argtypes_list(ctx *Function_with_argtypes_listContext) {
}

// ExitFunction_with_argtypes_list is called when production function_with_argtypes_list is exited.
func (s *BaseGaussdbParserListener) ExitFunction_with_argtypes_list(ctx *Function_with_argtypes_listContext) {
}

// EnterFunction_with_argtypes is called when production function_with_argtypes is entered.
func (s *BaseGaussdbParserListener) EnterFunction_with_argtypes(ctx *Function_with_argtypesContext) {}

// ExitFunction_with_argtypes is called when production function_with_argtypes is exited.
func (s *BaseGaussdbParserListener) ExitFunction_with_argtypes(ctx *Function_with_argtypesContext) {}

// EnterFunc_args_with_defaults is called when production func_args_with_defaults is entered.
func (s *BaseGaussdbParserListener) EnterFunc_args_with_defaults(ctx *Func_args_with_defaultsContext) {
}

// ExitFunc_args_with_defaults is called when production func_args_with_defaults is exited.
func (s *BaseGaussdbParserListener) ExitFunc_args_with_defaults(ctx *Func_args_with_defaultsContext) {
}

// EnterFunc_args_with_defaults_list is called when production func_args_with_defaults_list is entered.
func (s *BaseGaussdbParserListener) EnterFunc_args_with_defaults_list(ctx *Func_args_with_defaults_listContext) {
}

// ExitFunc_args_with_defaults_list is called when production func_args_with_defaults_list is exited.
func (s *BaseGaussdbParserListener) ExitFunc_args_with_defaults_list(ctx *Func_args_with_defaults_listContext) {
}

// EnterFunc_arg is called when production func_arg is entered.
func (s *BaseGaussdbParserListener) EnterFunc_arg(ctx *Func_argContext) {}

// ExitFunc_arg is called when production func_arg is exited.
func (s *BaseGaussdbParserListener) ExitFunc_arg(ctx *Func_argContext) {}

// EnterArg_class is called when production arg_class is entered.
func (s *BaseGaussdbParserListener) EnterArg_class(ctx *Arg_classContext) {}

// ExitArg_class is called when production arg_class is exited.
func (s *BaseGaussdbParserListener) ExitArg_class(ctx *Arg_classContext) {}

// EnterParam_name is called when production param_name is entered.
func (s *BaseGaussdbParserListener) EnterParam_name(ctx *Param_nameContext) {}

// ExitParam_name is called when production param_name is exited.
func (s *BaseGaussdbParserListener) ExitParam_name(ctx *Param_nameContext) {}

// EnterFunc_return is called when production func_return is entered.
func (s *BaseGaussdbParserListener) EnterFunc_return(ctx *Func_returnContext) {}

// ExitFunc_return is called when production func_return is exited.
func (s *BaseGaussdbParserListener) ExitFunc_return(ctx *Func_returnContext) {}

// EnterFunc_type is called when production func_type is entered.
func (s *BaseGaussdbParserListener) EnterFunc_type(ctx *Func_typeContext) {}

// ExitFunc_type is called when production func_type is exited.
func (s *BaseGaussdbParserListener) ExitFunc_type(ctx *Func_typeContext) {}

// EnterFunc_arg_with_default is called when production func_arg_with_default is entered.
func (s *BaseGaussdbParserListener) EnterFunc_arg_with_default(ctx *Func_arg_with_defaultContext) {}

// ExitFunc_arg_with_default is called when production func_arg_with_default is exited.
func (s *BaseGaussdbParserListener) ExitFunc_arg_with_default(ctx *Func_arg_with_defaultContext) {}

// EnterAggr_arg is called when production aggr_arg is entered.
func (s *BaseGaussdbParserListener) EnterAggr_arg(ctx *Aggr_argContext) {}

// ExitAggr_arg is called when production aggr_arg is exited.
func (s *BaseGaussdbParserListener) ExitAggr_arg(ctx *Aggr_argContext) {}

// EnterAggr_args is called when production aggr_args is entered.
func (s *BaseGaussdbParserListener) EnterAggr_args(ctx *Aggr_argsContext) {}

// ExitAggr_args is called when production aggr_args is exited.
func (s *BaseGaussdbParserListener) ExitAggr_args(ctx *Aggr_argsContext) {}

// EnterAggr_args_list is called when production aggr_args_list is entered.
func (s *BaseGaussdbParserListener) EnterAggr_args_list(ctx *Aggr_args_listContext) {}

// ExitAggr_args_list is called when production aggr_args_list is exited.
func (s *BaseGaussdbParserListener) ExitAggr_args_list(ctx *Aggr_args_listContext) {}

// EnterAggregate_with_argtypes is called when production aggregate_with_argtypes is entered.
func (s *BaseGaussdbParserListener) EnterAggregate_with_argtypes(ctx *Aggregate_with_argtypesContext) {
}

// ExitAggregate_with_argtypes is called when production aggregate_with_argtypes is exited.
func (s *BaseGaussdbParserListener) ExitAggregate_with_argtypes(ctx *Aggregate_with_argtypesContext) {
}

// EnterAggregate_with_argtypes_list is called when production aggregate_with_argtypes_list is entered.
func (s *BaseGaussdbParserListener) EnterAggregate_with_argtypes_list(ctx *Aggregate_with_argtypes_listContext) {
}

// ExitAggregate_with_argtypes_list is called when production aggregate_with_argtypes_list is exited.
func (s *BaseGaussdbParserListener) ExitAggregate_with_argtypes_list(ctx *Aggregate_with_argtypes_listContext) {
}

// EnterCreatefunc_opt_list is called when production createfunc_opt_list is entered.
func (s *BaseGaussdbParserListener) EnterCreatefunc_opt_list(ctx *Createfunc_opt_listContext) {}

// ExitCreatefunc_opt_list is called when production createfunc_opt_list is exited.
func (s *BaseGaussdbParserListener) ExitCreatefunc_opt_list(ctx *Createfunc_opt_listContext) {}

// EnterCommon_func_opt_item is called when production common_func_opt_item is entered.
func (s *BaseGaussdbParserListener) EnterCommon_func_opt_item(ctx *Common_func_opt_itemContext) {}

// ExitCommon_func_opt_item is called when production common_func_opt_item is exited.
func (s *BaseGaussdbParserListener) ExitCommon_func_opt_item(ctx *Common_func_opt_itemContext) {}

// EnterCreatefunc_opt_item is called when production createfunc_opt_item is entered.
func (s *BaseGaussdbParserListener) EnterCreatefunc_opt_item(ctx *Createfunc_opt_itemContext) {}

// ExitCreatefunc_opt_item is called when production createfunc_opt_item is exited.
func (s *BaseGaussdbParserListener) ExitCreatefunc_opt_item(ctx *Createfunc_opt_itemContext) {}

// EnterFunc_as is called when production func_as is entered.
func (s *BaseGaussdbParserListener) EnterFunc_as(ctx *Func_asContext) {}

// ExitFunc_as is called when production func_as is exited.
func (s *BaseGaussdbParserListener) ExitFunc_as(ctx *Func_asContext) {}

// EnterTransform_type_list is called when production transform_type_list is entered.
func (s *BaseGaussdbParserListener) EnterTransform_type_list(ctx *Transform_type_listContext) {}

// ExitTransform_type_list is called when production transform_type_list is exited.
func (s *BaseGaussdbParserListener) ExitTransform_type_list(ctx *Transform_type_listContext) {}

// EnterOpt_definition is called when production opt_definition is entered.
func (s *BaseGaussdbParserListener) EnterOpt_definition(ctx *Opt_definitionContext) {}

// ExitOpt_definition is called when production opt_definition is exited.
func (s *BaseGaussdbParserListener) ExitOpt_definition(ctx *Opt_definitionContext) {}

// EnterTable_func_column is called when production table_func_column is entered.
func (s *BaseGaussdbParserListener) EnterTable_func_column(ctx *Table_func_columnContext) {}

// ExitTable_func_column is called when production table_func_column is exited.
func (s *BaseGaussdbParserListener) ExitTable_func_column(ctx *Table_func_columnContext) {}

// EnterTable_func_column_list is called when production table_func_column_list is entered.
func (s *BaseGaussdbParserListener) EnterTable_func_column_list(ctx *Table_func_column_listContext) {}

// ExitTable_func_column_list is called when production table_func_column_list is exited.
func (s *BaseGaussdbParserListener) ExitTable_func_column_list(ctx *Table_func_column_listContext) {}

// EnterAlterfunctionstmt is called when production alterfunctionstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterfunctionstmt(ctx *AlterfunctionstmtContext) {}

// ExitAlterfunctionstmt is called when production alterfunctionstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterfunctionstmt(ctx *AlterfunctionstmtContext) {}

// EnterAlterfunc_opt_list is called when production alterfunc_opt_list is entered.
func (s *BaseGaussdbParserListener) EnterAlterfunc_opt_list(ctx *Alterfunc_opt_listContext) {}

// ExitAlterfunc_opt_list is called when production alterfunc_opt_list is exited.
func (s *BaseGaussdbParserListener) ExitAlterfunc_opt_list(ctx *Alterfunc_opt_listContext) {}

// EnterOpt_restrict is called when production opt_restrict is entered.
func (s *BaseGaussdbParserListener) EnterOpt_restrict(ctx *Opt_restrictContext) {}

// ExitOpt_restrict is called when production opt_restrict is exited.
func (s *BaseGaussdbParserListener) ExitOpt_restrict(ctx *Opt_restrictContext) {}

// EnterRemovefuncstmt is called when production removefuncstmt is entered.
func (s *BaseGaussdbParserListener) EnterRemovefuncstmt(ctx *RemovefuncstmtContext) {}

// ExitRemovefuncstmt is called when production removefuncstmt is exited.
func (s *BaseGaussdbParserListener) ExitRemovefuncstmt(ctx *RemovefuncstmtContext) {}

// EnterRemoveaggrstmt is called when production removeaggrstmt is entered.
func (s *BaseGaussdbParserListener) EnterRemoveaggrstmt(ctx *RemoveaggrstmtContext) {}

// ExitRemoveaggrstmt is called when production removeaggrstmt is exited.
func (s *BaseGaussdbParserListener) ExitRemoveaggrstmt(ctx *RemoveaggrstmtContext) {}

// EnterRemoveoperstmt is called when production removeoperstmt is entered.
func (s *BaseGaussdbParserListener) EnterRemoveoperstmt(ctx *RemoveoperstmtContext) {}

// ExitRemoveoperstmt is called when production removeoperstmt is exited.
func (s *BaseGaussdbParserListener) ExitRemoveoperstmt(ctx *RemoveoperstmtContext) {}

// EnterOper_argtypes is called when production oper_argtypes is entered.
func (s *BaseGaussdbParserListener) EnterOper_argtypes(ctx *Oper_argtypesContext) {}

// ExitOper_argtypes is called when production oper_argtypes is exited.
func (s *BaseGaussdbParserListener) ExitOper_argtypes(ctx *Oper_argtypesContext) {}

// EnterAny_operator is called when production any_operator is entered.
func (s *BaseGaussdbParserListener) EnterAny_operator(ctx *Any_operatorContext) {}

// ExitAny_operator is called when production any_operator is exited.
func (s *BaseGaussdbParserListener) ExitAny_operator(ctx *Any_operatorContext) {}

// EnterOperator_with_argtypes_list is called when production operator_with_argtypes_list is entered.
func (s *BaseGaussdbParserListener) EnterOperator_with_argtypes_list(ctx *Operator_with_argtypes_listContext) {
}

// ExitOperator_with_argtypes_list is called when production operator_with_argtypes_list is exited.
func (s *BaseGaussdbParserListener) ExitOperator_with_argtypes_list(ctx *Operator_with_argtypes_listContext) {
}

// EnterOperator_with_argtypes is called when production operator_with_argtypes is entered.
func (s *BaseGaussdbParserListener) EnterOperator_with_argtypes(ctx *Operator_with_argtypesContext) {}

// ExitOperator_with_argtypes is called when production operator_with_argtypes is exited.
func (s *BaseGaussdbParserListener) ExitOperator_with_argtypes(ctx *Operator_with_argtypesContext) {}

// EnterDostmt is called when production dostmt is entered.
func (s *BaseGaussdbParserListener) EnterDostmt(ctx *DostmtContext) {}

// ExitDostmt is called when production dostmt is exited.
func (s *BaseGaussdbParserListener) ExitDostmt(ctx *DostmtContext) {}

// EnterDostmt_opt_list is called when production dostmt_opt_list is entered.
func (s *BaseGaussdbParserListener) EnterDostmt_opt_list(ctx *Dostmt_opt_listContext) {}

// ExitDostmt_opt_list is called when production dostmt_opt_list is exited.
func (s *BaseGaussdbParserListener) ExitDostmt_opt_list(ctx *Dostmt_opt_listContext) {}

// EnterDostmt_opt_item is called when production dostmt_opt_item is entered.
func (s *BaseGaussdbParserListener) EnterDostmt_opt_item(ctx *Dostmt_opt_itemContext) {}

// ExitDostmt_opt_item is called when production dostmt_opt_item is exited.
func (s *BaseGaussdbParserListener) ExitDostmt_opt_item(ctx *Dostmt_opt_itemContext) {}

// EnterCreatecaststmt is called when production createcaststmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatecaststmt(ctx *CreatecaststmtContext) {}

// ExitCreatecaststmt is called when production createcaststmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatecaststmt(ctx *CreatecaststmtContext) {}

// EnterCast_context is called when production cast_context is entered.
func (s *BaseGaussdbParserListener) EnterCast_context(ctx *Cast_contextContext) {}

// ExitCast_context is called when production cast_context is exited.
func (s *BaseGaussdbParserListener) ExitCast_context(ctx *Cast_contextContext) {}

// EnterDropcaststmt is called when production dropcaststmt is entered.
func (s *BaseGaussdbParserListener) EnterDropcaststmt(ctx *DropcaststmtContext) {}

// ExitDropcaststmt is called when production dropcaststmt is exited.
func (s *BaseGaussdbParserListener) ExitDropcaststmt(ctx *DropcaststmtContext) {}

// EnterOpt_if_exists is called when production opt_if_exists is entered.
func (s *BaseGaussdbParserListener) EnterOpt_if_exists(ctx *Opt_if_existsContext) {}

// ExitOpt_if_exists is called when production opt_if_exists is exited.
func (s *BaseGaussdbParserListener) ExitOpt_if_exists(ctx *Opt_if_existsContext) {}

// EnterCreatetransformstmt is called when production createtransformstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatetransformstmt(ctx *CreatetransformstmtContext) {}

// ExitCreatetransformstmt is called when production createtransformstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatetransformstmt(ctx *CreatetransformstmtContext) {}

// EnterTransform_element_list is called when production transform_element_list is entered.
func (s *BaseGaussdbParserListener) EnterTransform_element_list(ctx *Transform_element_listContext) {}

// ExitTransform_element_list is called when production transform_element_list is exited.
func (s *BaseGaussdbParserListener) ExitTransform_element_list(ctx *Transform_element_listContext) {}

// EnterDroptransformstmt is called when production droptransformstmt is entered.
func (s *BaseGaussdbParserListener) EnterDroptransformstmt(ctx *DroptransformstmtContext) {}

// ExitDroptransformstmt is called when production droptransformstmt is exited.
func (s *BaseGaussdbParserListener) ExitDroptransformstmt(ctx *DroptransformstmtContext) {}

// EnterReindexstmt is called when production reindexstmt is entered.
func (s *BaseGaussdbParserListener) EnterReindexstmt(ctx *ReindexstmtContext) {}

// ExitReindexstmt is called when production reindexstmt is exited.
func (s *BaseGaussdbParserListener) ExitReindexstmt(ctx *ReindexstmtContext) {}

// EnterReindex_target_type is called when production reindex_target_type is entered.
func (s *BaseGaussdbParserListener) EnterReindex_target_type(ctx *Reindex_target_typeContext) {}

// ExitReindex_target_type is called when production reindex_target_type is exited.
func (s *BaseGaussdbParserListener) ExitReindex_target_type(ctx *Reindex_target_typeContext) {}

// EnterReindex_target_multitable is called when production reindex_target_multitable is entered.
func (s *BaseGaussdbParserListener) EnterReindex_target_multitable(ctx *Reindex_target_multitableContext) {
}

// ExitReindex_target_multitable is called when production reindex_target_multitable is exited.
func (s *BaseGaussdbParserListener) ExitReindex_target_multitable(ctx *Reindex_target_multitableContext) {
}

// EnterReindex_option_list is called when production reindex_option_list is entered.
func (s *BaseGaussdbParserListener) EnterReindex_option_list(ctx *Reindex_option_listContext) {}

// ExitReindex_option_list is called when production reindex_option_list is exited.
func (s *BaseGaussdbParserListener) ExitReindex_option_list(ctx *Reindex_option_listContext) {}

// EnterReindex_option_elem is called when production reindex_option_elem is entered.
func (s *BaseGaussdbParserListener) EnterReindex_option_elem(ctx *Reindex_option_elemContext) {}

// ExitReindex_option_elem is called when production reindex_option_elem is exited.
func (s *BaseGaussdbParserListener) ExitReindex_option_elem(ctx *Reindex_option_elemContext) {}

// EnterAltertblspcstmt is called when production altertblspcstmt is entered.
func (s *BaseGaussdbParserListener) EnterAltertblspcstmt(ctx *AltertblspcstmtContext) {}

// ExitAltertblspcstmt is called when production altertblspcstmt is exited.
func (s *BaseGaussdbParserListener) ExitAltertblspcstmt(ctx *AltertblspcstmtContext) {}

// EnterRenamestmt is called when production renamestmt is entered.
func (s *BaseGaussdbParserListener) EnterRenamestmt(ctx *RenamestmtContext) {}

// ExitRenamestmt is called when production renamestmt is exited.
func (s *BaseGaussdbParserListener) ExitRenamestmt(ctx *RenamestmtContext) {}

// EnterOpt_column is called when production opt_column is entered.
func (s *BaseGaussdbParserListener) EnterOpt_column(ctx *Opt_columnContext) {}

// ExitOpt_column is called when production opt_column is exited.
func (s *BaseGaussdbParserListener) ExitOpt_column(ctx *Opt_columnContext) {}

// EnterOpt_set_data is called when production opt_set_data is entered.
func (s *BaseGaussdbParserListener) EnterOpt_set_data(ctx *Opt_set_dataContext) {}

// ExitOpt_set_data is called when production opt_set_data is exited.
func (s *BaseGaussdbParserListener) ExitOpt_set_data(ctx *Opt_set_dataContext) {}

// EnterAlterobjectdependsstmt is called when production alterobjectdependsstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterobjectdependsstmt(ctx *AlterobjectdependsstmtContext) {}

// ExitAlterobjectdependsstmt is called when production alterobjectdependsstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterobjectdependsstmt(ctx *AlterobjectdependsstmtContext) {}

// EnterOpt_no is called when production opt_no is entered.
func (s *BaseGaussdbParserListener) EnterOpt_no(ctx *Opt_noContext) {}

// ExitOpt_no is called when production opt_no is exited.
func (s *BaseGaussdbParserListener) ExitOpt_no(ctx *Opt_noContext) {}

// EnterAlterobjectschemastmt is called when production alterobjectschemastmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterobjectschemastmt(ctx *AlterobjectschemastmtContext) {}

// ExitAlterobjectschemastmt is called when production alterobjectschemastmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterobjectschemastmt(ctx *AlterobjectschemastmtContext) {}

// EnterAlteroperatorstmt is called when production alteroperatorstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlteroperatorstmt(ctx *AlteroperatorstmtContext) {}

// ExitAlteroperatorstmt is called when production alteroperatorstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlteroperatorstmt(ctx *AlteroperatorstmtContext) {}

// EnterOperator_def_list is called when production operator_def_list is entered.
func (s *BaseGaussdbParserListener) EnterOperator_def_list(ctx *Operator_def_listContext) {}

// ExitOperator_def_list is called when production operator_def_list is exited.
func (s *BaseGaussdbParserListener) ExitOperator_def_list(ctx *Operator_def_listContext) {}

// EnterOperator_def_elem is called when production operator_def_elem is entered.
func (s *BaseGaussdbParserListener) EnterOperator_def_elem(ctx *Operator_def_elemContext) {}

// ExitOperator_def_elem is called when production operator_def_elem is exited.
func (s *BaseGaussdbParserListener) ExitOperator_def_elem(ctx *Operator_def_elemContext) {}

// EnterOperator_def_arg is called when production operator_def_arg is entered.
func (s *BaseGaussdbParserListener) EnterOperator_def_arg(ctx *Operator_def_argContext) {}

// ExitOperator_def_arg is called when production operator_def_arg is exited.
func (s *BaseGaussdbParserListener) ExitOperator_def_arg(ctx *Operator_def_argContext) {}

// EnterAltertypestmt is called when production altertypestmt is entered.
func (s *BaseGaussdbParserListener) EnterAltertypestmt(ctx *AltertypestmtContext) {}

// ExitAltertypestmt is called when production altertypestmt is exited.
func (s *BaseGaussdbParserListener) ExitAltertypestmt(ctx *AltertypestmtContext) {}

// EnterAlterownerstmt is called when production alterownerstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterownerstmt(ctx *AlterownerstmtContext) {}

// ExitAlterownerstmt is called when production alterownerstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterownerstmt(ctx *AlterownerstmtContext) {}

// EnterCreatepublicationstmt is called when production createpublicationstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatepublicationstmt(ctx *CreatepublicationstmtContext) {}

// ExitCreatepublicationstmt is called when production createpublicationstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatepublicationstmt(ctx *CreatepublicationstmtContext) {}

// EnterOpt_publication_for_tables is called when production opt_publication_for_tables is entered.
func (s *BaseGaussdbParserListener) EnterOpt_publication_for_tables(ctx *Opt_publication_for_tablesContext) {
}

// ExitOpt_publication_for_tables is called when production opt_publication_for_tables is exited.
func (s *BaseGaussdbParserListener) ExitOpt_publication_for_tables(ctx *Opt_publication_for_tablesContext) {
}

// EnterPublication_for_tables is called when production publication_for_tables is entered.
func (s *BaseGaussdbParserListener) EnterPublication_for_tables(ctx *Publication_for_tablesContext) {}

// ExitPublication_for_tables is called when production publication_for_tables is exited.
func (s *BaseGaussdbParserListener) ExitPublication_for_tables(ctx *Publication_for_tablesContext) {}

// EnterAlterpublicationstmt is called when production alterpublicationstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterpublicationstmt(ctx *AlterpublicationstmtContext) {}

// ExitAlterpublicationstmt is called when production alterpublicationstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterpublicationstmt(ctx *AlterpublicationstmtContext) {}

// EnterCreatesubscriptionstmt is called when production createsubscriptionstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatesubscriptionstmt(ctx *CreatesubscriptionstmtContext) {}

// ExitCreatesubscriptionstmt is called when production createsubscriptionstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatesubscriptionstmt(ctx *CreatesubscriptionstmtContext) {}

// EnterPublication_name_list is called when production publication_name_list is entered.
func (s *BaseGaussdbParserListener) EnterPublication_name_list(ctx *Publication_name_listContext) {}

// ExitPublication_name_list is called when production publication_name_list is exited.
func (s *BaseGaussdbParserListener) ExitPublication_name_list(ctx *Publication_name_listContext) {}

// EnterPublication_name_item is called when production publication_name_item is entered.
func (s *BaseGaussdbParserListener) EnterPublication_name_item(ctx *Publication_name_itemContext) {}

// ExitPublication_name_item is called when production publication_name_item is exited.
func (s *BaseGaussdbParserListener) ExitPublication_name_item(ctx *Publication_name_itemContext) {}

// EnterAltersubscriptionstmt is called when production altersubscriptionstmt is entered.
func (s *BaseGaussdbParserListener) EnterAltersubscriptionstmt(ctx *AltersubscriptionstmtContext) {}

// ExitAltersubscriptionstmt is called when production altersubscriptionstmt is exited.
func (s *BaseGaussdbParserListener) ExitAltersubscriptionstmt(ctx *AltersubscriptionstmtContext) {}

// EnterDropsubscriptionstmt is called when production dropsubscriptionstmt is entered.
func (s *BaseGaussdbParserListener) EnterDropsubscriptionstmt(ctx *DropsubscriptionstmtContext) {}

// ExitDropsubscriptionstmt is called when production dropsubscriptionstmt is exited.
func (s *BaseGaussdbParserListener) ExitDropsubscriptionstmt(ctx *DropsubscriptionstmtContext) {}

// EnterRulestmt is called when production rulestmt is entered.
func (s *BaseGaussdbParserListener) EnterRulestmt(ctx *RulestmtContext) {}

// ExitRulestmt is called when production rulestmt is exited.
func (s *BaseGaussdbParserListener) ExitRulestmt(ctx *RulestmtContext) {}

// EnterRuleactionlist is called when production ruleactionlist is entered.
func (s *BaseGaussdbParserListener) EnterRuleactionlist(ctx *RuleactionlistContext) {}

// ExitRuleactionlist is called when production ruleactionlist is exited.
func (s *BaseGaussdbParserListener) ExitRuleactionlist(ctx *RuleactionlistContext) {}

// EnterRuleactionmulti is called when production ruleactionmulti is entered.
func (s *BaseGaussdbParserListener) EnterRuleactionmulti(ctx *RuleactionmultiContext) {}

// ExitRuleactionmulti is called when production ruleactionmulti is exited.
func (s *BaseGaussdbParserListener) ExitRuleactionmulti(ctx *RuleactionmultiContext) {}

// EnterRuleactionstmt is called when production ruleactionstmt is entered.
func (s *BaseGaussdbParserListener) EnterRuleactionstmt(ctx *RuleactionstmtContext) {}

// ExitRuleactionstmt is called when production ruleactionstmt is exited.
func (s *BaseGaussdbParserListener) ExitRuleactionstmt(ctx *RuleactionstmtContext) {}

// EnterRuleactionstmtOrEmpty is called when production ruleactionstmtOrEmpty is entered.
func (s *BaseGaussdbParserListener) EnterRuleactionstmtOrEmpty(ctx *RuleactionstmtOrEmptyContext) {}

// ExitRuleactionstmtOrEmpty is called when production ruleactionstmtOrEmpty is exited.
func (s *BaseGaussdbParserListener) ExitRuleactionstmtOrEmpty(ctx *RuleactionstmtOrEmptyContext) {}

// EnterEvent is called when production event is entered.
func (s *BaseGaussdbParserListener) EnterEvent(ctx *EventContext) {}

// ExitEvent is called when production event is exited.
func (s *BaseGaussdbParserListener) ExitEvent(ctx *EventContext) {}

// EnterOpt_instead is called when production opt_instead is entered.
func (s *BaseGaussdbParserListener) EnterOpt_instead(ctx *Opt_insteadContext) {}

// ExitOpt_instead is called when production opt_instead is exited.
func (s *BaseGaussdbParserListener) ExitOpt_instead(ctx *Opt_insteadContext) {}

// EnterNotifystmt is called when production notifystmt is entered.
func (s *BaseGaussdbParserListener) EnterNotifystmt(ctx *NotifystmtContext) {}

// ExitNotifystmt is called when production notifystmt is exited.
func (s *BaseGaussdbParserListener) ExitNotifystmt(ctx *NotifystmtContext) {}

// EnterNotify_payload is called when production notify_payload is entered.
func (s *BaseGaussdbParserListener) EnterNotify_payload(ctx *Notify_payloadContext) {}

// ExitNotify_payload is called when production notify_payload is exited.
func (s *BaseGaussdbParserListener) ExitNotify_payload(ctx *Notify_payloadContext) {}

// EnterListenstmt is called when production listenstmt is entered.
func (s *BaseGaussdbParserListener) EnterListenstmt(ctx *ListenstmtContext) {}

// ExitListenstmt is called when production listenstmt is exited.
func (s *BaseGaussdbParserListener) ExitListenstmt(ctx *ListenstmtContext) {}

// EnterUnlistenstmt is called when production unlistenstmt is entered.
func (s *BaseGaussdbParserListener) EnterUnlistenstmt(ctx *UnlistenstmtContext) {}

// ExitUnlistenstmt is called when production unlistenstmt is exited.
func (s *BaseGaussdbParserListener) ExitUnlistenstmt(ctx *UnlistenstmtContext) {}

// EnterTransactionstmt is called when production transactionstmt is entered.
func (s *BaseGaussdbParserListener) EnterTransactionstmt(ctx *TransactionstmtContext) {}

// ExitTransactionstmt is called when production transactionstmt is exited.
func (s *BaseGaussdbParserListener) ExitTransactionstmt(ctx *TransactionstmtContext) {}

// EnterOpt_transaction is called when production opt_transaction is entered.
func (s *BaseGaussdbParserListener) EnterOpt_transaction(ctx *Opt_transactionContext) {}

// ExitOpt_transaction is called when production opt_transaction is exited.
func (s *BaseGaussdbParserListener) ExitOpt_transaction(ctx *Opt_transactionContext) {}

// EnterTransaction_mode_item is called when production transaction_mode_item is entered.
func (s *BaseGaussdbParserListener) EnterTransaction_mode_item(ctx *Transaction_mode_itemContext) {}

// ExitTransaction_mode_item is called when production transaction_mode_item is exited.
func (s *BaseGaussdbParserListener) ExitTransaction_mode_item(ctx *Transaction_mode_itemContext) {}

// EnterTransaction_mode_list is called when production transaction_mode_list is entered.
func (s *BaseGaussdbParserListener) EnterTransaction_mode_list(ctx *Transaction_mode_listContext) {}

// ExitTransaction_mode_list is called when production transaction_mode_list is exited.
func (s *BaseGaussdbParserListener) ExitTransaction_mode_list(ctx *Transaction_mode_listContext) {}

// EnterTransaction_mode_list_or_empty is called when production transaction_mode_list_or_empty is entered.
func (s *BaseGaussdbParserListener) EnterTransaction_mode_list_or_empty(ctx *Transaction_mode_list_or_emptyContext) {
}

// ExitTransaction_mode_list_or_empty is called when production transaction_mode_list_or_empty is exited.
func (s *BaseGaussdbParserListener) ExitTransaction_mode_list_or_empty(ctx *Transaction_mode_list_or_emptyContext) {
}

// EnterOpt_transaction_chain is called when production opt_transaction_chain is entered.
func (s *BaseGaussdbParserListener) EnterOpt_transaction_chain(ctx *Opt_transaction_chainContext) {}

// ExitOpt_transaction_chain is called when production opt_transaction_chain is exited.
func (s *BaseGaussdbParserListener) ExitOpt_transaction_chain(ctx *Opt_transaction_chainContext) {}

// EnterViewstmt is called when production viewstmt is entered.
func (s *BaseGaussdbParserListener) EnterViewstmt(ctx *ViewstmtContext) {}

// ExitViewstmt is called when production viewstmt is exited.
func (s *BaseGaussdbParserListener) ExitViewstmt(ctx *ViewstmtContext) {}

// EnterOpt_check_option is called when production opt_check_option is entered.
func (s *BaseGaussdbParserListener) EnterOpt_check_option(ctx *Opt_check_optionContext) {}

// ExitOpt_check_option is called when production opt_check_option is exited.
func (s *BaseGaussdbParserListener) ExitOpt_check_option(ctx *Opt_check_optionContext) {}

// EnterLoadstmt is called when production loadstmt is entered.
func (s *BaseGaussdbParserListener) EnterLoadstmt(ctx *LoadstmtContext) {}

// ExitLoadstmt is called when production loadstmt is exited.
func (s *BaseGaussdbParserListener) ExitLoadstmt(ctx *LoadstmtContext) {}

// EnterCreatedbstmt is called when production createdbstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatedbstmt(ctx *CreatedbstmtContext) {}

// ExitCreatedbstmt is called when production createdbstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatedbstmt(ctx *CreatedbstmtContext) {}

// EnterCreatedb_opt_list is called when production createdb_opt_list is entered.
func (s *BaseGaussdbParserListener) EnterCreatedb_opt_list(ctx *Createdb_opt_listContext) {}

// ExitCreatedb_opt_list is called when production createdb_opt_list is exited.
func (s *BaseGaussdbParserListener) ExitCreatedb_opt_list(ctx *Createdb_opt_listContext) {}

// EnterCreatedb_opt_items is called when production createdb_opt_items is entered.
func (s *BaseGaussdbParserListener) EnterCreatedb_opt_items(ctx *Createdb_opt_itemsContext) {}

// ExitCreatedb_opt_items is called when production createdb_opt_items is exited.
func (s *BaseGaussdbParserListener) ExitCreatedb_opt_items(ctx *Createdb_opt_itemsContext) {}

// EnterCreatedb_opt_item is called when production createdb_opt_item is entered.
func (s *BaseGaussdbParserListener) EnterCreatedb_opt_item(ctx *Createdb_opt_itemContext) {}

// ExitCreatedb_opt_item is called when production createdb_opt_item is exited.
func (s *BaseGaussdbParserListener) ExitCreatedb_opt_item(ctx *Createdb_opt_itemContext) {}

// EnterCreatedb_opt_name is called when production createdb_opt_name is entered.
func (s *BaseGaussdbParserListener) EnterCreatedb_opt_name(ctx *Createdb_opt_nameContext) {}

// ExitCreatedb_opt_name is called when production createdb_opt_name is exited.
func (s *BaseGaussdbParserListener) ExitCreatedb_opt_name(ctx *Createdb_opt_nameContext) {}

// EnterOpt_equal is called when production opt_equal is entered.
func (s *BaseGaussdbParserListener) EnterOpt_equal(ctx *Opt_equalContext) {}

// ExitOpt_equal is called when production opt_equal is exited.
func (s *BaseGaussdbParserListener) ExitOpt_equal(ctx *Opt_equalContext) {}

// EnterAlterdatabasestmt is called when production alterdatabasestmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterdatabasestmt(ctx *AlterdatabasestmtContext) {}

// ExitAlterdatabasestmt is called when production alterdatabasestmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterdatabasestmt(ctx *AlterdatabasestmtContext) {}

// EnterAlterdatabasesetstmt is called when production alterdatabasesetstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterdatabasesetstmt(ctx *AlterdatabasesetstmtContext) {}

// ExitAlterdatabasesetstmt is called when production alterdatabasesetstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterdatabasesetstmt(ctx *AlterdatabasesetstmtContext) {}

// EnterDropdbstmt is called when production dropdbstmt is entered.
func (s *BaseGaussdbParserListener) EnterDropdbstmt(ctx *DropdbstmtContext) {}

// ExitDropdbstmt is called when production dropdbstmt is exited.
func (s *BaseGaussdbParserListener) ExitDropdbstmt(ctx *DropdbstmtContext) {}

// EnterDrop_option_list is called when production drop_option_list is entered.
func (s *BaseGaussdbParserListener) EnterDrop_option_list(ctx *Drop_option_listContext) {}

// ExitDrop_option_list is called when production drop_option_list is exited.
func (s *BaseGaussdbParserListener) ExitDrop_option_list(ctx *Drop_option_listContext) {}

// EnterDrop_option is called when production drop_option is entered.
func (s *BaseGaussdbParserListener) EnterDrop_option(ctx *Drop_optionContext) {}

// ExitDrop_option is called when production drop_option is exited.
func (s *BaseGaussdbParserListener) ExitDrop_option(ctx *Drop_optionContext) {}

// EnterAltercollationstmt is called when production altercollationstmt is entered.
func (s *BaseGaussdbParserListener) EnterAltercollationstmt(ctx *AltercollationstmtContext) {}

// ExitAltercollationstmt is called when production altercollationstmt is exited.
func (s *BaseGaussdbParserListener) ExitAltercollationstmt(ctx *AltercollationstmtContext) {}

// EnterAltersystemstmt is called when production altersystemstmt is entered.
func (s *BaseGaussdbParserListener) EnterAltersystemstmt(ctx *AltersystemstmtContext) {}

// ExitAltersystemstmt is called when production altersystemstmt is exited.
func (s *BaseGaussdbParserListener) ExitAltersystemstmt(ctx *AltersystemstmtContext) {}

// EnterCreatedomainstmt is called when production createdomainstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreatedomainstmt(ctx *CreatedomainstmtContext) {}

// ExitCreatedomainstmt is called when production createdomainstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreatedomainstmt(ctx *CreatedomainstmtContext) {}

// EnterAlterdomainstmt is called when production alterdomainstmt is entered.
func (s *BaseGaussdbParserListener) EnterAlterdomainstmt(ctx *AlterdomainstmtContext) {}

// ExitAlterdomainstmt is called when production alterdomainstmt is exited.
func (s *BaseGaussdbParserListener) ExitAlterdomainstmt(ctx *AlterdomainstmtContext) {}

// EnterOpt_as is called when production opt_as is entered.
func (s *BaseGaussdbParserListener) EnterOpt_as(ctx *Opt_asContext) {}

// ExitOpt_as is called when production opt_as is exited.
func (s *BaseGaussdbParserListener) ExitOpt_as(ctx *Opt_asContext) {}

// EnterAltertsdictionarystmt is called when production altertsdictionarystmt is entered.
func (s *BaseGaussdbParserListener) EnterAltertsdictionarystmt(ctx *AltertsdictionarystmtContext) {}

// ExitAltertsdictionarystmt is called when production altertsdictionarystmt is exited.
func (s *BaseGaussdbParserListener) ExitAltertsdictionarystmt(ctx *AltertsdictionarystmtContext) {}

// EnterAltertsconfigurationstmt is called when production altertsconfigurationstmt is entered.
func (s *BaseGaussdbParserListener) EnterAltertsconfigurationstmt(ctx *AltertsconfigurationstmtContext) {
}

// ExitAltertsconfigurationstmt is called when production altertsconfigurationstmt is exited.
func (s *BaseGaussdbParserListener) ExitAltertsconfigurationstmt(ctx *AltertsconfigurationstmtContext) {
}

// EnterAny_with is called when production any_with is entered.
func (s *BaseGaussdbParserListener) EnterAny_with(ctx *Any_withContext) {}

// ExitAny_with is called when production any_with is exited.
func (s *BaseGaussdbParserListener) ExitAny_with(ctx *Any_withContext) {}

// EnterCreateconversionstmt is called when production createconversionstmt is entered.
func (s *BaseGaussdbParserListener) EnterCreateconversionstmt(ctx *CreateconversionstmtContext) {}

// ExitCreateconversionstmt is called when production createconversionstmt is exited.
func (s *BaseGaussdbParserListener) ExitCreateconversionstmt(ctx *CreateconversionstmtContext) {}

// EnterClusterstmt is called when production clusterstmt is entered.
func (s *BaseGaussdbParserListener) EnterClusterstmt(ctx *ClusterstmtContext) {}

// ExitClusterstmt is called when production clusterstmt is exited.
func (s *BaseGaussdbParserListener) ExitClusterstmt(ctx *ClusterstmtContext) {}

// EnterCluster_index_specification is called when production cluster_index_specification is entered.
func (s *BaseGaussdbParserListener) EnterCluster_index_specification(ctx *Cluster_index_specificationContext) {
}

// ExitCluster_index_specification is called when production cluster_index_specification is exited.
func (s *BaseGaussdbParserListener) ExitCluster_index_specification(ctx *Cluster_index_specificationContext) {
}

// EnterVacuumstmt is called when production vacuumstmt is entered.
func (s *BaseGaussdbParserListener) EnterVacuumstmt(ctx *VacuumstmtContext) {}

// ExitVacuumstmt is called when production vacuumstmt is exited.
func (s *BaseGaussdbParserListener) ExitVacuumstmt(ctx *VacuumstmtContext) {}

// EnterAnalyzestmt is called when production analyzestmt is entered.
func (s *BaseGaussdbParserListener) EnterAnalyzestmt(ctx *AnalyzestmtContext) {}

// ExitAnalyzestmt is called when production analyzestmt is exited.
func (s *BaseGaussdbParserListener) ExitAnalyzestmt(ctx *AnalyzestmtContext) {}

// EnterVac_analyze_option_list is called when production vac_analyze_option_list is entered.
func (s *BaseGaussdbParserListener) EnterVac_analyze_option_list(ctx *Vac_analyze_option_listContext) {
}

// ExitVac_analyze_option_list is called when production vac_analyze_option_list is exited.
func (s *BaseGaussdbParserListener) ExitVac_analyze_option_list(ctx *Vac_analyze_option_listContext) {
}

// EnterAnalyze_keyword is called when production analyze_keyword is entered.
func (s *BaseGaussdbParserListener) EnterAnalyze_keyword(ctx *Analyze_keywordContext) {}

// ExitAnalyze_keyword is called when production analyze_keyword is exited.
func (s *BaseGaussdbParserListener) ExitAnalyze_keyword(ctx *Analyze_keywordContext) {}

// EnterVac_analyze_option_elem is called when production vac_analyze_option_elem is entered.
func (s *BaseGaussdbParserListener) EnterVac_analyze_option_elem(ctx *Vac_analyze_option_elemContext) {
}

// ExitVac_analyze_option_elem is called when production vac_analyze_option_elem is exited.
func (s *BaseGaussdbParserListener) ExitVac_analyze_option_elem(ctx *Vac_analyze_option_elemContext) {
}

// EnterVac_analyze_option_name is called when production vac_analyze_option_name is entered.
func (s *BaseGaussdbParserListener) EnterVac_analyze_option_name(ctx *Vac_analyze_option_nameContext) {
}

// ExitVac_analyze_option_name is called when production vac_analyze_option_name is exited.
func (s *BaseGaussdbParserListener) ExitVac_analyze_option_name(ctx *Vac_analyze_option_nameContext) {
}

// EnterVac_analyze_option_arg is called when production vac_analyze_option_arg is entered.
func (s *BaseGaussdbParserListener) EnterVac_analyze_option_arg(ctx *Vac_analyze_option_argContext) {}

// ExitVac_analyze_option_arg is called when production vac_analyze_option_arg is exited.
func (s *BaseGaussdbParserListener) ExitVac_analyze_option_arg(ctx *Vac_analyze_option_argContext) {}

// EnterOpt_analyze is called when production opt_analyze is entered.
func (s *BaseGaussdbParserListener) EnterOpt_analyze(ctx *Opt_analyzeContext) {}

// ExitOpt_analyze is called when production opt_analyze is exited.
func (s *BaseGaussdbParserListener) ExitOpt_analyze(ctx *Opt_analyzeContext) {}

// EnterOpt_verbose is called when production opt_verbose is entered.
func (s *BaseGaussdbParserListener) EnterOpt_verbose(ctx *Opt_verboseContext) {}

// ExitOpt_verbose is called when production opt_verbose is exited.
func (s *BaseGaussdbParserListener) ExitOpt_verbose(ctx *Opt_verboseContext) {}

// EnterOpt_full is called when production opt_full is entered.
func (s *BaseGaussdbParserListener) EnterOpt_full(ctx *Opt_fullContext) {}

// ExitOpt_full is called when production opt_full is exited.
func (s *BaseGaussdbParserListener) ExitOpt_full(ctx *Opt_fullContext) {}

// EnterOpt_freeze is called when production opt_freeze is entered.
func (s *BaseGaussdbParserListener) EnterOpt_freeze(ctx *Opt_freezeContext) {}

// ExitOpt_freeze is called when production opt_freeze is exited.
func (s *BaseGaussdbParserListener) ExitOpt_freeze(ctx *Opt_freezeContext) {}

// EnterOpt_name_list is called when production opt_name_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_name_list(ctx *Opt_name_listContext) {}

// ExitOpt_name_list is called when production opt_name_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_name_list(ctx *Opt_name_listContext) {}

// EnterVacuum_relation is called when production vacuum_relation is entered.
func (s *BaseGaussdbParserListener) EnterVacuum_relation(ctx *Vacuum_relationContext) {}

// ExitVacuum_relation is called when production vacuum_relation is exited.
func (s *BaseGaussdbParserListener) ExitVacuum_relation(ctx *Vacuum_relationContext) {}

// EnterVacuum_relation_list is called when production vacuum_relation_list is entered.
func (s *BaseGaussdbParserListener) EnterVacuum_relation_list(ctx *Vacuum_relation_listContext) {}

// ExitVacuum_relation_list is called when production vacuum_relation_list is exited.
func (s *BaseGaussdbParserListener) ExitVacuum_relation_list(ctx *Vacuum_relation_listContext) {}

// EnterOpt_vacuum_relation_list is called when production opt_vacuum_relation_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_vacuum_relation_list(ctx *Opt_vacuum_relation_listContext) {
}

// ExitOpt_vacuum_relation_list is called when production opt_vacuum_relation_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_vacuum_relation_list(ctx *Opt_vacuum_relation_listContext) {
}

// EnterExplainplanstmt is called when production explainplanstmt is entered.
func (s *BaseGaussdbParserListener) EnterExplainplanstmt(ctx *ExplainplanstmtContext) {}

// ExitExplainplanstmt is called when production explainplanstmt is exited.
func (s *BaseGaussdbParserListener) ExitExplainplanstmt(ctx *ExplainplanstmtContext) {}

// EnterExplainstmt is called when production explainstmt is entered.
func (s *BaseGaussdbParserListener) EnterExplainstmt(ctx *ExplainstmtContext) {}

// ExitExplainstmt is called when production explainstmt is exited.
func (s *BaseGaussdbParserListener) ExitExplainstmt(ctx *ExplainstmtContext) {}

// EnterExplainablestmt is called when production explainablestmt is entered.
func (s *BaseGaussdbParserListener) EnterExplainablestmt(ctx *ExplainablestmtContext) {}

// ExitExplainablestmt is called when production explainablestmt is exited.
func (s *BaseGaussdbParserListener) ExitExplainablestmt(ctx *ExplainablestmtContext) {}

// EnterExplain_option_list is called when production explain_option_list is entered.
func (s *BaseGaussdbParserListener) EnterExplain_option_list(ctx *Explain_option_listContext) {}

// ExitExplain_option_list is called when production explain_option_list is exited.
func (s *BaseGaussdbParserListener) ExitExplain_option_list(ctx *Explain_option_listContext) {}

// EnterExplain_option_elem is called when production explain_option_elem is entered.
func (s *BaseGaussdbParserListener) EnterExplain_option_elem(ctx *Explain_option_elemContext) {}

// ExitExplain_option_elem is called when production explain_option_elem is exited.
func (s *BaseGaussdbParserListener) ExitExplain_option_elem(ctx *Explain_option_elemContext) {}

// EnterExplain_option_name is called when production explain_option_name is entered.
func (s *BaseGaussdbParserListener) EnterExplain_option_name(ctx *Explain_option_nameContext) {}

// ExitExplain_option_name is called when production explain_option_name is exited.
func (s *BaseGaussdbParserListener) ExitExplain_option_name(ctx *Explain_option_nameContext) {}

// EnterExplain_option_arg is called when production explain_option_arg is entered.
func (s *BaseGaussdbParserListener) EnterExplain_option_arg(ctx *Explain_option_argContext) {}

// ExitExplain_option_arg is called when production explain_option_arg is exited.
func (s *BaseGaussdbParserListener) ExitExplain_option_arg(ctx *Explain_option_argContext) {}

// EnterPreparestmt is called when production preparestmt is entered.
func (s *BaseGaussdbParserListener) EnterPreparestmt(ctx *PreparestmtContext) {}

// ExitPreparestmt is called when production preparestmt is exited.
func (s *BaseGaussdbParserListener) ExitPreparestmt(ctx *PreparestmtContext) {}

// EnterPrep_type_clause is called when production prep_type_clause is entered.
func (s *BaseGaussdbParserListener) EnterPrep_type_clause(ctx *Prep_type_clauseContext) {}

// ExitPrep_type_clause is called when production prep_type_clause is exited.
func (s *BaseGaussdbParserListener) ExitPrep_type_clause(ctx *Prep_type_clauseContext) {}

// EnterPreparablestmt is called when production preparablestmt is entered.
func (s *BaseGaussdbParserListener) EnterPreparablestmt(ctx *PreparablestmtContext) {}

// ExitPreparablestmt is called when production preparablestmt is exited.
func (s *BaseGaussdbParserListener) ExitPreparablestmt(ctx *PreparablestmtContext) {}

// EnterExecutestmt is called when production executestmt is entered.
func (s *BaseGaussdbParserListener) EnterExecutestmt(ctx *ExecutestmtContext) {}

// ExitExecutestmt is called when production executestmt is exited.
func (s *BaseGaussdbParserListener) ExitExecutestmt(ctx *ExecutestmtContext) {}

// EnterExecute_param_clause is called when production execute_param_clause is entered.
func (s *BaseGaussdbParserListener) EnterExecute_param_clause(ctx *Execute_param_clauseContext) {}

// ExitExecute_param_clause is called when production execute_param_clause is exited.
func (s *BaseGaussdbParserListener) ExitExecute_param_clause(ctx *Execute_param_clauseContext) {}

// EnterDeallocatestmt is called when production deallocatestmt is entered.
func (s *BaseGaussdbParserListener) EnterDeallocatestmt(ctx *DeallocatestmtContext) {}

// ExitDeallocatestmt is called when production deallocatestmt is exited.
func (s *BaseGaussdbParserListener) ExitDeallocatestmt(ctx *DeallocatestmtContext) {}

// EnterInsertstmt is called when production insertstmt is entered.
func (s *BaseGaussdbParserListener) EnterInsertstmt(ctx *InsertstmtContext) {}

// ExitInsertstmt is called when production insertstmt is exited.
func (s *BaseGaussdbParserListener) ExitInsertstmt(ctx *InsertstmtContext) {}

// EnterInsert_target is called when production insert_target is entered.
func (s *BaseGaussdbParserListener) EnterInsert_target(ctx *Insert_targetContext) {}

// ExitInsert_target is called when production insert_target is exited.
func (s *BaseGaussdbParserListener) ExitInsert_target(ctx *Insert_targetContext) {}

// EnterInsert_rest is called when production insert_rest is entered.
func (s *BaseGaussdbParserListener) EnterInsert_rest(ctx *Insert_restContext) {}

// ExitInsert_rest is called when production insert_rest is exited.
func (s *BaseGaussdbParserListener) ExitInsert_rest(ctx *Insert_restContext) {}

// EnterOverride_kind is called when production override_kind is entered.
func (s *BaseGaussdbParserListener) EnterOverride_kind(ctx *Override_kindContext) {}

// ExitOverride_kind is called when production override_kind is exited.
func (s *BaseGaussdbParserListener) ExitOverride_kind(ctx *Override_kindContext) {}

// EnterInsert_column_list is called when production insert_column_list is entered.
func (s *BaseGaussdbParserListener) EnterInsert_column_list(ctx *Insert_column_listContext) {}

// ExitInsert_column_list is called when production insert_column_list is exited.
func (s *BaseGaussdbParserListener) ExitInsert_column_list(ctx *Insert_column_listContext) {}

// EnterInsert_column_item is called when production insert_column_item is entered.
func (s *BaseGaussdbParserListener) EnterInsert_column_item(ctx *Insert_column_itemContext) {}

// ExitInsert_column_item is called when production insert_column_item is exited.
func (s *BaseGaussdbParserListener) ExitInsert_column_item(ctx *Insert_column_itemContext) {}

// EnterOpt_on_conflict is called when production opt_on_conflict is entered.
func (s *BaseGaussdbParserListener) EnterOpt_on_conflict(ctx *Opt_on_conflictContext) {}

// ExitOpt_on_conflict is called when production opt_on_conflict is exited.
func (s *BaseGaussdbParserListener) ExitOpt_on_conflict(ctx *Opt_on_conflictContext) {}

// EnterOpt_conf_expr is called when production opt_conf_expr is entered.
func (s *BaseGaussdbParserListener) EnterOpt_conf_expr(ctx *Opt_conf_exprContext) {}

// ExitOpt_conf_expr is called when production opt_conf_expr is exited.
func (s *BaseGaussdbParserListener) ExitOpt_conf_expr(ctx *Opt_conf_exprContext) {}

// EnterReturning_clause is called when production returning_clause is entered.
func (s *BaseGaussdbParserListener) EnterReturning_clause(ctx *Returning_clauseContext) {}

// ExitReturning_clause is called when production returning_clause is exited.
func (s *BaseGaussdbParserListener) ExitReturning_clause(ctx *Returning_clauseContext) {}

// EnterMergestmt is called when production mergestmt is entered.
func (s *BaseGaussdbParserListener) EnterMergestmt(ctx *MergestmtContext) {}

// ExitMergestmt is called when production mergestmt is exited.
func (s *BaseGaussdbParserListener) ExitMergestmt(ctx *MergestmtContext) {}

// EnterMerge_insert_clause is called when production merge_insert_clause is entered.
func (s *BaseGaussdbParserListener) EnterMerge_insert_clause(ctx *Merge_insert_clauseContext) {}

// ExitMerge_insert_clause is called when production merge_insert_clause is exited.
func (s *BaseGaussdbParserListener) ExitMerge_insert_clause(ctx *Merge_insert_clauseContext) {}

// EnterMerge_update_clause is called when production merge_update_clause is entered.
func (s *BaseGaussdbParserListener) EnterMerge_update_clause(ctx *Merge_update_clauseContext) {}

// ExitMerge_update_clause is called when production merge_update_clause is exited.
func (s *BaseGaussdbParserListener) ExitMerge_update_clause(ctx *Merge_update_clauseContext) {}

// EnterMerge_delete_clause is called when production merge_delete_clause is entered.
func (s *BaseGaussdbParserListener) EnterMerge_delete_clause(ctx *Merge_delete_clauseContext) {}

// ExitMerge_delete_clause is called when production merge_delete_clause is exited.
func (s *BaseGaussdbParserListener) ExitMerge_delete_clause(ctx *Merge_delete_clauseContext) {}

// EnterDeletestmt is called when production deletestmt is entered.
func (s *BaseGaussdbParserListener) EnterDeletestmt(ctx *DeletestmtContext) {}

// ExitDeletestmt is called when production deletestmt is exited.
func (s *BaseGaussdbParserListener) ExitDeletestmt(ctx *DeletestmtContext) {}

// EnterUsing_clause is called when production using_clause is entered.
func (s *BaseGaussdbParserListener) EnterUsing_clause(ctx *Using_clauseContext) {}

// ExitUsing_clause is called when production using_clause is exited.
func (s *BaseGaussdbParserListener) ExitUsing_clause(ctx *Using_clauseContext) {}

// EnterLockstmt is called when production lockstmt is entered.
func (s *BaseGaussdbParserListener) EnterLockstmt(ctx *LockstmtContext) {}

// ExitLockstmt is called when production lockstmt is exited.
func (s *BaseGaussdbParserListener) ExitLockstmt(ctx *LockstmtContext) {}

// EnterLockbucketsstmt is called when production lockbucketsstmt is entered.
func (s *BaseGaussdbParserListener) EnterLockbucketsstmt(ctx *LockbucketsstmtContext) {}

// ExitLockbucketsstmt is called when production lockbucketsstmt is exited.
func (s *BaseGaussdbParserListener) ExitLockbucketsstmt(ctx *LockbucketsstmtContext) {}

// EnterMarkbucketsstmt is called when production markbucketsstmt is entered.
func (s *BaseGaussdbParserListener) EnterMarkbucketsstmt(ctx *MarkbucketsstmtContext) {}

// ExitMarkbucketsstmt is called when production markbucketsstmt is exited.
func (s *BaseGaussdbParserListener) ExitMarkbucketsstmt(ctx *MarkbucketsstmtContext) {}

// EnterOpt_bucketlist is called when production opt_bucketlist is entered.
func (s *BaseGaussdbParserListener) EnterOpt_bucketlist(ctx *Opt_bucketlistContext) {}

// ExitOpt_bucketlist is called when production opt_bucketlist is exited.
func (s *BaseGaussdbParserListener) ExitOpt_bucketlist(ctx *Opt_bucketlistContext) {}

// EnterOpt_lock is called when production opt_lock is entered.
func (s *BaseGaussdbParserListener) EnterOpt_lock(ctx *Opt_lockContext) {}

// ExitOpt_lock is called when production opt_lock is exited.
func (s *BaseGaussdbParserListener) ExitOpt_lock(ctx *Opt_lockContext) {}

// EnterLock_type is called when production lock_type is entered.
func (s *BaseGaussdbParserListener) EnterLock_type(ctx *Lock_typeContext) {}

// ExitLock_type is called when production lock_type is exited.
func (s *BaseGaussdbParserListener) ExitLock_type(ctx *Lock_typeContext) {}

// EnterOpt_nowait is called when production opt_nowait is entered.
func (s *BaseGaussdbParserListener) EnterOpt_nowait(ctx *Opt_nowaitContext) {}

// ExitOpt_nowait is called when production opt_nowait is exited.
func (s *BaseGaussdbParserListener) ExitOpt_nowait(ctx *Opt_nowaitContext) {}

// EnterOpt_nowait_or_skip is called when production opt_nowait_or_skip is entered.
func (s *BaseGaussdbParserListener) EnterOpt_nowait_or_skip(ctx *Opt_nowait_or_skipContext) {}

// ExitOpt_nowait_or_skip is called when production opt_nowait_or_skip is exited.
func (s *BaseGaussdbParserListener) ExitOpt_nowait_or_skip(ctx *Opt_nowait_or_skipContext) {}

// EnterUpdatestmt is called when production updatestmt is entered.
func (s *BaseGaussdbParserListener) EnterUpdatestmt(ctx *UpdatestmtContext) {}

// ExitUpdatestmt is called when production updatestmt is exited.
func (s *BaseGaussdbParserListener) ExitUpdatestmt(ctx *UpdatestmtContext) {}

// EnterSet_clause_list is called when production set_clause_list is entered.
func (s *BaseGaussdbParserListener) EnterSet_clause_list(ctx *Set_clause_listContext) {}

// ExitSet_clause_list is called when production set_clause_list is exited.
func (s *BaseGaussdbParserListener) ExitSet_clause_list(ctx *Set_clause_listContext) {}

// EnterSet_clause is called when production set_clause is entered.
func (s *BaseGaussdbParserListener) EnterSet_clause(ctx *Set_clauseContext) {}

// ExitSet_clause is called when production set_clause is exited.
func (s *BaseGaussdbParserListener) ExitSet_clause(ctx *Set_clauseContext) {}

// EnterSet_target is called when production set_target is entered.
func (s *BaseGaussdbParserListener) EnterSet_target(ctx *Set_targetContext) {}

// ExitSet_target is called when production set_target is exited.
func (s *BaseGaussdbParserListener) ExitSet_target(ctx *Set_targetContext) {}

// EnterSet_target_list is called when production set_target_list is entered.
func (s *BaseGaussdbParserListener) EnterSet_target_list(ctx *Set_target_listContext) {}

// ExitSet_target_list is called when production set_target_list is exited.
func (s *BaseGaussdbParserListener) ExitSet_target_list(ctx *Set_target_listContext) {}

// EnterDeclarecursorstmt is called when production declarecursorstmt is entered.
func (s *BaseGaussdbParserListener) EnterDeclarecursorstmt(ctx *DeclarecursorstmtContext) {}

// ExitDeclarecursorstmt is called when production declarecursorstmt is exited.
func (s *BaseGaussdbParserListener) ExitDeclarecursorstmt(ctx *DeclarecursorstmtContext) {}

// EnterCursor_name is called when production cursor_name is entered.
func (s *BaseGaussdbParserListener) EnterCursor_name(ctx *Cursor_nameContext) {}

// ExitCursor_name is called when production cursor_name is exited.
func (s *BaseGaussdbParserListener) ExitCursor_name(ctx *Cursor_nameContext) {}

// EnterCursor_options is called when production cursor_options is entered.
func (s *BaseGaussdbParserListener) EnterCursor_options(ctx *Cursor_optionsContext) {}

// ExitCursor_options is called when production cursor_options is exited.
func (s *BaseGaussdbParserListener) ExitCursor_options(ctx *Cursor_optionsContext) {}

// EnterOpt_hold is called when production opt_hold is entered.
func (s *BaseGaussdbParserListener) EnterOpt_hold(ctx *Opt_holdContext) {}

// ExitOpt_hold is called when production opt_hold is exited.
func (s *BaseGaussdbParserListener) ExitOpt_hold(ctx *Opt_holdContext) {}

// EnterSelectstmt is called when production selectstmt is entered.
func (s *BaseGaussdbParserListener) EnterSelectstmt(ctx *SelectstmtContext) {}

// ExitSelectstmt is called when production selectstmt is exited.
func (s *BaseGaussdbParserListener) ExitSelectstmt(ctx *SelectstmtContext) {}

// EnterSelect_with_parens is called when production select_with_parens is entered.
func (s *BaseGaussdbParserListener) EnterSelect_with_parens(ctx *Select_with_parensContext) {}

// ExitSelect_with_parens is called when production select_with_parens is exited.
func (s *BaseGaussdbParserListener) ExitSelect_with_parens(ctx *Select_with_parensContext) {}

// EnterSelect_no_parens is called when production select_no_parens is entered.
func (s *BaseGaussdbParserListener) EnterSelect_no_parens(ctx *Select_no_parensContext) {}

// ExitSelect_no_parens is called when production select_no_parens is exited.
func (s *BaseGaussdbParserListener) ExitSelect_no_parens(ctx *Select_no_parensContext) {}

// EnterSelect_clause is called when production select_clause is entered.
func (s *BaseGaussdbParserListener) EnterSelect_clause(ctx *Select_clauseContext) {}

// ExitSelect_clause is called when production select_clause is exited.
func (s *BaseGaussdbParserListener) ExitSelect_clause(ctx *Select_clauseContext) {}

// EnterSimple_select_intersect is called when production simple_select_intersect is entered.
func (s *BaseGaussdbParserListener) EnterSimple_select_intersect(ctx *Simple_select_intersectContext) {
}

// ExitSimple_select_intersect is called when production simple_select_intersect is exited.
func (s *BaseGaussdbParserListener) ExitSimple_select_intersect(ctx *Simple_select_intersectContext) {
}

// EnterSimple_select_pramary is called when production simple_select_pramary is entered.
func (s *BaseGaussdbParserListener) EnterSimple_select_pramary(ctx *Simple_select_pramaryContext) {}

// ExitSimple_select_pramary is called when production simple_select_pramary is exited.
func (s *BaseGaussdbParserListener) ExitSimple_select_pramary(ctx *Simple_select_pramaryContext) {}

// EnterWith_clause is called when production with_clause is entered.
func (s *BaseGaussdbParserListener) EnterWith_clause(ctx *With_clauseContext) {}

// ExitWith_clause is called when production with_clause is exited.
func (s *BaseGaussdbParserListener) ExitWith_clause(ctx *With_clauseContext) {}

// EnterCte_list is called when production cte_list is entered.
func (s *BaseGaussdbParserListener) EnterCte_list(ctx *Cte_listContext) {}

// ExitCte_list is called when production cte_list is exited.
func (s *BaseGaussdbParserListener) ExitCte_list(ctx *Cte_listContext) {}

// EnterCommon_table_expr is called when production common_table_expr is entered.
func (s *BaseGaussdbParserListener) EnterCommon_table_expr(ctx *Common_table_exprContext) {}

// ExitCommon_table_expr is called when production common_table_expr is exited.
func (s *BaseGaussdbParserListener) ExitCommon_table_expr(ctx *Common_table_exprContext) {}

// EnterOpt_materialized is called when production opt_materialized is entered.
func (s *BaseGaussdbParserListener) EnterOpt_materialized(ctx *Opt_materializedContext) {}

// ExitOpt_materialized is called when production opt_materialized is exited.
func (s *BaseGaussdbParserListener) ExitOpt_materialized(ctx *Opt_materializedContext) {}

// EnterOpt_with_clause is called when production opt_with_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_with_clause(ctx *Opt_with_clauseContext) {}

// ExitOpt_with_clause is called when production opt_with_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_with_clause(ctx *Opt_with_clauseContext) {}

// EnterInto_clause is called when production into_clause is entered.
func (s *BaseGaussdbParserListener) EnterInto_clause(ctx *Into_clauseContext) {}

// ExitInto_clause is called when production into_clause is exited.
func (s *BaseGaussdbParserListener) ExitInto_clause(ctx *Into_clauseContext) {}

// EnterOpt_strict is called when production opt_strict is entered.
func (s *BaseGaussdbParserListener) EnterOpt_strict(ctx *Opt_strictContext) {}

// ExitOpt_strict is called when production opt_strict is exited.
func (s *BaseGaussdbParserListener) ExitOpt_strict(ctx *Opt_strictContext) {}

// EnterOpttempTableName is called when production opttempTableName is entered.
func (s *BaseGaussdbParserListener) EnterOpttempTableName(ctx *OpttempTableNameContext) {}

// ExitOpttempTableName is called when production opttempTableName is exited.
func (s *BaseGaussdbParserListener) ExitOpttempTableName(ctx *OpttempTableNameContext) {}

// EnterOpt_table is called when production opt_table is entered.
func (s *BaseGaussdbParserListener) EnterOpt_table(ctx *Opt_tableContext) {}

// ExitOpt_table is called when production opt_table is exited.
func (s *BaseGaussdbParserListener) ExitOpt_table(ctx *Opt_tableContext) {}

// EnterAll_or_distinct is called when production all_or_distinct is entered.
func (s *BaseGaussdbParserListener) EnterAll_or_distinct(ctx *All_or_distinctContext) {}

// ExitAll_or_distinct is called when production all_or_distinct is exited.
func (s *BaseGaussdbParserListener) ExitAll_or_distinct(ctx *All_or_distinctContext) {}

// EnterDistinct_clause is called when production distinct_clause is entered.
func (s *BaseGaussdbParserListener) EnterDistinct_clause(ctx *Distinct_clauseContext) {}

// ExitDistinct_clause is called when production distinct_clause is exited.
func (s *BaseGaussdbParserListener) ExitDistinct_clause(ctx *Distinct_clauseContext) {}

// EnterOpt_all_clause is called when production opt_all_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_all_clause(ctx *Opt_all_clauseContext) {}

// ExitOpt_all_clause is called when production opt_all_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_all_clause(ctx *Opt_all_clauseContext) {}

// EnterOpt_sort_clause is called when production opt_sort_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_sort_clause(ctx *Opt_sort_clauseContext) {}

// ExitOpt_sort_clause is called when production opt_sort_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_sort_clause(ctx *Opt_sort_clauseContext) {}

// EnterSort_clause is called when production sort_clause is entered.
func (s *BaseGaussdbParserListener) EnterSort_clause(ctx *Sort_clauseContext) {}

// ExitSort_clause is called when production sort_clause is exited.
func (s *BaseGaussdbParserListener) ExitSort_clause(ctx *Sort_clauseContext) {}

// EnterSortby_list is called when production sortby_list is entered.
func (s *BaseGaussdbParserListener) EnterSortby_list(ctx *Sortby_listContext) {}

// ExitSortby_list is called when production sortby_list is exited.
func (s *BaseGaussdbParserListener) ExitSortby_list(ctx *Sortby_listContext) {}

// EnterSortby is called when production sortby is entered.
func (s *BaseGaussdbParserListener) EnterSortby(ctx *SortbyContext) {}

// ExitSortby is called when production sortby is exited.
func (s *BaseGaussdbParserListener) ExitSortby(ctx *SortbyContext) {}

// EnterSelect_limit is called when production select_limit is entered.
func (s *BaseGaussdbParserListener) EnterSelect_limit(ctx *Select_limitContext) {}

// ExitSelect_limit is called when production select_limit is exited.
func (s *BaseGaussdbParserListener) ExitSelect_limit(ctx *Select_limitContext) {}

// EnterOpt_select_limit is called when production opt_select_limit is entered.
func (s *BaseGaussdbParserListener) EnterOpt_select_limit(ctx *Opt_select_limitContext) {}

// ExitOpt_select_limit is called when production opt_select_limit is exited.
func (s *BaseGaussdbParserListener) ExitOpt_select_limit(ctx *Opt_select_limitContext) {}

// EnterLimit_clause is called when production limit_clause is entered.
func (s *BaseGaussdbParserListener) EnterLimit_clause(ctx *Limit_clauseContext) {}

// ExitLimit_clause is called when production limit_clause is exited.
func (s *BaseGaussdbParserListener) ExitLimit_clause(ctx *Limit_clauseContext) {}

// EnterOffset_clause is called when production offset_clause is entered.
func (s *BaseGaussdbParserListener) EnterOffset_clause(ctx *Offset_clauseContext) {}

// ExitOffset_clause is called when production offset_clause is exited.
func (s *BaseGaussdbParserListener) ExitOffset_clause(ctx *Offset_clauseContext) {}

// EnterSelect_limit_value is called when production select_limit_value is entered.
func (s *BaseGaussdbParserListener) EnterSelect_limit_value(ctx *Select_limit_valueContext) {}

// ExitSelect_limit_value is called when production select_limit_value is exited.
func (s *BaseGaussdbParserListener) ExitSelect_limit_value(ctx *Select_limit_valueContext) {}

// EnterSelect_offset_value is called when production select_offset_value is entered.
func (s *BaseGaussdbParserListener) EnterSelect_offset_value(ctx *Select_offset_valueContext) {}

// ExitSelect_offset_value is called when production select_offset_value is exited.
func (s *BaseGaussdbParserListener) ExitSelect_offset_value(ctx *Select_offset_valueContext) {}

// EnterSelect_fetch_first_value is called when production select_fetch_first_value is entered.
func (s *BaseGaussdbParserListener) EnterSelect_fetch_first_value(ctx *Select_fetch_first_valueContext) {
}

// ExitSelect_fetch_first_value is called when production select_fetch_first_value is exited.
func (s *BaseGaussdbParserListener) ExitSelect_fetch_first_value(ctx *Select_fetch_first_valueContext) {
}

// EnterI_or_f_const is called when production i_or_f_const is entered.
func (s *BaseGaussdbParserListener) EnterI_or_f_const(ctx *I_or_f_constContext) {}

// ExitI_or_f_const is called when production i_or_f_const is exited.
func (s *BaseGaussdbParserListener) ExitI_or_f_const(ctx *I_or_f_constContext) {}

// EnterRow_or_rows is called when production row_or_rows is entered.
func (s *BaseGaussdbParserListener) EnterRow_or_rows(ctx *Row_or_rowsContext) {}

// ExitRow_or_rows is called when production row_or_rows is exited.
func (s *BaseGaussdbParserListener) ExitRow_or_rows(ctx *Row_or_rowsContext) {}

// EnterFirst_or_next is called when production first_or_next is entered.
func (s *BaseGaussdbParserListener) EnterFirst_or_next(ctx *First_or_nextContext) {}

// ExitFirst_or_next is called when production first_or_next is exited.
func (s *BaseGaussdbParserListener) ExitFirst_or_next(ctx *First_or_nextContext) {}

// EnterGroup_clause is called when production group_clause is entered.
func (s *BaseGaussdbParserListener) EnterGroup_clause(ctx *Group_clauseContext) {}

// ExitGroup_clause is called when production group_clause is exited.
func (s *BaseGaussdbParserListener) ExitGroup_clause(ctx *Group_clauseContext) {}

// EnterGroup_by_list is called when production group_by_list is entered.
func (s *BaseGaussdbParserListener) EnterGroup_by_list(ctx *Group_by_listContext) {}

// ExitGroup_by_list is called when production group_by_list is exited.
func (s *BaseGaussdbParserListener) ExitGroup_by_list(ctx *Group_by_listContext) {}

// EnterGroup_by_item is called when production group_by_item is entered.
func (s *BaseGaussdbParserListener) EnterGroup_by_item(ctx *Group_by_itemContext) {}

// ExitGroup_by_item is called when production group_by_item is exited.
func (s *BaseGaussdbParserListener) ExitGroup_by_item(ctx *Group_by_itemContext) {}

// EnterEmpty_grouping_set is called when production empty_grouping_set is entered.
func (s *BaseGaussdbParserListener) EnterEmpty_grouping_set(ctx *Empty_grouping_setContext) {}

// ExitEmpty_grouping_set is called when production empty_grouping_set is exited.
func (s *BaseGaussdbParserListener) ExitEmpty_grouping_set(ctx *Empty_grouping_setContext) {}

// EnterRollup_clause is called when production rollup_clause is entered.
func (s *BaseGaussdbParserListener) EnterRollup_clause(ctx *Rollup_clauseContext) {}

// ExitRollup_clause is called when production rollup_clause is exited.
func (s *BaseGaussdbParserListener) ExitRollup_clause(ctx *Rollup_clauseContext) {}

// EnterCube_clause is called when production cube_clause is entered.
func (s *BaseGaussdbParserListener) EnterCube_clause(ctx *Cube_clauseContext) {}

// ExitCube_clause is called when production cube_clause is exited.
func (s *BaseGaussdbParserListener) ExitCube_clause(ctx *Cube_clauseContext) {}

// EnterGrouping_sets_clause is called when production grouping_sets_clause is entered.
func (s *BaseGaussdbParserListener) EnterGrouping_sets_clause(ctx *Grouping_sets_clauseContext) {}

// ExitGrouping_sets_clause is called when production grouping_sets_clause is exited.
func (s *BaseGaussdbParserListener) ExitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BaseGaussdbParserListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BaseGaussdbParserListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterFor_locking_clause is called when production for_locking_clause is entered.
func (s *BaseGaussdbParserListener) EnterFor_locking_clause(ctx *For_locking_clauseContext) {}

// ExitFor_locking_clause is called when production for_locking_clause is exited.
func (s *BaseGaussdbParserListener) ExitFor_locking_clause(ctx *For_locking_clauseContext) {}

// EnterOpt_for_locking_clause is called when production opt_for_locking_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_for_locking_clause(ctx *Opt_for_locking_clauseContext) {}

// ExitOpt_for_locking_clause is called when production opt_for_locking_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_for_locking_clause(ctx *Opt_for_locking_clauseContext) {}

// EnterFor_locking_items is called when production for_locking_items is entered.
func (s *BaseGaussdbParserListener) EnterFor_locking_items(ctx *For_locking_itemsContext) {}

// ExitFor_locking_items is called when production for_locking_items is exited.
func (s *BaseGaussdbParserListener) ExitFor_locking_items(ctx *For_locking_itemsContext) {}

// EnterFor_locking_item is called when production for_locking_item is entered.
func (s *BaseGaussdbParserListener) EnterFor_locking_item(ctx *For_locking_itemContext) {}

// ExitFor_locking_item is called when production for_locking_item is exited.
func (s *BaseGaussdbParserListener) ExitFor_locking_item(ctx *For_locking_itemContext) {}

// EnterFor_locking_strength is called when production for_locking_strength is entered.
func (s *BaseGaussdbParserListener) EnterFor_locking_strength(ctx *For_locking_strengthContext) {}

// ExitFor_locking_strength is called when production for_locking_strength is exited.
func (s *BaseGaussdbParserListener) ExitFor_locking_strength(ctx *For_locking_strengthContext) {}

// EnterLocked_rels_list is called when production locked_rels_list is entered.
func (s *BaseGaussdbParserListener) EnterLocked_rels_list(ctx *Locked_rels_listContext) {}

// ExitLocked_rels_list is called when production locked_rels_list is exited.
func (s *BaseGaussdbParserListener) ExitLocked_rels_list(ctx *Locked_rels_listContext) {}

// EnterValues_clause is called when production values_clause is entered.
func (s *BaseGaussdbParserListener) EnterValues_clause(ctx *Values_clauseContext) {}

// ExitValues_clause is called when production values_clause is exited.
func (s *BaseGaussdbParserListener) ExitValues_clause(ctx *Values_clauseContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseGaussdbParserListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseGaussdbParserListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterFrom_list is called when production from_list is entered.
func (s *BaseGaussdbParserListener) EnterFrom_list(ctx *From_listContext) {}

// ExitFrom_list is called when production from_list is exited.
func (s *BaseGaussdbParserListener) ExitFrom_list(ctx *From_listContext) {}

// EnterTable_ref is called when production table_ref is entered.
func (s *BaseGaussdbParserListener) EnterTable_ref(ctx *Table_refContext) {}

// ExitTable_ref is called when production table_ref is exited.
func (s *BaseGaussdbParserListener) ExitTable_ref(ctx *Table_refContext) {}

// EnterJoined_table is called when production joined_table is entered.
func (s *BaseGaussdbParserListener) EnterJoined_table(ctx *Joined_tableContext) {}

// ExitJoined_table is called when production joined_table is exited.
func (s *BaseGaussdbParserListener) ExitJoined_table(ctx *Joined_tableContext) {}

// EnterAlias_clause is called when production alias_clause is entered.
func (s *BaseGaussdbParserListener) EnterAlias_clause(ctx *Alias_clauseContext) {}

// ExitAlias_clause is called when production alias_clause is exited.
func (s *BaseGaussdbParserListener) ExitAlias_clause(ctx *Alias_clauseContext) {}

// EnterOpt_alias_clause is called when production opt_alias_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_alias_clause(ctx *Opt_alias_clauseContext) {}

// ExitOpt_alias_clause is called when production opt_alias_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_alias_clause(ctx *Opt_alias_clauseContext) {}

// EnterTable_alias_clause is called when production table_alias_clause is entered.
func (s *BaseGaussdbParserListener) EnterTable_alias_clause(ctx *Table_alias_clauseContext) {}

// ExitTable_alias_clause is called when production table_alias_clause is exited.
func (s *BaseGaussdbParserListener) ExitTable_alias_clause(ctx *Table_alias_clauseContext) {}

// EnterFunc_alias_clause is called when production func_alias_clause is entered.
func (s *BaseGaussdbParserListener) EnterFunc_alias_clause(ctx *Func_alias_clauseContext) {}

// ExitFunc_alias_clause is called when production func_alias_clause is exited.
func (s *BaseGaussdbParserListener) ExitFunc_alias_clause(ctx *Func_alias_clauseContext) {}

// EnterJoin_type is called when production join_type is entered.
func (s *BaseGaussdbParserListener) EnterJoin_type(ctx *Join_typeContext) {}

// ExitJoin_type is called when production join_type is exited.
func (s *BaseGaussdbParserListener) ExitJoin_type(ctx *Join_typeContext) {}

// EnterJoin_qual is called when production join_qual is entered.
func (s *BaseGaussdbParserListener) EnterJoin_qual(ctx *Join_qualContext) {}

// ExitJoin_qual is called when production join_qual is exited.
func (s *BaseGaussdbParserListener) ExitJoin_qual(ctx *Join_qualContext) {}

// EnterRelation_expr is called when production relation_expr is entered.
func (s *BaseGaussdbParserListener) EnterRelation_expr(ctx *Relation_exprContext) {}

// ExitRelation_expr is called when production relation_expr is exited.
func (s *BaseGaussdbParserListener) ExitRelation_expr(ctx *Relation_exprContext) {}

// EnterRelation_expr_list is called when production relation_expr_list is entered.
func (s *BaseGaussdbParserListener) EnterRelation_expr_list(ctx *Relation_expr_listContext) {}

// ExitRelation_expr_list is called when production relation_expr_list is exited.
func (s *BaseGaussdbParserListener) ExitRelation_expr_list(ctx *Relation_expr_listContext) {}

// EnterRelation_expr_opt_alias is called when production relation_expr_opt_alias is entered.
func (s *BaseGaussdbParserListener) EnterRelation_expr_opt_alias(ctx *Relation_expr_opt_aliasContext) {
}

// ExitRelation_expr_opt_alias is called when production relation_expr_opt_alias is exited.
func (s *BaseGaussdbParserListener) ExitRelation_expr_opt_alias(ctx *Relation_expr_opt_aliasContext) {
}

// EnterTablesample_clause is called when production tablesample_clause is entered.
func (s *BaseGaussdbParserListener) EnterTablesample_clause(ctx *Tablesample_clauseContext) {}

// ExitTablesample_clause is called when production tablesample_clause is exited.
func (s *BaseGaussdbParserListener) ExitTablesample_clause(ctx *Tablesample_clauseContext) {}

// EnterOpt_repeatable_clause is called when production opt_repeatable_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_repeatable_clause(ctx *Opt_repeatable_clauseContext) {}

// ExitOpt_repeatable_clause is called when production opt_repeatable_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_repeatable_clause(ctx *Opt_repeatable_clauseContext) {}

// EnterFunc_table is called when production func_table is entered.
func (s *BaseGaussdbParserListener) EnterFunc_table(ctx *Func_tableContext) {}

// ExitFunc_table is called when production func_table is exited.
func (s *BaseGaussdbParserListener) ExitFunc_table(ctx *Func_tableContext) {}

// EnterRowsfrom_item is called when production rowsfrom_item is entered.
func (s *BaseGaussdbParserListener) EnterRowsfrom_item(ctx *Rowsfrom_itemContext) {}

// ExitRowsfrom_item is called when production rowsfrom_item is exited.
func (s *BaseGaussdbParserListener) ExitRowsfrom_item(ctx *Rowsfrom_itemContext) {}

// EnterRowsfrom_list is called when production rowsfrom_list is entered.
func (s *BaseGaussdbParserListener) EnterRowsfrom_list(ctx *Rowsfrom_listContext) {}

// ExitRowsfrom_list is called when production rowsfrom_list is exited.
func (s *BaseGaussdbParserListener) ExitRowsfrom_list(ctx *Rowsfrom_listContext) {}

// EnterOpt_col_def_list is called when production opt_col_def_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_col_def_list(ctx *Opt_col_def_listContext) {}

// ExitOpt_col_def_list is called when production opt_col_def_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_col_def_list(ctx *Opt_col_def_listContext) {}

// EnterOpt_ordinality is called when production opt_ordinality is entered.
func (s *BaseGaussdbParserListener) EnterOpt_ordinality(ctx *Opt_ordinalityContext) {}

// ExitOpt_ordinality is called when production opt_ordinality is exited.
func (s *BaseGaussdbParserListener) ExitOpt_ordinality(ctx *Opt_ordinalityContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseGaussdbParserListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseGaussdbParserListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterWhere_or_current_clause is called when production where_or_current_clause is entered.
func (s *BaseGaussdbParserListener) EnterWhere_or_current_clause(ctx *Where_or_current_clauseContext) {
}

// ExitWhere_or_current_clause is called when production where_or_current_clause is exited.
func (s *BaseGaussdbParserListener) ExitWhere_or_current_clause(ctx *Where_or_current_clauseContext) {
}

// EnterOpttablefuncelementlist is called when production opttablefuncelementlist is entered.
func (s *BaseGaussdbParserListener) EnterOpttablefuncelementlist(ctx *OpttablefuncelementlistContext) {
}

// ExitOpttablefuncelementlist is called when production opttablefuncelementlist is exited.
func (s *BaseGaussdbParserListener) ExitOpttablefuncelementlist(ctx *OpttablefuncelementlistContext) {
}

// EnterTablefuncelementlist is called when production tablefuncelementlist is entered.
func (s *BaseGaussdbParserListener) EnterTablefuncelementlist(ctx *TablefuncelementlistContext) {}

// ExitTablefuncelementlist is called when production tablefuncelementlist is exited.
func (s *BaseGaussdbParserListener) ExitTablefuncelementlist(ctx *TablefuncelementlistContext) {}

// EnterTablefuncelement is called when production tablefuncelement is entered.
func (s *BaseGaussdbParserListener) EnterTablefuncelement(ctx *TablefuncelementContext) {}

// ExitTablefuncelement is called when production tablefuncelement is exited.
func (s *BaseGaussdbParserListener) ExitTablefuncelement(ctx *TablefuncelementContext) {}

// EnterXmltable is called when production xmltable is entered.
func (s *BaseGaussdbParserListener) EnterXmltable(ctx *XmltableContext) {}

// ExitXmltable is called when production xmltable is exited.
func (s *BaseGaussdbParserListener) ExitXmltable(ctx *XmltableContext) {}

// EnterXmltable_column_list is called when production xmltable_column_list is entered.
func (s *BaseGaussdbParserListener) EnterXmltable_column_list(ctx *Xmltable_column_listContext) {}

// ExitXmltable_column_list is called when production xmltable_column_list is exited.
func (s *BaseGaussdbParserListener) ExitXmltable_column_list(ctx *Xmltable_column_listContext) {}

// EnterXmltable_column_el is called when production xmltable_column_el is entered.
func (s *BaseGaussdbParserListener) EnterXmltable_column_el(ctx *Xmltable_column_elContext) {}

// ExitXmltable_column_el is called when production xmltable_column_el is exited.
func (s *BaseGaussdbParserListener) ExitXmltable_column_el(ctx *Xmltable_column_elContext) {}

// EnterXmltable_column_option_list is called when production xmltable_column_option_list is entered.
func (s *BaseGaussdbParserListener) EnterXmltable_column_option_list(ctx *Xmltable_column_option_listContext) {
}

// ExitXmltable_column_option_list is called when production xmltable_column_option_list is exited.
func (s *BaseGaussdbParserListener) ExitXmltable_column_option_list(ctx *Xmltable_column_option_listContext) {
}

// EnterXmltable_column_option_el is called when production xmltable_column_option_el is entered.
func (s *BaseGaussdbParserListener) EnterXmltable_column_option_el(ctx *Xmltable_column_option_elContext) {
}

// ExitXmltable_column_option_el is called when production xmltable_column_option_el is exited.
func (s *BaseGaussdbParserListener) ExitXmltable_column_option_el(ctx *Xmltable_column_option_elContext) {
}

// EnterXml_namespace_list is called when production xml_namespace_list is entered.
func (s *BaseGaussdbParserListener) EnterXml_namespace_list(ctx *Xml_namespace_listContext) {}

// ExitXml_namespace_list is called when production xml_namespace_list is exited.
func (s *BaseGaussdbParserListener) ExitXml_namespace_list(ctx *Xml_namespace_listContext) {}

// EnterXml_namespace_el is called when production xml_namespace_el is entered.
func (s *BaseGaussdbParserListener) EnterXml_namespace_el(ctx *Xml_namespace_elContext) {}

// ExitXml_namespace_el is called when production xml_namespace_el is exited.
func (s *BaseGaussdbParserListener) ExitXml_namespace_el(ctx *Xml_namespace_elContext) {}

// EnterTypename is called when production typename is entered.
func (s *BaseGaussdbParserListener) EnterTypename(ctx *TypenameContext) {}

// ExitTypename is called when production typename is exited.
func (s *BaseGaussdbParserListener) ExitTypename(ctx *TypenameContext) {}

// EnterOpt_array_bounds is called when production opt_array_bounds is entered.
func (s *BaseGaussdbParserListener) EnterOpt_array_bounds(ctx *Opt_array_boundsContext) {}

// ExitOpt_array_bounds is called when production opt_array_bounds is exited.
func (s *BaseGaussdbParserListener) ExitOpt_array_bounds(ctx *Opt_array_boundsContext) {}

// EnterSimpletypename is called when production simpletypename is entered.
func (s *BaseGaussdbParserListener) EnterSimpletypename(ctx *SimpletypenameContext) {}

// ExitSimpletypename is called when production simpletypename is exited.
func (s *BaseGaussdbParserListener) ExitSimpletypename(ctx *SimpletypenameContext) {}

// EnterConsttypename is called when production consttypename is entered.
func (s *BaseGaussdbParserListener) EnterConsttypename(ctx *ConsttypenameContext) {}

// ExitConsttypename is called when production consttypename is exited.
func (s *BaseGaussdbParserListener) ExitConsttypename(ctx *ConsttypenameContext) {}

// EnterGenerictype is called when production generictype is entered.
func (s *BaseGaussdbParserListener) EnterGenerictype(ctx *GenerictypeContext) {}

// ExitGenerictype is called when production generictype is exited.
func (s *BaseGaussdbParserListener) ExitGenerictype(ctx *GenerictypeContext) {}

// EnterOpt_type_modifiers is called when production opt_type_modifiers is entered.
func (s *BaseGaussdbParserListener) EnterOpt_type_modifiers(ctx *Opt_type_modifiersContext) {}

// ExitOpt_type_modifiers is called when production opt_type_modifiers is exited.
func (s *BaseGaussdbParserListener) ExitOpt_type_modifiers(ctx *Opt_type_modifiersContext) {}

// EnterNumeric is called when production numeric is entered.
func (s *BaseGaussdbParserListener) EnterNumeric(ctx *NumericContext) {}

// ExitNumeric is called when production numeric is exited.
func (s *BaseGaussdbParserListener) ExitNumeric(ctx *NumericContext) {}

// EnterOpt_float is called when production opt_float is entered.
func (s *BaseGaussdbParserListener) EnterOpt_float(ctx *Opt_floatContext) {}

// ExitOpt_float is called when production opt_float is exited.
func (s *BaseGaussdbParserListener) ExitOpt_float(ctx *Opt_floatContext) {}

// EnterBit is called when production bit is entered.
func (s *BaseGaussdbParserListener) EnterBit(ctx *BitContext) {}

// ExitBit is called when production bit is exited.
func (s *BaseGaussdbParserListener) ExitBit(ctx *BitContext) {}

// EnterConstbit is called when production constbit is entered.
func (s *BaseGaussdbParserListener) EnterConstbit(ctx *ConstbitContext) {}

// ExitConstbit is called when production constbit is exited.
func (s *BaseGaussdbParserListener) ExitConstbit(ctx *ConstbitContext) {}

// EnterBitwithlength is called when production bitwithlength is entered.
func (s *BaseGaussdbParserListener) EnterBitwithlength(ctx *BitwithlengthContext) {}

// ExitBitwithlength is called when production bitwithlength is exited.
func (s *BaseGaussdbParserListener) ExitBitwithlength(ctx *BitwithlengthContext) {}

// EnterBitwithoutlength is called when production bitwithoutlength is entered.
func (s *BaseGaussdbParserListener) EnterBitwithoutlength(ctx *BitwithoutlengthContext) {}

// ExitBitwithoutlength is called when production bitwithoutlength is exited.
func (s *BaseGaussdbParserListener) ExitBitwithoutlength(ctx *BitwithoutlengthContext) {}

// EnterCharacter is called when production character is entered.
func (s *BaseGaussdbParserListener) EnterCharacter(ctx *CharacterContext) {}

// ExitCharacter is called when production character is exited.
func (s *BaseGaussdbParserListener) ExitCharacter(ctx *CharacterContext) {}

// EnterConstcharacter is called when production constcharacter is entered.
func (s *BaseGaussdbParserListener) EnterConstcharacter(ctx *ConstcharacterContext) {}

// ExitConstcharacter is called when production constcharacter is exited.
func (s *BaseGaussdbParserListener) ExitConstcharacter(ctx *ConstcharacterContext) {}

// EnterCharacter_c is called when production character_c is entered.
func (s *BaseGaussdbParserListener) EnterCharacter_c(ctx *Character_cContext) {}

// ExitCharacter_c is called when production character_c is exited.
func (s *BaseGaussdbParserListener) ExitCharacter_c(ctx *Character_cContext) {}

// EnterOpt_varying is called when production opt_varying is entered.
func (s *BaseGaussdbParserListener) EnterOpt_varying(ctx *Opt_varyingContext) {}

// ExitOpt_varying is called when production opt_varying is exited.
func (s *BaseGaussdbParserListener) ExitOpt_varying(ctx *Opt_varyingContext) {}

// EnterConstdatetime is called when production constdatetime is entered.
func (s *BaseGaussdbParserListener) EnterConstdatetime(ctx *ConstdatetimeContext) {}

// ExitConstdatetime is called when production constdatetime is exited.
func (s *BaseGaussdbParserListener) ExitConstdatetime(ctx *ConstdatetimeContext) {}

// EnterConstinterval is called when production constinterval is entered.
func (s *BaseGaussdbParserListener) EnterConstinterval(ctx *ConstintervalContext) {}

// ExitConstinterval is called when production constinterval is exited.
func (s *BaseGaussdbParserListener) ExitConstinterval(ctx *ConstintervalContext) {}

// EnterOpt_timezone is called when production opt_timezone is entered.
func (s *BaseGaussdbParserListener) EnterOpt_timezone(ctx *Opt_timezoneContext) {}

// ExitOpt_timezone is called when production opt_timezone is exited.
func (s *BaseGaussdbParserListener) ExitOpt_timezone(ctx *Opt_timezoneContext) {}

// EnterOpt_interval is called when production opt_interval is entered.
func (s *BaseGaussdbParserListener) EnterOpt_interval(ctx *Opt_intervalContext) {}

// ExitOpt_interval is called when production opt_interval is exited.
func (s *BaseGaussdbParserListener) ExitOpt_interval(ctx *Opt_intervalContext) {}

// EnterInterval_second is called when production interval_second is entered.
func (s *BaseGaussdbParserListener) EnterInterval_second(ctx *Interval_secondContext) {}

// ExitInterval_second is called when production interval_second is exited.
func (s *BaseGaussdbParserListener) ExitInterval_second(ctx *Interval_secondContext) {}

// EnterOpt_escape is called when production opt_escape is entered.
func (s *BaseGaussdbParserListener) EnterOpt_escape(ctx *Opt_escapeContext) {}

// ExitOpt_escape is called when production opt_escape is exited.
func (s *BaseGaussdbParserListener) ExitOpt_escape(ctx *Opt_escapeContext) {}

// EnterA_expr is called when production a_expr is entered.
func (s *BaseGaussdbParserListener) EnterA_expr(ctx *A_exprContext) {}

// ExitA_expr is called when production a_expr is exited.
func (s *BaseGaussdbParserListener) ExitA_expr(ctx *A_exprContext) {}

// EnterA_expr_qual is called when production a_expr_qual is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_qual(ctx *A_expr_qualContext) {}

// ExitA_expr_qual is called when production a_expr_qual is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_qual(ctx *A_expr_qualContext) {}

// EnterA_expr_lessless is called when production a_expr_lessless is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_lessless(ctx *A_expr_lesslessContext) {}

// ExitA_expr_lessless is called when production a_expr_lessless is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_lessless(ctx *A_expr_lesslessContext) {}

// EnterA_expr_or is called when production a_expr_or is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_or(ctx *A_expr_orContext) {}

// ExitA_expr_or is called when production a_expr_or is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_or(ctx *A_expr_orContext) {}

// EnterA_expr_and is called when production a_expr_and is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_and(ctx *A_expr_andContext) {}

// ExitA_expr_and is called when production a_expr_and is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_and(ctx *A_expr_andContext) {}

// EnterA_expr_between is called when production a_expr_between is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_between(ctx *A_expr_betweenContext) {}

// ExitA_expr_between is called when production a_expr_between is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_between(ctx *A_expr_betweenContext) {}

// EnterA_expr_in is called when production a_expr_in is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_in(ctx *A_expr_inContext) {}

// ExitA_expr_in is called when production a_expr_in is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_in(ctx *A_expr_inContext) {}

// EnterA_expr_unary_not is called when production a_expr_unary_not is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_unary_not(ctx *A_expr_unary_notContext) {}

// ExitA_expr_unary_not is called when production a_expr_unary_not is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_unary_not(ctx *A_expr_unary_notContext) {}

// EnterA_expr_isnull is called when production a_expr_isnull is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_isnull(ctx *A_expr_isnullContext) {}

// ExitA_expr_isnull is called when production a_expr_isnull is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_isnull(ctx *A_expr_isnullContext) {}

// EnterA_expr_is_not is called when production a_expr_is_not is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_is_not(ctx *A_expr_is_notContext) {}

// ExitA_expr_is_not is called when production a_expr_is_not is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_is_not(ctx *A_expr_is_notContext) {}

// EnterA_expr_compare is called when production a_expr_compare is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_compare(ctx *A_expr_compareContext) {}

// ExitA_expr_compare is called when production a_expr_compare is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_compare(ctx *A_expr_compareContext) {}

// EnterA_expr_like is called when production a_expr_like is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_like(ctx *A_expr_likeContext) {}

// ExitA_expr_like is called when production a_expr_like is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_like(ctx *A_expr_likeContext) {}

// EnterA_expr_qual_op is called when production a_expr_qual_op is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_qual_op(ctx *A_expr_qual_opContext) {}

// ExitA_expr_qual_op is called when production a_expr_qual_op is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_qual_op(ctx *A_expr_qual_opContext) {}

// EnterA_expr_unary_qualop is called when production a_expr_unary_qualop is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_unary_qualop(ctx *A_expr_unary_qualopContext) {}

// ExitA_expr_unary_qualop is called when production a_expr_unary_qualop is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_unary_qualop(ctx *A_expr_unary_qualopContext) {}

// EnterA_expr_add is called when production a_expr_add is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_add(ctx *A_expr_addContext) {}

// ExitA_expr_add is called when production a_expr_add is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_add(ctx *A_expr_addContext) {}

// EnterA_expr_mul is called when production a_expr_mul is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_mul(ctx *A_expr_mulContext) {}

// ExitA_expr_mul is called when production a_expr_mul is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_mul(ctx *A_expr_mulContext) {}

// EnterA_expr_caret is called when production a_expr_caret is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_caret(ctx *A_expr_caretContext) {}

// ExitA_expr_caret is called when production a_expr_caret is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_caret(ctx *A_expr_caretContext) {}

// EnterA_expr_unary_sign is called when production a_expr_unary_sign is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_unary_sign(ctx *A_expr_unary_signContext) {}

// ExitA_expr_unary_sign is called when production a_expr_unary_sign is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_unary_sign(ctx *A_expr_unary_signContext) {}

// EnterA_expr_at_time_zone is called when production a_expr_at_time_zone is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_at_time_zone(ctx *A_expr_at_time_zoneContext) {}

// ExitA_expr_at_time_zone is called when production a_expr_at_time_zone is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_at_time_zone(ctx *A_expr_at_time_zoneContext) {}

// EnterA_expr_collate is called when production a_expr_collate is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_collate(ctx *A_expr_collateContext) {}

// ExitA_expr_collate is called when production a_expr_collate is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_collate(ctx *A_expr_collateContext) {}

// EnterA_expr_typecast is called when production a_expr_typecast is entered.
func (s *BaseGaussdbParserListener) EnterA_expr_typecast(ctx *A_expr_typecastContext) {}

// ExitA_expr_typecast is called when production a_expr_typecast is exited.
func (s *BaseGaussdbParserListener) ExitA_expr_typecast(ctx *A_expr_typecastContext) {}

// EnterB_expr is called when production b_expr is entered.
func (s *BaseGaussdbParserListener) EnterB_expr(ctx *B_exprContext) {}

// ExitB_expr is called when production b_expr is exited.
func (s *BaseGaussdbParserListener) ExitB_expr(ctx *B_exprContext) {}

// EnterC_expr_exists is called when production c_expr_exists is entered.
func (s *BaseGaussdbParserListener) EnterC_expr_exists(ctx *C_expr_existsContext) {}

// ExitC_expr_exists is called when production c_expr_exists is exited.
func (s *BaseGaussdbParserListener) ExitC_expr_exists(ctx *C_expr_existsContext) {}

// EnterC_expr_expr is called when production c_expr_expr is entered.
func (s *BaseGaussdbParserListener) EnterC_expr_expr(ctx *C_expr_exprContext) {}

// ExitC_expr_expr is called when production c_expr_expr is exited.
func (s *BaseGaussdbParserListener) ExitC_expr_expr(ctx *C_expr_exprContext) {}

// EnterC_expr_case is called when production c_expr_case is entered.
func (s *BaseGaussdbParserListener) EnterC_expr_case(ctx *C_expr_caseContext) {}

// ExitC_expr_case is called when production c_expr_case is exited.
func (s *BaseGaussdbParserListener) ExitC_expr_case(ctx *C_expr_caseContext) {}

// EnterPlsqlvariablename is called when production plsqlvariablename is entered.
func (s *BaseGaussdbParserListener) EnterPlsqlvariablename(ctx *PlsqlvariablenameContext) {}

// ExitPlsqlvariablename is called when production plsqlvariablename is exited.
func (s *BaseGaussdbParserListener) ExitPlsqlvariablename(ctx *PlsqlvariablenameContext) {}

// EnterFunc_application is called when production func_application is entered.
func (s *BaseGaussdbParserListener) EnterFunc_application(ctx *Func_applicationContext) {}

// ExitFunc_application is called when production func_application is exited.
func (s *BaseGaussdbParserListener) ExitFunc_application(ctx *Func_applicationContext) {}

// EnterFunc_expr is called when production func_expr is entered.
func (s *BaseGaussdbParserListener) EnterFunc_expr(ctx *Func_exprContext) {}

// ExitFunc_expr is called when production func_expr is exited.
func (s *BaseGaussdbParserListener) ExitFunc_expr(ctx *Func_exprContext) {}

// EnterFunc_expr_windowless is called when production func_expr_windowless is entered.
func (s *BaseGaussdbParserListener) EnterFunc_expr_windowless(ctx *Func_expr_windowlessContext) {}

// ExitFunc_expr_windowless is called when production func_expr_windowless is exited.
func (s *BaseGaussdbParserListener) ExitFunc_expr_windowless(ctx *Func_expr_windowlessContext) {}

// EnterFunc_expr_common_subexpr is called when production func_expr_common_subexpr is entered.
func (s *BaseGaussdbParserListener) EnterFunc_expr_common_subexpr(ctx *Func_expr_common_subexprContext) {
}

// ExitFunc_expr_common_subexpr is called when production func_expr_common_subexpr is exited.
func (s *BaseGaussdbParserListener) ExitFunc_expr_common_subexpr(ctx *Func_expr_common_subexprContext) {
}

// EnterXml_root_version is called when production xml_root_version is entered.
func (s *BaseGaussdbParserListener) EnterXml_root_version(ctx *Xml_root_versionContext) {}

// ExitXml_root_version is called when production xml_root_version is exited.
func (s *BaseGaussdbParserListener) ExitXml_root_version(ctx *Xml_root_versionContext) {}

// EnterOpt_xml_root_standalone is called when production opt_xml_root_standalone is entered.
func (s *BaseGaussdbParserListener) EnterOpt_xml_root_standalone(ctx *Opt_xml_root_standaloneContext) {
}

// ExitOpt_xml_root_standalone is called when production opt_xml_root_standalone is exited.
func (s *BaseGaussdbParserListener) ExitOpt_xml_root_standalone(ctx *Opt_xml_root_standaloneContext) {
}

// EnterXml_attributes is called when production xml_attributes is entered.
func (s *BaseGaussdbParserListener) EnterXml_attributes(ctx *Xml_attributesContext) {}

// ExitXml_attributes is called when production xml_attributes is exited.
func (s *BaseGaussdbParserListener) ExitXml_attributes(ctx *Xml_attributesContext) {}

// EnterXml_attribute_list is called when production xml_attribute_list is entered.
func (s *BaseGaussdbParserListener) EnterXml_attribute_list(ctx *Xml_attribute_listContext) {}

// ExitXml_attribute_list is called when production xml_attribute_list is exited.
func (s *BaseGaussdbParserListener) ExitXml_attribute_list(ctx *Xml_attribute_listContext) {}

// EnterXml_attribute_el is called when production xml_attribute_el is entered.
func (s *BaseGaussdbParserListener) EnterXml_attribute_el(ctx *Xml_attribute_elContext) {}

// ExitXml_attribute_el is called when production xml_attribute_el is exited.
func (s *BaseGaussdbParserListener) ExitXml_attribute_el(ctx *Xml_attribute_elContext) {}

// EnterDocument_or_content is called when production document_or_content is entered.
func (s *BaseGaussdbParserListener) EnterDocument_or_content(ctx *Document_or_contentContext) {}

// ExitDocument_or_content is called when production document_or_content is exited.
func (s *BaseGaussdbParserListener) ExitDocument_or_content(ctx *Document_or_contentContext) {}

// EnterXml_whitespace_option is called when production xml_whitespace_option is entered.
func (s *BaseGaussdbParserListener) EnterXml_whitespace_option(ctx *Xml_whitespace_optionContext) {}

// ExitXml_whitespace_option is called when production xml_whitespace_option is exited.
func (s *BaseGaussdbParserListener) ExitXml_whitespace_option(ctx *Xml_whitespace_optionContext) {}

// EnterXmlexists_argument is called when production xmlexists_argument is entered.
func (s *BaseGaussdbParserListener) EnterXmlexists_argument(ctx *Xmlexists_argumentContext) {}

// ExitXmlexists_argument is called when production xmlexists_argument is exited.
func (s *BaseGaussdbParserListener) ExitXmlexists_argument(ctx *Xmlexists_argumentContext) {}

// EnterXml_passing_mech is called when production xml_passing_mech is entered.
func (s *BaseGaussdbParserListener) EnterXml_passing_mech(ctx *Xml_passing_mechContext) {}

// ExitXml_passing_mech is called when production xml_passing_mech is exited.
func (s *BaseGaussdbParserListener) ExitXml_passing_mech(ctx *Xml_passing_mechContext) {}

// EnterWithin_group_clause is called when production within_group_clause is entered.
func (s *BaseGaussdbParserListener) EnterWithin_group_clause(ctx *Within_group_clauseContext) {}

// ExitWithin_group_clause is called when production within_group_clause is exited.
func (s *BaseGaussdbParserListener) ExitWithin_group_clause(ctx *Within_group_clauseContext) {}

// EnterFilter_clause is called when production filter_clause is entered.
func (s *BaseGaussdbParserListener) EnterFilter_clause(ctx *Filter_clauseContext) {}

// ExitFilter_clause is called when production filter_clause is exited.
func (s *BaseGaussdbParserListener) ExitFilter_clause(ctx *Filter_clauseContext) {}

// EnterWindow_clause is called when production window_clause is entered.
func (s *BaseGaussdbParserListener) EnterWindow_clause(ctx *Window_clauseContext) {}

// ExitWindow_clause is called when production window_clause is exited.
func (s *BaseGaussdbParserListener) ExitWindow_clause(ctx *Window_clauseContext) {}

// EnterWindow_definition_list is called when production window_definition_list is entered.
func (s *BaseGaussdbParserListener) EnterWindow_definition_list(ctx *Window_definition_listContext) {}

// ExitWindow_definition_list is called when production window_definition_list is exited.
func (s *BaseGaussdbParserListener) ExitWindow_definition_list(ctx *Window_definition_listContext) {}

// EnterWindow_definition is called when production window_definition is entered.
func (s *BaseGaussdbParserListener) EnterWindow_definition(ctx *Window_definitionContext) {}

// ExitWindow_definition is called when production window_definition is exited.
func (s *BaseGaussdbParserListener) ExitWindow_definition(ctx *Window_definitionContext) {}

// EnterOver_clause is called when production over_clause is entered.
func (s *BaseGaussdbParserListener) EnterOver_clause(ctx *Over_clauseContext) {}

// ExitOver_clause is called when production over_clause is exited.
func (s *BaseGaussdbParserListener) ExitOver_clause(ctx *Over_clauseContext) {}

// EnterWindow_specification is called when production window_specification is entered.
func (s *BaseGaussdbParserListener) EnterWindow_specification(ctx *Window_specificationContext) {}

// ExitWindow_specification is called when production window_specification is exited.
func (s *BaseGaussdbParserListener) ExitWindow_specification(ctx *Window_specificationContext) {}

// EnterOpt_existing_window_name is called when production opt_existing_window_name is entered.
func (s *BaseGaussdbParserListener) EnterOpt_existing_window_name(ctx *Opt_existing_window_nameContext) {
}

// ExitOpt_existing_window_name is called when production opt_existing_window_name is exited.
func (s *BaseGaussdbParserListener) ExitOpt_existing_window_name(ctx *Opt_existing_window_nameContext) {
}

// EnterOpt_partition_clause is called when production opt_partition_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_partition_clause(ctx *Opt_partition_clauseContext) {}

// ExitOpt_partition_clause is called when production opt_partition_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_partition_clause(ctx *Opt_partition_clauseContext) {}

// EnterOpt_frame_clause is called when production opt_frame_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_frame_clause(ctx *Opt_frame_clauseContext) {}

// ExitOpt_frame_clause is called when production opt_frame_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_frame_clause(ctx *Opt_frame_clauseContext) {}

// EnterFrame_extent is called when production frame_extent is entered.
func (s *BaseGaussdbParserListener) EnterFrame_extent(ctx *Frame_extentContext) {}

// ExitFrame_extent is called when production frame_extent is exited.
func (s *BaseGaussdbParserListener) ExitFrame_extent(ctx *Frame_extentContext) {}

// EnterFrame_bound is called when production frame_bound is entered.
func (s *BaseGaussdbParserListener) EnterFrame_bound(ctx *Frame_boundContext) {}

// ExitFrame_bound is called when production frame_bound is exited.
func (s *BaseGaussdbParserListener) ExitFrame_bound(ctx *Frame_boundContext) {}

// EnterOpt_window_exclusion_clause is called when production opt_window_exclusion_clause is entered.
func (s *BaseGaussdbParserListener) EnterOpt_window_exclusion_clause(ctx *Opt_window_exclusion_clauseContext) {
}

// ExitOpt_window_exclusion_clause is called when production opt_window_exclusion_clause is exited.
func (s *BaseGaussdbParserListener) ExitOpt_window_exclusion_clause(ctx *Opt_window_exclusion_clauseContext) {
}

// EnterRow is called when production row is entered.
func (s *BaseGaussdbParserListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BaseGaussdbParserListener) ExitRow(ctx *RowContext) {}

// EnterExplicit_row is called when production explicit_row is entered.
func (s *BaseGaussdbParserListener) EnterExplicit_row(ctx *Explicit_rowContext) {}

// ExitExplicit_row is called when production explicit_row is exited.
func (s *BaseGaussdbParserListener) ExitExplicit_row(ctx *Explicit_rowContext) {}

// EnterImplicit_row is called when production implicit_row is entered.
func (s *BaseGaussdbParserListener) EnterImplicit_row(ctx *Implicit_rowContext) {}

// ExitImplicit_row is called when production implicit_row is exited.
func (s *BaseGaussdbParserListener) ExitImplicit_row(ctx *Implicit_rowContext) {}

// EnterSub_type is called when production sub_type is entered.
func (s *BaseGaussdbParserListener) EnterSub_type(ctx *Sub_typeContext) {}

// ExitSub_type is called when production sub_type is exited.
func (s *BaseGaussdbParserListener) ExitSub_type(ctx *Sub_typeContext) {}

// EnterAll_op is called when production all_op is entered.
func (s *BaseGaussdbParserListener) EnterAll_op(ctx *All_opContext) {}

// ExitAll_op is called when production all_op is exited.
func (s *BaseGaussdbParserListener) ExitAll_op(ctx *All_opContext) {}

// EnterMathop is called when production mathop is entered.
func (s *BaseGaussdbParserListener) EnterMathop(ctx *MathopContext) {}

// ExitMathop is called when production mathop is exited.
func (s *BaseGaussdbParserListener) ExitMathop(ctx *MathopContext) {}

// EnterQual_op is called when production qual_op is entered.
func (s *BaseGaussdbParserListener) EnterQual_op(ctx *Qual_opContext) {}

// ExitQual_op is called when production qual_op is exited.
func (s *BaseGaussdbParserListener) ExitQual_op(ctx *Qual_opContext) {}

// EnterQual_all_op is called when production qual_all_op is entered.
func (s *BaseGaussdbParserListener) EnterQual_all_op(ctx *Qual_all_opContext) {}

// ExitQual_all_op is called when production qual_all_op is exited.
func (s *BaseGaussdbParserListener) ExitQual_all_op(ctx *Qual_all_opContext) {}

// EnterSubquery_Op is called when production subquery_Op is entered.
func (s *BaseGaussdbParserListener) EnterSubquery_Op(ctx *Subquery_OpContext) {}

// ExitSubquery_Op is called when production subquery_Op is exited.
func (s *BaseGaussdbParserListener) ExitSubquery_Op(ctx *Subquery_OpContext) {}

// EnterExpr_list is called when production expr_list is entered.
func (s *BaseGaussdbParserListener) EnterExpr_list(ctx *Expr_listContext) {}

// ExitExpr_list is called when production expr_list is exited.
func (s *BaseGaussdbParserListener) ExitExpr_list(ctx *Expr_listContext) {}

// EnterFunc_arg_list is called when production func_arg_list is entered.
func (s *BaseGaussdbParserListener) EnterFunc_arg_list(ctx *Func_arg_listContext) {}

// ExitFunc_arg_list is called when production func_arg_list is exited.
func (s *BaseGaussdbParserListener) ExitFunc_arg_list(ctx *Func_arg_listContext) {}

// EnterFunc_arg_expr is called when production func_arg_expr is entered.
func (s *BaseGaussdbParserListener) EnterFunc_arg_expr(ctx *Func_arg_exprContext) {}

// ExitFunc_arg_expr is called when production func_arg_expr is exited.
func (s *BaseGaussdbParserListener) ExitFunc_arg_expr(ctx *Func_arg_exprContext) {}

// EnterType_list is called when production type_list is entered.
func (s *BaseGaussdbParserListener) EnterType_list(ctx *Type_listContext) {}

// ExitType_list is called when production type_list is exited.
func (s *BaseGaussdbParserListener) ExitType_list(ctx *Type_listContext) {}

// EnterArray_expr is called when production array_expr is entered.
func (s *BaseGaussdbParserListener) EnterArray_expr(ctx *Array_exprContext) {}

// ExitArray_expr is called when production array_expr is exited.
func (s *BaseGaussdbParserListener) ExitArray_expr(ctx *Array_exprContext) {}

// EnterArray_expr_list is called when production array_expr_list is entered.
func (s *BaseGaussdbParserListener) EnterArray_expr_list(ctx *Array_expr_listContext) {}

// ExitArray_expr_list is called when production array_expr_list is exited.
func (s *BaseGaussdbParserListener) ExitArray_expr_list(ctx *Array_expr_listContext) {}

// EnterExtract_list is called when production extract_list is entered.
func (s *BaseGaussdbParserListener) EnterExtract_list(ctx *Extract_listContext) {}

// ExitExtract_list is called when production extract_list is exited.
func (s *BaseGaussdbParserListener) ExitExtract_list(ctx *Extract_listContext) {}

// EnterExtract_arg is called when production extract_arg is entered.
func (s *BaseGaussdbParserListener) EnterExtract_arg(ctx *Extract_argContext) {}

// ExitExtract_arg is called when production extract_arg is exited.
func (s *BaseGaussdbParserListener) ExitExtract_arg(ctx *Extract_argContext) {}

// EnterUnicode_normal_form is called when production unicode_normal_form is entered.
func (s *BaseGaussdbParserListener) EnterUnicode_normal_form(ctx *Unicode_normal_formContext) {}

// ExitUnicode_normal_form is called when production unicode_normal_form is exited.
func (s *BaseGaussdbParserListener) ExitUnicode_normal_form(ctx *Unicode_normal_formContext) {}

// EnterOverlay_list is called when production overlay_list is entered.
func (s *BaseGaussdbParserListener) EnterOverlay_list(ctx *Overlay_listContext) {}

// ExitOverlay_list is called when production overlay_list is exited.
func (s *BaseGaussdbParserListener) ExitOverlay_list(ctx *Overlay_listContext) {}

// EnterPosition_list is called when production position_list is entered.
func (s *BaseGaussdbParserListener) EnterPosition_list(ctx *Position_listContext) {}

// ExitPosition_list is called when production position_list is exited.
func (s *BaseGaussdbParserListener) ExitPosition_list(ctx *Position_listContext) {}

// EnterSubstr_list is called when production substr_list is entered.
func (s *BaseGaussdbParserListener) EnterSubstr_list(ctx *Substr_listContext) {}

// ExitSubstr_list is called when production substr_list is exited.
func (s *BaseGaussdbParserListener) ExitSubstr_list(ctx *Substr_listContext) {}

// EnterTrim_list is called when production trim_list is entered.
func (s *BaseGaussdbParserListener) EnterTrim_list(ctx *Trim_listContext) {}

// ExitTrim_list is called when production trim_list is exited.
func (s *BaseGaussdbParserListener) ExitTrim_list(ctx *Trim_listContext) {}

// EnterIn_expr_select is called when production in_expr_select is entered.
func (s *BaseGaussdbParserListener) EnterIn_expr_select(ctx *In_expr_selectContext) {}

// ExitIn_expr_select is called when production in_expr_select is exited.
func (s *BaseGaussdbParserListener) ExitIn_expr_select(ctx *In_expr_selectContext) {}

// EnterIn_expr_list is called when production in_expr_list is entered.
func (s *BaseGaussdbParserListener) EnterIn_expr_list(ctx *In_expr_listContext) {}

// ExitIn_expr_list is called when production in_expr_list is exited.
func (s *BaseGaussdbParserListener) ExitIn_expr_list(ctx *In_expr_listContext) {}

// EnterCase_expr is called when production case_expr is entered.
func (s *BaseGaussdbParserListener) EnterCase_expr(ctx *Case_exprContext) {}

// ExitCase_expr is called when production case_expr is exited.
func (s *BaseGaussdbParserListener) ExitCase_expr(ctx *Case_exprContext) {}

// EnterWhen_clause_list is called when production when_clause_list is entered.
func (s *BaseGaussdbParserListener) EnterWhen_clause_list(ctx *When_clause_listContext) {}

// ExitWhen_clause_list is called when production when_clause_list is exited.
func (s *BaseGaussdbParserListener) ExitWhen_clause_list(ctx *When_clause_listContext) {}

// EnterWhen_clause is called when production when_clause is entered.
func (s *BaseGaussdbParserListener) EnterWhen_clause(ctx *When_clauseContext) {}

// ExitWhen_clause is called when production when_clause is exited.
func (s *BaseGaussdbParserListener) ExitWhen_clause(ctx *When_clauseContext) {}

// EnterCase_default is called when production case_default is entered.
func (s *BaseGaussdbParserListener) EnterCase_default(ctx *Case_defaultContext) {}

// ExitCase_default is called when production case_default is exited.
func (s *BaseGaussdbParserListener) ExitCase_default(ctx *Case_defaultContext) {}

// EnterCase_arg is called when production case_arg is entered.
func (s *BaseGaussdbParserListener) EnterCase_arg(ctx *Case_argContext) {}

// ExitCase_arg is called when production case_arg is exited.
func (s *BaseGaussdbParserListener) ExitCase_arg(ctx *Case_argContext) {}

// EnterColumnref is called when production columnref is entered.
func (s *BaseGaussdbParserListener) EnterColumnref(ctx *ColumnrefContext) {}

// ExitColumnref is called when production columnref is exited.
func (s *BaseGaussdbParserListener) ExitColumnref(ctx *ColumnrefContext) {}

// EnterIndirection_el is called when production indirection_el is entered.
func (s *BaseGaussdbParserListener) EnterIndirection_el(ctx *Indirection_elContext) {}

// ExitIndirection_el is called when production indirection_el is exited.
func (s *BaseGaussdbParserListener) ExitIndirection_el(ctx *Indirection_elContext) {}

// EnterOpt_slice_bound is called when production opt_slice_bound is entered.
func (s *BaseGaussdbParserListener) EnterOpt_slice_bound(ctx *Opt_slice_boundContext) {}

// ExitOpt_slice_bound is called when production opt_slice_bound is exited.
func (s *BaseGaussdbParserListener) ExitOpt_slice_bound(ctx *Opt_slice_boundContext) {}

// EnterIndirection is called when production indirection is entered.
func (s *BaseGaussdbParserListener) EnterIndirection(ctx *IndirectionContext) {}

// ExitIndirection is called when production indirection is exited.
func (s *BaseGaussdbParserListener) ExitIndirection(ctx *IndirectionContext) {}

// EnterOpt_indirection is called when production opt_indirection is entered.
func (s *BaseGaussdbParserListener) EnterOpt_indirection(ctx *Opt_indirectionContext) {}

// ExitOpt_indirection is called when production opt_indirection is exited.
func (s *BaseGaussdbParserListener) ExitOpt_indirection(ctx *Opt_indirectionContext) {}

// EnterOpt_target_list is called when production opt_target_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_target_list(ctx *Opt_target_listContext) {}

// ExitOpt_target_list is called when production opt_target_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_target_list(ctx *Opt_target_listContext) {}

// EnterTarget_list is called when production target_list is entered.
func (s *BaseGaussdbParserListener) EnterTarget_list(ctx *Target_listContext) {}

// ExitTarget_list is called when production target_list is exited.
func (s *BaseGaussdbParserListener) ExitTarget_list(ctx *Target_listContext) {}

// EnterTarget_label is called when production target_label is entered.
func (s *BaseGaussdbParserListener) EnterTarget_label(ctx *Target_labelContext) {}

// ExitTarget_label is called when production target_label is exited.
func (s *BaseGaussdbParserListener) ExitTarget_label(ctx *Target_labelContext) {}

// EnterTarget_star is called when production target_star is entered.
func (s *BaseGaussdbParserListener) EnterTarget_star(ctx *Target_starContext) {}

// ExitTarget_star is called when production target_star is exited.
func (s *BaseGaussdbParserListener) ExitTarget_star(ctx *Target_starContext) {}

// EnterTarget_alias is called when production target_alias is entered.
func (s *BaseGaussdbParserListener) EnterTarget_alias(ctx *Target_aliasContext) {}

// ExitTarget_alias is called when production target_alias is exited.
func (s *BaseGaussdbParserListener) ExitTarget_alias(ctx *Target_aliasContext) {}

// EnterQualified_name_list is called when production qualified_name_list is entered.
func (s *BaseGaussdbParserListener) EnterQualified_name_list(ctx *Qualified_name_listContext) {}

// ExitQualified_name_list is called when production qualified_name_list is exited.
func (s *BaseGaussdbParserListener) ExitQualified_name_list(ctx *Qualified_name_listContext) {}

// EnterQualified_name is called when production qualified_name is entered.
func (s *BaseGaussdbParserListener) EnterQualified_name(ctx *Qualified_nameContext) {}

// ExitQualified_name is called when production qualified_name is exited.
func (s *BaseGaussdbParserListener) ExitQualified_name(ctx *Qualified_nameContext) {}

// EnterLink_name is called when production link_name is entered.
func (s *BaseGaussdbParserListener) EnterLink_name(ctx *Link_nameContext) {}

// ExitLink_name is called when production link_name is exited.
func (s *BaseGaussdbParserListener) ExitLink_name(ctx *Link_nameContext) {}

// EnterAt_link_name is called when production at_link_name is entered.
func (s *BaseGaussdbParserListener) EnterAt_link_name(ctx *At_link_nameContext) {}

// ExitAt_link_name is called when production at_link_name is exited.
func (s *BaseGaussdbParserListener) ExitAt_link_name(ctx *At_link_nameContext) {}

// EnterName_list is called when production name_list is entered.
func (s *BaseGaussdbParserListener) EnterName_list(ctx *Name_listContext) {}

// ExitName_list is called when production name_list is exited.
func (s *BaseGaussdbParserListener) ExitName_list(ctx *Name_listContext) {}

// EnterName is called when production name is entered.
func (s *BaseGaussdbParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseGaussdbParserListener) ExitName(ctx *NameContext) {}

// EnterAttr_name is called when production attr_name is entered.
func (s *BaseGaussdbParserListener) EnterAttr_name(ctx *Attr_nameContext) {}

// ExitAttr_name is called when production attr_name is exited.
func (s *BaseGaussdbParserListener) ExitAttr_name(ctx *Attr_nameContext) {}

// EnterFile_name is called when production file_name is entered.
func (s *BaseGaussdbParserListener) EnterFile_name(ctx *File_nameContext) {}

// ExitFile_name is called when production file_name is exited.
func (s *BaseGaussdbParserListener) ExitFile_name(ctx *File_nameContext) {}

// EnterFunc_name is called when production func_name is entered.
func (s *BaseGaussdbParserListener) EnterFunc_name(ctx *Func_nameContext) {}

// ExitFunc_name is called when production func_name is exited.
func (s *BaseGaussdbParserListener) ExitFunc_name(ctx *Func_nameContext) {}

// EnterAexprconst is called when production aexprconst is entered.
func (s *BaseGaussdbParserListener) EnterAexprconst(ctx *AexprconstContext) {}

// ExitAexprconst is called when production aexprconst is exited.
func (s *BaseGaussdbParserListener) ExitAexprconst(ctx *AexprconstContext) {}

// EnterXconst is called when production xconst is entered.
func (s *BaseGaussdbParserListener) EnterXconst(ctx *XconstContext) {}

// ExitXconst is called when production xconst is exited.
func (s *BaseGaussdbParserListener) ExitXconst(ctx *XconstContext) {}

// EnterBconst is called when production bconst is entered.
func (s *BaseGaussdbParserListener) EnterBconst(ctx *BconstContext) {}

// ExitBconst is called when production bconst is exited.
func (s *BaseGaussdbParserListener) ExitBconst(ctx *BconstContext) {}

// EnterFconst is called when production fconst is entered.
func (s *BaseGaussdbParserListener) EnterFconst(ctx *FconstContext) {}

// ExitFconst is called when production fconst is exited.
func (s *BaseGaussdbParserListener) ExitFconst(ctx *FconstContext) {}

// EnterIconst is called when production iconst is entered.
func (s *BaseGaussdbParserListener) EnterIconst(ctx *IconstContext) {}

// ExitIconst is called when production iconst is exited.
func (s *BaseGaussdbParserListener) ExitIconst(ctx *IconstContext) {}

// EnterSconst is called when production sconst is entered.
func (s *BaseGaussdbParserListener) EnterSconst(ctx *SconstContext) {}

// ExitSconst is called when production sconst is exited.
func (s *BaseGaussdbParserListener) ExitSconst(ctx *SconstContext) {}

// EnterAnysconst is called when production anysconst is entered.
func (s *BaseGaussdbParserListener) EnterAnysconst(ctx *AnysconstContext) {}

// ExitAnysconst is called when production anysconst is exited.
func (s *BaseGaussdbParserListener) ExitAnysconst(ctx *AnysconstContext) {}

// EnterOpt_uescape is called when production opt_uescape is entered.
func (s *BaseGaussdbParserListener) EnterOpt_uescape(ctx *Opt_uescapeContext) {}

// ExitOpt_uescape is called when production opt_uescape is exited.
func (s *BaseGaussdbParserListener) ExitOpt_uescape(ctx *Opt_uescapeContext) {}

// EnterSignediconst is called when production signediconst is entered.
func (s *BaseGaussdbParserListener) EnterSignediconst(ctx *SignediconstContext) {}

// ExitSignediconst is called when production signediconst is exited.
func (s *BaseGaussdbParserListener) ExitSignediconst(ctx *SignediconstContext) {}

// EnterRoleid is called when production roleid is entered.
func (s *BaseGaussdbParserListener) EnterRoleid(ctx *RoleidContext) {}

// ExitRoleid is called when production roleid is exited.
func (s *BaseGaussdbParserListener) ExitRoleid(ctx *RoleidContext) {}

// EnterRolespec is called when production rolespec is entered.
func (s *BaseGaussdbParserListener) EnterRolespec(ctx *RolespecContext) {}

// ExitRolespec is called when production rolespec is exited.
func (s *BaseGaussdbParserListener) ExitRolespec(ctx *RolespecContext) {}

// EnterRole_list is called when production role_list is entered.
func (s *BaseGaussdbParserListener) EnterRole_list(ctx *Role_listContext) {}

// ExitRole_list is called when production role_list is exited.
func (s *BaseGaussdbParserListener) ExitRole_list(ctx *Role_listContext) {}

// EnterColid is called when production colid is entered.
func (s *BaseGaussdbParserListener) EnterColid(ctx *ColidContext) {}

// ExitColid is called when production colid is exited.
func (s *BaseGaussdbParserListener) ExitColid(ctx *ColidContext) {}

// EnterTable_alias is called when production table_alias is entered.
func (s *BaseGaussdbParserListener) EnterTable_alias(ctx *Table_aliasContext) {}

// ExitTable_alias is called when production table_alias is exited.
func (s *BaseGaussdbParserListener) ExitTable_alias(ctx *Table_aliasContext) {}

// EnterType_function_name is called when production type_function_name is entered.
func (s *BaseGaussdbParserListener) EnterType_function_name(ctx *Type_function_nameContext) {}

// ExitType_function_name is called when production type_function_name is exited.
func (s *BaseGaussdbParserListener) ExitType_function_name(ctx *Type_function_nameContext) {}

// EnterNonreservedword is called when production nonreservedword is entered.
func (s *BaseGaussdbParserListener) EnterNonreservedword(ctx *NonreservedwordContext) {}

// ExitNonreservedword is called when production nonreservedword is exited.
func (s *BaseGaussdbParserListener) ExitNonreservedword(ctx *NonreservedwordContext) {}

// EnterCollabel is called when production collabel is entered.
func (s *BaseGaussdbParserListener) EnterCollabel(ctx *CollabelContext) {}

// ExitCollabel is called when production collabel is exited.
func (s *BaseGaussdbParserListener) ExitCollabel(ctx *CollabelContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseGaussdbParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseGaussdbParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterPlsqlidentifier is called when production plsqlidentifier is entered.
func (s *BaseGaussdbParserListener) EnterPlsqlidentifier(ctx *PlsqlidentifierContext) {}

// ExitPlsqlidentifier is called when production plsqlidentifier is exited.
func (s *BaseGaussdbParserListener) ExitPlsqlidentifier(ctx *PlsqlidentifierContext) {}

// EnterUnreserved_keyword is called when production unreserved_keyword is entered.
func (s *BaseGaussdbParserListener) EnterUnreserved_keyword(ctx *Unreserved_keywordContext) {}

// ExitUnreserved_keyword is called when production unreserved_keyword is exited.
func (s *BaseGaussdbParserListener) ExitUnreserved_keyword(ctx *Unreserved_keywordContext) {}

// EnterCol_name_keyword is called when production col_name_keyword is entered.
func (s *BaseGaussdbParserListener) EnterCol_name_keyword(ctx *Col_name_keywordContext) {}

// ExitCol_name_keyword is called when production col_name_keyword is exited.
func (s *BaseGaussdbParserListener) ExitCol_name_keyword(ctx *Col_name_keywordContext) {}

// EnterType_func_name_keyword is called when production type_func_name_keyword is entered.
func (s *BaseGaussdbParserListener) EnterType_func_name_keyword(ctx *Type_func_name_keywordContext) {}

// ExitType_func_name_keyword is called when production type_func_name_keyword is exited.
func (s *BaseGaussdbParserListener) ExitType_func_name_keyword(ctx *Type_func_name_keywordContext) {}

// EnterReserved_keyword is called when production reserved_keyword is entered.
func (s *BaseGaussdbParserListener) EnterReserved_keyword(ctx *Reserved_keywordContext) {}

// ExitReserved_keyword is called when production reserved_keyword is exited.
func (s *BaseGaussdbParserListener) ExitReserved_keyword(ctx *Reserved_keywordContext) {}

// EnterBuiltin_function_name is called when production builtin_function_name is entered.
func (s *BaseGaussdbParserListener) EnterBuiltin_function_name(ctx *Builtin_function_nameContext) {}

// ExitBuiltin_function_name is called when production builtin_function_name is exited.
func (s *BaseGaussdbParserListener) ExitBuiltin_function_name(ctx *Builtin_function_nameContext) {}

// EnterPl_function is called when production pl_function is entered.
func (s *BaseGaussdbParserListener) EnterPl_function(ctx *Pl_functionContext) {}

// ExitPl_function is called when production pl_function is exited.
func (s *BaseGaussdbParserListener) ExitPl_function(ctx *Pl_functionContext) {}

// EnterComp_options is called when production comp_options is entered.
func (s *BaseGaussdbParserListener) EnterComp_options(ctx *Comp_optionsContext) {}

// ExitComp_options is called when production comp_options is exited.
func (s *BaseGaussdbParserListener) ExitComp_options(ctx *Comp_optionsContext) {}

// EnterComp_option is called when production comp_option is entered.
func (s *BaseGaussdbParserListener) EnterComp_option(ctx *Comp_optionContext) {}

// ExitComp_option is called when production comp_option is exited.
func (s *BaseGaussdbParserListener) ExitComp_option(ctx *Comp_optionContext) {}

// EnterSharp is called when production sharp is entered.
func (s *BaseGaussdbParserListener) EnterSharp(ctx *SharpContext) {}

// ExitSharp is called when production sharp is exited.
func (s *BaseGaussdbParserListener) ExitSharp(ctx *SharpContext) {}

// EnterOption_value is called when production option_value is entered.
func (s *BaseGaussdbParserListener) EnterOption_value(ctx *Option_valueContext) {}

// ExitOption_value is called when production option_value is exited.
func (s *BaseGaussdbParserListener) ExitOption_value(ctx *Option_valueContext) {}

// EnterOpt_semi is called when production opt_semi is entered.
func (s *BaseGaussdbParserListener) EnterOpt_semi(ctx *Opt_semiContext) {}

// ExitOpt_semi is called when production opt_semi is exited.
func (s *BaseGaussdbParserListener) ExitOpt_semi(ctx *Opt_semiContext) {}

// EnterPl_block is called when production pl_block is entered.
func (s *BaseGaussdbParserListener) EnterPl_block(ctx *Pl_blockContext) {}

// ExitPl_block is called when production pl_block is exited.
func (s *BaseGaussdbParserListener) ExitPl_block(ctx *Pl_blockContext) {}

// EnterDecl_sect is called when production decl_sect is entered.
func (s *BaseGaussdbParserListener) EnterDecl_sect(ctx *Decl_sectContext) {}

// ExitDecl_sect is called when production decl_sect is exited.
func (s *BaseGaussdbParserListener) ExitDecl_sect(ctx *Decl_sectContext) {}

// EnterDecl_start is called when production decl_start is entered.
func (s *BaseGaussdbParserListener) EnterDecl_start(ctx *Decl_startContext) {}

// ExitDecl_start is called when production decl_start is exited.
func (s *BaseGaussdbParserListener) ExitDecl_start(ctx *Decl_startContext) {}

// EnterDecl_stmts is called when production decl_stmts is entered.
func (s *BaseGaussdbParserListener) EnterDecl_stmts(ctx *Decl_stmtsContext) {}

// ExitDecl_stmts is called when production decl_stmts is exited.
func (s *BaseGaussdbParserListener) ExitDecl_stmts(ctx *Decl_stmtsContext) {}

// EnterLabel_decl is called when production label_decl is entered.
func (s *BaseGaussdbParserListener) EnterLabel_decl(ctx *Label_declContext) {}

// ExitLabel_decl is called when production label_decl is exited.
func (s *BaseGaussdbParserListener) ExitLabel_decl(ctx *Label_declContext) {}

// EnterDecl_stmt is called when production decl_stmt is entered.
func (s *BaseGaussdbParserListener) EnterDecl_stmt(ctx *Decl_stmtContext) {}

// ExitDecl_stmt is called when production decl_stmt is exited.
func (s *BaseGaussdbParserListener) ExitDecl_stmt(ctx *Decl_stmtContext) {}

// EnterDecl_statement is called when production decl_statement is entered.
func (s *BaseGaussdbParserListener) EnterDecl_statement(ctx *Decl_statementContext) {}

// ExitDecl_statement is called when production decl_statement is exited.
func (s *BaseGaussdbParserListener) ExitDecl_statement(ctx *Decl_statementContext) {}

// EnterOpt_scrollable is called when production opt_scrollable is entered.
func (s *BaseGaussdbParserListener) EnterOpt_scrollable(ctx *Opt_scrollableContext) {}

// ExitOpt_scrollable is called when production opt_scrollable is exited.
func (s *BaseGaussdbParserListener) ExitOpt_scrollable(ctx *Opt_scrollableContext) {}

// EnterDecl_cursor_query is called when production decl_cursor_query is entered.
func (s *BaseGaussdbParserListener) EnterDecl_cursor_query(ctx *Decl_cursor_queryContext) {}

// ExitDecl_cursor_query is called when production decl_cursor_query is exited.
func (s *BaseGaussdbParserListener) ExitDecl_cursor_query(ctx *Decl_cursor_queryContext) {}

// EnterDecl_cursor_args is called when production decl_cursor_args is entered.
func (s *BaseGaussdbParserListener) EnterDecl_cursor_args(ctx *Decl_cursor_argsContext) {}

// ExitDecl_cursor_args is called when production decl_cursor_args is exited.
func (s *BaseGaussdbParserListener) ExitDecl_cursor_args(ctx *Decl_cursor_argsContext) {}

// EnterDecl_cursor_arglist is called when production decl_cursor_arglist is entered.
func (s *BaseGaussdbParserListener) EnterDecl_cursor_arglist(ctx *Decl_cursor_arglistContext) {}

// ExitDecl_cursor_arglist is called when production decl_cursor_arglist is exited.
func (s *BaseGaussdbParserListener) ExitDecl_cursor_arglist(ctx *Decl_cursor_arglistContext) {}

// EnterDecl_cursor_arg is called when production decl_cursor_arg is entered.
func (s *BaseGaussdbParserListener) EnterDecl_cursor_arg(ctx *Decl_cursor_argContext) {}

// ExitDecl_cursor_arg is called when production decl_cursor_arg is exited.
func (s *BaseGaussdbParserListener) ExitDecl_cursor_arg(ctx *Decl_cursor_argContext) {}

// EnterDecl_is_for is called when production decl_is_for is entered.
func (s *BaseGaussdbParserListener) EnterDecl_is_for(ctx *Decl_is_forContext) {}

// ExitDecl_is_for is called when production decl_is_for is exited.
func (s *BaseGaussdbParserListener) ExitDecl_is_for(ctx *Decl_is_forContext) {}

// EnterDecl_aliasitem is called when production decl_aliasitem is entered.
func (s *BaseGaussdbParserListener) EnterDecl_aliasitem(ctx *Decl_aliasitemContext) {}

// ExitDecl_aliasitem is called when production decl_aliasitem is exited.
func (s *BaseGaussdbParserListener) ExitDecl_aliasitem(ctx *Decl_aliasitemContext) {}

// EnterDecl_varname is called when production decl_varname is entered.
func (s *BaseGaussdbParserListener) EnterDecl_varname(ctx *Decl_varnameContext) {}

// ExitDecl_varname is called when production decl_varname is exited.
func (s *BaseGaussdbParserListener) ExitDecl_varname(ctx *Decl_varnameContext) {}

// EnterDecl_const is called when production decl_const is entered.
func (s *BaseGaussdbParserListener) EnterDecl_const(ctx *Decl_constContext) {}

// ExitDecl_const is called when production decl_const is exited.
func (s *BaseGaussdbParserListener) ExitDecl_const(ctx *Decl_constContext) {}

// EnterDecl_datatype is called when production decl_datatype is entered.
func (s *BaseGaussdbParserListener) EnterDecl_datatype(ctx *Decl_datatypeContext) {}

// ExitDecl_datatype is called when production decl_datatype is exited.
func (s *BaseGaussdbParserListener) ExitDecl_datatype(ctx *Decl_datatypeContext) {}

// EnterDecl_collate is called when production decl_collate is entered.
func (s *BaseGaussdbParserListener) EnterDecl_collate(ctx *Decl_collateContext) {}

// ExitDecl_collate is called when production decl_collate is exited.
func (s *BaseGaussdbParserListener) ExitDecl_collate(ctx *Decl_collateContext) {}

// EnterDecl_notnull is called when production decl_notnull is entered.
func (s *BaseGaussdbParserListener) EnterDecl_notnull(ctx *Decl_notnullContext) {}

// ExitDecl_notnull is called when production decl_notnull is exited.
func (s *BaseGaussdbParserListener) ExitDecl_notnull(ctx *Decl_notnullContext) {}

// EnterDecl_defval is called when production decl_defval is entered.
func (s *BaseGaussdbParserListener) EnterDecl_defval(ctx *Decl_defvalContext) {}

// ExitDecl_defval is called when production decl_defval is exited.
func (s *BaseGaussdbParserListener) ExitDecl_defval(ctx *Decl_defvalContext) {}

// EnterDecl_defkey is called when production decl_defkey is entered.
func (s *BaseGaussdbParserListener) EnterDecl_defkey(ctx *Decl_defkeyContext) {}

// ExitDecl_defkey is called when production decl_defkey is exited.
func (s *BaseGaussdbParserListener) ExitDecl_defkey(ctx *Decl_defkeyContext) {}

// EnterAssign_operator is called when production assign_operator is entered.
func (s *BaseGaussdbParserListener) EnterAssign_operator(ctx *Assign_operatorContext) {}

// ExitAssign_operator is called when production assign_operator is exited.
func (s *BaseGaussdbParserListener) ExitAssign_operator(ctx *Assign_operatorContext) {}

// EnterProc_sect is called when production proc_sect is entered.
func (s *BaseGaussdbParserListener) EnterProc_sect(ctx *Proc_sectContext) {}

// ExitProc_sect is called when production proc_sect is exited.
func (s *BaseGaussdbParserListener) ExitProc_sect(ctx *Proc_sectContext) {}

// EnterProc_stmt is called when production proc_stmt is entered.
func (s *BaseGaussdbParserListener) EnterProc_stmt(ctx *Proc_stmtContext) {}

// ExitProc_stmt is called when production proc_stmt is exited.
func (s *BaseGaussdbParserListener) ExitProc_stmt(ctx *Proc_stmtContext) {}

// EnterStmt_perform is called when production stmt_perform is entered.
func (s *BaseGaussdbParserListener) EnterStmt_perform(ctx *Stmt_performContext) {}

// ExitStmt_perform is called when production stmt_perform is exited.
func (s *BaseGaussdbParserListener) ExitStmt_perform(ctx *Stmt_performContext) {}

// EnterStmt_call is called when production stmt_call is entered.
func (s *BaseGaussdbParserListener) EnterStmt_call(ctx *Stmt_callContext) {}

// ExitStmt_call is called when production stmt_call is exited.
func (s *BaseGaussdbParserListener) ExitStmt_call(ctx *Stmt_callContext) {}

// EnterOpt_expr_list is called when production opt_expr_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_expr_list(ctx *Opt_expr_listContext) {}

// ExitOpt_expr_list is called when production opt_expr_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_expr_list(ctx *Opt_expr_listContext) {}

// EnterStmt_assign is called when production stmt_assign is entered.
func (s *BaseGaussdbParserListener) EnterStmt_assign(ctx *Stmt_assignContext) {}

// ExitStmt_assign is called when production stmt_assign is exited.
func (s *BaseGaussdbParserListener) ExitStmt_assign(ctx *Stmt_assignContext) {}

// EnterStmt_getdiag is called when production stmt_getdiag is entered.
func (s *BaseGaussdbParserListener) EnterStmt_getdiag(ctx *Stmt_getdiagContext) {}

// ExitStmt_getdiag is called when production stmt_getdiag is exited.
func (s *BaseGaussdbParserListener) ExitStmt_getdiag(ctx *Stmt_getdiagContext) {}

// EnterGetdiag_area_opt is called when production getdiag_area_opt is entered.
func (s *BaseGaussdbParserListener) EnterGetdiag_area_opt(ctx *Getdiag_area_optContext) {}

// ExitGetdiag_area_opt is called when production getdiag_area_opt is exited.
func (s *BaseGaussdbParserListener) ExitGetdiag_area_opt(ctx *Getdiag_area_optContext) {}

// EnterGetdiag_list is called when production getdiag_list is entered.
func (s *BaseGaussdbParserListener) EnterGetdiag_list(ctx *Getdiag_listContext) {}

// ExitGetdiag_list is called when production getdiag_list is exited.
func (s *BaseGaussdbParserListener) ExitGetdiag_list(ctx *Getdiag_listContext) {}

// EnterGetdiag_list_item is called when production getdiag_list_item is entered.
func (s *BaseGaussdbParserListener) EnterGetdiag_list_item(ctx *Getdiag_list_itemContext) {}

// ExitGetdiag_list_item is called when production getdiag_list_item is exited.
func (s *BaseGaussdbParserListener) ExitGetdiag_list_item(ctx *Getdiag_list_itemContext) {}

// EnterGetdiag_item is called when production getdiag_item is entered.
func (s *BaseGaussdbParserListener) EnterGetdiag_item(ctx *Getdiag_itemContext) {}

// ExitGetdiag_item is called when production getdiag_item is exited.
func (s *BaseGaussdbParserListener) ExitGetdiag_item(ctx *Getdiag_itemContext) {}

// EnterGetdiag_target is called when production getdiag_target is entered.
func (s *BaseGaussdbParserListener) EnterGetdiag_target(ctx *Getdiag_targetContext) {}

// ExitGetdiag_target is called when production getdiag_target is exited.
func (s *BaseGaussdbParserListener) ExitGetdiag_target(ctx *Getdiag_targetContext) {}

// EnterAssign_var is called when production assign_var is entered.
func (s *BaseGaussdbParserListener) EnterAssign_var(ctx *Assign_varContext) {}

// ExitAssign_var is called when production assign_var is exited.
func (s *BaseGaussdbParserListener) ExitAssign_var(ctx *Assign_varContext) {}

// EnterStmt_if is called when production stmt_if is entered.
func (s *BaseGaussdbParserListener) EnterStmt_if(ctx *Stmt_ifContext) {}

// ExitStmt_if is called when production stmt_if is exited.
func (s *BaseGaussdbParserListener) ExitStmt_if(ctx *Stmt_ifContext) {}

// EnterStmt_elsifs is called when production stmt_elsifs is entered.
func (s *BaseGaussdbParserListener) EnterStmt_elsifs(ctx *Stmt_elsifsContext) {}

// ExitStmt_elsifs is called when production stmt_elsifs is exited.
func (s *BaseGaussdbParserListener) ExitStmt_elsifs(ctx *Stmt_elsifsContext) {}

// EnterStmt_else is called when production stmt_else is entered.
func (s *BaseGaussdbParserListener) EnterStmt_else(ctx *Stmt_elseContext) {}

// ExitStmt_else is called when production stmt_else is exited.
func (s *BaseGaussdbParserListener) ExitStmt_else(ctx *Stmt_elseContext) {}

// EnterStmt_case is called when production stmt_case is entered.
func (s *BaseGaussdbParserListener) EnterStmt_case(ctx *Stmt_caseContext) {}

// ExitStmt_case is called when production stmt_case is exited.
func (s *BaseGaussdbParserListener) ExitStmt_case(ctx *Stmt_caseContext) {}

// EnterOpt_expr_until_when is called when production opt_expr_until_when is entered.
func (s *BaseGaussdbParserListener) EnterOpt_expr_until_when(ctx *Opt_expr_until_whenContext) {}

// ExitOpt_expr_until_when is called when production opt_expr_until_when is exited.
func (s *BaseGaussdbParserListener) ExitOpt_expr_until_when(ctx *Opt_expr_until_whenContext) {}

// EnterCase_when_list is called when production case_when_list is entered.
func (s *BaseGaussdbParserListener) EnterCase_when_list(ctx *Case_when_listContext) {}

// ExitCase_when_list is called when production case_when_list is exited.
func (s *BaseGaussdbParserListener) ExitCase_when_list(ctx *Case_when_listContext) {}

// EnterCase_when is called when production case_when is entered.
func (s *BaseGaussdbParserListener) EnterCase_when(ctx *Case_whenContext) {}

// ExitCase_when is called when production case_when is exited.
func (s *BaseGaussdbParserListener) ExitCase_when(ctx *Case_whenContext) {}

// EnterOpt_case_else is called when production opt_case_else is entered.
func (s *BaseGaussdbParserListener) EnterOpt_case_else(ctx *Opt_case_elseContext) {}

// ExitOpt_case_else is called when production opt_case_else is exited.
func (s *BaseGaussdbParserListener) ExitOpt_case_else(ctx *Opt_case_elseContext) {}

// EnterStmt_loop is called when production stmt_loop is entered.
func (s *BaseGaussdbParserListener) EnterStmt_loop(ctx *Stmt_loopContext) {}

// ExitStmt_loop is called when production stmt_loop is exited.
func (s *BaseGaussdbParserListener) ExitStmt_loop(ctx *Stmt_loopContext) {}

// EnterStmt_while is called when production stmt_while is entered.
func (s *BaseGaussdbParserListener) EnterStmt_while(ctx *Stmt_whileContext) {}

// ExitStmt_while is called when production stmt_while is exited.
func (s *BaseGaussdbParserListener) ExitStmt_while(ctx *Stmt_whileContext) {}

// EnterStmt_for is called when production stmt_for is entered.
func (s *BaseGaussdbParserListener) EnterStmt_for(ctx *Stmt_forContext) {}

// ExitStmt_for is called when production stmt_for is exited.
func (s *BaseGaussdbParserListener) ExitStmt_for(ctx *Stmt_forContext) {}

// EnterFor_control is called when production for_control is entered.
func (s *BaseGaussdbParserListener) EnterFor_control(ctx *For_controlContext) {}

// ExitFor_control is called when production for_control is exited.
func (s *BaseGaussdbParserListener) ExitFor_control(ctx *For_controlContext) {}

// EnterOpt_for_using_expression is called when production opt_for_using_expression is entered.
func (s *BaseGaussdbParserListener) EnterOpt_for_using_expression(ctx *Opt_for_using_expressionContext) {
}

// ExitOpt_for_using_expression is called when production opt_for_using_expression is exited.
func (s *BaseGaussdbParserListener) ExitOpt_for_using_expression(ctx *Opt_for_using_expressionContext) {
}

// EnterOpt_cursor_parameters is called when production opt_cursor_parameters is entered.
func (s *BaseGaussdbParserListener) EnterOpt_cursor_parameters(ctx *Opt_cursor_parametersContext) {}

// ExitOpt_cursor_parameters is called when production opt_cursor_parameters is exited.
func (s *BaseGaussdbParserListener) ExitOpt_cursor_parameters(ctx *Opt_cursor_parametersContext) {}

// EnterOpt_reverse is called when production opt_reverse is entered.
func (s *BaseGaussdbParserListener) EnterOpt_reverse(ctx *Opt_reverseContext) {}

// ExitOpt_reverse is called when production opt_reverse is exited.
func (s *BaseGaussdbParserListener) ExitOpt_reverse(ctx *Opt_reverseContext) {}

// EnterOpt_by_expression is called when production opt_by_expression is entered.
func (s *BaseGaussdbParserListener) EnterOpt_by_expression(ctx *Opt_by_expressionContext) {}

// ExitOpt_by_expression is called when production opt_by_expression is exited.
func (s *BaseGaussdbParserListener) ExitOpt_by_expression(ctx *Opt_by_expressionContext) {}

// EnterFor_variable is called when production for_variable is entered.
func (s *BaseGaussdbParserListener) EnterFor_variable(ctx *For_variableContext) {}

// ExitFor_variable is called when production for_variable is exited.
func (s *BaseGaussdbParserListener) ExitFor_variable(ctx *For_variableContext) {}

// EnterStmt_foreach_a is called when production stmt_foreach_a is entered.
func (s *BaseGaussdbParserListener) EnterStmt_foreach_a(ctx *Stmt_foreach_aContext) {}

// ExitStmt_foreach_a is called when production stmt_foreach_a is exited.
func (s *BaseGaussdbParserListener) ExitStmt_foreach_a(ctx *Stmt_foreach_aContext) {}

// EnterForeach_slice is called when production foreach_slice is entered.
func (s *BaseGaussdbParserListener) EnterForeach_slice(ctx *Foreach_sliceContext) {}

// ExitForeach_slice is called when production foreach_slice is exited.
func (s *BaseGaussdbParserListener) ExitForeach_slice(ctx *Foreach_sliceContext) {}

// EnterStmt_exit is called when production stmt_exit is entered.
func (s *BaseGaussdbParserListener) EnterStmt_exit(ctx *Stmt_exitContext) {}

// ExitStmt_exit is called when production stmt_exit is exited.
func (s *BaseGaussdbParserListener) ExitStmt_exit(ctx *Stmt_exitContext) {}

// EnterExit_type is called when production exit_type is entered.
func (s *BaseGaussdbParserListener) EnterExit_type(ctx *Exit_typeContext) {}

// ExitExit_type is called when production exit_type is exited.
func (s *BaseGaussdbParserListener) ExitExit_type(ctx *Exit_typeContext) {}

// EnterStmt_return is called when production stmt_return is entered.
func (s *BaseGaussdbParserListener) EnterStmt_return(ctx *Stmt_returnContext) {}

// ExitStmt_return is called when production stmt_return is exited.
func (s *BaseGaussdbParserListener) ExitStmt_return(ctx *Stmt_returnContext) {}

// EnterOpt_return_result is called when production opt_return_result is entered.
func (s *BaseGaussdbParserListener) EnterOpt_return_result(ctx *Opt_return_resultContext) {}

// ExitOpt_return_result is called when production opt_return_result is exited.
func (s *BaseGaussdbParserListener) ExitOpt_return_result(ctx *Opt_return_resultContext) {}

// EnterStmt_raise is called when production stmt_raise is entered.
func (s *BaseGaussdbParserListener) EnterStmt_raise(ctx *Stmt_raiseContext) {}

// ExitStmt_raise is called when production stmt_raise is exited.
func (s *BaseGaussdbParserListener) ExitStmt_raise(ctx *Stmt_raiseContext) {}

// EnterOpt_stmt_raise_level is called when production opt_stmt_raise_level is entered.
func (s *BaseGaussdbParserListener) EnterOpt_stmt_raise_level(ctx *Opt_stmt_raise_levelContext) {}

// ExitOpt_stmt_raise_level is called when production opt_stmt_raise_level is exited.
func (s *BaseGaussdbParserListener) ExitOpt_stmt_raise_level(ctx *Opt_stmt_raise_levelContext) {}

// EnterOpt_raise_list is called when production opt_raise_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_raise_list(ctx *Opt_raise_listContext) {}

// ExitOpt_raise_list is called when production opt_raise_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_raise_list(ctx *Opt_raise_listContext) {}

// EnterOpt_raise_using is called when production opt_raise_using is entered.
func (s *BaseGaussdbParserListener) EnterOpt_raise_using(ctx *Opt_raise_usingContext) {}

// ExitOpt_raise_using is called when production opt_raise_using is exited.
func (s *BaseGaussdbParserListener) ExitOpt_raise_using(ctx *Opt_raise_usingContext) {}

// EnterOpt_raise_using_elem is called when production opt_raise_using_elem is entered.
func (s *BaseGaussdbParserListener) EnterOpt_raise_using_elem(ctx *Opt_raise_using_elemContext) {}

// ExitOpt_raise_using_elem is called when production opt_raise_using_elem is exited.
func (s *BaseGaussdbParserListener) ExitOpt_raise_using_elem(ctx *Opt_raise_using_elemContext) {}

// EnterOpt_raise_using_elem_list is called when production opt_raise_using_elem_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_raise_using_elem_list(ctx *Opt_raise_using_elem_listContext) {
}

// ExitOpt_raise_using_elem_list is called when production opt_raise_using_elem_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_raise_using_elem_list(ctx *Opt_raise_using_elem_listContext) {
}

// EnterStmt_assert is called when production stmt_assert is entered.
func (s *BaseGaussdbParserListener) EnterStmt_assert(ctx *Stmt_assertContext) {}

// ExitStmt_assert is called when production stmt_assert is exited.
func (s *BaseGaussdbParserListener) ExitStmt_assert(ctx *Stmt_assertContext) {}

// EnterOpt_stmt_assert_message is called when production opt_stmt_assert_message is entered.
func (s *BaseGaussdbParserListener) EnterOpt_stmt_assert_message(ctx *Opt_stmt_assert_messageContext) {
}

// ExitOpt_stmt_assert_message is called when production opt_stmt_assert_message is exited.
func (s *BaseGaussdbParserListener) ExitOpt_stmt_assert_message(ctx *Opt_stmt_assert_messageContext) {
}

// EnterLoop_body is called when production loop_body is entered.
func (s *BaseGaussdbParserListener) EnterLoop_body(ctx *Loop_bodyContext) {}

// ExitLoop_body is called when production loop_body is exited.
func (s *BaseGaussdbParserListener) ExitLoop_body(ctx *Loop_bodyContext) {}

// EnterStmt_execsql is called when production stmt_execsql is entered.
func (s *BaseGaussdbParserListener) EnterStmt_execsql(ctx *Stmt_execsqlContext) {}

// ExitStmt_execsql is called when production stmt_execsql is exited.
func (s *BaseGaussdbParserListener) ExitStmt_execsql(ctx *Stmt_execsqlContext) {}

// EnterStmt_dynexecute is called when production stmt_dynexecute is entered.
func (s *BaseGaussdbParserListener) EnterStmt_dynexecute(ctx *Stmt_dynexecuteContext) {}

// ExitStmt_dynexecute is called when production stmt_dynexecute is exited.
func (s *BaseGaussdbParserListener) ExitStmt_dynexecute(ctx *Stmt_dynexecuteContext) {}

// EnterOpt_execute_using is called when production opt_execute_using is entered.
func (s *BaseGaussdbParserListener) EnterOpt_execute_using(ctx *Opt_execute_usingContext) {}

// ExitOpt_execute_using is called when production opt_execute_using is exited.
func (s *BaseGaussdbParserListener) ExitOpt_execute_using(ctx *Opt_execute_usingContext) {}

// EnterOpt_execute_using_list is called when production opt_execute_using_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_execute_using_list(ctx *Opt_execute_using_listContext) {}

// ExitOpt_execute_using_list is called when production opt_execute_using_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_execute_using_list(ctx *Opt_execute_using_listContext) {}

// EnterOpt_execute_into is called when production opt_execute_into is entered.
func (s *BaseGaussdbParserListener) EnterOpt_execute_into(ctx *Opt_execute_intoContext) {}

// ExitOpt_execute_into is called when production opt_execute_into is exited.
func (s *BaseGaussdbParserListener) ExitOpt_execute_into(ctx *Opt_execute_intoContext) {}

// EnterStmt_open is called when production stmt_open is entered.
func (s *BaseGaussdbParserListener) EnterStmt_open(ctx *Stmt_openContext) {}

// ExitStmt_open is called when production stmt_open is exited.
func (s *BaseGaussdbParserListener) ExitStmt_open(ctx *Stmt_openContext) {}

// EnterOpt_open_bound_list_item is called when production opt_open_bound_list_item is entered.
func (s *BaseGaussdbParserListener) EnterOpt_open_bound_list_item(ctx *Opt_open_bound_list_itemContext) {
}

// ExitOpt_open_bound_list_item is called when production opt_open_bound_list_item is exited.
func (s *BaseGaussdbParserListener) ExitOpt_open_bound_list_item(ctx *Opt_open_bound_list_itemContext) {
}

// EnterOpt_open_bound_list is called when production opt_open_bound_list is entered.
func (s *BaseGaussdbParserListener) EnterOpt_open_bound_list(ctx *Opt_open_bound_listContext) {}

// ExitOpt_open_bound_list is called when production opt_open_bound_list is exited.
func (s *BaseGaussdbParserListener) ExitOpt_open_bound_list(ctx *Opt_open_bound_listContext) {}

// EnterOpt_open_using is called when production opt_open_using is entered.
func (s *BaseGaussdbParserListener) EnterOpt_open_using(ctx *Opt_open_usingContext) {}

// ExitOpt_open_using is called when production opt_open_using is exited.
func (s *BaseGaussdbParserListener) ExitOpt_open_using(ctx *Opt_open_usingContext) {}

// EnterOpt_scroll_option is called when production opt_scroll_option is entered.
func (s *BaseGaussdbParserListener) EnterOpt_scroll_option(ctx *Opt_scroll_optionContext) {}

// ExitOpt_scroll_option is called when production opt_scroll_option is exited.
func (s *BaseGaussdbParserListener) ExitOpt_scroll_option(ctx *Opt_scroll_optionContext) {}

// EnterOpt_scroll_option_no is called when production opt_scroll_option_no is entered.
func (s *BaseGaussdbParserListener) EnterOpt_scroll_option_no(ctx *Opt_scroll_option_noContext) {}

// ExitOpt_scroll_option_no is called when production opt_scroll_option_no is exited.
func (s *BaseGaussdbParserListener) ExitOpt_scroll_option_no(ctx *Opt_scroll_option_noContext) {}

// EnterStmt_fetch is called when production stmt_fetch is entered.
func (s *BaseGaussdbParserListener) EnterStmt_fetch(ctx *Stmt_fetchContext) {}

// ExitStmt_fetch is called when production stmt_fetch is exited.
func (s *BaseGaussdbParserListener) ExitStmt_fetch(ctx *Stmt_fetchContext) {}

// EnterInto_target is called when production into_target is entered.
func (s *BaseGaussdbParserListener) EnterInto_target(ctx *Into_targetContext) {}

// ExitInto_target is called when production into_target is exited.
func (s *BaseGaussdbParserListener) ExitInto_target(ctx *Into_targetContext) {}

// EnterOpt_cursor_from is called when production opt_cursor_from is entered.
func (s *BaseGaussdbParserListener) EnterOpt_cursor_from(ctx *Opt_cursor_fromContext) {}

// ExitOpt_cursor_from is called when production opt_cursor_from is exited.
func (s *BaseGaussdbParserListener) ExitOpt_cursor_from(ctx *Opt_cursor_fromContext) {}

// EnterOpt_fetch_direction is called when production opt_fetch_direction is entered.
func (s *BaseGaussdbParserListener) EnterOpt_fetch_direction(ctx *Opt_fetch_directionContext) {}

// ExitOpt_fetch_direction is called when production opt_fetch_direction is exited.
func (s *BaseGaussdbParserListener) ExitOpt_fetch_direction(ctx *Opt_fetch_directionContext) {}

// EnterStmt_move is called when production stmt_move is entered.
func (s *BaseGaussdbParserListener) EnterStmt_move(ctx *Stmt_moveContext) {}

// ExitStmt_move is called when production stmt_move is exited.
func (s *BaseGaussdbParserListener) ExitStmt_move(ctx *Stmt_moveContext) {}

// EnterStmt_close is called when production stmt_close is entered.
func (s *BaseGaussdbParserListener) EnterStmt_close(ctx *Stmt_closeContext) {}

// ExitStmt_close is called when production stmt_close is exited.
func (s *BaseGaussdbParserListener) ExitStmt_close(ctx *Stmt_closeContext) {}

// EnterStmt_null is called when production stmt_null is entered.
func (s *BaseGaussdbParserListener) EnterStmt_null(ctx *Stmt_nullContext) {}

// ExitStmt_null is called when production stmt_null is exited.
func (s *BaseGaussdbParserListener) ExitStmt_null(ctx *Stmt_nullContext) {}

// EnterStmt_commit is called when production stmt_commit is entered.
func (s *BaseGaussdbParserListener) EnterStmt_commit(ctx *Stmt_commitContext) {}

// ExitStmt_commit is called when production stmt_commit is exited.
func (s *BaseGaussdbParserListener) ExitStmt_commit(ctx *Stmt_commitContext) {}

// EnterStmt_rollback is called when production stmt_rollback is entered.
func (s *BaseGaussdbParserListener) EnterStmt_rollback(ctx *Stmt_rollbackContext) {}

// ExitStmt_rollback is called when production stmt_rollback is exited.
func (s *BaseGaussdbParserListener) ExitStmt_rollback(ctx *Stmt_rollbackContext) {}

// EnterPlsql_opt_transaction_chain is called when production plsql_opt_transaction_chain is entered.
func (s *BaseGaussdbParserListener) EnterPlsql_opt_transaction_chain(ctx *Plsql_opt_transaction_chainContext) {
}

// ExitPlsql_opt_transaction_chain is called when production plsql_opt_transaction_chain is exited.
func (s *BaseGaussdbParserListener) ExitPlsql_opt_transaction_chain(ctx *Plsql_opt_transaction_chainContext) {
}

// EnterStmt_set is called when production stmt_set is entered.
func (s *BaseGaussdbParserListener) EnterStmt_set(ctx *Stmt_setContext) {}

// ExitStmt_set is called when production stmt_set is exited.
func (s *BaseGaussdbParserListener) ExitStmt_set(ctx *Stmt_setContext) {}

// EnterCursor_variable is called when production cursor_variable is entered.
func (s *BaseGaussdbParserListener) EnterCursor_variable(ctx *Cursor_variableContext) {}

// ExitCursor_variable is called when production cursor_variable is exited.
func (s *BaseGaussdbParserListener) ExitCursor_variable(ctx *Cursor_variableContext) {}

// EnterException_sect is called when production exception_sect is entered.
func (s *BaseGaussdbParserListener) EnterException_sect(ctx *Exception_sectContext) {}

// ExitException_sect is called when production exception_sect is exited.
func (s *BaseGaussdbParserListener) ExitException_sect(ctx *Exception_sectContext) {}

// EnterProc_exceptions is called when production proc_exceptions is entered.
func (s *BaseGaussdbParserListener) EnterProc_exceptions(ctx *Proc_exceptionsContext) {}

// ExitProc_exceptions is called when production proc_exceptions is exited.
func (s *BaseGaussdbParserListener) ExitProc_exceptions(ctx *Proc_exceptionsContext) {}

// EnterProc_exception is called when production proc_exception is entered.
func (s *BaseGaussdbParserListener) EnterProc_exception(ctx *Proc_exceptionContext) {}

// ExitProc_exception is called when production proc_exception is exited.
func (s *BaseGaussdbParserListener) ExitProc_exception(ctx *Proc_exceptionContext) {}

// EnterProc_conditions is called when production proc_conditions is entered.
func (s *BaseGaussdbParserListener) EnterProc_conditions(ctx *Proc_conditionsContext) {}

// ExitProc_conditions is called when production proc_conditions is exited.
func (s *BaseGaussdbParserListener) ExitProc_conditions(ctx *Proc_conditionsContext) {}

// EnterProc_condition is called when production proc_condition is entered.
func (s *BaseGaussdbParserListener) EnterProc_condition(ctx *Proc_conditionContext) {}

// ExitProc_condition is called when production proc_condition is exited.
func (s *BaseGaussdbParserListener) ExitProc_condition(ctx *Proc_conditionContext) {}

// EnterOpt_block_label is called when production opt_block_label is entered.
func (s *BaseGaussdbParserListener) EnterOpt_block_label(ctx *Opt_block_labelContext) {}

// ExitOpt_block_label is called when production opt_block_label is exited.
func (s *BaseGaussdbParserListener) ExitOpt_block_label(ctx *Opt_block_labelContext) {}

// EnterOpt_loop_label is called when production opt_loop_label is entered.
func (s *BaseGaussdbParserListener) EnterOpt_loop_label(ctx *Opt_loop_labelContext) {}

// ExitOpt_loop_label is called when production opt_loop_label is exited.
func (s *BaseGaussdbParserListener) ExitOpt_loop_label(ctx *Opt_loop_labelContext) {}

// EnterOpt_label is called when production opt_label is entered.
func (s *BaseGaussdbParserListener) EnterOpt_label(ctx *Opt_labelContext) {}

// ExitOpt_label is called when production opt_label is exited.
func (s *BaseGaussdbParserListener) ExitOpt_label(ctx *Opt_labelContext) {}

// EnterOpt_exitcond is called when production opt_exitcond is entered.
func (s *BaseGaussdbParserListener) EnterOpt_exitcond(ctx *Opt_exitcondContext) {}

// ExitOpt_exitcond is called when production opt_exitcond is exited.
func (s *BaseGaussdbParserListener) ExitOpt_exitcond(ctx *Opt_exitcondContext) {}

// EnterAny_identifier is called when production any_identifier is entered.
func (s *BaseGaussdbParserListener) EnterAny_identifier(ctx *Any_identifierContext) {}

// ExitAny_identifier is called when production any_identifier is exited.
func (s *BaseGaussdbParserListener) ExitAny_identifier(ctx *Any_identifierContext) {}

// EnterPlsql_unreserved_keyword is called when production plsql_unreserved_keyword is entered.
func (s *BaseGaussdbParserListener) EnterPlsql_unreserved_keyword(ctx *Plsql_unreserved_keywordContext) {
}

// ExitPlsql_unreserved_keyword is called when production plsql_unreserved_keyword is exited.
func (s *BaseGaussdbParserListener) ExitPlsql_unreserved_keyword(ctx *Plsql_unreserved_keywordContext) {
}

// EnterSql_expression is called when production sql_expression is entered.
func (s *BaseGaussdbParserListener) EnterSql_expression(ctx *Sql_expressionContext) {}

// ExitSql_expression is called when production sql_expression is exited.
func (s *BaseGaussdbParserListener) ExitSql_expression(ctx *Sql_expressionContext) {}

// EnterExpr_until_then is called when production expr_until_then is entered.
func (s *BaseGaussdbParserListener) EnterExpr_until_then(ctx *Expr_until_thenContext) {}

// ExitExpr_until_then is called when production expr_until_then is exited.
func (s *BaseGaussdbParserListener) ExitExpr_until_then(ctx *Expr_until_thenContext) {}

// EnterExpr_until_semi is called when production expr_until_semi is entered.
func (s *BaseGaussdbParserListener) EnterExpr_until_semi(ctx *Expr_until_semiContext) {}

// ExitExpr_until_semi is called when production expr_until_semi is exited.
func (s *BaseGaussdbParserListener) ExitExpr_until_semi(ctx *Expr_until_semiContext) {}

// EnterExpr_until_rightbracket is called when production expr_until_rightbracket is entered.
func (s *BaseGaussdbParserListener) EnterExpr_until_rightbracket(ctx *Expr_until_rightbracketContext) {
}

// ExitExpr_until_rightbracket is called when production expr_until_rightbracket is exited.
func (s *BaseGaussdbParserListener) ExitExpr_until_rightbracket(ctx *Expr_until_rightbracketContext) {
}

// EnterExpr_until_loop is called when production expr_until_loop is entered.
func (s *BaseGaussdbParserListener) EnterExpr_until_loop(ctx *Expr_until_loopContext) {}

// ExitExpr_until_loop is called when production expr_until_loop is exited.
func (s *BaseGaussdbParserListener) ExitExpr_until_loop(ctx *Expr_until_loopContext) {}

// EnterMake_execsql_stmt is called when production make_execsql_stmt is entered.
func (s *BaseGaussdbParserListener) EnterMake_execsql_stmt(ctx *Make_execsql_stmtContext) {}

// ExitMake_execsql_stmt is called when production make_execsql_stmt is exited.
func (s *BaseGaussdbParserListener) ExitMake_execsql_stmt(ctx *Make_execsql_stmtContext) {}

// EnterOpt_returning_clause_into is called when production opt_returning_clause_into is entered.
func (s *BaseGaussdbParserListener) EnterOpt_returning_clause_into(ctx *Opt_returning_clause_intoContext) {
}

// ExitOpt_returning_clause_into is called when production opt_returning_clause_into is exited.
func (s *BaseGaussdbParserListener) ExitOpt_returning_clause_into(ctx *Opt_returning_clause_intoContext) {
}

// EnterGaussdb_create_generic_position is called when production gaussdb_create_generic_position is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_generic_position(ctx *Gaussdb_create_generic_positionContext) {
}

// ExitGaussdb_create_generic_position is called when production gaussdb_create_generic_position is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_generic_position(ctx *Gaussdb_create_generic_positionContext) {
}

// EnterGaussdb_create_foreigntable_extend is called when production gaussdb_create_foreigntable_extend is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_foreigntable_extend(ctx *Gaussdb_create_foreigntable_extendContext) {
}

// ExitGaussdb_create_foreigntable_extend is called when production gaussdb_create_foreigntable_extend is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_foreigntable_extend(ctx *Gaussdb_create_foreigntable_extendContext) {
}

// EnterGaussdb_error_table_name is called when production gaussdb_error_table_name is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_error_table_name(ctx *Gaussdb_error_table_nameContext) {
}

// ExitGaussdb_error_table_name is called when production gaussdb_error_table_name is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_error_table_name(ctx *Gaussdb_error_table_nameContext) {
}

// EnterGaussdb_nodenamelist is called when production gaussdb_nodenamelist is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_nodenamelist(ctx *Gaussdb_nodenamelistContext) {}

// ExitGaussdb_nodenamelist is called when production gaussdb_nodenamelist is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_nodenamelist(ctx *Gaussdb_nodenamelistContext) {}

// EnterGaussdb_groupname is called when production gaussdb_groupname is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_groupname(ctx *Gaussdb_groupnameContext) {}

// ExitGaussdb_groupname is called when production gaussdb_groupname is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_groupname(ctx *Gaussdb_groupnameContext) {}

// EnterGaussdb_to_group is called when production gaussdb_to_group is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_to_group(ctx *Gaussdb_to_groupContext) {}

// ExitGaussdb_to_group is called when production gaussdb_to_group is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_to_group(ctx *Gaussdb_to_groupContext) {}

// EnterGaussdb_limit_clause is called when production gaussdb_limit_clause is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_limit_clause(ctx *Gaussdb_limit_clauseContext) {}

// ExitGaussdb_limit_clause is called when production gaussdb_limit_clause is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_limit_clause(ctx *Gaussdb_limit_clauseContext) {}

// EnterGaussdb_sort_clause is called when production gaussdb_sort_clause is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_sort_clause(ctx *Gaussdb_sort_clauseContext) {}

// ExitGaussdb_sort_clause is called when production gaussdb_sort_clause is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_sort_clause(ctx *Gaussdb_sort_clauseContext) {}

// EnterGaussdb_from_clause is called when production gaussdb_from_clause is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_from_clause(ctx *Gaussdb_from_clauseContext) {}

// ExitGaussdb_from_clause is called when production gaussdb_from_clause is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_from_clause(ctx *Gaussdb_from_clauseContext) {}

// EnterGaussdb_insert_when_clause is called when production gaussdb_insert_when_clause is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_insert_when_clause(ctx *Gaussdb_insert_when_clauseContext) {
}

// ExitGaussdb_insert_when_clause is called when production gaussdb_insert_when_clause is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_insert_when_clause(ctx *Gaussdb_insert_when_clauseContext) {
}

// EnterGaussdb_truncate_table_partition is called when production gaussdb_truncate_table_partition is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_truncate_table_partition(ctx *Gaussdb_truncate_table_partitionContext) {
}

// ExitGaussdb_truncate_table_partition is called when production gaussdb_truncate_table_partition is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_truncate_table_partition(ctx *Gaussdb_truncate_table_partitionContext) {
}

// EnterGaussdb_partition_extended_names is called when production gaussdb_partition_extended_names is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_extended_names(ctx *Gaussdb_partition_extended_namesContext) {
}

// ExitGaussdb_partition_extended_names is called when production gaussdb_partition_extended_names is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_extended_names(ctx *Gaussdb_partition_extended_namesContext) {
}

// EnterGaussdb_partition_name is called when production gaussdb_partition_name is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_name(ctx *Gaussdb_partition_nameContext) {}

// ExitGaussdb_partition_name is called when production gaussdb_partition_name is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_name(ctx *Gaussdb_partition_nameContext) {}

// EnterGaussdb_new_partition_name is called when production gaussdb_new_partition_name is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_new_partition_name(ctx *Gaussdb_new_partition_nameContext) {
}

// ExitGaussdb_new_partition_name is called when production gaussdb_new_partition_name is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_new_partition_name(ctx *Gaussdb_new_partition_nameContext) {
}

// EnterGaussdb_partition_key_value is called when production gaussdb_partition_key_value is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_key_value(ctx *Gaussdb_partition_key_valueContext) {
}

// ExitGaussdb_partition_key_value is called when production gaussdb_partition_key_value is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_key_value(ctx *Gaussdb_partition_key_valueContext) {
}

// EnterGaussdb_PURGE is called when production gaussdb_PURGE is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_PURGE(ctx *Gaussdb_PURGEContext) {}

// ExitGaussdb_PURGE is called when production gaussdb_PURGE is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_PURGE(ctx *Gaussdb_PURGEContext) {}

// EnterGaussdb_or_replace is called when production gaussdb_or_replace is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_or_replace(ctx *Gaussdb_or_replaceContext) {}

// ExitGaussdb_or_replace is called when production gaussdb_or_replace is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_or_replace(ctx *Gaussdb_or_replaceContext) {}

// EnterGaussdb_index_name is called when production gaussdb_index_name is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_index_name(ctx *Gaussdb_index_nameContext) {}

// ExitGaussdb_index_name is called when production gaussdb_index_name is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_index_name(ctx *Gaussdb_index_nameContext) {}

// EnterGaussdb_comment_list is called when production gaussdb_comment_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_comment_list(ctx *Gaussdb_comment_listContext) {}

// ExitGaussdb_comment_list is called when production gaussdb_comment_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_comment_list(ctx *Gaussdb_comment_listContext) {}

// EnterGaussdb_comment is called when production gaussdb_comment is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_comment(ctx *Gaussdb_commentContext) {}

// ExitGaussdb_comment is called when production gaussdb_comment is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_comment(ctx *Gaussdb_commentContext) {}

// EnterGaussdb_alter_index_ops_set is called when production gaussdb_alter_index_ops_set is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_alter_index_ops_set(ctx *Gaussdb_alter_index_ops_setContext) {
}

// ExitGaussdb_alter_index_ops_set is called when production gaussdb_alter_index_ops_set is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_alter_index_ops_set(ctx *Gaussdb_alter_index_ops_setContext) {
}

// EnterGaussdb_rebuild_clause is called when production gaussdb_rebuild_clause is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_rebuild_clause(ctx *Gaussdb_rebuild_clauseContext) {}

// ExitGaussdb_rebuild_clause is called when production gaussdb_rebuild_clause is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_rebuild_clause(ctx *Gaussdb_rebuild_clauseContext) {}

// EnterGaussdb_alter_index_partitioning is called when production gaussdb_alter_index_partitioning is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_alter_index_partitioning(ctx *Gaussdb_alter_index_partitioningContext) {
}

// ExitGaussdb_alter_index_partitioning is called when production gaussdb_alter_index_partitioning is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_alter_index_partitioning(ctx *Gaussdb_alter_index_partitioningContext) {
}

// EnterGaussdb_modify_index_partition is called when production gaussdb_modify_index_partition is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_modify_index_partition(ctx *Gaussdb_modify_index_partitionContext) {
}

// ExitGaussdb_modify_index_partition is called when production gaussdb_modify_index_partition is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_modify_index_partition(ctx *Gaussdb_modify_index_partitionContext) {
}

// EnterGaussdb_rename_index_partition is called when production gaussdb_rename_index_partition is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_rename_index_partition(ctx *Gaussdb_rename_index_partitionContext) {
}

// ExitGaussdb_rename_index_partition is called when production gaussdb_rename_index_partition is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_rename_index_partition(ctx *Gaussdb_rename_index_partitionContext) {
}

// EnterGaussdb_move_index_partition is called when production gaussdb_move_index_partition is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_move_index_partition(ctx *Gaussdb_move_index_partitionContext) {
}

// ExitGaussdb_move_index_partition is called when production gaussdb_move_index_partition is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_move_index_partition(ctx *Gaussdb_move_index_partitionContext) {
}

// EnterGaussdb_online is called when production gaussdb_online is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_online(ctx *Gaussdb_onlineContext) {}

// ExitGaussdb_online is called when production gaussdb_online is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_online(ctx *Gaussdb_onlineContext) {}

// EnterGaussdb_index_partition_name is called when production gaussdb_index_partition_name is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_index_partition_name(ctx *Gaussdb_index_partition_nameContext) {
}

// ExitGaussdb_index_partition_name is called when production gaussdb_index_partition_name is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_index_partition_name(ctx *Gaussdb_index_partition_nameContext) {
}

// EnterGaussdb_partition_value is called when production gaussdb_partition_value is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_value(ctx *Gaussdb_partition_valueContext) {
}

// ExitGaussdb_partition_value is called when production gaussdb_partition_value is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_value(ctx *Gaussdb_partition_valueContext) {
}

// EnterGaussdb_index_partition_tablespace is called when production gaussdb_index_partition_tablespace is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_index_partition_tablespace(ctx *Gaussdb_index_partition_tablespaceContext) {
}

// ExitGaussdb_index_partition_tablespace is called when production gaussdb_index_partition_tablespace is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_index_partition_tablespace(ctx *Gaussdb_index_partition_tablespaceContext) {
}

// EnterGaussdb_index_partition_list is called when production gaussdb_index_partition_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_index_partition_list(ctx *Gaussdb_index_partition_listContext) {
}

// ExitGaussdb_index_partition_list is called when production gaussdb_index_partition_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_index_partition_list(ctx *Gaussdb_index_partition_listContext) {
}

// EnterGaussdb_index_partition is called when production gaussdb_index_partition is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_index_partition(ctx *Gaussdb_index_partitionContext) {
}

// ExitGaussdb_index_partition is called when production gaussdb_index_partition is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_index_partition(ctx *Gaussdb_index_partitionContext) {
}

// EnterGaussdb_partition_value_list is called when production gaussdb_partition_value_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_value_list(ctx *Gaussdb_partition_value_listContext) {
}

// ExitGaussdb_partition_value_list is called when production gaussdb_partition_value_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_value_list(ctx *Gaussdb_partition_value_listContext) {
}

// EnterGaussdb_index_subpartition_list is called when production gaussdb_index_subpartition_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_index_subpartition_list(ctx *Gaussdb_index_subpartition_listContext) {
}

// ExitGaussdb_index_subpartition_list is called when production gaussdb_index_subpartition_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_index_subpartition_list(ctx *Gaussdb_index_subpartition_listContext) {
}

// EnterGaussdb_index_subppartition is called when production gaussdb_index_subppartition is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_index_subppartition(ctx *Gaussdb_index_subppartitionContext) {
}

// ExitGaussdb_index_subppartition is called when production gaussdb_index_subppartition is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_index_subppartition(ctx *Gaussdb_index_subppartitionContext) {
}

// EnterGaussdb_connect_clause is called when production gaussdb_connect_clause is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_connect_clause(ctx *Gaussdb_connect_clauseContext) {}

// ExitGaussdb_connect_clause is called when production gaussdb_connect_clause is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_connect_clause(ctx *Gaussdb_connect_clauseContext) {}

// EnterGaussdb_engine_list is called when production gaussdb_engine_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_engine_list(ctx *Gaussdb_engine_listContext) {}

// ExitGaussdb_engine_list is called when production gaussdb_engine_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_engine_list(ctx *Gaussdb_engine_listContext) {}

// EnterGaussdb_engine_item is called when production gaussdb_engine_item is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_engine_item(ctx *Gaussdb_engine_itemContext) {}

// ExitGaussdb_engine_item is called when production gaussdb_engine_item is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_engine_item(ctx *Gaussdb_engine_itemContext) {}

// EnterGaussdb_table_option_list is called when production gaussdb_table_option_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_table_option_list(ctx *Gaussdb_table_option_listContext) {
}

// ExitGaussdb_table_option_list is called when production gaussdb_table_option_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_table_option_list(ctx *Gaussdb_table_option_listContext) {
}

// EnterGaussdb_table_option is called when production gaussdb_table_option is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_table_option(ctx *Gaussdb_table_optionContext) {}

// ExitGaussdb_table_option is called when production gaussdb_table_option is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_table_option(ctx *Gaussdb_table_optionContext) {}

// EnterGaussdb_htap_option is called when production gaussdb_htap_option is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_htap_option(ctx *Gaussdb_htap_optionContext) {}

// ExitGaussdb_htap_option is called when production gaussdb_htap_option is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_htap_option(ctx *Gaussdb_htap_optionContext) {}

// EnterGaussdb_distribute is called when production gaussdb_distribute is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_distribute(ctx *Gaussdb_distributeContext) {}

// ExitGaussdb_distribute is called when production gaussdb_distribute is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_distribute(ctx *Gaussdb_distributeContext) {}

// EnterGaussdb_slice_less_than_item_list is called when production gaussdb_slice_less_than_item_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_slice_less_than_item_list(ctx *Gaussdb_slice_less_than_item_listContext) {
}

// ExitGaussdb_slice_less_than_item_list is called when production gaussdb_slice_less_than_item_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_slice_less_than_item_list(ctx *Gaussdb_slice_less_than_item_listContext) {
}

// EnterGaussdb_slice_less_than_item is called when production gaussdb_slice_less_than_item is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_slice_less_than_item(ctx *Gaussdb_slice_less_than_itemContext) {
}

// ExitGaussdb_slice_less_than_item is called when production gaussdb_slice_less_than_item is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_slice_less_than_item(ctx *Gaussdb_slice_less_than_itemContext) {
}

// EnterGaussdb_slice_start_end_item_list is called when production gaussdb_slice_start_end_item_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_slice_start_end_item_list(ctx *Gaussdb_slice_start_end_item_listContext) {
}

// ExitGaussdb_slice_start_end_item_list is called when production gaussdb_slice_start_end_item_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_slice_start_end_item_list(ctx *Gaussdb_slice_start_end_item_listContext) {
}

// EnterGaussdb_slice_start_end_item is called when production gaussdb_slice_start_end_item is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_slice_start_end_item(ctx *Gaussdb_slice_start_end_itemContext) {
}

// ExitGaussdb_slice_start_end_item is called when production gaussdb_slice_start_end_item is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_slice_start_end_item(ctx *Gaussdb_slice_start_end_itemContext) {
}

// EnterGaussdb_slice_values_item_list is called when production gaussdb_slice_values_item_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_slice_values_item_list(ctx *Gaussdb_slice_values_item_listContext) {
}

// ExitGaussdb_slice_values_item_list is called when production gaussdb_slice_values_item_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_slice_values_item_list(ctx *Gaussdb_slice_values_item_listContext) {
}

// EnterGaussdb_slice_values_item_item is called when production gaussdb_slice_values_item_item is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_slice_values_item_item(ctx *Gaussdb_slice_values_item_itemContext) {
}

// ExitGaussdb_slice_values_item_item is called when production gaussdb_slice_values_item_item is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_slice_values_item_item(ctx *Gaussdb_slice_values_item_itemContext) {
}

// EnterGaussdb_ilm_add_policy is called when production gaussdb_ilm_add_policy is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_ilm_add_policy(ctx *Gaussdb_ilm_add_policyContext) {}

// ExitGaussdb_ilm_add_policy is called when production gaussdb_ilm_add_policy is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_ilm_add_policy(ctx *Gaussdb_ilm_add_policyContext) {}

// EnterLabel_name is called when production label_name is entered.
func (s *BaseGaussdbParserListener) EnterLabel_name(ctx *Label_nameContext) {}

// ExitLabel_name is called when production label_name is exited.
func (s *BaseGaussdbParserListener) ExitLabel_name(ctx *Label_nameContext) {}

// EnterException_name is called when production exception_name is entered.
func (s *BaseGaussdbParserListener) EnterException_name(ctx *Exception_nameContext) {}

// ExitException_name is called when production exception_name is exited.
func (s *BaseGaussdbParserListener) ExitException_name(ctx *Exception_nameContext) {}

// EnterImplementation_type_name is called when production implementation_type_name is entered.
func (s *BaseGaussdbParserListener) EnterImplementation_type_name(ctx *Implementation_type_nameContext) {
}

// ExitImplementation_type_name is called when production implementation_type_name is exited.
func (s *BaseGaussdbParserListener) ExitImplementation_type_name(ctx *Implementation_type_nameContext) {
}

// EnterParameter_name is called when production parameter_name is entered.
func (s *BaseGaussdbParserListener) EnterParameter_name(ctx *Parameter_nameContext) {}

// ExitParameter_name is called when production parameter_name is exited.
func (s *BaseGaussdbParserListener) ExitParameter_name(ctx *Parameter_nameContext) {}

// EnterNumeric_negative is called when production numeric_negative is entered.
func (s *BaseGaussdbParserListener) EnterNumeric_negative(ctx *Numeric_negativeContext) {}

// ExitNumeric_negative is called when production numeric_negative is exited.
func (s *BaseGaussdbParserListener) ExitNumeric_negative(ctx *Numeric_negativeContext) {}

// EnterPlsql_body is called when production plsql_body is entered.
func (s *BaseGaussdbParserListener) EnterPlsql_body(ctx *Plsql_bodyContext) {}

// ExitPlsql_body is called when production plsql_body is exited.
func (s *BaseGaussdbParserListener) ExitPlsql_body(ctx *Plsql_bodyContext) {}

// EnterBody is called when production body is entered.
func (s *BaseGaussdbParserListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseGaussdbParserListener) ExitBody(ctx *BodyContext) {}

// EnterException_handler is called when production exception_handler is entered.
func (s *BaseGaussdbParserListener) EnterException_handler(ctx *Exception_handlerContext) {}

// ExitException_handler is called when production exception_handler is exited.
func (s *BaseGaussdbParserListener) ExitException_handler(ctx *Exception_handlerContext) {}

// EnterSeq_of_statements is called when production seq_of_statements is entered.
func (s *BaseGaussdbParserListener) EnterSeq_of_statements(ctx *Seq_of_statementsContext) {}

// ExitSeq_of_statements is called when production seq_of_statements is exited.
func (s *BaseGaussdbParserListener) ExitSeq_of_statements(ctx *Seq_of_statementsContext) {}

// EnterLabel_declaration is called when production label_declaration is entered.
func (s *BaseGaussdbParserListener) EnterLabel_declaration(ctx *Label_declarationContext) {}

// ExitLabel_declaration is called when production label_declaration is exited.
func (s *BaseGaussdbParserListener) ExitLabel_declaration(ctx *Label_declarationContext) {}

// EnterSeq_of_declare_specs is called when production seq_of_declare_specs is entered.
func (s *BaseGaussdbParserListener) EnterSeq_of_declare_specs(ctx *Seq_of_declare_specsContext) {}

// ExitSeq_of_declare_specs is called when production seq_of_declare_specs is exited.
func (s *BaseGaussdbParserListener) ExitSeq_of_declare_specs(ctx *Seq_of_declare_specsContext) {}

// EnterDeclare_spec is called when production declare_spec is entered.
func (s *BaseGaussdbParserListener) EnterDeclare_spec(ctx *Declare_specContext) {}

// ExitDeclare_spec is called when production declare_spec is exited.
func (s *BaseGaussdbParserListener) ExitDeclare_spec(ctx *Declare_specContext) {}

// EnterPragma_declaration is called when production pragma_declaration is entered.
func (s *BaseGaussdbParserListener) EnterPragma_declaration(ctx *Pragma_declarationContext) {}

// ExitPragma_declaration is called when production pragma_declaration is exited.
func (s *BaseGaussdbParserListener) ExitPragma_declaration(ctx *Pragma_declarationContext) {}

// EnterException_declaration is called when production exception_declaration is entered.
func (s *BaseGaussdbParserListener) EnterException_declaration(ctx *Exception_declarationContext) {}

// ExitException_declaration is called when production exception_declaration is exited.
func (s *BaseGaussdbParserListener) ExitException_declaration(ctx *Exception_declarationContext) {}

// EnterDefault_value_part is called when production default_value_part is entered.
func (s *BaseGaussdbParserListener) EnterDefault_value_part(ctx *Default_value_partContext) {}

// ExitDefault_value_part is called when production default_value_part is exited.
func (s *BaseGaussdbParserListener) ExitDefault_value_part(ctx *Default_value_partContext) {}

// EnterVariable_declaration is called when production variable_declaration is entered.
func (s *BaseGaussdbParserListener) EnterVariable_declaration(ctx *Variable_declarationContext) {}

// ExitVariable_declaration is called when production variable_declaration is exited.
func (s *BaseGaussdbParserListener) ExitVariable_declaration(ctx *Variable_declarationContext) {}

// EnterSubtype_declaration is called when production subtype_declaration is entered.
func (s *BaseGaussdbParserListener) EnterSubtype_declaration(ctx *Subtype_declarationContext) {}

// ExitSubtype_declaration is called when production subtype_declaration is exited.
func (s *BaseGaussdbParserListener) ExitSubtype_declaration(ctx *Subtype_declarationContext) {}

// EnterCursor_declaration is called when production cursor_declaration is entered.
func (s *BaseGaussdbParserListener) EnterCursor_declaration(ctx *Cursor_declarationContext) {}

// ExitCursor_declaration is called when production cursor_declaration is exited.
func (s *BaseGaussdbParserListener) ExitCursor_declaration(ctx *Cursor_declarationContext) {}

// EnterParameter_spec is called when production parameter_spec is entered.
func (s *BaseGaussdbParserListener) EnterParameter_spec(ctx *Parameter_specContext) {}

// ExitParameter_spec is called when production parameter_spec is exited.
func (s *BaseGaussdbParserListener) ExitParameter_spec(ctx *Parameter_specContext) {}

// EnterRef_cursor_type_def is called when production ref_cursor_type_def is entered.
func (s *BaseGaussdbParserListener) EnterRef_cursor_type_def(ctx *Ref_cursor_type_defContext) {}

// ExitRef_cursor_type_def is called when production ref_cursor_type_def is exited.
func (s *BaseGaussdbParserListener) ExitRef_cursor_type_def(ctx *Ref_cursor_type_defContext) {}

// EnterType_declaration is called when production type_declaration is entered.
func (s *BaseGaussdbParserListener) EnterType_declaration(ctx *Type_declarationContext) {}

// ExitType_declaration is called when production type_declaration is exited.
func (s *BaseGaussdbParserListener) ExitType_declaration(ctx *Type_declarationContext) {}

// EnterTable_type_def is called when production table_type_def is entered.
func (s *BaseGaussdbParserListener) EnterTable_type_def(ctx *Table_type_defContext) {}

// ExitTable_type_def is called when production table_type_def is exited.
func (s *BaseGaussdbParserListener) ExitTable_type_def(ctx *Table_type_defContext) {}

// EnterTable_indexed_by_part is called when production table_indexed_by_part is entered.
func (s *BaseGaussdbParserListener) EnterTable_indexed_by_part(ctx *Table_indexed_by_partContext) {}

// ExitTable_indexed_by_part is called when production table_indexed_by_part is exited.
func (s *BaseGaussdbParserListener) ExitTable_indexed_by_part(ctx *Table_indexed_by_partContext) {}

// EnterVarray_type_def is called when production varray_type_def is entered.
func (s *BaseGaussdbParserListener) EnterVarray_type_def(ctx *Varray_type_defContext) {}

// ExitVarray_type_def is called when production varray_type_def is exited.
func (s *BaseGaussdbParserListener) ExitVarray_type_def(ctx *Varray_type_defContext) {}

// EnterRecord_type_def is called when production record_type_def is entered.
func (s *BaseGaussdbParserListener) EnterRecord_type_def(ctx *Record_type_defContext) {}

// ExitRecord_type_def is called when production record_type_def is exited.
func (s *BaseGaussdbParserListener) ExitRecord_type_def(ctx *Record_type_defContext) {}

// EnterField_spec is called when production field_spec is entered.
func (s *BaseGaussdbParserListener) EnterField_spec(ctx *Field_specContext) {}

// ExitField_spec is called when production field_spec is exited.
func (s *BaseGaussdbParserListener) ExitField_spec(ctx *Field_specContext) {}

// EnterProcedure_spec is called when production procedure_spec is entered.
func (s *BaseGaussdbParserListener) EnterProcedure_spec(ctx *Procedure_specContext) {}

// ExitProcedure_spec is called when production procedure_spec is exited.
func (s *BaseGaussdbParserListener) ExitProcedure_spec(ctx *Procedure_specContext) {}

// EnterAccessible_by_clause is called when production accessible_by_clause is entered.
func (s *BaseGaussdbParserListener) EnterAccessible_by_clause(ctx *Accessible_by_clauseContext) {}

// ExitAccessible_by_clause is called when production accessible_by_clause is exited.
func (s *BaseGaussdbParserListener) ExitAccessible_by_clause(ctx *Accessible_by_clauseContext) {}

// EnterAccessor is called when production accessor is entered.
func (s *BaseGaussdbParserListener) EnterAccessor(ctx *AccessorContext) {}

// ExitAccessor is called when production accessor is exited.
func (s *BaseGaussdbParserListener) ExitAccessor(ctx *AccessorContext) {}

// EnterFunction_spec is called when production function_spec is entered.
func (s *BaseGaussdbParserListener) EnterFunction_spec(ctx *Function_specContext) {}

// ExitFunction_spec is called when production function_spec is exited.
func (s *BaseGaussdbParserListener) ExitFunction_spec(ctx *Function_specContext) {}

// EnterParallel_enable_clause is called when production parallel_enable_clause is entered.
func (s *BaseGaussdbParserListener) EnterParallel_enable_clause(ctx *Parallel_enable_clauseContext) {}

// ExitParallel_enable_clause is called when production parallel_enable_clause is exited.
func (s *BaseGaussdbParserListener) ExitParallel_enable_clause(ctx *Parallel_enable_clauseContext) {}

// EnterPartition_by_clause is called when production partition_by_clause is entered.
func (s *BaseGaussdbParserListener) EnterPartition_by_clause(ctx *Partition_by_clauseContext) {}

// ExitPartition_by_clause is called when production partition_by_clause is exited.
func (s *BaseGaussdbParserListener) ExitPartition_by_clause(ctx *Partition_by_clauseContext) {}

// EnterStreaming_clause is called when production streaming_clause is entered.
func (s *BaseGaussdbParserListener) EnterStreaming_clause(ctx *Streaming_clauseContext) {}

// ExitStreaming_clause is called when production streaming_clause is exited.
func (s *BaseGaussdbParserListener) ExitStreaming_clause(ctx *Streaming_clauseContext) {}

// EnterColumn_list is called when production column_list is entered.
func (s *BaseGaussdbParserListener) EnterColumn_list(ctx *Column_listContext) {}

// ExitColumn_list is called when production column_list is exited.
func (s *BaseGaussdbParserListener) ExitColumn_list(ctx *Column_listContext) {}

// EnterParen_column_list is called when production paren_column_list is entered.
func (s *BaseGaussdbParserListener) EnterParen_column_list(ctx *Paren_column_listContext) {}

// ExitParen_column_list is called when production paren_column_list is exited.
func (s *BaseGaussdbParserListener) ExitParen_column_list(ctx *Paren_column_listContext) {}

// EnterProcedure_body is called when production procedure_body is entered.
func (s *BaseGaussdbParserListener) EnterProcedure_body(ctx *Procedure_bodyContext) {}

// ExitProcedure_body is called when production procedure_body is exited.
func (s *BaseGaussdbParserListener) ExitProcedure_body(ctx *Procedure_bodyContext) {}

// EnterFunction_body is called when production function_body is entered.
func (s *BaseGaussdbParserListener) EnterFunction_body(ctx *Function_bodyContext) {}

// ExitFunction_body is called when production function_body is exited.
func (s *BaseGaussdbParserListener) ExitFunction_body(ctx *Function_bodyContext) {}

// EnterResult_cache_clause is called when production result_cache_clause is entered.
func (s *BaseGaussdbParserListener) EnterResult_cache_clause(ctx *Result_cache_clauseContext) {}

// ExitResult_cache_clause is called when production result_cache_clause is exited.
func (s *BaseGaussdbParserListener) ExitResult_cache_clause(ctx *Result_cache_clauseContext) {}

// EnterRelies_on_part is called when production relies_on_part is entered.
func (s *BaseGaussdbParserListener) EnterRelies_on_part(ctx *Relies_on_partContext) {}

// ExitRelies_on_part is called when production relies_on_part is exited.
func (s *BaseGaussdbParserListener) ExitRelies_on_part(ctx *Relies_on_partContext) {}

// EnterInvoker_rights_clause is called when production invoker_rights_clause is entered.
func (s *BaseGaussdbParserListener) EnterInvoker_rights_clause(ctx *Invoker_rights_clauseContext) {}

// ExitInvoker_rights_clause is called when production invoker_rights_clause is exited.
func (s *BaseGaussdbParserListener) ExitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) {}

// EnterPlsql_statement is called when production plsql_statement is entered.
func (s *BaseGaussdbParserListener) EnterPlsql_statement(ctx *Plsql_statementContext) {}

// ExitPlsql_statement is called when production plsql_statement is exited.
func (s *BaseGaussdbParserListener) ExitPlsql_statement(ctx *Plsql_statementContext) {}

// EnterSql_statement is called when production sql_statement is entered.
func (s *BaseGaussdbParserListener) EnterSql_statement(ctx *Sql_statementContext) {}

// ExitSql_statement is called when production sql_statement is exited.
func (s *BaseGaussdbParserListener) ExitSql_statement(ctx *Sql_statementContext) {}

// EnterCursor_manipulation_statements is called when production cursor_manipulation_statements is entered.
func (s *BaseGaussdbParserListener) EnterCursor_manipulation_statements(ctx *Cursor_manipulation_statementsContext) {
}

// ExitCursor_manipulation_statements is called when production cursor_manipulation_statements is exited.
func (s *BaseGaussdbParserListener) ExitCursor_manipulation_statements(ctx *Cursor_manipulation_statementsContext) {
}

// EnterClose_statement is called when production close_statement is entered.
func (s *BaseGaussdbParserListener) EnterClose_statement(ctx *Close_statementContext) {}

// ExitClose_statement is called when production close_statement is exited.
func (s *BaseGaussdbParserListener) ExitClose_statement(ctx *Close_statementContext) {}

// EnterOpen_statement is called when production open_statement is entered.
func (s *BaseGaussdbParserListener) EnterOpen_statement(ctx *Open_statementContext) {}

// ExitOpen_statement is called when production open_statement is exited.
func (s *BaseGaussdbParserListener) ExitOpen_statement(ctx *Open_statementContext) {}

// EnterFetch_statement is called when production fetch_statement is entered.
func (s *BaseGaussdbParserListener) EnterFetch_statement(ctx *Fetch_statementContext) {}

// ExitFetch_statement is called when production fetch_statement is exited.
func (s *BaseGaussdbParserListener) ExitFetch_statement(ctx *Fetch_statementContext) {}

// EnterVariable_or_collection is called when production variable_or_collection is entered.
func (s *BaseGaussdbParserListener) EnterVariable_or_collection(ctx *Variable_or_collectionContext) {}

// ExitVariable_or_collection is called when production variable_or_collection is exited.
func (s *BaseGaussdbParserListener) ExitVariable_or_collection(ctx *Variable_or_collectionContext) {}

// EnterCollection_expression is called when production collection_expression is entered.
func (s *BaseGaussdbParserListener) EnterCollection_expression(ctx *Collection_expressionContext) {}

// ExitCollection_expression is called when production collection_expression is exited.
func (s *BaseGaussdbParserListener) ExitCollection_expression(ctx *Collection_expressionContext) {}

// EnterOpen_for_statement is called when production open_for_statement is entered.
func (s *BaseGaussdbParserListener) EnterOpen_for_statement(ctx *Open_for_statementContext) {}

// ExitOpen_for_statement is called when production open_for_statement is exited.
func (s *BaseGaussdbParserListener) ExitOpen_for_statement(ctx *Open_for_statementContext) {}

// EnterCollection_method_call is called when production collection_method_call is entered.
func (s *BaseGaussdbParserListener) EnterCollection_method_call(ctx *Collection_method_callContext) {}

// ExitCollection_method_call is called when production collection_method_call is exited.
func (s *BaseGaussdbParserListener) ExitCollection_method_call(ctx *Collection_method_callContext) {}

// EnterExecute_immediate is called when production execute_immediate is entered.
func (s *BaseGaussdbParserListener) EnterExecute_immediate(ctx *Execute_immediateContext) {}

// ExitExecute_immediate is called when production execute_immediate is exited.
func (s *BaseGaussdbParserListener) ExitExecute_immediate(ctx *Execute_immediateContext) {}

// EnterDynamic_returning_clause is called when production dynamic_returning_clause is entered.
func (s *BaseGaussdbParserListener) EnterDynamic_returning_clause(ctx *Dynamic_returning_clauseContext) {
}

// ExitDynamic_returning_clause is called when production dynamic_returning_clause is exited.
func (s *BaseGaussdbParserListener) ExitDynamic_returning_clause(ctx *Dynamic_returning_clauseContext) {
}

// EnterRaise_statement is called when production raise_statement is entered.
func (s *BaseGaussdbParserListener) EnterRaise_statement(ctx *Raise_statementContext) {}

// ExitRaise_statement is called when production raise_statement is exited.
func (s *BaseGaussdbParserListener) ExitRaise_statement(ctx *Raise_statementContext) {}

// EnterPipe_row_statement is called when production pipe_row_statement is entered.
func (s *BaseGaussdbParserListener) EnterPipe_row_statement(ctx *Pipe_row_statementContext) {}

// ExitPipe_row_statement is called when production pipe_row_statement is exited.
func (s *BaseGaussdbParserListener) ExitPipe_row_statement(ctx *Pipe_row_statementContext) {}

// EnterForall_statement is called when production forall_statement is entered.
func (s *BaseGaussdbParserListener) EnterForall_statement(ctx *Forall_statementContext) {}

// ExitForall_statement is called when production forall_statement is exited.
func (s *BaseGaussdbParserListener) ExitForall_statement(ctx *Forall_statementContext) {}

// EnterBounds_clause is called when production bounds_clause is entered.
func (s *BaseGaussdbParserListener) EnterBounds_clause(ctx *Bounds_clauseContext) {}

// ExitBounds_clause is called when production bounds_clause is exited.
func (s *BaseGaussdbParserListener) ExitBounds_clause(ctx *Bounds_clauseContext) {}

// EnterBetween_bound is called when production between_bound is entered.
func (s *BaseGaussdbParserListener) EnterBetween_bound(ctx *Between_boundContext) {}

// ExitBetween_bound is called when production between_bound is exited.
func (s *BaseGaussdbParserListener) ExitBetween_bound(ctx *Between_boundContext) {}

// EnterData_manipulation_language_statements is called when production data_manipulation_language_statements is entered.
func (s *BaseGaussdbParserListener) EnterData_manipulation_language_statements(ctx *Data_manipulation_language_statementsContext) {
}

// ExitData_manipulation_language_statements is called when production data_manipulation_language_statements is exited.
func (s *BaseGaussdbParserListener) ExitData_manipulation_language_statements(ctx *Data_manipulation_language_statementsContext) {
}

// EnterNull_statement is called when production null_statement is entered.
func (s *BaseGaussdbParserListener) EnterNull_statement(ctx *Null_statementContext) {}

// ExitNull_statement is called when production null_statement is exited.
func (s *BaseGaussdbParserListener) ExitNull_statement(ctx *Null_statementContext) {}

// EnterGoto_statement is called when production goto_statement is entered.
func (s *BaseGaussdbParserListener) EnterGoto_statement(ctx *Goto_statementContext) {}

// ExitGoto_statement is called when production goto_statement is exited.
func (s *BaseGaussdbParserListener) ExitGoto_statement(ctx *Goto_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseGaussdbParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseGaussdbParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterElsif_part is called when production elsif_part is entered.
func (s *BaseGaussdbParserListener) EnterElsif_part(ctx *Elsif_partContext) {}

// ExitElsif_part is called when production elsif_part is exited.
func (s *BaseGaussdbParserListener) ExitElsif_part(ctx *Elsif_partContext) {}

// EnterElse_part is called when production else_part is entered.
func (s *BaseGaussdbParserListener) EnterElse_part(ctx *Else_partContext) {}

// ExitElse_part is called when production else_part is exited.
func (s *BaseGaussdbParserListener) ExitElse_part(ctx *Else_partContext) {}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BaseGaussdbParserListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BaseGaussdbParserListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterCursor_loop_param is called when production cursor_loop_param is entered.
func (s *BaseGaussdbParserListener) EnterCursor_loop_param(ctx *Cursor_loop_paramContext) {}

// ExitCursor_loop_param is called when production cursor_loop_param is exited.
func (s *BaseGaussdbParserListener) ExitCursor_loop_param(ctx *Cursor_loop_paramContext) {}

// EnterLower_bound is called when production lower_bound is entered.
func (s *BaseGaussdbParserListener) EnterLower_bound(ctx *Lower_boundContext) {}

// ExitLower_bound is called when production lower_bound is exited.
func (s *BaseGaussdbParserListener) ExitLower_bound(ctx *Lower_boundContext) {}

// EnterUpper_bound is called when production upper_bound is entered.
func (s *BaseGaussdbParserListener) EnterUpper_bound(ctx *Upper_boundContext) {}

// ExitUpper_bound is called when production upper_bound is exited.
func (s *BaseGaussdbParserListener) ExitUpper_bound(ctx *Upper_boundContext) {}

// EnterExit_statement is called when production exit_statement is entered.
func (s *BaseGaussdbParserListener) EnterExit_statement(ctx *Exit_statementContext) {}

// ExitExit_statement is called when production exit_statement is exited.
func (s *BaseGaussdbParserListener) ExitExit_statement(ctx *Exit_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseGaussdbParserListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseGaussdbParserListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGaussdbParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGaussdbParserListener) ExitBlock(ctx *BlockContext) {}

// EnterAssignment_statement is called when production assignment_statement is entered.
func (s *BaseGaussdbParserListener) EnterAssignment_statement(ctx *Assignment_statementContext) {}

// ExitAssignment_statement is called when production assignment_statement is exited.
func (s *BaseGaussdbParserListener) ExitAssignment_statement(ctx *Assignment_statementContext) {}

// EnterGeneral_element is called when production general_element is entered.
func (s *BaseGaussdbParserListener) EnterGeneral_element(ctx *General_elementContext) {}

// ExitGeneral_element is called when production general_element is exited.
func (s *BaseGaussdbParserListener) ExitGeneral_element(ctx *General_elementContext) {}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *BaseGaussdbParserListener) EnterContinue_statement(ctx *Continue_statementContext) {}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *BaseGaussdbParserListener) ExitContinue_statement(ctx *Continue_statementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseGaussdbParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseGaussdbParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseGaussdbParserListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseGaussdbParserListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGaussdbParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGaussdbParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCursor_expression is called when production cursor_expression is entered.
func (s *BaseGaussdbParserListener) EnterCursor_expression(ctx *Cursor_expressionContext) {}

// ExitCursor_expression is called when production cursor_expression is exited.
func (s *BaseGaussdbParserListener) ExitCursor_expression(ctx *Cursor_expressionContext) {}

// EnterLogical_expression is called when production logical_expression is entered.
func (s *BaseGaussdbParserListener) EnterLogical_expression(ctx *Logical_expressionContext) {}

// ExitLogical_expression is called when production logical_expression is exited.
func (s *BaseGaussdbParserListener) ExitLogical_expression(ctx *Logical_expressionContext) {}

// EnterUnary_logical_expression is called when production unary_logical_expression is entered.
func (s *BaseGaussdbParserListener) EnterUnary_logical_expression(ctx *Unary_logical_expressionContext) {
}

// ExitUnary_logical_expression is called when production unary_logical_expression is exited.
func (s *BaseGaussdbParserListener) ExitUnary_logical_expression(ctx *Unary_logical_expressionContext) {
}

// EnterUnary_logical_operation is called when production unary_logical_operation is entered.
func (s *BaseGaussdbParserListener) EnterUnary_logical_operation(ctx *Unary_logical_operationContext) {
}

// ExitUnary_logical_operation is called when production unary_logical_operation is exited.
func (s *BaseGaussdbParserListener) ExitUnary_logical_operation(ctx *Unary_logical_operationContext) {
}

// EnterLogical_operation is called when production logical_operation is entered.
func (s *BaseGaussdbParserListener) EnterLogical_operation(ctx *Logical_operationContext) {}

// ExitLogical_operation is called when production logical_operation is exited.
func (s *BaseGaussdbParserListener) ExitLogical_operation(ctx *Logical_operationContext) {}

// EnterMultiset_expression is called when production multiset_expression is entered.
func (s *BaseGaussdbParserListener) EnterMultiset_expression(ctx *Multiset_expressionContext) {}

// ExitMultiset_expression is called when production multiset_expression is exited.
func (s *BaseGaussdbParserListener) ExitMultiset_expression(ctx *Multiset_expressionContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BaseGaussdbParserListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BaseGaussdbParserListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterRelational_expression is called when production relational_expression is entered.
func (s *BaseGaussdbParserListener) EnterRelational_expression(ctx *Relational_expressionContext) {}

// ExitRelational_expression is called when production relational_expression is exited.
func (s *BaseGaussdbParserListener) ExitRelational_expression(ctx *Relational_expressionContext) {}

// EnterCompound_expression is called when production compound_expression is entered.
func (s *BaseGaussdbParserListener) EnterCompound_expression(ctx *Compound_expressionContext) {}

// ExitCompound_expression is called when production compound_expression is exited.
func (s *BaseGaussdbParserListener) ExitCompound_expression(ctx *Compound_expressionContext) {}

// EnterRelational_operator is called when production relational_operator is entered.
func (s *BaseGaussdbParserListener) EnterRelational_operator(ctx *Relational_operatorContext) {}

// ExitRelational_operator is called when production relational_operator is exited.
func (s *BaseGaussdbParserListener) ExitRelational_operator(ctx *Relational_operatorContext) {}

// EnterIn_elements is called when production in_elements is entered.
func (s *BaseGaussdbParserListener) EnterIn_elements(ctx *In_elementsContext) {}

// ExitIn_elements is called when production in_elements is exited.
func (s *BaseGaussdbParserListener) ExitIn_elements(ctx *In_elementsContext) {}

// EnterBetween_elements is called when production between_elements is entered.
func (s *BaseGaussdbParserListener) EnterBetween_elements(ctx *Between_elementsContext) {}

// ExitBetween_elements is called when production between_elements is exited.
func (s *BaseGaussdbParserListener) ExitBetween_elements(ctx *Between_elementsContext) {}

// EnterInterval_expression is called when production interval_expression is entered.
func (s *BaseGaussdbParserListener) EnterInterval_expression(ctx *Interval_expressionContext) {}

// ExitInterval_expression is called when production interval_expression is exited.
func (s *BaseGaussdbParserListener) ExitInterval_expression(ctx *Interval_expressionContext) {}

// EnterModel_expression is called when production model_expression is entered.
func (s *BaseGaussdbParserListener) EnterModel_expression(ctx *Model_expressionContext) {}

// ExitModel_expression is called when production model_expression is exited.
func (s *BaseGaussdbParserListener) ExitModel_expression(ctx *Model_expressionContext) {}

// EnterModel_expression_element is called when production model_expression_element is entered.
func (s *BaseGaussdbParserListener) EnterModel_expression_element(ctx *Model_expression_elementContext) {
}

// ExitModel_expression_element is called when production model_expression_element is exited.
func (s *BaseGaussdbParserListener) ExitModel_expression_element(ctx *Model_expression_elementContext) {
}

// EnterSingle_column_for_loop is called when production single_column_for_loop is entered.
func (s *BaseGaussdbParserListener) EnterSingle_column_for_loop(ctx *Single_column_for_loopContext) {}

// ExitSingle_column_for_loop is called when production single_column_for_loop is exited.
func (s *BaseGaussdbParserListener) ExitSingle_column_for_loop(ctx *Single_column_for_loopContext) {}

// EnterMulti_column_for_loop is called when production multi_column_for_loop is entered.
func (s *BaseGaussdbParserListener) EnterMulti_column_for_loop(ctx *Multi_column_for_loopContext) {}

// ExitMulti_column_for_loop is called when production multi_column_for_loop is exited.
func (s *BaseGaussdbParserListener) ExitMulti_column_for_loop(ctx *Multi_column_for_loopContext) {}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BaseGaussdbParserListener) EnterUnary_expression(ctx *Unary_expressionContext) {}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BaseGaussdbParserListener) ExitUnary_expression(ctx *Unary_expressionContext) {}

// EnterImplicit_cursor_expression is called when production implicit_cursor_expression is entered.
func (s *BaseGaussdbParserListener) EnterImplicit_cursor_expression(ctx *Implicit_cursor_expressionContext) {
}

// ExitImplicit_cursor_expression is called when production implicit_cursor_expression is exited.
func (s *BaseGaussdbParserListener) ExitImplicit_cursor_expression(ctx *Implicit_cursor_expressionContext) {
}

// EnterCase_statement is called when production case_statement is entered.
func (s *BaseGaussdbParserListener) EnterCase_statement(ctx *Case_statementContext) {}

// ExitCase_statement is called when production case_statement is exited.
func (s *BaseGaussdbParserListener) ExitCase_statement(ctx *Case_statementContext) {}

// EnterSimple_case_statement is called when production simple_case_statement is entered.
func (s *BaseGaussdbParserListener) EnterSimple_case_statement(ctx *Simple_case_statementContext) {}

// ExitSimple_case_statement is called when production simple_case_statement is exited.
func (s *BaseGaussdbParserListener) ExitSimple_case_statement(ctx *Simple_case_statementContext) {}

// EnterSimple_case_when_part is called when production simple_case_when_part is entered.
func (s *BaseGaussdbParserListener) EnterSimple_case_when_part(ctx *Simple_case_when_partContext) {}

// ExitSimple_case_when_part is called when production simple_case_when_part is exited.
func (s *BaseGaussdbParserListener) ExitSimple_case_when_part(ctx *Simple_case_when_partContext) {}

// EnterSearched_case_statement is called when production searched_case_statement is entered.
func (s *BaseGaussdbParserListener) EnterSearched_case_statement(ctx *Searched_case_statementContext) {
}

// ExitSearched_case_statement is called when production searched_case_statement is exited.
func (s *BaseGaussdbParserListener) ExitSearched_case_statement(ctx *Searched_case_statementContext) {
}

// EnterSearched_case_when_part is called when production searched_case_when_part is entered.
func (s *BaseGaussdbParserListener) EnterSearched_case_when_part(ctx *Searched_case_when_partContext) {
}

// ExitSearched_case_when_part is called when production searched_case_when_part is exited.
func (s *BaseGaussdbParserListener) ExitSearched_case_when_part(ctx *Searched_case_when_partContext) {
}

// EnterCase_else_part is called when production case_else_part is entered.
func (s *BaseGaussdbParserListener) EnterCase_else_part(ctx *Case_else_partContext) {}

// ExitCase_else_part is called when production case_else_part is exited.
func (s *BaseGaussdbParserListener) ExitCase_else_part(ctx *Case_else_partContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseGaussdbParserListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseGaussdbParserListener) ExitAtom(ctx *AtomContext) {}

// EnterOuter_join_sign is called when production outer_join_sign is entered.
func (s *BaseGaussdbParserListener) EnterOuter_join_sign(ctx *Outer_join_signContext) {}

// ExitOuter_join_sign is called when production outer_join_sign is exited.
func (s *BaseGaussdbParserListener) ExitOuter_join_sign(ctx *Outer_join_signContext) {}

// EnterQuantified_expression is called when production quantified_expression is entered.
func (s *BaseGaussdbParserListener) EnterQuantified_expression(ctx *Quantified_expressionContext) {}

// ExitQuantified_expression is called when production quantified_expression is exited.
func (s *BaseGaussdbParserListener) ExitQuantified_expression(ctx *Quantified_expressionContext) {}

// EnterStandard_function is called when production standard_function is entered.
func (s *BaseGaussdbParserListener) EnterStandard_function(ctx *Standard_functionContext) {}

// ExitStandard_function is called when production standard_function is exited.
func (s *BaseGaussdbParserListener) ExitStandard_function(ctx *Standard_functionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseGaussdbParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseGaussdbParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterTrigger_anonymous_block is called when production trigger_anonymous_block is entered.
func (s *BaseGaussdbParserListener) EnterTrigger_anonymous_block(ctx *Trigger_anonymous_blockContext) {
}

// ExitTrigger_anonymous_block is called when production trigger_anonymous_block is exited.
func (s *BaseGaussdbParserListener) ExitTrigger_anonymous_block(ctx *Trigger_anonymous_blockContext) {
}

// EnterAnonymous_block is called when production anonymous_block is entered.
func (s *BaseGaussdbParserListener) EnterAnonymous_block(ctx *Anonymous_blockContext) {}

// ExitAnonymous_block is called when production anonymous_block is exited.
func (s *BaseGaussdbParserListener) ExitAnonymous_block(ctx *Anonymous_blockContext) {}

// EnterPackage_name is called when production package_name is entered.
func (s *BaseGaussdbParserListener) EnterPackage_name(ctx *Package_nameContext) {}

// ExitPackage_name is called when production package_name is exited.
func (s *BaseGaussdbParserListener) ExitPackage_name(ctx *Package_nameContext) {}

// EnterAlter_package is called when production alter_package is entered.
func (s *BaseGaussdbParserListener) EnterAlter_package(ctx *Alter_packageContext) {}

// ExitAlter_package is called when production alter_package is exited.
func (s *BaseGaussdbParserListener) ExitAlter_package(ctx *Alter_packageContext) {}

// EnterCreate_package is called when production create_package is entered.
func (s *BaseGaussdbParserListener) EnterCreate_package(ctx *Create_packageContext) {}

// ExitCreate_package is called when production create_package is exited.
func (s *BaseGaussdbParserListener) ExitCreate_package(ctx *Create_packageContext) {}

// EnterCreate_package_body is called when production create_package_body is entered.
func (s *BaseGaussdbParserListener) EnterCreate_package_body(ctx *Create_package_bodyContext) {}

// ExitCreate_package_body is called when production create_package_body is exited.
func (s *BaseGaussdbParserListener) ExitCreate_package_body(ctx *Create_package_bodyContext) {}

// EnterPackage_obj_spec is called when production package_obj_spec is entered.
func (s *BaseGaussdbParserListener) EnterPackage_obj_spec(ctx *Package_obj_specContext) {}

// ExitPackage_obj_spec is called when production package_obj_spec is exited.
func (s *BaseGaussdbParserListener) ExitPackage_obj_spec(ctx *Package_obj_specContext) {}

// EnterPackage_obj_body is called when production package_obj_body is entered.
func (s *BaseGaussdbParserListener) EnterPackage_obj_body(ctx *Package_obj_bodyContext) {}

// ExitPackage_obj_body is called when production package_obj_body is exited.
func (s *BaseGaussdbParserListener) ExitPackage_obj_body(ctx *Package_obj_bodyContext) {}

// EnterGaussdb_cleanconnection is called when production gaussdb_cleanconnection is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_cleanconnection(ctx *Gaussdb_cleanconnectionContext) {
}

// ExitGaussdb_cleanconnection is called when production gaussdb_cleanconnection is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_cleanconnection(ctx *Gaussdb_cleanconnectionContext) {
}

// EnterPolicy_name is called when production policy_name is entered.
func (s *BaseGaussdbParserListener) EnterPolicy_name(ctx *Policy_nameContext) {}

// ExitPolicy_name is called when production policy_name is exited.
func (s *BaseGaussdbParserListener) ExitPolicy_name(ctx *Policy_nameContext) {}

// EnterGaussdb_createpolicystmt is called when production gaussdb_createpolicystmt is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_createpolicystmt(ctx *Gaussdb_createpolicystmtContext) {
}

// ExitGaussdb_createpolicystmt is called when production gaussdb_createpolicystmt is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_createpolicystmt(ctx *Gaussdb_createpolicystmtContext) {
}

// EnterGaussdb_audit_clause_list is called when production gaussdb_audit_clause_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_audit_clause_list(ctx *Gaussdb_audit_clause_listContext) {
}

// ExitGaussdb_audit_clause_list is called when production gaussdb_audit_clause_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_audit_clause_list(ctx *Gaussdb_audit_clause_listContext) {
}

// EnterPrivilege_access_audit_clause is called when production privilege_access_audit_clause is entered.
func (s *BaseGaussdbParserListener) EnterPrivilege_access_audit_clause(ctx *Privilege_access_audit_clauseContext) {
}

// ExitPrivilege_access_audit_clause is called when production privilege_access_audit_clause is exited.
func (s *BaseGaussdbParserListener) ExitPrivilege_access_audit_clause(ctx *Privilege_access_audit_clauseContext) {
}

// EnterDdl_clause is called when production ddl_clause is entered.
func (s *BaseGaussdbParserListener) EnterDdl_clause(ctx *Ddl_clauseContext) {}

// ExitDdl_clause is called when production ddl_clause is exited.
func (s *BaseGaussdbParserListener) ExitDdl_clause(ctx *Ddl_clauseContext) {}

// EnterResource_label_name is called when production resource_label_name is entered.
func (s *BaseGaussdbParserListener) EnterResource_label_name(ctx *Resource_label_nameContext) {}

// ExitResource_label_name is called when production resource_label_name is exited.
func (s *BaseGaussdbParserListener) ExitResource_label_name(ctx *Resource_label_nameContext) {}

// EnterDml_clause is called when production dml_clause is entered.
func (s *BaseGaussdbParserListener) EnterDml_clause(ctx *Dml_clauseContext) {}

// ExitDml_clause is called when production dml_clause is exited.
func (s *BaseGaussdbParserListener) ExitDml_clause(ctx *Dml_clauseContext) {}

// EnterFilter_group_clause is called when production filter_group_clause is entered.
func (s *BaseGaussdbParserListener) EnterFilter_group_clause(ctx *Filter_group_clauseContext) {}

// ExitFilter_group_clause is called when production filter_group_clause is exited.
func (s *BaseGaussdbParserListener) ExitFilter_group_clause(ctx *Filter_group_clauseContext) {}

// EnterFilter_type_list is called when production filter_type_list is entered.
func (s *BaseGaussdbParserListener) EnterFilter_type_list(ctx *Filter_type_listContext) {}

// ExitFilter_type_list is called when production filter_type_list is exited.
func (s *BaseGaussdbParserListener) ExitFilter_type_list(ctx *Filter_type_listContext) {}

// EnterFilter_type is called when production filter_type is entered.
func (s *BaseGaussdbParserListener) EnterFilter_type(ctx *Filter_typeContext) {}

// ExitFilter_type is called when production filter_type is exited.
func (s *BaseGaussdbParserListener) ExitFilter_type(ctx *Filter_typeContext) {}

// EnterFilter_value_list is called when production filter_value_list is entered.
func (s *BaseGaussdbParserListener) EnterFilter_value_list(ctx *Filter_value_listContext) {}

// ExitFilter_value_list is called when production filter_value_list is exited.
func (s *BaseGaussdbParserListener) ExitFilter_value_list(ctx *Filter_value_listContext) {}

// EnterFilter_value is called when production filter_value is entered.
func (s *BaseGaussdbParserListener) EnterFilter_value(ctx *Filter_valueContext) {}

// ExitFilter_value is called when production filter_value is exited.
func (s *BaseGaussdbParserListener) ExitFilter_value(ctx *Filter_valueContext) {}

// EnterGaussdb_createbarrierstmt is called when production gaussdb_createbarrierstmt is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_createbarrierstmt(ctx *Gaussdb_createbarrierstmtContext) {
}

// ExitGaussdb_createbarrierstmt is called when production gaussdb_createbarrierstmt is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_createbarrierstmt(ctx *Gaussdb_createbarrierstmtContext) {
}

// EnterGaussdb_create_client_master_key is called when production gaussdb_create_client_master_key is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_client_master_key(ctx *Gaussdb_create_client_master_keyContext) {
}

// ExitGaussdb_create_client_master_key is called when production gaussdb_create_client_master_key is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_client_master_key(ctx *Gaussdb_create_client_master_keyContext) {
}

// EnterClient_master_key_name is called when production client_master_key_name is entered.
func (s *BaseGaussdbParserListener) EnterClient_master_key_name(ctx *Client_master_key_nameContext) {}

// ExitClient_master_key_name is called when production client_master_key_name is exited.
func (s *BaseGaussdbParserListener) ExitClient_master_key_name(ctx *Client_master_key_nameContext) {}

// EnterClient_master_key_list is called when production client_master_key_list is entered.
func (s *BaseGaussdbParserListener) EnterClient_master_key_list(ctx *Client_master_key_listContext) {}

// ExitClient_master_key_list is called when production client_master_key_list is exited.
func (s *BaseGaussdbParserListener) ExitClient_master_key_list(ctx *Client_master_key_listContext) {}

// EnterClient_master_key is called when production client_master_key is entered.
func (s *BaseGaussdbParserListener) EnterClient_master_key(ctx *Client_master_keyContext) {}

// ExitClient_master_key is called when production client_master_key is exited.
func (s *BaseGaussdbParserListener) ExitClient_master_key(ctx *Client_master_keyContext) {}

// EnterGaussdb_create_column_encryption_key is called when production gaussdb_create_column_encryption_key is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_column_encryption_key(ctx *Gaussdb_create_column_encryption_keyContext) {
}

// ExitGaussdb_create_column_encryption_key is called when production gaussdb_create_column_encryption_key is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_column_encryption_key(ctx *Gaussdb_create_column_encryption_keyContext) {
}

// EnterColumn_encryption_key_name is called when production column_encryption_key_name is entered.
func (s *BaseGaussdbParserListener) EnterColumn_encryption_key_name(ctx *Column_encryption_key_nameContext) {
}

// ExitColumn_encryption_key_name is called when production column_encryption_key_name is exited.
func (s *BaseGaussdbParserListener) ExitColumn_encryption_key_name(ctx *Column_encryption_key_nameContext) {
}

// EnterColumn_encryption_key_list is called when production column_encryption_key_list is entered.
func (s *BaseGaussdbParserListener) EnterColumn_encryption_key_list(ctx *Column_encryption_key_listContext) {
}

// ExitColumn_encryption_key_list is called when production column_encryption_key_list is exited.
func (s *BaseGaussdbParserListener) ExitColumn_encryption_key_list(ctx *Column_encryption_key_listContext) {
}

// EnterColumn_encryption_key is called when production column_encryption_key is entered.
func (s *BaseGaussdbParserListener) EnterColumn_encryption_key(ctx *Column_encryption_keyContext) {}

// ExitColumn_encryption_key is called when production column_encryption_key is exited.
func (s *BaseGaussdbParserListener) ExitColumn_encryption_key(ctx *Column_encryption_keyContext) {}

// EnterGaussdb_create_database_link is called when production gaussdb_create_database_link is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_database_link(ctx *Gaussdb_create_database_linkContext) {
}

// ExitGaussdb_create_database_link is called when production gaussdb_create_database_link is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_database_link(ctx *Gaussdb_create_database_linkContext) {
}

// EnterUser_object_name is called when production user_object_name is entered.
func (s *BaseGaussdbParserListener) EnterUser_object_name(ctx *User_object_nameContext) {}

// ExitUser_object_name is called when production user_object_name is exited.
func (s *BaseGaussdbParserListener) ExitUser_object_name(ctx *User_object_nameContext) {}

// EnterPassword_value is called when production password_value is entered.
func (s *BaseGaussdbParserListener) EnterPassword_value(ctx *Password_valueContext) {}

// ExitPassword_value is called when production password_value is exited.
func (s *BaseGaussdbParserListener) ExitPassword_value(ctx *Password_valueContext) {}

// EnterDatabase_link_using is called when production database_link_using is entered.
func (s *BaseGaussdbParserListener) EnterDatabase_link_using(ctx *Database_link_usingContext) {}

// ExitDatabase_link_using is called when production database_link_using is exited.
func (s *BaseGaussdbParserListener) ExitDatabase_link_using(ctx *Database_link_usingContext) {}

// EnterDatabase_link_using_opt is called when production database_link_using_opt is entered.
func (s *BaseGaussdbParserListener) EnterDatabase_link_using_opt(ctx *Database_link_using_optContext) {
}

// ExitDatabase_link_using_opt is called when production database_link_using_opt is exited.
func (s *BaseGaussdbParserListener) ExitDatabase_link_using_opt(ctx *Database_link_using_optContext) {
}

// EnterGaussdb_alter_database_link is called when production gaussdb_alter_database_link is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_alter_database_link(ctx *Gaussdb_alter_database_linkContext) {
}

// ExitGaussdb_alter_database_link is called when production gaussdb_alter_database_link is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_alter_database_link(ctx *Gaussdb_alter_database_linkContext) {
}

// EnterGaussdb_create_directory is called when production gaussdb_create_directory is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_directory(ctx *Gaussdb_create_directoryContext) {
}

// ExitGaussdb_create_directory is called when production gaussdb_create_directory is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_directory(ctx *Gaussdb_create_directoryContext) {
}

// EnterDirectory_name is called when production directory_name is entered.
func (s *BaseGaussdbParserListener) EnterDirectory_name(ctx *Directory_nameContext) {}

// ExitDirectory_name is called when production directory_name is exited.
func (s *BaseGaussdbParserListener) ExitDirectory_name(ctx *Directory_nameContext) {}

// EnterGaussdb_create_masking_policy is called when production gaussdb_create_masking_policy is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_masking_policy(ctx *Gaussdb_create_masking_policyContext) {
}

// ExitGaussdb_create_masking_policy is called when production gaussdb_create_masking_policy is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_masking_policy(ctx *Gaussdb_create_masking_policyContext) {
}

// EnterMasking_clause_list is called when production masking_clause_list is entered.
func (s *BaseGaussdbParserListener) EnterMasking_clause_list(ctx *Masking_clause_listContext) {}

// ExitMasking_clause_list is called when production masking_clause_list is exited.
func (s *BaseGaussdbParserListener) ExitMasking_clause_list(ctx *Masking_clause_listContext) {}

// EnterMasking_clause is called when production masking_clause is entered.
func (s *BaseGaussdbParserListener) EnterMasking_clause(ctx *Masking_clauseContext) {}

// ExitMasking_clause is called when production masking_clause is exited.
func (s *BaseGaussdbParserListener) ExitMasking_clause(ctx *Masking_clauseContext) {}

// EnterMasking_function is called when production masking_function is entered.
func (s *BaseGaussdbParserListener) EnterMasking_function(ctx *Masking_functionContext) {}

// ExitMasking_function is called when production masking_function is exited.
func (s *BaseGaussdbParserListener) ExitMasking_function(ctx *Masking_functionContext) {}

// EnterPolicy_filter_clause is called when production policy_filter_clause is entered.
func (s *BaseGaussdbParserListener) EnterPolicy_filter_clause(ctx *Policy_filter_clauseContext) {}

// ExitPolicy_filter_clause is called when production policy_filter_clause is exited.
func (s *BaseGaussdbParserListener) ExitPolicy_filter_clause(ctx *Policy_filter_clauseContext) {}

// EnterGaussdb_create_node is called when production gaussdb_create_node is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_node(ctx *Gaussdb_create_nodeContext) {}

// ExitGaussdb_create_node is called when production gaussdb_create_node is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_node(ctx *Gaussdb_create_nodeContext) {}

// EnterGaussdb_nodename is called when production gaussdb_nodename is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_nodename(ctx *Gaussdb_nodenameContext) {}

// ExitGaussdb_nodename is called when production gaussdb_nodename is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_nodename(ctx *Gaussdb_nodenameContext) {}

// EnterGaussdb_create_node_group is called when production gaussdb_create_node_group is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_node_group(ctx *Gaussdb_create_node_groupContext) {
}

// ExitGaussdb_create_node_group is called when production gaussdb_create_node_group is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_node_group(ctx *Gaussdb_create_node_groupContext) {
}

// EnterGaussdb_bucketnumber is called when production gaussdb_bucketnumber is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_bucketnumber(ctx *Gaussdb_bucketnumberContext) {}

// ExitGaussdb_bucketnumber is called when production gaussdb_bucketnumber is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_bucketnumber(ctx *Gaussdb_bucketnumberContext) {}

// EnterGaussdb_create_resource_pool is called when production gaussdb_create_resource_pool is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_resource_pool(ctx *Gaussdb_create_resource_poolContext) {
}

// ExitGaussdb_create_resource_pool is called when production gaussdb_create_resource_pool is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_resource_pool(ctx *Gaussdb_create_resource_poolContext) {
}

// EnterGaussdb_resource_pool_name is called when production gaussdb_resource_pool_name is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_resource_pool_name(ctx *Gaussdb_resource_pool_nameContext) {
}

// ExitGaussdb_resource_pool_name is called when production gaussdb_resource_pool_name is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_resource_pool_name(ctx *Gaussdb_resource_pool_nameContext) {
}

// EnterGaussdb_create_resource_label is called when production gaussdb_create_resource_label is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_resource_label(ctx *Gaussdb_create_resource_labelContext) {
}

// ExitGaussdb_create_resource_label is called when production gaussdb_create_resource_label is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_resource_label(ctx *Gaussdb_create_resource_labelContext) {
}

// EnterLabel_item_list is called when production label_item_list is entered.
func (s *BaseGaussdbParserListener) EnterLabel_item_list(ctx *Label_item_listContext) {}

// ExitLabel_item_list is called when production label_item_list is exited.
func (s *BaseGaussdbParserListener) ExitLabel_item_list(ctx *Label_item_listContext) {}

// EnterLabel_item is called when production label_item is entered.
func (s *BaseGaussdbParserListener) EnterLabel_item(ctx *Label_itemContext) {}

// ExitLabel_item is called when production label_item is exited.
func (s *BaseGaussdbParserListener) ExitLabel_item(ctx *Label_itemContext) {}

// EnterResource_path is called when production resource_path is entered.
func (s *BaseGaussdbParserListener) EnterResource_path(ctx *Resource_pathContext) {}

// ExitResource_path is called when production resource_path is exited.
func (s *BaseGaussdbParserListener) ExitResource_path(ctx *Resource_pathContext) {}

// EnterResource_type is called when production resource_type is entered.
func (s *BaseGaussdbParserListener) EnterResource_type(ctx *Resource_typeContext) {}

// ExitResource_type is called when production resource_type is exited.
func (s *BaseGaussdbParserListener) ExitResource_type(ctx *Resource_typeContext) {}

// EnterGaussdb_create_security_label is called when production gaussdb_create_security_label is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_security_label(ctx *Gaussdb_create_security_labelContext) {
}

// ExitGaussdb_create_security_label is called when production gaussdb_create_security_label is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_security_label(ctx *Gaussdb_create_security_labelContext) {
}

// EnterGaussdb_create_synonym is called when production gaussdb_create_synonym is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_synonym(ctx *Gaussdb_create_synonymContext) {}

// ExitGaussdb_create_synonym is called when production gaussdb_create_synonym is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_synonym(ctx *Gaussdb_create_synonymContext) {}

// EnterSynonym_name is called when production synonym_name is entered.
func (s *BaseGaussdbParserListener) EnterSynonym_name(ctx *Synonym_nameContext) {}

// ExitSynonym_name is called when production synonym_name is exited.
func (s *BaseGaussdbParserListener) ExitSynonym_name(ctx *Synonym_nameContext) {}

// EnterGaussdb_partition_less_than_item_list is called when production gaussdb_partition_less_than_item_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_less_than_item_list(ctx *Gaussdb_partition_less_than_item_listContext) {
}

// ExitGaussdb_partition_less_than_item_list is called when production gaussdb_partition_less_than_item_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_less_than_item_list(ctx *Gaussdb_partition_less_than_item_listContext) {
}

// EnterGaussdb_partition_less_than_item is called when production gaussdb_partition_less_than_item is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_less_than_item(ctx *Gaussdb_partition_less_than_itemContext) {
}

// ExitGaussdb_partition_less_than_item is called when production gaussdb_partition_less_than_item is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_less_than_item(ctx *Gaussdb_partition_less_than_itemContext) {
}

// EnterGaussdb_partition_start_end_item_list is called when production gaussdb_partition_start_end_item_list is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_start_end_item_list(ctx *Gaussdb_partition_start_end_item_listContext) {
}

// ExitGaussdb_partition_start_end_item_list is called when production gaussdb_partition_start_end_item_list is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_start_end_item_list(ctx *Gaussdb_partition_start_end_item_listContext) {
}

// EnterGaussdb_partition_start_end_item is called when production gaussdb_partition_start_end_item is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_start_end_item(ctx *Gaussdb_partition_start_end_itemContext) {
}

// ExitGaussdb_partition_start_end_item is called when production gaussdb_partition_start_end_item is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_start_end_item(ctx *Gaussdb_partition_start_end_itemContext) {
}

// EnterOpttablespace_list is called when production opttablespace_list is entered.
func (s *BaseGaussdbParserListener) EnterOpttablespace_list(ctx *Opttablespace_listContext) {}

// ExitOpttablespace_list is called when production opttablespace_list is exited.
func (s *BaseGaussdbParserListener) ExitOpttablespace_list(ctx *Opttablespace_listContext) {}

// EnterGaussdb_partition_hash_item is called when production gaussdb_partition_hash_item is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_hash_item(ctx *Gaussdb_partition_hash_itemContext) {
}

// ExitGaussdb_partition_hash_item is called when production gaussdb_partition_hash_item is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_hash_item(ctx *Gaussdb_partition_hash_itemContext) {
}

// EnterGaussdb_partition_list_item is called when production gaussdb_partition_list_item is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_partition_list_item(ctx *Gaussdb_partition_list_itemContext) {
}

// ExitGaussdb_partition_list_item is called when production gaussdb_partition_list_item is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_partition_list_item(ctx *Gaussdb_partition_list_itemContext) {
}

// EnterGaussdb_create_weak_password_dic is called when production gaussdb_create_weak_password_dic is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_create_weak_password_dic(ctx *Gaussdb_create_weak_password_dicContext) {
}

// ExitGaussdb_create_weak_password_dic is called when production gaussdb_create_weak_password_dic is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_create_weak_password_dic(ctx *Gaussdb_create_weak_password_dicContext) {
}

// EnterGaussdb_purge_stmt is called when production gaussdb_purge_stmt is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_purge_stmt(ctx *Gaussdb_purge_stmtContext) {}

// ExitGaussdb_purge_stmt is called when production gaussdb_purge_stmt is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_purge_stmt(ctx *Gaussdb_purge_stmtContext) {}

// EnterGaussdb_replacestmt is called when production gaussdb_replacestmt is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_replacestmt(ctx *Gaussdb_replacestmtContext) {}

// ExitGaussdb_replacestmt is called when production gaussdb_replacestmt is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_replacestmt(ctx *Gaussdb_replacestmtContext) {}

// EnterGaussdb_shutdown_stmt is called when production gaussdb_shutdown_stmt is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_shutdown_stmt(ctx *Gaussdb_shutdown_stmtContext) {}

// ExitGaussdb_shutdown_stmt is called when production gaussdb_shutdown_stmt is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_shutdown_stmt(ctx *Gaussdb_shutdown_stmtContext) {}

// EnterGaussdb_show_events_stmt is called when production gaussdb_show_events_stmt is entered.
func (s *BaseGaussdbParserListener) EnterGaussdb_show_events_stmt(ctx *Gaussdb_show_events_stmtContext) {
}

// ExitGaussdb_show_events_stmt is called when production gaussdb_show_events_stmt is exited.
func (s *BaseGaussdbParserListener) ExitGaussdb_show_events_stmt(ctx *Gaussdb_show_events_stmtContext) {
}
