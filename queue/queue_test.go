package queue

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

type TaskExample struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func TestQueue_Setup_Queue(t *testing.T) {
	fmt.Println("opening redis connection...")

	c := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	defer c.Close()

	if _, err := c.Ping().Result(); err != nil {
		t.Fatal("unable to connect to redis", err)
	}

	New(c)
	go SetAsSubscriber()

	time.AfterFunc(time.Second, func() {
		fmt.Println("enqueing something")

		err := Enqueue(TaskEmail, TaskExample{
			From:    "me@testing.com",
			To:      "unit@test.com",
			Subject: "unit test",
			Body:    "<h1>unit test</h1>",
		})

		if err != nil {
			t.Fatal("unable to publish to channel", err)
		}
	})

	time.Sleep(3 * time.Second)
}