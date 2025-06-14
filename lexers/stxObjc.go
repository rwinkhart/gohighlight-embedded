//go:build stxObjc || stxAll

package lexers

func init() {
	syntaxMap["objc"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: objective-c
rules:
    - type: "\\b(float|double|CGFloat|id|bool|BOOL|Boolean|char|int|short|long|sizeof|enum|void|static|const|struct|union|typedef|extern|(un)?signed|inline|Class|SEL|IMP|NS(U)?Integer)\\b"
    - type: "\\b((s?size)|((u_?)?int(8|16|32|64|ptr)))_t\\b"
    - type: "\\b[A-Z][A-Z][[:alnum:]]*\\b"
    - type: "\\b[A-Za-z0-9_]*_t\\b"
    - type: "\\bdispatch_[a-zA-Z0-9_]*_t\\b"
    - statement: "(__attribute__[[:space:]]*\\(\\([^)]*\\)\\)|__(aligned|asm|builtin|hidden|inline|packed|restrict|section|typeof|weak)__|__unused|_Nonnull|_Nullable|__block|__builtin.*)"
    - statement: "\\b(class|namespace|template|public|protected|private|typename|this|friend|virtual|using|mutable|volatile|register|explicit)\\b"
    - statement: "\\b(for|if|while|do|else|case|default|switch)\\b"
    - statement: "\\b(try|throw|catch|operator|new|delete)\\b"
    - statement: "\\b(goto|continue|break|return)\\b"
    - statement: "\\b(nonatomic|atomic|readonly|readwrite|strong|weak|assign)\\b"
    - statement: "@(encode|end|interface|implementation|class|selector|protocol|synchronized|try|catch|finally|property|optional|required|import|autoreleasepool)"
    - preproc: "^[[:space:]]*#[[:space:]]*(define|include|import|(un|ifn?)def|endif|el(if|se)|if|warning|error|pragma).*$"
    - preproc: "__[A-Z0-9_]*__"
    - special: "^[[:space:]]*[#|@][[:space:]]*(import|include)[[:space:]]*[\"|<].*\\/?[>|\"][[:space:]]*$"
    - statement: "([.:;,+*|=!\\%\\[\\]]|<|>|/|-|&)"
    - constant.number: "(\\b(-?)?[0-9]+\\b|\\b\\[0-9]+\\.[0-9]+\\b|\\b0x[0-9A-F]+\\b)"
    - constant: "(@\\[(\\\\.|[^\\]])*\\]|@\\{(\\\\.|[^\\}])*\\}|@\\((\\\\.|[^\\)])*\\))"
    - constant: "\\b<(\\\\.[^\\>])*\\>\\b"
    - constant: "\\b(nil|NULL|YES|NO|TRUE|true|FALSE|false|self)\\b"
    - constant: "\\bk[[:alnum]]*\\b"
    - constant.string: "'.'"
    - constant.string:
        start: "@\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
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
        start: "/\\*"
        end: "\\*/"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
