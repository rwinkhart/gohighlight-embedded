//go:build stxPhp || stxAll

package lexers

func init() {
	syntaxMap["php"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: php
rules:
    - symbol.operator: "<|>"
    - error: "<[^!].*?>"
    - symbol: "(?i)<[/]?(a(bbr|cronym|ddress|pplet|rea|rticle|side|udio)?|b(ase(font)?|d(i|o)|ig|lockquote|r)?|ca(nvas|ption)|center|cite|co(de|l|lgroup)|d(ata(list)?|d|el|etails|fn|ialog|ir|l|t)|em(bed)?|fieldset|fig(caption|ure)|font|form|(i)?frame|frameset|h[1-6]|hr|i|img|in(put|s)|kbd|keygen|label|legend|li(nk)?|ma(in|p|rk)|menu(item)?|met(a|er)|nav|no(frames|script)|o(l|pt(group|ion)|utput)|p(aram|icture|re|rogress)?|q|r(p|t|uby)|s(trike)?|samp|se(ction|lect)|small|source|span|strong|su(b|p|mmary)|textarea|time|track|u(l)?|var|video|wbr)( .*|>)*?>"
    # todo used in place of type.extended to reduce complexity
    - todo: "(?i)<[/]?(body|div|html|head(er)?|footer|title|table|t(body|d|h(ead)?|r|foot))( .*|>)*?>"
    - preproc: "(?i)<[/]?(script|style)( .*|>)*?>"
    - special: "&[^;[[:space:]]]*;"
    - symbol: "[:=]"
    - identifier: "(alt|bgcolor|height|href|label|longdesc|name|onclick|onfocus|onload|onmouseover|size|span|src|style|target|type|value|width)="
    - constant.string: "\"[^\"]*\""
    - constant.number: "(?i)#[0-9A-F]{6,6}"
    # constant.number is used in place of constant.string.url to reduce complexity
    - constant.number: "(ftp(s)?|http(s)?|git|chrome)://[^ 	]+"
    - comment: "<!--.+?-->"
    - default: "<\\?(php|=)\" end=\"\\?>"
    - identifier.macro: "([a-zA-Z0-9_-]+)\\("
    - preproc: "^\\s*(require|include|require_once|include_once)"
    - identifier.macro: "[a-zA-Z\\\\]+::"
    - identifier: "([A-Z][a-zA-Z0-9_]+)\\s"
    - identifier: "([A-Z0-9_]+)[;|\\s|\\)|,]"
    - keyword: "(global|public|private|protected|static|const)"
    - type: "\\b(var|class|extends|function|function_exists|__construct|phpinfo|echo|case|default|exit|switch|extends|as|define|do|declare|in|trait|interface|[E|e]xception|array|int|string|bool|iterable|void)\\b"
    - statement: "(^|[[:space:]])(implements|abstract|instanceof|if|else(if)?|endif|namespace|use|as|new|throw|catch|try|while|print|for|(end)?(foreach)?)\\b"
    - identifier: "new\\s([a-zA-Z0-9\\\\]+)"
    - special: "(break|continue|goto|return)"
    - constant.bool: "(\\:|\\?|=|[[:space:]])(true|false|null|TRUE|FALSE|NULL)"
    - constant: "[\\s|=|\\s|\\(|/|+|-|\\*|\\[]"
    - constant.number: "[0-9]"
    - identifier: "(\\$this|parent|self|\\$this->)"
    - symbol.operator: "(=>|===|!==|==|!=|&&|\\|\\||::|=|->|\\!)"
    - identifier.var: "(\\$[a-zA-Z0-9\\-_]+)"
    - symbol.operator: "[\\(|\\)|/|+|\\-|\\*|\\[|.|,|;]"
    - constant.string: "\"(\\\\.|[^\"])*\"|'(\\\\.|[^'])*'"
    - constant.specialChar: "\\\\[abfnrtv'\\\"\\\\]"
    - symbol.brackets: "(\\[|\\]|\\{|\\}|[()])"
    - comment: "(^|[[:space:]])//.*"
    - comment: "(^|[[:space:]])#.*"
    - comment:
        start: "/\\*"
        end: "\\*/"
        rules: []
    - preproc: "<\\?(php|=)?"
    - preproc: "\\?>"
    - preproc: "<!DOCTYPE.+?>"`)
	}}
}
