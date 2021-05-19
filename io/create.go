package io

import (
	"math/rand"
	"time"
)

func Rand() string {
	rand.Seed(time.Now().UnixNano())
	balance1:= [20]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t'}
	//balance2:= [20]byte{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T'}
	result:=[2]byte{balance1[rand.Intn(20)],balance1[rand.Intn(20)]}
	var str string = string(result[:])
	return str

}

/*func CreatFile()  {
	var str string=Rand()
	str=str+" "+myfile.Randresult(myfile.RndNature())+"\n"
	for i := 0; i < 10; i++ {
		str=str+Rand()+" "+myfile.Randresult(myfile.RndNature())+"\n"
	}
	myfile.Mywritefile(str)

}*/
