package timeutils

import (
	"fmt"
	"time"
)

var taskTimes = make(map[string]time.Time)

func StartTimer(taskLabel string) {
	taskTimes[taskLabel] = time.Now()
	fmt.Println("Tracking execution time for: ", taskLabel)
}

func TimeElapsed(taskLabel string, completeTask bool) {
	startTime, exists := taskTimes[taskLabel]
	if !exists {
		fmt.Printf("Task %s has not been started\n", taskLabel)
		return
	}
	duration := time.Since(startTime)
	if !completeTask {
		fmt.Printf("%s has been running for %s\n", taskLabel, formatDuration(duration))
	} else {
		fmt.Printf("%s completed in %s\n", taskLabel, formatDuration(duration))
		// if the task is completed, we can remove it from our tracking time map
		delete(taskTimes, taskLabel)
	}
}

func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	milliseconds := int(d.Milliseconds()) % 1000
	microseconds := int(d.Microseconds()) % 1000
	nanoseconds := int(d.Nanoseconds()) % 1000

	if hours > 0 {
		return fmt.Sprintf("%d hours, %d minutes, %d seconds", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%d minutes, %d seconds", minutes, seconds)
	} else if seconds > 0 {
		return fmt.Sprintf("%d seconds, %d milliseconds", seconds, milliseconds)
	} else if milliseconds > 0 {
		return fmt.Sprintf("%d milliseconds, %d microseconds", milliseconds, microseconds)
	} else if microseconds > 0 {
		return fmt.Sprintf("%d microseconds, %d nanoseconds", microseconds, nanoseconds)
	} else {
		return fmt.Sprintf("%d nanoseconds", nanoseconds)
	}
}
