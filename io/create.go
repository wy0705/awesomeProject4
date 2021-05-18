package io

import (
	"math/rand"
	"time"
)

func Rand() string {
	rand.Seed(time.Now().UnixNano())
	balance1:= [20]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t'}
	balance2:= [20]byte{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T'}
	result:=[5]byte{balance2[rand.Intn(20)],balance1[rand.Intn(20)],' ',balance1[rand.Intn(20)],balance1[rand.Intn(20)]}
	var str string = string(result[:])
	return str

}
