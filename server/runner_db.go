package server

import (
	"github.com/golang/glog"
	m "github.com/tradingAI/go/db/postgres/model"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

func (s *Servlet) CreateRunner(runnerID string) (runner *m.Runner, err error) {
	runner = &m.Runner{
		RunnerID: runnerID,
		Status:   int(pb.RunnerStatus_UNKNOWN),
	}

	if err = s.DB.Create(runner).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func (s *Servlet) SelectIdleRunner() (runner *m.Runner, err error) {
	runner = &m.Runner{}
	err = s.DB.Where("status = ?", pb.RunnerStatus_IDLE).First(&runner).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
