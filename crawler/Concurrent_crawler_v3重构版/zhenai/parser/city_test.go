package parser

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"io/ioutil"

	//"io/ioutil"
	"reflect"
	"testing"
)

var usernameList = []string{
	"靠近WO温暖你",
	"喜欢就好",
}

var expectedUrls = []string{
	"http://album.zhenai.com/u/1157102175",
}

func TestParseCity(t *testing.T) {
	t.SkipNow()
	//	str:="http://www.zhenai.com/zhenghun/beijing"
	contents, err := ioutil.ReadFile("city.html")
	if err != nil {
		t.Errorf("fetch failed:%s\n", err.Error())
	}

	result := ParseCity(contents)
	for i, _ := range usernameList {

		name := getName(result.Items[i])
		if usernameList[i] != name {
			t.Errorf("爬取名字失败:期待%s ,实际：%s",
				usernameList[i], name)
		}
		fmt.Printf("test:%d,期待name:%s，爬取name:%s\n", i, usernameList[i], name)
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expect url %s;but "+
				"was %s",
				url, result.Requests[i].Url)
		}
		fmt.Printf("%s", url)
	}
}

func getName(uname interface{}) string {
	switch uname.(type) {
	case string:
		return uname.(string)
	case []uint8:
		return string(uname.([]uint8))
	default:
		log.Printf("类型为:%s,处理失败，返回空字符串", reflect.TypeOf(uname))
		return ""
	}
}
