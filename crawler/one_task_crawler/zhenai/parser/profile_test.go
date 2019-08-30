package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents,err:= ioutil.ReadFile("city.html")
	if err!=nil{
		t.Errorf("fetch failed:%s\n",err.Error())
	}

	var usernameList = []string{
		"靠近WO温暖你",
		"喜欢就好",
	}

	result:=ParseProfile(contents,usernameList[0])
	fmt.Printf("%v",result.Items)
}