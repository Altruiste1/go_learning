package parser

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

)

func TestParseCityList(t *testing.T) {
	t.SkipNow()
	contents,err:= ioutil.ReadFile("text.html")
	if err!=nil{
		t.Errorf("Fetch failed:%s\n",err.Error())
	}

	expectedUrls:=[]string{
		"http://www.zhenai.com/zhenghun/beijing",
	}

	expectedCities:=[]string{
		"北京",
	}
	result :=ParseCityList(contents)
	fmt.Printf("result :%d",len(result.Items))
	for i,url:=range expectedUrls{
		if result.Requests[i].Url!=url{
			t.Errorf("expect url %s;but "+
				"was %s",
				url,result.Requests[i].Url)
		}
		fmt.Printf("%s",url)
	}


	for i,city:= range expectedCities{
		item,ok:=result.Items[i].([]uint8)
		if !ok{
			fmt.Println("not string,type is",reflect.ValueOf(result.Items[i]),item)

		}
		if string(item)!=city{
			t.Errorf("expect city :%s;but "+
				"was %s",
				city,string(item))
		}
		fmt.Println(city)

	}

}
