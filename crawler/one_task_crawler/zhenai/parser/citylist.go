package parser

import (
	"github.com/Altruiste1/go_learning/crawler/one_task_crawler/engine"
	"regexp"
)

//<a target="_blank" href="http://www.zhenai.com/zhenghun/beijing" data-v-28ffd398="">北京</a>
const cityListRe = `<a target="_blank" href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
const limit = 10

// 爬取城市列表
func ParseCityList(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	re := regexp.MustCompile(cityListRe)
	mathes := re.FindAllSubmatch(contents, -1)

	for i, m := range mathes {
		result.Items = append(
			result.Items, "City "+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
		if i == limit-1 {
			break
		}
	}

	return result
}
