//go:build stxIno || stxAll

package lexers

func init() {
	syntaxMap["ino"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: ino
rules:
    - identifier: "\\b[A-Z_][0-9A-Z_]+\\b"
    - type: "\\b((s?size)|((u_?)?int(8|16|32|64|ptr)))_t\\b"
    - constant: "(?i)\\b(HIGH|LOW|INPUT|OUTPUT)\\b"
    - constant: "(?i)\\b(DEC|BIN|HEX|OCT|BYTE)\\b"
    - constant: "(?i)\\b(PI|HALF_PI|TWO_PI)\\b"
    - constant: "(?i)\\b(LSBFIRST|MSBFIRST)\\b"
    - constant: "(?i)\\b(CHANGE|FALLING|RISING)\\b"
    - constant: "(?i)\\b(DEFAULT|EXTERNAL|INTERNAL|INTERNAL1V1|INTERNAL2V56)\\b"
    - type: "\\b(boolean|byte|char|float|int|long|word)\\b"
    - statement: "\\b(case|class|default|do|double|else|false|for|if|new|null|private|protected|public|short|signed|static|String|switch|this|throw|try|true|unsigned|void|while)\\b"
    - statement: "\\b(goto|continue|break|return)\\b"
    - identifier: "\\b(abs|acos|asin|atan|atan2|ceil|constrain|cos|degrees|exp|floor|log|map|max|min|radians|random|randomSeed|round|sin|sq|sqrt|tan)\\b"
    - identifier: "\\b(bitRead|bitWrite|bitSet|bitClear|bit|highByte|lowByte)\\b"
    - identifier: "\\b(analogReference|analogRead|analogWrite)\\b"
    - identifier: "\\b(attachInterrupt|detachInterrupt)\\b"
    - identifier: "\\b(delay|delayMicroseconds|millis|micros)\\b"
    - identifier: "\\b(pinMode|digitalWrite|digitalRead)\\b"
    - identifier: "\\b(interrupts|noInterrupts)\\b"
    - identifier: "\\b(noTone|pulseIn|shiftIn|shiftOut|tone)\\b"
    - identifier: "\\b(Serial|Serial1|Serial2|Serial3|begin|end|peek|read|print|println|available|flush)\\b"
    - identifier: "\\b(setup|loop)\\b"
    - statement: "^[[:space:]]*#[[:space:]]*(define|include(_next)?|(un|ifn?)def|endif|el(if|se)|if|warning|error|pragma)"
    - constant: "(__attribute__[[:space:]]*\\(\\([^)]*\\)\\)|__(aligned|asm|builtin|hidden|inline|packed|restrict|section|typeof|weak)__)"
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
            - preproc: "..+"
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
