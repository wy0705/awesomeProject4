package io

import (
	"fmt"
	"reflect"
)

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
func Myref()  {
	t:=Test{am:1}
	a:=reflect.TypeOf(t)
	b:=reflect.ValueOf(t)
	for i := 0; i < a.NumField(); i++ {
		key:=a.Field(i)
		val:=b.Field(i)
		fmt.Println(key.Name,key.Type,"|",val)
	}
}
