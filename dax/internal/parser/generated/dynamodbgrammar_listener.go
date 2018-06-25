// Code generated from DynamoDbGrammar.g4 by ANTLR 4.7.1. DO NOT EDIT.

package generated // DynamoDbGrammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DynamoDbGrammarListener is a complete listener for a parse tree produced by DynamoDbGrammarParser.
type DynamoDbGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProjection_ is called when entering the projection_ production.
	EnterProjection_(c *Projection_Context)

	// EnterProjection is called when entering the projection production.
	EnterProjection(c *ProjectionContext)

	// EnterCondition_ is called when entering the condition_ production.
	EnterCondition_(c *Condition_Context)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterNegation is called when entering the Negation production.
	EnterNegation(c *NegationContext)

	// EnterIn is called when entering the In production.
	EnterIn(c *InContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterBetween is called when entering the Between production.
	EnterBetween(c *BetweenContext)

	// EnterFunctionCondition is called when entering the FunctionCondition production.
	EnterFunctionCondition(c *FunctionConditionContext)

	// EnterComparator is called when entering the Comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterConditionGrouping is called when entering the ConditionGrouping production.
	EnterConditionGrouping(c *ConditionGroupingContext)

	// EnterComparator_symbol is called when entering the comparator_symbol production.
	EnterComparator_symbol(c *Comparator_symbolContext)

	// EnterUpdate_ is called when entering the update_ production.
	EnterUpdate_(c *Update_Context)

	// EnterUpdate is called when entering the update production.
	EnterUpdate(c *UpdateContext)

	// EnterSet_section is called when entering the set_section production.
	EnterSet_section(c *Set_sectionContext)

	// EnterSet_action is called when entering the set_action production.
	EnterSet_action(c *Set_actionContext)

	// EnterAdd_section is called when entering the add_section production.
	EnterAdd_section(c *Add_sectionContext)

	// EnterAdd_action is called when entering the add_action production.
	EnterAdd_action(c *Add_actionContext)

	// EnterDelete_section is called when entering the delete_section production.
	EnterDelete_section(c *Delete_sectionContext)

	// EnterDelete_action is called when entering the delete_action production.
	EnterDelete_action(c *Delete_actionContext)

	// EnterRemove_section is called when entering the remove_section production.
	EnterRemove_section(c *Remove_sectionContext)

	// EnterRemove_action is called when entering the remove_action production.
	EnterRemove_action(c *Remove_actionContext)

	// EnterOperandValue is called when entering the OperandValue production.
	EnterOperandValue(c *OperandValueContext)

	// EnterArithmeticValue is called when entering the ArithmeticValue production.
	EnterArithmeticValue(c *ArithmeticValueContext)

	// EnterPlusMinus is called when entering the PlusMinus production.
	EnterPlusMinus(c *PlusMinusContext)

	// EnterArithmeticParens is called when entering the ArithmeticParens production.
	EnterArithmeticParens(c *ArithmeticParensContext)

	// EnterPathOperand is called when entering the PathOperand production.
	EnterPathOperand(c *PathOperandContext)

	// EnterLiteralOperand is called when entering the LiteralOperand production.
	EnterLiteralOperand(c *LiteralOperandContext)

	// EnterFunctionOperand is called when entering the FunctionOperand production.
	EnterFunctionOperand(c *FunctionOperandContext)

	// EnterParenOperand is called when entering the ParenOperand production.
	EnterParenOperand(c *ParenOperandContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterMapAccess is called when entering the MapAccess production.
	EnterMapAccess(c *MapAccessContext)

	// EnterListAccess is called when entering the ListAccess production.
	EnterListAccess(c *ListAccessContext)

	// EnterLiteralSub is called when entering the LiteralSub production.
	EnterLiteralSub(c *LiteralSubContext)

	// EnterExpression_attr_names_sub is called when entering the expression_attr_names_sub production.
	EnterExpression_attr_names_sub(c *Expression_attr_names_subContext)

	// EnterExpression_attr_values_sub is called when entering the expression_attr_values_sub production.
	EnterExpression_attr_values_sub(c *Expression_attr_values_subContext)

	// EnterStatement_ is called when entering the statement_ production.
	EnterStatement_(c *Statement_Context)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDropTableStatement is called when entering the dropTableStatement production.
	EnterDropTableStatement(c *DropTableStatementContext)

	// EnterDescribeStatement is called when entering the describeStatement production.
	EnterDescribeStatement(c *DescribeStatementContext)

	// EnterAlterTableStatement is called when entering the alterTableStatement production.
	EnterAlterTableStatement(c *AlterTableStatementContext)

	// EnterAlterTableStatementType is called when entering the alterTableStatementType production.
	EnterAlterTableStatementType(c *AlterTableStatementTypeContext)

	// EnterSetCapacity is called when entering the setCapacity production.
	EnterSetCapacity(c *SetCapacityContext)

	// EnterAddIndex is called when entering the addIndex production.
	EnterAddIndex(c *AddIndexContext)

	// EnterDropIndex is called when entering the dropIndex production.
	EnterDropIndex(c *DropIndexContext)

	// EnterAlterIndex is called when entering the alterIndex production.
	EnterAlterIndex(c *AlterIndexContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterCreateTableStatement is called when entering the createTableStatement production.
	EnterCreateTableStatement(c *CreateTableStatementContext)

	// EnterShowTablesStatement is called when entering the showTablesStatement production.
	EnterShowTablesStatement(c *ShowTablesStatementContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterSelectProjection is called when entering the selectProjection production.
	EnterSelectProjection(c *SelectProjectionContext)

	// EnterOptionBlock is called when entering the optionBlock production.
	EnterOptionBlock(c *OptionBlockContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterSingleOption is called when entering the singleOption production.
	EnterSingleOption(c *SingleOptionContext)

	// EnterKeyValueOption is called when entering the keyValueOption production.
	EnterKeyValueOption(c *KeyValueOptionContext)

	// EnterOptionKey is called when entering the optionKey production.
	EnterOptionKey(c *OptionKeyContext)

	// EnterOptionValue is called when entering the optionValue production.
	EnterOptionValue(c *OptionValueContext)

	// EnterOptionValueString is called when entering the optionValueString production.
	EnterOptionValueString(c *OptionValueStringContext)

	// EnterOptionValueNumber is called when entering the optionValueNumber production.
	EnterOptionValueNumber(c *OptionValueNumberContext)

	// EnterStar is called when entering the star production.
	EnterStar(c *StarContext)

	// EnterHint is called when entering the hint production.
	EnterHint(c *HintContext)

	// EnterIndexHint is called when entering the indexHint production.
	EnterIndexHint(c *IndexHintContext)

	// EnterIndexHintName is called when entering the indexHintName production.
	EnterIndexHintName(c *IndexHintNameContext)

	// EnterScanInfo is called when entering the scanInfo production.
	EnterScanInfo(c *ScanInfoContext)

	// EnterTotalSegment is called when entering the totalSegment production.
	EnterTotalSegment(c *TotalSegmentContext)

	// EnterSegment is called when entering the segment production.
	EnterSegment(c *SegmentContext)

	// EnterWhere is called when entering the where production.
	EnterWhere(c *WhereContext)

	// EnterPrimaryKeyDecl is called when entering the primaryKeyDecl production.
	EnterPrimaryKeyDecl(c *PrimaryKeyDeclContext)

	// EnterSecondaryIndexDecl is called when entering the secondaryIndexDecl production.
	EnterSecondaryIndexDecl(c *SecondaryIndexDeclContext)

	// EnterSecondaryIndexType is called when entering the secondaryIndexType production.
	EnterSecondaryIndexType(c *SecondaryIndexTypeContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterProjectionIndex is called when entering the projectionIndex production.
	EnterProjectionIndex(c *ProjectionIndexContext)

	// EnterProjectionIndexKeysOnly is called when entering the projectionIndexKeysOnly production.
	EnterProjectionIndexKeysOnly(c *ProjectionIndexKeysOnlyContext)

	// EnterProjectionIndexVector is called when entering the projectionIndexVector production.
	EnterProjectionIndexVector(c *ProjectionIndexVectorContext)

	// EnterCapacity is called when entering the capacity production.
	EnterCapacity(c *CapacityContext)

	// EnterReadUnits is called when entering the readUnits production.
	EnterReadUnits(c *ReadUnitsContext)

	// EnterWriteUnits is called when entering the writeUnits production.
	EnterWriteUnits(c *WriteUnitsContext)

	// EnterIndexDecl is called when entering the indexDecl production.
	EnterIndexDecl(c *IndexDeclContext)

	// EnterAttributeDecl is called when entering the attributeDecl production.
	EnterAttributeDecl(c *AttributeDeclContext)

	// EnterHashKey is called when entering the hashKey production.
	EnterHashKey(c *HashKeyContext)

	// EnterRangeKey is called when entering the rangeKey production.
	EnterRangeKey(c *RangeKeyContext)

	// EnterAttributeName is called when entering the attributeName production.
	EnterAttributeName(c *AttributeNameContext)

	// EnterAttributeType is called when entering the attributeType production.
	EnterAttributeType(c *AttributeTypeContext)

	// EnterReturning is called when entering the returning production.
	EnterReturning(c *ReturningContext)

	// EnterReturningValue is called when entering the returningValue production.
	EnterReturningValue(c *ReturningValueContext)

	// EnterOnDuplicateKeyUpdate is called when entering the onDuplicateKeyUpdate production.
	EnterOnDuplicateKeyUpdate(c *OnDuplicateKeyUpdateContext)

	// EnterIfClause is called when entering the ifClause production.
	EnterIfClause(c *IfClauseContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterDdlName is called when entering the ddlName production.
	EnterDdlName(c *DdlNameContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterUnknown is called when entering the unknown production.
	EnterUnknown(c *UnknownContext)

	// ExitProjection_ is called when exiting the projection_ production.
	ExitProjection_(c *Projection_Context)

	// ExitProjection is called when exiting the projection production.
	ExitProjection(c *ProjectionContext)

	// ExitCondition_ is called when exiting the condition_ production.
	ExitCondition_(c *Condition_Context)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitNegation is called when exiting the Negation production.
	ExitNegation(c *NegationContext)

	// ExitIn is called when exiting the In production.
	ExitIn(c *InContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitBetween is called when exiting the Between production.
	ExitBetween(c *BetweenContext)

	// ExitFunctionCondition is called when exiting the FunctionCondition production.
	ExitFunctionCondition(c *FunctionConditionContext)

	// ExitComparator is called when exiting the Comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitConditionGrouping is called when exiting the ConditionGrouping production.
	ExitConditionGrouping(c *ConditionGroupingContext)

	// ExitComparator_symbol is called when exiting the comparator_symbol production.
	ExitComparator_symbol(c *Comparator_symbolContext)

	// ExitUpdate_ is called when exiting the update_ production.
	ExitUpdate_(c *Update_Context)

	// ExitUpdate is called when exiting the update production.
	ExitUpdate(c *UpdateContext)

	// ExitSet_section is called when exiting the set_section production.
	ExitSet_section(c *Set_sectionContext)

	// ExitSet_action is called when exiting the set_action production.
	ExitSet_action(c *Set_actionContext)

	// ExitAdd_section is called when exiting the add_section production.
	ExitAdd_section(c *Add_sectionContext)

	// ExitAdd_action is called when exiting the add_action production.
	ExitAdd_action(c *Add_actionContext)

	// ExitDelete_section is called when exiting the delete_section production.
	ExitDelete_section(c *Delete_sectionContext)

	// ExitDelete_action is called when exiting the delete_action production.
	ExitDelete_action(c *Delete_actionContext)

	// ExitRemove_section is called when exiting the remove_section production.
	ExitRemove_section(c *Remove_sectionContext)

	// ExitRemove_action is called when exiting the remove_action production.
	ExitRemove_action(c *Remove_actionContext)

	// ExitOperandValue is called when exiting the OperandValue production.
	ExitOperandValue(c *OperandValueContext)

	// ExitArithmeticValue is called when exiting the ArithmeticValue production.
	ExitArithmeticValue(c *ArithmeticValueContext)

	// ExitPlusMinus is called when exiting the PlusMinus production.
	ExitPlusMinus(c *PlusMinusContext)

	// ExitArithmeticParens is called when exiting the ArithmeticParens production.
	ExitArithmeticParens(c *ArithmeticParensContext)

	// ExitPathOperand is called when exiting the PathOperand production.
	ExitPathOperand(c *PathOperandContext)

	// ExitLiteralOperand is called when exiting the LiteralOperand production.
	ExitLiteralOperand(c *LiteralOperandContext)

	// ExitFunctionOperand is called when exiting the FunctionOperand production.
	ExitFunctionOperand(c *FunctionOperandContext)

	// ExitParenOperand is called when exiting the ParenOperand production.
	ExitParenOperand(c *ParenOperandContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitMapAccess is called when exiting the MapAccess production.
	ExitMapAccess(c *MapAccessContext)

	// ExitListAccess is called when exiting the ListAccess production.
	ExitListAccess(c *ListAccessContext)

	// ExitLiteralSub is called when exiting the LiteralSub production.
	ExitLiteralSub(c *LiteralSubContext)

	// ExitExpression_attr_names_sub is called when exiting the expression_attr_names_sub production.
	ExitExpression_attr_names_sub(c *Expression_attr_names_subContext)

	// ExitExpression_attr_values_sub is called when exiting the expression_attr_values_sub production.
	ExitExpression_attr_values_sub(c *Expression_attr_values_subContext)

	// ExitStatement_ is called when exiting the statement_ production.
	ExitStatement_(c *Statement_Context)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDropTableStatement is called when exiting the dropTableStatement production.
	ExitDropTableStatement(c *DropTableStatementContext)

	// ExitDescribeStatement is called when exiting the describeStatement production.
	ExitDescribeStatement(c *DescribeStatementContext)

	// ExitAlterTableStatement is called when exiting the alterTableStatement production.
	ExitAlterTableStatement(c *AlterTableStatementContext)

	// ExitAlterTableStatementType is called when exiting the alterTableStatementType production.
	ExitAlterTableStatementType(c *AlterTableStatementTypeContext)

	// ExitSetCapacity is called when exiting the setCapacity production.
	ExitSetCapacity(c *SetCapacityContext)

	// ExitAddIndex is called when exiting the addIndex production.
	ExitAddIndex(c *AddIndexContext)

	// ExitDropIndex is called when exiting the dropIndex production.
	ExitDropIndex(c *DropIndexContext)

	// ExitAlterIndex is called when exiting the alterIndex production.
	ExitAlterIndex(c *AlterIndexContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitCreateTableStatement is called when exiting the createTableStatement production.
	ExitCreateTableStatement(c *CreateTableStatementContext)

	// ExitShowTablesStatement is called when exiting the showTablesStatement production.
	ExitShowTablesStatement(c *ShowTablesStatementContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitSelectProjection is called when exiting the selectProjection production.
	ExitSelectProjection(c *SelectProjectionContext)

	// ExitOptionBlock is called when exiting the optionBlock production.
	ExitOptionBlock(c *OptionBlockContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitSingleOption is called when exiting the singleOption production.
	ExitSingleOption(c *SingleOptionContext)

	// ExitKeyValueOption is called when exiting the keyValueOption production.
	ExitKeyValueOption(c *KeyValueOptionContext)

	// ExitOptionKey is called when exiting the optionKey production.
	ExitOptionKey(c *OptionKeyContext)

	// ExitOptionValue is called when exiting the optionValue production.
	ExitOptionValue(c *OptionValueContext)

	// ExitOptionValueString is called when exiting the optionValueString production.
	ExitOptionValueString(c *OptionValueStringContext)

	// ExitOptionValueNumber is called when exiting the optionValueNumber production.
	ExitOptionValueNumber(c *OptionValueNumberContext)

	// ExitStar is called when exiting the star production.
	ExitStar(c *StarContext)

	// ExitHint is called when exiting the hint production.
	ExitHint(c *HintContext)

	// ExitIndexHint is called when exiting the indexHint production.
	ExitIndexHint(c *IndexHintContext)

	// ExitIndexHintName is called when exiting the indexHintName production.
	ExitIndexHintName(c *IndexHintNameContext)

	// ExitScanInfo is called when exiting the scanInfo production.
	ExitScanInfo(c *ScanInfoContext)

	// ExitTotalSegment is called when exiting the totalSegment production.
	ExitTotalSegment(c *TotalSegmentContext)

	// ExitSegment is called when exiting the segment production.
	ExitSegment(c *SegmentContext)

	// ExitWhere is called when exiting the where production.
	ExitWhere(c *WhereContext)

	// ExitPrimaryKeyDecl is called when exiting the primaryKeyDecl production.
	ExitPrimaryKeyDecl(c *PrimaryKeyDeclContext)

	// ExitSecondaryIndexDecl is called when exiting the secondaryIndexDecl production.
	ExitSecondaryIndexDecl(c *SecondaryIndexDeclContext)

	// ExitSecondaryIndexType is called when exiting the secondaryIndexType production.
	ExitSecondaryIndexType(c *SecondaryIndexTypeContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitProjectionIndex is called when exiting the projectionIndex production.
	ExitProjectionIndex(c *ProjectionIndexContext)

	// ExitProjectionIndexKeysOnly is called when exiting the projectionIndexKeysOnly production.
	ExitProjectionIndexKeysOnly(c *ProjectionIndexKeysOnlyContext)

	// ExitProjectionIndexVector is called when exiting the projectionIndexVector production.
	ExitProjectionIndexVector(c *ProjectionIndexVectorContext)

	// ExitCapacity is called when exiting the capacity production.
	ExitCapacity(c *CapacityContext)

	// ExitReadUnits is called when exiting the readUnits production.
	ExitReadUnits(c *ReadUnitsContext)

	// ExitWriteUnits is called when exiting the writeUnits production.
	ExitWriteUnits(c *WriteUnitsContext)

	// ExitIndexDecl is called when exiting the indexDecl production.
	ExitIndexDecl(c *IndexDeclContext)

	// ExitAttributeDecl is called when exiting the attributeDecl production.
	ExitAttributeDecl(c *AttributeDeclContext)

	// ExitHashKey is called when exiting the hashKey production.
	ExitHashKey(c *HashKeyContext)

	// ExitRangeKey is called when exiting the rangeKey production.
	ExitRangeKey(c *RangeKeyContext)

	// ExitAttributeName is called when exiting the attributeName production.
	ExitAttributeName(c *AttributeNameContext)

	// ExitAttributeType is called when exiting the attributeType production.
	ExitAttributeType(c *AttributeTypeContext)

	// ExitReturning is called when exiting the returning production.
	ExitReturning(c *ReturningContext)

	// ExitReturningValue is called when exiting the returningValue production.
	ExitReturningValue(c *ReturningValueContext)

	// ExitOnDuplicateKeyUpdate is called when exiting the onDuplicateKeyUpdate production.
	ExitOnDuplicateKeyUpdate(c *OnDuplicateKeyUpdateContext)

	// ExitIfClause is called when exiting the ifClause production.
	ExitIfClause(c *IfClauseContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitDdlName is called when exiting the ddlName production.
	ExitDdlName(c *DdlNameContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitUnknown is called when exiting the unknown production.
	ExitUnknown(c *UnknownContext)
}
