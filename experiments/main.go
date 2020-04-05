package main

import (
	"context"
	"flag"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/scheduler/experiments/client"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	client := client.New()

	err := registerRunner(client.Client)
	if err != nil {
		glog.Fatal(err)
	}

}

func registerRunner(c pb.SchedulerClient) (err error) {
	runner := pb.Runner{
		Id:                 "6666",
		Status:             pb.RunnerStatus_IDLE,
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
	resp, err := c.HeartBeat(context.Background(), &pb.HeartBeatRequest{Runner: &runner})
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
