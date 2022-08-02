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
	out, err, _ := expandExpr1(tokens)
	if out == nil {
		os.Stderr.WriteString(err + "\n")
	}
	return out
}

func expandExpr1(tokens []lexToken) (*parseTreeNode, string, []lexToken) {
	out, err, newtokens := expandBinExpr(tokens)
	if out != nil {
		return out, err, newtokens
	}

	return expandExpr2(tokens)
}

func expandExpr2(tokens []lexToken) (*parseTreeNode, string, []lexToken) {
	out, err, newtokens := expandName(tokens);
	if out != nil {
		return out, err, newtokens
	}

	//expand brackets here
	return out, err, newtokens
}

func expandName(tokens []lexToken) (*parseTreeNode, string, []lexToken) {
	if !expect(tokens, []lexTokenType{lexTokenType_name}) {
		return nil, "expected table name but got " + topLiteral(tokens), tokens
	}

	return &parseTreeNode{token: tokens[0]}, "", tokens[1:]
}

func expandBinExpr(tokens []lexToken) (*parseTreeNode, string, []lexToken) {
	var out parseTreeNode
	
	//try for first param
	child, err, newTokens := expandExpr2(tokens)
	if child == nil {
		return nil, err, tokens
	}
	tokens = newTokens
	out.children = append(out.children, child)

	//consume our operator
	if !expect(tokens, []lexTokenType{lexTokenType_and, lexTokenType_or}) {
		return nil, "expected binary operation (and, or) but got " + topLiteral(tokens), tokens
	}
	out.token = tokens[0]
	tokens = tokens[1:]

	//try for secend param
	child, err, newTokens = expandExpr1(tokens)
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

func expect(tokens []lexToken, targets []lexTokenType) bool {
	if len(tokens) > 0 {
		for _, element := range targets {
			if element == tokens[0].tokenType {
				return true
			}
		}
	}

	return false
}

func topLiteral(tokens []lexToken) string {
	if len(tokens) > 0 {
		return tokens[0].literal
	} else {
		return "nothing"
	}
}