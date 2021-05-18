package io

import "fmt"

func Debalk(s string)  string{
	var data []byte = []byte(s)
	/*s1 := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

	//在原切片的末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)*/
	result :=[]byte{}
	for i := 0; i < len(data); i++ {
		if data[i]!=' '&&data[i]!='\n' {
			result=append(result,data[i])
		}
	}
	var str string = string(result[:])
	return str
}
func SplitArray(s string)  []string{
	var data []byte = []byte(s)
	str:=[2]byte{}
	strs:=[]string{}
	for i := 0; i < len(data); i++ {
		str[0]=data[i]
		i++
		str[1]=data[i]
		var result string = string(str[:])
		strs=append(strs,result)
	}
	return strs
}
// 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}