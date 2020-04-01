package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Runner struct {
	gorm.Model
	RunnerID           string        `gorm:"column:runner_id; not null" json:"version"`
	Status             int           `gorm:"column:status; not null" json:"status"`
	JobsID             pq.Int64Array `gorm:"type:integer[]; column:jobs_id" json:"jobs_id"`
	CPUCoreNum         int           `gorm:"column:cpu_core_num; not null" json:"cpu_core_num"`
	CPUUtilization     float64       `gorm:"column:cpu_utilization; not null" json:"cpu_utilization"`
	GPUNum             int           `gorm:"column:gpu_num; not null" json:"gpu_num"`
	GPUsIndex          pq.Int64Array `gorm:"type:integer[]; column:gpus_index" json:"gpus_index"`
	GPUUtilization     float64       `gorm:"column:gpu_utilization; not null" json:"gpu_utilization"`
	CPUMemory          int64         `gorm:"column:cpu_memory; not null" json:"cpu_memory"`
	AvaliableCPUMemory int64         `gorm:"column:available_cpu_memory; not null" json:"available_cpu_memory"`
	GPUMemory          int64         `gorm:"column:gpu_memory; not null" json:"gpu_memory"`
	AvaliableGPUMemory int64         `gorm:"column:available_gpu_memory; not null" json:"available_gpu_memory"`
}

func (Runner) TableName() string {
	return "runners"
}
