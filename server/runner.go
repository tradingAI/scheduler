package server

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"

	"github.com/golang/glog"
	m "github.com/tradingAI/go/db/postgres/model"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/scheduler/common"
)

// HeartBeat update job and runner info stored in database.
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

	err = s.CheckTokenExisted(runnerPb.GetToken())
	if err != nil {
		glog.Error(err)
		return
	}

	// Update runner
	var runner m.Runner
	err = s.DB.Where("runner_id = ?", runnerID).Find(&runner).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			glog.Error(err)
			return
		}

		// register runner if not found
		glog.Infof("runner not found, register new runner %s", runnerID)

		newRunner, err := s.CreateRunner(runnerID)
		if err != nil {
			glog.Error(err)
			return resp, err
		}

		runner = *newRunner
	}

	runner.Status = int(runnerPb.Status)
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

	// Update jobs
	jobsPb := runnerPb.GetJobs()
	for _, jobPb := range jobsPb {
		var job m.Job

		err = s.DB.Where("id = ?", jobPb.Id).Find(&job).Error
		if err != nil {
			glog.Error(err)
			if gorm.IsRecordNotFoundError(err) {
				glog.Warningf("job [%d] is not existed", jobPb.Id)
				continue
			}
		}

		job.Status = int(jobPb.Status)
		job.Retry = jobPb.Retry
		job.FinishedTimeUsec = jobPb.FinishedTimeUsec
		job.TotalSteps = jobPb.TotalSteps
		job.CurrentStep = jobPb.CurrentStep
		job.GPUsIndex = int32ArrToInt64Arr(jobPb.GetGpusIndex())
	}

	return
}

func (s *Servlet) DestoryRunner(ctx context.Context, req *pb.DestoryRunnerRequest) (resp *pb.DestoryRunnerResponse, err error) {
	if req == nil {
		err = common.ErrNilDestoryRunnerRequest
		glog.Error(err)
		return
	}

	err = s.CheckTokenExisted(req.GetToken())
	if err != nil {
		glog.Error(err)
		return
	}

	runnerID := req.GetRunnerId()

	var runner m.Runner
	// check whether runner is existed
	err = s.DB.Where("runner_id = ?", runnerID).Find(&runner).Error
	if err != nil {
		glog.Error(err)
		return
	}

	_, err = s.Redis.Do("LPUSH", genRedisKey(DESTORY_RUNNER, runnerID), runnerID)
	if err != nil {
		glog.Error(err)
		return
	}

	err = s.DB.Delete(&runner).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
