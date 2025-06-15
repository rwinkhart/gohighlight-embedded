//go:build stxPython || stxAll

package lexers

func init() {
	syntaxMap["python"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: python
rules:
    - constant: "\\b(None|self|True|False)\\b"
    - constant: "\\b(__bases__|__builtin__|__class__|__debug__|__dict__|__doc__|__file__|__members__|__methods__|__name__|__self__)\\b"
    - identifier: "\\b(abs|all|any|ascii|bin|callable|chr|compile|delattr|dir|divmod|eval|exec|format|getattr|globals|hasattr|hash|help|hex|id|input|isinstance|issubclass|iter|len|locals|max|min|next|oct|open|ord|pow|print|repr|round|setattr|sorted|sum|vars|__import__)\\b"
    - identifier: "\\b(__abs__|__add__|__and__|__call__|__cmp__|__coerce__|__complex__|__concat__|__contains__|__del__|__delattr__|__delitem__|__delslice__|__div__|__divmod__|__float__|__getattr__|__getitem__|__getslice__|__hash__|__hex__|__init__|__int__|__inv__|__invert__|__len__|__dict__|__long__|__lshift__|__mod__|__mul__|__neg__|__next__|__nonzero__|__oct__|__or__|__pos__|__pow__|__radd__|__rand__|__rcmp__|__rdiv__|__rdivmod__|__repeat__|__repr__|__rlshift__|__rmod__|__rmul__|__ror__|__rpow__|__rrshift__|__rshift__|__rsub__|__rxor__|__setattr__|__setitem__|__setslice__|__str__|__sub__|__xor__)\\b"
    - type: "\\b(bool|bytearray|bytes|classmethod|complex|dict|enumerate|filter|float|frozenset|int|list|map|memoryview|object|property|range|reversed|set|slice|staticmethod|str|super|tuple|type|zip)\\b"
    - identifier: "def [a-zA-Z_0-9]+"
    - statement: "\\b(and|as|assert|break|class|continue|def|del|elif|else|except|finally|for|from|global|if|import|in|is|lambda|nonlocal|not|or|pass|raise|return|try|while|with|yield)\\b"
    - special: "@.*[(]"
    - statement: "([.:;,+*|=!\\%@]|<|>|/|-|&)"
    - statement: "([(){}]|\\[|\\])"
    - constant.number: "\\b[0-9]+\\b"
    - comment:
        start: "\"\"\""
        end: "\"\"\""
        rules: []
    - comment:
        start: "'''"
        end: "'''"
        rules: []
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "'"
        end: "'"
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - comment:
        start: "#"
        end: "$"
        rules: []`)
	}}
}
