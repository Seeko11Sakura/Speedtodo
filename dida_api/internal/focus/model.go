package focus

import "time"

type Session struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	TaskID         *uint      `json:"taskId"`
	PlannedMinutes int        `json:"plannedMinutes"`
	StartedAt      time.Time  `json:"startedAt"`
	EndedAt        *time.Time `json:"endedAt"`
	Status         string     `json:"status"`
	CreatedAt      time.Time  `json:"createdAt"`
}

func (s *Session) Complete() {
	now := time.Now().UTC()
	s.Status = "completed"
	s.EndedAt = &now
}

func (s *Session) Cancel() {
	now := time.Now().UTC()
	s.Status = "cancelled"
	s.EndedAt = &now
}
