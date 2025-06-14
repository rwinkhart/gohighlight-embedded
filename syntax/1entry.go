package syntax

import "sync"

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
