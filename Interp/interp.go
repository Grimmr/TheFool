package Interp

import (
	"github.com/Grimmr/TheFool/Parser"
	"github.com/Grimmr/TheFool/Csv"
)

func InterpProgramme(root *Parser.ParseTreeNode, buffers *BufferTable) *Csv.Csv {
	if buffers == nil {
		buffers = NewBufferTable()
	}

	return interpNode(root, buffers)
}

func interpNode(node *Parser.ParseTreeNode, buffers *BufferTable) *Csv.Csv {
	switch node.Token.TokenType {
	case Parser.LexTokenType_name: 
		return buffers.GetOrLoad(node.Token.Literal)
	case Parser.LexTokenType_or:
		lhs := interpNode(node.Children[0], buffers)
		rhs := interpNode(node.Children[1], buffers)
		return lhs.OperatorOr(rhs)
	case Parser.LexTokenType_and:
		lhs := interpNode(node.Children[0], buffers)
		rhs := interpNode(node.Children[1], buffers)
		return lhs.OperatorAnd(rhs)
	case Parser.LexTokenType_less:
		lhs := interpNode(node.Children[0], buffers)
		rhs := interpNode(node.Children[1], buffers)
		return lhs.OperatorLess(rhs)
	case Parser.LexTokenType_lParen:
		return interpNode(node.Children[0], buffers)
	} 

	return nil
}