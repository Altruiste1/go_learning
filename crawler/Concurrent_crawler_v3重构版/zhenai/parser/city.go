package parser

import (
	"github.com/Altruiste1/go_learning/crawler/Concurrent_crawler/engine"
	"regexp"
)

//var cityRe = `<th><a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a></th>`
var(
	cityRe = regexp.MustCompile(
	`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`,
	)

    cityUrlRe = regexp.MustCompile(
    	`<a href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`,
    )
)

//var cityRe=`<th><a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a></th>`
// 获取某个city中人名信息
func ParseCity(contents []byte) engine.ParseResult {
	//re := regexp.MustCompile(cityRe)
	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(m[2]))
			},
		})
	}
	matches = cityUrlRe.FindAllSubmatch(contents,-1)
	for _,m:=range matches{
		result.Requests =append(result.Requests,
			engine.Request{
				Url:string(m[1]),
				ParserFunc:ParseCity,
			})
	}
	return result
}

func ParsePerson(contents []byte) engine.ParseResult {

	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(m[2]))
			},
		})
	}


	return result
}
