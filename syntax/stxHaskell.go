//go:build stxHaskell || stxAll

package syntax

func init() {
	syntaxMap["haskell"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: haskell
rules:
    - statement: "[ ](as|case|of|class|data|default|deriving|do|forall|foreign|hiding|if|then|else|import|infix|infixl|infixr|instance|let|in|mdo|module|newtype|qualified|type|where)[ ]"
    - statement: "(^data|^foreign|^import|^infix|^infixl|^infixr|^instance|^module|^newtype|^type)[ ]"
    - statement: "[ ](as$|case$|of$|class$|data$|default$|deriving$|do$|forall$|foreign$|hiding$|if$|then$|else$|import$|infix$|infixl$|infixr$|instance$|let$|in$|mdo$|module$|newtype$|qualified$|type$|where$)"
    - symbol: "(\\||@|!|:|_|~|=|\\\\|;|\\(\\)|,|\\[|\\]|\\{|\\})"
    - symbol.operator: "(==|/=|&&|\\|\\||<|>|<=|>=)"
    - special: "(->|<-)"
    - symbol: "\\.|\\$"
    - constant.bool: "\\b(True|False)\\b"
    - constant: "(Nothing|Just|Left|Right|LT|EQ|GT)"
    - identifier.class: "[ ](Read|Show|Enum|Eq|Ord|Data|Bounded|Typeable|Num|Real|Fractional|Integral|RealFrac|Floating|RealFloat|Monad|MonadPlus|Functor)"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - comment:
        start: "--"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"
    - comment:
        start: "\\{-"
        end: "-\\}"
        rules:
            - todo: "(TODO|XXX|FIXME):?"
    - identifier.micro: "undefined"`)
	}}
}
