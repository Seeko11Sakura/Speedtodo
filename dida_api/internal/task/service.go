package task

import "time"

func ViewForTask(item Task, now time.Time) string {
	if item.DueAt == nil {
		return "inbox"
	}
	y1, m1, d1 := item.DueAt.In(now.Location()).Date()
	y2, m2, d2 := now.Date()
	if y1 == y2 && m1 == m2 && d1 == d2 {
		return "today"
	}
	return "upcoming"
}
