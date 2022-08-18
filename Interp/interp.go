package Interp

import (
	"github.com/Grimmr/TheFool/Parser"
	"github.com/Grimmr/TheFool/Csv"
	"math/rand"
	"time"
)

func InterpProgramme(root *Parser.ParseTreeNode, buffers *BufferTable, seed *rand.Source) *Csv.Csv {
	if seed == nil {
		v := rand.NewSource(time.Now().UnixNano())
		seed = &v
	}
	
	if buffers == nil {
		buffers = NewBufferTable()
	}

	return interpNode(root, buffers, rand.New(*seed))
}

func interpNode(node *Parser.ParseTreeNode, buffers *BufferTable, random *rand.Rand) *Csv.Csv {
	switch node.Token.TokenType {
	case Parser.LexTokenType_name: 
		return buffers.GetOrLoad(node.Token.Literal)
	case Parser.LexTokenType_or:
		lhs := interpNode(node.Children[0], buffers, random)
		rhs := interpNode(node.Children[1], buffers, random)
		return lhs.OperatorOr(rhs)
	case Parser.LexTokenType_and:
		lhs := interpNode(node.Children[0], buffers, random)
		rhs := interpNode(node.Children[1], buffers, random)
		return lhs.OperatorAnd(rhs)
	case Parser.LexTokenType_less:
		lhs := interpNode(node.Children[0], buffers, random)
		rhs := interpNode(node.Children[1], buffers, random)
		return lhs.OperatorLess(rhs)
	case Parser.LexTokenType_lParen:
		return interpNode(node.Children[0], buffers, random)
	case Parser.LexTokenType_percent:
		lhs := interpNode(node.Children[0], buffers, random)
		rhs := node.Children[1].Token.Literal
		return lhs.OperatorRandomSubset(rhs, random)
	case Parser.LexTokenType_plus:
		lhs := interpNode(node.Children[0], buffers, random)
		rhs := interpNode(node.Children[1], buffers, random)
		return lhs.OperatorPlus(rhs)
	case Parser.LexTokenType_minus:
		lhs := interpNode(node.Children[0], buffers, random)
		rhs := interpNode(node.Children[1], buffers, random)
		return lhs.OperatorMinus(rhs)
	} 

	return nil
}