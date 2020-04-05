package server

import (
	"context"

	"github.com/golang/protobuf/proto"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	m "github.com/tradingAI/go/db/postgres/model"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/scheduler/common"
)

func (s *Servlet) CreateJob(ctx context.Context, req *pb.CreateJobRequest) (resp *pb.CreateJobResponse, err error) {
	if req == nil {
		err = common.ErrNilCreateJobRequest
		glog.Error(err)
		return
	}

	err = s.CheckTokenExisted(req.GetToken())
	if err != nil {
		glog.Error(err)
		return
	}

	// Select idle runner
	runner, err := s.SelectIdleRunner()
	if err != nil {
		glog.Error(err)
		if gorm.IsRecordNotFoundError(err) {
			glog.Errorf("idle runner not found")
			return
		}
		return
	}

	// Assign job
	jobID := req.GetJobId()

	var job m.Job
	err = s.DB.Where("id = ?", jobID).Find(&job).Error
	if err != nil {
		glog.Error(err)
		return
	}

	jobPb, err := convertJobModelToJobProto(job)
	if err != nil {
		glog.Error(err)
		return
	}

	jobBytes, err := proto.Marshal(jobPb)
	if err != nil {
		glog.Error(err)
		return
	}

	_, err = s.Redis.Do("LPUSH", genRedisKey(CREATE_JOB, runner.RunnerID), jobBytes)
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

	err = s.CheckTokenExisted(req.GetToken())
	if err != nil {
		glog.Error(err)
		return
	}

	// Assign job
	jobID := req.GetJobId()

	var job m.Job
	err = s.DB.Where("id = ?", jobID).Find(&job).Error
	if err != nil {
		glog.Error(err)
		return
	}

	_, err = s.Redis.Do("LPUSH", genRedisKey(STOP_JOB, job.RunnerID), jobID)
	if err != nil {
		glog.Error(err)
		return
	}

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

	err = s.CheckTokenExisted(req.GetToken())
	if err != nil {
		glog.Error(err)
		return
	}

	var job m.Job
	err = s.DB.Where("id = ?", req.GetJobId()).Find(&job).Error
	if err != nil {
		glog.Error(err)
		return
	}

	runnerID := job.RunnerID

	queue := CREATE_JOB
	if req.GetQueue() == pb.JobQueue_STOP {
		queue = STOP_JOB
	}

	_, err = s.Redis.Do("LREM", genRedisKey(queue, runnerID), job.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &pb.RemoveJobResponse{
		Ok: true,
	}

	return
}

func convertJobModelToJobProto(job m.Job) (jobPb *pb.Job, err error) {
	jobInput, err := convertJobInput(pb.JobType(job.Type), job.Input)
	if err != nil {
		glog.Error(err)
		return
	}

	jobPb = &pb.Job{
		Id:               uint64(job.ID),
		RunnerId:         job.RunnerID,
		Type:             pb.JobType(job.Type),
		Status:           pb.JobStatus(job.Status),
		Retry:            job.Retry,
		MaxRetry:         job.MaxRetry,
		CreateTimeUsec:   job.CreateTimeUsec,
		FinishedTimeUsec: job.FinishedTimeUsec,
		TotalSteps:       job.TotalSteps,
		CurrentStep:      job.CurrentStep,
		GpusIndex:        int64ArrToInt32Arr(job.GPUsIndex),
		Input:            jobInput,
	}
	return
}

func convertJobInput(jobType pb.JobType, input []byte) (jobInput *pb.JobInput, err error) {
	err = proto.Unmarshal(input, jobInput)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
