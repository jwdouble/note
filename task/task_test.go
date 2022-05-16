package task

import (
	"context"
	"fmt"
	"testing"
	"time"

	"arctron.lib/task"
)

func TestTask(t *testing.T) {
	ctx := context.Background()
	work := task.New(ctx)
	work.Add(func() error {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			//work1 := task.New(ctx)
			work.Add(func() error {
				for l := 0; l < 9; l++ {
					fmt.Printf("%d", l)
				}
				fmt.Println("")
				return nil
			})
			//work.Run()
		}
		return nil
	})
	work.Run()
}

func Test_mess(t *testing.T) {
	go func() {
		ta := task.New(context.Background())
		ta.Add(func() error {
			fmt.Println("1")
			return nil
		})
		ta.Run()
	}()
	go func() {
		ta := task.New(context.Background())
		ta.Add(func() error {
			fmt.Println("2")
			return nil
		})
		ta.Run()
	}()

	time.Sleep(1 * time.Second)
}
