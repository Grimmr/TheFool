package Parser

import (
	"os"
)

type ParseTreeNode struct {
	token lexToken
	children []*ParseTreeNode
}

func ParseProgramme(tokens []lexToken) *ParseTreeNode {
	out, err, _ := expandExpr1(tokens)
	if out == nil {
		os.Stderr.WriteString(err + "\n")
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
	if !expect(tokens, []lexTokenType{lexTokenType_name}) {
		return nil, "expected table name but got " + topLiteral(tokens), tokens
	}

	return &ParseTreeNode{token: tokens[0]}, "", tokens[1:]
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
	var out *ParseTreeNode = new(ParseTreeNode)
	out.token = operators[0]
	out.children = append(out.children, children[0])
	out.children = append(out.children, children[1])
	for i := range operators[1:] {
		var holder *ParseTreeNode = new(ParseTreeNode)
		holder.token = operators[i+1]
		holder.children = append(holder.children, out)
		holder.children = append(holder.children, children[2+i])
		out = holder
	}


	//return the generated node and consume tokens
	return out, "", tokens
}

func WalkParseTree(node *ParseTreeNode) string {
	out := "(" + node.token.literal 
	for _, element := range node.children {
		out += " " + WalkParseTree(element)
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