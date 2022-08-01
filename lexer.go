package main

type lexTokenType int

type lexToken struct {
	literal string
	tokenType lexTokenType
}

const (
	lexTokenType_name = iota
	lexTokenType_or
	lexTokenType_and
)

func lexProgramme(prog string) []lexToken {
	var out []lexToken //10 is an abitry starting value for a dynamic array
	var buffer string //used to store multicharacter token literals during lexing

	for _, letter := range prog {
		if(letter != ' ' && letter != '\n') {
			buffer += string(letter)
		} else {
			switch buffer {
			case "and": 
				out = append(out, lexToken{buffer, lexTokenType_and})
			case "or": 
				out = append(out, lexToken{buffer, lexTokenType_or})
			default:
				out = append(out, lexToken{buffer, lexTokenType_name})
			}
			buffer = ""
		}
	}

	return out
}