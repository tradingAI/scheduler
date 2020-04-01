package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Job struct {
	gorm.Model
	RunnerID         string        `gorm:"column:runner_id; not null" json:"runner_id"`
	Type             int           `gorm:"column:type; not null" json:"type"`
	Status           int           `gorm:"column:status; not null" json:"status"`
	Retry            uint32        `gorm:"column:retry; not null" json:"retry"`
	MaxRetry         uint32        `gorm:"column:max_retry; not null" json:"max_retry"`
	CreateTimeUsec   int64         `gorm:"column:create_time_usec; not null" json:"create_time_usec"`
	FinishedTimeUsec int64         `gorm:"column:finished_time_usec" json:"finished_time_usec"`
	TotalSteps       uint32        `gorm:"column:total_steps; not null" json:"total_steps"`
	CurrentStep      uint32        `gorm:"column:current_step; not null" json:"current_step"`
	GPUsIndex        pq.Int64Array `gorm:"type:integer[]; column:gpus_index" json:"gpus_index"`
}

func (Job) TableName() string {
	return "jobs"
}
