package scheduler

import (
	"errors"
	"github.com/gorhill/cronexpr"
	"time"
)

type CommandFunction func (arg ...interface{})


type CronJob struct {
	expression *cronexpr.Expression
	next time.Time
	handle CommandFunction
	name string
	arg []interface{}
	ok chan bool
}

func InitCronJob() *CronJob{
	var (
		c chan bool
	)
	c = make(chan bool,1)
	c <- true
	return &CronJob{ok:c}
}

func (job *CronJob) GetChan() chan bool {
	return job.ok
}
func (job *CronJob) SetChan(t bool) {
	job.ok <- t
}

func (job *CronJob) SetName(name string) *CronJob {
	job.name = name
	return job
}

func (job *CronJob) GetName() string {
	return job.name
}

func (job *CronJob) SetArgs (args ...interface{}) *CronJob{
	job.arg = args
	return job
}

func (job *CronJob) GetArgs() []interface{} {
	return job.arg
}

func (job *CronJob) SetExpression(expression string, now time.Time) *CronJob {
	var (
		expire *cronexpr.Expression
		err error
	)

	if expire, err = cronexpr.Parse(expression); err != nil {
		panic(err)
	}
	job.expression = expire
	job.next = expire.Next(now)
	return job
}

func (job *CronJob) GetExpression() *cronexpr.Expression {
	return job.expression
}


func (job *CronJob) SetHandle(handle CommandFunction) *CronJob{
	job.handle = handle
	return job
}

func (job *CronJob) GetHandle() CommandFunction {
	return job.handle
}

func (job *CronJob) GetNext() time.Time {
	return job.next
}

func (job *CronJob) SetNext(next time.Time) {
	job.next = next
}







type Scheduler struct {

	scheduler map[string]*CronJob
}


func  GetScheduler() *Scheduler {
	return &Scheduler{scheduler:make(map[string]*CronJob)}
}

func (scheduler *Scheduler) GetJobManager() map[string]*CronJob {
	return scheduler.scheduler
}

func (scheduler *Scheduler) PushJob(job... *CronJob) error {
	var (
		name string
		j *CronJob
		ok bool
	)
	for _, j = range job {
		name = j.name
		if _, ok = scheduler.scheduler[name]; ok {
			return errors.New("存在同名的job！")
		}
		scheduler.scheduler[name] = j
	}
	return nil
}



