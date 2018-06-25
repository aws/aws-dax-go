/*
  Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

  Licensed under the Apache License, Version 2.0 (the "License").
  You may not use this file except in compliance with the License.
  A copy of the License is located at

      http://www.apache.org/licenses/LICENSE-2.0

  or in the "license" file accompanying this file. This file is distributed
  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
  express or implied. See the License for the specific language governing
  permissions and limitations under the License.
*/

package parser

import (
	"bytes"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-dax-go/dax/internal/parser/generated"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"io"
	"strconv"
	"strings"
)

const (
	encodingVersion    = 1
	variablePrefix     = ":"
	substitutionPrefix = "#"
)

const tagDocumentPathOrdinal = 3324 // TODO remove once merged

type ExpressionEncoder struct {
	generated.BaseDynamoDbGrammarListener

	// input
	expressions map[int]string
	substitutes map[string]*string
	variables   map[string]*dynamodb.AttributeValue

	// output
	encoded        map[int][]byte
	variableValues []dynamodb.AttributeValue

	// book keeping
	stack             []sexpr
	unusedSubstitutes stringSet
	unusedVariables   stringSet
	variableIdByName  map[string]int
	exprType          int
	nestingLevel      int
	err               error

	// temporary buffer/writer
	cborWriter *cbor.Writer
	buf        *bytes.Buffer
}

func NewExpressionEncoder(expr map[int]string, subs map[string]*string, vars map[string]*dynamodb.AttributeValue) *ExpressionEncoder {
	var b bytes.Buffer
	us := make(stringSet, len(subs))
	us.addKeysStrVal(subs)
	uv := make(stringSet, len(vars))
	uv.addKeysAttrVal(vars)
	return &ExpressionEncoder{
		expressions: expr,
		substitutes: subs,
		variables:   vars,

		encoded:           make(map[int][]byte),
		variableIdByName:  make(map[string]int),
		variableValues:    make([]dynamodb.AttributeValue, 0, len(vars)),
		unusedSubstitutes: us,
		unusedVariables:   uv,

		cborWriter: cbor.NewWriter(&b),
		buf:        &b,
	}
}

func (e *ExpressionEncoder) Parse() (map[int][]byte, error) {
	if len(e.expressions) == 0 || len(e.encoded) == len(e.expressions) {
		return e.encoded, nil
	}
	var err error
	for k, v := range e.expressions {
		e.reset(k)
		if err = walkDynamoDbExpr(k, v, e); err != nil {
			return nil, err
		}
		if err = e.validate(false); err != nil {
			return nil, err
		}
		exprRaw := e.pop()
		expr := e.genSExpr(exprRaw)
		if e.encoded[k], err = e.fullExpr(k, expr); err != nil {
			return nil, err
		}
	}
	if err = e.validate(true); err != nil {
		return nil, err
	}
	return e.encoded, nil
}

func (e *ExpressionEncoder) Write(typ int, writer io.Writer) error {
	if _, err := e.Parse(); err != nil {
		return err
	}
	_, err := writer.Write(e.encoded[typ])
	return err
}

func (e *ExpressionEncoder) reset(typ int) {
	e.exprType = typ
	e.nestingLevel = 0
	e.variableIdByName = make(map[string]int)
	e.variableValues = make([]dynamodb.AttributeValue, 0, len(e.variables))
	e.err = nil
}

func (e *ExpressionEncoder) validate(final bool) error {
	if e.err != nil {
		return e.err
	}
	if final {
		if len(e.unusedSubstitutes) != 0 {
			return newInvalidParameterError(fmt.Sprintf("Value provided in ExpressionAttributeNames unused in expressions: keys: {%v}", e.unusedSubstitutes))
		}
		if len(e.unusedVariables) != 0 {
			return newInvalidParameterError(fmt.Sprintf("Value provided in ExpressionAttributeValues unused in expressions: keys: {%v}", e.unusedVariables))
		}
	} else {
		if sz := len(e.stack); sz != 1 {
			return newInvalidParameterError(fmt.Sprintf("Invalid %sExpression, Stack size = %d", exprTypeString(e.exprType), sz))
		}
		if e.nestingLevel != 0 {
			return newInvalidParameterError(fmt.Sprintf("Invalid %sExpression, Nesting level = %d", exprTypeString(e.exprType), e.nestingLevel))
		}
	}
	return nil
}

func (e *ExpressionEncoder) genSExpr(expr sexpr) []byte {
	e.writeSExpr(expr)
	return e.bytes()
}

func (e *ExpressionEncoder) writeSExpr(expr sexpr) {
	switch {
	case expr.atom != nil:
		e.cborWriter.Write(expr.atom)
	case expr.expr != nil:
		e.cborWriter.WriteArrayHeader(len(expr.expr))
		for _, a := range expr.expr {
			e.writeSExpr(a)
		}
	default:
	}
}

func (e *ExpressionEncoder) fullExpr(typ int, expr []byte) ([]byte, error) {
	if typ == ProjectionExpr {
		e.cborWriter.WriteArrayHeader(2)
	} else {
		e.cborWriter.WriteArrayHeader(3)
	}
	e.cborWriter.WriteInt(encodingVersion)
	e.cborWriter.Write(expr)

	if typ != ProjectionExpr {
		e.cborWriter.WriteArrayHeader(len(e.variableValues))
		for _, v := range e.variableValues {
			if err := cbor.EncodeAttributeValue(&v, e.cborWriter); err != nil {
				return nil, err
			}
		}
	}

	return e.bytes(), nil
}

func (e *ExpressionEncoder) ExitId(ctx *generated.IdContext) {
	if e.err != nil { // bail out
		return
	}
	id := ctx.GetText()
	if id[0:1] == substitutionPrefix {
		s, ok := e.substitutes[id]
		if !ok {
			e.err = newInvalidParameterError(fmt.Sprintf("Invalid %sExpression. Substitution value not provided for %s", exprTypeString(e.exprType), id))
			return
		}
		e.unusedSubstitutes.remove(id)
		e.push(e.encodeDocumentPathElement(*s))
	} else {
		e.push(e.encodeDocumentPathElement(id))
	}
}

func (e *ExpressionEncoder) ExitListAccess(ctx *generated.ListAccessContext) {
	if e.err != nil { // bail out
		return
	}
	s := ctx.GetText()
	o, _ := strconv.ParseInt(s[1:len(s)-1], 10, 64) // parser will detect syntax error
	e.push(e.encodeListAccess(o))
}

func (e *ExpressionEncoder) ExitPath(ctx *generated.PathContext) {
	if e.err != nil { // bail out
		return
	}
	c := ctx.GetChildCount()
	comp := make([]sexpr, c)
	for c > 0 {
		c--
		comp[c] = e.pop()
	}
	b := e.encodeFunction(opDocumentPath, comp)
	e.push(b)
}

func (e *ExpressionEncoder) ExitProjection(ctx *generated.ProjectionContext) {
	if e.err != nil { // bail out
		return
	}
	c := int((ctx.GetChildCount() + 1) / 2) // path, path, ... path
	fields := make([]sexpr, c)
	for c > 0 {
		c--
		fields[c] = e.pop()
	}
	b := e.encodeArray(fields)
	e.push(b)
}

func (e *ExpressionEncoder) ExitLiteralSub(ctx *generated.LiteralSubContext) {
	if e.err != nil { // bail out
		return
	}
	l := ctx.GetText()
	v := e.encodeVariable(l[1:])
	e.push(v)
}

func (e *ExpressionEncoder) ExitComparator_symbol(ctx *generated.Comparator_symbolContext) {
	if e.err != nil { // bail out
		return
	}
	var f sexpr
	switch ctx.GetStart().GetTokenType() {
	case generated.DynamoDbGrammarParserEQ:
		f = e.encodeFunctionCode(opEqual)
	case generated.DynamoDbGrammarParserNE:
		f = e.encodeFunctionCode(opNotEqual)
	case generated.DynamoDbGrammarParserLT:
		f = e.encodeFunctionCode(opLessThan)
	case generated.DynamoDbGrammarParserLE:
		f = e.encodeFunctionCode(opLessEqual)
	case generated.DynamoDbGrammarParserGT:
		f = e.encodeFunctionCode(opGreaterThan)
	case generated.DynamoDbGrammarParserGE:
		f = e.encodeFunctionCode(opGreaterEqual)
	default:
		// parser will detect syntax error
	}
	e.push(f)
}

func (e *ExpressionEncoder) EnterComparator(ctx *generated.ComparatorContext) {
	if e.err != nil { // bail out
		return
	}
	e.nestingLevel++
}

func (e *ExpressionEncoder) ExitComparator(ctx *generated.ComparatorContext) {
	if e.err != nil { // bail out
		return
	}
	e.nestingLevel--
	a2 := e.pop()
	f := e.pop()
	a1 := e.pop()
	e.push(e.encodeArray([]sexpr{f, a1, a2}))
}

func (e *ExpressionEncoder) ExitAnd(ctx *generated.AndContext) {
	if e.err != nil { // bail out
		return
	}
	a2 := e.pop()
	a1 := e.pop()
	e.push(e.encodeFunction(opAnd, []sexpr{a1, a2}))
}

func (e *ExpressionEncoder) ExitOr(ctx *generated.OrContext) {
	if e.err != nil { // bail out
		return
	}
	a2 := e.pop()
	a1 := e.pop()
	e.push(e.encodeFunction(opOr, []sexpr{a1, a2}))
}

func (e *ExpressionEncoder) ExitNegation(ctx *generated.NegationContext) {
	if e.err != nil { // bail out
		return
	}
	a := e.pop()
	e.push(e.encodeFunction(opNot, []sexpr{a}))
}

func (e *ExpressionEncoder) EnterIn(ctx *generated.InContext) {
	if e.err != nil { // bail out
		return
	}
	e.nestingLevel++
}

func (e *ExpressionEncoder) ExitIn(ctx *generated.InContext) {
	if e.err != nil { // bail out
		return
	}
	e.nestingLevel--
	n := (ctx.GetChildCount() - 3) / 2 // arg + IN + '(' + args*2-1 + ')'
	args := make([]sexpr, n)
	for n > 0 {
		n--
		args[n] = e.pop()
	}
	a := e.pop()
	e.push(e.encodeFunction(opIn, []sexpr{a, sexpr{expr: args}}))
}

func (e *ExpressionEncoder) EnterBetween(ctx *generated.BetweenContext) {
	if e.err != nil { // bail out
		return
	}
	e.nestingLevel++
}

func (e *ExpressionEncoder) ExitBetween(ctx *generated.BetweenContext) {
	if e.err != nil { // bail out
		return
	}
	e.nestingLevel--
	a3 := e.pop()
	a2 := e.pop()
	a1 := e.pop()
	e.push(e.encodeFunction(opBetween, []sexpr{a1, a2, a3}))
}

func (e *ExpressionEncoder) EnterFunctionCall(ctx *generated.FunctionCallContext) {
	if e.err != nil { // bail out
		return
	}
	fname := strings.ToLower(ctx.ID().GetText())
	eName := exprTypeString(e.exprType)
	switch e.exprType {
	case UpdateExpr:
		if !e.funcAllowedInUpdate(fname) {
			e.err = newInvalidParameterError(fmt.Sprintf("Invalid %sExpression: The function is not allowed in a %s expression", eName, strings.ToLower(eName)))
			return
		}
		if e.nestingLevel > 0 && fname != "if_not_exists" {
			e.err = newInvalidParameterError(fmt.Sprintf("Only if_not_exists() function can be nested"))
			return
		}
	case FilterExpr, ConditionExpr, KeyConditionExpr:
		if !e.funcAllowedInFilter(fname) {
			e.err = newInvalidParameterError(fmt.Sprintf("Invalid %sExpression: The function is not allowed in a %s expression", eName, strings.ToLower(eName)))
			return
		}
		if e.nestingLevel == 0 && fname == "size" {
			e.err = newInvalidParameterError(fmt.Sprintf("Invalid %sExpression: The function is not allowed in a %s expression", eName, strings.ToLower(eName)))
			return
		}
		if e.nestingLevel > 0 && fname != "size" {
			e.err = newInvalidParameterError(fmt.Sprintf("Only size() function is allowed to be nested"))
		}
	default:
	}
	e.nestingLevel++
}

func (e *ExpressionEncoder) ExitFunctionCall(ctx *generated.FunctionCallContext) {
	if e.err != nil { // bail out
		return
	}

	fname := strings.ToLower(ctx.ID().GetText())
	var fn int
	switch fname {
	case "attribute_exists":
		fn = opAttributeExists
		break
	case "attribute_not_exists":
		fn = opAttributeNotExists
		break
	case "attribute_type":
		fn = opAttributeType
		break
	case "begins_with":
		fn = opBeginsWith
		break
	case "contains":
		fn = opContains
		break
	case "size":
		fn = opSize
		break
	case "if_not_exists":
		fn = opIfNotExists
		break
	case "list_append":
		fn = opListAppend
		break
	default:
		// parser will detect syntax error
	}

	n := (ctx.GetChildCount() - 2) / 2 // children = fname + '(' + numOperands*2-1 + ')'
	args := make([]sexpr, n)
	for n > 0 {
		n--
		args[n] = e.pop()
	}
	e.push(e.encodeFunction(fn, args))
	e.nestingLevel--
}

func (e *ExpressionEncoder) ExitSet_action(ctx *generated.Set_actionContext) {
	if e.err != nil { // bail out
		return
	}
	a := e.pop()
	p := e.pop()
	e.push(e.encodeFunction(opSetAction, []sexpr{p, a}))
}

func (e *ExpressionEncoder) ExitRemove_action(ctx *generated.Remove_actionContext) {
	if e.err != nil { // bail out
		return
	}
	p := e.pop()
	e.push(e.encodeFunction(opRemoveAction, []sexpr{p}))
}

func (e *ExpressionEncoder) ExitAdd_action(ctx *generated.Add_actionContext) {
	if e.err != nil { // bail out
		return
	}
	v := e.pop()
	p := e.pop()
	e.push(e.encodeFunction(opAddAction, []sexpr{p, v}))
}

func (e *ExpressionEncoder) ExitDelete_section(ctx *generated.Delete_sectionContext) {
	if e.err != nil { // bail out
		return
	}
	v := e.pop()
	p := e.pop()
	e.push(e.encodeFunction(opDeleteAction, []sexpr{p, v}))
}

func (e *ExpressionEncoder) EnterPlusMinus(ctx *generated.PlusMinusContext) {
	if e.err != nil { // bail out
		return
	}
	e.nestingLevel++
}

func (e *ExpressionEncoder) ExitPlusMinus(ctx *generated.PlusMinusContext) {
	if e.err != nil { // bail out
		return
	}
	e.nestingLevel--
	a2 := e.pop()
	a1 := e.pop()

	var fn int
	pt := ctx.GetChild(1).(antlr.ParseTree) // leak antlr construct
	switch pt.GetText() {
	case "+":
		fn = opPlus
	case "-":
		fn = opMinus
	default:
		e.err = newInvalidParameterError(fmt.Sprintf("Must be +/-"))
		return
	}
	e.push(e.encodeFunction(fn, []sexpr{a1, a2}))
}

func (e *ExpressionEncoder) ExitUpdate(ctx *generated.UpdateContext) {
	if e.err != nil { // bail out
		return
	}
	sz := len(e.stack)
	us := make([]sexpr, sz)
	for sz > 0 {
		u := e.pop()
		sz--
		us[sz] = u
	}
	e.push(e.encodeArray(us))
}

func (e *ExpressionEncoder) encodeFunction(code int, args []sexpr) sexpr {
	b := make([]sexpr, len(args)+1)

	b[0] = e.encodeFunctionCode(code)
	for i, a := range args {
		b[i+1] = a
	}
	return sexpr{expr: b}
}

func (e *ExpressionEncoder) encodeDocumentPathElement(s string) sexpr {
	e.cborWriter.WriteString(s)
	return sexpr{atom: e.bytes()}
}

func (e *ExpressionEncoder) encodeListAccess(o int64) sexpr {
	e.cborWriter.WriteTag(tagDocumentPathOrdinal)
	e.cborWriter.WriteInt64(o)
	return sexpr{atom: e.bytes()}
}

func (e *ExpressionEncoder) encodeVariable(l string) sexpr {
	n := variablePrefix + l
	v, ok := e.variables[n]
	if !ok {
		e.err = newInvalidParameterError(fmt.Sprintf("Invalid %sExpression: An expression attribute value used in expression is not defined: attribute value %s", exprTypeString(e.exprType), n))
		return sexpr{}
	}
	e.unusedVariables.remove(n)
	id, ok := e.variableIdByName[n]
	if !ok {
		id = len(e.variableValues)
		e.variableIdByName[n] = id
		e.variableValues = append(e.variableValues, *v)
	}
	return e.encodeFunction(opVariable, []sexpr{e.encodeId(id)})
}

func (e *ExpressionEncoder) encodeId(id int) sexpr {
	e.cborWriter.WriteInt(id)
	return sexpr{atom: e.bytes()}
}

func (e *ExpressionEncoder) encodeFunctionCode(c int) sexpr {
	e.cborWriter.WriteInt(c)
	return sexpr{atom: e.bytes()}
}

func (e *ExpressionEncoder) encodeArray(args []sexpr) sexpr {
	b := make([]sexpr, len(args))
	for i, a := range args {
		b[i] = a
	}
	return sexpr{expr: b}
}

func (e *ExpressionEncoder) bytes() []byte {
	e.cborWriter.Flush()
	l := e.buf.Len()
	b := make([]byte, l)
	e.buf.Read(b)
	e.buf.Reset()
	return b
}

func (e *ExpressionEncoder) funcAllowedInUpdate(fname string) bool {
	switch fname {
	case "attribute_exists", "attribute_not_exists", "attribute_type", "begins_with", "contains", "size":
		return false
	default:
		return true
	}
}

func (e *ExpressionEncoder) funcAllowedInFilter(fname string) bool {
	switch fname {
	case "if_not_exists", "list_append":
		return false
	default:
		return true
	}
}

func (e *ExpressionEncoder) push(v sexpr) {
	e.stack = append(e.stack, v)
}

func (e *ExpressionEncoder) pop() sexpr {
	// stack cannot be empty as it will be identified as syntax error by parser
	v := e.stack[len(e.stack)-1]
	e.stack = e.stack[:len(e.stack)-1]
	return v
}

func newInvalidParameterError(msg string) awserr.Error {
	return awserr.New(request.InvalidParameterErrCode, msg, nil)
}

type sexpr struct {
	atom []byte
	expr []sexpr
}

type stringSet map[string]struct{}

func (s stringSet) addKeysStrVal(m map[string]*string) {
	for k, _ := range m {
		s[k] = struct{}{}
	}
}

func (s stringSet) addKeysAttrVal(m map[string]*dynamodb.AttributeValue) {
	for k, _ := range m {
		s[k] = struct{}{}
	}
}

func (s stringSet) remove(v string) {
	delete(s, v)
}

func (s stringSet) String() string {
	out := ""
	first := true
	for k, _ := range s {
		if first {
			first = false
		} else {
			out = out + ", "
		}
		out = out + k
	}
	return out
}

const (
	// Comparison operators
	opEqual = iota
	opNotEqual
	opLessThan
	opGreaterEqual
	opGreaterThan
	opLessEqual

	// Logical operators
	opAnd
	opOr
	opNot

	// Range operators
	opBetween

	// Enumeration operators
	opIn

	// Functions
	opAttributeExists
	opAttributeNotExists
	opAttributeType
	opBeginsWith
	opContains
	opSize

	// Document path elements
	opVariable     // takes 1 argument which is a placeholder for a value
	opDocumentPath // takes an array of string/number which form a document path

	// Update Actions
	opSetAction
	opAddAction
	opDeleteAction
	opRemoveAction

	// Update operations
	opIfNotExists
	opListAppend
	opPlus
	opMinus
)
