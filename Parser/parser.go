package Parser

import (
	"errors"
)

type ParseTreeNode struct {
	Token lexToken
	Children []*ParseTreeNode
}

func ParseProgramme(tokens []lexToken) *ParseTreeNode {
	out, err, _ := expandExpr1(tokens)
	if out == nil {
		panic(errors.New(err + "\n"))
	}
	return out
}

func expandExpr1(tokens []lexToken) (*ParseTreeNode, string, []lexToken) {
	return expandBinExpr(tokens)
}

func expandExpr2(tokens []lexToken) (*ParseTreeNode, string, []lexToken) {
	out, err, newtokens := expandName(tokens);
	if out != nil {
		return out, err, newtokens
	}

	//expand brackets here
	return out, err, newtokens
}

func expandName(tokens []lexToken) (*ParseTreeNode, string, []lexToken) {
	if !expect(tokens, []lexTokenType{LexTokenType_name}) {
		return nil, "expected table name but got " + topLiteral(tokens), tokens
	}

	return &ParseTreeNode{Token: tokens[0]}, "", tokens[1:]
}

func expandBinExpr(tokens []lexToken) (*ParseTreeNode, string, []lexToken) {
	var children []*ParseTreeNode
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
		if !expect(tokens, []lexTokenType{LexTokenType_and, LexTokenType_or}) {
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
		return children[0], "", tokens
	}

	//otherwise construct the tree
	var out *ParseTreeNode = new(ParseTreeNode)
	out.Token = operators[0]
	out.Children = append(out.Children, children[0])
	out.Children = append(out.Children, children[1])
	for i := range operators[1:] {
		var holder *ParseTreeNode = new(ParseTreeNode)
		holder.Token = operators[i+1]
		holder.Children = append(holder.Children, out)
		holder.Children = append(holder.Children, children[2+i])
		out = holder
	}


	//return the generated node and consume tokens
	return out, "", tokens
}

func WalkParseTree(node *ParseTreeNode) string {
	out := "(" + node.Token.Literal 
	for _, element := range node.Children {
		out += " " + WalkParseTree(element)
	}
	out += ")"
	return out;
}

func expect(tokens []lexToken, targets []lexTokenType) bool {
	if len(tokens) > 0 {
		for _, element := range targets {
			if element == tokens[0].TokenType {
				return true
			}
		}
	}

	return false
}

func topLiteral(tokens []lexToken) string {
	if len(tokens) > 0 {
		return tokens[0].Literal
	} else {
		return "nothing"
	}
}