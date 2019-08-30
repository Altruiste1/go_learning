package main

import (
	"bufio"
	"fmt"
	"github.com/Altruiste1/go_learning/crawler/one_task_crawler/engine"
	"github.com/Altruiste1/go_learning/crawler/one_task_crawler/zhenai/parser"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main()  {
	engine.Run(engine.Request{
		Url:"https://www.zhenai.com/n/register?channelId=901388&subChannelId=15026",
		ParserFunc:parser.ParseCityList,
	})
}

func judgeHtmlEncoding(
	reader io.Reader)  encoding.Encoding{
		bytes,err:=bufio.NewReader(reader).Peek(1024)
		if err!=nil{
			panic(err)
		}
		e,_,_:= charset.DetermineEncoding(bytes,"")
		return e

}

func printCityList(contents []byte){
	str:=`<a target="_blank" href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`
	re:=regexp.MustCompile(str)
	mathes:=re.FindAll(contents,-1)
	for _,m:=range mathes{
		fmt.Printf("%s\n",m)
	}

	fmt.Println("Matches found:",len(mathes))
}

// 城市列表解析器
func printCityListSub(contents []byte){
	//str:=`<a target="_blank" href="(http://www.zhenai.com/zhenghun/[0-9a-z])+"[^>]*>([^<])+</a>`
	log.Printf("Get Data:%s\n",contents)
	re:=regexp.MustCompile(`<a target="_blank" href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	mathes:=re.FindAllSubmatch(contents,-1)
	for _,m:=range mathes{
		fmt.Printf("City:%s, URL: %s\n",m[2],m[1])
		//fmt.Printf("%s",m)
	}

	fmt.Println("Matches found:",len(mathes))
}

func fetcher(){
	resp,err:=http.Get("https://www.zhenai.com/n/register?channelId=901388&subChannelId=15026")
	if err!=nil{
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		fmt.Println("err code:",resp.StatusCode)
		return
	}

	e:=judgeHtmlEncoding(resp.Body)
	utf8Reader:=transform.NewReader(resp.Body,e.NewDecoder())
	if resp.StatusCode ==http.StatusOK{
		all,err:=ioutil.ReadAll(utf8Reader)
		if err!=nil{
			panic(err)
		}

		//fmt.Printf("%s\n",all)
		printCityListSub(all)
		//printCityList(all)

	}
}