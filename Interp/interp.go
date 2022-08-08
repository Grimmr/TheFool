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
	} 

	return nil
}