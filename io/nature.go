package io

import (
	"fmt"
	"reflect"
)

type Test struct {
	am string
}

func (t *Test)setAm(am string)  {
	t.am=am
}
func (t *Test)getAm() string {
	return t.am
}
func Myref()  {
	t:=Test{am:"123"}
	a:=reflect.TypeOf(t)
	b:=reflect.ValueOf(t)
	for i := 0; i < a.NumField(); i++ {
		key:=a.Field(i)
		val:=b.Field(i)
		fmt.Println(key.Name,key.Type,"|",val)
	}

}
