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
	return expandBinExpr(tokens)
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
	var children []*parseTreeNode
	var operators []lexToken

	//LHS
	child, err, newTokens := expandExpr2(tokens)
	if child == nil {
		return nil, err, tokens
	}
	tokens = newTokens
	children = append(children, child)
	//now look for (op child)+
	for true {
		//consume operator
		if !expect(tokens, []lexTokenType{lexTokenType_and, lexTokenType_or}) {
			break
		}
		operators = append(operators, tokens[0])
		tokens = tokens[1:]

		//consume child
		child, err, newTokens = expandExpr2(tokens)
		if child == nil {
			return nil, err, tokens
		}
		tokens = newTokens
		children = append(children, child)
	}

	//if theres only one child just return it (effective bypass to Expr2)
	if len(children) == 1 {
		return children[1], "", tokens
	}

	//otherwise construct the tree
	var out *parseTreeNode = new(parseTreeNode)
	out.token = operators[0]
	out.children = append(out.children, children[0])
	out.children = append(out.children, children[1])
	for i := range operators[1:] {
		var holder *parseTreeNode = new(parseTreeNode)
		holder.token = operators[i]
		holder.children = append(holder.children, out)
		holder.children = append(holder.children, children[2+i])
		out = holder
	}


	//return the generated node and consume tokens
	return out, "", tokens
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