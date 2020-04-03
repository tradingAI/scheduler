package server

import (
	"context"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/scheduler/common"
)

func (s *Servlet) CreateJob(ctx context.Context, req *pb.CreateJobRequest) (resp *pb.CreateJobResponse, err error) {
	if req == nil {
		err = common.ErrNilCreateJobRequest
		glog.Error(err)
		return
	}

	// Select idle runner
	runner, err := s.SelectIdleRunner()
	if err != nil {
		glog.Error(err)
		if err == gorm.ErrRecordNotFound {
			glog.Errorf("idle runner not found")
			return
		}
		return
	}

	// Assign job
	jobID := req.GetJobId()

	_, err = s.Redis.Do("LPUSH", runner.RunnerID, jobID)
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &pb.CreateJobResponse{
		Ok: true,
	}

	return
}

func (s *Servlet) StopJob(ctx context.Context, req *pb.StopJobRequest) (resp *pb.StopJobResponse, err error) {
	if req == nil {
		err = common.ErrNilStopJobRequest
		glog.Error(err)
		return
	}

	// TODO(mickey): stop job

	resp = &pb.StopJobResponse{
		Ok: true,
	}

	return
}

func (s *Servlet) RemoveJob(ctx context.Context, req *pb.RemoveJobRequest) (resp *pb.RemoveJobResponse, err error) {
	if req == nil {
		err = common.ErrNilRemoveJobRequest
		glog.Error(err)
		return
	}

	// TODO(mickey): stop job

	resp = &pb.RemoveJobResponse{
		Ok: true,
	}

	return
}
