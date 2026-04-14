package task

import (
	"testing"
	"time"
)

func TestClassifyTaskView(t *testing.T) {
	now := time.Date(2026, 4, 14, 9, 0, 0, 0, time.UTC)
	dueToday := now.Add(2 * time.Hour)
	dueTomorrow := now.Add(24 * time.Hour)

	todayTask := Task{Title: "today", DueAt: &dueToday, Status: "active"}
	upcomingTask := Task{Title: "upcoming", DueAt: &dueTomorrow, Status: "active"}

	if ViewForTask(todayTask, now) != "today" {
		t.Fatal("expected today task to map to today view")
	}
	if ViewForTask(upcomingTask, now) != "upcoming" {
		t.Fatal("expected future task to map to upcoming view")
	}
}
