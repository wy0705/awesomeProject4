package io

type Aa int
type Test struct {
	am Aa
}

func (t *Test)setAm(am Aa)  {
	t.am=am
}
func (t *Test)getAm() Aa {
	return t.am
}

