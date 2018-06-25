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
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/aws/aws-dax-go/dax/internal/parser/generated"
)

const (
	ProjectionExpr = iota
	ConditionExpr
	KeyConditionExpr
	FilterExpr
	UpdateExpr
)

var errUnexpected = newInvalidParameterError("unexpected error when parsing expression")

// 13.7 Maximizing Parser Speed
// ANTLR v4’s adaptive parsing strategy is more powerful than v3’s, but it comes
// at the cost of a little bit of speed. If you need the most speed and the smallest
// memory footprint possible, you can do a two-step parsing strategy. The first
// step uses a slightly weaker parsing strategy, SLL(*), that almost always works.
// (It’s very similar to v3’s strategy, except it doesn’t need to backtrack.) If the
// first parsing step fails, you have to try the full LL(*) parse. After failing the
// first step, we don’t know whether it’s a true syntax error or whether it’s
// because the SLL(*) strategy wasn’t strong enough. Input that passes the SLL(*)
// step is guaranteed to pass the full LL(*), so there’s no point in trying out that
// more expensive strategy.
//
// Input that fails the second step is truly syntactically invalid.
//
// Antlr use panic/recover as error handling mechanism. Converting it into Go style error.
func walkDynamoDbExpr(typ int, expression string, listener generated.DynamoDbGrammarListener) (err error) {
	if expression == "" {
		return newInvalidParameterError("expression cannot be empty")
	}

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = errUnexpected
			}
		}
	}()

	errList := newErrorListener(typ)
	tree, err := newDynamoDbParseTree(typ, expression, errList)
	if err != nil {
		return err
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return nil
}

func newDynamoDbParseTree(typ int, expression string, listener antlr.ErrorListener) (tree antlr.Tree, err error) {
	is := antlr.NewInputStream(expression)
	lexer := generated.NewDynamoDbGrammarLexer(is)
	lexer.RemoveErrorListeners()

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := generated.NewDynamoDbGrammarParser(tokens)
	parser.BuildParseTrees = true
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(newBailErrorStrategy())
	parser.Interpreter.SetPredictionMode(antlr.PredictionModeSLL)

	defer func() { // Retry with PredictionModeLL
		if r := recover(); r != nil {
			var ok bool
			if _, ok = r.(*antlr.ParseCancellationException); ok {
				tokens.Seek(0)
				parser2 := generated.NewDynamoDbGrammarParser(tokens)
				parser2.BuildParseTrees = true
				parser2.RemoveErrorListeners()
				parser2.AddErrorListener(listener)
				parser2.SetErrorHandler(antlr.NewDefaultErrorStrategy())
				parser2.Interpreter.SetPredictionMode(antlr.PredictionModeLL)
				tree = parseDynamoDbExpr(typ, parser2)
				err = nil
			} else if err, ok = r.(error); !ok {
				err = errUnexpected
			}
		}
	}()
	tree = parseDynamoDbExpr(typ, parser)
	err = nil
	return
}

func parseDynamoDbExpr(typ int, parser *generated.DynamoDbGrammarParser) antlr.Tree {
	switch typ {
	case ProjectionExpr:
		return parser.Projection()
	case ConditionExpr, KeyConditionExpr, FilterExpr:
		return parser.Condition()
	case UpdateExpr:
		return parser.Update()
	default:
		return nil
	}
}

func exprTypeString(typ int) string {
	switch typ {
	case ProjectionExpr:
		return "Projection"
	case ConditionExpr:
		return "Condition"
	case KeyConditionExpr:
		return "KeyCondition"
	case FilterExpr:
		return "Filter"
	case UpdateExpr:
		return "Update"
	default:
		return "Unknown"
	}
}

func newErrorListener(t int) antlr.ErrorListener {
	return &errorListener{typ: t}
}

type errorListener struct {
	*antlr.DefaultErrorListener
	typ int
}

func (el *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	ct, ok := offendingSymbol.(*antlr.CommonToken)
	var tok string
	if ok {
		tok = ct.GetText()
	} else {
		tok = "unknown"
	}
	panic(newInvalidParameterError(fmt.Sprintf("Invalid %s: Syntax error; token: %v, near line %d char %d",
		exprTypeString(el.typ), tok, line, column)))
}

// copied from antlr/error_strategy.go, to workaround bug in antlr.BailErrorStrategy.Recover()
type bailErrorStrategy struct {
	*antlr.DefaultErrorStrategy
}

func newBailErrorStrategy() *bailErrorStrategy {
	b := new(bailErrorStrategy)
	b.DefaultErrorStrategy = antlr.NewDefaultErrorStrategy()
	return b
}

func (b *bailErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	panic(antlr.NewParseCancellationException())
}

func (b *bailErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	b.Recover(recognizer, antlr.NewInputMisMatchException(recognizer))
	return nil
}

func (b *bailErrorStrategy) Sync(recognizer antlr.Parser) {
	// pass
}
