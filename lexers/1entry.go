package lexers

import "sync"

// Syntax Rules Master List
// comment
// constant
// constant.bool
// constant.number
// constant.specialChar
// constant.string
// error
// identifier
// identifier.macro
// identifier.var
// keyword
// preproc
// special
// statement
// symbol
// symbol.brackets
// symbol.operator
// todo
// type

var syntaxMap = make(map[string]*lazySyntax)

type lazySyntax struct {
	once        sync.Once
	syntaxLayer []byte
	init        func() []byte
}

func (ls *lazySyntax) get() []byte {
	ls.once.Do(func() {
		ls.syntaxLayer = ls.init()
	})
	return ls.syntaxLayer
}

func Get(id string) []byte {
	if syntax, exists := syntaxMap[id]; exists {
		return syntax.get()
	}
	return nil
}
