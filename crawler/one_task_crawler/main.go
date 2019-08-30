package main

import (
	"github.com/Altruiste1/go_learning/crawler/one_task_crawler/engine"
	"github.com/Altruiste1/go_learning/crawler/one_task_crawler/zhenai/parser"
)


func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/n/register?channelId=901388&subChannelId=15026",
		ParserFunc: parser.ParseCity,
	})
}

