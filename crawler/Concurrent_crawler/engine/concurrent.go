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
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine)Run(seeds ...Request){
	//in :=make(chan Request)
	out :=make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	//channel里读数据
	for i:=0;i<e.WorkerCount;i++{
		createWorker(out,e.Scheduler)
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
	out chan ParseResult,s Scheduler){
		in:=make(chan Request)
		go func(){
			for{
				// tell scheduler ready
				s.WorkerReady(in)
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