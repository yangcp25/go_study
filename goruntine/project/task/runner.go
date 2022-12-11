package main

import (
	"context"
	"fmt"
	"os"
	"task/runner"
	"time"
)

func main() {
	fmt.Print("来了的\n")
	ctx, cancelFun := context.WithCancel(context.Background())
	task := runner.New(3*time.Second, cancelFun)

	task.Add(taskFunc(), taskFunc(), taskFunc(), taskFunc(), taskFunc(), taskFunc(), taskFunc())
	err := task.Start()
	if err != nil {
		switch err {
		case runner.ErrorTimeOut:
			fmt.Printf("%v", runner.ErrorTimeOut)
			os.Exit(1)
		case runner.InterruptOut:
			fmt.Printf("%v", runner.InterruptOut)
			os.Exit(2)
		}
	}
	<-ctx.Done()
	fmt.Println("\nfinished")
}

func taskFunc() func(int, func()) {
	return func(i int, funcDone func()) {
		date := getDateTime()
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Printf(date+";任务：%d顺利执行！\n", i)
		if i == 6 {
			funcDone()
		}
	}
}

func getDateTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}
