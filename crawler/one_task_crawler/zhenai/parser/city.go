package parser

import (
	"github.com/Altruiste1/go_learning/crawler/one_task_crawler/engine"
	"regexp"
)

//var cityRe = `<th><a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a></th>`
var cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`
var pmessageRe = `<a href="(http://album.zhenai.com/u/[0-9]+" [^>]*>([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>[^>]*[^<]+</td> <td><span class="grayL">居住地：</span>北京</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>([0-9]+)</td> <!----> <td><span class="grayL">月&nbsp;&nbsp;&nbsp;薪：</span>5001-8000元</td></tr> <tr><td width="180"><span class="grayL">婚况：</span>未婚</td> <td width="180"><span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>173</td></tr></tbody>`

//var cityRe=`<th><a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a></th>`
// 获取某个city中人名信息
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
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

func ParsePerson(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
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
}
