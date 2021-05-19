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
	userFile := "/home/wy/go/src/awesomeProject4/res/myfile1.txt" //文件路径
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
func RndNature() string {
	rand.Seed(time.Now().UnixNano())
	balance1:= [3]string{"int","string","float64"}
	return balance1[rand.Intn(3)]
}
//随机生成字符串
func Randstring() string {
	rand.Seed(time.Now().UnixNano())
	balance1:= [20]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t'}
	balance2:= [20]byte{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T'}
	result:=[6]byte{'"',balance2[rand.Intn(20)],balance1[rand.Intn(20)],balance1[rand.Intn(20)],balance1[rand.Intn(20)],'"'}
	var str string = string(result[:])
	return str

}
//随机整数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
func RandFloat64(min,max float64) float64{
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
		jj := RandInt64(qian1,qian2)
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
func Randresult(nature string) string {
	var s string
	switch nature {
	case "int": s=strconv.Itoa(rand.Intn(100))//随机生成数字
	case "string":s=Randstring()//随机生成字符串
	case "float64": s=strconv.FormatFloat(RandFloat64(0.001234,0.009874), 'g', 5, 32)
	}
	return s
}
//写入go文件
func WriteGoFile(strs []string,natures []string)  {
	userFile := "/home/wy/go/src/awesomeProject4/test/test.go" //文件路径
	fout,err := os.Create(userFile) //根据路径创建File的内存地址
	defer fout.Close() //延迟关闭资源
	if err != nil{
		fmt.Println(userFile,err)
		return
	}
	var s string="package test\n"

	s=s+"type Test struct {\n"
	for i := 0; i < 10; i++ {
		s=s+strs[i*2]+" "+natures[i]+"\n"
	}
	s=s+"}\n"
	for i := 0; i < 10; i++ {
		s=s+"func (t *Test)Set"+io.Capitalize(strs[i*2])+"("+strs[i*2]+" "+natures[i]+") {\n"+"t."+strs[i*2]+"="+strs[i*2]+"\n}\n"
		s=s+"func (t *Test)Get"+io.Capitalize(strs[i*2])+"() "+natures[i]+"{\n"+"return t."+strs[i*2]+"\n}\n"
	}
	s=s+"func Myref() {\nt:=Test{\n"
	for i := 0; i < 11; i++ {
		if i!=10 {
			s=s+strs[i*2]+":"+strs[i*2+1]+","
		}else {
			s=s+strs[i*2]+":"+strs[i*2+1]+"}\n"
		}

	}
	s=s+"a:=reflect.TypeOf(t)\n\tb:=reflect.ValueOf(t)\nfor i := 0; i < a.NumField(); i++ {\n\t\tkey:=a.Field(i)\n\t\tval:=b.Field(i)\n\t\tfmt.Println(key.Name,key.Type,\"|\",val)\n\t}\n}"
	fout.WriteString(s)

}

func CreatFile()  []string{
	natures:=[]string{}
	var str string=io.Rand()
	str=str+" "+Randresult(RndNature())+"\n"
	for i := 0; i < 10; i++ {
		nature:=RndNature()
		str=str+io.Rand()+" "+Randresult(nature)+"\n"
		natures=append(natures, nature)
	}
	Mywritefile(str)
	return natures
}

//后期读文件生成对象
func CreatObject(strs []string)  {
	userFile := "/home/wy/go/src/awesomeProject4/test/creatObject.go" //文件路径
	fout,err := os.Create(userFile) //根据路径创建File的内存地址
	defer fout.Close() //延迟关闭资源
	if err != nil{
		fmt.Println(userFile,err)
		return
	}
	var s string="package test\n"

	s=s+"func Mycreat() {\nt1:=new(Test)\n"
	for i := 0; i < 10; i++ {
		s=s+"t1.Set"+io.Capitalize(strs[i*2])+"("+strs[i*2+1]+")\n"
	}
	s=s+"\n}"
	fout.WriteString(s)
}