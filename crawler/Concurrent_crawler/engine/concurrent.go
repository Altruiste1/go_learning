package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
//	WorkerReady(chan Request)
//	Run()
}

func (e *ConcurrentEngine)Run(seeds ...Request){
	in :=make(chan Request)
	out :=make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	//channel里读数据
	for i:=0;i<e.WorkerCount;i++{
		createWorker(in,out)
	}

	// 往channel写入数据,往scheduler发任务
	for _,r:=range seeds{
		// chan<-
		e.Scheduler.Submit(r)
	}

	for {
		result:=<-out
		for _,item :=range result.Items{
			log.Printf("Got item :%v",item)
		}

		for _,request:=range result.Requests{
			e.Scheduler.Submit(request)
		}
	}

}


func createWorker(
	in chan Request,out chan ParseResult){
		go func(){
			for{
				// tell scheduler
				request:= <- in
				result,err:= worker(request)
				if err!=nil{
					log.Printf("worker error :%s",err.Error())
					continue
				}
				out <-result
			}
		}()
}