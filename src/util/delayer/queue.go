package delayer

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cast"
)

type QueueExecutor func(i interface{})

// newQueue 创建队列
func newQueue(name string, scannIterval time.Duration, poolSize int, opts ...ants.Option) (*Queue, error) {
	pools, err := ants.NewPoolWithFunc(poolSize, topicFuncSelector, opts...)
	if err != nil {
		return new(Queue), err
	}

	return &Queue{
		Name:         name,
		ScanInterval: scannIterval,
		Pools:        pools,
	}, nil
}

func topicFuncSelector(i interface{}) {
	ctx := context.Background()

	// 获取任务id
	jobid := i.(string)

	job := &Job{
		ID: jobid,
	}

	// 获取任务信息
	if err := getJobInfo(ctx, job); err != nil {
		return
	}

	// 判断任务状态
	if job.status != delayerStatusReady {
		_ = job.withFunc(ctx, delJobInfo)
		return
	}

	// 设置任务超时时间
	ctx, _ = context.WithTimeout(ctx, time.Millisecond*time.Duration(job.TTR))

	// 执行任务
	_ = job.withTopicFunc(ctx)

	// 成功时，执行
	_ = job.withFunc(ctx, delJobInfo)
}

// Queue 延迟队列
type Queue struct {
	Name         string             // 队列名称
	ScanInterval time.Duration      // 扫描间隔
	Pools        *ants.PoolWithFunc // 协程池
}

func (q *Queue) delayKeyName() string {
	return fmt.Sprintf("%s:queue:%s:delay-list", rdcPrefix, q.Name)
}

func (q *Queue) readyKeyName() string {
	return fmt.Sprintf("%s:queue:%s:ready-list", rdcPrefix, q.Name)
}

func (q *Queue) startDelayScan() {
	ctx := context.Background()

	go func() {
		heartbeat := time.NewTicker(q.ScanInterval)
		defer heartbeat.Stop()

		for range heartbeat.C {
			jobs, _ := q.getMaturityDelayJob(ctx)

			for _, jobid := range jobs {
				job := &Job{ID: jobid}

				// 获取任务信息
				if err := getJobInfo(ctx, job); err != nil {
					return
				}

				if job.status != delayerStatusWait {
					// 状态不对，直接删除
					if err := job.withFunc(ctx, delJobInfo); err != nil {
						fmt.Println("删除任务失败，错误信息：", err.Error())
					}
					continue
				}

				_ = job.changeStatus(delayerStatusReady)
				_ = q.addReadyJob(ctx, jobid)
			}
		}
	}()
}

func (q *Queue) startReadyScan() {
	ctx := context.Background()
	go func() {
		heartbeat := time.NewTicker(q.ScanInterval)
		defer heartbeat.Stop()

		for range heartbeat.C {
			jobid, err := q.getReadyJob(ctx)
			if err != nil || jobid == "" {
				continue
			}

			_ = q.Pools.Invoke(jobid)
		}
	}()
}

// addDelayJob 添加排期任务
func (q *Queue) addDelayJob(ctx context.Context, jobID string, execUnix int64) error {
	return rdc.ZAdd(ctx, q.delayKeyName(), &redis.Z{
		Score:  cast.ToFloat64(execUnix),
		Member: jobID,
	}).Err()
}

// getMaturityDelayJob 获取指定最大数量的已经到期的排期任务
func (q *Queue) getMaturityDelayJob(ctx context.Context) ([]string, error) {
	maxUnix := cast.ToString(time.Now().Unix())

	pipeline := rdc.Pipeline()
	pipeline.ZRangeByScore(ctx, q.delayKeyName(), &redis.ZRangeBy{
		Min: "0",
		Max: maxUnix,
	})
	pipeline.ZRemRangeByScore(ctx, q.delayKeyName(), "0", maxUnix)
	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return cmders[0].(*redis.StringSliceCmd).Result()
}

// addReadyJob 添加待执行任务
func (q *Queue) addReadyJob(ctx context.Context, jobid string) error {
	return rdc.LPush(ctx, q.readyKeyName(), jobid).Err()
}

// getReadyJob 获取一个待执行任务
func (q *Queue) getReadyJob(ctx context.Context) (string, error) {
	return rdc.RPop(ctx, q.readyKeyName()).Result()
}

// StartQueue 启动队列
func (q *Queue) StartQueue() {
	q.startDelayScan()
	q.startReadyScan()
}

// AddJob 添加一个任务
func (q *Queue) AddJob(ctx context.Context, topic string, delayer int64, jobInfo string) (string, error) {
	now := time.Now()

	// 存储任务初始状态和指定时间
	job := &Job{
		queue:  q.Name,
		Topic:  topic,
		ID:     snowflakeNode.Generate().String(),
		Delay:  delayer,
		TTR:    30000,
		Body:   jobInfo,
		status: delayerStatusWait,
	}

	// 保存任务信息
	if err := job.withFunc(ctx, addJobInfo); err != nil {
		return "", err
	}

	// 推任务进排期队列
	return job.ID, q.addDelayJob(ctx, job.ID, now.Add(time.Millisecond*time.Duration(job.Delay)).Unix())
}

func (q *Queue) CancleJob(ctx context.Context, jobid string) error {
	job := &Job{
		queue: q.Name,
		ID:    jobid,
	}

	if err := job.withFunc(ctx, cancledJob); err != nil {
		return err
	}

	return nil
}

// AddReadyJob 添加一个待执行任务
func (q *Queue) AddReadyJob(ctx context.Context, topic string, delayer int64, jobInfo string) (string, error) {
	job := &Job{
		queue:  q.Name,
		Topic:  topic,
		ID:     snowflakeNode.Generate().String(),
		Delay:  delayer,
		TTR:    30000,
		Body:   jobInfo,
		status: delayerStatusWait,
	}

	// 保存任务信息
	if err := job.withFunc(ctx, addJobInfo); err != nil {
		return "", err
	}

	// 推任务进排期队列
	return job.ID, q.addReadyJob(ctx, job.ID)
}

func (q *Queue) AddTopicWithFunc(topic string, fn topicFunc) {
	topicFuncMaps[q.Name+"-"+topic] = fn
}
