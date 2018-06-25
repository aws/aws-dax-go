// Code generated from DynamoDbGrammar.g4 by ANTLR 4.7.1. DO NOT EDIT.

package generated // DynamoDbGrammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDynamoDbGrammarListener is a complete listener for a parse tree produced by DynamoDbGrammarParser.
type BaseDynamoDbGrammarListener struct{}

var _ DynamoDbGrammarListener = &BaseDynamoDbGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDynamoDbGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDynamoDbGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDynamoDbGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDynamoDbGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProjection_ is called when production projection_ is entered.
func (s *BaseDynamoDbGrammarListener) EnterProjection_(ctx *Projection_Context) {}

// ExitProjection_ is called when production projection_ is exited.
func (s *BaseDynamoDbGrammarListener) ExitProjection_(ctx *Projection_Context) {}

// EnterProjection is called when production projection is entered.
func (s *BaseDynamoDbGrammarListener) EnterProjection(ctx *ProjectionContext) {}

// ExitProjection is called when production projection is exited.
func (s *BaseDynamoDbGrammarListener) ExitProjection(ctx *ProjectionContext) {}

// EnterCondition_ is called when production condition_ is entered.
func (s *BaseDynamoDbGrammarListener) EnterCondition_(ctx *Condition_Context) {}

// ExitCondition_ is called when production condition_ is exited.
func (s *BaseDynamoDbGrammarListener) ExitCondition_(ctx *Condition_Context) {}

// EnterOr is called when production Or is entered.
func (s *BaseDynamoDbGrammarListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BaseDynamoDbGrammarListener) ExitOr(ctx *OrContext) {}

// EnterNegation is called when production Negation is entered.
func (s *BaseDynamoDbGrammarListener) EnterNegation(ctx *NegationContext) {}

// ExitNegation is called when production Negation is exited.
func (s *BaseDynamoDbGrammarListener) ExitNegation(ctx *NegationContext) {}

// EnterIn is called when production In is entered.
func (s *BaseDynamoDbGrammarListener) EnterIn(ctx *InContext) {}

// ExitIn is called when production In is exited.
func (s *BaseDynamoDbGrammarListener) ExitIn(ctx *InContext) {}

// EnterAnd is called when production And is entered.
func (s *BaseDynamoDbGrammarListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BaseDynamoDbGrammarListener) ExitAnd(ctx *AndContext) {}

// EnterBetween is called when production Between is entered.
func (s *BaseDynamoDbGrammarListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production Between is exited.
func (s *BaseDynamoDbGrammarListener) ExitBetween(ctx *BetweenContext) {}

// EnterFunctionCondition is called when production FunctionCondition is entered.
func (s *BaseDynamoDbGrammarListener) EnterFunctionCondition(ctx *FunctionConditionContext) {}

// ExitFunctionCondition is called when production FunctionCondition is exited.
func (s *BaseDynamoDbGrammarListener) ExitFunctionCondition(ctx *FunctionConditionContext) {}

// EnterComparator is called when production Comparator is entered.
func (s *BaseDynamoDbGrammarListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production Comparator is exited.
func (s *BaseDynamoDbGrammarListener) ExitComparator(ctx *ComparatorContext) {}

// EnterConditionGrouping is called when production ConditionGrouping is entered.
func (s *BaseDynamoDbGrammarListener) EnterConditionGrouping(ctx *ConditionGroupingContext) {}

// ExitConditionGrouping is called when production ConditionGrouping is exited.
func (s *BaseDynamoDbGrammarListener) ExitConditionGrouping(ctx *ConditionGroupingContext) {}

// EnterComparator_symbol is called when production comparator_symbol is entered.
func (s *BaseDynamoDbGrammarListener) EnterComparator_symbol(ctx *Comparator_symbolContext) {}

// ExitComparator_symbol is called when production comparator_symbol is exited.
func (s *BaseDynamoDbGrammarListener) ExitComparator_symbol(ctx *Comparator_symbolContext) {}

// EnterUpdate_ is called when production update_ is entered.
func (s *BaseDynamoDbGrammarListener) EnterUpdate_(ctx *Update_Context) {}

// ExitUpdate_ is called when production update_ is exited.
func (s *BaseDynamoDbGrammarListener) ExitUpdate_(ctx *Update_Context) {}

// EnterUpdate is called when production update is entered.
func (s *BaseDynamoDbGrammarListener) EnterUpdate(ctx *UpdateContext) {}

// ExitUpdate is called when production update is exited.
func (s *BaseDynamoDbGrammarListener) ExitUpdate(ctx *UpdateContext) {}

// EnterSet_section is called when production set_section is entered.
func (s *BaseDynamoDbGrammarListener) EnterSet_section(ctx *Set_sectionContext) {}

// ExitSet_section is called when production set_section is exited.
func (s *BaseDynamoDbGrammarListener) ExitSet_section(ctx *Set_sectionContext) {}

// EnterSet_action is called when production set_action is entered.
func (s *BaseDynamoDbGrammarListener) EnterSet_action(ctx *Set_actionContext) {}

// ExitSet_action is called when production set_action is exited.
func (s *BaseDynamoDbGrammarListener) ExitSet_action(ctx *Set_actionContext) {}

// EnterAdd_section is called when production add_section is entered.
func (s *BaseDynamoDbGrammarListener) EnterAdd_section(ctx *Add_sectionContext) {}

// ExitAdd_section is called when production add_section is exited.
func (s *BaseDynamoDbGrammarListener) ExitAdd_section(ctx *Add_sectionContext) {}

// EnterAdd_action is called when production add_action is entered.
func (s *BaseDynamoDbGrammarListener) EnterAdd_action(ctx *Add_actionContext) {}

// ExitAdd_action is called when production add_action is exited.
func (s *BaseDynamoDbGrammarListener) ExitAdd_action(ctx *Add_actionContext) {}

// EnterDelete_section is called when production delete_section is entered.
func (s *BaseDynamoDbGrammarListener) EnterDelete_section(ctx *Delete_sectionContext) {}

// ExitDelete_section is called when production delete_section is exited.
func (s *BaseDynamoDbGrammarListener) ExitDelete_section(ctx *Delete_sectionContext) {}

// EnterDelete_action is called when production delete_action is entered.
func (s *BaseDynamoDbGrammarListener) EnterDelete_action(ctx *Delete_actionContext) {}

// ExitDelete_action is called when production delete_action is exited.
func (s *BaseDynamoDbGrammarListener) ExitDelete_action(ctx *Delete_actionContext) {}

// EnterRemove_section is called when production remove_section is entered.
func (s *BaseDynamoDbGrammarListener) EnterRemove_section(ctx *Remove_sectionContext) {}

// ExitRemove_section is called when production remove_section is exited.
func (s *BaseDynamoDbGrammarListener) ExitRemove_section(ctx *Remove_sectionContext) {}

// EnterRemove_action is called when production remove_action is entered.
func (s *BaseDynamoDbGrammarListener) EnterRemove_action(ctx *Remove_actionContext) {}

// ExitRemove_action is called when production remove_action is exited.
func (s *BaseDynamoDbGrammarListener) ExitRemove_action(ctx *Remove_actionContext) {}

// EnterOperandValue is called when production OperandValue is entered.
func (s *BaseDynamoDbGrammarListener) EnterOperandValue(ctx *OperandValueContext) {}

// ExitOperandValue is called when production OperandValue is exited.
func (s *BaseDynamoDbGrammarListener) ExitOperandValue(ctx *OperandValueContext) {}

// EnterArithmeticValue is called when production ArithmeticValue is entered.
func (s *BaseDynamoDbGrammarListener) EnterArithmeticValue(ctx *ArithmeticValueContext) {}

// ExitArithmeticValue is called when production ArithmeticValue is exited.
func (s *BaseDynamoDbGrammarListener) ExitArithmeticValue(ctx *ArithmeticValueContext) {}

// EnterPlusMinus is called when production PlusMinus is entered.
func (s *BaseDynamoDbGrammarListener) EnterPlusMinus(ctx *PlusMinusContext) {}

// ExitPlusMinus is called when production PlusMinus is exited.
func (s *BaseDynamoDbGrammarListener) ExitPlusMinus(ctx *PlusMinusContext) {}

// EnterArithmeticParens is called when production ArithmeticParens is entered.
func (s *BaseDynamoDbGrammarListener) EnterArithmeticParens(ctx *ArithmeticParensContext) {}

// ExitArithmeticParens is called when production ArithmeticParens is exited.
func (s *BaseDynamoDbGrammarListener) ExitArithmeticParens(ctx *ArithmeticParensContext) {}

// EnterPathOperand is called when production PathOperand is entered.
func (s *BaseDynamoDbGrammarListener) EnterPathOperand(ctx *PathOperandContext) {}

// ExitPathOperand is called when production PathOperand is exited.
func (s *BaseDynamoDbGrammarListener) ExitPathOperand(ctx *PathOperandContext) {}

// EnterLiteralOperand is called when production LiteralOperand is entered.
func (s *BaseDynamoDbGrammarListener) EnterLiteralOperand(ctx *LiteralOperandContext) {}

// ExitLiteralOperand is called when production LiteralOperand is exited.
func (s *BaseDynamoDbGrammarListener) ExitLiteralOperand(ctx *LiteralOperandContext) {}

// EnterFunctionOperand is called when production FunctionOperand is entered.
func (s *BaseDynamoDbGrammarListener) EnterFunctionOperand(ctx *FunctionOperandContext) {}

// ExitFunctionOperand is called when production FunctionOperand is exited.
func (s *BaseDynamoDbGrammarListener) ExitFunctionOperand(ctx *FunctionOperandContext) {}

// EnterParenOperand is called when production ParenOperand is entered.
func (s *BaseDynamoDbGrammarListener) EnterParenOperand(ctx *ParenOperandContext) {}

// ExitParenOperand is called when production ParenOperand is exited.
func (s *BaseDynamoDbGrammarListener) ExitParenOperand(ctx *ParenOperandContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseDynamoDbGrammarListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseDynamoDbGrammarListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterPath is called when production path is entered.
func (s *BaseDynamoDbGrammarListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseDynamoDbGrammarListener) ExitPath(ctx *PathContext) {}

// EnterId is called when production id is entered.
func (s *BaseDynamoDbGrammarListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseDynamoDbGrammarListener) ExitId(ctx *IdContext) {}

// EnterMapAccess is called when production MapAccess is entered.
func (s *BaseDynamoDbGrammarListener) EnterMapAccess(ctx *MapAccessContext) {}

// ExitMapAccess is called when production MapAccess is exited.
func (s *BaseDynamoDbGrammarListener) ExitMapAccess(ctx *MapAccessContext) {}

// EnterListAccess is called when production ListAccess is entered.
func (s *BaseDynamoDbGrammarListener) EnterListAccess(ctx *ListAccessContext) {}

// ExitListAccess is called when production ListAccess is exited.
func (s *BaseDynamoDbGrammarListener) ExitListAccess(ctx *ListAccessContext) {}

// EnterLiteralSub is called when production LiteralSub is entered.
func (s *BaseDynamoDbGrammarListener) EnterLiteralSub(ctx *LiteralSubContext) {}

// ExitLiteralSub is called when production LiteralSub is exited.
func (s *BaseDynamoDbGrammarListener) ExitLiteralSub(ctx *LiteralSubContext) {}

// EnterExpression_attr_names_sub is called when production expression_attr_names_sub is entered.
func (s *BaseDynamoDbGrammarListener) EnterExpression_attr_names_sub(ctx *Expression_attr_names_subContext) {
}

// ExitExpression_attr_names_sub is called when production expression_attr_names_sub is exited.
func (s *BaseDynamoDbGrammarListener) ExitExpression_attr_names_sub(ctx *Expression_attr_names_subContext) {
}

// EnterExpression_attr_values_sub is called when production expression_attr_values_sub is entered.
func (s *BaseDynamoDbGrammarListener) EnterExpression_attr_values_sub(ctx *Expression_attr_values_subContext) {
}

// ExitExpression_attr_values_sub is called when production expression_attr_values_sub is exited.
func (s *BaseDynamoDbGrammarListener) ExitExpression_attr_values_sub(ctx *Expression_attr_values_subContext) {
}

// EnterStatement_ is called when production statement_ is entered.
func (s *BaseDynamoDbGrammarListener) EnterStatement_(ctx *Statement_Context) {}

// ExitStatement_ is called when production statement_ is exited.
func (s *BaseDynamoDbGrammarListener) ExitStatement_(ctx *Statement_Context) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDynamoDbGrammarListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDynamoDbGrammarListener) ExitStatement(ctx *StatementContext) {}

// EnterDropTableStatement is called when production dropTableStatement is entered.
func (s *BaseDynamoDbGrammarListener) EnterDropTableStatement(ctx *DropTableStatementContext) {}

// ExitDropTableStatement is called when production dropTableStatement is exited.
func (s *BaseDynamoDbGrammarListener) ExitDropTableStatement(ctx *DropTableStatementContext) {}

// EnterDescribeStatement is called when production describeStatement is entered.
func (s *BaseDynamoDbGrammarListener) EnterDescribeStatement(ctx *DescribeStatementContext) {}

// ExitDescribeStatement is called when production describeStatement is exited.
func (s *BaseDynamoDbGrammarListener) ExitDescribeStatement(ctx *DescribeStatementContext) {}

// EnterAlterTableStatement is called when production alterTableStatement is entered.
func (s *BaseDynamoDbGrammarListener) EnterAlterTableStatement(ctx *AlterTableStatementContext) {}

// ExitAlterTableStatement is called when production alterTableStatement is exited.
func (s *BaseDynamoDbGrammarListener) ExitAlterTableStatement(ctx *AlterTableStatementContext) {}

// EnterAlterTableStatementType is called when production alterTableStatementType is entered.
func (s *BaseDynamoDbGrammarListener) EnterAlterTableStatementType(ctx *AlterTableStatementTypeContext) {
}

// ExitAlterTableStatementType is called when production alterTableStatementType is exited.
func (s *BaseDynamoDbGrammarListener) ExitAlterTableStatementType(ctx *AlterTableStatementTypeContext) {
}

// EnterSetCapacity is called when production setCapacity is entered.
func (s *BaseDynamoDbGrammarListener) EnterSetCapacity(ctx *SetCapacityContext) {}

// ExitSetCapacity is called when production setCapacity is exited.
func (s *BaseDynamoDbGrammarListener) ExitSetCapacity(ctx *SetCapacityContext) {}

// EnterAddIndex is called when production addIndex is entered.
func (s *BaseDynamoDbGrammarListener) EnterAddIndex(ctx *AddIndexContext) {}

// ExitAddIndex is called when production addIndex is exited.
func (s *BaseDynamoDbGrammarListener) ExitAddIndex(ctx *AddIndexContext) {}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseDynamoDbGrammarListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseDynamoDbGrammarListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterAlterIndex is called when production alterIndex is entered.
func (s *BaseDynamoDbGrammarListener) EnterAlterIndex(ctx *AlterIndexContext) {}

// ExitAlterIndex is called when production alterIndex is exited.
func (s *BaseDynamoDbGrammarListener) ExitAlterIndex(ctx *AlterIndexContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseDynamoDbGrammarListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseDynamoDbGrammarListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseDynamoDbGrammarListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseDynamoDbGrammarListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseDynamoDbGrammarListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseDynamoDbGrammarListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseDynamoDbGrammarListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseDynamoDbGrammarListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterShowTablesStatement is called when production showTablesStatement is entered.
func (s *BaseDynamoDbGrammarListener) EnterShowTablesStatement(ctx *ShowTablesStatementContext) {}

// ExitShowTablesStatement is called when production showTablesStatement is exited.
func (s *BaseDynamoDbGrammarListener) ExitShowTablesStatement(ctx *ShowTablesStatementContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseDynamoDbGrammarListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseDynamoDbGrammarListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterSelectProjection is called when production selectProjection is entered.
func (s *BaseDynamoDbGrammarListener) EnterSelectProjection(ctx *SelectProjectionContext) {}

// ExitSelectProjection is called when production selectProjection is exited.
func (s *BaseDynamoDbGrammarListener) ExitSelectProjection(ctx *SelectProjectionContext) {}

// EnterOptionBlock is called when production optionBlock is entered.
func (s *BaseDynamoDbGrammarListener) EnterOptionBlock(ctx *OptionBlockContext) {}

// ExitOptionBlock is called when production optionBlock is exited.
func (s *BaseDynamoDbGrammarListener) ExitOptionBlock(ctx *OptionBlockContext) {}

// EnterOption is called when production option is entered.
func (s *BaseDynamoDbGrammarListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseDynamoDbGrammarListener) ExitOption(ctx *OptionContext) {}

// EnterSingleOption is called when production singleOption is entered.
func (s *BaseDynamoDbGrammarListener) EnterSingleOption(ctx *SingleOptionContext) {}

// ExitSingleOption is called when production singleOption is exited.
func (s *BaseDynamoDbGrammarListener) ExitSingleOption(ctx *SingleOptionContext) {}

// EnterKeyValueOption is called when production keyValueOption is entered.
func (s *BaseDynamoDbGrammarListener) EnterKeyValueOption(ctx *KeyValueOptionContext) {}

// ExitKeyValueOption is called when production keyValueOption is exited.
func (s *BaseDynamoDbGrammarListener) ExitKeyValueOption(ctx *KeyValueOptionContext) {}

// EnterOptionKey is called when production optionKey is entered.
func (s *BaseDynamoDbGrammarListener) EnterOptionKey(ctx *OptionKeyContext) {}

// ExitOptionKey is called when production optionKey is exited.
func (s *BaseDynamoDbGrammarListener) ExitOptionKey(ctx *OptionKeyContext) {}

// EnterOptionValue is called when production optionValue is entered.
func (s *BaseDynamoDbGrammarListener) EnterOptionValue(ctx *OptionValueContext) {}

// ExitOptionValue is called when production optionValue is exited.
func (s *BaseDynamoDbGrammarListener) ExitOptionValue(ctx *OptionValueContext) {}

// EnterOptionValueString is called when production optionValueString is entered.
func (s *BaseDynamoDbGrammarListener) EnterOptionValueString(ctx *OptionValueStringContext) {}

// ExitOptionValueString is called when production optionValueString is exited.
func (s *BaseDynamoDbGrammarListener) ExitOptionValueString(ctx *OptionValueStringContext) {}

// EnterOptionValueNumber is called when production optionValueNumber is entered.
func (s *BaseDynamoDbGrammarListener) EnterOptionValueNumber(ctx *OptionValueNumberContext) {}

// ExitOptionValueNumber is called when production optionValueNumber is exited.
func (s *BaseDynamoDbGrammarListener) ExitOptionValueNumber(ctx *OptionValueNumberContext) {}

// EnterStar is called when production star is entered.
func (s *BaseDynamoDbGrammarListener) EnterStar(ctx *StarContext) {}

// ExitStar is called when production star is exited.
func (s *BaseDynamoDbGrammarListener) ExitStar(ctx *StarContext) {}

// EnterHint is called when production hint is entered.
func (s *BaseDynamoDbGrammarListener) EnterHint(ctx *HintContext) {}

// ExitHint is called when production hint is exited.
func (s *BaseDynamoDbGrammarListener) ExitHint(ctx *HintContext) {}

// EnterIndexHint is called when production indexHint is entered.
func (s *BaseDynamoDbGrammarListener) EnterIndexHint(ctx *IndexHintContext) {}

// ExitIndexHint is called when production indexHint is exited.
func (s *BaseDynamoDbGrammarListener) ExitIndexHint(ctx *IndexHintContext) {}

// EnterIndexHintName is called when production indexHintName is entered.
func (s *BaseDynamoDbGrammarListener) EnterIndexHintName(ctx *IndexHintNameContext) {}

// ExitIndexHintName is called when production indexHintName is exited.
func (s *BaseDynamoDbGrammarListener) ExitIndexHintName(ctx *IndexHintNameContext) {}

// EnterScanInfo is called when production scanInfo is entered.
func (s *BaseDynamoDbGrammarListener) EnterScanInfo(ctx *ScanInfoContext) {}

// ExitScanInfo is called when production scanInfo is exited.
func (s *BaseDynamoDbGrammarListener) ExitScanInfo(ctx *ScanInfoContext) {}

// EnterTotalSegment is called when production totalSegment is entered.
func (s *BaseDynamoDbGrammarListener) EnterTotalSegment(ctx *TotalSegmentContext) {}

// ExitTotalSegment is called when production totalSegment is exited.
func (s *BaseDynamoDbGrammarListener) ExitTotalSegment(ctx *TotalSegmentContext) {}

// EnterSegment is called when production segment is entered.
func (s *BaseDynamoDbGrammarListener) EnterSegment(ctx *SegmentContext) {}

// ExitSegment is called when production segment is exited.
func (s *BaseDynamoDbGrammarListener) ExitSegment(ctx *SegmentContext) {}

// EnterWhere is called when production where is entered.
func (s *BaseDynamoDbGrammarListener) EnterWhere(ctx *WhereContext) {}

// ExitWhere is called when production where is exited.
func (s *BaseDynamoDbGrammarListener) ExitWhere(ctx *WhereContext) {}

// EnterPrimaryKeyDecl is called when production primaryKeyDecl is entered.
func (s *BaseDynamoDbGrammarListener) EnterPrimaryKeyDecl(ctx *PrimaryKeyDeclContext) {}

// ExitPrimaryKeyDecl is called when production primaryKeyDecl is exited.
func (s *BaseDynamoDbGrammarListener) ExitPrimaryKeyDecl(ctx *PrimaryKeyDeclContext) {}

// EnterSecondaryIndexDecl is called when production secondaryIndexDecl is entered.
func (s *BaseDynamoDbGrammarListener) EnterSecondaryIndexDecl(ctx *SecondaryIndexDeclContext) {}

// ExitSecondaryIndexDecl is called when production secondaryIndexDecl is exited.
func (s *BaseDynamoDbGrammarListener) ExitSecondaryIndexDecl(ctx *SecondaryIndexDeclContext) {}

// EnterSecondaryIndexType is called when production secondaryIndexType is entered.
func (s *BaseDynamoDbGrammarListener) EnterSecondaryIndexType(ctx *SecondaryIndexTypeContext) {}

// ExitSecondaryIndexType is called when production secondaryIndexType is exited.
func (s *BaseDynamoDbGrammarListener) ExitSecondaryIndexType(ctx *SecondaryIndexTypeContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseDynamoDbGrammarListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseDynamoDbGrammarListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterProjectionIndex is called when production projectionIndex is entered.
func (s *BaseDynamoDbGrammarListener) EnterProjectionIndex(ctx *ProjectionIndexContext) {}

// ExitProjectionIndex is called when production projectionIndex is exited.
func (s *BaseDynamoDbGrammarListener) ExitProjectionIndex(ctx *ProjectionIndexContext) {}

// EnterProjectionIndexKeysOnly is called when production projectionIndexKeysOnly is entered.
func (s *BaseDynamoDbGrammarListener) EnterProjectionIndexKeysOnly(ctx *ProjectionIndexKeysOnlyContext) {
}

// ExitProjectionIndexKeysOnly is called when production projectionIndexKeysOnly is exited.
func (s *BaseDynamoDbGrammarListener) ExitProjectionIndexKeysOnly(ctx *ProjectionIndexKeysOnlyContext) {
}

// EnterProjectionIndexVector is called when production projectionIndexVector is entered.
func (s *BaseDynamoDbGrammarListener) EnterProjectionIndexVector(ctx *ProjectionIndexVectorContext) {}

// ExitProjectionIndexVector is called when production projectionIndexVector is exited.
func (s *BaseDynamoDbGrammarListener) ExitProjectionIndexVector(ctx *ProjectionIndexVectorContext) {}

// EnterCapacity is called when production capacity is entered.
func (s *BaseDynamoDbGrammarListener) EnterCapacity(ctx *CapacityContext) {}

// ExitCapacity is called when production capacity is exited.
func (s *BaseDynamoDbGrammarListener) ExitCapacity(ctx *CapacityContext) {}

// EnterReadUnits is called when production readUnits is entered.
func (s *BaseDynamoDbGrammarListener) EnterReadUnits(ctx *ReadUnitsContext) {}

// ExitReadUnits is called when production readUnits is exited.
func (s *BaseDynamoDbGrammarListener) ExitReadUnits(ctx *ReadUnitsContext) {}

// EnterWriteUnits is called when production writeUnits is entered.
func (s *BaseDynamoDbGrammarListener) EnterWriteUnits(ctx *WriteUnitsContext) {}

// ExitWriteUnits is called when production writeUnits is exited.
func (s *BaseDynamoDbGrammarListener) ExitWriteUnits(ctx *WriteUnitsContext) {}

// EnterIndexDecl is called when production indexDecl is entered.
func (s *BaseDynamoDbGrammarListener) EnterIndexDecl(ctx *IndexDeclContext) {}

// ExitIndexDecl is called when production indexDecl is exited.
func (s *BaseDynamoDbGrammarListener) ExitIndexDecl(ctx *IndexDeclContext) {}

// EnterAttributeDecl is called when production attributeDecl is entered.
func (s *BaseDynamoDbGrammarListener) EnterAttributeDecl(ctx *AttributeDeclContext) {}

// ExitAttributeDecl is called when production attributeDecl is exited.
func (s *BaseDynamoDbGrammarListener) ExitAttributeDecl(ctx *AttributeDeclContext) {}

// EnterHashKey is called when production hashKey is entered.
func (s *BaseDynamoDbGrammarListener) EnterHashKey(ctx *HashKeyContext) {}

// ExitHashKey is called when production hashKey is exited.
func (s *BaseDynamoDbGrammarListener) ExitHashKey(ctx *HashKeyContext) {}

// EnterRangeKey is called when production rangeKey is entered.
func (s *BaseDynamoDbGrammarListener) EnterRangeKey(ctx *RangeKeyContext) {}

// ExitRangeKey is called when production rangeKey is exited.
func (s *BaseDynamoDbGrammarListener) ExitRangeKey(ctx *RangeKeyContext) {}

// EnterAttributeName is called when production attributeName is entered.
func (s *BaseDynamoDbGrammarListener) EnterAttributeName(ctx *AttributeNameContext) {}

// ExitAttributeName is called when production attributeName is exited.
func (s *BaseDynamoDbGrammarListener) ExitAttributeName(ctx *AttributeNameContext) {}

// EnterAttributeType is called when production attributeType is entered.
func (s *BaseDynamoDbGrammarListener) EnterAttributeType(ctx *AttributeTypeContext) {}

// ExitAttributeType is called when production attributeType is exited.
func (s *BaseDynamoDbGrammarListener) ExitAttributeType(ctx *AttributeTypeContext) {}

// EnterReturning is called when production returning is entered.
func (s *BaseDynamoDbGrammarListener) EnterReturning(ctx *ReturningContext) {}

// ExitReturning is called when production returning is exited.
func (s *BaseDynamoDbGrammarListener) ExitReturning(ctx *ReturningContext) {}

// EnterReturningValue is called when production returningValue is entered.
func (s *BaseDynamoDbGrammarListener) EnterReturningValue(ctx *ReturningValueContext) {}

// ExitReturningValue is called when production returningValue is exited.
func (s *BaseDynamoDbGrammarListener) ExitReturningValue(ctx *ReturningValueContext) {}

// EnterOnDuplicateKeyUpdate is called when production onDuplicateKeyUpdate is entered.
func (s *BaseDynamoDbGrammarListener) EnterOnDuplicateKeyUpdate(ctx *OnDuplicateKeyUpdateContext) {}

// ExitOnDuplicateKeyUpdate is called when production onDuplicateKeyUpdate is exited.
func (s *BaseDynamoDbGrammarListener) ExitOnDuplicateKeyUpdate(ctx *OnDuplicateKeyUpdateContext) {}

// EnterIfClause is called when production ifClause is entered.
func (s *BaseDynamoDbGrammarListener) EnterIfClause(ctx *IfClauseContext) {}

// ExitIfClause is called when production ifClause is exited.
func (s *BaseDynamoDbGrammarListener) ExitIfClause(ctx *IfClauseContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseDynamoDbGrammarListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseDynamoDbGrammarListener) ExitTableName(ctx *TableNameContext) {}

// EnterDdlName is called when production ddlName is entered.
func (s *BaseDynamoDbGrammarListener) EnterDdlName(ctx *DdlNameContext) {}

// ExitDdlName is called when production ddlName is exited.
func (s *BaseDynamoDbGrammarListener) ExitDdlName(ctx *DdlNameContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseDynamoDbGrammarListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseDynamoDbGrammarListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterUnknown is called when production unknown is entered.
func (s *BaseDynamoDbGrammarListener) EnterUnknown(ctx *UnknownContext) {}

// ExitUnknown is called when production unknown is exited.
func (s *BaseDynamoDbGrammarListener) ExitUnknown(ctx *UnknownContext) {}
