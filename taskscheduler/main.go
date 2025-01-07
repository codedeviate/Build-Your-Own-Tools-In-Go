// main.go
package main

import (
	"fmt"
	"time"

	"github.com/username/taskscheduler/lib"
)

func main() {
	scheduler := lib.NewScheduler()

	// Example task: Print a message every 2 seconds
	scheduler.AddTask(lib.Task{
		Name:     "PrintMessage",
		Interval: 2 * time.Second,
		Action: func() {
			fmt.Println("Task executed: PrintMessage")
		},
	})

	// Example task: Print a different message every 5 seconds
	scheduler.AddTask(lib.Task{
		Name:     "PrintAnotherMessage",
		Interval: 5 * time.Second,
		Action: func() {
			fmt.Println("Task executed: PrintAnotherMessage")
		},
	})

	// Example task: Print a message after a 10-second delay
	scheduler.AddTask(lib.Task{
		Name:  "DelayedMessage",
		Delay: 10 * time.Second,
		Action: func() {
			fmt.Println("Task executed: DelayedMessage")
		},
	})

	// Example task: Print a message based on a cron expression
	scheduler.AddTask(lib.Task{
		Name:     "CronMessage",
		CronExpr: "*/1 * * * *", // Every minute
		Action: func() {
			fmt.Println("Task executed: CronMessage")
		},
	})

	// Example task: Print a message if a condition is met
	scheduler.AddTask(lib.Task{
		Name: "ConditionalMessage",
		Condition: func() bool {
			return time.Now().Second()%2 == 0
		},
		Action: func() {
			fmt.Println("Task executed: ConditionalMessage")
		},
	})

	// Start the scheduler
	scheduler.Start()

	// Keep the main function running
	select {}
}
