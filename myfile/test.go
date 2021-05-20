package myfile

import (
	"fmt"
	"reflect"
)

type Test struct {
Cq  string
Lq  int
Ci  string
Rf  int
Rh  string
Ko  int
Ed  string
Ni  string
Kr  string
Li  int
}
func (t *Test)SetCq (cq  string) {
t.Cq =cq
}
func (t *Test)GetCq () string{
return t.Cq
}
func (t *Test)SetLq (lq  int) {
	t.Lq =lq
}
func (t *Test)GetLq () int{
	return t.Lq
}
func (t *Test)SetCi (ci  string) {
t.Ci =ci
}
func (t *Test)GetCi () string{
return t.Ci
}
func (t *Test)SetRf (rf  int) {
	t.Rf =rf
}
func (t *Test)GetRf () int{
	return t.Rf
}
func (t *Test)SetRh (rh  string) {
t.Rh =rh
}
func (t *Test)GetRh () string{
return t.Rh
}
func (t *Test)SetKo (ko  int) {
	t.Ko =ko
}
func (t *Test)GetKo () int{
	return t.Ko
}
func (t *Test)SetEd (ed  string) {
t.Ed =ed
}
func (t *Test)GetEd () string{
return t.Ed
}
func (t *Test)GetNi () string{
	return t.Ni
}
func (t *Test)SetNi (ni  string) {
	t.Ni =ni
}
func (t *Test)SetKr (kr  string) {
	t.Kr =kr
}
func (t *Test)GetKr () string{
return t.Kr
}
func Myref() {
t:=Test{
Cq :"Obcm",Lq :43,Ci :"Qebh",Rf :23,Rh :"Tlrg",Ko :51,Ed :"Pres",Ni :"Pqdj",Kr :"Raak",Li :5}
a:=reflect.TypeOf(t)
	b:=reflect.ValueOf(t)
for i := 0; i < a.NumField(); i++ {
		key:=a.Field(i)
		val:=b.Field(i)
		fmt.Println(key.Name,key.Type,"|",val)
	}

}