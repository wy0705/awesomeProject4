package test

import (
	"fmt"
	"reflect"
)

type Oa int
type Ci int
type Jm int
type Ri int
type Os int
type Rr float64
type Ck string
type Tc float64
type Fe int
type Gf int
type Cs float64
type Test struct {
pq Oa
re Ci
ej Jm
hl Ri
fl Os
ak Rr
df Ck
hn Tc
js Fe
ma Gf
gd Cs
}
func (t *Test)setPq(pq Oa) {
t.pq=pq
}
func (t *Test)getPq() Oa{
return t.pq
}
func (t *Test)setRe(re Ci) {
t.re=re
}
func (t *Test)getRe() Ci{
return t.re
}
func (t *Test)setEj(ej Jm) {
t.ej=ej
}
func (t *Test)getEj() Jm{
return t.ej
}
func (t *Test)setHl(hl Ri) {
t.hl=hl
}
func (t *Test)getHl() Ri{
return t.hl
}
func (t *Test)setFl(fl Os) {
t.fl=fl
}
func (t *Test)getFl() Os{
return t.fl
}
func Myref() {
t:=Test{pq:75,re:95,ej:41,hl:29,fl:23,ak:0.006066,df:"Hhhl",hn:0.008828,js:22,ma:37,gd:0.007192}
a:=reflect.TypeOf(t)
	b:=reflect.ValueOf(t)
for i := 0; i < a.NumField(); i++ {
		key:=a.Field(i)
		val:=b.Field(i)
		fmt.Println(key.Name,key.Type,"|",val)
	}
}