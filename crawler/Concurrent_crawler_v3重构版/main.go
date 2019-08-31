package main

import (
	"github.com/Altruiste1/go_learning/crawler/Concurrent_crawler/engine"
	"github.com/Altruiste1/go_learning/crawler/Concurrent_crawler/scheduler"
	"github.com/Altruiste1/go_learning/crawler/Concurrent_crawler/zhenai/parser"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/n/register?channelId=901388&subChannelId=15026",
	//	ParserFunc: parser.ParseCityList,
	//})
	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:10,
	}
	//e.Run(engine.Request{
	//		Url:        "https://www.zhenai.com/n/register?channelId=901388&subChannelId=15026",
	//		ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
				Url:        "http://www.zhenai.com/zhenghun/shanghai",
				ParserFunc: parser.ParseCity,
	})
}

