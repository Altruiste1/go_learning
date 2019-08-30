package fetcher

import (
	"fmt"
	"net/http"
	"testing"
)

var urlList=[]string{
	"https://album.zhenai.com/u/1781349545",
	"https://album.zhenai.com/u/1157102175",
}

func TestFetch(t *testing.T) {
	for i,_:=range urlList{
		resp,err:=http.Get(urlList[i])
		fmt.Println(resp,err)
		all,err:=Fetch(urlList[i])
		if err!=nil{
			t.Errorf("Error :%s",err.Error())
			return
		}

		fmt.Println(all)
	}

}
