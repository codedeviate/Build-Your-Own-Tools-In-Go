// scheduler.go
package lib

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// Task represents a scheduled task.
type Task struct {
    Name       string
    Interval   time.Duration
    CronExpr   string
    Delay      time.Duration
    Condition  func() bool
    Action     func()
    LastStatus string
}

// Scheduler manages and runs tasks.
type Scheduler struct {
    tasks []Task
    cron  *cron.Cron
}

// NewScheduler creates a new Scheduler.
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(),
    }
}

// AddTask adds a new task to the scheduler.
func (s *Scheduler) AddTask(task Task) {
    s.tasks = append(s.tasks, task)
    if task.CronExpr != "" {
        s.cron.AddFunc(task.CronExpr, func() {
            s.executeTask(task)
        })
    } else if task.Interval > 0 {
        go func(t Task) {
            ticker := time.NewTicker(t.Interval)
            for {
                select {
                case <-ticker.C:
                    s.executeTask(t)
                }
            }
        }(task)
    } else if task.Delay > 0 {
        go func(t Task) {
            time.Sleep(t.Delay)
            s.executeTask(t)
        }(task)
    }
}

// Start runs the scheduler and executes tasks at their specified intervals.
func (s *Scheduler) Start() {
    s.cron.Start()
    for _, task := range s.tasks {
        if task.CronExpr == "" && task.Interval == 0 && task.Delay == 0 {
            go s.executeTask(task)
        }
    }
}

// Stop stops the scheduler.
func (s *Scheduler) Stop() {
    s.cron.Stop()
    fmt.Println("Scheduler stopped")
}

// executeTask executes a task and logs its status.
func (s *Scheduler) executeTask(task Task) {
    if task.Condition != nil && !task.Condition() {
        log.Printf("Task %s skipped due to condition", task.Name)
        return
    }
    log.Printf("Task %s started", task.Name)
    task.Action()
    task.LastStatus = "completed"
    log.Printf("Task %s completed", task.Name)
}
