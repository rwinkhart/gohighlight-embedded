//go:build stxMakefile || stxAll

package lexers

func init() {
	syntaxMap["makefile"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: makefile
rules:
    - preproc: "(ifeq|ifdef|ifneq|ifndef|else|endif)"
    - statement: "^(export|include|override)\\>"
    - type: "(^[A-Za-z_]+[[:space:]]+)(\\?|:)?="
    - operator: "^[^:=	]+:"
    - operator: "([=,%]|\\+=|\\?=|:=|&&|\\|\\|)"
    - statement: "(abspath|addprefix|addsuffix|and|basename|call|([[:space:]]|^)([a-z]+)?dir|mkdir)[[:space:]]"
    - statement: "(([a-z]+(-))?(un)?install(-)?([a-z]+)?(-)?([a-z]+))"
    - statement: "\\$\\((error|eval|filter|filter-out|findstring|firstword)[[:space:]]"
    - statement: "(flavor|echo|foreach|if|else|fi(;)?|elif|case|esac(;)?|for|done|shift|chmod|chown(info)?|pwd|join|lastword|notdir|or)[[:space:]]"
    - statement: "\\$\\((origin|patsubst|realpath|shell|sort|strip|suffix)[[:space:]]"
    - statement: "\\$\\((value|warning|wildcard|word|wordlist|words)[[:space:]]"
    - keyword: "([[:space:]]|^)(rm|rmdir|exit)[[:space:]]"
    - special: "([[:space:]]|^)(gcc|ld)[[:space:]]"
    - type: "^.+:"
    - identifier: "[()$]"
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
    - identifier: "\\$+(\\{[^} ]+\\}|\\([^) ]+\\))"
    - identifier: "\\$[@^<*?%|+]|\\$\\([@^<*?%+-][DF]\\)"
    - identifier: "\\$\\$|\\\\.?"
    - comment:
        start: "#"
        end: "$"
        rules: []`)
	}}
}
