package backtracking

type Pattern struct {
	matched bool   // 是否匹配
	pattern []rune //正则表达式
	plen    int    // 正则表达式长度
}

func (p *Pattern) match(text []rune, tlen int) bool {
	p.matched = false
	p.rmatch(0, 0, text, tlen)
	return p.matched
}

func (p *Pattern) rmatch(ti, pj int, text []rune, tlen int) {
	if p.matched {
		return
	}
	if pj == p.plen {
		if ti == tlen {
			p.matched = true
		}
		return
	}
	if p.pattern[pj] == '*' {
		for k := 0; k <= tlen-ti; k++ {
		}
	}
}
