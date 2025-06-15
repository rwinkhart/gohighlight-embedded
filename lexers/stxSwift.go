//go:build stxSwift || stxAll

package lexers

func init() {
	syntaxMap["swift"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: swift
rules:
    - statement: "([.:;,+*|=!?\\%]|<|>|/|-|&)"
    - statement:  "(class|import|let|var|struct|enum|func|if|else|switch|case|default|for|in|internal|external|unowned|private|public|throws)\\ "
    - statement:  "(prefix|postfix|operator|extension|lazy|get|set|self|willSet|didSet|override|super|convenience|weak|strong|mutating|return|guard)\\ "
    - statement: "(print)"
    - statement: "(init)"
    - constant.number: "([0-9]+)"
    - type: "\\ ((U)?Int(8|16|32|64))"
    - constant: "(true|false|nil)"
    - type: "\\ (Double|String|Float|Boolean|Dictionary|Array|Int)"
    - type: "\\ (AnyObject)"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - comment:
        start: "//"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"
    - comment:
        start: "///"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"
    - comment:
        start: "/\\*\\*"
        end: "\\*/"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
