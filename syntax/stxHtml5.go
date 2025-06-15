//go:build stxHtml5 || stxAll

package syntax

func init() {
	syntaxMap["html5"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: html5
rules:
    - error: "<[^!].*?>"
    - symbol: "(?i)<[/]?(a|a(bbr|ddress|rea|rticle|side|udio)|b|b(ase|d(i|o)|lockquote|r|utton)|ca(nvas|ption)|center|cite|co(de|l|lgroup)|d(ata|atalist|d|el|etails|fn|ialog|l|t)|em|embed|fieldset|fig(caption|ure)|form|iframe|h[1-6]|hr|i|img|in(put|s)|kbd|keygen|label|legend|li|link|ma(in|p|rk)|menu|menuitem|met(a|er)|nav|noscript|o(bject|l|pt(group|ion)|utput)|p|param|picture|pre|progress|q|r(p|t|uby)|s|samp|se(ction|lect)|small|source|span|strong|su(b|p|mmary)|textarea|time|track|u|ul|var|video|wbr)( .*)*?>"
    # todo used in place of type.extended to reduce complexity (only unused rule in stxPhp.go)
    - todo: "(?i)<[/]?(body|div|html|head(er)?|footer|title|table|t(body|d|h(ead)?|r|foot))( .*)*?>"
    - preproc: "(?i)<[/]?(script|style)( .*)*?>"
    - special: "&[^;[[:space:]]]*;"
    - symbol: "[:=]"
    - identifier: "(alt|bgcolor|height|href|id|label|longdesc|name|on(click|focus|load|mouseover)|size|span|src|style|target|type|value|width)="
    - constant.string: "\"[^\"]*\""
    - constant.number: "(?i)#[0-9A-F]{6,6}"
    - default:
        start: ">"
        end: "<"
        rules: []
    - symbol: "<|>"
    # constant.number is used in place of constant.string.url to reduce complexity
    - constant.number: "(ftp(s)?|http(s)?|git|chrome)://[^ 	]+"
    - comment: "<!--.+?-->"
    - preproc: "<!DOCTYPE.+?>"`)
	}}
}
