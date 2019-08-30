package engine

import (
	"github.com/Altruiste1/go_learning/crawler/one_task_crawler/fetcher"
	"github.com/labstack/gommon/log"
)
func Run(seeds ...Request){
	var requests []Request
	for _,r:=range seeds{
		requests = append(requests,r)
	}

	for len(requests)>0{
		r :=requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s\n",r.Url)
		body ,err:=fetcher.Fetch(r.Url)
		if err!=nil{
			log.Printf("Fetcher:error "+
				"fetching url %s,%v",r.Url,err)
			continue
		}

		parestResult:=r.ParserFunc(body)
		requests =append(requests,
			parestResult.Requests...)
		for _,item:=range parestResult.Items{
			log.Printf("Got item %s",item)
		}
	}
}
