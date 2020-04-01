package server

import (
	"context"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/scheduler/common"
)

func (s *Servlet) CreateJob(ctx context.Context, req *pb.CreateJobRequest) (resp *pb.CreateJobResponse, err error) {
	if req == nil {
		err = common.ErrNilCreateJobRequest
		glog.Error(err)
		return
	}

	// TODO(mickey): create job
	return
}

func (s *Servlet) StopJob(ctx context.Context, req *pb.StopJobRequest) (resp *pb.StopJobResponse, err error) {
	if req == nil {
		err = common.ErrNilStopJobRequest
		glog.Error(err)
		return
	}

	jobID := req.GetJobId()
	if jobID == "" {
		err = common.ErrEmptyJobID
		glog.Error(err)
		return
	}

	// TODO(mickey): stop job

	return
}
