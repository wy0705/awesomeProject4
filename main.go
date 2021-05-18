package main

import (
	"awesomeProject4/io"
	"awesomeProject4/myfile"
	"fmt"
)

func main() {
	//生成随机属性
	var ss string=io.Rand()
	fmt.Println(ss)
	for i := 0; i < 10; i++ {
		s:=io.Rand()
		ss=ss+"\n"+s
		fmt.Println(s)
	}
	//写入文件夹
	myfile.Mywritefile(ss)
	//从文件夹中读出
	str:=io.Debalk(myfile.Myreadfile())
	//分成字符串数组
	strs:=io.SplitArray(str)
	//写入test.go文件中
	myfile.WriteGoFile(strs)

}

