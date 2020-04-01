package server

import (
	"context"

	"github.com/lib/pq"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/scheduler/common"
	m "github.com/tradingAI/scheduler/server/model"
)

func (s *Servlet) RegisterRunner(ctx context.Context, req *pb.RegisterRunnerRequest) (resp *pb.RegisterRunnerResponse, err error) {
	runnerID := req.RunnerId

	if runnerID == "" {
		err = common.ErrEmptyRunnerID
		glog.Error(err)
		return
	}

	_, err = s.CreateRunner(runnerID)
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &pb.RegisterRunnerResponse{
		Ok: true,
	}

	return
}

func (s *Servlet) HeartBeat(ctx context.Context, req *pb.HeartBeatRequest) (resp *pb.HeartBeatResponse, err error) {
	runnerPb := req.GetRunner()
	if runnerPb == nil {
		err = common.ErrNilRunner
		glog.Error(err)
		return
	}

	runnerID := runnerPb.GetId()

	if runnerID == "" {
		err = common.ErrEmptyRunnerID
		glog.Error(err)
		return
	}

	var runner m.Runner
	err = s.DB.Where("runner_id = ?", runnerID).Find(&runner).Error
	if err != nil {
		glog.Error(err)
		return
	}

	runner.Status = int(runnerPb.Status)
	runner.JobsID = pq.Int64Array(uint64ArrToInt64Arr(runnerPb.JobsId))
	runner.CPUCoreNum = int(runnerPb.CpuCoreNum)
	runner.CPUUtilization = runnerPb.CpuUtilization
	runner.GPUNum = int(runnerPb.GpuNum)
	runner.GPUsIndex = pq.Int64Array(int32ArrToInt64Arr(runnerPb.GpusIndex))
	runner.GPUUtilization = runnerPb.GpuUtilization
	runner.CPUMemory = runnerPb.Memory
	runner.AvaliableCPUMemory = runnerPb.AvailableMemory
	runner.GPUMemory = runnerPb.GpuMemory
	runner.AvaliableGPUMemory = runnerPb.AvailableGpuMemory

	err = s.DB.Save(&runner).Error
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &pb.HeartBeatResponse{
		Ok: true,
	}

	return
}

func (s *Servlet) DestoryRunner(ctx context.Context, req *pb.DestoryRunnerRequest) (resp *pb.DestoryRunnerResponse, err error) {

	return
}
