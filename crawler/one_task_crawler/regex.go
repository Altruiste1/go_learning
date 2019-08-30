package main

import (
	"fmt"
	"regexp"
)

func main(){
	//re:=regexp.MustCompile("ccmouse@gmail.com")
	texts:= []struct {
		text string
	}{
		{"My email is ccmouse@gmail.com"},
		{"ccmouse@gmail.com@abc.com"},
		{`ccmouse@gmail.com@abc.com
  c@qq.com
  dd@dad.com`},
	}

	test(texts,GetEmail)
	test(texts,GetAllEmail)
	test(texts,GainNeedMessage)
}

// 在一个字符串中获得一个满足xxx@xxx的字符串
func GetEmail(text string){
		re:=regexp.MustCompile(`[a-zA-Z0-9]+@.+\..`)
		match:=re.FindString(text)
		fmt.Println(match)
}

// 获得字符串中符合的所有子串
func GetAllEmail(text string){
		re:=regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9\.]+\.[a-zA-Z0-9]+`)
		match:=re.FindAllString(text,-1)
		fmt.Println(match)

}

// 将爬取的email分解:例如aaa@qq.com 分解成  aaa  qq  com
// ()可以将所需要的数据保留，再通过FindAllStringSubmatch可以获得()预留的字符串
func GainNeedMessage(text string){
	re:=regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+)\.([a-zA-Z0-9]+)`)
	match:=re.FindAllStringSubmatch(text,-1)
	fmt.Println("all:",match)
	for _,m:=range match{
		fmt.Println("sub:",m)
	}
}

type Reg   func(string)
type textList []struct{text string}

func test(texts []struct{text string},reg func(string)){
	for _,v:=range texts{
		reg(v.text)
	}
}

