package Parser

type lexTokenType int

type lexToken struct {
	Literal string
	TokenType lexTokenType
}

const (
	LexTokenType_name = iota
	LexTokenType_or
	LexTokenType_and
	LexTokenType_less
	LexTokenType_lParen
	LexTokenType_rParen
	LexTokenType_percent
	LexTokenType_plus
	LexTokenType_minus
)

func LexProgramme(prog string) []lexToken {
	var out []lexToken //10 is an abitry starting value for a dynamic array
	var buffer string //used to store multicharacter token literals during lexing

	//append a dummy whitespace to the back of the programme so we don't have to manually flush the buffer at EOF
	prog += " "

	for _, letter := range prog {
		if(isBufferChar(letter)) {
			buffer += string(letter)
		} else {
			switch buffer {
			case "and": 
				out = append(out, lexToken{buffer, LexTokenType_and})
			case "or": 
				out = append(out, lexToken{buffer, LexTokenType_or})
			case "less":
				out = append(out, lexToken{buffer, LexTokenType_less})
			case "-":
				out = append(out, lexToken{buffer, LexTokenType_minus})
			default:
				if len(buffer) != 0 {
					out = append(out, lexToken{buffer, LexTokenType_name})
				}
			}
			buffer = ""
		}

		switch letter {
		case '(':
			out = append(out, lexToken{"'('", LexTokenType_lParen})
		case ')':
			out = append(out, lexToken{"')'", LexTokenType_rParen})
		case '%':
			out = append(out, lexToken{"%", LexTokenType_percent})
		case '+':
			out = append(out, lexToken{"+", LexTokenType_plus})
		}
	}

	return out
}

func isBufferChar(target rune) bool {
	for _, char := range []rune{'(',')','%',' ','\n', '+'} {
		if char == target {
			return false
		}
	}
	return true
}