package client

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	pg "github.com/tradingAI/go/db/postgres"
	m "github.com/tradingAI/go/db/postgres/model"
	mpb "github.com/tradingAI/proto/gen/go/model"
	spb "github.com/tradingAI/proto/gen/go/scheduler"
	"google.golang.org/grpc"
)

type Client struct {
	Client spb.SchedulerClient
	Conn   *grpc.ClientConn
	DB     *gorm.DB
}

func New() (client Client) {
	conn, err := grpc.Dial("localhost:8889", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		glog.Fatal(err)
	}

	client.Conn = conn
	client.Client = spb.NewSchedulerClient(conn)
	client.DB, err = pg.NewPostgreSQL(pg.DBConf{
		Database:     "tweb_db",
		Username:     "tweb_test",
		Password:     "tweb_test",
		Port:         5432,
		Host:         "localhost",
		Reset:        false,
		ReconnectSec: 3,
	})

	if err != nil {
		glog.Fatal(err)
	}

	return
}

func (c *Client) Free() (err error) {
	err = c.Conn.Close()
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func (c *Client) RegisterRunner(runnerID string) (err error) {
	runner := spb.Runner{
		Id:                 runnerID,
		Status:             spb.RunnerStatus_IDLE,
		CpuCoreNum:         4,
		CpuUtilization:     0.1,
		GpuNum:             0,
		GpuUtilization:     0,
		Memory:             2 * 1024 * 1024 * 1024,
		AvailableMemory:    1 * 1024 * 1024 * 1024,
		GpuMemory:          0,
		AvailableGpuMemory: 0,
		Token:              "admin",
	}
	resp, err := c.Client.HeartBeat(context.Background(), &spb.HeartBeatRequest{Runner: &runner})
	if err != nil {
		glog.Error(err)
		return
	}

	if resp.GetOk() {
		glog.Info("Register runner successfully")
	} else {
		glog.Fatalf("Register runner failed")
	}

	return
}

func (c *Client) CreateJob(runnerID string) (err error) {
	input := &spb.JobInput{
		Input: &spb.JobInput_EvalInput{
			EvalInput: &mpb.TbaseEvaluateInput{
				Start: "20120202",
				End:   "20120303",
			},
		},
	}

	jobInput, err := proto.Marshal(input)
	if err != nil {
		glog.Error(err)
		return
	}

	job := &m.Job{
		RunnerID: runnerID,
		Status:   int(spb.RunnerStatus_UNKNOWN),
		Input:    jobInput,
	}

	if err = c.DB.Create(job).Error; err != nil {
		glog.Error(err)
		return
	}

	resp, err := c.Client.CreateJob(context.Background(), &spb.CreateJobRequest{
		JobId:    int64(job.ID),
		Token:    "admin",
		MaxRetry: 0,
		Input:    input,
	})

	if err != nil {
		glog.Error(err)
		return
	}
	if resp.GetOk() {
		glog.Info("Create job successfully")
	} else {
		glog.Fatalf("Create job failed")
	}

	return
}
