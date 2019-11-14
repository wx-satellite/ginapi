package main

import (
	"fmt"
	"gin/command/handler"
	"gin/command/scheduler"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {

	var (
		jobManager *scheduler.Scheduler
		now        time.Time
		err        error
		errors     chan interface{}
	)

	jobManager = scheduler.GetScheduler()

	now = time.Now()

	if err = jobManager.PushJob(
		scheduler.
			InitCronJob().
			SetExpression("* * * * *", now).
			SetName("每五秒输出hello").
			SetHandle(handler.SayHello).
			SetArgs(33333),
		scheduler.
			InitCronJob().
			SetExpression("* * * * *", now).
			SetName("每十秒输出hello").
			SetHandle(handler.SayHello).
			SetArgs(66666),
		scheduler.
			InitCronJob().
			SetExpression("* * * * *", now).
			SetName("每20分钟执行一次").
			SetHandle(handler.SayHello).
			SetArgs(888888),
	); err != nil {
		fmt.Println(err)
		return
	}

	// 调度
	go func() {
		var (
			name       string
			job        *scheduler.CronJob
			expression *cronexpr.Expression
			next       time.Time
			now        time.Time
		)
		defer func() {
			if err := recover(); err != nil {
				errors <- err
			}
		}()
		for {
			now = time.Now()
			for name, job = range jobManager.GetJobManager() {
				expression = job.GetExpression()
				next = job.GetNext()
				if next.Before(now) || next.Equal(now) {
					select {
					case <-job.GetChan():
						go func(job *scheduler.CronJob, name string) {
							defer func() {
								if err := recover(); err != nil {
									fmt.Println("任务：", name, "执行失败了！")
								}
							}()
							job.GetHandle()(job.GetArgs()...)
							job.SetNext(expression.Next(now))
							job.SetChan(true)
						}(job,name)
					default:

					}
				}

				select {
				case <-time.NewTimer(100 * time.Microsecond).C:
				}

			}
		}
	}()

	select {
	case err := <-errors:
		fmt.Println(err)
	}
}
