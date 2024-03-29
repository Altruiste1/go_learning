package scheduler

import "github.com/Altruiste1/go_learning/crawler/Concurrent_crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueueScheduler)Submit(r engine.Request){
	s.requestChan <-r
}

func (s *QueueScheduler)WorkerReady(
	w chan engine.Request){
		s.workerChan<-w
}

func (s *QueueScheduler)ConfigureMasterWorkerChan(){
	panic("implement me")
}

func (s *QueueScheduler)Run(){
	s.workerChan =make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func(){
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0 &&
				len(workerQ)>0{
				activeWorker =workerQ[0]
				activeRequest =requestQ[0]
			}

			select {
			case r:=<-s.requestChan:
				requestQ = append(requestQ,r)
				// send r to a ?worker
			case w:=<-s.workerChan:
				// send ?next request to w
				workerQ =append(workerQ,w)
				case activeWorker <-activeRequest:
					workerQ =workerQ[1:]
					requestQ =requestQ[1:]
			}
		}
	}()
}