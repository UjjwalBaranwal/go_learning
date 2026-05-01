package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID          string
	Name        string
	IsRecurring bool
	stopChan    chan struct{}
	wg          sync.WaitGroup
}
type Scheduler struct {
	tasks    map[string]*Task
	mu       sync.Mutex
	globalWg sync.WaitGroup
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(map[string]*Task),
	}
}

func generateTaskID() string {
	return fmt.Sprintf("task-%d", time.Now().UnixNano())
}

// ScheduleOnce schedules a one-off task to run after a specified delay. It generates a unique task ID, logs the scheduling action, and uses time.AfterFunc to execute the provided task function after the delay. The global wait group is used to track the completion of the task for graceful shutdown purposes.
func (s *Scheduler) ScheduleOnce(name string, delay time.Duration, taskFunc func()) string {
	taskID := generateTaskID()
	fmt.Printf("[%s] SCHEDULER: Scheduling one-off task '%s' (ID: %s) to run after %s\n",
		time.Now().Format("15:04:05.000"), name, taskID, delay)
	s.globalWg.Add(1)
	time.AfterFunc(delay, func() {
		defer s.globalWg.Done()
		fmt.Printf("[%s] TASK '%s' (ID: %s): Executing one-off action.\n",
			time.Now().Format("15:04:05.000"), name, taskID)
		taskFunc()
		fmt.Printf("[%s] TASK '%s' (ID: %s): Finished one-off action.\n",
			time.Now().Format("15:04:05.000"), name, taskID)
	})
	return taskID
}

// ScheduleInterval schedules a recurring task that starts after an initial delay and then executes at regular intervals. It creates a Task struct to manage the task's state and uses a goroutine to handle the timing and execution of the task function. The function listens for stop signals to allow for graceful termination of the recurring task when needed.

func (s *Scheduler) ScheduleInterval(name string, initialDelay time.Duration, interval time.Duration, taskFunc func()) (string, error) {
	if interval <= 0 {
		return "", fmt.Errorf("interval must be positive , got %v", interval)
	}
	taskID := generateTaskID()
	task := &Task{
		ID:          taskID,
		Name:        name,
		IsRecurring: true,
		stopChan:    make(chan struct{}),
	}
	s.mu.Lock()
	s.tasks[taskID] = task
	s.mu.Unlock()
	fmt.Printf("[%s] SCHEDULER: Scheduling interval task '%s' (ID: %s) with initial delay %s, interval %s\n",
		time.Now().Format("15:04:05.000"), name, taskID, initialDelay, interval)
	s.globalWg.Add(1)
	task.wg.Add(1)
	go func() {
		defer s.globalWg.Done()
		defer task.wg.Done()
		fmt.Printf("[%s] TASK '%s' (ID: %s): Goroutine started. Initial delay: %s.\n",
			time.Now().Format("15:04:05.000"), task.Name, task.ID, initialDelay)
		initTimer := time.NewTimer(initialDelay)
		select {
		case <-initTimer.C:
			// allow to flow down
		case <-task.stopChan:
			initTimer.Stop()
			fmt.Printf("[%s] TASK '%s' (ID: %s): Stopped before initial run.\n",
				time.Now().Format("15:04:05.000"), task.Name, task.ID)
			return
		}
		fmt.Printf("[%s] TASK '%s' (ID: %s): Executing first action.\n",
			time.Now().Format("15:04:05.000"), task.Name, task.ID)
		taskFunc()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		fmt.Printf("[%s] TASK '%s' (ID: %s): Ticker started, interval %s.\n",
			time.Now().Format("15:04:05.000"), task.Name, task.ID, interval)
		for {
			select {
			case t := <-ticker.C:
				fmt.Printf("[%s] TASK '%s' (ID: %s): Executing recurring action at %v.\n",
					time.Now().Format("15:04:05.000"), task.Name, task.ID, t.Format("15:04:05.000"))
				taskFunc()
			case <-task.stopChan:
				fmt.Printf("[%s] TASK '%s' (ID: %s): Received stop signal. Exiting goroutine.\n",
					time.Now().Format("15:04:05.000"), task.Name, task.ID)
				return
			}

		}
	}()
	return taskID, nil
}
func (s *Scheduler) stopTask(taskID string) bool {
	s.mu.Lock()
	task, exists := s.tasks[taskID]
	if !exists || !task.IsRecurring {
		s.mu.Unlock()
		fmt.Printf("[%s] SCHEDULER: Task ID '%s' not found or not a stoppable recurring task.\n",
			time.Now().Format("15:04:05.000"), taskID)
		return false
	}
	delete(s.tasks, taskID)
	s.mu.Unlock()
	fmt.Printf("[%s] SCHEDULER: Sending stop signal to task '%s' (ID: %s).\n",
		time.Now().Format("15:04:05.000"), task.Name, taskID)
	close(task.stopChan) // am assumption that this would be closed once and only once, so no need to worry about multiple close panic here
	task.wg.Wait()       // wait for the task goroutine to finish
	fmt.Printf("[%s] SCHEDULER: Task '%s' (ID: %s) has been stopped and cleaned up.\n",
		time.Now().Format("15:04:05.000"), task.Name, taskID)
	return true
}

// StopAllTasks gracefully stops all recurring tasks by sending stop signals to each task's stop channel and waiting for their goroutines to finish. It first collects the IDs of all tasks, then iterates through them to stop each one. After initiating the shutdown of all tasks, it waits for the global wait group to ensure that all task goroutines have completed before confirming that the scheduler has fully stopped.
func (s *Scheduler) StopAllTasks() {
	fmt.Printf("[%s] SCHEDULER: Initiating shutdown of all tasks...\n", time.Now().Format("15:04:05.000"))
	s.mu.Lock()
	var taskIDs []string
	for id := range s.tasks {
		taskIDs = append(taskIDs, id)
	}
	s.mu.Unlock()
	for _, id := range taskIDs {
		s.stopTask(id)
	}
	fmt.Printf("[%s] SCHEDULER: Waiting for all task goroutines to complete...\n", time.Now().Format("15:04:05.000"))
	s.globalWg.Wait()

	fmt.Printf("[%s] SCHEDULER: All tasks shut down. Scheduler stopped.\n", time.Now().Format("15:04:05.000"))
}

// AutoStopAll starts a goroutine that waits for a specified duration before automatically stopping all tasks in the scheduler. It uses a time.Timer to trigger the shutdown process after the given duration, allowing for a graceful exit of all scheduled tasks without requiring manual intervention.
func (s *Scheduler) AutoStopAll(after time.Duration) {
	go func() {
		timer := time.NewTimer(after)
		defer timer.Stop()
		<-timer.C
		fmt.Printf("[%s] SCHEDULER: AutoStopAll triggered after %s.\n",
			time.Now().Format("15:04:05.000"), after)
		s.StopAllTasks()
	}()
}
func main() {
	scheduler := NewScheduler()
	// Schedule a one-off task to run after 3 seconds
	scheduler.ScheduleOnce("One-Off Task", 3*time.Second, func() {
		fmt.Println("This is a one-off task executing after 3 seconds.")
	})
	// Schedule a recurring task to run every 5 seconds, starting after an initial delay of 2 seconds
	scheduler.ScheduleInterval("Recurring Task", 2*time.Second, 5*time.Second, func() {
		fmt.Println("This is a recurring task executing every 5 seconds.")
	})
	// Automatically stop all tasks after 20 seconds
	// scheduler.AutoStopAll(20 * time.Second)
	// time.AfterFunc(5*time.Second, scheduler.StopAllTasks)
	scheduler.globalWg.Wait()
	fmt.Println("All scheduled tasks have completed. Exiting program.")
}
