package main

import (
	"fmt"
	"os"
	"task/runner"
	"time"
)

func main() {
	fmt.Print("来了的\n")
	task := runner.New(3 * time.Second)

	task.Add(taskFunc(), taskFunc(), taskFunc())
	err := task.Start()
	//fmt.Printf("%v", err)
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
}

func taskFunc() func(int) {
	return func(i int) {
		fmt.Printf("任务：%d顺利执行！\n", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
