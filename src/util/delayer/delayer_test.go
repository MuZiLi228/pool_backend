package delayer

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestNewClient(t *testing.T) {
	rdc := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.6:6379",
		Password: "redis_pass_2020",
		DB:       10,
	})

	if rdc.Ping(context.Background()).Err() != nil {
		t.Error("创建连接失败")
	}

	client, err := NewClient(
		WithRedis(rdc),
		WithSnowflakeNode(1),
	)
	if err != nil {
		t.Error(err.Error())
	}

	_ = client.NewQueue("oreder_canal", time.Second*10, 5)
	testQueue := client.UseQueue("oreder_canal")
	testQueue.AddTopicWithFunc("oreder_canal", func(ctx context.Context, job *Job) error {
		fmt.Printf("jobid:%s | 执行时间点：%s \n", job.ID, time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})
	testQueue.StartQueue()

	jobid1, err := testQueue.AddJob(context.Background(), "oreder_canal", 1000*10, "111")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("jobid:%s | 添加时间点：%s \n", jobid1, time.Now().Format("2006-01-02 15:04:05"))

	time.Sleep(time.Second * 10)
}
