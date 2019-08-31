package parser

import (
	"github.com/Altruiste1/go_learning/crawler/Concurrent_crawler/engine"
	"github.com/Altruiste1/go_learning/crawler/Concurrent_crawler/model"
	"regexp"
	"strconv"
)

//<td width="180"><span class="grayL">年龄：</span>24</td>
// [\d]匹配数字
var ageRe = regexp.MustCompile(
	`<span class="grayL">年龄：</span>([\d]+)</td>`,
)

var marriageRe = regexp.MustCompile(
	`<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`,
)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	// 有错误不做处理
	if err == nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}
