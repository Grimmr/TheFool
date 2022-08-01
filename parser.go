package main

import (
	"os"
	"fmt"
)

type parseTreeNode struct {
	token lexToken
	children []*parseTreeNode
}

func parseProgramme(tokens []lexToken) *parseTreeNode {
	fmt.Println(tokens)
	out, err, _ := expandExpr(tokens)
	if out == nil {
		os.Stderr.WriteString(err)
	}
	return out
}

func expandExpr(tokens []lexToken) (*parseTreeNode, string, []lexToken) {
	out, _, newTokens := expandBinExpr(tokens)
	if out != nil {
		return out, "", newTokens
	}

	out, _, newTokens = expandName(tokens)
	if out != nil {
		return out, "", newTokens
	}

	return nil, "expeced binary expresion got " + tokens[0].literal + " " + tokens[1].literal + " " + tokens[2].literal, tokens
}

func expandName(tokens []lexToken) (*parseTreeNode, string, []lexToken) {
	if tokens[0].tokenType != lexTokenType_name {
		return nil, "expected table name but got " + tokens[0].literal, tokens
	}

	return &parseTreeNode{token: tokens[0]}, "", tokens[1:]
}

func expandBinExpr(tokens []lexToken) (*parseTreeNode, string, []lexToken) {
	var out parseTreeNode
	
	//try for bin operator at look ahead 1
	if len(tokens) <= 2 {
		return nil, "expected binary operation (and, or) but got " + tokens[0].literal, tokens
	}
	if tokens[1].tokenType != lexTokenType_and && tokens[1].tokenType != lexTokenType_or {
		return nil, "expected binary operation (and, or) but got " + tokens[1].literal, tokens
	}
	out.token = tokens[1]

	//try for first param
	child, err, newTokens := expandExpr(tokens)
	if child == nil {
		return nil, err, tokens
	} else {
		tokens = newTokens
	}

	out.children = append(out.children, child)

	//consume our operator
	tokens = tokens[1:]

	//try for secend param
	child, err, newTokens = expandExpr(tokens)
	if child == nil {
		return nil, err, tokens
	} else {
		tokens = newTokens
	}

	out.children = append(out.children, child)

	//return the generated node and consume tokens
	return &out, "", tokens
}

func walkParseTree(node *parseTreeNode) string {
	out := "(" + node.token.literal 
	for _, element := range node.children {
		out += " " + walkParseTree(element)
	}
	out += ")"
	return out;
}