package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	// Identifiers + Literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operators
	ASSIGN     = "="
	PLUS       = "+"
	MINUS      = "-"
	STAR       = "*"
	SLASH      = "/"
	DOUBLESTAR = "**"

	LT        = "<"
	LTEQ      = "<="
	EQ        = "=="
	GTEQ      = ">="
	GT        = ">"
	NOTEQ     = "!="
	BOOLAND   = "&&"
	BOOLOR    = "||"
	BOOLXOR   = "^^"
	BANG      = "!"
	BOOLNOTEQ = "!="

	BITAND    = "&"
	BITOR     = "|"
	BITXOR    = "^"
	BITNOT    = "~"
	LBITSHIFT = "<<"
	RBITSHIFT = ">>"

	HASH     = "#"
	QUESTION = "?"
	COALESCE = "??"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	PACKAGE   = "PACKAGE"
	IMPORT    = "IMPORT"
	LET       = "LET"
	INT_KW    = "INT_KW"
	BOOL_KW   = "BOOL_KW"
	STRING_KW = "STRING_KW"
	FUNCTION  = "FUNCTION"
	CLASS     = "CLASS"
	THIS      = "THIS"
	TYPE      = "TYPE"
	ENUM      = "ENUM"
	UNION     = "UNION"
	COROUTINE = "COROUTINE"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	NULL      = "NULL"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
)

var keywords = map[string]TokenType{
	"package": PACKAGE,   // Priority 4
	"import":  IMPORT,    // Priority 4
	"let":     LET,       // Priority 1
	"int":     INT_KW,    // Priority 2
	"bool":    BOOL_KW,   // Priority 2
	"string":  STRING_KW, // Priority 2
	"fn":      FUNCTION,  // Priority 1
	"class":   CLASS,     // Priority 3
	"this":    THIS,      // Priority 3
	"type":    TYPE,      // Priority 3
	"enum":    ENUM,      // Priority 3
	"union":   UNION,     // Priority 3
	"co":      COROUTINE, // Priority 4
	"true":    TRUE,      // Priority 1
	"false":   FALSE,     // Priority 1
	"null":    NULL,      // Priority 1
	"if":      IF,        // Priority 1
	"else":    ELSE,      // Priority 1
	"return":  RETURN,    // Priority 1
}

func MatchIdent(ident string) TokenType {
	tokType, ok := keywords[ident]
	if ok {
		return tokType
	}

	return IDENT
}
