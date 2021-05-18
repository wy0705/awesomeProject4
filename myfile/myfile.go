package myfile

import (
	"awesomeProject4/io"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Mywritefile(s string) {
	userFile := "/home/wy/go/src/awesomeProject4/res/myfile.txt" //文件路径
	fout,err := os.Create(userFile) //根据路径创建File的内存地址
	defer fout.Close() //延迟关闭资源
	if err != nil{
		fmt.Println(userFile,err)
		return
	}

	fout.WriteString(s+"\n") //写入字符串
	//fout.Write([]byte("abcd!\n"))//强转成byte slice后再写入

}
func Myreadfile() string {
	userFile := "/home/wy/go/src/awesomeProject4/res/myfile.txt" //文件路径
	fin, err := os.Open(userFile)                                //打开文件,返回File的内存地址
	defer fin.Close()                                            //延迟关闭资源
	if err != nil {
		fmt.Println(userFile, err)
		return " "
	}
	buf := make([]byte, 1024) //创建一个初始容量为1024的slice,作为缓冲容器
	for {
		//循环读取文件数据到缓冲容器中,返回读取到的个数
		n, _ := fin.Read(buf)

		if 0 == n {
			break //如果读到个数为0,则读取完毕,跳出循环
		}
		//从缓冲slice中写出数据,从slice下标0到n,通过os.Stdout写出到控制台
		os.Stdout.Write(buf[:n])

	}
	var str string = string(buf[:])
	return str
}
//随机生成类型
func randNature() string {
	rand.Seed(time.Now().UnixNano())
	balance1:= [3]string{"int","string","float"}
	return balance1[rand.Intn(3)]
}
//随机生成字符串
func randstring() string {
	rand.Seed(time.Now().UnixNano())
	balance1:= [20]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t'}
	balance2:= [20]byte{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T'}
	result:=[4]byte{balance2[rand.Intn(20)],balance1[rand.Intn(20)],balance1[rand.Intn(20)],balance1[rand.Intn(20)]}
	var str string = string(result[:])
	return str

}
//随机整数
func randInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
func randFloat64(min,max float64) float64{
	//转为字符串
	my_min := strconv.FormatFloat(min,'f',-1,64)
	my_max := strconv.FormatFloat(max,'f',-1,64)

	if len(my_min)!=len(my_max){
		return 0
		//fmt.Println("错误")
	}else{
		qianzhui := ""
		houzhui1 := ""
		houzhui2 := ""
		sw := 0
		for i := 0;i<len(my_min);i++ {
			str1 := string(my_min[i])
			str2 := string(my_max[i])
			if sw==0{
				if my_min[i]==my_max[i]{
					qianzhui = qianzhui + str1
				}else{
					sw = 1
					houzhui1 = houzhui1 + str1
					houzhui2 = houzhui2 + str2
				}
			}else{
				houzhui1 = houzhui1 + str1
				houzhui2 = houzhui2 + str2
			}
		}
		//return 0,qianzhui+"**"+houzhui1+"**"+houzhui2
		//fmt.Println(qianzhui+"**"+houzhui1+"**"+houzhui2)
		//前缀字符串转整数
		qian1,_ := strconv.ParseInt(houzhui1,10,64)
		qian2,_ := strconv.ParseInt(houzhui2,10,64)
		//生成随机整数
		jj := randInt64(qian1,qian2)
		//随机整数转字符串
		jj_str := strconv.FormatInt(jj,10)
		//前缀拼接随机字符串
		all_str := qianzhui+jj_str
		//全字符串转浮点数
		jieguo,_ := strconv.ParseFloat(all_str,64)
		return jieguo
	}
	return 0
}

//随机生成类型值
func randresult(nature string) string {
	var s string
	switch nature {
	case "int": s=strconv.Itoa(rand.Intn(100))//随机生成数字
	case "string":s=randstring()//随机生成字符串
	case "float": s=strconv.FormatFloat(randFloat64(0.001234,0.009874), 'g', 5, 32)
	}
	return s
}
//写入go文件
func WriteGoFile(strs []string)  {
	userFile := "/home/wy/go/src/awesomeProject4/test/test.go" //文件路径
	fout,err := os.Create(userFile) //根据路径创建File的内存地址
	defer fout.Close() //延迟关闭资源
	if err != nil{
		fmt.Println(userFile,err)
		return
	}
	var s string="package test\n"

	natures:=[]string{}
	for i := 0; i < 22; i++ {
		nature:=randNature()
		s=s+"type "+strs[i]+nature+"\n"
		ii:=i+1
		natures=append(natures,strs[ii]+":"+randresult(nature))
		i++
	}
	s=s+"type Test struct {\n"
	for i := 0; i < 21; i++ {
		ii:=i+1
		s=s+strs[ii]
		s=s+" "+strs[i]+"\n"
		i++
	}
	s=s+"}\n"
	for i := 0; i < 10; i++ {
		ii:=i+1
		s=s+"func (t *Test)set"+io.Capitalize(strs[ii])+"("+strs[ii]+" "+strs[i]+") {\n"+"t."+strs[ii]+"="+strs[ii]+"\n}\n"
		s=s+"func (t *Test)get"+io.Capitalize(strs[ii])+"() "+strs[i]+"{\n"+"return t."+strs[ii]+"\n}\n"
        i++
	}
	fout.WriteString(s)

}

