package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	//ConfigureMasterWorkerChan(chan Request)
	WorkerChan()chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine)Run(seeds ...Request){
	//in :=make(chan Request)
	out :=make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	//channel里读数据
	for i:=0;i<e.WorkerCount;i++{
		createWorker(e.Scheduler.WorkerChan(),
			out,e.Scheduler)
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


func createWorker(in chan Request,
	out chan ParseResult,ready ReadyNotifier){
		//in:=make(chan Request)
		go func(){
			for{
				// tell scheduler ready
				ready.WorkerReady(in)
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