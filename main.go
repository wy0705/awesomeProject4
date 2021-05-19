package main

import (
	"awesomeProject4/io"
	"awesomeProject4/myfile"
)

func main() {
	/*//生成随机属性
	var ss string=io.Rand()
	fmt.Println(ss)
	for i := 0; i < 10; i++ {
		s:=io.Rand()
		ss=ss+"\n"+s
		fmt.Println(s)
	}
	//写入文件夹
	myfile.Mywritefile(ss)*/
	//natures:=myfile.CreatFile()
	//从文件夹中读出

	//str:=myfile.Myreadfile()
	//分成字符串数组
	//strs:=io.SplitArray(str)
	//写入test.go文件中
	//myfile.WriteGoFile(strs,natures)

	str:=myfile.Myreadfile()
	//分成字符串数组
	strs:=io.SplitArray(str)
	//生成新对象
	myfile.CreatObject(strs)
	//test.Myref()
}

