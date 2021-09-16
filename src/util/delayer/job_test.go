package delayer

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestJob_withFunc(t *testing.T) {
	rdc = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.6:6379",
		Password: "redis_pass_2020",
		DB:       10,
	})
	job := &Job{
		Topic: "oreder_canal",
		ID:    "1430025868733845504",
		Delay: 100,
		TTR:   30,
		Body:  "Hello World!",
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*6)

	if err := job.withFunc(ctx, addJobInfo); err != nil {
		t.Error(err.Error())
	}

	fmt.Println("OK")
}
