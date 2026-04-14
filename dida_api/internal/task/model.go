package task

import "time"

type Task struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `json:"title"`
	Notes       string     `json:"notes"`
	Status      string     `json:"status"`
	Priority    int        `json:"priority"`
	ListID      *uint      `json:"listId"`
	DueAt       *time.Time `json:"dueAt"`
	RemindAt    *time.Time `json:"remindAt"`
	IsAllDay    bool       `json:"isAllDay"`
	RepeatRule  string     `json:"repeatRule"`
	CompletedAt *time.Time `json:"completedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
