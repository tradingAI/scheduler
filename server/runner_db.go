package server

import (
	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	m "github.com/tradingAI/scheduler/server/model"
)

func (s *Servlet) CreateRunner(runnerID string) (id int, err error) {
	runner := &m.Runner{
		RunnerID: runnerID,
		Status:   int(pb.RunnerStatus_UNKNOWN),
	}

	if err = s.DB.Create(runner).Error; err != nil {
		glog.Error(err)
		return
	}

	id = int(runner.ID)

	return
}
