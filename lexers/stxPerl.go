//go:build stxPerl || stxAll

package lexers

func init() {
	syntaxMap["perl"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: perl
rules:
    - type: "\\b(accept|alarm|atan2|bin(d|mode)|c(aller|h(dir|mod|op|own|root)|lose(dir)?|onnect|os|rypt)|d(bm(close|open)|efined|elete|ie|o|ump)|e(ach|of|val|x(ec|ists|it|p))|f(cntl|ileno|lock|ork))\\b|\\b(get(c|login|peername|pgrp|ppid|priority|pwnam|(host|net|proto|serv)byname|pwuid|grgid|(host|net)byaddr|protobynumber|servbyport)|([gs]et|end)(pw|gr|host|net|proto|serv)ent|getsock(name|opt)|gmtime|goto|grep|hex|index|int|ioctl|join)\\b|\\b(keys|kill|last|length|link|listen|local(time)?|log|lstat|m|mkdir|msg(ctl|get|snd|rcv)|next|oct|open(dir)?|ord|pack|pipe|pop|printf?|push|q|qq|qx|rand|re(ad(dir|link)?|cv|do|name|quire|set|turn|verse|winddir)|rindex|rmdir|s|scalar|seek(dir)?)\\b|\\b(se(lect|mctl|mget|mop|nd|tpgrp|tpriority|tsockopt)|shift|shm(ctl|get|read|write)|shutdown|sin|sleep|socket(pair)?|sort|spli(ce|t)|sprintf|sqrt|srand|stat|study|substr|symlink|sys(call|read|tem|write)|tell(dir)?|time|tr(y)?|truncate|umask)\\b|\\b(un(def|link|pack|shift)|utime|values|vec|wait(pid)?|wantarray|warn|write)\\b"
    - statement: "\\b(continue|else|elsif|do|for|foreach|if|unless|until|while|eq|ne|lt|gt|le|ge|cmp|x|my|sub|use|package|can|isa)\\b"
    - identifier:
        start: "[$@%]"
        end: "((?i) |[^0-9A-Z_]|-)"
        rules: []
    - constant.string: "\".*\"|qq\\|.*\\|"
    - default: "[sm]/.*/"
    - preproc:
        start: "(^use| = new)"
        end: ";"
        rules: []
    - comment: "#.*"
    - identifier.macro:
        start: "<< 'STOP'"
        end: "STOP"
        rules: []`)
	}}
}
