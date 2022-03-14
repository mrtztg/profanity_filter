package filterProfanity

import (
	"regexp"
)

var RegexOnlyWords = regexp.MustCompile(`[A-Za-zآ-ی]`)

type Engine struct {
	blacklist *[]rune
	whitelist *[]rune
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) SetWhiteList(ls []rune) *Engine {
	e.whitelist = &ls
	return e
}

func (e *Engine) SetBlacklist(ls []rune) *Engine {
	e.blacklist = &ls
	return e
}

func (e Engine) Filter(val string) string {
	return "" //TODO
}
