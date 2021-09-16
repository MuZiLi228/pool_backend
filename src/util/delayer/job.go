package delayer

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cast"
)

const (
	delayerStatusWait      = "WAIT"
	delayerStatusReady     = "READY"
	delayerStatusCancelled = "CANCELLED"
)

type topicFunc func(ctx context.Context, job *Job) error

var topicFuncMaps map[string]topicFunc

func init() {
	topicFuncMaps = make(map[string]topicFunc)
}

type Job struct {
	queue  string // 队列名
	Topic  string `json:"topic"`  // 任务类型
	ID     string `json:"ID"`     // 任务id
	Delay  int64  `json:"delay"`  // 延迟时间 ms
	TTR    int64  `json:"ttr"`    // 任务最大执行时间 ms
	Body   string `json:"body"`   // 任务详细信息
	status string `json:"status"` // 任务状态
}

func (j *Job) jobKeyName() string {
	return fmt.Sprintf("%s:queue:%s", rdcPrefix, j.ID)
}

func (j *Job) changeStatus(status string) error {
	return rdc.HSet(context.Background(), j.jobKeyName(), "status", status).Err()
}

func (j *Job) withFunc(ctx context.Context, fn func(ctx context.Context, job *Job) error) error {
	errc := make(chan error, 1)
	go func() {
		errc <- fn(ctx, j)
	}()

	select {
	case <-ctx.Done():
		// Wait for the goroutine to finish and send something.
		<-errc
		return ctx.Err()
	case err := <-errc:
		return err
	}
}

// withTopicFunc 执行指定TOPIC对应的方法
func (j *Job) withTopicFunc(ctx context.Context) error {
	if j.Topic == "" {
		return errors.New("job.topic 为空")
	}

	fn, exist := topicFuncMaps[j.queue+"-"+j.Topic]
	if !exist {
		return errors.New("该topic不存在对应的执行方法")
	}

	return j.withFunc(ctx, fn)
}

// addJobInfo 添加任务元信息
func addJobInfo(ctx context.Context, job *Job) error {
	if job.ID == "" {
		return errors.New("任务id为空")
	}
	return rdc.HMSet(ctx, job.jobKeyName(), map[string]interface{}{
		"topic":  job.Topic,
		"ID":     job.ID,
		"delay":  job.Delay,
		"ttr":    job.TTR,
		"body":   job.Body,
		"status": job.status,
		"queue":  job.queue,
	}).Err()
}

// cancledJob 将任务取消
func cancledJob(ctx context.Context, job *Job) error {
	return rdc.HSet(ctx, job.jobKeyName(), "status", delayerStatusCancelled).Err()
}

// delJobInfo 删除任务信息
func delJobInfo(ctx context.Context, job *Job) error {
	return rdc.Del(ctx, job.jobKeyName()).Err()
}

// getJobInfo 获取任务信息
func getJobInfo(ctx context.Context, job *Job) error {
	if job.ID == "" {
		return errors.New("job.ID 为空")
	}

	res, err := rdc.HGetAll(ctx, job.jobKeyName()).Result()
	if err != nil {
		return err
	}

	job.Topic = res["topic"]
	job.ID = res["ID"]
	job.Body = res["body"]
	job.Delay = cast.ToInt64(res["delay"])
	job.TTR = cast.ToInt64(res["ttr"])
	job.status = res["status"]
	job.queue = res["queue"]
	return nil
}
