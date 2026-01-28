package worker_pool

import (
	"context"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

type Task interface {
	Do(ctx context.Context) error
}

////////////////////////////////////////////////////
////////////////////////////////////////////////////
//////////////////////Worker////////////////////////
////////////////////////////////////////////////////
////////////////////////////////////////////////////

type Worker struct {
	task     Task
	schedule cron.Schedule
}

func NewWorker(task Task, cronTime string) Worker {
	schedule, err := cron.ParseStandard(cronTime)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при парсинге КРОНА %s", err.Error()))
		return Worker{}
	}

	return Worker{
		task:     task,
		schedule: schedule,
	}
}

func (w Worker) Start(ctx context.Context) {
LOOP:
	for {
		now := time.Now().UTC()
		select {
		case <-ctx.Done():
			break LOOP
		case <-time.After(w.nextScheduledTime(now).Sub(now)):
			if err := w.task.Do(ctx); err != nil {
				log.Printf("task process error: %v", err) // todo: log
			}
		}
	}
}

func (w Worker) nextScheduledTime(now time.Time) time.Time {
	return w.schedule.Next(now)
}

////////////////////////////////////////////////////
////////////////////////////////////////////////////
////////////////////WorkerPool//////////////////////
////////////////////////////////////////////////////
////////////////////////////////////////////////////

type WorkerPool struct {
	workers []Worker
	limit   int
}

func NewWorkerPool(workers []Worker) WorkerPool {
	return WorkerPool{
		workers: workers,
	}
}

func (wp WorkerPool) Run(ctx context.Context) {
	for _, w := range wp.workers {
		go w.Start(ctx)
	}
}
