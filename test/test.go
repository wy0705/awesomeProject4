package test

import (
	"fmt"
	"reflect"
)

type Test struct {
cq  string
lq  int
ci  string
rf  int
rh  string
ko  int
ed  string
ni  string
kr  string
li  int
}
func (t *Test)SetCq (cq  string) {
t.cq =cq 
}
func (t *Test)GetCq () string{
return t.cq 
}
func (t *Test)SetLq (lq  int) {
	t.lq =lq
}
func (t *Test)GetLq () int{
	return t.lq
}
func (t *Test)SetCi (ci  string) {
t.ci =ci 
}
func (t *Test)GetCi () string{
return t.ci 
}
func (t *Test)SetRf (rf  int) {
	t.rf =rf
}
func (t *Test)GetRf () int{
	return t.rf
}
func (t *Test)SetRh (rh  string) {
t.rh =rh 
}
func (t *Test)GetRh () string{
return t.rh 
}
func (t *Test)SetKo (ko  int) {
	t.ko =ko
}
func (t *Test)GetKo () int{
	return t.ko
}
func (t *Test)SetEd (ed  string) {
t.ed =ed 
}
func (t *Test)GetEd () string{
return t.ed 
}
func (t *Test)GetNi () string{
	return t.ni
}
func (t *Test)SetNi (ni  string) {
	t.ni =ni
}
func (t *Test)SetKr (kr  string) {
	t.kr =kr
}
func (t *Test)GetKr () string{
return t.kr 
}
func Myref() {
t:=Test{
cq :"Obcm",lq :43,ci :"Qebh",rf :23,rh :"Tlrg",ko :51,ed :"Pres",ni :"Pqdj",kr :"Raak",li :5}
a:=reflect.TypeOf(t)
	b:=reflect.ValueOf(t)
for i := 0; i < a.NumField(); i++ {
		key:=a.Field(i)
		val:=b.Field(i)
		fmt.Println(key.Name,key.Type,"|",val)
	}

}